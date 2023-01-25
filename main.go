package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type User struct {
	Name     string
	LastName string
	Contact  string
}

func main() {
	user1 := User{faker.FirstName(), faker.LastName(), faker.Phonenumber()}

	fmt.Printf("%#v", user1)
}
