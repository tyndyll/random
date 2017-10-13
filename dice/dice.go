package dice

import (
	"math/rand"
)

// Roll generates a random number between 1 and the number of sides
func Roll(numberOfSides int) int {
	return rand.Intn(numberOfSides-1) + 1
}

// RollPool generates a rolls a number of X sided dice and returns the generated
// results as a list of ints
func RollPool(numberOfDice, numberOfSides int) []int {
	pool := make([]int, numberOfDice, numberOfDice)
	for i := 0; i < numberOfDice; i++ {
		// Intn returns an int between 0 and n. We need to subtract one
		// to provide a value between 1 and the number of sides
		pool[i] = rand.Intn(numberOfSides-1) + 1
	}
	return pool
}
