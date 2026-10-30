package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error al abrir la conexion: ", err)
	}

	defer db.Close()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Listar Contactos")
		fmt.Println("2. Obtener Contacto por ID")
		fmt.Println("3. Crear Contacto")
		fmt.Println("4. Actualizar Contacto")
		fmt.Println("5. Eliminar Contacto")
		fmt.Println("6. Salir")

		fmt.Println("Selecione una opcion:")
		var opc int
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del Contacto:")
			var id int
			fmt.Scanln(&id)
			handlers.GetContactByID(db, id)
		case 3:
			newContacto := inputContactDetails(opc)
			handlers.CreateContact(db, newContacto)
		case 4:
			updateContacto := inputContactDetails(opc)
			handlers.UpdateContact(db, updateContacto)
			handlers.ListContacts(db)
		case 5:
			fmt.Println("Ingrese el Id del contacto que quiere eliminar:")
			var id int
			fmt.Scanln(&id)
			handlers.DeleteContact(db, id)
			handlers.ListContacts(db)
		case 6:
			fmt.Println("Saliendo del Programa")
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func inputContactDetails(option int) models.Contact {
	var contact models.Contact

	reader := bufio.NewReader(os.Stdin)

	if option == 4 {

		fmt.Print("Ingrese ID: ")
		var id int
		fmt.Scanln(&id)
		contact.Id = id

	}

	fmt.Print("Ingrese Nombre: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese Email: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese Telefono: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact

}
