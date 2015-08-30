// Package runner runs all tasks for all commands associated with a repository.
package runner

import (
	"log"
	"os/exec"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/notification"
)

// RunQueue receives job IDs for which commands should run.
var RunQueue = make(chan int64, 100)

// Runner waits for jobs to be pushed on RunQueue and runs all commands. It also
// creates the command logs and sends the necessary notifications.
func Runner() {
	for {
		jobID := <-RunQueue

		job := database.GetJob(jobID)
		repository, err := database.GetRepositoryByID(job.RepositoryID)

		if err != nil {
			log.Println("Could not find repository for", job.Repository.URL)
			return
		}

		run(job, repository, database.CommandKindTest)
		notification.Notify(job, notification.EventTest)

		if job.Passed() {
			run(job, repository, database.CommandKindBuild)
			notification.Notify(job, notification.EventBuild)
		}

		job.TasksDone()

		if job.Passed() {
			notification.Notify(job, notification.EventDeployStart)
			run(job, repository, database.CommandKindDeploy)
			job.DeployDone()
			notification.Notify(job, notification.EventDeployEnd)
		}
	}
}

// run runs the command that is specified in Command.Execute and creates the
// command log with the results of the command.
func run(job *database.Job, repository *database.Repository, kind string) {
	commands := repository.GetCommands(job.Branch, kind)

	for _, command := range commands {
		if command.Kind == kind {
			if (command.Branch != "" && command.Branch == job.Branch) || command.Branch == "" {
				repository := job.Repository.Name
				branch := job.Branch

				log.Println("Running", command.Name, "for", repository, branch)

				cmd := exec.Command(command.Execute, repository, branch)
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
