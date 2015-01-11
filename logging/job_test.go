package logging

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

func TestMD5(t *testing.T) {
	j := Job{
		URL: "https://github.com/fallenhitokiri/pushtest",
	}

	hex := "68747470733a2f2f6769746875622e636f6d2f66616c6c656e6869746f6b6972692f7075736874657374"

	if j.Hex() != hex {
		t.Error("Wrong hex", j.Hex())
	}
}

func TestStatusURL(t *testing.T) {
	j := Job{
		URL:    "https://github.com/fallenhitokiri/pushtest",
		Commit: "foo",
	}

	u := j.StatusURL("asdf/")
	e := "asdf/status/commit/68747470733a2f2f6769746875622e636f6d2f66616c6c656e6869746f6b6972692f7075736874657374/foo"

	if u != e {
		t.Error("Wrong URL", u)
	}
}

func TestDeploySuccess(t *testing.T) {
	j := Job{
		Deployed: "0",
	}

	if j.DeploySuccess() != true {
		t.Error("DeploySuccess is not true but should be.")
	}

	j.Deployed = "1"

	if j.DeploySuccess() != false {
		t.Error("DeploySuccess is not false but should be.")
	}
}
