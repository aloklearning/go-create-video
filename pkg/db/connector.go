package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnect() (*sql.DB, string) {
	db, err := sql.Open("sqlite3", "./videos.db")
	if err != nil {
		return nil, fmt.Sprintf("Failed to connect to the DB due to %v", err)
	}

	return db, "Successfully Connected to the DB"
}
