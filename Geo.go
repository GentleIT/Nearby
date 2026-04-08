package main

import (
	"fmt"
	"math"
)

const (
	upperLeft  = 'a'
	upperRight = 'b'
	downLeft   = 'c'
	downRight  = 'd'
)

type Options struct {
	areaWidth  float64
	areaLength float64
	precision  uint8
}

// Bad commentation and english
// To hash the coords. Gives a "cell" name.
func GetHashFromCoords(x, y float64, width, length float64, precision uint8) ([]rune, Options) {

	hash := make([]rune, 0, precision)

	options := Options{
		areaWidth:  width,
		areaLength: length,
		precision:  precision,
	}

	midx := width / 2 // 8 8
	midy := length / 2
	stepX := width / 2 // iterations
	stepY := length / 2

	for i := uint8(1); i <= precision; i++ {
		isRight := x-midx >= 0
		isUp := y-midy >= 0

		stepX /= 2 // <=========--------- Shall I devide it to cellX so that it would get more precise? Or am I tripping?
		stepY /= 2

		switch {
		case !isRight && isUp: // left / on line && up
			// hash += upperLeft
			hash = append(hash, upperLeft)
			midx -= stepX
			midy += stepY
		case isRight && isUp: // on right && up
			// hash += upperRight
			hash = append(hash, upperRight)
			midx += stepX
			midy += stepY
		case !isRight && !isUp: // left / on line && down / on line
			// hash += downLeft
			hash = append(hash, downLeft)
			midx -= stepX
			midy -= stepY
		case isRight && !isUp: // right && on line / down
			// hash += downRight
			hash = append(hash, downRight)
			midx += stepX
			midy -= stepY
		}
	}

	return hash, options
}

/*
	To-do:
		1. Fix the problems with types
		2. Check for allocations
		3. Possible problem with user position
*/

// Finds and gives an array of neighbouring hashes.
func FindHashNeighbours(user User, options Options) [][]rune {
	// Find 8 coords and then check with loop for -values (below zero). Check only +values.
	// 1. Check left/right, up-down
	// 2. Store all of the 8 coords in array
	// 3. Check and remove the ones with either -x or -y values and then hash them.

	hashList := make([][]rune, 0, 8)
	storedCoords := make([][]float64, 0, 8)

	areaWidth := options.areaWidth
	areaLength := options.areaLength
	precision := options.precision

	cellX, cellY := GetCell(options)
	fmt.Println("Cells: ", cellX, cellY)
	xList := []float64{-cellX, 0, +cellX, -cellX, +cellX, -cellX, 0, +cellX} // <====
	yList := []float64{+cellY, +cellY, +cellY, 0, 0, -cellY, -cellY, -cellY} // <====

	/*
		Recent Problems:
			- Rays: Logic can break if "rays" start from wrong coords. For example, x:15,y:15 is top-right of bbb. If cell is 1.875 and not 2 then rays would never meet other hash zones.
			- Cell: Should I really allow the float cell values? I mean, it may be wrong in some cases. Or... I dont really know..
	*/

	// To get all 8 possible neighbours
	for i := 0; i < cap(storedCoords); i++ {
		x := float64(user.position.x) - xList[i]
		y := float64(user.position.y) - yList[i]

		storedCoords = append(storedCoords, []float64{x, y})
	}

	// To hash 8 possible coords that are neighbours to user
	for _, coord := range storedCoords {
		// For cases when neighbour coords go beyond map borders.
		if coord[0] < 0 || coord[1] < 0 || coord[0] > areaWidth || coord[1] > areaLength {
			continue
		}
		hash, _ := GetHashFromCoords(coord[0], coord[1], areaWidth, areaLength, precision)
		hashList = append(hashList, hash)
	}
	return hashList
}

// CellSize: XxY(map): XxY(cell) | 16x16: 2x2 | 32x32: 4x4 | 64x64: 8x8 | 128x128: 16x16 | ...
func GetCell(options Options) (float64, float64) {
	cell := struct {
		x float64
		y float64
	}{}
	mapWidth := float64(options.areaWidth) // Intently saved it in another variable
	mapLength := float64(options.areaLength)
	precision := float64(options.precision)

	cell.x = mapWidth / math.Pow(2, precision)
	cell.y = mapLength / math.Pow(2, precision)

	return cell.x, cell.y
}

/*
Problems: - O(N^2) difficulty. I check neighbours too much. I can check Neighbour A for neighbour B and then Neihbour B for neighbour A
- Too much operations. I should somehow just check only one step further from my distant hash and dont check others.
- map[0][0] - is very fragile, I should somehow create structure to prevent a panic (if value in key would be empty or smthng other)
*/
// I writed this function to save output once and forever (the data of neighbours for one hash) into database from which I could take information of what neighbouring cells users can see depending on his location (that saved in hash)
func HashNeighboursForAll(hashMap map[string][][]int) map[string][]string {
	list := make(map[string][]string)

	for k, v := range hashMap { // <-- O(N^2) loop in loop.
		for nk, nv := range hashMap { // <-- nasty
			resX := v[0][0] - nv[0][0] // <-- nasty
			resY := v[0][1] - nv[0][1]

			if (resX == 2 || resX == -2 || resX == 0) && (resY == 2 || resY == -2 || resY == 0) && k != nk { // <-- idk, seems somewhat normal, I dont see alternatives.
				list[k] = append(list[k], nk)
			}
		}
	}

	return list
}

// This function seems like working not quite right, need to test
// Can be useful in future after refactoring
func GetCoordsFromHash(hash []rune, options map[string]int) {
	searchScope := struct {
		x int
		y int
	}{
		x: options["areaWidth"],
		y: options["areaLength"],
	}
	stepX := searchScope.x / 2 // iteration
	stepY := searchScope.y / 2
	searchScope.x /= 2
	searchScope.y /= 2

	for i := 0; i < options["precision"]; i++ {
		stepX /= 2
		stepY /= 2
		switch {
		case hash[i] == 'a':
			searchScope.x -= stepX
			searchScope.y += stepY
		case hash[i] == 'b':
			searchScope.x += stepX
			searchScope.y += stepY
		case hash[i] == 'c':
			searchScope.x -= stepX
			searchScope.y -= stepY
		case hash[i] == 'd':
			searchScope.x += stepX
			searchScope.y -= stepY
		}
	}
	fmt.Println(searchScope)
}
