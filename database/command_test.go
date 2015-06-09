package database

import (
	"testing"
)

func TestCommandCRUD(t *testing.T) {
	repo := CreateRepository("", "", "", false, false, false)

	add, _ := AddCommand(repo, "name", "execute", "branch", CommandKindBuild)
	get1 := GetCommand(add.ID)
	get1.Update("name", "kind", "branch", CommandKindDeploy)
	get2 := GetCommand(add.ID)
	get1.Delete()
	get3 := GetCommand(add.ID)

	if get1.ID != get2.ID {
		t.Error("ID mismatch")
	}

	if get1.Kind != "kind" {
		t.Error("Kind not updated")
	}

	if get3.ID != 0 {
		t.Error("Not deleted")
	}
}
