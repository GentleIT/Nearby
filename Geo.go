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
	a := []string{"hello", "world"}

	for k, v := range hashMap {
		fmt.Println(k, v[0])
	}

	// keys := make([]string, 0, len(hashMap))
	// for v := range
	return a
}
