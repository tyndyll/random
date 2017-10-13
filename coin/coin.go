package coin

import (
	"math/rand"
)

const (
	// Head represents the Head side of a coin
	Head = iota
	// Tail represents the Tail side of a coin
	Tail = iota
)

// Flip returns an int to signify HEADS (0) or TAILS(1)
func Flip() int {
	return rand.Intn(1)
}
