package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// Listar contactos
func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contact"

	//Ejecutar query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("------------------------------------------------------------------")

	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString
		rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin Correo Electrónico"
		}

		fmt.Printf("Id: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("----------------------------------------------------")
	}
}

func GetContactById(db *sql.DB, contactID int) {

	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	contact := models.Contact{}
	var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("no se encontró ningún contacto con el ID %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin Correo Electrónico"
	}

	fmt.Println("\nLISTA DE UN CONTACTO:")
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("Id: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("----------------------------------------------------")
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?,)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto creado con exito")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con exito")
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id= ?"

	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con exito")
}
