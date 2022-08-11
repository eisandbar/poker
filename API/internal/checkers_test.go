package internal

import (
	"eisandbar/poker/typing"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValues(t *testing.T) {
	cards := []int{1, 2, 3, 27, 43}
	temp := getValues(cards)
	assert.NotEqual(t, cards, temp)
	assert.Equal(t, []int{0, 0, 0, 6, 10}, temp)
}

func TestIsStraight(t *testing.T) {
	testData := [][]string{
		[]string{"9D", "8S", "7C", "6D", "5D"},
		[]string{"JC", "TC", "QH", "8S", "9D"},
	}
	for _, tt := range testData {
		cards, err := typing.ConvertStrings(tt)
		assert.NoError(t, err)
		sort.Slice(cards, func(i, j int) bool { return cards[i] > cards[j] })
		combo, _, is := isStraight(cards)
		assert.Equal(t, true, is)
		assert.Equal(t, Straight, combo)
	}
}

func TestIsTwoPair(t *testing.T) {
	cards, err := typing.ConvertStrings([]string{"9D", "9S", "7C", "5S", "5D"})
	assert.NoError(t, err)
	combo, _, is := isTwoPair(cards)
	assert.Equal(t, true, is)
	assert.Equal(t, TwoPair, combo)
}
