package pin

import (
	"log"
	"testing"

	TESTSERVER "github.com/computes/go-test-server"
)

var server TESTSERVER.Server

func TestMain(m *testing.M) {
	server = TESTSERVER.New()
	err := server.Open()
	if err != nil {
		log.Fatalln("Error on server.Open()", err.Error())
	}
	defer server.Close()

	m.Run()
}
