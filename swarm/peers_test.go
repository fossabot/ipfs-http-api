package swarm

import (
	"io/ioutil"
	"testing"
)

func TestPeers(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/swarm/peers?", "foo")

	reader, err := Peers(server.URL())
	if err != nil {
		t.Fatal("Error on Peers()", err.Error())
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
