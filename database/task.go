// Package database provides a wrapper between the database and stucts
package database

// Task defines a completed Command.
type Task struct {
	ID     int64
	Name   string // name of the command
	Return string
	Output string

	Job   Job
	JobID int64
}

// AddTask adds a new task.
func AddTask(command *Command, job *Job, ret, out string) *Task {
	task := Task{
		Name:   command.Name,
		Return: ret,
		Output: out,
		Job:    *job,
	}

	db.Save(&task)

	return &task
}

// Passed returns true if the task completed successfully.
func (t *Task) Passed() bool {
	if t.Return == "" {
		return true
	}

	return false
}
