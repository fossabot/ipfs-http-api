package pubsub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Subscription is a stateful connection to IPFS
type Subscription struct {
	Errors   chan error
	Messages chan []byte

	ipfsURL  url.URL
	topic    string
	response *http.Response
	closed   bool
}

// Close closes an open connection. This will return an error if
// the connection has already been closed.
func (s *Subscription) Close() error {
	s.closed = true

	close(s.Messages)
	close(s.Errors)

	return s.response.Body.Close()
}

// Connect establishes an IPFS connection. This method will panic
// if it is called after Close because it'll try to write to closed
// channels
func (s *Subscription) Connect() error {
	ipfsURL := s.ipfsURL
	query := ipfsURL.Query()
	query.Add("arg", s.topic)
	query.Add("discover", "true")

	ipfsURL.Path = "/api/v0/id"
	idresponse, err := http.Get(ipfsURL.String())
	if err != nil {
		return err
	}
	idMessage := struct {
		ID string `json:"ID"`
	}{}
	byt, err := ioutil.ReadAll(idresponse.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byt, &idMessage)
	if err != nil {
		return err
	}

	ipfsURL.Path = "/api/v0/pubsub/sub"
	ipfsURL.RawQuery = query.Encode()

	debug("Subscribe %v", ipfsURL.String())
	response, err := http.Get(ipfsURL.String())
	if err != nil {
		return err
	}
	s.response = response

	go func() {
		decoder := json.NewDecoder(response.Body)
		for decoder.More() {
			ipfsMessage := struct {
				From string `json:"from"`
				Data []byte `json:"data"`
			}{}
			err := decoder.Decode(&ipfsMessage)
			if err != nil {
				s.emitError(err)
				continue
			}

			if len(ipfsMessage.Data) == 0 {
				continue
			}

			s.emitMessage(ipfsMessage.Data)
		}
		s.emitError(&DisconnectError{})
	}()

	return nil
}

func (s *Subscription) emitError(err error) {
	if s.closed {
		return
	}

	s.Errors <- err
}

func (s *Subscription) emitMessage(msg []byte) {
	if s.closed {
		return
	}

	s.Messages <- msg
}

// DisconnectError is returned when a pubsub sub connection
// is severed on the server side
type DisconnectError struct{}

func (e *DisconnectError) Error() string {
	return "Disconnected"
}

// IsDisconnectError tests to see if an error is a
// disconnect error. This can be used to implement
// reconnection logic. This error will also sent if
// the subscription.Close method was called.
func IsDisconnectError(err error) bool {
	_, ok := err.(*DisconnectError)
	return ok
}
