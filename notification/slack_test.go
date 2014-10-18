package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestBuildSlack(t *testing.T) {
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

	_, err := buildSlack(&c, &job, "foo")

	if err != nil {
		t.Error(err)
	}
}
