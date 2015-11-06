package database

import (
	"testing"
	"time"
)

func TestCGDoneJob(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)

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
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
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

func TestGetJobs(t *testing.T) {
	db.Exec("DELETE FROM jobs WHERE id > 0")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job1 := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	job2 := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	job3 := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	jobs := GetJobs(0, 1)

	if len(jobs) != 1 {
		t.Error("Wrong length", len(jobs))
	}

	if jobs[0].ID != job3.ID {
		t.Error("Wrong job", jobs[0].ID)
	}

	jobs = GetJobs(1, 2)

	if len(jobs) != 2 {
		t.Error("Wrong length", len(jobs))
	}

	if jobs[0].ID != job2.ID {
		t.Error("Wrong job", jobs[0].ID)
	}

	if jobs[1].ID != job1.ID {
		t.Error("Wrong job", jobs[1].ID)
	}
}

func TestNumberOfJobs(t *testing.T) {
	NewInMemoryDatabase()
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	count := NumberOfJobs()

	if count != 3 {
		t.Error("Wrong number of jobs", count)
	}
}

func TestJobPassed(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	if job.Passed() != true {
		t.Error("Job did not pass")
	}

	com, _ := CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)
	CreateCommandLog(com, job, "1", "foo")

	if job.Passed() != false {
		t.Error("Job did pass")
	}
}

func TestJobStatus(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	if job.Status() != JobStatusPending {
		t.Error("Job not pending")
	}

	job.TasksDone()

	if job.Status() != JobStatusSuccess {
		t.Error("Job not success")
	}

	com, _ := CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)
	CreateCommandLog(com, job, "1", "foo")

	if job.Status() != JobStatusError {
		t.Error("Job not error")
	}
}

func TestJobDeployDone(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	job.DeployDone()

	blank := time.Time{}
	if job.DeployFinished == blank {
		t.Error("Deploy time not set")
	}
}

func TestJobURL(t *testing.T) {
	db.Exec("DELETE FROM jobs WHERE id > 0")
	AddConfig("secret", "url", "cert", "key")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	if job.URL() != "url/1" {
		t.Error("Wrong URL", job.URL())
	}
}

func TestJobShouldBuild(t *testing.T) {
	AddConfig("secret", "url", "cert", "key")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	CreateCommand(repo, "name", "execute", "branch", CommandKindTest)

	if job.ShouldBuild() == true {
		t.Error("ShouldBuild = true")
	}

	CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)

	if job.ShouldBuild() == false {
		t.Error("ShouldBuild = false")
	}
}

func TestJobShouldDeploy(t *testing.T) {
	AddConfig("secret", "url", "cert", "key")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	CreateCommand(repo, "name", "execute", "branch", CommandKindTest)

	if job.ShouldDeploy() == true {
		t.Error("ShouldDeploy = true")
	}

	CreateCommand(repo, "name", "execute", "branch", CommandKindDeploy)

	if job.ShouldDeploy() == false {
		t.Error("ShouldDeploy = false")
	}
}

func TestJobStarted(t *testing.T) {
	AddConfig("secret", "url", "cert", "key")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	job.Started()

	if !job.TasksStarted.After(time.Time{}) {
		t.Error("No started time.")
	}
}

func TestJobIsRunningTasks(t *testing.T) {
	NewInMemoryDatabase()
	AddConfig("secret", "url", "cert", "key")
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false)
	job := CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	if job.IsRunning() {
		t.Error("Job is running, but should not - no start / end")
	}

	job.Started()

	if !job.IsRunning() {
		t.Error("Job is not running, but should - start")
	}

	job.TasksDone()

	if job.IsRunning() {
		t.Error("Job is running, but should not - start / end")
	}

	CreateCommand(repo, "name", "execute", "branch", CommandKindDeploy)

	if !job.IsRunning() {
		t.Error("Job should be running - deploy, but is not")
	}

	job.DeployDone()

	if job.IsRunning() {
		t.Error("Job is running, but should not")
	}
}
