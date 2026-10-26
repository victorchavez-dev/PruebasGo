package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contacto struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}

func agregarContactosArchivo(contactos []Contacto) error {
	archivo, err := os.Create("contactos.json")
	if err != nil {
		return err
	}
	defer archivo.Close()
	encoder := json.NewEncoder(archivo)
	err = encoder.Encode(contactos)
	if err != nil {
		return err
	}
	return nil
}

func cargarContactosArchivo(contactos *[]Contacto) error {
	archivo, err := os.Open("contactos.json")
	if err != nil {
		return err
	}
	defer archivo.Close()

	decoder := json.NewDecoder(archivo)
	err = decoder.Decode(contactos)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contactos []Contacto
	err := cargarContactosArchivo(&contactos)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	leer := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("====GESTOR DE CONTACTOS=====\n")
		fmt.Print("1. AGREGAR CONTACTO\n")
		fmt.Print("2. VER CONTACTO\n")
		fmt.Print("3. SALIR\n")
		fmt.Print("Selecciona una opcion: ")

		var opcion int
		_, err = fmt.Scanln(&opcion)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		switch opcion {
		case 1:
			var c Contacto
			fmt.Print("Ingrese el nombre: ")
			c.Nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese el email: ")
			c.Email, _ = leer.ReadString('\n')

			fmt.Print("Ingrese el telefono: ")
			c.Telefono, _ = leer.ReadString('\n')

			contactos = append(contactos, c)

			if err := agregarContactosArchivo(contactos); err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			fmt.Println("Contacto agregado exitosamente")
		case 2:
			fmt.Println("==============================================")
			for index, contacto := range contactos {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contacto.Nombre, contacto.Email, contacto.Telefono)
			}
			fmt.Println("==============================================")
		case 3:
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}
