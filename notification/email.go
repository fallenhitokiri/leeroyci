// Implement email notifications.
package notification

import (
	"encoding/base64"
	"fmt"
	"ironman/config"
	"ironman/logging"
	"log"
	"net/mail"
	"net/smtp"
)

// Send an email to `toName <toEmail>` with the details of the failed build.
func email(c *config.Config, j *logging.Job, to string) {
	message := buildEmail(c, j)
	auth := smtp.PlainAuth("", c.EmailUser, c.EmailPassword, c.EmailHost)

	err := smtp.SendMail(
		c.MailServer(),
		auth,
		c.EmailFrom,
		[]string{to},
		message,
	)

	if err != nil {
		log.Println(err)
	}
}

// Notify the person who pushed the changes
func buildEmail(c *config.Config, j *logging.Job) []byte {
	from := mail.Address{"Ironman", c.EmailFrom}
	to := mail.Address{j.Name, j.Email}

	subject := "Build for " + j.Branch + " finished "

	if j.Success() == true {
		subject = subject + "successfully"
	} else {
		subject = subject + "with errors"
	}

	body := "Repo: " + j.URL + "\n"
	body = body + "Branch: " + j.Branch + "\n"
	body = body + "Time: " + j.Timestamp.String() + "\n"
	body = body + "Return: " + j.Status() + "\n\n\n"

	for _, t := range j.Tasks {
		body = body + "Command: " + t.Command + "\n"
		body = body + "Return: " + t.Status() + "\n\n"
		body = body + "Output: \n" + t.Output + "\n\n\n"
	}

	message := addHeaders(from.String(), to.String(), subject, body)

	return message
}

// Build a string to be used as argument for net/smtp to send as mail.
func addHeaders(from, to, subject, body string) []byte {
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}
