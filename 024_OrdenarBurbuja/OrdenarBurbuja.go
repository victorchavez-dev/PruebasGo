package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la cantidad de Elementos a registrar:")
	if lectura.Scan() {
		numero, err := strconv.Atoi(lectura.Text())
		if err != nil {
			fmt.Println("Error al convertir el número")
		}

		sNumero := make([]int, numero)

		for i := 0; i < numero; i++ {
			fmt.Printf("Ingresar el numero %d: ", i+1)
			if lectura.Scan() {
				sNumero[i], err = strconv.Atoi(lectura.Text())
				if err != nil {
					fmt.Println("Error al convertir el número")
				}
			}
		}

		fmt.Println("Arreglo original: ", sNumero)
		sNumero = ordenarBurbuja(sNumero)
		fmt.Println("Arreglo ordenado: ", sNumero)
	}

}

func ordenarBurbuja(sArreglo []int) []int {
	for i := 0; i < len(sArreglo); i++ {
		for j := 0; j < len(sArreglo)-1-i; j++ {
			if sArreglo[j] > sArreglo[j+1] {
				sArreglo[j], sArreglo[j+1] = sArreglo[j+1], sArreglo[j]
			}
		}
	}
	return sArreglo
}
