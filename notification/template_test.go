package notification

import (
	"strings"
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestGetTemplate(t *testing.T) {
	_, err := getTemplate("email", EventTest, "text")

	if err != nil {
		t.Error(err)
	}

	_, err = getTemplate("foo", EventTest, "email-asdf")

	if err == nil {
		t.Error("No error")
	}
}

func TestMessage(t *testing.T) {
	j := database.Job{
		Branch:    "foo",
		Commit:    "bar",
		CommitURL: "url",
		Name:      "baz",
		Email:     "foo@bar.tld",
	}

	tmpl := message(&j, "email", EventTest, "text")

	if !strings.Contains(tmpl, "branch: foo") {
		t.Error("branch name not found", tmpl)
	}
}
