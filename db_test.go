package main

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	db := ConnectToDB()
	defer db.Close()

	fmt.Println("Database Ping():", db.Ping())

	testuser := User{
		name: "Edward",
		hash: "aaa",
	}

	SendPostToDb(db, testuser.CreatePost("hello"))
}
