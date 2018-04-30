package pin

import (
	"io"
	"io/ioutil"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Add pins a an IPFS object recursively
func Add(ipfsURL *url.URL, address string) error {
	query := url.Values{}
	query.Add("arg", address)

	pinAddURL := *ipfsURL
	pinAddURL.Path = "/api/v0/pin/add"
	pinAddURL.RawQuery = query.Encode()

	debug("Add %v", pinAddURL.String())
	resp, err := http.Get(pinAddURL.String())
	if err != nil {
		return err
	}
	io.Copy(ioutil.Discard, resp)
	defer resp.Close()
	return nil
}
