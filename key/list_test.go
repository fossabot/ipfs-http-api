package key

import (
	"encoding/json"
	"testing"
)

func TestList(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/key/list?", `{"Keys": []}`)

	reader, err := List(server.URL())
	if err != nil {
		t.Fatal("Error on List()", err.Error())
	}
	defer reader.Close()

	message := json.RawMessage{}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&message)
	if err != nil {
		t.Fatal("Error on decoder.Decode()", err.Error())
	}

	if string(message) != `{"Keys": []}` {
		t.Fatalf(`Expected body == '{"Keys": []}', Actual body == '%s'`, string(message))
	}
}
