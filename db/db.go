package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and sets up the database tables.
// It opens a connection to the SQLite3 database file "api.db" and sets the maximum number of open and idle connections.
// If the connection fails, it panics with an error message.
// After setting up the connection, it calls the createTables function to create the necessary tables.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}
	
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// createTables creates the necessary tables in the database if they do not already exist.
// It creates a "users" table with columns for id, email, and password.
// The email column is unique and cannot be null.
// It also creates an "events" table with columns for id, name, description, location, dateTime, and user_id.
// The user_id column is a foreign key that references the id column in the users table.
func createTables() {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)  
   `

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTables := `
	 CREATE TABLE IF NOT EXISTS events (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 location TEXT NOT NULL,
	 dateTime DATETIME NOT NULL,
	 user_id INTEGER,
	 FOREIGN KEY(user_id) REFERENCES users(id)
     )
	`
	_, err = DB.Exec(createEventsTables)

	if err != nil {
		panic("Could not create events table")
	}


	createRegistrationsRable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id  INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id)
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsRable)

	if err != nil {
		panic("Could not create registrations table")
	}
}