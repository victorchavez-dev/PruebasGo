package main

import (
	"fmt"
	"os"
)

func main() {

	fichero, err := os.Create("datos.txt")

	if err != nil {
		fmt.Println("Error al crear el fichero: ", err)
		return
	}
	defer fichero.Close()

	_, err = fichero.WriteString("Hola, esta es una prueba de escritura en un fichero.\n")
	if err != nil {
		fmt.Println("Error al escribir en el fichero: ", err)
		return
	}
	pi := 3.14159
	_, err = fmt.Fprintf(fichero, "El valor de pi es: %f\n", pi)

	if err != nil {
		fmt.Println("Error al escribir en el fichero: ", err)
		return
	}

	_, err = fmt.Fprintln(fichero, "Fin del programa.")

	if err != nil {
		fmt.Println("Error al escribir en el fichero: ", err)
		return
	}

	fmt.Println("Fichero creado y escrito exitosamente.")

}
