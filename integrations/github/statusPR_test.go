package github

import (
	"testing"
)

func TestNewStatus(t *testing.T) {
	s := newStatus(statusSuccess, "foo")

	if s.State != "success" {
		t.Error("Wrong state")
	}
}
