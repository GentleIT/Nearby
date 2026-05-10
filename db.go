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

func SendPostToDb(post *Post) {
	db := ConnectToDB()
	defer db.Close()

	query := "INSERT INTO posts (hash, owner_name, title, content) VALUES ($1, $2, $3, $4) RETURNING id"
	id, err := db.Exec(query, post.hash, post.owner, post.title, post.content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
} // Should I somehow save id in some map? But this map doesnt make any sense if it will wipe itself after every run
// Or should I return post id to save this id in another column or some other data variant to save posts belonging to user?

func GetPostsFromDb(db *sql.DB) ([]Post, error) {
	query := "SELECT * FROM posts"
	posts := make([]Post, 0, 100)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		postRow := Post{}
		if err := rows.Scan(&postRow.id, &postRow.hash, &postRow.created_at, &postRow.owner, &postRow.title, &postRow.content); err != nil {
			return nil, err
		}

		posts = append(posts, postRow)
	}

	return posts, nil
}
