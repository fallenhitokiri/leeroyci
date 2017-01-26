// Package database provides a wrapper between the database and stucts
package database

import (
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"           // Postgres driver
	_ "github.com/mattn/go-sqlite3" // Sqlite3 driver (used for testing)
)

var db *gorm.DB

// Configured indicates if there is a valid configuration.
var Configured bool

// NewDatabase established a database connection and stores it in `db`.
func NewDatabase(driver, options string) error {
	if driver == "" {
		driver, options = envURL()
	}

	sql, err := gorm.Open(driver, options)

	if err != nil {
		return err
	}

	sql.DB()
	db = sql

	db.AutoMigrate(
		&Command{},
		&Config{},
		&Job{},
		&MailServer{},
		&Notification{},
		&Repository{},
		&CommandLog{},
		&User{},
	)

	cfg := GetConfig()

	if cfg.ID == 0 {
		Configured = false
	} else {
		Configured = true
	}

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
		log.Println("Invalid DATABASE_URL - using sqlite3 `leeroy.sqlite3`")
		return "sqlite3", "leeroy.sqlite3"
	}

	return s[0], s[1]
}

// NewInMemoryDatabase creates a new database using :memory:
func NewInMemoryDatabase() {
	NewDatabase("sqlite3", ":memory:")
}
