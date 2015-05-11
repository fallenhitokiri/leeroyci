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
	DeletedAt *time.Time
}

// AddRepository adds a new repository.
func AddRepository(name, url, accessKey string, commentPR, closePR, statusPR bool) *Repository {
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
