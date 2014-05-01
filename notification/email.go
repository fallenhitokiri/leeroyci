// Implement email notifications.
package notification

import (
	"encoding/base64"
	"fmt"
	"ironman/config"
	"log"
	"net/smtp"
)

// Send an email to `toName <toEmail>` with the details of the failed build.
func email(c *config.Config, to string, message []byte) {
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

// Build a string to be used as argument for net/smtp to send as mail.
func buildEmail(from, to, subject, body string) []byte {
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
