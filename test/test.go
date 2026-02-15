package main

import (
	"fmt"
	"time"
)

/*
	- Реализация геохешей???

	Попробовать создать программу, которая оптимизировано будет проверять схожесть цифр из 9 чисел.
	Если не совпали 9 цифр, то смотрит сравнение 8 цифр, если нет 7 цифр и т.д.
*/

/*
// Теперь попробовать написать код так, дабы он возвращал, в каком моменте числа начали не совпадать.
Что если попробовать сделать формат map с определенной длинной (9 элементов) и записывать туда значения?
Или же просто переобразовывать формат string в map и наоборот.
*/
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
	for i := 0; i < len(list[0]); i++ {
		if list[0][i] == list[1][i] {
			result += string(list[0][i])
		} else {
			break
		}
	}
	return result

	// notInShape := make([]string, 0, 2)
	// if len(a) != 9 {
	// 	notInShape = append(notInShape, a)
	// }
	// if len(b) != 9 {
	// 	notInShape = append(notInShape, b)
	// }
	// for _, v := range notInShape {
	// 	Normalize(v)
	// }

	// if 9 < len(a) || len(a) > 9 {
	// 	Normalize(a)
	// }
	// if 9 < len(b) || len(b) > 9 {
	// 	Normalize(b)
	// }
	// return a, b
}

/*
	= Попробовать нормализовать длину символов в каждой строке под 9-значный формат.
*/

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

func main() {
	startTime := time.Now()
	a := "1ab4dcc84"
	b := "1ab4dcc84fdgdfgr"
	c := "1ab4ddc8"

	fmt.Println(a, b, CheckSimilarity(a, b))
	fmt.Println(b, c, CheckSimilarity(b, c))

	// fmt.Println(Normalize(a))
	// fmt.Println(Normalize(b))
	// fmt.Println(Normalize(c))

	// fmt.Println(CheckSimilarity(&a, &b))
	// fmt.Println(CheckSimilarity(&a, &c))

	fmt.Println(time.Since(startTime) * time.Second)
}

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
