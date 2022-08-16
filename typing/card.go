package typing

import (
	"github.com/eisandbar/poker/util"

	"regexp"
)

type Card struct {
	card string
}

const CardsInDeck = 52

var values = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var suits = [4]string{"D", "C", "H", "S"}

func NewCard(input string) (Card, error) {
	match, _ := regexp.MatchString("[2-9TJQKA][DCHS]", input)
	if !match {
		return Card{}, util.BadInput
	}
	return Card{input}, nil
}

func (c Card) Value() string {
	return c.card
}

func (c Card) ToInt() int {
	res := 0
	for i, val := range values {
		if string(c.card[0]) == val {
			res += i * 4
		}
	}
	for i, suit := range suits {
		if string(c.card[1]) == suit {
			res += i
		}
	}
	return res
}

func CardFromInt(input int) (Card, error) {
	if input < 0 || input >= CardsInDeck {
		return Card{}, util.BadInput
	}
	return NewCard(values[input/4] + suits[input%4])
}

func ToInts(cards []Card) []int {
	res := make([]int, len(cards))
	for i, card := range cards {
		res[i] = card.ToInt()
	}
	return res
}

func FromStrings(cards []string) ([]Card, error) {
	res := make([]Card, len(cards))
	var err error
	for i, card := range cards {
		res[i], err = NewCard(card)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
