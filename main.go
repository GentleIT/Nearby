package main

import (
	"fmt"
	"sync"
)

/*
	Я хотел попробовать в слайс создать сотню юзеров и попробовать для каждого определить расстояние до следующего.
	Их первоначальные позиции определялись бы рандомно.

	Есть ли смысл сейчас пробовать это применять?
	- Попробовать создать способ определения в реальном времени в каком квадрате находится User.
	  Либо сделать это и попробовать сделать через while и switch case либо же как-то по другому.
	- Но вообще я бы хотел создать всё же способ визуализации этих данных каким-нибудь способом.
	  Пора ли? Но пока еще несколько функций на беке не было создано
	- Что если создать систему квадратов и в этих квадратах юзер сможет открыть и прочитать посты которые написали
	  Другие пользователи. Только как это все создать?

	= Реализации системы зон в которой юзер находится и которая будет проверяться через while()
*/

// Как вообще оно будет само запускаться и всегда смотреть в какой зоне находиться юзер? Есть ли возможно создавать это
// не для каждого юзера, а попробовать создать хендлер, который будет обрабатывать всех юзеров и возвращать им через
// канал данные где они находятся?

func LocateZoneOfUser(users []User, wg *sync.WaitGroup) {
	for true {
		for _, v := range users {
			switch {
			case v.position.x > 0:
				// Do something
			case v.position.x < 0:
				// Do something
			default:
				// Do something
			}
			fmt.Println(v)
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	user1 := User{
		name: "Herald",
		position: Position{
			x: 2,
			y: -5,
		},
	}
	user2 := User{
		name: "Zero-Point",
		position: Position{
			x: 0,
			y: 0,
		},
	}

	users := []User{user1, user2}
	wg.Add(1)
	go LocateZoneOfUser(users, &wg)
	wg.Wait()
	fmt.Printf("Location of User: x: %d, y: %d\n", user1.position.x, user1.position.y)
	fmt.Printf("Location of Zero-Point: x: %d, y: %d\n", user2.position.x, user2.position.y)
	fmt.Println("Result:")
	fmt.Printf("Distance from zero-point: %v\n", CalculateDistance(&user1, &user2))
	fmt.Printf("Direction to reach zero-point: %v\n", CalculateDirection(&user1, &user2))
}
