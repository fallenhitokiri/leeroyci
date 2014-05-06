// Build runs the build commands for a repository.
package build

import (
	"ironman/config"
	"ironman/logging"
	"ironman/notification"
	"log"
	"os/exec"
)

// Build waits for new notifications and runs the build process after
// receiving one.
func Build(jobs chan logging.Job, c *config.Config, b *logging.Buildlog) {
	for {
		j := <-jobs
		run(j, c, b)
	}
}

// Run a build porcess.
func run(j logging.Job, c *config.Config, b *logging.Buildlog) {
	config, err := c.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("could not find repo", j.URL)
		return
	}

	log.Println("Starting build process for", j.URL, j.Branch)
	for _, cmd := range config.Commands {
		log.Println("Building", cmd.Name)
		out, code := call(cmd.Execute, j.URL, j.Branch)
		j.Command = cmd.Name
		j.ReturnCode = code
		j.Output = out
		b.Add(j)
		go notification.Notify(c, &j)
	}
	log.Println("Finished building", j.URL, j.Branch)
}

// Call a build script and return the output.
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
