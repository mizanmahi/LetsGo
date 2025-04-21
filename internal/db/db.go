package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitPostgres() *sql.DB {
	dsn := "postgres://postgres:541990@localhost:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	return db
}