package main

import (
	"fmt"
	"github.com/goinaction/code/chapter5/listing68/counters"
	"github.com/goinaction/code/chapter5/listing74/entities"
)

type account struct {
	Name     string
	UserId   int32
	Email    string
	password string
}

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)

	admin := entities.Admin{
		Rights: 123,
	}
	admin.Name = "lcy"
	admin.Email = "lcyzzz@foxmail.com"
	fmt.Printf("admin: %v \n", admin)
	account := account{
		Name:     "1",
		UserId:   12,
		Email:    "2",
		password: "3",
	}
	fmt.Printf("account: %v \n", account)
}
