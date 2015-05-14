package database

import (
	"testing"
)

func TestAddCommandGetCommand(t *testing.T) {
	repo := AddRepository("", "", "", false, false, false)
	com1, err := AddCommand(repo, "", "", "", CommandKindBuild)

	if err != nil {
		t.Error(err)
	}

	coms := GetCommands(repo, "", CommandKindBuild)

	if coms[0].ID != com1.ID {
		t.Error("ID mismatch")
	}
}

func TestAddCommandGetCommandDifferentKind(t *testing.T) {
	repo := AddRepository("", "", "", false, false, false)
	_, err := AddCommand(repo, "", "", "", CommandKindBuild)

	if err != nil {
		t.Error(err)
	}

	coms := GetCommands(repo, "", CommandKindTest)

	if len(coms) != 0 {
		t.Error("Wrong number of commands", len(coms))
	}
}

func TestCommandCRUD(t *testing.T) {
	repo := AddRepository("", "", "", false, false, false)

	add, _ := AddCommand(repo, "name", "execute", "branch", CommandKindBuild)
	get1 := GetCommand(add.ID)
	updated := UpdateCommand(add.ID, "name", "kind", "branch", CommandKindDeploy)
	get2 := GetCommand(add.ID)
	DeleteCommand(add.ID)
	get3 := GetCommand(add.ID)

	if get1.ID != get2.ID || updated.ID != get2.ID {
		t.Error("ID mismatch")
	}

	if get1.Kind == get2.Kind {
		t.Error("Kind not updated")
	}

	if get3.ID != 0 {
		t.Error("Not deleted")
	}
}
