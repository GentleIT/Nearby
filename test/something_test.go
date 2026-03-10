package test

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	someArray := []int{1, 2, 3, 4}
	deletingInteger := 2
	someArray = append(someArray[:deletingInteger-1], someArray[2:]...)
	fmt.Println(someArray) // [1, 3, 4]
}
