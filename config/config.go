// Package config contains all data models used for LeeroyCI.
package config

// Config stores everything configuraiton related to the process, including
// users and job configurations.
type Config struct {
	URL            string
	Secret         string
	SSLCert        string
	SSLKey         string
	ParallelBuilds int

	Users []*User

	MailServer *MailServer

	Projects []*Project
}
