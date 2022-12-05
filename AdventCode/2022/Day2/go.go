package main

import (
	"fmt"
	"time"
)

const file = "actual.txt"

const (
	ROCK = iota + 1
	PAPER
	SCISSORS
)

const (
	WIN  = 6
	DRAW = 3
	LOSS = 0
)

const (
	OPPONENT_ELF_LOCATION = 0
	PLAYER_ELF_LOCATION   = 2
)

var (
	opponentElf = map[byte]int{
		'A': ROCK,
		'B': PAPER,
		'C': SCISSORS,
	}
	playerElf = map[byte]int{
		'X': ROCK,
		'Y': PAPER,
		'Z': SCISSORS,
	}
)

func scoreSelections(player, opponent int) int {
	if player == opponent {
		return DRAW
	}

	if player == ROCK && opponent == SCISSORS {
		return WIN
	}

	if player == PAPER && opponent == ROCK {
		return WIN
	}

	if player == SCISSORS && opponent == PAPER {
		return WIN
	}

	return LOSS
}

func DoQuestion(input []string) int {
	total := 0

	for _, line := range input {
		playerSelection := playerElf[line[PLAYER_ELF_LOCATION]]

		total += playerSelection

		opponentSelection := opponentElf[line[OPPONENT_ELF_LOCATION]]

		score := scoreSelections(playerSelection, opponentSelection)

		total += score
	}

	return total
}

func main() {
	start := time.Now()

	lines, err := readFile(file)
	if err != nil {
		fmt.Println("reading failed " + err.Error())
	}

	fmt.Printf("loaded file with %d lines\n", len(lines))

	fmt.Println("Running main")
	result := DoQuestion(lines)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
