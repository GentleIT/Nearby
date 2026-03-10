package main

import "fmt"

const (
	upperLeft  = "a"
	upperRight = "b"
	downLeft   = "c"
	downRight  = "d"
)

func HashCoords(y, x int, width, length int, precision int) string {
	hash := ""
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
			hash += upperLeft
			midx -= stepX
			midy += stepY
		case isRight && isUp: // on right && up
			hash += upperRight
			midx += stepX
			midy += stepY
		case !isRight && !isUp: // left / on line && down / on line
			hash += downLeft
			midx -= stepX
			midy -= stepY
		case isRight && !isUp: // right && on line / down
			hash += downRight
			midx += stepX
			midy -= stepY
		}
	}
	return hash
}

func FindNeighborsForHash(hashMap map[string][][]int, target []int) []string { // Здесь нужно принять ключ таргетного хеша и чекнуть только его нечетные коорды (справа сверху) (Хотя, в теории всё равно)
	list := make(map[string][]string)
	for k, v := range hashMap {
		for nk, nv := range hashMap {
			resX := v[0][0] - nv[0][0]
			resY := v[0][1] - nv[0][1]

			res := resX + resY

			if res == 4 || res == 2 || res == 0 || res == -2 || res == -4 { // Здесь что-то не так.
				list[k] = append(list[k], nk)
			}
			// if resX == 2 || resX == -2 {}

			listOfHashes := make([][][]int, 0, 1000)
			listOfHashes = append(listOfHashes, v)
			// for k, v := range hashMap {
			// 	// fmt.Println(k, v[0])
			// }
		}
	}
	fmt.Println(list)
	// First Object - ishodnie coordy, dlya kotorih ishu sosedei
	// Second Object - dannye hashi sosedei, budu brat' pervie cordy from all hashes.

	return []string{"hello", "world"}
}

// Сначала нужно было бы написать функцию которая бы сортировала мапу и копировала её в отдельный лист.
