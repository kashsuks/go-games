package randInt
import "fmt"

func main() int {
	rand.Seed(time.Now().UnixNano())

	var number int = rand.Intn(100) //0-99
	return number
}