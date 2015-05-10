package github

import (
	"leeroy/database"
	"strings"
	"testing"
)

func TestNewComment(t *testing.T) {
	j := database.Job{}
	c := newComment(&j, "")

	if strings.Contains(c.Body, "successful") == false {
		t.Error("wrong comment - not successful")
	}

	ta := database.Task{
		Return: "foo",
	}
	j.Tasks = append(j.Tasks, ta)

	c = newComment(&j, "")

	if strings.Contains(c.Body, "failed") == false {
		t.Error("wrong comment - not failed")
	}
}
