package database

import (
	"strconv"
	"testing"
)

func TestCommandCRUD(t *testing.T) {
	repo, _ := CreateRepository("", "", "", false, false, false)

	add, _ := CreateCommand(repo, "name", "execute", "branch", CommandKindBuild)
	get1, _ := GetCommand(strconv.FormatInt(add.ID, 10))
	get1.Update("name", "kind", "branch", CommandKindDeploy)
	get2, _ := GetCommand(strconv.FormatInt(add.ID, 10))
	get1.Delete()
	get3, _ := GetCommand(strconv.FormatInt(add.ID, 10))

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
