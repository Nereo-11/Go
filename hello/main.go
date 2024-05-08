package main

import (
	"fmt"
	"log"

	"github.com/Nereo-11/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Nereo", "Uriel", "Uri"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	/*
		message, err := greetings.Hello("Alex")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)
	*/
	fmt.Println(messages)
}
