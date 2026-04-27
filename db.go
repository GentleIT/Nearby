package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// For future: Not good that I save my password and login in there and not in environment variables
func ConnectToDB() *sql.DB {
	// "postgres://user:pass@localhost:5432/dbname"
	var dsn = "postgres://postgres:postgres@localhost:5432/Nearby_db?sslmode=disable"
	var db, err = sql.Open("pgx", dsn)

	if err != nil {
		fmt.Println("db error:", err)
	}
	fmt.Println(db)
	return db
}

// func SendPostToDb() {

// }
