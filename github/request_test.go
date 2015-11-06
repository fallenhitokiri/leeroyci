package github

import (
	"net/http"
	"testing"
)

func TestAddHeaders(t *testing.T) {
	r, _ := http.NewRequest("GET", "foo", nil)
	addHeaders("foo", r)

	if r.Header["Authorization"][0] != "token foo" {
		t.Error("Wrong authorization headers ", r.Header["Authorization"][0])
	}
}
