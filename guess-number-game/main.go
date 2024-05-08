package main

import (
	"fmt"
	"math/rand"
)

func main() {

	play()

}

func play() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un Numero (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicidades, adivinaste el número")
			playAgain()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")
		}
	}

	fmt.Println("Se terminaron los intentos. El numero era:", numAleatorio)
	playAgain()
}

func playAgain() {
	var eleccion string
	fmt.Println("¿Quieres jugar nuevamente?  (S/n):")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Elección invalida")
		playAgain()
	}
}
