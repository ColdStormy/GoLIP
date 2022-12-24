package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Checks if an initial setup is required.
// If it has not yet been performed, it will do so and return false.
// Else it returns true (initial setup is not needed).
func initialSetup(conf Config) bool {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return false
}
