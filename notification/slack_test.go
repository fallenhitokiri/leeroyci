package notification

import (
	"errors"
	"ironman/config"
	"ironman/logging"
	"testing"
)

func TestBuildSlack(t *testing.T) {
	c := config.Config{
		SlackChannel: "#devel",
	}

	task := logging.Task{
		Return: errors.New("0"),
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
