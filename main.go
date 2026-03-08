package main

func main() {

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
			hash := HashCoords(y, x, Area.width, Area.length, 3)
			// hashMap[hash] = append(hashMap[hash], strconv.Itoa(x)+"|"+strconv.Itoa(y))
			hashMap[hash] = append(hashMap[hash], []int{x, y}) // []int{x, y} искал как сделать(сначала в голове далее в инете, нужно было в обсидиане) достаточно долго.
		}
	}

	// Ниже буду получать соседей от хеша и далее помещать это в базу.
	for i := 0; i <= 10; i++ {
		// Поместить список существующих хешей и проверка их.
		FindNeighborsForHash(hashMap, []int{1, 2, 3})
	}
}
