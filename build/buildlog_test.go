package build

import (
	"testing"
)

func TestBuildlogAdd(t *testing.T) {
	log := Buildlog{}
	log.Add("url", "branch", "pusher", "output", 1)

	if len(log.Jobs) != 1 {
		t.Error("build not added")
	}

	j := log.Jobs[0]

	if j.URL != "url" {
		t.Error("wrong URL", j.URL)
	}

	if j.Branch != "branch" {
		t.Error("wrong Branch", j.Branch)
	}

	if j.ReturnCode != 1 {
		t.Error("wrong ReturnCode", j.ReturnCode)
	}

	if j.Output != "output" {
		t.Error("wrong Output", j.Output)
	}

	if j.Pusher != "pusher" {
		t.Error("wrong Pusher", j.Pusher)
	}
}
