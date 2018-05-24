package dag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// Put uploads a file to IPFS as a DAG Object and returns
// the address
func Put(ipfsURL *url.URL, reader io.Reader) (string, error) {
	var buffer bytes.Buffer

	dagPutURL := *ipfsURL
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

	debug("Put %v %s", dagPutURL.String(), buffer.Bytes())
	response, err := http.Post(dagPutURL.String(), writer.FormDataContentType(), &buffer)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code returned. Expected 200, received %v", response.StatusCode)
	}

	dagPutResponse := struct {
		Cid struct {
			Value string `json:"/"`
		}
	}{}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&dagPutResponse)
	if err != nil {
		return "", err
	}

	debug("Put %v '%v'", dagPutURL.String(), dagPutResponse.Cid.Value)
	return dagPutResponse.Cid.Value, nil
}

// PutInterface will take an interface and convert it to a buffer
// using json.Marshal
func PutInterface(ipfsURL *url.URL, data interface{}) (string, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return Put(ipfsURL, bytes.NewBuffer(buf))
}
