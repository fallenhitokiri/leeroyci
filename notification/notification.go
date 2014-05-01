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
	code := "0"
	if job.ReturnCode != nil {
		code = job.ReturnCode.Error()
	}

	from := mail.Address{"Ironman", c.EmailFrom}
	to := mail.Address{job.Name, job.Email}
	subject := "Build finished"

	body := "Finished building your commits.\n\n"
	body = body + "Repo: " + job.URL + "\n"
	body = body + "Branch: " + job.Branch + "\n"
	body = body + "Time: " + job.Timestamp.String() + "\n"
	body = body + "Command: " + job.Command + "\n"
	body = body + "Return Code: " + code + "\n\n"
	body = body + "Output: \n" + job.Output + "\n"

	message := buildEmail(from.String(), to.String(), subject, body)

	email(c, job.Email, message)
}
