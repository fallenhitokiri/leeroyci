package notification

import (
	"strings"
	"testing"
)

func TestBuildHipChat(t *testing.T) {
	n := notification{
		Repo:   "repo",
		Branch: "branch",
		Name:   "name",
		Email:  "email",
		Status: false,
		Url:    "url",
		kind:   "build",
	}

	p := notToHipChapt(&n, "foo")

	if p.Room != "foo" {
		t.Error("Wrong room", p.Room)
	}
}

func TestToURLEncoded(t *testing.T) {
	h := hipchatPayload{
		Room:    "foo",
		From:    "bar",
		Message: "baz",
		Notify:  true,
		Format:  "text",
		Status:  true,
	}

	e := string(h.toURLEncoded())

	if strings.Contains(e, "notify=1") == false {
		t.Error("Wrong notification settings")
	}

	if strings.Contains(e, "color=green") == false {
		t.Error("Wrong notification color")
	}

	h.Status = false
	h.Notify = false

	e = string(h.toURLEncoded())

	if strings.Contains(e, "notify=2") == false {
		t.Error("Wrong notification settings")
	}

	if strings.Contains(e, "color=red") == false {
		t.Error("Wrong notification color")
	}
}

func TestEndpointHipChat(t *testing.T) {
	exp := "https://www.hipchat.com/v1/rooms/message?auth_token=foo"
	e := endpointHipChat("foo")

	if e != exp {
		t.Error("Wrong API endpoint ", e)
	}
}
