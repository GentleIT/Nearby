package main

import (
	"fmt"
	"testing"
)

func TestHashNeighbours(t *testing.T) {
	Area := struct {
		width  float64
		length float64
	}{
		width:  15,
		length: 15,
	}
	fmt.Println("{x y}:", Area)

	user := User{
		name: "Edward",
		position: Position{
			x: 0,
			y: 0,
		},
	}

	_, options := GetHashFromCoords(0, 0, Area.width, Area.length, 3)

	for ix := 0.0; ix <= Area.width; ix++ {
		for iy := 0.0; iy <= Area.length; iy++ {
			user.position.x = ix
			user.position.y = iy

			hash, _ := GetHashFromCoords(user.position.x, user.position.y, options.areaWidth, options.areaLength, 3)
			neisRune := FindHashNeighbours(user, options)
			neisString := make([]string, 0, 8)
			for _, neis := range neisRune {
				neisString = append(neisString, string(neis))
			}
			fmt.Printf("Find: x:%v, y: %v | hash: %v, neis: %v\n", user.position.x, user.position.y, string(hash), neisString)
		}
	}
}

/*
	Problems:
		GetHashFromCoords()
			idk why but when I print all the hash combinations sometimes I can get 3
		FindHashNeighbours()
*/
