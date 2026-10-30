package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Incluido[T comparable](lista []T, item T) bool {
	for _, v := range lista {
		if v == item {
			return true
		}
	}
	return false
}

func Filtro[T constraints.Ordered](lista []T, callback func(T) bool) []T {
	resultado := make([]T, 0, len(lista))
	for _, v := range lista {
		if callback(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func main() {
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println("Incluido", Incluido(lista, 3))
	fmt.Println("Incluido", Incluido(lista, 6))

	lista2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Incluido", Incluido(lista2, "c"))
	fmt.Println("Incluido", Incluido(lista2, "f"))

	fmt.Println(Filtro(lista, func(value int) bool { return value > 3 }))
	fmt.Println(Filtro(lista2, func(value string) bool { return value > "c" }))
}
