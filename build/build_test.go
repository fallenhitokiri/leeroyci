package build

import (
	"ironman/config"
	"ironman/logging"
	"testing"
)

func TestCall(t *testing.T) {
	// TODO: find a way to make this work for everyone
	out, _ := call("ls", "-a", "/dev/null")

	if out != "/dev/null\n" {
		t.Error("wrong output", out)
	}
}

func TestRun(t *testing.T) {
	b := logging.Buildlog{}
	b.Done = make(chan bool, 5)

	cc := config.Command{
		Name:    "cmd",
		Execute: "ls",
	}

	cr := config.Repository{
		URL:      "http://test.tld",
		Commands: []config.Command{cc, cc},
	}

	c := config.Config{
		Repositories: []config.Repository{cr},
	}

	j := logging.Job{
		URL:    "http://test.tld",
		Branch: "branch",
		Commit: "commit",
		Name:   "name",
		Email:  "email",
	}

	run(j, &c, &b)

	if len(b.Jobs) != 1 {
		t.Error("Wrong number of jobs", len(b.Jobs))
	}

	if len(b.Jobs[0].Tasks) != 2 {
		t.Error("Wrong number of tasks", len(b.Jobs[0].Tasks))
	}
}
