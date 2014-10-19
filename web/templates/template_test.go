package templates

import (
	"testing"
)

func TestFullPath(t *testing.T) {
	fqp := fullPath("foo", "~/bar")

	if fqp != "~/bar/foo.html" {
		t.Error("Wrong FQP", fqp)
	}
}
