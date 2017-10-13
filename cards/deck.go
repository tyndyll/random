package cards

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	// DeckSize is the number of cards in the deck
	DeckSize = 52
	// SuitSize is the numer of cards in the suit
	SuitSize = 13
	// Ace is the numeric value of the Ace card
	Ace = 1
	// Jack is the numeric value of the Jack card
	Jack = 11
	// Queen is the numeric value of the Queen card
	Queen = 12
	// King is the numeric value of the King card
	King = 13
)

func init() {
	// Set the seed so we get random shuffling
	rand.Seed(time.Now().UTC().UnixNano())
}

// Deck is a structure that contains a fresh collection of 52 standard playing
// cards (Ace -> King in Clubs, Diamonds, Hearts and Spades)
type Deck struct {
	cards []*Card
}

// CardsRemaining reports how many cards are remaining in the deck
func (d *Deck) CardsRemaining() int {
	return len(d.cards)
}

// Cut slices the deck in two at an arbitrary location. It returns two slices
// of cards
func (d *Deck) Cut() ([]*Card, []*Card) {
	cut := rand.Intn(len(d.cards) - 1)
	return d.cards[:cut], d.cards[cut:]
}

// Deal returns a card from the deck. The card is subsequently removed from the
// deck and therefore unavailable.
func (d *Deck) Deal() *Card {
	var card *Card
	if d.CardsRemaining() > 0 {
		card, d.cards = d.cards[0], d.cards[1:]
	}
	return card
}

// Shuffle reorganises the cards remaining in the deck into a random order
func (d *Deck) Shuffle() {
	shuffleList := rand.Perm(DeckSize)
	var from *Card
	var to *Card
	for newPos, oldPos := range shuffleList {
		from = d.cards[oldPos]
		to = d.cards[newPos]

		d.cards[oldPos] = to
		d.cards[newPos] = from
	}
}

// NewDeck returns a new deck of 52 cards
func NewDeck() *Deck {
	deck := &Deck{
		cards: make([]*Card, DeckSize, DeckSize),
	}

	count := 0
	var value = ""

	for _, suit := range []Suit{Club, Diamond, Heart, Spade} {
		for i := 1; i <= SuitSize; i++ {
			switch i {
			case Ace:
				value = "A"
			case Jack:
				value = "J"
			case Queen:
				value = "Q"
			case King:
				value = "K"
			default:
				value = strconv.Itoa(i)
			}
			deck.cards[count] = &Card{Suit: suit, Value: value}
			count++
		}
	}
	return deck
}
