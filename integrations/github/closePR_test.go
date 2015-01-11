package github

import (
	"net/http"
	"testing"
)

func TestNewUpdate(t *testing.T) {
	u := newUpdate()

	if u.State != "closed" {
		t.Error("Update state is not closed.")
	}
}

func TestAddHeaders(t *testing.T) {
	r, _ := http.NewRequest("GET", "foo", nil)
	addHeaders("foo", r)

	if r.Header["Authorization"][0] != "Basic Zm9v" {
		t.Error("Wrong authorization headers ", r.Header["Authorization"][0])
	}
}
