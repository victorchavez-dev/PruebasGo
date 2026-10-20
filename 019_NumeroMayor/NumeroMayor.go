package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n1, n2 int
	var err error

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresa Cantidad de Numeros a Comparar: ")
	lectura.Scan()
	n1, err = strconv.Atoi(strings.TrimSpace(lectura.Text()))
	if err != nil {
		fmt.Println("Error al convertir el numero")
		return
	}

	sCompara := make([]int, n1)
	for i := 1; i <= n1; i++ {
		fmt.Printf("Ingresa Numero (%d de %d): ", i, n1)
		lectura.Scan()
		n2, err = strconv.Atoi(strings.TrimSpace(lectura.Text()))
		if err != nil {
			fmt.Println("Error al convertir el numero")
			return
		}
		sCompara[i-1] = n2
	}

	fmt.Println("Numeros Ingresados: ", sCompara)

	n2 = 0
	for _, numero := range sCompara {
		if numero > n2 {
			n2 = numero
		}
	}

	fmt.Println("Numero Mayor: ", n2)

}
