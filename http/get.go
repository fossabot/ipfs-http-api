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
		return nil, NewResponseError(res, "unexpected non 200 status code on %v: %v", getURL, res.StatusCode)
	}

	return res.Body, nil
}

// ResponseError wraps a go error in a struct that exposes
// the response
type ResponseError struct {
	message  string
	Response *http.Response
}

// NewResponseError returns a new ResponseError instance
func NewResponseError(response *http.Response, message string, args ...interface{}) *ResponseError {
	return &ResponseError{
		message:  fmt.Sprintf(message, args...),
		Response: response,
	}
}

func (e *ResponseError) Error() string {
	return e.message
}
