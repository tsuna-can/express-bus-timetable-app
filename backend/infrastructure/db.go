package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func NewDb() (*sqlx.DB, error) {
	var err error

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s", user, password, host, dbname)
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		panic(err)
	}

	return db, nil
}
