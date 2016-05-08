// Package config contains all data models used for LeeroyCI.
package config

// import "errors"

// Project represents one project / repository used to run jobs for.
type Project struct {
	Name      string
	URL       string
	AccessKey string

	// If ClosePR is true LeeroyCI will try to close the pull request if jobs
	// fail and if this is supported by the git hosting service.
	ClosePR bool

	// If StatusPR is true LeeroyCI will try to comment on commits after jobs ran
	// if this is supported by the git hosting service.
	StatusPR bool

	Tasks []*Task

	Notifications []*Notification
}

// func NewProject(name string) (*Project, error)  {
//
// }
