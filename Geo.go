package main

import (
	"fmt"
)

/*
Теперь что делать?

Есть:
  - Возвращение одинаковых символов начиная с первой до первой несостыковки. (Уже проверка схожести геохешей)
    = Нормализация строки под 9-ти символьный формат.
  - Проверка одинаковы ли два геохеша.

Теперь:
  - Попробовать создать геохеш с данным о координатах (x и y (которые замена ширине и долготе))

Для этого:
  - Расширить примеры, добавив возможности минусовых x и y
  - Посмотреть или придумать как имея только 2 значения понять в каком геохеше он находится.

- Ну и в будущем было бы классно попробовать визуализировать их.
*/

// Либо превращать координаты в хэши и тупа класть их в хэш-таблицу (мапу) и просто там проверять
// для поста кто находится в той же таблице, где и юзер, дабы его могли увидеть.
// Или же, попробовать всё же сделать алгоритм.

type extPosition struct {
	lenx int
	leny int
	midx float32
	midy float32
}

/*
	Что делать?:
		- Прочитать ответ Gemini и разобраться что делать дальше.
		(Я сколько думал и так часто мимо пропускал мысль что нужно сначала создать новый центр и
		потом сравнивать со стороны центра. Или я уже кое-как типа сделал это я уже хз. Котелок не хочет варить)
*/

func Hash(user User, width int, length int) string {
	hash := ""

	area := extPosition{
		lenx: width,
		leny: length,
	}

	/*
		! Мне нужно высчитывать новый центр зоны, как в своем подобии GeoHash с реализацией через Z-curve (хотя я сам особо не понимаю этот метод)
		> Сперва будет карта длиной 10 на 10, далее нужно поделить центр на 2, дабы, карта была разделена на 4 участка.
		> Далее, через готовую функцию проверяю местоположение и, например, если будет слева-снизу, то снова делю на 2.
		> Думаю даже добавить float значения для длины и центра. Юзеры(Посты их) будут то на определенных точках только.
	*/
	notDefined := true
	area.midx = float32(area.lenx) / 2        // divided to center
	area.midy = float32(area.leny) / 2        // divided to center
	iterationValueX := float32(area.lenx) / 2 // divided to optimize algoritm
	iterationValueY := float32(area.leny) / 2 // divided to optimize algoritm
	var i float32 = 1

	for notDefined {
		fmt.Println(i, area.midx, area.midy)
		iterationValueX /= 2
		iterationValueY /= 2
		fmt.Println("Iteration value:", iterationValueX, iterationValueY)
		switch extCalculateDirection(float32(user.position.x), float32(user.position.y), area.midx, area.midy) {
		case "left up ": // (каждая помеченная сторона даёт буквы хэшу)
			hash += "a"
			area.midx -= iterationValueX
			area.midy += iterationValueY
		case "right up ":
			hash += "b"
			area.midx += iterationValueX
			area.midy += iterationValueY
		case "left down ":
			hash += "c"
			area.midx -= iterationValueX
			area.midy -= iterationValueY
		case "right down ":
			hash += "d"
			area.midx += iterationValueX
			area.midy -= iterationValueY
		case "left ":
			hash += "<"
			area.midx -= iterationValueX
		case "right ":
			hash += ">"
			area.midx += iterationValueX
		case "up ":
			hash += "//"
			area.midy += iterationValueY
		case "down ":
			hash += "()"
			area.midy -= iterationValueY
		}

		i++

		if i > 3 {
			break
		}
	}
	return hash
}

// 	for notDefined {
// 		i++
// 		iterationValueX /= 2 * i
// 		iterationValueY /= 2 * i

// 		// Я уже почти закончил с этой функцией
// 		// Осталось только удалить штуку ниже, переписать CalculateDirection(), изменить возвращаемые значения
// 		// Дабы не сравнивала со строками (нагружает). Ну и в конце все же вернуть себе хэш.

// 		_, x, y := extCalculateDirection(float32(user.position.x), float32(user.position.y), area.midx, area.midy)
// 		switch x {
// 		case -1:
// 			//left
// 			area.midx -= iterationValueX
// 			hash += "left"
// 		case 1:
// 			//right
// 			area.midx += iterationValueX
// 			hash += "right"
// 		}

// 		switch y {
// 		case -1:
// 			//down
// 			area.midy -= iterationValueY
// 			hash += "down"
// 		case 1:
// 			//up
// 			area.midy += iterationValueY
// 			hash += "up"
// 		}

// 		if i > 3 {
// 			break
// 		}
// 	}
// 	return hash
// }

// // В этой функции я хочу хэшировать текущую координату юзера в зависимости от итераций в каких квадратах он находится.
// func Old2Hash(user *User, areaWidth int, areaLength int) string {
// 	notDefined := true
// 	hash := ""

