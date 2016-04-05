package runner

import (
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestNewTaskManager(t *testing.T) {
	database.NewInMemoryDatabase()
	database.AddConfig("secret", "url", "cert", "key", 5)
	newTaskManager()

	if manager.progress[1] != "" {
		t.Error("Wrong value at 1", manager.progress[1])
	}

	if manager.progress[2] != "" {
		t.Error("Wrong value at 2", manager.progress[2])
	}

	if manager.progress[3] != "" {
		t.Error("Wrong value at 3", manager.progress[3])
	}

	if manager.progress[4] != "" {
		t.Error("Wrong value at 4", manager.progress[4])
	}

	if manager.progress[5] != "" {
		t.Error("Wrong value at 5", manager.progress[5])
	}
}

func TestGetTaskID(t *testing.T) {
	database.NewInMemoryDatabase()
	database.AddConfig("secret", "url", "cert", "key", 2)
	newTaskManager()

	tID1 := manager.getTaskID("foo", "bar")

	if tID1 == 0 {
		t.Error("Got no ID")
	}

	if manager.getTaskID("foo", "bar") != 0 {
		t.Error("Got a second ID")
	}

	tID2 := manager.getTaskID("foo", "baz")

	if tID2 == 0 {
		t.Error("Got no ID")
	}

	if tID1 == tID2 {
		t.Error("Got same ID as tID1")
	}

	if manager.getTaskID("foo", "baz") != 0 {
		t.Error("Got a second ID")
	}

	tID3 := manager.getTaskID("foo", "zab")

	if tID3 != 0 {
		t.Error("Got an ID, no free IDs")
	}

	manager.doneWithID(tID1)

	tID3 = manager.getTaskID("foo", "zab")

	if tID3 == 0 {
		t.Error("Got no ID, but one is free")
		t.Error(manager.progress)
	}

	if tID1 != tID3 {
		t.Error("Got a wrong ID")
	}
}
