package callbacks

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSplitURL(t *testing.T) {
	u := url.URL{
		Path: "/callback/foo/github",
	}
	req := http.Request{
		URL: &u,
	}

	s, k := splitUrl(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}

	if k != "foo" {
		t.Error("wrong key", k)
	}

	u.Path = "/callback/foo/github/"

	s, _ = splitUrl(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}
}
