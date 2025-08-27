package main
import (
	"fmt"
	"math/rand"
	"time"
)

func numberGenerator() int {
	rand.Seed(time.Now().UnixNano())

	var number int = rand.Intn(100) //0-99
	return number
}

func numberGuesser() bool {
	var userInput int
	var randomInt int
	var state bool

	randomInt = numberGenerator()

	fmt.Print("Enter your guess (between 0 and 99): ")
	fmt.Scan(&userInput)

	state = (userInput == randomInt)

	return state
}

func rps() (int, string) {
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

func main() {
	var gameNumber int

	fmt.Print("Would you like to play [1]Number Guessing Game or [2]Rock, Paper, Scissors: ")
	fmt.Scan(&gameNumber)

	if gameNumber == 1{
		if numberGuesser() {
			fmt.Println("Good job! You got it right")
		} else {
			fmt.Println("That was wrong. Ahh bummer, maybe next time")
		}
	} else if gameNumber == 2 {
		ans, choice := rps()
		if ans == 1 {
			fmt.Printf("Good job! You won! The computer chose: %v", choice)
		} else if ans == 0 {
			fmt.Printf("You tied with the computer. The computer chose: %v", choice)
		} else if ans == -1 {
			fmt.Printf("You lost. Ahh bummer, maybe next time. The computer chose: %v", choice)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Invalid choice, try again.")
	}
}