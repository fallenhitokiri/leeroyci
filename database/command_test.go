package database

import (
	"testing"
)

func TestCommandCRUD(t *testing.T) {
	repo, _ := CreateRepository("", "", "", false, false)

	add, _ := CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)
	get1, _ := GetCommand(add.ID)
	get1.Update("name", "kind", "branch", CommandKindDeploy)
	get2, _ := GetCommand(add.ID)
	get1.Delete()
	get3, _ := GetCommand(add.ID)

	if get1.ID != get2.ID {
		t.Error("ID mismatch", get1.ID, get2.ID)
	}

	if get1.Kind != "kind" {
		t.Error("Kind not updated")
	}

	if get3.ID != 0 {
		t.Error("Not deleted")
	}
}

func TestWrongKind(t *testing.T) {
	repo, _ := CreateRepository("", "", "", false, false)
	_, err := CreateCommand(repo, "name", "execute", "branch", "baz")

	if err == nil {
		t.Error("No error")
	}
}

func TestCommandLogPassed(t *testing.T) {
	log := CommandLog{Return: ""}

	if log.Passed() == false {
		t.Error("Passed not true")
	}

	log.Return = "1"

	if log.Passed() == true {
		t.Error("Passed not false")
	}
}

func TestCommandLogCR(t *testing.T) {
	repo, _ := CreateRepository("asdf", "", "", false, false)
	com, _ := CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)
	job := CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	CreateCommandLog(com, job, "", "foo")

	got := GetCommandLogsForJob(job.ID)

	if len(got) != 1 {
		t.Error("Wrong length", len(got))
	}
}
