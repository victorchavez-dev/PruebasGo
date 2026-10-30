package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeros interface {
	~int | ~float64 | ~float32 | ~uint
}

func Sumar[T int | float64](nums ...T) T {
	var resultado T
	for _, num := range nums {
		resultado += num
	}
	return resultado
}

func Sumar2[T Numeros](nums ...T) T {
	var resultado T
	for _, num := range nums {
		resultado += num
	}
	return resultado
}

func Sumar3[T constraints.Integer | constraints.Float](nums ...T) T {
	var resultado T
	for _, num := range nums {
		resultado += num
	}
	return resultado
}

func main() {
	fmt.Println("Sumar", Sumar(100, 210, 450, 310))
	fmt.Println("Sumar", Sumar(100.5, 210.1, 450.5, 310.5))

	fmt.Println("Sumar2", Sumar2(100, 210, 450, 310))
	fmt.Println("Sumar2", Sumar2(100.5, 210.1, 450.5, 310.5))

	fmt.Println("Sumar3", Sumar3(100, 210, 450, 310))
	fmt.Println("Sumar3", Sumar3(100.5, 210.1, 450.5, 310.5))
}
