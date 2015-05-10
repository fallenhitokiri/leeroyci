package web

import (
	"leeroy/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewResponse(t *testing.T) {
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	r := newResponse(nil, req)

	if r.Format != formatHTML {
		t.Error("Wrong format: ", r.Format)
	}
}

func TestRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	r := newResponse(w, req)
	j := []*database.Job{
		&database.Job{
			Identifier: "foo",
		},
	}
	r.Context.Jobs = j

	r.renderJSON()

	if ct, ok := w.Header()["Content-Type"]; ok == true {
		if ct[0] != "application/json" {
			t.Error("Wrong content type: ", ct[0])
		}
	} else {
		t.Error("Content Type not found")
	}

	if strings.Contains(w.Body.String(), "foo") == false {
		t.Error("Wrong body: ", w.Body.String())
	}
}

func TestRenderHTML(t *testing.T) {
	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	r := newResponse(w, req)
	j := []*logging.Job{
		&logging.Job{
			Identifier: "foo",
		},
	}
	r.Context.Jobs = j

	r.renderHTML()

	if w.Code != 500 {
		t.Error("Wrong response code: ", w.Code)
	}

	r.Template = "status"
	r.TemplatePath = ""
	r.Writer = w2

	r.renderHTML()

	if w2.Code != 200 {
		t.Error("Wrong response code: ", w.Code)
	}
}

func TestRender(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	r := newResponse(w, req)
	r.Format = "foobar"

	r.render()

	if w.Code != 500 {
		t.Error("Wrong response code: ", w.Code)
	}
}
