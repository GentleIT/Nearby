package main

import (
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
