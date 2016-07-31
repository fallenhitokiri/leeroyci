package database

// MailServer stores the configuration for a mail server used to send
// notifications.
type MailServer struct {
	Host     string
	Sender   string
	Port     int
	User     string
	Password string
}
