package http

import (
	"fmt"
	"io"
	"net/http"
)

// Get will make an http Get request and handle common errors
func Get(getURL string) (io.ReadCloser, error) {
	res, err := http.Get(getURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected non 200 status code on %v: %v", getURL, res.StatusCode)
	}

	return res.Body, nil
}
