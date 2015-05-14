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

// AddJob adds a new job to the database.
func AddJob(repo *Repository, branch, commit, commitURL, name, email string) *Job {
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

/*// GetOpenJobs returns all jobs which are not finished.
func GetOpenJobs() []*Job {
	j := []*Job{}

	db.Where("BuildDone = NULL").Find(&j)

	return j
}

// GetAllJobs returns all jobs.
func GetAllJobs() []*Job {
	j := []*Job{}
	db.Find(&j)
	return j
}

// Status returns either the exit code of the triggered command or 0 if the
// command finished successfully.
func (j *Job) Status() string {
	code := "0"

	for _, task := range j.Tasks {
		if task.Return != "" {
			return task.Return
		}
	}

	return code
}

// Success returns true if the build was successful.
func (j *Job) Success() bool {
	if j.Status() == "0" {
		return true
	}
	return false
}
