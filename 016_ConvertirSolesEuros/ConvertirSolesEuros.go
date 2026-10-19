package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var soles, tasacambio float64
	var err error

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el monto en Soles:")
	if lectura.Scan() {
		soles, err = strconv.ParseFloat(lectura.Text(), 64)
		if err != nil {
			fmt.Println("Error: El monto debe ser un numero")
			os.Exit(1)
		}
	}

	fmt.Println("Ingrese la tasa de cambio:")
	if lectura.Scan() {
		tasacambio, err = strconv.ParseFloat(lectura.Text(), 64)
		if err != nil {
			fmt.Println("Error: La tasa de cambio debe ser un numero")
			os.Exit(1)
		}
	}

	if soles != 0 && tasacambio != 0 {
		euro := math.Round((soles/tasacambio)*100) / 100
		fmt.Println("El monto en Euros es:", euro)
	} else {
		fmt.Println("Error: El monto y la tasa de cambio deben ser distintos de cero")
		os.Exit(1)
	}

}
