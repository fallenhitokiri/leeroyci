package notification

import (
	"testing"
)

func TestBuildEmail(t *testing.T) {
	message := buildEmail("foo", "bar", "bla", "baz")

	if len(message) != 137 {
		t.Error("Message got the wrong length", len(message))
	}
}
