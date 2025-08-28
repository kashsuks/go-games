package numberGuesser

import (
	"fmt"
	"local/gogames/helper/utils"
)

func PlayGame() bool {
	var userInput int
	var randomInt int
	var state bool

	randomInt = utils.GenerateRandomNumber(99)

	fmt.Print("Enter your guess (between 0 and 99): ")
	fmt.Scan(&userInput)

	state = (userInput == randomInt)

	return state
}