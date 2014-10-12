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

	for _, d := range r.Deploy {
		if d.Branch == j.Branch {
			log.Println("deploying", d.Branch)
			go Deploy(j, c)
			return true
		}
	}

	return false
}
