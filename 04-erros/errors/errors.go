package main

import (
	"errors"
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (e SqrtError) Error() string {
	return e.msg
}

func RaizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"Não é possível calcular a raiz quadrada de um número negativo"}
	}
	resultado := math.Sqrt(x)
	return resultado, nil
}

var ErrNotFound = errors.New("Not Found")

func foo() error { return ErrNotFound }

func main() {
	a := 10
	b := 0

	resultado, err := dividir(a, b)
	if err != nil {
		println(err.Error())
		return
	}
	println(resultado)

	x := -10
	res, err := RaizQuadrada(float64(x))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	err2 := foo()
	if err2 != nil && errors.Is(err2, ErrNotFound) {
		fmt.Println("Erro do tipo ErrNotFound")
		return
	}
	fmt.Println("Foi para fora")

	var SqrtError *SqrtError
	if err2 != nil && errors.As(err2, &SqrtError) {
		fmt.Println(SqrtError.msg)
		return
	}
	fmt.Println("Foi para fora 2")
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")
	}
	return a / b, nil
}
