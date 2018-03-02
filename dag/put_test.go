package dag

import (
	"bytes"
	"fmt"
	"net/url"
	"testing"
)

type ErrorReader struct{}

func (e *ErrorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("I only error")
}

func TestPut(t *testing.T) {
	server.Reset()
	server.SetPOSTResponseBody(`/api/v0/dag/put? "foo"`, `{"Cid": {"/": "foo-addr"}}`)

	addr, err := Put(server.URL(), bytes.NewBuffer([]byte(`"foo"`)))
	if err != nil {
		t.Fatal("Error on Put()", err.Error())
	}

	if addr != "foo-addr" {
		t.Fatalf(`Expected addr == "foo-addr", Actual addr == "%s"`, addr)
	}
}

func TestPutInvalidJSONResponse(t *testing.T) {
	server.Reset()
	server.SetPOSTResponseBody(`/api/v0/dag/put? "foo"`, `{"Ci`)

	addr, err := Put(server.URL(), bytes.NewBuffer([]byte(`"foo"`)))
	if err == nil {
		t.Fatal("Expected err on Put(), received nil")
	}

	if addr != "" {
		t.Fatalf(`Expected empty addr, Actual addr == "%s"`, addr)
	}
}

func TestPutInvalidReader(t *testing.T) {
	server.Reset()

	addr, err := Put(server.URL(), &ErrorReader{})
	if err == nil {
		t.Fatal("Expected err on Put(), received nil")
	}

	if addr != "" {
		t.Fatalf(`Expected empty addr, Actual addr == "%s"`, addr)
	}
}

func TestPutNoServer(t *testing.T) {
	server.Reset()

	ipfsURL, err := url.Parse("http://notexist.example")
	if err != nil {
		t.Fatal("Error on url.Parse", err.Error())
	}

	addr, err := Put(*ipfsURL, bytes.NewBuffer([]byte(`"foo"`)))
	if err == nil {
		t.Fatal("Expected err on Put(), received nil")
	}

	if addr != "" {
		t.Fatalf(`Expected empty addr, Actual addr == "%s"`, addr)
	}
}

func TestPut404(t *testing.T) {
	server.Reset()

	addr, err := Put(server.URL(), bytes.NewBuffer([]byte(`"foo"`)))
	if err == nil {
		t.Fatal("Expected err on Put(), received nil")
	}

	if addr != "" {
		t.Fatalf(`Expected empty addr, Actual addr == "%s"`, addr)
	}
}
