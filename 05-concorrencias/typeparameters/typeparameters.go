package main

import "fmt"

func main() {
	foo("")
	foo(1)
	foo(1.2)
	foo(true)
	foo([]int{1, 2, 3})
	foo([]string{"a", "b", "c"})
}

func foo[T any](arg T) {
	fmt.Println(arg)
}
