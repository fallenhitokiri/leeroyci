package database

import (
	"testing"
)

func TestAddRepositoryGetRepository(t *testing.T) {
	r1 := AddRepository("foo", "bar", "accessKey", false, false, false)
	r2 := GetRepository("bar")

	if r1.ID != r2.ID {
		t.Error("IDs do not match.")
	}
}
