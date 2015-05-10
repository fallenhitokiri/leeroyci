package notification

import (
	"leeroy/database"
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
		URL:    "url",
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
		Repo:    "repo",
		Branch:  "branch",
		Name:    "name",
		Email:   "email",
		Status:  true,
		URL:     "url",
		kind:    "build",
		message: "foo",
	}

	ms := database.GetMailServer()
	ms.Sender = "foo@bar.tld"

	m := buildEmail(&n)

	if strings.Contains(string(m), "foo@bar.tld") == false {
		t.Error("Sender email not found")
	}
}
