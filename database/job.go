// Package database provides a wrapper between the database and stucts
package database

import (
	"time"
)

// Job stores all information about one commit and the executed tasks.
type Job struct {
	ID int64

	TasksFinished  time.Time
	DeployFinished time.Time

	Repository   Repository
	RepositoryID int64

	Branch    string
	Commit    string
	CommitURL string

	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateJob adds a new job to the database.
func CreateJob(repo *Repository, branch, commit, commitURL, name, email string) *Job {
	j := &Job{
		Repository: *repo,
		Branch:     branch,
		Commit:     commit,
		CommitURL:  commitURL,
		Name:       name,
		Email:      email,
	}

	db.Save(j)

	return j
}

// GetJob returns a job for a given ID.
func GetJob(id int64) *Job {
	j := &Job{}
	db.Where("ID = ?", id).First(&j)
	return j
}

// GetJobs returns a list of jobs for a given range.
func GetJobs(offset, limit int) []*Job {
	jobs := make([]*Job, 0)

	db.Offset(offset).Limit(limit).Find(&jobs)

	return jobs
}

// NumberOfJobs returns the number of all existing jobs.
func NumberOfJobs() int {
	var count int

	db.Table("jobs").Count(&count)

	return count
}

// TasksDone sets TasksDone
func (j *Job) TasksDone() {
	j.TasksFinished = time.Now()
	db.Save(j)
}

// DeployDone sets DeployDone
func (j *Job) DeployDone() {
	j.DeployFinished = time.Now()
	db.Save(j)
}
