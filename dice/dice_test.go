package dice_test

import (
	"testing"

	"github.com/tyndyll/random/dice"
)

func TestRoll(t *testing.T) {
	// Dice type is D6
	dice_type := 6

	// Now do it a million times
	//
	// Yes... this is a completely arbitrary test that may pass in a scenario
	// where the underlying code is wrong, but we just don't hit it in the
	// case of randomness, but it's better than a punch in the mouth so to
	// speak....
	for i := 0; i < 1000000; i++ {
		result := dice.Roll(dice_type)
		if result > dice_type || result == 0 {
			t.Error(`Result is higher than the dice_type`)
		}
	}

}

func TestRollPool(t *testing.T) {
	// Roll 10 dice
	number_of_dice := 10
	// Dice type is D6
	dice_type := 6

	// Now do it a million times
	//
	// Yes... this is a completely arbitrary test that may pass in a scenario
	// where the underlying code is wrong, but we just don't hit it in the
	// case of randomness, but it's better than a punch in the mouth so to
	// speak....
	for i := 0; i < 1000000; i++ {
		result := dice.RollPool(number_of_dice, dice_type)
		if len(result) != number_of_dice {
			t.Error(`Returned more results than asked for`)
		}
		for _, val := range result {
			if val > dice_type || val == 0 {
				t.Error(`Result is higher than the dice_type`)
			}
		}
	}
}
