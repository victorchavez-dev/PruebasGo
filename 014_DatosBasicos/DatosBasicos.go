package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nombre, apellido string
	var edad int
	var peso, altura float64
	var err error

	lectura := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el nombre:")
	lectura.Scan()
	nombre = lectura.Text()

	fmt.Println("Ingrese el apellido:")
	if lectura.Scan() {
		apellido = lectura.Text()
	}

	fmt.Println("Ingrese la edad:")
	if lectura.Scan() {
		edad, err = strconv.Atoi(lectura.Text())
		if err != nil {
			fmt.Println("Error: La Edad debe ser numero entero.")
			os.Exit(1)
		}
	}

	fmt.Println("Ingrese el peso:")
	lectura.Scan()
	peso, err = strconv.ParseFloat(lectura.Text(), 64)
	if err != nil {
		fmt.Println("Error: El peso debe ser un numero decimal.")
	}

	fmt.Println("Ingrese la altura:")
	lectura.Scan()
	altura, err = strconv.ParseFloat(lectura.Text(), 64)
	if err != nil {
		fmt.Println("Error: La altura debe ser un numero decimal.")
	}

	fmt.Println("Los Datos Ingresados son:")
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Altura: ", altura)
}
