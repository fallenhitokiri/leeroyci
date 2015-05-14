// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"
	"time"
)

const (
	CommandKindTest   = "test"
	CommandKindBuild  = "build"
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

// AddCommand adds a new command to a repository.
func AddCommand(repo *Repository, name, execute, branch, kind string) (*Command, error) {
	if kind != CommandKindTest && kind != CommandKindBuild && kind != CommandKindDeploy {
		return nil, errors.New("wrong kind.")
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

// GetCommands returns all commands for a repository, branch and kind
func GetCommands(repo *Repository, branch, kind string) []Command {
	noBranch := []Command{}
	branches := []Command{}

	db.Where(&Command{
		RepositoryID: repo.ID,
		Kind:         kind,
		Branch:       "",
	}).Find(&noBranch)

	if branch != "" {
		db.Where(&Command{
			RepositoryID: repo.ID,
			Kind:         kind,
			Branch:       branch,
		}).Find(&branches)
	}

	return append(noBranch, branches...)
}

// GetCommand returns a command for a specific ID.
func GetCommand(id int64) *Command {
	c := &Command{}
	db.Where("ID = ?", id).First(&c)
	return c
}

// UpdateCommand updates a command.
func UpdateCommand(id int64, name, kind, branch, execute string) *Command {
	c := GetCommand(id)

	c.Name = name
	c.Kind = kind
	c.Branch = branch
	c.Execute = execute

	db.Save(c)

	return c
}

// DeleteCommand deletes a command.
func DeleteCommand(id int64) {
	c := GetCommand(id)
	db.Delete(c)
}
