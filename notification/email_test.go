package notification

import (
	"testing"
)

func TestAddHeaders(t *testing.T) {
	message := addHeaders("foo", "bar", "bla", "baz")

	if len(message) != 137 {
		t.Error("Message got the wrong length", len(message))
	}
}
