package web

import (
	"leeroy/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	j := []*database.Job{
		&database.Job{},
	}

	Status(rw, req)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestRepo(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/repo/75726c", nil)
	j := []*database.Job{
		&database.Job{},
	}

	Repo(rw, req)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestBranch(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/branch/75726c/foo", nil)
	j := []*database.Job{
		&database.Job{},
	}

	Branch(rw, req)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestCommit(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/branch/75726c/foo", nil)
	j := []*database.Job{
		&database.Job{},
	}

	Commit(rw, req)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}
}

func TestBadge(t *testing.T) {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1/status/badge/75726c/foo", nil)
	j := []*database.Job{
		&database.Job{},
	}

	Badge(rw, req)

	if rw.Code != 200 {
		t.Error("Wrong status code: ", 200)
	}

	if ct, ok := rw.Header()["Content-Type"]; ok == true {
		if ct[0] != "image/svg+xml" {
			t.Error("Wrong content type: ", ct[0])
		}
	} else {
		t.Error("Content Type not found")
	}
}
