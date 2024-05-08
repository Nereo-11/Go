package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"ComputerChoice"`
	RoundResult       string `json:"RoundResult"`
	ComputerChoiceInt int    `json:"ComputerChoiceInt"`
	ComputerScore     string `json:"ComputerScore"`
	PlayerScore       string `json:"PlayerScore"`
}

var winMessages = []string{
	"Bien echo!",
	"Buen Trabajo!",
	"Felicidades!",
}

var loseMessages = []string{
	"Lastima!",
	"Sera la proxima!",
	"Intentalo de nuevo!",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual!",
	"Intentalo de nuevo!",
	"Nadie gana, sera la proxima!",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió Piedra"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PapEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió Tijeras"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"

		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El Jugador Gana!"
		message = winMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}

}
