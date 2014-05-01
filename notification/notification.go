package notification

import (
	"ironman/config"
	"ironman/logging"
	//"log"
	"net/mail"
)

// Run all notifications configured for a
func Notify(c *config.Config, job *logging.Job) {
	go notifyPusher(c, job)

	/*config, err := c.ConfigForRepo(job.URL)

	if err != nil {
		log.Println("could not find repo", job.URL)
		return
	}*/
}

// Notify the person who pushed the changes
func notifyPusher(c *config.Config, job *logging.Job) {
	from := mail.Address{"Ironman", c.EmailFrom}
	to := mail.Address{job.Name, job.Email}

	subject := "Build for " + job.Branch + " finished "

	if job.Success() == true {
		subject = subject + "successfully"
	} else {
		subject = subject + "with errors"
	}

	body := "Finished building your commits.\n\n"
	body = body + "Repo: " + job.URL + "\n"
	body = body + "Branch: " + job.Branch + "\n"
	body = body + "Time: " + job.Timestamp.String() + "\n"
	body = body + "Command: " + job.Command + "\n"
	body = body + "Return Code: " + job.Status() + "\n\n"
	body = body + "Output: \n" + job.Output + "\n"

	message := buildEmail(from.String(), to.String(), subject, body)

	email(c, job.Email, message)
}
