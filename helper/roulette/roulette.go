package roulette

import (
	"fmt"
	"local/gogames/helper/utils"
)

func PlayGame() (bool, int) {
	var userBet int
	var userChoice string
	var randomNumber int
	var won bool

	randomNumber = utils.GenerateRandomNumber(36) // 0-36 for roulette

	fmt.Print("Enter your bet amount: ")
	fmt.Scan(&userBet)
	
	fmt.Print("Bet on [even/odd] or [number]: ")
	fmt.Scan(&userChoice)

	if userChoice == "even" {
		won = (randomNumber != 0 && randomNumber%2 == 0)
	} else if userChoice == "odd" {
		won = (randomNumber != 0 && randomNumber%2 == 1)
	} else {
		won = (randomNumber == 7) 
	}

	return won, randomNumber
}