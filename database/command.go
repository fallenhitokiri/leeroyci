// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"
	"time"
)

const (
	// CommandKindTest is used when a command runs tests.
	CommandKindTest = "test"

	// CommandKindBuild is used when a command builds a package or project.
	CommandKindBuild = "build"

	// CommandKindDeploy is used when a command deploys a branch.
	CommandKindDeploy = "deploy"
)

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	ID      int64
	Name    string
	Kind    string
	Branch  string
	Execute string

	Repository   Repository
	RepositoryID int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

// CommandLog stored a finnished command and the output of the task.
type CommandLog struct {
	ID     int64
	Name   string // we only keep the name, no reference to the command, in case it changes.
	Return string
	Output string

	Job   Job
	JobID int64
}

// AddCommand adds a new command to a repository.
func CreateCommand(repo *Repository, name, execute, branch, kind string) (*Command, error) {
	if kind != CommandKindTest && kind != CommandKindBuild && kind != CommandKindDeploy {
		return nil, errors.New("wrong kind")
	}

	c := Command{
		Name:       name,
		Execute:    execute,
		Kind:       kind,
		Repository: *repo,
		Branch:     branch,
	}

	db.Save(&c)

	return &c, nil
}

// GetCommand returns a command for a specific ID.
func GetCommand(id string) (*Command, error) {
	c := &Command{}
	db.Where("ID = ?", id).First(&c)
	return c, nil
}

// UpdateCommand updates a command.
func (c *Command) Update(name, kind, branch, execute string) error {
	c.Name = name
	c.Kind = kind
	c.Branch = branch
	c.Execute = execute

	db.Save(c)

	return nil
}

// DeleteCommand deletes a command.
func (c *Command) Delete() {
	db.Delete(c)
}

// CreateCommandLog adds a new log.
func CreateCommandLog(command *Command, job *Job, ret, out string) *CommandLog {
	log := CommandLog{
		Name:   command.Name,
		Return: ret,
		Output: out,
		Job:    *job,
	}

	db.Save(&log)

	return &log
}

// Passed returns true if the command completed successfully.
func (t *CommandLog) Passed() bool {
	if t.Return == "" {
		return true
	}

	return false
}
