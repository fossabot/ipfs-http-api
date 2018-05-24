package dag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
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
		return "", errors.Wrap(err, "writer.CreateFormFile failed")
	}

	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		return "", errors.Wrap(err, "io.Copy failed")
	}

	err = writer.Close()
	if err != nil {
		return "", errors.Wrap(err, "Close failed")
	}

	debug("Put %v %s", dagPutURL.String(), buffer.Bytes())
	response, err := http.Post(dagPutURL.String(), writer.FormDataContentType(), &buffer)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", errors.Wrap(err, "Post failed")
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
		return "", errors.Wrap(err, "json Decode failed")
	}

	debug("Put %v '%v'", dagPutURL.String(), dagPutResponse.Cid.Value)
	return dagPutResponse.Cid.Value, nil
}

// PutInterface will take an interface and convert it to a buffer
// using json.Marshal
func PutInterface(ipfsURL *url.URL, data interface{}) (string, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "json.Marshal failed")
	}

	addr, err := Put(ipfsURL, bytes.NewBuffer(buf))
	if err != nil {
		return "", errors.Wrap(err, "Put failed")
	}

	return addr, nil
}
