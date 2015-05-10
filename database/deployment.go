// Package database provides a wrapper between the database and stucts
package database

// Deploy stores the command to execute and arguments to pass to the command
// when a users pushes to a certain branch.
type Deploy struct {
	ID        int64
	Name      string
	Branch    string
	Execute   string
	Arguments string

	Repository   Repository
	RepositoryID int64
}

func AddDeploy(r *Repository, name, branch, execute, arguments string) *Deploy {
	d := Deploy{
		Name:       name,
		Branch:     branch,
		Execute:    execute,
		Arguments:  arguments,
		Repository: r,
	}

	db.Save(&d)

	return &d
}
