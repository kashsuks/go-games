package main
import (
	"fmt"
	"math/rand"
	"time"
)

func numberGenerator() {
	rand.Seed(time.Now().UnixNano())

	var number int = rand.Intn(100)
	return number
}

func numberGuesser() {
	var int userInput
	var bool state

	randomInt = numberGenerator()

	fmt.Print("Enter your guess (between 0 and 99): ")
	fmt.Scan(&userInput)

	if userInput == randomInt {
		state = True
	} else {
		state = False
	}

	return state

}


func main() {
	numberGenerator()
}