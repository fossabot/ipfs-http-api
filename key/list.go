package key

import (
	"io"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// List will return a list of existing keys
func List(ipfsURL *url.URL) (io.ReadCloser, error) {
	query := url.Values{}

	keyListURL := *ipfsURL
	keyListURL.Path = "/api/v0/key/list"
	keyListURL.RawQuery = query.Encode()

	debug("Get %v", keyListURL.String())
	debugStack()
	return http.Get(keyListURL.String())
}
