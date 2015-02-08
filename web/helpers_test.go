package web

import (
	"leeroy/logging"
	"net/http"
	"testing"
)

func TestSplitFirst(t *testing.T) {
	p := "/status/repo/68747470733a2f2f6769746875622e636f6d2f66616c6c656e6869746f6b6972692f7075736874657374/"
	r := splitFirst(p)

	if r != "https://github.com/fallenhitokiri/pushtest" {
		t.Error("Wrong repo", r)
	}
}

func TestSplitSecond(t *testing.T) {
	p := "/status/repo/a/foo/"
	b := splitSecond(p)

	if b != "foo" {
		t.Error("Wrong repo", b)
	}
}

func TestPaginateGetPrevious(t *testing.T) {
	if r := paginateGetPrevious(0); r != "" {
		t.Error("Wrong previous: ", r)
	}

	// negative previous
	if r := paginateGetPrevious(1); r != "0" {
		t.Error("Wrong previous: ", r)
	}

	// start 50 -> first element = 49
	if r := paginateGetPrevious(49); r != "40" {
		t.Error("Wrong previous: ", r)
	}
}

func TestPaginateGetNext(t *testing.T) {
	if r := paginateGetNext(0, 100); r != "10" {
		t.Error("Wrong next: ", r)
	}

	if r := paginateGetNext(9, 100); r != "20" {
		t.Error("Wrong next: ", r)
	}

	if r := paginateGetNext(50, 10); r != "" {
		t.Error("Wrong next: ", r)
	}
}

func TestPaginateGetFirst(t *testing.T) {
	if r := paginateGetFirst("10", 50); r != 9 {
		t.Error("Wrong first: ", r)
	}

	if r := paginateGetFirst("0", 50); r != 0 {
		t.Error("Wrong first: ", r)
	}

	if r := paginateGetFirst("10", 50); r != 9 {
		t.Error("Wrong first: ", r)
	}
}

func TestPaginateGetLast(t *testing.T) {
	if r := paginateGetLast(10, 50); r != 20 {
		t.Error("Wrong last: ", r)
	}

	if r := paginateGetLast(50, 10); r != 10 {
		t.Error("Wrong last: ", r)
	}
}

func TestGetParameter(t *testing.T) {
	r, _ := http.NewRequest("Get", "127.0.0.1", nil)
	r.URL.RawQuery = "key=value"

	if v := getParameter(r, "key", "fail"); v != "value" {
		t.Error("Wrong value: ", v)
	}

	if v := getParameter(r, "foo", "bar"); v != "bar" {
		t.Error("Wrong value: ", v)
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

	jobs, _, _ := paginatedJobs(j, "4")

	if len(jobs) != 10 {
		t.Error("Wrong number of jobs: ", len(jobs))
	}

	if jobs[0].Identifier != "4" {
		t.Error("Wrong job identifier: ", jobs[0].Identifier)
	}
}
