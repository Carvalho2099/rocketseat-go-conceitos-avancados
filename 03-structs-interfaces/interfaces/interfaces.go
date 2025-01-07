package main

import (
	"fmt"
	"interfaces/bar"
)

type Animal interface {
	sound() string
}

type Dog struct{}

func (d Dog) sound() string {
	return "Au Au"
}

func (Dog) Interface() {
	fmt.Println("Interface")
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.sound())
}

func main() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)
	bar.TakeFoo(dog)

	var a Animal
	var dog2 *Dog
	fmt.Println(a == nil)
	fmt.Println(dog2 == nil)
}
