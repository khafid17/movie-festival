package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3306)/movie_festival"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to connect to the database")
		return nil, err
	}
	return db, nil
}