// Package database provides a wrapper between the database and stucts
package database

import (
	"time"
)

// Repository holds all information needed to identify a repository and run
// tests and builds.
type Repository struct {
	ID   int64
	Name string
	URL  string

	ClosePR   bool
	StatusPR  bool
	AccessKey string

	CreatedAt time.Time
	UpdatedAt time.Time

	Notifications []Notification
	Commands      []Command
}

// CreateRepository adds a new repository.
func CreateRepository(name, url, accessKey string, closePR, statusPR bool) (*Repository, error) {
	repo := Repository{
		Name:      name,
		URL:       url,
		AccessKey: accessKey,
		ClosePR:   closePR,
		StatusPR:  statusPR,
	}

	db.Save(&repo)

	return &repo, nil
}

// GetRepository returns the repository based on the URL that pushed changes.
func GetRepository(url string) *Repository {
	repo := &Repository{}
	db.Preload("Notifications").Preload("Commands").Where("URL = ?", url).First(&repo)
	return repo
}

// GetRepositoryByID returns the repository based on the ID.
func GetRepositoryByID(id int64) (*Repository, error) {
	repo := &Repository{}
	db.Preload("Notifications").Preload("Commands").Where("ID = ?", id).First(&repo)
	return repo, nil
}

// Update this repository.
func (r *Repository) Update(name, url, accessKey string, closePR, statusPR bool) (*Repository, error) {
	r.Name = name
	r.StatusPR = statusPR
	r.ClosePR = closePR
	r.AccessKey = accessKey

	db.Save(r)

	return r, nil
}

// ListRepositories returns all repositories.
func ListRepositories() []*Repository {
	var repos []*Repository
	db.Find(&repos)
	return repos
}

// Delete this repository.
func (r *Repository) Delete() error {
	db.Delete(r)

	return nil
}

// Jobs returns all jobs for this repository.
func (r *Repository) Jobs() []Job {
	jobs := []Job{}

	db.Where("repository_id = ?", r.ID).Find(&jobs)

	return jobs
}

// GetCommands returns all commands for a repository, branch and kind
func (r *Repository) GetCommands(branch, kind string) []Command {
	commands := []Command{}
	branchSpecific := []Command{}

	db.Where(&Command{
		RepositoryID: r.ID,
		Kind:         kind,
		Branch:       "",
	}).Find(&commands)

	if branch != "" {
		existing := []int64{}

		for _, command := range commands {
			existing = append(existing, command.ID)
		}

		db.Where(&Command{
			RepositoryID: r.ID,
			Kind:         kind,
			Branch:       branch,
		}).Not("id", existing).Find(&branchSpecific)

		commands = append(commands, branchSpecific...)
	}

	return commands
}
