package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //this driver allows us to use the sql API
)

var DB *sql.DB

func Init() {
	var err error
	dsn := "arjumand.ayub:example-password@tcp(127.0.0.1:3306)/movie_cinema"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database : %v\n", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Database ping failed %v\n", err)
	}
	log.Println("Database connection made!")
}
