// Package database provides a wrapper between the database and stucts
package database

type Task struct {
	ID      int64
	Command string
	Return  string
	Output  string
}

// Returns either the exit code of the triggered command or 0 if the command
// finished successfully.
func (t *Task) Status() string {
	code := "success"

	if t.Return != "" {
		return t.Return
	}

	return code
}
