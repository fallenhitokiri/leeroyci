package logging

import (
	"testing"
)

func TestBuildlogAdd(t *testing.T) {
	log := Buildlog{}
	log.Done = make(chan bool, 5)

	task := Task{
		Command: "command",
		Output:  "output",
		Return:  nil,
	}
	job := Job{
		URL:    "url",
		Branch: "branch",
		Commit: "commit",
		Name:   "name",
		Email:  "email",
		Tasks:  []Task{task},
	}
	log.Add(job)

	if len(log.Jobs) != 1 {
		t.Error("build not added")
	}

	j := log.Jobs[0]

	if j.URL != "url" {
		t.Error("wrong URL", j.URL)
	}

	if j.Branch != "branch" {
		t.Error("wrong Branch", j.Branch)
	}

	if j.Commit != "commit" {
		t.Error("wrong Commit", j.Commit)
	}

	if j.Tasks[0].Command != "command" {
		t.Error("wrong Command", j.Tasks[0].Command)
	}

	if j.Tasks[0].Return != nil {
		t.Error("wrong ReturnCode", j.Tasks[0].Return)
	}

	if j.Tasks[0].Output != "output" {
		t.Error("wrong Output", j.Tasks[0].Output)
	}

	if j.Name != "name" {
		t.Error("wrong Name", j.Name)
	}

	if j.Email != "email" {
		t.Error("wrong Email", j.Email)
	}
}
