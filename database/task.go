// Package database provides a wrapper between the database and stucts
package database

// Task defines a task which is executed when a branch is pushed.
type Task struct {
	ID     int64
	Return string
	Output string

	Command   Command
	CommandID int64

	Job   Job
	JobID int64
}

// AddTask adds a new task.
func AddTask(command *Command, job *Job, ret, out string) *Task {
	task := Task{
		Return:  ret,
		Output:  out,
		Command: *command,
		Job:     *job,
	}

	db.Save(&task)

	return &task
}

// Status returns either the exit code of the triggered command or 0
// 0 if the command finished successfully.
func (t *Task) Status() string {
	code := "success"

	if t.Return != "" {
		return t.Return
	}

	return code
}
