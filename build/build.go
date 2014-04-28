package build

import (
	"ironman/callbacks"
	"ironman/config"
	"log"
	"os/exec"
)

// Build waits for new notifications and runs the build process after
// receiving one.
func Build(not chan callbacks.Notification, c *config.Config, b *Buildlog) {
	for {
		n := <-not
		run(n, c, b)
	}
}

// Run a build porcess.
func run(n callbacks.Notification, c *config.Config, b *Buildlog) {
	repo := n.URL()
	branch := n.Branch()
	name, email := n.By()
	config, err := c.ConfigForRepo(repo)

	if err != nil {
		log.Println("could not find repo")
		return
	}

	for _, cmd := range config.Commands {
		out, code := call(cmd.Execute, repo, branch)
		b.Add(repo, branch, email, out, code)

		if code != nil {
			notify_fail(cmd, repo, branch, out, code, c, name, email)
		}
	}
}

// Call a build script and return the output
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.Output()
	return string(out), err
}

// Notify a pusher that a build failed
func notify_fail(cmd config.Command, repo string, branch string, out string,
	err error, c *config.Config, name string, mail string) {

}
