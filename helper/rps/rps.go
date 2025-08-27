package rps

import (
	"fmt"
	"math/rand"
)

func PlayGame() (int, string) {
	var state int
	var userInput string
	var choices = [3]string{"rock", "paper", "scissors"}

	var randIdx int = rand.Intn(len(choices))
	var randChoice string = choices[randIdx]

	fmt.Print("Pick rock, paper, or scissors: ")
	fmt.Scan(&userInput)

	if userInput == "rock" {
		if randChoice == "rock" {
			state = 0
		} else if randChoice == "paper" {
			state = -1
		} else {
			state = 1
		}
	} else if userInput == "paper" {
		if randChoice == "rock" {
			state = 1
		} else if randChoice == "paper" {
			state = 0
		} else {
			state = -1
		}
	} else if userInput == "scissors" {
		if randChoice == "rock" {
			state = -1
		} else if randChoice == "paper" {
			state = 1
		} else {
			state = 0
		}
	} else {
		state = -2
	}

	return state, randChoice
}