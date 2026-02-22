package main

import (
	"math"
)

type Position struct {
	x int
	y int
}

type Post struct {
	id            int8
	position      Position
	usersAttached []User // Здесь список юзеров которые смогут увидеть экземпляр поста
}

type User struct {
	name     string
	position Position
	// postThatCanSee []Post // Список посток которые юзер увидит в своей ленте если он был в 2-ух км от момента поста.
}

// Можно ли сделать так, дабы ко второму аргументу он мог принимать или User или Position? Полиморфизм?
// Можно если структура будет имплементировать какой-нибудь метод.
// Можно переделать CalDist и CalDir под методы пренадлежащие User и переназвать методы как (fromUser *User) CalDistTo / CalDirTo
func CalculateDistance(fromUser *User, toUser *User) float64 {
	resX := fromUser.position.x - toUser.position.x
	resY := fromUser.position.y - toUser.position.y

	distance := math.Sqrt(math.Pow(float64(resX), 2) + math.Pow(float64(resY), 2))
	return math.Floor(distance*100) / 100 // Округляю в меньшюю сторону и убераю лишние цифры после запятой. Правда я еще не понял как оно убрало
}
