package main

type Position struct {
	x float64
	y float64
}

type User struct {
	name     string
	position Position
	hash     []rune // All users while walking should
}

func (u *User) CreatePost(m string) Post {
	newPost := Post{
		owner:    *u,
		position: u.position,
		hash:     u.hash,
		message:  m,
	}
	StoreOfPosts(&newPost)
	return newPost
}

// This Phrankenshtein is scary af
func StoreOfPosts(p ...*Post) func() []Post {
	actualPost := Post{}
	if len(p[0].message) > 0 {
		actualPost = *p[0]
	}
	postSlice := make([]Post, 0, 10)
	return func() []Post {
		postSlice = append(postSlice, actualPost)
		return postSlice
	}
}

type Post struct {
	owner    User
	position Position
	hash     []rune
	message  string
}
