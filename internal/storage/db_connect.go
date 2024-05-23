package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {
	connStr := "user=sanved dbname=sanved password=ILgta77!@# host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected!")

	return db
}
