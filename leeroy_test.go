package main

import (
	"os"
	"testing"
)

func TestPort(t *testing.T) {
	existing := os.Getenv("PORT")
	os.Setenv("PORT", "")

	if port() != ":8082" {
		t.Error("Wrong port", port())
	}

	os.Setenv("PORT", "1234")

	if port() != ":1234" {
		t.Error("Wrong port", port())
	}

	os.Setenv("PORT", existing)
}
