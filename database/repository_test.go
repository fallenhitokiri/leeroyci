package database

import (
	"testing"
)

func TestAddRepositoryGetRepository(t *testing.T) {
	r1 := CreateRepository("foo", "bar", "accessKey", false, false, false)
	r2 := GetRepository("bar")
	r3 := UpdateRepository("baz", "bar", "accessKey", false, false, false)
	DeleteRepository("bar")
	r4 := GetRepository("bar")

	if r1.ID != r2.ID {
		t.Error("IDs do not match.")
	}

	if r3.Name == r2.Name {
		t.Error("Names are the same.")
	}

	if r4.ID == r1.ID || r4.ID != 0 {
		t.Error("Repository not deleted.")
	}
}

func TestJobs(t *testing.T) {
	r1 := CreateRepository("name", "url", "accessKey", false, false, false)
	r2 := CreateRepository("name2", "url2", "accessKey", false, false, false)

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
