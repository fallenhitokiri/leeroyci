package notification

import (
	"ironman/config"
	"ironman/logging"
	"testing"
)

func TestBuildSlack(t *testing.T) {
	c := config.Config{
		SlackChannel: "#devel",
	}

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

	_, err := buildSlack(&c, &job)

	if err != nil {
		t.Error(err)
	}
}
