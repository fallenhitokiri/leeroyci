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

func TestMailServerCRUD(t *testing.T) {
	_ = AddMailServer("host", "sender", "user", "password", 1234)
	get1 := GetMailServer()
	updated := UpdateMailServer("host", "sender", "user", "password", 4321)
	get2 := GetMailServer()
	DeleteMailServer()
	get3 := GetMailServer()

	if get1.ID != get2.ID || updated.ID != get2.ID {
		t.Error("ID mismatch")
	}

	if get1.Port == get2.Port {
		t.Error("Port not updated")
	}

	if get3.ID != 0 {
		t.Error("Not deleted")
	}
}
