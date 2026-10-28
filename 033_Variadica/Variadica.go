package main

import "fmt"

func main() {
	fmt.Println("Total1: ", suma("Victor", 9, 2, 3, 4, 5))
	fmt.Println("Total2: ", suma("Manuel", 11, 12, 3, 4, 5))

	imprimirDatos("Victor Manuel", 9, true, 3, 4.45, 5, "chao")
	imprimirDatos("Zaira Marizet", true, 2.34)
}

func suma(nombre string, numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Printf("Nombre: %s, Total: %d\n", nombre, total)
	return total
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("Tipo: %T, Dato: %v\n", dato, dato)
	}

}
