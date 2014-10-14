// Handle finished build jobs and trigger a deployment if configured.
package deployment

import (
	"leeroy/config"
	"leeroy/logging"
	"log"
)

// Check if the branch that was just build is configured for continuous
// deployment and trigger a deploy.
// Returns true if a deployment is triggered.
func ContinuousDeploy(j *logging.Job, c *config.Config) bool {
	r, err := c.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("Cannot deploy", j.Branch, "repository not found.")
		return false
	}

	_, err = r.DeployTarget(j.Branch)

	if err != nil {
		return false
	}

	log.Println("deploying", j.Branch)

	go Deploy(j, c)

	return true
}
