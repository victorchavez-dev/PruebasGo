package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	query := "Select * from contact"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("LISTA DE CONCTACTOS:")
	fmt.Println("-----------------------------------------------")
	for rows.Next() {
		contact := models.Contact{}
		var valueEmial sql.NullString

		if err := rows.Scan(&contact.Id, &contact.Name, &valueEmial, &contact.Phone); err != nil {
			log.Fatal(err)
		}

		if valueEmial.Valid {
			contact.Email = valueEmial.String
		} else {
			contact.Email = "Sin Correo Electronico"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("-----------------------------------------------")
}

func GetContactByID(db *sql.DB, id int) {
	contact := models.Contact{}
	var valueEmial sql.NullString

	query := "Select * from contact where id = ?"
	if err := db.QueryRow(query, id).Scan(&contact.Id, &contact.Name, &valueEmial, &contact.Phone); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Contacto no encontrado")
			return
		}
		log.Fatal(err)
	}

	if valueEmial.Valid {
		contact.Email = valueEmial.String
	} else {
		contact.Email = "Sin Correo Electronico"
	}

	fmt.Println("LISTA DE UN CONTACTO:")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "Insert into contact(name, email, phone) values (?, ?, ?)"
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se creo nuevo registro: ", id)
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "Update contact set name = ?, email = ?, phone = ? where id = ?"
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se actualizo el registro: ", rowsAffected)
}

func DeleteContact(db *sql.DB, id int) {
	query := "Delete from contact where id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se elimino el registro: ", rowsAffected)
}
