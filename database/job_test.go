package database

import (
	"testing"
)

func TestJobStatus(t *testing.T) {
	j := Job{}

	if j.Status() != "0" {
		t.Error("Wrong status", j.Status())
	}

	task := Task{
		Return: "foo",
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
		Return: "foo",
	}
	j.Tasks = append(j.Tasks, task)

	if j.Success() == true {
		t.Error("Returned no error for failed build")
	}
}

func TestAdd(t *testing.T) {
	j := Job{}
	task := Task{}

	j.Add(task)

	if len(j.Tasks) != 1 {
		t.Error("Wrong length of task list", len(j.Tasks))
	}
}

func TestDeploySuccess(t *testing.T) {
	j := Job{}

	if j.DeploySuccess() != false {
		t.Error("DeploySuccess is not false but should be.")
	}

	ta := Task{
		Return: "",
	}

	j.Deployed = &ta

	if j.DeploySuccess() != true {
		t.Error("DeploySuccess is not true but should be.")
	}

	j.Deployed.Return = "1"

	if j.DeploySuccess() != false {
		t.Error("DeploySuccess is not false but should be.")
	}
}
