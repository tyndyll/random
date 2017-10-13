package cards_test

import (
	"testing"

	"github.com/tyndyll/random/cards"
)

func TestNewDeck(t *testing.T) {

}

func TestDeck_Cut(t *testing.T) {
	d := cards.NewDeck()

	slice1, slice2 := d.Cut()

	if len(slice1)+len(slice2) != cards.DeckSize {
		t.Errorf(`Slice returned more cards than the deck contains`)
	}
}
