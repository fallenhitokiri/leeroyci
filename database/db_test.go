package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_URL", "sqlite3 :memory:")
	NewDatabase()

	i := m.Run()

	os.Unsetenv("DATABASE_URL")
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

	if s != ":memory:" {
		t.Error("Wrong connection string", s)
	}
}
