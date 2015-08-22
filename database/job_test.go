package database

import (
	"testing"
)

func TestCGDoneJob(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false, false)

	job := CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	job.TasksDone()
	job.DeployDone()
	get := GetJob(job.ID)

	if job.TasksFinished == get.TasksFinished {
		t.Error("tasks not finished")
	}

	if job.DeployFinished == get.DeployFinished {
		t.Error("deploy not finished")
	}
}

func TestGetJobByCommit(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	j1 := GetJobByCommit("foo")
	j2 := GetJobByCommit("bar")

	if j1.ID != 0 {
		t.Error("j1 not nil", j1)
	}

	if j2.Branch != job.Branch {
		t.Error("j2 branches do not match", j2.Branch)
	}
}
