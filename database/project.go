package database

import "errors"

var (
	ErrorProjectNotFound   = errors.New("Project not found")
	ErrorProjectNameExists = errors.New("Project with this name exists")
)

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

	Tasks         []*Task
	Notifications []*Notification
	Branches      []*Branch
}
