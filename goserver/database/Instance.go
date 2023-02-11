package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

// Singleton - initialize here
var db, err = sql.Open("sqlite3", "../data/db.db")

// Define mutex
var dbLock = &sync.Mutex{}

func GetInstance() *sql.DB {
	dbLock.Lock()
	defer dbLock.Unlock()

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
