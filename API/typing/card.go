package typing

import (
	"errors"
	"regexp"
)

type Card struct {
	Value int // These go from 0-12 with 0 being the deuce and 12 being the Ace
	Suit  int // These go from 0-3 in the following order: Diamonds, Clubs, Hearts, Spades
}

var values = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = [4]string{"D", "C", "H", "S"}

func (card Card) ToString() string {
	return values[card.Value] + suits[card.Suit]
}

func (card Card) ToInt() int {
	return card.Suit*13 + card.Value
}

func (card *Card) FromString(input string) error {
	match, _ := regexp.MatchString("[2-9TJQKA][DCHS]", input)
	if !match {
		return errors.New("Bad Input")
	}
	for i, val := range values {
		if input[0:1] == val {
			card.Value = i
		}
	}
	for i, suit := range suits {
		if input[1:2] == suit {
			card.Suit = i
		}
	}
	return nil
}

func (card *Card) FromInt(input int) error {
	if input < 0 || input >= 52 {
		return errors.New("Bad Input")
	}
	card.Suit, card.Value = input/13, input%13
	return nil
}