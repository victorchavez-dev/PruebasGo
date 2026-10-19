package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ingresa tu nombre")
	lectura := bufio.NewScanner(os.Stdin)
	lectura.Scan()
	nombre := lectura.Text()
	fmt.Println("Hola ", nombre)
}
