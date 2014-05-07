// Each command ran becomes a task.
package logging

type Task struct {
	Command string
	Return  error
	Output  string
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
