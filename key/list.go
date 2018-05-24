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
	res, err := http.Get(keyListURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	return res, nil
}
