package pubsub

import (
	"io"
	"io/ioutil"
	"net/url"

	"github.com/pkg/errors"

	"github.com/computes/ipfs-http-api/http"
)

// Publish will publish the content to a given URL
func Publish(ipfsURL *url.URL, topic, payload string) error {
	query := url.Values{}
	query.Add("arg", topic)
	query.Add("arg", payload)

	pubURL := *ipfsURL
	pubURL.Path = "/api/v0/pubsub/pub"
	pubURL.RawQuery = query.Encode()

	debug("Publish %v", pubURL.String())
	reader, err := http.Get(pubURL.String())
	if err != nil {
		return errors.Wrap(err, "http.Get failed")
	}
	io.Copy(ioutil.Discard, reader)
	defer reader.Close()
	return nil
}
