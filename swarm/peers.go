package swarm

import (
	"io"
	"net/url"

	"github.com/pkg/errors"

	"github.com/computes/ipfs-http-api/http"
)

// Peers list peers with open connections
func Peers(ipfsURL *url.URL) (io.ReadCloser, error) {
	query := url.Values{}

	swarmPeersURL := *ipfsURL
	swarmPeersURL.Path = "/api/v0/swarm/peers"
	swarmPeersURL.RawQuery = query.Encode()

	debug("Peers %v", swarmPeersURL.String())
	res, err := http.Get(swarmPeersURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	return res, nil
}
