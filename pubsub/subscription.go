package pubsub

import (
	"encoding/json"
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
}

// Close closes an open connection. This will return an error if
// the connection has already been closed.
func (s *Subscription) Close() error {
	err := s.response.Body.Close()
	if err != nil {
		return err
	}

	close(s.Messages)
	close(s.Errors)
	return nil
}

// Connect establishes an IPFS connection. This method will panic
// if it is called after Close because it'll try to write to closed
// channels
func (s *Subscription) Connect() error {
	ipfsURL := s.ipfsURL
	query := ipfsURL.Query()
	query.Add("arg", s.topic)

	ipfsURL.Path = "/api/v0/pubsub/sub"
	ipfsURL.RawQuery = query.Encode()

	response, err := http.Get(ipfsURL.String())
	if err != nil {
		return err
	}
	s.response = response

	go func() {
		decoder := json.NewDecoder(response.Body)
		for decoder.More() {
			ipfsMessage := struct {
				Data []byte `json:"data"`
			}{}
			err := decoder.Decode(&ipfsMessage)
			if err != nil {
				s.Errors <- err
				continue
			}

			if len(ipfsMessage.Data) == 0 {
				continue
			}

			s.Messages <- ipfsMessage.Data
		}
		s.Errors <- &DisconnectError{}
	}()

	return nil
}

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
