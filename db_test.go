package main

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	db := ConnectToDB()
	defer db.Close()

	fmt.Println(db.Ping())
}
