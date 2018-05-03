package key

import (
	"encoding/json"
	"testing"
)

func TestGen(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/key/gen?arg=foo&type=ed25519", `{"Name": "foo", "Id": "id"}`)

	reader, err := Gen(server.URL(), "foo")
	if err != nil {
		t.Fatal("Error on Gen()", err.Error())
	}
	defer reader.Close()

	message := json.RawMessage{}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&message)
	if err != nil {
		t.Fatal("Error on decoder.Decode()", err.Error())
	}

	if string(message) != `{"Name": "foo", "Id": "id"}` {
		t.Fatalf(`Expected body == '{"Name": "foo", "Id": "id"}', Actual body == '%s'`, string(message))
	}
}
