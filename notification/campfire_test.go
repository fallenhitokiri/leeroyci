package notification

import (
	"testing"
)

func TestBuildCampfire(t *testing.T) {
	n := notification{
		Repo:    "repo",
		Branch:  "branch",
		Name:    "name",
		Email:   "email",
		Status:  true,
		URL:     "url",
		kind:    "build",
		message: "foo",
	}

	_, err := buildCampfire(&n)

	if err != nil {
		t.Error(err)
	}
}

func TestEndpointCampfire(t *testing.T) {
	exp := "https://foo.campfirenow.com/room/bar/speak.json"
	e := endpointCampfire("foo", "bar")

	if exp != e {
		t.Error("Expected ", exp, " got ", e)
	}
}

func TestRequestCampfire(t *testing.T) {
	r := requestCampfire("foo", "bar", []byte("baz"))

	if r.Method != "POST" {
		t.Error("Wrong method ", r.Method)
	}

	u, _, _ := r.BasicAuth()

	if u != "bar" {
		t.Error("Wrong username", u)
	}
}
