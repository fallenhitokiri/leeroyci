// Package runner runs all tasks for all commands associated with a repository.
package runner

import (
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/notification"
)

// QueueJob represents a job put on the runner queue. The status channel is
// used to notify about finished builds.
type QueueJob struct {
	JobID  int64
	Status chan bool
}

// Enqueue a job for running.
func (q *QueueJob) Enqueue() {
	runQueue <- q
}

// RunQueue receives job IDs for which commands should run.
var runQueue = make(chan *QueueJob, 100)

// Runner waits for jobs to be pushed on RunQueue and runs all commands. It also
// creates the command logs and sends the necessary notifications.
func Runner() {
	newTaskManager()

	for {
		queueJob := <-runQueue

		job := database.GetJob(queueJob.JobID)
		repository, err := database.GetRepositoryByID(job.RepositoryID)

		if err != nil {
			log.Println("Could not find repository for", job.Repository.URL)
			continue
		}

		if job.Cancelled == true {
			log.Println("Job cancelled, not running commands", job.ID)
			continue
		}

		taskID := manager.getTaskID(repository.Name, job.Branch)

		if taskID != 0 {
			go handleJob(job, repository, queueJob, taskID)
		} else {
			runQueue <- queueJob
			// if we cannot handle any more jobs we wait for 10 seconds before
			// checking again.
			time.Sleep(10 * time.Second)
		}
	}
}

// handleJob runs all tasks for a job and updates the queueJob status field once
// builds and tests are done.
func handleJob(job *database.Job, repository *database.Repository,
	queueJob *QueueJob, taskID int) {

	job.Started()

	run(job, repository, database.CommandKindTest, taskID)
	notification.Notify(job, notification.EventTest)

	if job.Passed() && job.ShouldBuild() {
		run(job, repository, database.CommandKindBuild, taskID)
		notification.Notify(job, notification.EventBuild)
	}

	job.TasksDone()

	if queueJob.Status != nil {
		queueJob.Status <- true
	}

	manager.doneWithID(taskID)

	if job.Passed() && job.ShouldDeploy() {
		go deploy(job, repository, taskID)
	}
}

// deploy is a wrapper around the run commnad to make running the deploy commands
// in a separate go routine more convenient.
func deploy(job *database.Job, repository *database.Repository, taskID int) {
	notification.Notify(job, notification.EventDeployStart)
	run(job, repository, database.CommandKindDeploy, taskID)
	job.DeployDone()
	notification.Notify(job, notification.EventDeployEnd)
}

// run runs the command that is specified in Command.Execute and creates the
// command log with the results of the command.
func run(job *database.Job, repository *database.Repository, kind string, taskID int) {
	commands := repository.GetCommands(job.Branch, kind)

	for _, command := range commands {
		if command.Kind == kind {
			if (command.Branch != "" && command.Branch == job.Branch) || command.Branch == "" {
				repository := job.Repository.Name
				branch := job.Branch

				log.Println("Running", command.Name, "for", repository, branch)

				cmd := exec.Command(command.Execute, repository, branch, strconv.Itoa(taskID))
				out, err := cmd.CombinedOutput()

				returnValue := ""

				if err != nil {
					returnValue = err.Error()
				}

				database.CreateCommandLog(&command, job, returnValue, string(out))
			}
		}
	}
}
