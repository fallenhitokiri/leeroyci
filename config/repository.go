// Package config takes care of the whole configuration.
package config

import (
	"errors"
)

// Repository holds all information needed to identify a repository and run
// tests and builds.
type Repository struct {
	Name      string
	URL       string
	Commands  []Command
	Notify    []Notify
	CommentPR bool
	ClosePR   bool
	StatusPR  bool
	AccessKey string
	Deploy    []Deploy
}

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	Name    string
	Execute string
}

// Notify stores the configuration needed for a notification plugin to work. All
// optiones required by the services are stored as map - it is the job of the
// notification plugin to access them correctly and handle missing ones.
type Notify struct {
	Service   string
	Arguments map[string]string
}

// Deploy stores the command to execute and arguments to pass to the command
// when a users pushes to a certain branch.
type Deploy struct {
	Name      string
	Branch    string
	Execute   string
	Arguments map[string]string
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
