package ipfs

import (
	"io/ioutil"
	"testing"
)

func TestID(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/id?", `foo`)

	reader, err := ID(server.URL())
	if err != nil {
		t.Fatal("Error on ID()", err.Error())
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal("Error on ioutil.ReadAll()", err.Error())
	}

	if string(body) != "foo" {
		t.Fatalf(`Expected body == "foo", Actual body == "%s"`, body)
	}
}
