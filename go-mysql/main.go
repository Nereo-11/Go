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
	//Establecer conexión
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Listar Contactos")
		fmt.Println("2. Obtener Contacto por Id")
		fmt.Println("3. Crear Nuevo Contacto")
		fmt.Println("4. Actualizar Contacto")
		fmt.Println("5. Eliminar Contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese Id del contacto")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactById(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Print("Ingrese Id del contacto que quiere eliminar")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción incorrecta o no válida")
		}

	}

	/*
		newContact := models.Contact{
			Id:    6,
			Name:  "Dulce",
			Email: "nuevo@example.com",
			Phone: "123456789",
		}

		handlers.UpdateContact(db, newContact)
		handlers.DeleteContact(db, 6)
		handlers.ListContacts(db)
	*/

}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("Ingrese Id del contacto que quiere editar")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo del contacto")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el telefono del contacto")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
