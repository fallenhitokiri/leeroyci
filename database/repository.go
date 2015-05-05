// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"
)

// Repository holds all information needed to identify a repository and run
// tests and builds.
type Repository struct {
	ID            int
	Name          string
	URL           string
	Commands      []Command
	Notifications []Notify
	CommentPR     bool
	ClosePR       bool
	StatusPR      bool
	AccessKey     string
	Deploy        []Deploy
}

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	ID      int
	Name    string
	Execute string
}

// Notify stores the configuration needed for a notification plugin to work. All
// optiones required by the services are stored as map - it is the job of the
// notification plugin to access them correctly and handle missing ones.
type Notify struct {
	ID        int
	Service   string
	Arguments string
}

// Deploy stores the command to execute and arguments to pass to the command
// when a users pushes to a certain branch.
type Deploy struct {
	ID        int
	Name      string
	Branch    string
	Execute   string
	Arguments string
}

// RepositoryForURL returns the repository based on the URL that pushed changes.
func RepositoryForURL(url string) *Repository {
	r := &Repository{}
	db.Where("URL = ?", url).First(&r)
	return r
}

// Identifier returns the name or the URL
func (r *Repository) Identifier() string {
	if r.Name != "" {
		return r.Name
	}
	return r.URL
}

// DeployTarget returns the deployment target for a branch
func (r *Repository) DeployTarget(branch string) (Deploy, error) {
	for _, d := range r.Deploy {
		if d.Branch == branch {
			return d, nil
		}
	}
	return Deploy{}, errors.New("No deployment target for branch")
}
