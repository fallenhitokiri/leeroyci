package logging

import (
	"errors"
	"testing"
)

func TestStatus(t *testing.T) {
	j := Job{}

	if j.Status() != "0" {
		t.Error("Wrong status", j.Status())
	}

	task := Task{
		Return: errors.New("foo"),
	}
	j.Tasks = append(j.Tasks, task)

	if j.Status() == "0" {
		t.Error("Wrong status", j.Status())
	}
}

func TestSuccess(t *testing.T) {
	j := Job{}

	if j.Success() == false {
		t.Error("Returned error for successful build")
	}

	task := Task{
		Return: errors.New("foo"),
	}
	j.Tasks = append(j.Tasks, task)

	if j.Success() == true {
		t.Error("Returned no error for failed build")
	}
}

func TestAddTask(t *testing.T) {
	j := Job{}
	task := Task{}

	j.AddTask(task)

	if len(j.Tasks) != 1 {
		t.Error("Wrong length of task list", len(j.Tasks))
	}
}

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
