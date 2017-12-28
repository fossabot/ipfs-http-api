package pubsub

import (
	"encoding/json"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// IPFSMessage defines the structure of IPFS pubsub messages
type IPFSMessage struct {
	Data []byte `json:"data"`
}

// Subscribe will subscribe to a given topic and returns
// a channel for messages and a channel for errors
func Subscribe(ipfsURL url.URL, topic string) (<-chan []byte, chan error, error) {
	messages := make(chan []byte)
	errors := make(chan error)
	query := ipfsURL.Query()
	query.Add("arg", topic)

	ipfsURL.Path = "/api/v0/pubsub/sub"
	ipfsURL.RawQuery = query.Encode()

	body, err := http.Get(ipfsURL.String())
	if err != nil {
		return messages, errors, err
	}

	go func() {
		decoder := json.NewDecoder(body)
		for decoder.More() {
			ipfsMessage := IPFSMessage{}
			err := decoder.Decode(&ipfsMessage)
			if err != nil {
				errors <- err
				continue
			}

			if len(ipfsMessage.Data) == 0 {
				continue
			}

			messages <- ipfsMessage.Data
		}
	}()

	return messages, errors, nil
}
