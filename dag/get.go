package dag

import (
	"io"
	"io/ioutil"
	"net/url"
	"runtime"

	"github.com/computes/ipfs-http-api/http"
)

// Get retrieves a dag object from IPFS
func Get(ipfsURL url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagGetURL := ipfsURL
	dagGetURL.Path = "/api/v0/dag/get"
	dagGetURL.RawQuery = query.Encode()

	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	debug("Get %v: %v", dagGetURL.String(), string(buf))
	return http.Get(dagGetURL.String())
}

// GetBytes retrieves a dag object from IPFS and reads the whole buffer
// into memory
func GetBytes(ipfsURL url.URL, address string) ([]byte, error) {
	reader, err := Get(ipfsURL, address)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
