package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Global variable for the database connection

func InitDB() {
	var err error                           // Declare err separately
	DB, err = sql.Open("sqlite3", "api.db") // Assign to the global DB
	if err != nil {
		panic("Could not connect to Database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL, 
    description TEXT NOT NULL, 
    location TEXT NOT NULL, 
    datetime DATETIME NOT NULL, 
    user_id INTEGER
    )
    `
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create event table: " + err.Error()) // Print detailed error
	}
}
