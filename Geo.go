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
	x0, x1 int
	y0, y1 int
}

func Hash(user *User, areaWidth int, areaLength int) string {
	notDefined := true
	hash := ""

	newZeroPoint := User{
		position: extPosition{
			x0: 0,
			x1: areaWidth,
			y0: 0,
			y1: areaLength,
		},
	}
	// newZeroPoint.position.x = areaWidth / 2
	// newZeroPoint.position.y = areaLength / 2

	// А если создать словарь (мапу), куда положу цифры 0-9 и далее буквы алфавита.
	// А потом просто буду в зависимости от возвращаемого ближайшего i присуждать пустой строке hash индекс
	// равный ключу в мапе?
	// Типа, если поделить один квадрат (карту) на 32 квадрта (в котором тоже внутри 32 квадрата),
	// то каждому будет присужден свой индекс в зависимости от того по какомсу счёту он стоит
	// Вопрос, а как такое использовать по отношению к координатам?

	for notDefined {
		// Ниже не сработает просто из за того что оно может считать положительные x и y. То есть считывают ток право-верхнюю сторону.
		// Решит наверн если я откажусь от минусов и остановлюсь только на право-верхней стороне по всему пространству.
		newZeroPoint.position.x -= newZeroPoint.position.x / 2
		newZeroPoint.position.y -= newZeroPoint.position.y / 2
		switch CalculateDirection(user, &newZeroPoint) {
		case "left up ":
			hash += "a"
		case "right up ":
			hash += "b"
		case "left down ":
			hash += "c"
		case "right down ":
			hash += "d"
		case "left ":
			hash += "<"
		case "right ":
			hash += ">"
		case "up ":
			hash += "//"
		case "down ":
			hash += "()"
		}
		nearX := newZeroPoint.position.x - user.position.x // 0 -
		nearY := newZeroPoint.position.y - user.position.y
		if nearX < 2 && nearY < 2 { // Если находится в пределах 2-ух точек
			fmt.Println("Finished: ", areaWidth, areaLength, user.position, Normalize(hash))
			notDefined = false
			break
		}
		fmt.Println("Ingoing: ", Normalize(hash), newZeroPoint.position.x, newZeroPoint.position.y)
	}
	return hash
}

func OldHash(user *User) string {
	// Здесь я просто обновляю новосозданный zero-point пока он не будет максимально близок к Юзеру.

	newX := 0
	newY := 0
	notDefined := true

	for notDefined {
		zeroPoint := User{
			position: Position{
				x: newX,
				y: newY,
			},
		}
		switch CalculateDirection(user, &zeroPoint) {
		case "right up ":
			newX += 1 * 2
			newY += 1 * 2
		case "left up ":
			newX += -(1 * 2)
			newY += 1 * 2
		case "right down ":
			newX += 1 * 2
			newY += -(1 * 2)
		case "left down ":
			newX += -(1 * 2)
			newY += -(1 * 2)
		case "right ":
			newX += (1 * 2)
		case "left ":
			newX += -(1 * 2)
		case "up ":
			newY += (1 * 2)
		case "down ":
			newY += -(1 * 2)
		}
		fmt.Println(newX, newY, user.position)

		// Нижний код закончил цикл на моменте когда x и y zeroPoint стал максимально близко к юзерским координатам.
		// Зачем мне это? Если так подумать то я просто приблизил zeroPoint к юзеру. Но не определил для него ходы хеша.
		nearX := newX - user.position.x
		nearY := newY - user.position.y
		if nearX < 2 && nearY < 2 { // Если находится в пределах 2-ух точек
			fmt.Println("Finished: ", newX, newY, user.position)
			break
		}
	}
	return "hash"
}

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
