package db

import (
	"database/sql"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

// InitializeDatabase initializes the SQLite database and creates necessary tables if they don't exist
func InitializeDatabase() error {
	var err error
	DB, err = sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		return err // Return the error to the caller
	}

	// Create Blogs table if not exists
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		author TEXT,
		timestamp INT
	);`)
	if err != nil {
		return err // Return the error to the caller
	}

	// Create Users table if not exists
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);`)
	if err != nil {
		return err // Return the error to the caller
	}

	return nil // Return nil if there are no errors
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}
