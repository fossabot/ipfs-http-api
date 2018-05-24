package key

import (
	"io"
	"net/url"

	"github.com/pkg/errors"

	"github.com/computes/ipfs-http-api/http"
)

// Gen will create a new IPFS key
func Gen(ipfsURL *url.URL, name string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", name)
	query.Add("type", "ed25519")

	keyGenURL := *ipfsURL
	keyGenURL.Path = "/api/v0/key/gen"
	keyGenURL.RawQuery = query.Encode()

	debug("Get %v", keyGenURL.String())
	res, err := http.Get(keyGenURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	return res, nil
}
