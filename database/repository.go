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

	CommentPR bool
	ClosePR   bool
	StatusPR  bool
	AccessKey string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateRepository adds a new repository.
func CreateRepository(name, url, accessKey string, commentPR, closePR, statusPR bool) (*Repository, error) {
	repo := Repository{
		Name:      name,
		URL:       url,
		AccessKey: accessKey,
		CommentPR: commentPR,
		ClosePR:   closePR,
		StatusPR:  statusPR,
	}

	db.Save(&repo)

	return &repo, nil
}

// GetRepository returns the repository based on the URL that pushed changes.
func GetRepository(url string) *Repository {
	repo := &Repository{}
	db.Where("URL = ?", url).First(&repo)
	return repo
}

// GetRepositoryByID returns the repository based on the ID.
func GetRepositoryByID(id string) (*Repository, error) {
	repo := &Repository{}
	db.Where("ID = ?", id).First(&repo)
	return repo, nil
}

// UpdateRepository updates an existing repository.
func (r *Repository) Update(name, url, accessKey string, commentPR, closePR, statusPR bool) (*Repository, error) {
	r.Name = name
	r.CommentPR = commentPR
	r.StatusPR = statusPR
	r.ClosePR = closePR
	r.AccessKey = accessKey

	db.Save(r)

	return r, nil
}

// ListRepositories returns all repositories.
func ListRepositories() []*Repository {
	repos := make([]*Repository, 0)
	db.Find(&repos)
	return repos
}

// DeleteRepository deletes an existing repository.
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

// Commands returns all commands for a repository, branch and kind
func (r *Repository) Commands(repo *Repository, branch, kind string) []Command {
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
