// Package database provides a wrapper between the database and stucts
package database

// Task defines a task which is executed when a branch is pushed.
type Task struct {
	ID      int64
	Command string
	Return  string
	Output  string
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
