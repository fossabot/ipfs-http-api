package ipfs

import (
	"io"
	"net/url"

	"github.com/pkg/errors"

	"github.com/computes/ipfs-http-api/http"
)

// ID returns a reader of the IPFS node info
func ID(ipfsURL *url.URL) (io.ReadCloser, error) {
	idURL := *ipfsURL
	idURL.Path = "/api/v0/id"

	debug("ID %v", idURL.String())
	res, err := http.Get(idURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	return res, nil
}
