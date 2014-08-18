// Implement email notifications.
package notification

import (
	"encoding/base64"
	"fmt"
	"leeroy/config"
	"leeroy/logging"
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
	f := mail.Address{Name: "leeroy", Address: c.EmailFrom}
	t := mail.Address{Name: j.Name, Address: j.Email}
	s := subject(j)
	b := body(j, c.URL)
	m := addHeaders(f.String(), t.String(), s, b)
	return m
}

// Build a string to be used as argument for net/smtp to send as mail.
func addHeaders(from, to, subject, body string) []byte {
	h := make(map[string]string)
	h["From"] = from
	h["To"] = to
	h["Subject"] = subject
	h["MIME-Version"] = "1.0"
	h["Content-Type"] = "text/plain; charset=\"utf-8\""
	h["Content-Transfer-Encoding"] = "base64"

	m := ""

	for k, v := range h {
		m += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	m += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(m)
}

// Returns the subject for the mail.
func subject(j *logging.Job) string {
	if j.Success() == true {
		return fmt.Sprintf("Build for %s finished successfully", j.Branch)
	}
	return fmt.Sprintf("Build for %s finished with errors", j.Branch)
}

// Returns the body for the mail.
func body(j *logging.Job, u string) string {
	b := fmt.Sprintf(
		"Repo: %s - Branch: %s\nTime: %s\nReturn: %s\n\n\n",
		j.URL,
		j.Branch,
		j.Timestamp.String(),
		j.Status(),
	)

	for _, t := range j.Tasks {
		b = b + fmt.Sprintf(
			"Command: %s -> %s\n",
			t.Command,
			t.Status(),
		)
	}

	b = b + fmt.Sprintf("\n\n%s\n", j.StatusURL(u))

	return b
}
