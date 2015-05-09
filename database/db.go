// Package database provides a wrapper between the database and stucts
package database

import (
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"           // Postgres driver
	_ "github.com/mattn/go-sqlite3" // Sqlite3 driver (used for testing)
)

var db gorm.DB

// NewDatabase established a database connection and stores it in `db`.
func NewDatabase() error {
	d, s := envURL()

	db, err := gorm.Open(d, s)

	if err != nil {
		return err
	}

	db.DB()

	db.AutoMigrate(
		&Command{},
		&Config{},
		&Deploy{},
		&Job{},
		&MailServer{},
		&Notify{},
		&Repository{},
		&Task{},
		&User{},
	)

	return nil
}

// envURL returns the database type and connection settings read from the environment
// variable `DATABASE_URL`.
//
// Format: "SQLDRIVER connection settings for driver"
func envURL() (string, string) {
	dbURL := os.Getenv("DATABASE_URL")

	s := strings.SplitN(dbURL, " ", 2)

	if len(s) != 2 {
		panic("Invalid DATABASE_URL")
	}

	return s[0], s[1]
}
