package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var lado float64
	var err error

	lectura := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el lado:")
	if lectura.Scan() {
		lado, err = strconv.ParseFloat(lectura.Text(), 64)
		if err != nil {
			fmt.Println("Error: El lado debe ser un numero.")
			os.Exit(1)
		}

		area := lado * lado
		fmt.Println("El Area del Cuadrado es :", area)
	}

}
