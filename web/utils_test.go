package web

import (
	"testing"
)

func TestStringToIntPanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("No panic when passing in character string")
		}
	}()

	stringToInt("asdf")
}

func TestStringToInt(t *testing.T) {
	num := stringToInt("123")

	if num != 123 {
		t.Error("Wrong number", num)
	}
}
