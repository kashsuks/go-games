package main

import (
	"fmt"
	"local/gogames/helper/numberGuesser"
	"local/gogames/helper/roulette"
	"local/gogames/helper/rps"
)

func main() {
	var gameNumber int

	fmt.Print("Would you like to play [1]Number Guessing Game, [2]Rock, Paper, Scissors, or [3]Roulette: ")
	fmt.Scan(&gameNumber)

	if gameNumber == 1 {
		if numberGuesser.PlayGame() {
			fmt.Println("Good job! You got it right")
		} else {
			fmt.Println("That was wrong. Ahh bummer, maybe next time")
		}
	} else if gameNumber == 2 {
		ans, choice := rps.PlayGame()
		if ans == 1 {
			fmt.Printf("Good job! You won! The computer chose: %v\n", choice)
		} else if ans == 0 {
			fmt.Printf("You tied with the computer. The computer chose: %v\n", choice)
		} else if ans == -1 {
			fmt.Printf("You lost. Ahh bummer, maybe next time. The computer chose: %v\n", choice)
		} else {
			fmt.Println("Error")
		}
	} else if gameNumber == 3 {
		won, number := roulette.PlayGame()
		if won {
			fmt.Printf("Congratulations! You won! The number was: %d\n", number)
		} else {
			fmt.Printf("Sorry, you lost. The number was: %d\n", number)
		}
	} else {
		fmt.Println("Invalid choice, try again.")
	}
}