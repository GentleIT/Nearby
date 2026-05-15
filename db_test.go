package main

import (
	"fmt"
	"testing"
)

func TestReadfromDb(t *testing.T) {
	db := ConnectToDB()
	defer db.Close()

	fmt.Println("Database Ping():", db.Ping())

	// resp, err := GetPostsFromDb(db)
	// if err != nil {
	// 	fmt.Println("err with GetPostFromDb(): ", err)
	// }
	// fmt.Println(resp)
}

func TestSendtoDb(t *testing.T) {
	db := ConnectToDB()
	defer db.Close()

	fmt.Println("Database Ping():", db.Ping())

	// testuser := User{
	// 	name: "Edward",
	// 	hash: "aaa",
	// }

	// SendPostToDb(testuser.CreatePost("hello"))
}
