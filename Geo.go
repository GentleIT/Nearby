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

/*
	Что делать?:
		- Прочитать ответ Gemini и разобраться что делать дальше.
		(Я сколько думал и так часто мимо пропускал мысль что нужно сначала создать новый центр и
		потом сравнивать со стороны центра. Или я уже кое-как типа сделал это я уже хз. Котелок не хочет варить)
*/

const (
	upperLeft  = "a"
	upperRight = "b"
	downLeft   = "c"
	downRight  = "d"
)

// Функция поиска пути и превращения её в строку по котороу можно будет находить людей в близости.
/*
	Над поиском алгоритма (продумыванием и созданием) а также имплементацией её с абстрактных ходов
	работал дня 3-4 мне кажется.
*/
func Hash(user *User, width, length int, precision int) string {
	hash := ""

	midx := width / 2 // 8 8
	midy := length / 2
	stepX := width / 2 // Централизация + дополнительное деление для правильного центрирования в дальнейших итерациях
	stepY := length / 2

	for i := 1; i <= precision; i++ {
		isRight := user.position.x-midx >= 0
		isUp := user.position.y-midy >= 0

		stepX /= 2
		stepY /= 2
		fmt.Println(i, "iteration:", stepX, stepY)
		fmt.Println(i, "mid:", midx, midy)

		switch {
		case !isRight && isUp:
			hash += upperLeft
			midx -= stepX
			midy += stepY
		case isRight && isUp:
			hash += upperRight
			midx += stepX
			midy += stepY
		case !isRight && !isUp:
			hash += downLeft
			midx -= stepX
			midy -= stepY
		case isRight && !isUp:
			hash += downRight
			midx += stepX
			midy -= stepY
		}
	}
	return hash
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

func Normalize(s string) string {
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
	return s
}
