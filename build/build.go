// Package build runs the build commands for a repository.
package build

import (
	"leeroy/database"
	"leeroy/deployment"
	"leeroy/notification"
	"log"
	"os/exec"
)

// Build waits for new notifications and runs the build process after
// receiving one.
func Build(jobs chan database.Job) {
	for {
		j := <-jobs
		run(j)
	}
}

// Run a build porcess.
func run(j database.Job) {
	r, err := config.CONFIG.ConfigForRepo(j.URL)
	j.Identifier = r.Identifier()

	if err != nil {
		log.Println("could not find repo", j.URL)
		return
	}

	log.Println("Starting build process for", j.URL, j.Branch)

	for _, cmd := range r.Commands {
		log.Println("Building", cmd.Name)

		o, err := call(cmd.Execute, j.URL, j.Branch)
		t := logging.Task{
			Command: cmd.Name,
			Output:  o,
		}

		if err != nil {
			t.Return = err.Error()
		}

		j.Add(t)
	}

	logging.BUILDLOG.Add(&j)
	go notification.Notify(&j, notification.KindBuild)
	go deployment.ContinuousDeploy(&j)

	log.Println("Finished building", j.URL, j.Branch)
}

// Call a build script and return the output.
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
