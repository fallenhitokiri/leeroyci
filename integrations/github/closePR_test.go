package github

import (
	"testing"
)

func TestNewUpdate(t *testing.T) {
	u := newUpdate()

	if u.State != "closed" {
		t.Error("Update state is not closed.")
	}
}
