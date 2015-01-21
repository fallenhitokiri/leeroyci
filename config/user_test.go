package config

import (
	"testing"
)

func TestHashComparePassword(t *testing.T) {
	pass := "asdf"

	h, err := hashPassword(pass)

	if err != nil {
		t.Error(err)
	}

	v := comparePassword(pass, h)

	if v != true {
		t.Error("password did not match")
	}

	v = comparePassword("foo", h)

	if v != false {
		t.Error("passwords did match")
	}
}
