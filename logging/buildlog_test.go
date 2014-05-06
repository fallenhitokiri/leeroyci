package logging

import (
	"testing"
)

func TestBuildlogAdd(t *testing.T) {
	log := Buildlog{}
	job := Job{
		URL:        "url",
		Branch:     "branch",
		Commit:     "commit",
		Command:    "command",
		Name:       "name",
		Email:      "email",
		Output:     "output",
		ReturnCode: nil,
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

	if j.Command != "command" {
		t.Error("wrong Command", j.Command)
	}

	if j.ReturnCode != nil {
		t.Error("wrong ReturnCode", j.ReturnCode)
	}

	if j.Output != "output" {
		t.Error("wrong Output", j.Output)
	}

	if j.Name != "name" {
		t.Error("wrong Name", j.Name)
	}

	if j.Email != "email" {
		t.Error("wrong Email", j.Email)
	}
}
