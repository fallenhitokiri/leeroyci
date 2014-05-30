package logging

import (
	"testing"
	"time"
)

func TestBuildlogAdd(t *testing.T) {
	log := Buildlog{}

	task := Task{
		Command: "command",
		Output:  "output",
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

	if j.Tasks[0].Return != "" {
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

func TestBuildlogSorted(t *testing.T) {
	log := Buildlog{}
	j1 := Job{
		URL:       "url",
		Branch:    "branch",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	j2 := Job{
		URL:       "lur",
		Branch:    "branch",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	log.Add(j1)
	log.Add(j2)

	log.Sort()

	if log.Jobs[0].URL != j2.URL {
		t.Error("wrong URL", log.Jobs[0].URL)
	}

	if log.Jobs[1].URL != j1.URL {
		t.Error("wrong URL", log.Jobs[1].URL)
	}
}

func TestBuildlogJobsForRepo(t *testing.T) {
	log := Buildlog{}
	j1 := Job{
		URL:       "url",
		Branch:    "branch",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	j2 := Job{
		URL:       "lur",
		Branch:    "branch",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	log.Add(j1)
	log.Add(j2)

	j := log.JobsForRepo("url")

	if len(j) != 1 {
		t.Error("Wrong length", len(j))
	}

	if j[0].URL != "url" {
		t.Error("Wrong URL", j[0].URL)
	}
}

func TestBuildlogJobsForRepoBranch(t *testing.T) {
	log := Buildlog{}
	j1 := Job{
		URL:       "url",
		Branch:    "foo",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	j2 := Job{
		URL:       "url",
		Branch:    "branch",
		Commit:    "commit",
		Name:      "name",
		Email:     "email",
		Timestamp: time.Now(),
	}
	log.Add(j1)
	log.Add(j2)

	j := log.JobsForRepoBranch("url", "foo")

	if len(j) != 1 {
		t.Error("Wrong length", len(j))
	}

	if j[0].Branch != "foo" {
		t.Error("Wrong branch", j[0].Branch)
	}
}
