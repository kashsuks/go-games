package numberGuesser
import (
	"fmt"
	"github.com/kashsuks/go-games/helper/utils/randInt"
)

func main() bool {
	var userInput int
	var randomInt int
	var state bool

	randomInt = numberGenerator()

	fmt.Print("Enter your guess (between 0 and 99): ")
	fmt.Scan(&userInput)

	state = (userInput == randomInt)

	return state
}