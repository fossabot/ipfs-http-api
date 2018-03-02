package dag

import "testing"

func TestResolve(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/dag/resolve?arg=foo-addr", `{"Cid": {"/": "foo-addr"}}`)

	addr, err := Resolve(server.URL(), "foo-addr")
	if err != nil {
		t.Fatal("Error on Resolve()", err.Error())
	}

	if addr != "foo-addr" {
		t.Fatalf(`Expected addr == "foo-addr", Actual addr == "%v"`, addr)
	}
}

func TestResolve404(t *testing.T) {
	server.Reset()

	addr, err := Resolve(server.URL(), "foo-addr")
	if err == nil {
		t.Fatal("Expected error on Resolve, got nil")
	}

	if addr != "" {
		t.Fatalf(`Expected addr == "", Actual addr == "%v"`, addr)
	}
}

func TestResolveInvalidJSON(t *testing.T) {
	server.Reset()
	server.SetGETResponseBody("/api/v0/dag/resolve?arg=foo-addr", `{"Cid"`)

	addr, err := Resolve(server.URL(), "foo-addr")
	if err == nil {
		t.Fatal("Expected error on Resolve, got nil")
	}

	if addr != "" {
		t.Fatalf(`Expected addr == "", Actual addr == "%v"`, addr)
	}
}
