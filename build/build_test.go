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
	r := database.AddRepository("foo", "http://test.tld", "", false, false, false)
	c1 := database.AddCommand(r, "cmd", "ls")
	c2 := database.AddCommand(r, "cmd", "ls")

	j := database.Job{
		URL:    "http://test.tld",
		Branch: "branch",
		Commit: "commit",
		Name:   "name",
		Email:  "email",
	}

	run(j)

	jobs := database.GetAllJobs()

	if len(jobs) != 1 {
		t.Error("Wrong number of jobs", len(logging.BUILDLOG.Jobs))
	}

	if len(jobs[0].Tasks) != 2 {
		t.Error("Wrong number of tasks", len(logging.BUILDLOG.Jobs[0].Tasks))
	}
}
