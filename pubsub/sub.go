package pubsub

import (
	"net/url"
)

// Subscribe will construct a subscription instance and
// call connect on it.
func Subscribe(ipfsURL url.URL, topic string) (*Subscription, error) {
	subscription := NewSubscription(ipfsURL, topic)

	err := subscription.Connect()
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
