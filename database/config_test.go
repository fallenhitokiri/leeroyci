package database

import (
	"testing"
)

func TestScheme(t *testing.T) {
	c := Config{
		URL: "https://foo:8080",
	}

	if c.Scheme() != "https" {
		t.Error("Wrong scheme", c.Scheme())
	}
}

func TestHost(t *testing.T) {
	c := Config{
		URL: "https://foo:8080",
	}

	if c.Host() != "foo:8080" {
		t.Error("Wrong host", c.Host())
	}
}
