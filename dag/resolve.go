package dag

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/computes/ipfs-http-api/http"
)

// Resolve resolves a ipld reference in IPFS
func Resolve(ipfsURL url.URL, address string) (string, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagResolveURL := ipfsURL
	dagResolveURL.Path = "/api/v0/dag/resolve"
	dagResolveURL.RawQuery = query.Encode()

	debug("Resolve %v", dagResolveURL.String())
	reader, err := http.Get(dagResolveURL.String())
	if err != nil {
		return "", err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	resolveResponse := struct {
		Cid struct {
			Address string `json:"/"`
		}
	}{}

	err = json.Unmarshal(data, &resolveResponse)
	if err != nil {
		return "", err
	}

	return resolveResponse.Cid.Address, nil
}
