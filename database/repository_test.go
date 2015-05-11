package database

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	r := Repository{
		URL: "foo",
	}

	i := r.Identifier()

	if i != "foo" {
		t.Error("Wrong identifier", i)
	}

	r.Name = "bar"

	i = r.Identifier()

	if i != "bar" {
		t.Error("Wrong identifier", i)
	}
}
