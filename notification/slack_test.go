package notification

import (
	"testing"
)

func TestBuildSlack(t *testing.T) {
	n := notification{
		Repo:   "repo",
		Branch: "branch",
		Name:   "name",
		Email:  "email",
		Status: true,
		URL:    "url",
		kind:   "build",
	}

	_, err := buildSlack(&n, "foo")

	if err != nil {
		t.Error(err)
	}
}
