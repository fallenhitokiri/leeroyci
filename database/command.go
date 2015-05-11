// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"
)

const (
	CommandKindTest   = "test"
	CommandKindBuild  = "build"
	CommandKindDeploy = "deploy"
)

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	ID     int64
	Name   string
	Kind   string
	Branch string

	Execute string

	Repository   Repository
	RepositoryID int64
}

// AddCommand adds a new command to a repository.
func AddCommand(repo *Repository, name, execute, branch, kind string) (*Command, error) {
	if kind != CommandKindTest && kind != CommandKindBuild && kind != CommandKindDeploy {
		return nil, errors.New("wrong kind.")
	}

	command := Command{
		Name:       name,
		Execute:    execute,
		Kind:       kind,
		Repository: *repo,
		Branch:     branch,
	}

	db.Save(&command)

	return &command, nil
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
