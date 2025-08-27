package main
import (
	"fmt"
	"math/rand"
	"time"
    "github.com/kashsuks/go-games/helper/numberGuesser"
    "github.com/kashsuks/go-games/helper/roulette"
    "github.com/kashsuks/go-games/helper/rps"
)

func main() {
	var gameNumber int

	fmt.Print("Would you like to play [1]Number Guessing Game, [2]Rock, Paper, Scissors, or [3]Roulette: ")
	fmt.Scan(&gameNumber)

	if gameNumber == 1{
		if numberGuesser.main() {
			fmt.Println("Good job! You got it right")
		} else {
			fmt.Println("That was wrong. Ahh bummer, maybe next time")
		}
	} else if gameNumber == 2 {
		ans, choice := rps.main()
		if ans == 1 {
			fmt.Printf("Good job! You won! The computer chose: %v", choice)
		} else if ans == 0 {
			fmt.Printf("You tied with the computer. The computer chose: %v", choice)
		} else if ans == -1 {
			fmt.Printf("You lost. Ahh bummer, maybe next time. The computer chose: %v", choice)
		} else {
			fmt.Println("Error")
		}
	} else if {

	} else {
		fmt.Println("Invalid choice, try again.")
	}
}