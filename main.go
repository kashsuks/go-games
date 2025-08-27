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

func main() {
	if numberGuesser() {
		fmt.Println("Good job! You got it right")
	} else {
		fmt.Println("That was wrong. Ahh bummer, maybe next time")
	}
}