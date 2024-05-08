package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de Contactos
type Contacto struct {
	Nombre string `json:"name"` //Para esas comillas Altgr + } + Espacio
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

func SaveContactsToFile(Contactos []Contacto) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(Contactos)
	if err != nil {
		return err
	}

	return nil
}

//Cargar contactos dedes archivo Json

func LoadContactsFromFile(contactos *[]Contacto) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contactos)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	var contactos []Contacto

	err := LoadContactsFromFile(&contactos)
	if err != nil {
		fmt.Println("Error al cargar contactos", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar Contactos\n",
			"2, Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una Opción")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción")
			return
		}

		switch option {
		case 1:

			var c Contacto
			fmt.Print("Nombre: ")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			contactos = append(contactos, c)

			if err := SaveContactsToFile(contactos); err != nil {
				fmt.Println("Error al guardar contactos")
			}

		case 2:
			fmt.Println("============================")
			for index, Contacto := range contactos {
				fmt.Println("%d. Nombre: %s Email: %s Telefono: %s \n",
					index+1, Contacto.Nombre, Contacto.Email, Contacto.Phone)
			}
			fmt.Println("=========================")

		case 3:
			fmt.Println("Saliendo...")
			return

		default:
			fmt.Println("Opción invalida")
		}

	}

}
