package ipfs

import (
	"io/ioutil"
	"testing"
)

func TestCat(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/cat?arg=foo-addr", `foo`)

	reader, err := Cat(server.URL(), "foo-addr")
	if err != nil {
		t.Fatal("Error on Cat()", err.Error())
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
