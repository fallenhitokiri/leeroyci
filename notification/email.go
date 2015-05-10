// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"encoding/base64"
	"fmt"
	"leeroy/database"
	"log"
	"net/mail"
	"net/smtp"
)

// Send an email to `toName <toEmail>` with the details of the failed build.
func email(n *notification, to string) {
	message := buildEmail(n)

	ms := database.GetMailServer()

	auth := smtp.PlainAuth(
		"",
		ms.User,
		ms.Password,
		ms.Host,
	)

	err := smtp.SendMail(
		ms.Server(),
		auth,
		ms.Sender,
		[]string{to},
		message,
	)

	if err != nil {
		log.Println(err)
	}
}

// Notify the person who pushed the changes
func buildEmail(n *notification) []byte {
	ms := database.GetMailServer()

	f := mail.Address{Name: "leeroy", Address: ms.Sender}
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
