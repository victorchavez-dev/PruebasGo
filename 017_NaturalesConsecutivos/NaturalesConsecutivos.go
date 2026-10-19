package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var err error

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un numero:")
	if lectura.Scan() {
		n, err = strconv.Atoi(lectura.Text())
		if err != nil {
			fmt.Println("Error: El numero debe ser un entero.")
			os.Exit(1)
		}

		if n <= 0 {
			fmt.Println("Error: El numero debe ser natural.")
			os.Exit(1)
		}

		suma := (n * (n + 1)) / 2
		fmt.Println("La suma de los numeros naturales consecutivos es:", suma)
	}
}
