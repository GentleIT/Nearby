package main

const (
	upperLeft  = 'a'
	upperRight = 'b'
	downLeft   = 'c'
	downRight  = 'd'
)

func GetHashFromCoords(y, x int, width, length int, precision int) ([]rune, map[string]int) {
	// hash := "" // <-- I can make it array with capacity of 6 or more, depending on precision
	hash := make([]rune, 0, precision)
	// hashSlice := make([]string, 0, precision)
	options := make(map[string]int)
	options["areaWidth"], options["areaLength"], options["precision"] = width, length, precision
	// Ошибка - стоило полировать и убрать не нужное самому а не проверять
	// как написать лаконичнееs
	midx := width / 2 // 8 8
	midy := length / 2
	stepX := width / 2 // Для правильного центрирования под каждую итерацияю
	stepY := length / 2

	for i := 1; i <= precision; i++ {
		isRight := x-midx >= 0
		isUp := y-midy >= 0

		stepX /= 2
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

func GetCoordsFromHash(hash []rune, options map[string]int) {
	// Maybe there I need to scale the logic and turn the "sides" (upperLeft, upperRight, etc...) into some arrayyyy?? Idk why,
	// but maybe in future I would have to add new runes for hash instead of 4 (a, b, c, d) or idk. My sleepy mind dont really know
	searchScope := struct {
		x int
		y int
	}{
		x: options["areaWidth"],
		y: options["areaLength"],
	}
	finX := 0
	finY := 0

	for i := 0; i <= options["presicion"]; i++ {
		// There I need to play with iterations again
		iteration := ((searchScope.x + searchScope.y) / 2) / 2
		switch {
		case hash[i] == 'a':
			searchScope.x -= iteration
			searchScope.y -= iteration
		case hash[i] == 'b':
		case hash[i] == 'c':
		case hash[i] == 'd':
		}
	}
}

// This function should be optimized version for calculating the neighbours of hash (hash of User) without using saved tables.
func NewHashNeighboursForAll(hashMap map[string][][]int) []string {
	resMap := make(map[string][]string)
	hashList := make([]string, 0, len(hashMap)) // <-- 1. Should I change the data type from []string to f.e. [][]string? To put inside values of hash keys
	for k := range hashMap {                    // or would be it recreating old solution? (function in function in another way?)
		hashList = append(hashList, k) // it always will be randomized, so
		// hashList = hashLists // <-- 2. Creating sorting function is manadotory to keep solving problem of defining neighbours.
	}

	i := 0 // <- to go from one hash to another in sorted order.
	for k, v := range hashMap {
		i++
		k = k
		v = v
		resMap[hashList[i]] = append(resMap[hashList[i]])
	}

	return hashList
}

/*
Идея: получить позицию текущую у юзера и сделать 8 шагов вдаль от текущей позиции. С этого получить 8 координат, и
с помощью поиска хеша узнать хеш-ную строку этих 8-ми координат. Так в целом можно узнать хешные соседи текущего хеша пользователя.
*/
func FindHashNeighbours() []string {
	return []string{}
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

/*
	1. Neighbours of hashes can be saved in some map and neigbours can be save to some list of a user (he saves the list of his neighbours),
	   but when he switches his position ang geohash automatically will be made operation that gives to his list new neighbours.
	2.
*/

// Сначала нужно было бы написать функцию которая бы сортировала мапу и копировала её в отдельный лист.
