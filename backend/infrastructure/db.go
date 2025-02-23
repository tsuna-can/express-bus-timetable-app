package infrastructure

import (
  "log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func NewDb() (*sqlx.DB, error) {
	var err error
	connStr := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
    log.Printf("Error connecting to database: %v", err)
		panic(err)
	}

	return db, nil
}
