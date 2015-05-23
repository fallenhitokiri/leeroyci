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
