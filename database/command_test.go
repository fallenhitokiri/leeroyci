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
