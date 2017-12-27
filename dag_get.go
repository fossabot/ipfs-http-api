package ipfs

import (
	"io"
	"io/ioutil"
	"net/url"
)

// DagGet retrieves a dag object from IPFS
func DagGet(ipfsURL url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagGetURL := ipfsURL
	dagGetURL.Path = "/api/v0/dag/get"
	dagGetURL.RawQuery = query.Encode()

	return httpGet(dagGetURL.String())
}

// DagGetBytes retrieves a dag object from IPFS and reads the whole buffer
// into memory
func DagGetBytes(ipfsURL url.URL, address string) ([]byte, error) {
	reader, err := DagGet(ipfsURL, address)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(reader)
}
