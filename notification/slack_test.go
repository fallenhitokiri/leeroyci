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

	job := logging.Job{
		URL:        "foo",
		Branch:     "bar",
		Name:       "baz",
		Email:      "zab",
		ReturnCode: errors.New("0"),
	}

	_, err := buildSlack(&c, &job)

	if err != nil {
		t.Error(err)
	}
}