// 	newZeroPoint := extPosition{
// 		x0: 0,
// 		x1: areaWidth, // Они показывают 1. Всю длину линии, 2. Конечное число в конце
// 		y0: 0,
// 		y1: areaLength,
// 		// Возможно нужно будет разделить логику и добавить ниже два значения, max длину x и y
// 	}
// 	// А можно ли поделить всё так, чтоб юзеры могли быть ток на цельных x и y (int), а границы шли по (float) числам?

// 	// Здесь уже определяю хэш в зависимости в какой стороне находится координаты юзера.

// 		fmt.Println(newZeroPoint)
// 		if newZeroPoint.y0 > 10 {
// 			break
// 		}
// 	}
// 	return hash
// }

// func OldHash(user *User) string {
// 	// Здесь я просто обновляю новосозданный zero-point пока он не будет максимально близок к Юзеру.

// 	newX := 0
// 	newY := 0
// 	notDefined := true

// 	for notDefined {
// 		zeroPoint := User{
// 			position: Position{
// 				x: newX,
// 				y: newY,
// 			},
// 		}
// 		switch CalculateDirection(user, &zeroPoint) {
// 		case "right up ":
// 			newX += 1 * 2
// 			newY += 1 * 2
// 		case "left up ":
// 			newX += -(1 * 2)
// 			newY += 1 * 2
// 		case "right down ":
// 			newX += 1 * 2
// 			newY += -(1 * 2)
// 		case "left down ":
// 			newX += -(1 * 2)
// 			newY += -(1 * 2)
// 		case "right ":
// 			newX += (1 * 2)
// 		case "left ":
// 			newX += -(1 * 2)
// 		case "up ":
// 			newY += (1 * 2)
// 		case "down ":
// 			newY += -(1 * 2)
// 		}
// 		fmt.Println(newX, newY, user.position)

// 		// Нижний код закончил цикл на моменте когда x и y zeroPoint стал максимально близко к юзерским координатам.
// 		// Зачем мне это? Если так подумать то я просто приблизил zeroPoint к юзеру. Но не определил для него ходы хеша.
// 		nearX := newX - user.position.x
// 		nearY := newY - user.position.y
// 		if nearX < 2 && nearY < 2 { // Если находится в пределах 2-ух точек
// 			fmt.Println("Finished: ", newX, newY, user.position)
// 			break
// 		}
// 	}
// 	return "hash"
// }

func CheckSimilarity(a string, b string) string {
	list := []string{a, b}
	for i, v := range list {
		if len(v) != 9 {
			list[i] = Normalize(v)
		}
	}
	// Нужно сделать 9 итераций и проверить схожи ли цифры. Вернуть строку с одинаковыми значениями. Если первый не одинаков то всё.
	// Хотя было бы интересно попробовать сделать задачку где нужно вернуть строку только с теми символами, которые были одинаковы. Мне кажется это не так сложно. Нужно тупо result += v (v - это символы которые совпали начиная с первого. Потом еще вернуть индексы (итерации) которые совпали)
	var result string
	for i := 0; i < len(list[0]); i++ { // Насчёт проверки длинны через ток одно значение не сильно уверен, но прикол в том что я уже нормализовал их под формат 9-ти символов, так что разницы нет.
		if list[0][i] == list[1][i] {
			result += string(list[0][i])
		} else {
			break
		}
	}
	return result
}

func Normalize(s string) string { // Работает неплохо. Нормализовано нормально
	// fmt.Println("Working with", s)
	notInForm := true
	switch {
	case len(s) > 9:
		// i := 0
		for notInForm {
			s = s[0 : len(s)-1]
			if len(s) == 9 {
				notInForm = false
			}
			// i++
			// fmt.Println(i)
		}
	case len(s) < 9:
		for notInForm {
			s += "_"
			if len(s) == 9 {
				notInForm = false
			}
		}
	}
	// fmt.Println(s)
	return s
}

// Также попробовать создать функцию хеширования под формат 9-символов в строке (Мб тогда многое сверху станет не нужно)

// Говнокод который еще свет не ведал. Шучу. Хочу переписать его без повторений и лишних объявлений.
func saySimilar(a *string, b *string) bool {
	var longHash string
	var shortHash string
	if len(*a) > len(*b) {
		longHash = *a
		shortHash = *b
	} else {
		longHash = *b
		shortHash = *a
	}

	ind := len(shortHash)
	count := 0
	for i := range shortHash {
		fmt.Println(i)

		if longHash[i] == shortHash[i] {
			count += 1
			continue
		}
	}
	if count == ind {
		return true
	} else {
		return false
	}
}
