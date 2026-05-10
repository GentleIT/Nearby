package main

import (
	"fmt"
	"time"
)

type Position struct {
	x float64
	y float64
}

type User struct {
	name     string
	position Position // should be not null
	hash     string   // All users while walking should // hash can be like this: always GetHashFromCoords(Position...)
	// hash was []rune format before
}

type Post struct { // data types <--
	// hash     []run
	id         int
	hash       string
	created_at time.Time
	owner      string
	title      string
	content    string
}

// id, hash, created_at(From Postgre), owner_name, title, content
func (u *User) CreatePost(m string) *Post {
	newPost := Post{
		hash:    u.hash,
		owner:   u.name,
		title:   fmt.Sprintf("%v''s post", u.name),
		content: m,
	}
	// Here should be the function that sends new post to a default store of all the posts. Store should be in fact a db.

	SendPostToDb(&newPost)
	return &newPost
}

// This Phrankenshtein is scary af
// func StoreOfPosts(p ...Post) func() []Post {
// 	actualPost := Post{}
// 	if len(p[0].message) > 0 {
// 		actualPost = p[0]
// 	}
// 	postSlice := make([]Post, 0, 10)
// 	return func() []Post {
// 		postSlice = append(postSlice, actualPost)
// 		return postSlice
// 	}
// }
