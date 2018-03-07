package ipfs

import (
	"io"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// ID returns a reader of the IPFS node info
func ID(ipfsURL url.URL) (io.ReadCloser, error) {
	query := url.Values{}

	idURL := ipfsURL
	idURL.Path = "/api/v0/id"
	idURL.RawQuery = query.Encode()

	debug("ID %v", idURL.String())
	return http.Get(idURL.String())
}
