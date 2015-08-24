// Package database provides a wrapper between the database and stucts
package database

import (
	"fmt"
	//"log"
	"time"
)

// Define all statuses a job can have.
const (
	JobStatusSuccess = "success"
	JobStatusError   = "error"
	JobStatusPending = "pending"
)

// Job stores all information about one commit and the executed tasks.
type Job struct {
	ID int64

	TasksFinished  time.Time
	DeployFinished time.Time

	Repository   Repository
	RepositoryID int64

	Branch    string
	Commit    string `gorm:"column:commit_sha"`
	CommitURL string

	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time

	CommandLogs []CommandLog
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
	db.Preload("Repository").Preload("CommandLogs").Where("ID = ?", id).First(&j)
	return j
}

// GetJobs returns a list of jobs for a given range.
func GetJobs(offset, limit int) []*Job {
	var jobs []*Job

	db.Preload(
		"Repository",
	).Preload(
		"CommandLogs",
	).Offset(
		offset,
	).Limit(
		limit,
	).Order(
		"created_at desc",
	).Find(&jobs)

	return jobs
}

// GetJobByCommit returns with a specific commit ID or nil.
func GetJobByCommit(commit string) *Job {
	job := &Job{}

	db.Where("commit_sha = ?", commit).Last(&job)

	return job
}

// NumberOfJobs returns the number of all existing jobs.
func NumberOfJobs() int {
	var count int

	db.Table("jobs").Count(&count)

	return count
}

// Passed returns true if all commands succeeded.
func (j *Job) Passed() bool {
	logs := GetCommandLogsForJob(j.ID)

	for _, c := range logs {
		if !c.Passed() {
			return false
		}
	}
	return true
}

// Status returns the current status fo the job.
func (j *Job) Status() string {
	n := time.Time{}
	if j.TasksFinished != n {
		if j.Passed() {
			return JobStatusSuccess
		}

		return JobStatusError
	}
	return JobStatusPending
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

func (j *Job) URL() string {
	config := GetConfig()
	return fmt.Sprintf("%s/%d", config.URL, j.ID)
}
