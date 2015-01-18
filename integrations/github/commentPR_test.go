package github

import (
	"leeroy/logging"
	"strings"
	"testing"
)

func TestNewComment(t *testing.T) {
	j := logging.Job{}
	c := newComment(&j, "")

	if strings.Contains(c.Body, "successful") == false {
		t.Error("wrong comment - not successful")
	}

	ta := logging.Task{
		Return: "foo",
	}
	j.Tasks = append(j.Tasks, ta)

	c = newComment(&j, "")

	if strings.Contains(c.Body, "failed") == false {
		t.Error("wrong comment - not failed")
	}
}
