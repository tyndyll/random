package cards

import "fmt"

type Suit string

type Card struct {
	Value string
	Suit  Suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Value, c.Suit)
}

const (
	Club    Suit = `♣`
	Diamond Suit = "♦"
	Heart   Suit = "♥"
	Spade   Suit = "♠"
)
