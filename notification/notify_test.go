package notification

import (
	"testing"
)

func TestKindSupported(t *testing.T) {
	s := kindSupported(KindBuild)

	if s == false {
		t.Error(s, "not supported, but should.")
	}

	s = kindSupported("foo")

	if s == true {
		t.Error(s, "supported, but should not.")
	}
}
