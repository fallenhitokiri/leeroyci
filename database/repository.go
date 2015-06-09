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
func CreateRepository(name, url, accessKey string, commentPR, closePR, statusPR bool) *Repository {
	repo := Repository{
		Name:      name,
		URL:       url,
		AccessKey: accessKey,
		CommentPR: commentPR,
		ClosePR:   closePR,
		StatusPR:  statusPR,
	}

	db.Save(&repo)

	return &repo
}

// GetRepository returns the repository based on the URL that pushed changes.
func GetRepository(url string) *Repository {
	repo := &Repository{}
	db.Where("URL = ?", url).First(&repo)
	return repo
}

// UpdateRepository updates an existing repository.
func (r *Repository) Update(name, url, accessKey string, commentPR, closePR, statusPR bool) {
	r.Name = name
	r.CommentPR = commentPR
	r.StatusPR = statusPR
	r.ClosePR = closePR
	r.AccessKey = accessKey

	db.Save(r)
}

// DeleteRepository deletes an existing repository.
func (r *Repository) Delete(url string) {
	db.Delete(r)
}

// Jobs returns all jobs for this repository.
func (r *Repository) Jobs() []Job {
	jobs := []Job{}

	db.Where("repository_id = ?", r.ID).Find(&jobs)

	return jobs
}
