package req

import (
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("https://httpbin.org/get")

	if err != nil {
		t.Error(err)
	}

	if !resp.Ok {
		t.Errorf("Got status code: %d", resp.StatusCode)
	}
}

func TestPost(t *testing.T) {
	resp, err := Post("https://httpbin.org/post")

	if err != nil {
		t.Error(err)
	}

	if !resp.Ok {
		t.Errorf("Got status code: %d", resp.StatusCode)
	}
}

func TestDelete(t *testing.T) {
	resp, err := Delete("https://httpbin.org/delete")

	if err != nil {
		t.Error(err)
	}

	if !resp.Ok {
		t.Errorf("Got status code: %d", resp.StatusCode)
	}
}

func TestPatch(t *testing.T) {
	resp, err := Patch("https://httpbin.org/patch")

	if err != nil {
		t.Error(err)
	}

	if !resp.Ok {
		t.Errorf("Got status code: %d", resp.StatusCode)
	}
}

func TestPut(t *testing.T) {
	resp, err := Put("https://httpbin.org/put")

	if err != nil {
		t.Error(err)
	}

	if !resp.Ok {
		t.Errorf("Got status code: %d", resp.StatusCode)
	}
}

func TestNew(t *testing.T) {
	method := "PUT"

	resp, err := New(method, "http://httpbin.org/anything")

	if err != nil {
		t.Error(err)
	}

	var m map[string]string
	resp.JSON(&m)

	if !(m["method"] == method) {
		t.Errorf("Method Mismatch\nExpected: %s\nGot: %s", method, m["method"])
	}
}

func TestToResponseFails(t *testing.T) {
	// This closes the body
	resp, err := Get("http://httpbin.org/get")

	if err != nil {
		t.Error(err)
	}

	// Fails because body has been consumed
	resp, err = toResponse(resp.Request, resp.Raw)
	if err == nil {
		t.Errorf("toResponse() = (%v, nil), want error", resp)
	}
}
