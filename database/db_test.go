package database

import (
	"os"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	os.Setenv("DATABASE_URL", "sqlite3 :memory:")

	NewDatabase()

	err := db.DB().Ping()

	if err != nil {
		t.Error("No database connection")
	}

	os.Unsetenv("DATABASE_URL")
}

func TestEnvURL(t *testing.T) {
	os.Setenv("DATABASE_URL", "foo bar baz")

	d, s := envURL()

	if d != "foo" {
		t.Error("Wrong driver", d)
	}

	if s != "bar baz" {
		t.Error("Wrong connection string", s)
	}

	os.Unsetenv("DATABASE_URL")
}
