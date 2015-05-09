// Package deployment handles finished build jobs and triggers a deployment
// script if configured.
package deployment

import (
	"leeroy/config"
	"leeroy/database"
	"log"
)

// ContinuousDeploy checks if the branch that was just built is configured for
// continuous deployment and triggers a deploy if that is the case.
// Returns true if a deployment is triggered.
func ContinuousDeploy(j *database.Job) bool {
	r, err := config.CONFIG.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("Cannot deploy", j.Branch, "repository not found.")
		return false
	}

	_, err = r.DeployTarget(j.Branch)

	if err != nil {
		return false
	}

	log.Println("deploying", j.Branch)

	go Deploy(j)

	return true
}
