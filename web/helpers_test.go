package web

import (
	//"leeroy/logging"
	"testing"
)

func TestSplitRepo(t *testing.T) {
	p := "/status/repo/68747470733a2f2f6769746875622e636f6d2f66616c6c656e6869746f6b6972692f7075736874657374/"
	r := splitFirst(p)

	if r != "https://github.com/fallenhitokiri/pushtest" {
		t.Error("Wrong repo", r)
	}
}

func TestSplitBranch(t *testing.T) {
	p := "/status/repo/a/foo/"
	b := splitSecond(p)

	if b != "foo" {
		t.Error("Wrong repo", b)
	}
}

/*func TestResponseFormat(t *testing.T) {
	f := responseFormat("")

	if f != "html" {
		t.Error("Wrong format", f)
	}

	f = responseFormat("JSON")

	if f != "json" {
		t.Error("Wrong format", f)
	}
}

func TestPaginatedJobs(t *testing.T) {
	j := []*logging.Job{
		&logging.Job{
			Identifier: "1",
		},
		&logging.Job{
			Identifier: "2",
		},
		&logging.Job{
			Identifier: "3",
		},
		&logging.Job{
			Identifier: "4",
		},
		&logging.Job{
			Identifier: "5",
		},
		&logging.Job{
			Identifier: "6",
		},
		&logging.Job{
			Identifier: "7",
		},
		&logging.Job{
			Identifier: "8",
		},
		&logging.Job{
			Identifier: "9",
		},
		&logging.Job{
			Identifier: "10",
		},
		&logging.Job{
			Identifier: "11",
		},
		&logging.Job{
			Identifier: "12",
		},
		&logging.Job{
			Identifier: "13",
		},
		&logging.Job{
			Identifier: "14",
		},
		&logging.Job{
			Identifier: "15",
		},
	}

	jobs, next := paginatedJobs(j, "4")

	if next != 14 {
		t.Error("Wrong next: ", next)
	}

	if len(jobs) != 10 {
		t.Error("Wrong number of jobs: ", len(jobs))
	}

	if jobs[0].Identifier != "4" {
		t.Error("Wrong job identifier: ", jobs[0].Identifier)
	}
}

func TestPaginateOutOfRange(t *testing.T) {
	j := []*logging.Job{
		&logging.Job{
			Identifier: "1",
		},
	}

	jobs, next := paginatedJobs(j, "4")

	if next != 0 {
		t.Error("Wrong next: ", next)
	}

	if len(jobs) != 1 {
		t.Error("Wrong number of jobs: ", len(jobs))
	}
}

func TestPaginateEnd(t *testing.T) {
	j := []*logging.Job{
		&logging.Job{
			Identifier: "1",
		},
		&logging.Job{
			Identifier: "2",
		},
		&logging.Job{
			Identifier: "3",
		},
		&logging.Job{
			Identifier: "4",
		},
		&logging.Job{
			Identifier: "5",
		},
		&logging.Job{
			Identifier: "6",
		},
		&logging.Job{
			Identifier: "7",
		},
		&logging.Job{
			Identifier: "8",
		},
		&logging.Job{
			Identifier: "9",
		},
		&logging.Job{
			Identifier: "10",
		},
		&logging.Job{
			Identifier: "11",
		},
		&logging.Job{
			Identifier: "12",
		},
		&logging.Job{
			Identifier: "13",
		},
		&logging.Job{
			Identifier: "14",
		},
		&logging.Job{
			Identifier: "15",
		},
	}

	jobs, next := paginatedJobs(j, "10")

	if next != 0 {
		t.Error("Wrong next: ", next)
	}

	if len(jobs) != 6 {
		t.Error("Wrong number of jobs: ", len(jobs))
	}
}

func TestPaginateOutOfRangeZero(t *testing.T) {
	j := []*logging.Job{
		&logging.Job{
			Identifier: "1",
		},
	}

	jobs, next := paginatedJobs(j, "0")

	if next != 0 {
		t.Error("Wrong next: ", next)
	}

	if len(jobs) != 1 {
		t.Error("Wrong number of jobs: ", len(jobs))
	}
}*/
