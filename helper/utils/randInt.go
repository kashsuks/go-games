package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	var number int = rand.Intn(max) // 0 to max inclusive
	return number
}