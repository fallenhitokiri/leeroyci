// GitHub provides all structs to unmarshal a GitHub webhook.
package github

import (
	"encoding/json"
	"ironman/logging"
	"log"
	"time"
)

// Parse a GitHub request body and add it to the build queue.
func Parse(jobs chan logging.Job, body []byte) {
	var cb GitHubCallback
	err := json.Unmarshal(body, &cb)

	name, email := cb.By()

	j := logging.Job{
		URL:       cb.URL(),
		Branch:    cb.Branch(),
		Timestamp: time.Now(),
		Commit:    cb.Commit(),
		Name:      name,
		Email:     email,
	}

	if err != nil {
		log.Println(string(body))
		panic("Could not unmarshal request")
	}

	if cb.ShouldBuild() == true {
		jobs <- j
	} else {
		log.Println("Not adding", cb.URL(), cb.Branch(), "to build queue")
	}
}
