package database

import (
	"testing"
)

func TestHashComparePassword(t *testing.T) {
	pass := "asdf"

	h, err := hashPassword(pass)

	if err != nil {
		t.Error(err)
	}

	v := ComparePassword(pass, h)

	if v != true {
		t.Error("password did not match")
	}

	v = ComparePassword("foo", h)

	if v != false {
		t.Error("passwords did match")
	}
}

func TestCreateGetUpdateDelete(t *testing.T) {
	u, err := CreateUser("foo@bar.tld", "foo", "bar", "adsf", false)

	if err != nil {
		t.Error(err)
	}

	u2, err := GetUser("foo@bar.tld")

	if err != nil {
		t.Error(err)
	}

	if u.ID != u2.ID {
		t.Error("IDs do not match")
	}

	u3, err := u.Update("foo@bar.tld", "foo", "baz", "adsf", false)

	if err != nil {
		t.Error(err)
	}

	if u3.LastName == u2.LastName {
		t.Error("Name not updated")
	}

	u.Delete()

	_, err = GetUser("foo@bar.tld")

	if err == nil {
		t.Error("User found, should be deleted.")
	}
}
