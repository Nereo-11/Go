package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//book := book.Book{"Sangre de Campeón", "Sanchez", 114}

	myBook := book.NewBook("Ready Player One", "Nereo", 111)

	/*	var otherBook = book.Book{
			title: "Moby Dick",
			author: "Herman Melville",
		    pages: 300,
		}
	*/

	myBook.PrintInfo()
	myBook.SetTitle("Ready Player Two")
	fmt.Println(myBook.GetTitle())

	myBook.PrintInfo()

	myTextBook := book.NewTextBook("Telematica", "Hector", 243, "Santillana", "Universidad")
	myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	// Ejemplo de Interfaces

	/*
		miPerro := animal.Perro{Nombre: "Max"}
		miGato := animal.Gato{Nombre: "Zure"}

		animal.HacerSonido(&miPerro)
		animal.HacerSonido(&miGato)
	*/

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Toy"},
		&animal.Perro{Nombre: "Tom"},
		&animal.Gato{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
