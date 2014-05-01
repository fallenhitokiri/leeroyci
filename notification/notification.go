package notification

import (
	"ironman/config"
	"ironman/logging"
	//"log"
)

// Run all notifications configured for a
func Notify(c *config.Config, j *logging.Job) {
	// always notify the person who comitted
	go email(c, j, j.Email)

	/*config, err := c.ConfigForRepo(job.URL)

	if err != nil {
		log.Println("could not find repo", job.URL)
		return
	}*/
}
