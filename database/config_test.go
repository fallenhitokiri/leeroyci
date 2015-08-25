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

func TestConfigCRUD(t *testing.T) {
	AddConfig("secret", "url", "cert", "key")
	get1 := GetConfig()
	updated := UpdateConfig("secret2", "url", "cert", "key")
	get2 := GetConfig()
	DeleteConfig()
	get3 := GetConfig()

	if get1.ID != get2.ID || updated.ID != get2.ID {
		t.Error("ID mismatch")
	}

	if get1.Secret == get2.Secret {
		t.Error("Secret not updated")
	}

	if get3.ID != 0 {
		t.Error("Not deleted")
	}
}
