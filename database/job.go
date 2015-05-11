// Package database provides a wrapper between the database and stucts
package database

import (
	"time"
)

// Job stores all information about one commit and the executed tasks.
type Job struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	BuildDone time.Time

	Repository   Repository
	RepositoryID int64

	Branch    string
	Commit    string
	CommitURL string

	Name  string
	Email string

	Tasks    []Task `gorm:"many2many:job_tasks;"`
	Deployed *Task
}

// AddJob adds a new job to the database.
func AddJob(url, branch, commit, name, email, commitURL string) *Job {
	r := RepositoryForURL(url)

	j := Job{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Repository: *r,
		Branch:     branch,
		Commit:     commit,
		CommitURL:  commitURL,
		Name:       name,
		Email:      email,
	}

	return &j
}

// GetOpenJobs returns all jobs which are not finished.
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

// Add a task to the job.
func (j *Job) Add(t Task) {
	j.Tasks = append(j.Tasks, t)
}

// DeploySuccess returns if the deploy was successful.
func (j *Job) DeploySuccess() bool {
	if j.Deployed == nil {
		return false
	}

	if j.Deployed.Return == "" {
		return true
	}
	return false
}
