// Package db provides database initialization and connection management.
package db

import (
	"database/sql"
	"log"
	"udemy-multi-api-golang/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and creates necessary tables.
func InitDB(cfg *config.Config) error {
	var err error
	DB, err = sql.Open("sqlite3", cfg.Database.Path)
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return err
	}

	log.Printf("database connected successfully: %s\n", cfg.Database.Path)

	if err := createTables(); err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database connection.
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// createTables creates all necessary database tables if they don't exist.
func createTables() error {
	tables := []string{
		createUsersTable,
		createEventsTable,
		createRegistrationsTable,
	}

	for _, query := range tables {
		if _, err := DB.Exec(query); err != nil {
			log.Printf("failed to create table: %v\n", err)
			return err
		}
	}

	log.Println("all tables created/verified successfully")
	return nil
}

const (
	// createUsersTable defines the users table schema.
	createUsersTable = `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	// createEventsTable defines the events table schema.
	createEventsTable = `
		CREATE TABLE IF NOT EXISTS events(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	// createRegistrationsTable defines the registrations table schema.
	createRegistrationsTable = `
		CREATE TABLE IF NOT EXISTS registrations(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id),
			UNIQUE(event_id, user_id)
		)
	`
)
