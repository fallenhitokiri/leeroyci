package database

import (
	"testing"
)

func TestServer(t *testing.T) {
	m := &MailServer{}
	m.Host = "foo"
	m.Port = 1234

	s := m.Server()

	if s != "foo:1234" {
		t.Error("Wrong server", s)
	}
}
