package callbacks

import (
	"net/http"
	"net/url"
	"testing"
)

func TestService(t *testing.T) {
	u := url.URL{
		Path: "/callback/github",
	}
	req := http.Request{
		URL: &u,
	}

	s := service(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}

	u.Path = "/callback/github/"

	s = service(&req)

	if s != "github" {
		t.Error("service returned a wrong string", s)
	}
}
