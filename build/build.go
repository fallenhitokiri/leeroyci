// Build runs the build commands for a repository.
package build

import (
	"ironman/callbacks"
	"ironman/config"
	"ironman/logging"
	"ironman/notification"
	"log"
	"os/exec"
)

// Build waits for new notifications and runs the build process after
// receiving one.
func Build(not chan callbacks.Notification, c *config.Config, b *logging.Buildlog) {
	for {
		n := <-not
		run(n, c, b)
	}
}

// Run a build porcess.
func run(n callbacks.Notification, c *config.Config, b *logging.Buildlog) {
	repo := n.URL()
	branch := n.Branch()
	name, email := n.By()
	config, err := c.ConfigForRepo(repo)

	if err != nil {
		log.Println("could not find repo", repo)
		return
	}

	log.Println("Starting build process for", repo, branch)
	for _, cmd := range config.Commands {
		log.Println("Building", cmd.Name)
		out, code := call(cmd.Execute, repo, branch)
		job := b.Add(repo, branch, cmd.Name, name, email, out, code)
		go notification.Notify(c, job)
	}
	log.Println("Finished building", repo, branch)
}

// Call a build script and return the output.
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
