package internal

import (
	"eisandbar/poker/typing"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStraight(t *testing.T) {
	testData := [][]string{
		[]string{"9D", "8S", "7C", "6D", "5D"},
		[]string{"JC", "TC", "QH", "8S", "9D"},
	}
	for _, tt := range testData {
		cards, err := typing.FromStrings(tt)
		if err != nil {
			t.Fatal(err)
		}
		intCards := typing.ToInts(cards)
		assert.NoError(t, err)
		sort.Slice(intCards, func(i, j int) bool { return intCards[i] > intCards[j] })
		combo, _, is := isStraight(intCards)
		assert.Equal(t, true, is)
		assert.Equal(t, Straight, combo)
	}
}

func TestIsTwoPair(t *testing.T) {
	cards, err := typing.FromStrings([]string{"9D", "9S", "7C", "5S", "5D"})
	intCards := typing.ToInts(cards)
	assert.NoError(t, err)
	combo, _, is := isTwoPair(intCards)
	assert.Equal(t, true, is)
	assert.Equal(t, TwoPair, combo)
}
