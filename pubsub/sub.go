package pubsub

import (
	"net/url"
)

// Subscribe will subscribe to a given topic and returns
// a channel for messages and a channel for errors
func Subscribe(ipfsURL url.URL, topic string) (*Subscription, error) {
	messages := make(chan []byte)
	errors := make(chan error)

	subscription := &Subscription{
		Errors:   errors,
		Messages: messages,
		ipfsURL:  ipfsURL,
		topic:    topic,
	}

	err := subscription.Connect()
	return subscription, err
}
