package cards

import "fmt"

const (
	// Club or Clover
	Club Suit = `♣`
	// Diamond or Tile
	Diamond Suit = "♦"
	// Heart or ...
	Heart Suit = "♥"
	// Spade or Pike
	Spade Suit = "♠"

	// Ace is the numeric value of the Ace card
	Ace = 1
	// Jack is the numeric value of the Jack card
	Jack = 11
	// Queen is the numeric value of the Queen card
	Queen = 12
	// King is the numeric value of the King card
	King = 13
)

// Suit is one of several categories into which the cards of a deck are divided
type Suit string

// Card is a playing card, it's identify provided by the combination of a Value
// and a Suit
type Card struct {
	value string
	suit  Suit
}

// Value returns the string value of a playing card. This consists of Ace (1),
// 2 through to 10, Jack (11), Queen (12) and King (13)
func (c *Card) Value() string {
	return c.value
}

// Suit returns the suit of the card
func (c *Card) Suit() Suit {
	return c.suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Value(), c.Suit())
}

const ()
