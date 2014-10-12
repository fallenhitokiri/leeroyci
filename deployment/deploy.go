// Deploy implements the actual deployment process for successfully build jobs.
package deployment

import (
	"leeroy/config"
	"leeroy/logging"
	"log"
	"os/exec"
)

// Deploy a job if all tests are passed.
func Deploy(j *logging.Job, c *config.Config) {
	if j.Success() != true {
		log.Println("Not deploying", j.Branch, "build did not succeed.")
		return
	}

	r, err := c.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("Repo", j.URL, "does not exist.")
		return
	}

	d, err := r.DeployTarget(j.Branch)

	if err != nil {
		log.Println("Deployment target for", j.Branch, "does not exist")
		return
	}

	announceStart(j, c, &d)

	o, err := call(d.Execute, r.URL, j.Branch)

	if err != nil {
		log.Println(err.Error())
	}

	announceComplete(j, c, &d, o)
}

// Log and announce a started deployment.
func announceStart(j *logging.Job, c *config.Config, d *config.Deploy) {
	log.Println(j.Name, "triggered a deploy for", j.Branch, "to", d.Name)
}

// Log and announce a completed deployment.
func announceComplete(j *logging.Job, c *config.Config, d *config.Deploy, output string) {
	log.Println("finished deploying", j.Branch, "to", d.Name, "with", output)
}

// Call a deployment script and return the output.
func call(app string, repo string, branch string) (string, error) {
	cmd := exec.Command(app, repo, branch)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
