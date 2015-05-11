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
	r := Repository{
		Name:      name,
		URL:       url,
		AccessKey: accessKey,
		CommentPR: commentPR,
		ClosePR:   closePR,
		StatusPR:  statusPR,
	}

	db.Save(&r)

	return &r
}

// RepositoryForURL returns the repository based on the URL that pushed changes.
func RepositoryForURL(url string) *Repository {
	r := &Repository{}
	db.Where("URL = ?", url).First(&r)
	return r
}

// Identifier returns the name or the URL
func (r *Repository) Identifier() string {
	if r.Name != "" {
		return r.Name
	}
	return r.URL
}

// DeployTarget returns the deployment target for a branch
/*func (r *Repository) DeployTarget(branch string) (Deploy, error) {
	for _, d := range r.Deploy {
		if d.Branch == branch {
			return d, nil
		}
	}
	return Deploy{}, errors.New("No deployment target for branch")
}*/
