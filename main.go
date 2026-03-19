package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()
	Area := struct {
		width  int
		length int
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

	hash, options := GetHashFromCoords(int(user.position.x), int(user.position.y), Area.width, Area.length, 3)
	fmt.Println(string(hash), options)
	hashN := FindHashNeighbours(user, options)
	for i := range hashN {
		fmt.Printf("%v | ", string(hashN[i]))
	}

	fmt.Printf("\n%v", time.Since(timeStart))
}
