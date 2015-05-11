// Package database provides a wrapper between the database and stucts
package database

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	ID      int64
	Name    string
	Execute string

	Repository   Repository
	RepositoryID int64
}

// AddCommand adds a new command to a repository.
func AddCommand(r *Repository, name, execute string) *Command {
	c := Command{
		Name:       name,
		Execute:    execute,
		Repository: *r,
	}

	db.Save(&c)

	return &c
}
