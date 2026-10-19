package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numero int
	var err error

	fmt.Println("Ingresa Numero:")
	lectura := bufio.NewScanner(os.Stdin)
	if lectura.Scan() {
		numero, err = strconv.Atoi(lectura.Text())
		if err != nil {
			fmt.Println("Error al convertir el numero", err)
			os.Exit(1)
		}
		if numero%2 == 0 {
			fmt.Println("El numero es par")
		} else {
			fmt.Println("El numero es impar")
		}
	}
}
