package main

import "fmt"

var filmesNoDB = []string{
	"O Poderoso Chefão",
	"Interestelar",
	"Batman",
	"Superman",
	"Vingadores",
	"O Senhor dos Anéis",
	"Harry Potter",
	"Star Wars",
	"Matrix",
	"O Rei Leão",
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	resultFromApi := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	var filmes []string

	for _, id := range resultFromApi {
		filmes = append(filmes, filmesNoDB[id])
	}

	fmt.Println("Filmes:", filmes)

	fmt.Println(slice, cap(slice), len(slice))
}
