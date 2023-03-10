package luckygen

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt() int {
	return rand.Int()
}

func RandomRune() rune {
	// Choose a number representing the ASCII uppercase alphabet
	num := 65 + rand.Intn(25)
	return rune(num)
}
