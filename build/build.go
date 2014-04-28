// Build runs the build commands for a repository.
package build

import (
	"ironman/callbacks"
	"ironman/config"
	"log"
	"net/smtp"
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
		log.Println("could not find repo", repo)
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

// Call a build script and return the output.
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.Output()
	return string(out), err
}

// Notify a pusher that a build failed.
func notify_fail(cmd config.Command, repo string, branch string, out string,
	err error, c *config.Config, name string, mail string) error {
	mime := "MIME-version: 1.0;Content-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Build Failed!\n"
	body := repo + "\n" + branch + "\n" + cmd.Name + "\n\n"
	body = body + err.Error() + "\n\n" + out
	// TODO: this should look better

	msg := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", c.EmailUser, c.EmailPassword, c.EmailHost)

	err = smtp.SendMail(c.MailServer(), auth, c.EmailFrom, []string{mail}, msg)

	if err != nil {
		log.Println(err)
	}

	return err
}
