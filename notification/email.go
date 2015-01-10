// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"encoding/base64"
	"fmt"
	"leeroy/config"
	"log"
	"net/mail"
	"net/smtp"
)

// Send an email to `toName <toEmail>` with the details of the failed build.
func email(c *config.Config, n *notification, to string) {
	message := buildEmail(c, n)
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
func buildEmail(c *config.Config, n *notification) []byte {
	f := mail.Address{Name: "leeroy", Address: c.EmailFrom}
	t := mail.Address{Name: n.Name, Address: n.Email}
	s := subject(n)
	b := n.message
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
func subject(n *notification) string {
	if n.Status == true {
		return fmt.Sprintf("%s: success", n.Branch)
	}
	return fmt.Sprintf("%s: failed", n.Branch)
}
