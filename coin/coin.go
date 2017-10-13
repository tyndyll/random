package coin

import (
	"math/rand"
)

const (
	Head = iota
	Tail = iota
)

// Flip returns an int to signify HEADS (0) or TAILS(1)
func Flip() int {
	return rand.Intn(1)
}
