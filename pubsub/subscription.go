package pubsub

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
)

// Handler is the subscription handler interface
type Handler interface {
	Recv(*SubscriptionMessage)
}

// Subscription is a stateful connection to IPFS
type Subscription struct {
	Wait    *sync.WaitGroup
	handler Handler
	ipfsURL *url.URL
	topic   string
	closed  bool
}

// NewSubscription constructs a new subscription
func NewSubscription(ipfsURL *url.URL, topic string) *Subscription {
	wait := &sync.WaitGroup{}
	wait.Add(1)
	return &Subscription{
		Wait:    wait,
		ipfsURL: ipfsURL,
		topic:   topic,
		closed:  false,
	}
}

// Handle will register a message handler
func (s *Subscription) Handle(h Handler) {
	s.handler = h
}

// Close closes an open connection. This will return an error if
// the connection has already been closed.
func (s *Subscription) Close() error {
	s.closed = true
	return nil
}

// Start establishes an IPFS connection and passes messages
// to the handlers
func (s *Subscription) Start() error {
	query := url.Values{}
	query.Add("arg", s.topic)
	query.Add("discover", "true")

	subURL := *s.ipfsURL
	subURL.Path = "/api/v0/pubsub/sub"
	subURL.RawQuery = query.Encode()

	debug("Subscribe %v", subURL.String())
	response, err := http.Get(subURL.String())
	s.Wait.Done()
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(response.Body)
	defer response.Body.Close()
	for decoder.More() {
		if s.closed {
			debug("Closed by client")
			return nil
		}

		var msg SubscriptionMessage
		err := decoder.Decode(&msg)
		if err != nil {
			debugError(err, "Decode Error")
			continue
		}

		if s.handler != nil {
			go s.handler.Recv(&msg)
		}
	}

	debug("Closed by server")
	return &DisconnectError{}
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
