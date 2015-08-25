package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	NewDatabase("sqlite3", ":memory:")

	i := m.Run()

	os.Exit(i)
}

func TestNewDatabase(t *testing.T) {
	err := db.DB().Ping()

	if err != nil {
		t.Error("No database connection")
	}
}

func TestEnvURL(t *testing.T) {
	d, s := envURL()

	if d != "sqlite3" {
		t.Error("Wrong driver", d)
	}

	if s != "/Users/timo/tmp/leeory.sqlite3" {
		t.Error("Wrong connection string", s)
	}
}
