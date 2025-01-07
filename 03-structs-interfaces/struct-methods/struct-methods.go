package main

import (
	"encoding/json"
	"fmt"
)

type MinhaString string

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func UpdateName(u *User, newName string) {
	u.Name = newName
}

func main() {
	user := User{}
	user2 := User{Name: "João", ID: 1}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user2.Name)
	fmt.Println(user2.ID)
	user.UpdateName("Jaquim")
	fmt.Println(user)
	UpdateName(&user2, "Maria")
	fmt.Println(user2)

	res, err := json.Marshal(user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
