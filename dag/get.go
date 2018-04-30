package dag

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Get retrieves a dag object from IPFS
func Get(ipfsURL *url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagGetURL := ipfsURL
	dagGetURL.Path = "/api/v0/dag/get"
	dagGetURL.RawQuery = query.Encode()

	debug("Get %v", dagGetURL.String())
	debugStack()
	return http.Get(dagGetURL.String())
}

// GetBytes retrieves a dag object from IPFS and reads the whole buffer
// into memory
func GetBytes(ipfsURL *url.URL, address string) ([]byte, error) {
	reader, err := Get(ipfsURL, address)
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		return nil, err
	}

	message := json.RawMessage{}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&message)
	if err != nil {
		return nil, err
	}

	return []byte(message), nil
}
