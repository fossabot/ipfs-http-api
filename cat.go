package ipfs

import (
	"io"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Cat returns a reader for the data in IPFS located at address
func Cat(ipfsURL *url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	catURL := *ipfsURL
	catURL.Path = "/api/v0/cat"
	catURL.RawQuery = query.Encode()

	debug("Cat %v", catURL.String())
	return http.Get(catURL.String())
}
