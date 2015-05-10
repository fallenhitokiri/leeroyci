package build

import (
	"leeroy/database"
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
	cc := database.Command{
		Name:    "cmd",
		Execute: "ls",
	}

	cr := database.Repository{
		URL:      "http://test.tld",
		Commands: []database.Command{cc, cc},
	}

	j := logging.Job{
		URL:    "http://test.tld",
		Branch: "branch",
		Commit: "commit",
		Name:   "name",
		Email:  "email",
	}

	run(j)

	if len(logging.BUILDLOG.Jobs) != 1 {
		t.Error("Wrong number of jobs", len(logging.BUILDLOG.Jobs))
	}

	if len(logging.BUILDLOG.Jobs[0].Tasks) != 2 {
		t.Error("Wrong number of tasks", len(logging.BUILDLOG.Jobs[0].Tasks))
	}
}
