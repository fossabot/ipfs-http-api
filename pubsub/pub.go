package pubsub

import (
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Publish will publish the content to a given URL
func Publish(ipfsURL url.URL, topic, payload string) error {
	query := ipfsURL.Query()
	query.Add("arg", topic)
	query.Add("arg", payload)

	ipfsURL.Path = "/api/v0/pubsub/pub"
	ipfsURL.RawQuery = query.Encode()

	debug("Publish %v", ipfsURL.String())
	reader, err := http.Get(ipfsURL.String())
	if err != nil {
		return err
	}

	defer reader.Close()
	return nil
}
