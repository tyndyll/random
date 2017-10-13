package dice

import (
	"math/rand"
)

// Roll generates a random number between 1 and the number of sides
func Roll(number_of_sides int) int {
	return rand.Intn(number_of_sides-1) + 1
}

// RollPool generates a rolls a number of X sided dice and returns the generated
// results as a list of ints
func RollPool(number_of_dice, number_of_sides int) []int {
	pool := make([]int, number_of_dice, number_of_dice)
	for i := 0; i < number_of_dice; i++ {
		// Intn returns an int between 0 and n. We need to subtract one
		// to provide a value between 1 and the number of sides
		pool[i] = rand.Intn(number_of_sides-1) + 1
	}
	return pool
}
