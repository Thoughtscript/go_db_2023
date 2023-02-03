package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB = nil

func GetInstance() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "../data/db.db")
	log.Println("Database connection opened...")
	if err != nil {
		log.Fatal(err)
	}
	// Don't uncomment this - thought it would prevent db.Close()
	// Incorrect! This will call close at the end of the method / return.
	// defer db.Close()
	return db
}
