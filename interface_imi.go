package main

import "fmt"

func MainMenu() {
	var name string
	var x, y float64

	var title string
	var msg string

	fmt.Println("User, what is your name?")
	fmt.Scan(&name)
	fmt.Println("Where are you now in coordinates?")
	fmt.Scan(&x, &y)

	user := User{
		name:     name,
		position: Position{x: x, y: y},
		hash:     GetOptions(0, 0, 0).GetHashFromCoords(Position{x: x, y: y}),
	}

	fmt.Printf("User %v from position %v is in hash(%v)\n", user.name, user.position, user.hash)

	fmt.Println("If you want to write the post. Create title and message!")
	fmt.Print("title: ")
	fmt.Scan(&title)
	fmt.Print("message: ")
	fmt.Scan(&msg)

	_, rslt := user.CreatePost(title, msg)
	fmt.Println("Let's check if your post is in database!")
	fmt.Println(rslt)
}
