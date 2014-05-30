package logging

import (
	"testing"
)

func TestTaskStatus(t *testing.T) {
	task := Task{}

	if task.Status() != "success" {
		t.Error("Wrong status", task.Status())
	}

	task.Return = "foo"

	if task.Status() == "success" {
		t.Error("Wrong status", task.Status())
	}
}
