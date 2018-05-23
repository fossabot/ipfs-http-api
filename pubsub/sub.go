package pubsub

import (
	"net/url"
)

// Subscribe will construct a subscription instance and
// call connect on it.
func Subscribe(ipfsURL *url.URL, topic string) (*Subscription, error) {
	subscription := NewSubscription(ipfsURL, topic)

	go subscription.Start()

	return subscription, nil
}
