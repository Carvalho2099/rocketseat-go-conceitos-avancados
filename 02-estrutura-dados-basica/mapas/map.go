package main

import "fmt"

func main() {
	var m map[string]string
	m2 := make(map[string]string)
	m3 := map[string]string{
		"João":    "Pessoa",
		"Joaquim": "Pedro",
	}
	m4 := map[string][]int{
		"João": {1, 2, 3},
	}
	m5 := make(map[string]string)
	m5["Pedro"] = "Pessoa"
	valor, ok := m5["Pedro"]
	valor2, ok2 := m5["Pedro2"]

	fmt.Println(m == nil)
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5["Pedro"])
	fmt.Println(valor, ok)
	fmt.Println(valor2, ok2)
}
