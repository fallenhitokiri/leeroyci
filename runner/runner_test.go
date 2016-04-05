package runner

import (
	"testing"
	"time"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestQueueJobEnqueue(t *testing.T) {
	qj := QueueJob{
		JobID: 1,
	}

	qj.Enqueue()

	got := <-runQueue

	if got.JobID != qj.JobID {
		t.Error("Got wrong job ID", got.JobID)
	}
}

func TestRun(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	cmd, _ := database.CreateCommand(repo, "foo", "", "", database.CommandKindBuild)

	run(job, repo, database.CommandKindBuild, 1)

	logs := database.GetCommandLogsForJob(job.ID)

	if len(logs) != 1 {
		t.Error("Wrong logs length", len(logs))
	}

	if logs[0].Name != cmd.Name {
		t.Error("Wrong name", logs[0].Name)
	}
}

func TestRunWrongKind(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	database.CreateCommand(repo, "foo", "", "", database.CommandKindBuild)

	run(job, repo, database.CommandKindTest, 1)

	logs := database.GetCommandLogsForJob(job.ID)

	if len(logs) != 0 {
		t.Error("Wrong logs length", len(logs))
	}
}

func TestRunWrongBranch(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	database.CreateCommand(repo, "foo", "", "bar", database.CommandKindBuild)

	run(job, repo, database.CommandKindBuild, 1)

	logs := database.GetCommandLogsForJob(job.ID)

	if len(logs) != 0 {
		t.Error("Wrong logs length", len(logs))
	}
}

func TestRunMatchingBranch(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	cmd, _ := database.CreateCommand(repo, "foo", "", "branch", database.CommandKindBuild)

	run(job, repo, database.CommandKindBuild, 1)

	logs := database.GetCommandLogsForJob(job.ID)

	if len(logs) != 1 {
		t.Error("Wrong logs length", len(logs))
	}

	if logs[0].Name != cmd.Name {
		t.Error("Wrong name", logs[0].Name)
	}
}

func TestDeploy(t *testing.T) {
	database.NewInMemoryDatabase()

	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")

	deploy(job, repo, 1)

	job = database.GetJob(job.ID)

	blank := time.Time{}
	if job.DeployFinished == blank {
		t.Error("Deploy not finished - time not set.")
	}
}

func TestHandleJob(t *testing.T) {
	database.NewInMemoryDatabase()

	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	status := make(chan bool, 5)
	qJob := QueueJob{
		JobID:  job.ID,
		Status: status,
	}
	blank := time.Time{}

	handleJob(job, repo, &qJob, 1)

	if job.TasksStarted == blank {
		t.Error("Tasks not started")
	}

	if job.TasksFinished == blank {
		t.Error("Tasks not finished")
	}
}
