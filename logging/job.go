// Each triggered build is represented as a Job.
package logging

import (
	"time"
)

type Job struct {
	URL       string
	Branch    string
	Commit    string
	Timestamp time.Time
	Name      string
	Email     string
	Tasks     []Task
}

type Task struct {
	Command string
	Return  error
	Output  string
}

// Returns either the exit code of the triggered command or 0 if the command
// finished successfully.
func (j *Job) Status() string {
	code := "0"

	for _, task := range j.Tasks {
		if task.Return != nil {
			return task.Return.Error()
		}
	}

	return code
}

// Returns true if the build was successful.
func (j *Job) Success() bool {
	if j.Status() == "0" {
		return true
	}
	return false
}

// Add a task to the job
func (j *Job) AddTask(t Task) {
	j.Tasks = append(j.Tasks, t)
}

// Returns either the exit code of the triggered command or 0 if the command
// finished successfully.
func (t *Task) Status() string {
	code := "success"

	if t.Return != nil {
		return t.Return.Error()
	}

	return code
}
