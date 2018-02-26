package pin

import (
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Add pins a an IPFS object recursively
func Add(ipfsURL url.URL, address string) error {
	query := url.Values{}
	query.Add("arg", address)

	pinAddURL := ipfsURL
	pinAddURL.Path = "/api/v0/pin/add"
	pinAddURL.RawQuery = query.Encode()

	debug("Add %v", pinAddURL.String())
	request, err := http.Get(pinAddURL.String())
	if err != nil {
		return err
	}
	defer request.Close()
	return nil
}
