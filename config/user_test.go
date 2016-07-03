package config

import "testing"

func TestListUsers(t *testing.T) {
	cfg = &Config{}

	if len(ListUsers()) != 0 {
		t.Error("Users exist, but none was added")
	}

	cfg.Users = append(cfg.Users, &User{})

	if len(ListUsers()) != 1 {
		t.Error("User was added but not returned")
	}
}

func TestGetUserByEmail(t *testing.T) {
	cfg = &Config{}
	cfg.Users = append(cfg.Users, &User{Email: "foo@bar.tld"})

	if _, err := GetUserByEmail("bar@foo.tld"); err == nil {
		t.Error("Found a user that does not exist")
	}

	if _, err := GetUserByEmail("foo@bar.tld"); err != nil {
		t.Error("Did not find the user")
	}
}

func TestGetUserBySessionKey(t *testing.T) {
	cfg = &Config{}
	session := &SessionKey{Key: "foo"}
	user := &User{}
	user.Sessions = append(user.Sessions, session)
	cfg.Users = append(cfg.Users, user)

	if _, err := GetUserBySessionKey("bar"); err == nil {
		t.Error("Found a user that does not exist")
	}

	if _, err := GetUserBySessionKey("foo"); err != nil {
		t.Error("Did not find the user")
	}
}

func TestGetUserByAPIKey(t *testing.T) {
	cfg = &Config{}
	api := &APIKey{Key: "foo"}
	user := &User{}
	user.APIKeys = append(user.APIKeys, api)
	cfg.Users = append(cfg.Users, user)

	if _, err := GetUserByAPIKey("bar"); err == nil {
		t.Error("Found a user that does not exist")
	}

	if _, err := GetUserByAPIKey("foo"); err != nil {
		t.Error("Did not find the user")
	}
}

func TestGetUserByID(t *testing.T) {
	cfg = &Config{}
	cfg.Users = append(cfg.Users, &User{ID: "foo"})

	if _, err := GetUserByID("bar"); err == nil {
		t.Error("Found a user that does not exist")
	}

	if _, err := GetUserByID("foo"); err != nil {
		t.Error("Did not find the user")
	}
}

func TestCreateUser(t *testing.T) {
	cfg = &Config{}
	user := &User{}
	if err := CreateUser(user); err.Error() != errorMissingValue {
		t.Error("No missing value error")
	}

	user.Email = "foo@bar.tld"
	if err := CreateUser(user); err.Error() != errorMissingValue {
		t.Error("No missing value error")
	}

	user.Password = "foo"
	if err := CreateUser(user); err.Error() != errorNoConfigPath {
		t.Error(err)
	}

	u := cfg.Users[0]
	if u.Password == "foo" {
		t.Error("Password not hashed")
	}

	if u.ID == "" {
		t.Error("No UUID")
	}
}

func TestUpdatePassword(t *testing.T) {
	user := &User{Password: "foo"}
	user.UpdatePassword("foo")
	if user.Password == "foo" {
		t.Error("Password was not hashed")
	}
}

func TestDelete(t *testing.T) {
	cfg = &Config{}
	user := &User{ID: "foo"}
	cfg.Users = append(cfg.Users, user)
	user.Delete()
	if len(cfg.Users) != 0 {
		t.Error("User not deleted")
	}
}

func TestNewSession(t *testing.T) {
	cfg = &Config{inMemory: true}
	user := &User{ID: "foo", Email: "foo@bar.tld"}
	cfg.Users = append(cfg.Users, user)

	key, err := user.NewSession()

	if user.Sessions[0].Key != key {
		t.Error("Wrong session key", key)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestNewAccessKey(t *testing.T) {
	cfg = &Config{inMemory: true}
	user := &User{ID: "foo", Email: "foo@bar.tld"}
	cfg.Users = append(cfg.Users, user)

	key, err := user.NewAccessKey()

	if user.APIKeys[0].Key != key {
		t.Error("Wrong session key", key)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestHachPassword(t *testing.T) {
	cfg = &Config{inMemory: true}
	user := &User{
		Email:    "foo@bar.tld",
		Password: "foo",
	}
	err := CreateUser(user)

	if err != nil {
		t.Error(err)
	}

	valid := user.CheckPassword("foo")
	invalid := user.CheckPassword("bar")

	if valid == false {
		t.Error("foo considered invalid")
	}

	if invalid == true {
		t.Error("bar considered valid")
	}
}
