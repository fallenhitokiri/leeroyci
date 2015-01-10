package notification

import (
	"leeroy/config"
	"strings"
	"testing"
)

func TestAddHeaders(t *testing.T) {
	message := addHeaders("foo", "bar", "bla", "baz")

	if len(message) != 137 {
		t.Error("Message got the wrong length", len(message))
	}
}

func TestSubject(t *testing.T) {
	n := notification{
		Repo:   "repo",
		Branch: "branch",
		Name:   "name",
		Email:  "email",
		Status: true,
		Url:    "url",
		kind:   "build",
	}

	s := subject(&n)

	if s != "branch: success" {
		t.Error("Wrong subject", s)
	}

	n.Status = false
	s = subject(&n)

	if s != "branch: failed" {
		t.Error("Wrong subject", s)
	}
}

func TestBuildEmail(t *testing.T) {
	n := notification{
		Repo:     "repo",
		Branch:   "branch",
		Name:     "name",
		Email:    "email",
		Status:   true,
		Url:      "url",
		kind:     "build",
		rendered: "foo",
	}

	c := config.Config{
		EmailFrom: "foo@bar.tld",
	}

	m := buildEmail(&c, &n)

	if strings.Contains(string(m), "foo@bar.tld") == false {
		t.Error("Sender email not found")
	}
}
