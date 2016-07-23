// Package config contains all data models used for LeeroyCI.
package config

import "errors"

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

// NewProject creates a new project. An AccessKey is also created.
func NewProject(name, url string, closePR, statusPR bool) (*Project, error) {
	_, err := ProjectByName(name)
	if err == nil {
		return nil, errors.New("Project with this name already exists.")
	}

	accessKey := generateAccessKey(accessKeyDictionary, accessKeyLength)

	project := &Project{
		Name:      name,
		URL:       url,
		AccessKey: accessKey,
		ClosePR:   closePR,
		StatusPR:  statusPR,
	}

	cfg.Projects = append(cfg.Projects, project)

	return project, nil
}

// ProjectByName returns a project with the matching name or an error.
func ProjectByName(name string) (*Project, error) {
	for _, project := range cfg.Projects {
		if project.Name == name {
			return project, nil
		}
	}

	return nil, errors.New("Project not found.")
}

// ProjectByAccessKey returns a project with the matching access key or an
// error.
func ProjectByAccessKey(key string) (*Project, error) {
	for _, project := range cfg.Projects {
		if project.AccessKey == key {
			return project, nil
		}
	}

	return nil, errors.New("Project not found.")
}
