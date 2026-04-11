package main

import (
	"database/sql"
	"fmt"
)

// For future: Not good that I save my password and login in there and not in environment variables

var dsn = "posrges://postges:postgress@localhost:5432/dbname?sslmode=disable"

func ConnectToDB() *sql.DB {
	var db, err = sql.Open("pgx", dsn)

	if err != nil {
		fmt.Println("db error:", err)
	}
	return db
}

func SendPostToDb