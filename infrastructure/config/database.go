package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "macbook"
	DBPassword = "password"
	DBName     = "quiz"
)

func DbConn() (*sql.DB, error) {
	connStr := "dbname=" + DBName + " user=" + DBUser + " password=" + DBPassword + " host=" + DBHost + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
