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

func TestListUsers(t *testing.T) {
	db.Exec("DELETE FROM users WHERE id > 0")
	CreateUser("foo1@bar.tld", "foo", "bar", "adsf", false)
	CreateUser("foo2@bar.tld", "foo", "bar", "adsf", false)
	CreateUser("foo3@bar.tld", "foo", "bar", "adsf", false)

	users := ListUsers()

	if len(users) != 3 {
		t.Error("Wrong user count", len(users))
	}
}

func TestGetUserBySession(t *testing.T) {
	db.Exec("DELETE FROM users WHERE id > 0")
	user, _ := CreateUser("foo1@bar.tld", "foo", "bar", "adsf", false)
	user.NewSession()

	_, err := GetUserBySession("asdf")

	if err == nil {
		t.Error("Got user with incorrect session.")
	}

	got, _ := GetUserBySession(user.Session)

	if got.ID != user.ID {
		t.Error("Got wrong user")
	}
}

func TestGetUserbyID(t *testing.T) {
	db.Exec("DELETE FROM users WHERE id > 0")
	user, _ := CreateUser("foo1@bar.tld", "foo", "bar", "adsf", false)

	_, err := GetUserByID(9999999)

	if err == nil {
		t.Error("Found user with impossible ID")
	}

	_, err = GetUserByID(user.ID)

	if err != nil {
		t.Error(err)
	}
}
