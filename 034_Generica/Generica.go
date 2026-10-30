package main

import "fmt"

func PrintList(lista ...any) {
	for _, valor := range lista {
		fmt.Printf(" Valor : %v \n", valor)
	}
}

func main() {
	PrintList("Victor", "Manuel", "Juan", "Zaira")
	PrintList(100, 210, 450, 310)
	PrintList("HOLA", 450, 0.2334433, true)
}
