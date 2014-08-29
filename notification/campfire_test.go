package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestBuildCampfire(t *testing.T) {
	c := config.Config{
		CampfireId:   "c",
		CampfireRoom: "b",
		CampfireKey:  "a",
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

	_, err := buildCampfire(&c, &job)

	if err != nil {
		t.Error(err)
	}
}
