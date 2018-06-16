package dag

import (
	"encoding/json"
	"testing"
)

func TestGet(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/dag/get?arg=foo-addr", `"foo"`)

	reader, err := Get(server.URL(), "foo-addr")
	if err != nil {
		t.Fatal("Error on Cat()", err.Error())
	}
	defer reader.Close()

	message := json.RawMessage{}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&message)
	if err != nil {
		t.Fatal("Error on decoder.Decode()", err.Error())
	}

	if string(message) != `"foo"` {
		t.Fatalf(`Expected body == '"foo"', Actual body == '%s'`, string(message))
	}
}

func TestGetBytes(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/dag/get?arg=foo-addr", `"foo"`)

	message, err := GetBytes(server.URL(), "foo-addr")
	if err != nil {
		t.Fatal("Error on Cat()", err.Error())
	}

	if string(message) != `"foo"` {
		t.Fatalf(`Expected body == '"foo"', Actual body == '%s'`, string(message))
	}
}

func TestGetInterface(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/dag/get?arg=foo-addr", `"foo"`)

	var message string
	err := GetInterface(server.URL(), "foo-addr", &message)
	if err != nil {
		t.Fatal("Error on GetInterface()", err.Error())
	}

	if string(message) != "foo" {
		t.Fatalf(`Expected body == 'foo', Actual body == '%s'`, string(message))
	}
}
