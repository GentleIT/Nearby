package main

import (
	"fmt"
	"time"
)

/*
	TO DO:
		- Rewrite comments to english.
		- Make normal Readme file.
		- Finnish with finding neighbours of hashes.
		- After that: integrate and save it in db.
*/

func main() {
	timeStart := time.Now()
	Area := struct {
		width  int
		length int
	}{
		width:  16,
		length: 16,
	}

	hashMap := make(map[string][][]int)
	for y := range Area.length {
		for x := range Area.width {
			// hash := HashCoords(y, x, Area.width, Area.length, 3)
			hashrune, _ := GetHashFromCoords(y, x, Area.width, Area.length, 3)
			hash := string(hashrune)
			// hashMap[hash] = append(hashMap[hash], strconv.Itoa(x)+"|"+strconv.Itoa(y))
			hashMap[hash] = append(hashMap[hash], []int{x, y}) // []int{x, y} искал как сделать(сначала в голове далее в инете, нужно было в обсидиане) достаточно долго.
		}
	}

	// Ниже буду получать соседей от хеша и далее помещать это в базу.
	// for i := 0; i <= 10; i++ {
	// 	// Поместить список существующих хешей и проверка их.
	// 	FindNeighborsForHash(hashMap, []int{1, 2, 3})
	// }
	fmt.Println(HashNeighboursForAll(hashMap))
	fmt.Println(time.Since(timeStart))
}
