package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestNotificationFromJob(t *testing.T) {
	c := config.Config{}
	j := logging.Job{
		Identifier: "ident",
		URL:        "url",
		Branch:     "branch",
		Commit:     "commit",
		CommitURL:  "curl",
		Name:       "name",
		Email:      "email",
	}

	n := notificationFromJob(&j, &c)

	if n.Status != true {
		t.Error("Status false, should be true")
	}
}

func TestRender(t *testing.T) {
	n := notification{
		Repo:   "repo",
		Branch: "branch",
		Name:   "name",
		Email:  "email",
		Status: true,
		URL:    "url",
		kind:   "build",
	}

	n.render()

	exp := "Repository: repo Branch: branch by name <email> -> Build success\nDetails: url"

	if n.message != exp {
		t.Error("Got ", n.message, "Expected ", exp)
	}
}
