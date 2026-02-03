package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {

	var err error
	DB, err = sql.Open("sqlite3", "../api.db")
	if err != nil {
		panic("Could not connect to the Database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {

	var tables []string

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL)`

	createEventTable := `
		CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER UNIQUE,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	tables = append(tables, createUsersTable, createEventTable)

	for _, table := range tables {
		_, err := DB.Exec(table)
		if err != nil {
			fmt.Println(err)
			panic("Could not create table.")
		}
	}

}
