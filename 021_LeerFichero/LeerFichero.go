package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fichero, err := os.Open("datos.txt")

	if err != nil {
		fmt.Println("Error al abrir el fichero: ", err)
		return
	}
	defer fichero.Close()

	fmt.Println("\nContenido de 'datos.txt':")
	scanner := bufio.NewScanner(fichero)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error leyendo el fichero: %v\n", err)
	}
}
