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
