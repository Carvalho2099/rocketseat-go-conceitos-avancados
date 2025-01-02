package main

import "fmt"

func main() {
	x := 10
	take(&x)
	//fmt.Println("Valor de x:", *p)
	//fmt.Println("Endereço de x:", p)
	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)

	p := create()
	fmt.Println("Valor de p:", *p)
}

func take(x *int) {
	*x = 100
}

func create() *int {
	x := 10
	return &x
}
