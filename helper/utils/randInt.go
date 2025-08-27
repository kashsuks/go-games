package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	var number int = rand.Intn(100) //0-99
	return number
}