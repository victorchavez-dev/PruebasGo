package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingresar la cantidad de Elementos a Comparar.")
	if lectura.Scan() {
		numero, err := strconv.Atoi(lectura.Text())
		if err != nil {
			fmt.Println("Error al convertir el numero: ", err)
			return
		}

		sNumero := make([]int, numero)

		for i := 0; i < numero; i++ {
			fmt.Printf("Ingresar el numero (%d): ", i+1)
			if lectura.Scan() {
				numero, err := strconv.Atoi(lectura.Text())
				if err != nil {
					fmt.Println("Error al convertir el numero: ", err)
					return
				}
				sNumero[i] = numero
			}
		}

		fmt.Println("Slide antes de Ordenar")
		fmt.Println(sNumero)
		slices.Sort(sNumero)
		fmt.Println("Slide despues de Ordenar")
		fmt.Println(sNumero)
	}

}
