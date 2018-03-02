package pin

import "testing"

func TestAdd(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/pin/add?arg=foo-addr", "")

	err := Add(server.URL(), "foo-addr")
	if err != nil {
		t.Fatal("Error on Add()", err.Error())
	}

	requests := server.GetGETRequests("/api/v0/pin/add?arg=foo-addr")
	if len(requests) != 1 {
		t.Fatalf("Expected len(requests) == 1, Actual len(requests) == %v", len(requests))
	}
}

func TestAdd404(t *testing.T) {
	server.Reset()

	err := Add(server.URL(), "foo-addr")
	if err == nil {
		t.Fatal("Expected Add() to return an error, received nil")
	}
}
