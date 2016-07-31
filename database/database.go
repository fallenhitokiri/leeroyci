package database

import "fmt"

var db *Database
var test string

func init() {
	fmt.Println(test)
	fmt.Println("init")
	test = "foo"
}

type Database interface {
	Config() (*Config, error)
	ConfigUpdate(cfg *Config) error

	UserList() ([]*User, error)
	UserByEmail(email string) (*User, error)
	UserBySessionKey(key string) (*User, error)
	UserByAPIKey(key string) (*User, error)
	UserByID(id string) (*User, error)
	UserCreate(user *User) error
	UserUpdatePassword(id string, password string) error
	UserDelete(user *User) error

	ProjectCreate(project *Project) error
	ProjectList() ([]*Project, error)
	ProjectByName(name string) (*Project, error)
	ProjectByAccessKey(key string) (*Project, error)

	MailServer() (*MailServer, error)
	MailServerUpdate(ms *MailServer) error
}
