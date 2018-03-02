package http

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/foo?", "bar")

	ipfsURL := server.URL()
	ipfsURL.Path = "/foo"

	reader, err := Get(ipfsURL.String())
	if err != nil {
		resErr, ok := err.(*ResponseError)
		if !ok {
			t.Fatal("Error on Get()", err.Error())
		}
		body, readErr := ioutil.ReadAll(resErr.Response.Body)
		t.Fatal("Error on Get()", err.Error(), string(body), readErr)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal("Error on ioutil.ReadAll()", err.Error())
	}

	if string(data) != "bar" {
		t.Fatalf(`Expected data == "bar", Actual data == "%s"`, data)
	}
}

func TestGet404(t *testing.T) {
	server.Reset()

	ipfsURL := server.URL()
	ipfsURL.Path = "/foo"

	_, err := Get(ipfsURL.String())
	if err == nil {
		t.Fatal("Expected Get() to return an error, received nil")
	}

	if !strings.Contains(err.Error(), "unexpected non 200") {
		t.Fatalf(`Expected "%v" to contain "unexpected non 200"`, err.Error())
	}
}

func TestGetNoServer(t *testing.T) {
	server.Reset()

	_, err := Get("http://notexist.example")
	if err == nil {
		t.Fatal("Expected Get() to return an error, received nil")
	}
}
