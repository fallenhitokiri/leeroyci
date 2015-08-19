package runner

import (
	"log"
	"os/exec"

	"github.com/fallenhitokiri/leeroyci/database"
)

var RunQueue = make(chan int64, 100)

func Runner() {
	for {
		jobID := <-RunQueue

		job := database.GetJob(jobID)
		repository, err := database.GetRepositoryByID(string(job.RepositoryID))

		if err != nil {
			log.Println("Could not find repository for", job.Repository.URL)
			return
		}

		run(job, repository, database.CommandKindTest)
		run(job, repository, database.CommandKindBuild)
		job.TasksDone()

		run(job, repository, database.CommandKindDeploy)
		job.DeployDone()
	}
}

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
