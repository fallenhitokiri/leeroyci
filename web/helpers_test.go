package web

import (
	"net/url"
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

func TestResponseFormat(t *testing.T) {
	v := url.Values{}
	f := responseFormat(v)

	if f != "html" {
		t.Error("Wrong format", f)
	}

	v.Add("format", "JSON")
	f = responseFormat(v)

	if f != "json" {
		t.Error("Wrong format", f)
	}
}
