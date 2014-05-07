package logging

import (
	"errors"
	"testing"
)

func TestTaskStatus(t *testing.T) {
	task := Task{}

	if task.Status() != "success" {
		t.Error("Wrong status", task.Status())
	}

	task.Return = errors.New("foo")

	if task.Status() == "success" {
		t.Error("Wrong status", task.Status())
	}
}
