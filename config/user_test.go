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

func TestGetUser(t *testing.T) {
	c := Config{
		Users: []*User{
			&User{
				Email:    "foo",
				Name:     "a",
				Password: "x",
			},
		},
	}

	_, err := c.GetUser("foo")

	if err != nil {
		t.Error("Could not find user, but should exist")
	}

	_, err = c.GetUser("bar")

	if err == nil {
		t.Error("Found a user, but user should not exist")
	}
}

func TestCreateUser(t *testing.T) {
	c := Config{}

	c.CreateUser("foo", "a", "x", false)

	u := c.Users[0]

	if u.Email != "foo" {
		t.Error("Wrong email address, ", u.Email)
	}

	if u.Password == "x" {
		t.Error("Password was not encrypted")
	}
}

func TestUpdateUser(t *testing.T) {
	c := Config{
		Users: []*User{
			&User{
				Email:    "foo",
				Name:     "a",
				Password: "x",
			},
		},
	}

	c.UpdateUser("foo", "name", "x", false)

	u := c.Users[0]

	if u.Name != "name" {
		t.Error("Wrong name ", u.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	c := Config{
		Users: []*User{
			&User{
				Email:    "foo",
				Name:     "a",
				Password: "x",
			},
			&User{
				Email: "bar",

				Name:     "b",
				Password: "x",
			},
			&User{
				Email: "baz",

				Name:     "b",
				Password: "x",
			},
			&User{
				Email: "foobar",

				Name:     "b",
				Password: "x",
			},
		},
	}

	c.DeleteUser("baz")

	_, err := c.GetUser("baz")

	if err == nil {
		t.Error("Found the deleted user.")
	}

	err = c.DeleteUser("asdf")

	if err == nil {
		t.Error("User not found but no error returned.")
	}
}
