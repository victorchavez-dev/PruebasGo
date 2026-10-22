package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println("Slide despues de Ordenar")
		fmt.Println(mergeSort(sNumero, 0, numero-1))
	}

}

func mergeSort(arr []int, inicio, fin int) []int {
	if inicio < fin {
		medio := inicio + (fin-inicio)/2
		mergeSort(arr, inicio, medio)
		mergeSort(arr, medio+1, fin)
		merge(arr, inicio, medio, fin)
	}
	return arr
}

func merge(arr []int, inicio, medio, fin int) {
	tamano := fin - inicio + 1
	temp := make([]int, tamano)
	izq := inicio
	der := medio + 1
	pos := 0

	for izq <= medio && der <= fin {
		if arr[izq] <= arr[der] {
			temp[pos] = arr[izq]
			izq++
		} else {
			temp[pos] = arr[der]
			der++
		}
		pos++
	}

	for izq <= medio {
		temp[pos] = arr[izq]
		izq++
		pos++
	}

	for der <= fin {
		temp[pos] = arr[der]
		der++
		pos++
	}

	copy(arr[inicio:fin+1], temp)
}
