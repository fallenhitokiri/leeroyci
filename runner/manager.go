// Package runner runs all tasks for all commands associated with a repository.
package runner

import (
	"sync"

	"github.com/fallenhitokiri/leeroyci/database"
)

var manager taskManager

type taskManager struct {
	progress map[int]string
	mutex    sync.Mutex
}

func newTaskManager() {
	config := database.GetConfig()

	manager = taskManager{
		progress: map[int]string{},
	}

	for index := 1; index <= config.Parallel; index++ {
		manager.progress[index] = ""
	}
}

func (t *taskManager) getTaskID(repository, branch string) int {
	ident := repository + branch

	t.mutex.Lock()
	defer t.mutex.Unlock()

	// only one task per branch
	for _, val := range t.progress {
		if val == ident {
			return 0
		}
	}

	for id, val := range t.progress {
		if val == "" {
			t.progress[id] = ident
			return id
		}
	}

	return 0
}

func (t *taskManager) doneWithID(id int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.progress[id] = ""
}
