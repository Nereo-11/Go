package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 0; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("El jugador eligió piedra")
		case 1:
			fmt.Println("El jugador eligió papel")
		case 2:
			fmt.Println("El jugador eligió tijeras")
		}

		fmt.Printf("Message:  %s\n", round.Message)
		fmt.Printf("ComputerChoice:  %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult:  %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt:  %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore:  %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore:  %s\n", round.PlayerScore)
	}

}
