package database

import (
	"testing"
)

func TestAddRepositoryGetRepository(t *testing.T) {
	r1, _ := CreateRepository("foo", "bar", "accessKey", false, false)
	r2 := GetRepository("bar")
	r2.Update("baz", "bar", "accessKey", false, false)
	r2.Delete()
	r3 := GetRepository("bar")

	if r1.ID != r2.ID {
		t.Error("IDs do not match.")
	}

	if r2.Name == "bar" {
		t.Error("Names are the same.")
	}

	if r3.ID == r1.ID || r3.ID != 0 {
		t.Error("Repository not deleted.")
	}
}

func TestJobs(t *testing.T) {
	r1, _ := CreateRepository("name", "url", "accessKey", false, false)
	r2, _ := CreateRepository("name2", "url2", "accessKey", false, false)

	j1 := CreateJob(r1, "branch", "commit", "commitURL", "name", "email")
	j2 := CreateJob(r1, "branch2", "commit", "commitURL", "name", "email")
	CreateJob(r2, "branch3", "commit", "commitURL", "name", "email")

	j := r1.Jobs()

	if len(j) != 2 {
		t.Error("Wrong number of jobs", len(j))
	}

	for _, v := range j {
		if v.ID != j1.ID && v.ID != j2.ID {
			t.Error("Wrong job")
		}
	}
}

func TestAddCommandCommand(t *testing.T) {
	repo, _ := CreateRepository("", "", "", false, false)
	com1, err := CreateCommand(repo, "", "", "", CommandKindBuild)

	if err != nil {
		t.Error(err)
	}

	coms := repo.GetCommands("", CommandKindBuild)

	if coms[0].ID != com1.ID {
		t.Error("ID mismatch")
	}
}

func TestAddCommandGetCommandDifferentKind(t *testing.T) {
	repo, _ := CreateRepository("", "", "", false, false)
	_, err := CreateCommand(repo, "", "", "", CommandKindBuild)

	if err != nil {
		t.Error(err)
	}

	coms := repo.GetCommands("", CommandKindTest)

	if len(coms) != 0 {
		t.Error("Wrong number of commands", len(coms))
	}
}
