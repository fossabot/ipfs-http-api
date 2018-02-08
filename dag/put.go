package dag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// Put uploads a file to IPFS as a DAG Object and returns
// the address
func Put(ipfsURL url.URL, reader io.Reader) (string, error) {
	var buffer bytes.Buffer

	dagPutURL := ipfsURL
	dagPutURL.Path = "/api/v0/dag/put"

	writer := multipart.NewWriter(&buffer)
	fileWriter, err := writer.CreateFormFile("file", "result.json")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	response, err := http.Post(dagPutURL.String(), writer.FormDataContentType(), &buffer)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code returned. Expected 200, received %v.\n%v", response.StatusCode, string(body))
	}

	dagPutResponse := struct {
		Cid struct {
			Value string `json:"/"`
		}
	}{}
	err = json.Unmarshal(body, &dagPutResponse)
	if err != nil {
		return "", err
	}

	debug("Put %v '%v'", dagPutURL.String(), dagPutResponse.Cid.Value)
	return dagPutResponse.Cid.Value, nil
}
