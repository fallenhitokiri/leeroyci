package callbacks

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSplitURL(t *testing.T) {
	u := url.URL{
		Path: "/callback/github/foo",
	}
	req := http.Request{
		URL: &u,
	}

	s := getService(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}

	u.Path = "/callback/github/foo/"

	s = getService(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}
}
