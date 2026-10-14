package main

import "fmt"

func main() {
	var (
		nombre string
		edad   int
		sexo   string
		estado bool
	)

	nombre = "Victor"
	edad = 40
	sexo = "Masculino"
	estado = true

	fmt.Println(nombre, edad, sexo, estado)
}
