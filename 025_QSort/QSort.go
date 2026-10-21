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
		sNumero = ordenarQSort(sNumero, 0, numero-1)
		fmt.Println("Arreglo ordenado: ", sNumero)
	}
}

func ordenarQSort(sArreglo []int, nInicio int, nFin int) []int {
	if nInicio < nFin {
		nPivot := particionQSort(sArreglo, nInicio, nFin)
		sArreglo = ordenarQSort(sArreglo, nInicio, nPivot-1)
		sArreglo = ordenarQSort(sArreglo, nPivot+1, nFin)
	}
	return sArreglo
}

func particionQSort(sArreglo []int, nInicio int, nFin int) int {
	nPivot := sArreglo[nFin]
	nI := nInicio - 1
	for i := nInicio; i < nFin; i++ {
		if sArreglo[i] <= nPivot {
			nI++
			sArreglo[nI], sArreglo[i] = sArreglo[i], sArreglo[nI]
		}
	}
	sArreglo[nI+1], sArreglo[nFin] = sArreglo[nFin], sArreglo[nI+1]
	return nI + 1
}
