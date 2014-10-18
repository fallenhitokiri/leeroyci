package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestBuildHipChat(t *testing.T) {
	c := config.Config{}

	task := logging.Task{
		Return: "",
	}
	job := logging.Job{
		URL:    "foo",
		Branch: "bar",
		Name:   "baz",
		Email:  "zab",
		Tasks:  []logging.Task{task},
	}

	payload := buildHipChat(&c, &job, "foo")

	if payload.Color != "green" {
		t.Error("Wrong color", payload.Color)
	}
}
