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

// 11.05.26: Should I really store  them (areaWidth, areaLength, precisioN) in different data types?
func GetOptions(areaWidth, areaLength float64, precision uint8) *Options {
	if areaWidth == 0 && areaLength == 0 && precision == 0 {
		return &Options{
			areaWidth:  16,
			areaLength: 16,
			precision:  3,
		}
	}
	return &Options{
		areaWidth:  areaWidth,
		areaLength: areaLength,
		precision:  precision,
	}
}

// Bad commentation and english
// To hash the coords. Gives a "cell" name.
/*11.05.26: parametrs: (position Position, options Options) -- I should store options somewhere and only give it to this function, not get it*/
//11.05.26: func GetHashFromCoords(x, y float64, width, length float64, precision uint8) ([]rune, Options) {
func (options *Options) GetHashFromCoords(position Position) string {

	width := options.areaWidth
	length := options.areaLength
	precision := options.precision

	x, y := position.x, position.y

	hash := make([]rune, 0, precision+1)

	midx := width / 2
	midy := length / 2

	// iterations
	stepX := width / 2
	stepY := length / 2

	for i := uint8(1); i <= precision; i++ {
		isRight := x-midx >= 0
		isUp := y-midy >= 0

		stepX /= 2 // <=========--------- Shall I devide it to cellX so that it would get more precise? Or am I tripping?
		stepY /= 2

		switch {
		case !isRight && isUp:
			hash = append(hash, upperLeft)
			midx -= stepX
			midy += stepY
		case isRight && isUp:
			hash = append(hash, upperRight)
			midx += stepX
			midy += stepY
		case !isRight && !isUp:
			hash = append(hash, downLeft)
			midx -= stepX
			midy -= stepY
		case isRight && !isUp:
			hash = append(hash, downRight)
			midx += stepX
			midy -= stepY
		}
	}

	return string(hash)
}

/*
	To-do:
		1. Fix the problems with types
		2. Check for allocations
		3. Possible problem with user position
*/

// 15.05.26 commentary: Maybe I use too complicated version and rather need to ponder about offset with bytes or some other
// mathemathical method, but I'm fine with this version at the moment.
// Finds and gives an array of neighbouring hashes.
func (options *Options) FindHashNeighbours(position Position) []string { // hash should be in string format
	// Find 8 coords and then check with loop for -values (below zero). Check only +values.
	// 1. Check left/right, up-down
	// 2. Store all of the 8 coords in array
	// 3. Check and remove the ones with either -x or -y values and then hash them.

	hashList := make([]string, 0, 8)
	storedCoords := make([][]float64, 0, 8)

	areaWidth := options.areaWidth
	areaLength := options.areaLength

	cellX, cellY := GetCell(options)

	// Used only for check for duplicates: to dont put current hash in neighbouring hashes list
	posHash := options.GetHashFromCoords(position)

	xList := []float64{-cellX, 0, +cellX, -cellX, +cellX, -cellX, 0, +cellX} // <====
	yList := []float64{+cellY, +cellY, +cellY, 0, 0, -cellY, -cellY, -cellY} // <====

	/*
		Recent Problems:
			- Rays: Logic can break if "rays" start from wrong coords. For example, x:15,y:15 is top-right of bbb. If cell is 1.875 and not 2 then rays would never meet other hash zones.
				(that thing that rays from center of coord to 8 directions)
			- Cell: Should I really allow the float cell values? I mean, it may be wrong in some cases. Or... I dont really know..
	*/

	// --- 15.05.26: Can't I pass this creation of slice and just hop to the next calculating and 'hash'-ing calculated
	// cords and then store it in hashList?
	// To get all 8 possible neighbours
	for i := 0; i < cap(storedCoords); i++ {
		x := float64(position.x) - xList[i]
		y := float64(position.y) - yList[i]

		storedCoords = append(storedCoords, []float64{x, y}) // <-- 8 allocations: 8 times I create new massive pointers.
	}

	// To hash 8 possible coords that are neighbours to user
Outerloop:
	for _, coord := range storedCoords {
		// For cases when neighbour coords go beyond map borders.
		if coord[0] < 0 || coord[1] < 0 || coord[0] > areaWidth || coord[1] > areaLength {
			continue
		}
		hash := options.GetHashFromCoords(Position{x: coord[0], y: coord[1]})

		// 14.05.26: There I should check for the repeats in hashes. 15.05.26:
		if hash != posHash {
			for _, sthash := range hashList {
				if hash == posHash || hash == sthash {
					// fmt.Println("Duplicate: ", hash, " of ", posHash)
					continue Outerloop
				}
			}
			hashList = append(hashList, hash)
		} else {
			continue Outerloop
		}
	}
	return hashList
}

// CellSize: XxY(map): XxY(cell) | 16x16: 2x2 | 32x32: 4x4 | 64x64: 8x8 | 128x128: 16x16 | ...
// Returns the size of cell needed to define inner borders
func GetCell(options *Options) (float64, float64) {
	cell := struct {
		x int
		y int
	}{}

	mapWidth := options.areaWidth // Intently saved it in another variable
	mapLength := options.areaLength
	precision := float64(options.precision) // Turned with float64() because of math.Pow() down here

	cell.x = int(mapWidth / math.Pow(2, precision))
	cell.y = int(mapLength / math.Pow(2, precision))

	return float64(cell.x), float64(cell.y)
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
