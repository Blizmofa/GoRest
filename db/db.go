package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Table creation SQL constants
const (
	createUsersTable = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	createEventsTable = `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	createRegistrationsTable = `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	defaultDBPath        = "api.db"
	maxOpenedConnections = 10
	maxIdleConnections   = 5
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", defaultDBPath)

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err.Error())
	}

	DB.SetMaxOpenConns(maxOpenedConnections)
	DB.SetMaxIdleConns(maxIdleConnections)

	createTable(createUsersTable)
	createTable(createEventsTable)
	createTable(createRegistrationsTable)
}

func createTable(tableQuery string) {
	_, err := DB.Exec(tableQuery)

	if err != nil {
		log.Fatalf("Could not create table: %v", err.Error())
	}
}
