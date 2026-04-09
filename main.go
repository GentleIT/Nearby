package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	Area := struct {
		width  float64
		length float64
	}{
		width:  16,
		length: 16,
	}

	user := User{
		name: "Edward",
		position: Position{
			x: 0,
			y: 0,
		},
	}

	fmt.Scan(&user.position.x)
	fmt.Scan(&user.position.y)
	// Getting the hash from the coords of users position with data of area
	hash, options := GetHashFromCoords(user.position.x, user.position.y, Area.width, Area.length, 3)
	fmt.Println(string(hash), options)
	// To get slice of 1<=n<=8 neighbours from the current hash (position) of user
	hashN := FindHashNeighbours(user, options)
	for i := range hashN {
		// Print all neighbours from the slice
		fmt.Printf("%v | ", string(hashN[i]))
	}

	user.CreatePost("Hello, World!")
	user.CreatePost("Hello, World!!!!!")
	fmt.Println(StoreOfPosts())

	// To get time of process
	fmt.Printf("\n%v", time.Since(timeStart))
}

/*
	Message for myself: (09.04)
		- User.hash
			= It should somehow recheck its hash every 5 seconds
		- Posts
			= How to store them?
*/

/* Print Formating explanation (09.04.26)
Note:
	| - fmt from main.go
	| -
Print Format:
	|Scan for x:
	|Scan for y:
	|hash {width_of_area, length_of_area, precision}

	|neighbouring hashes
	|time of a process
*/
