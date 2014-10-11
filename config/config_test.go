package config

import (
	"encoding/json"
	"testing"
)

var payload = []byte(`{"Secret":"asdf","Repositories":[{"URL":"foo","Commands":[{"Name":"command1","Execute":"exec1"},{"Name":"command2","Execute":"exec2 --wusa"}],"Notify":[{"Name":"name 1","Email":"email1"},{"Name":"name 2","Email":"email2"}]}]}`)

func TestConfigLoad(t *testing.T) {
	var c Config

	json.Unmarshal(payload, &c)

	if c.Secret != "asdf" {
		t.Error("Wrong secret", c.Secret)
	}

	// TODO: test all values
}

func TestConfigForRepo(t *testing.T) {
	var c Config

	json.Unmarshal(payload, &c)

	r, err := c.ConfigForRepo("foo")

	if err != nil {
		t.Error(err)
	}

	if len(r.Commands) != 2 {
		t.Error("Wrong number of commands", r.Commands)
	}

	r, err = c.ConfigForRepo("bazbaz")

	if err == nil {
		t.Error("Error is nil")
	}
}

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

func TestIdentifier(t *testing.T) {
	r1 := Repository{
		Name: "foobar",
		URL:  "baz",
	}
	r2 := Repository{
		URL: "baz",
	}

	if r1.Identifier() != "foobar" {
		t.Error("Wrong identifier", r1.Identifier())
	}

	if r2.Identifier() != "baz" {
		t.Error("Wrong identifier", r2.Identifier())
	}
}
