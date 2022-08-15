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
		[]string{"AC", "5C", "4H", "3S", "2D"},
	}
	for _, tt := range testData {
		cards, err := typing.FromStrings(tt)
		if err != nil {
			t.Fatal(err)
		}
		intCards := typing.ToInts(cards)
		assert.NoError(t, err)
		sort.Slice(intCards, func(i, j int) bool { return intCards[i] > intCards[j] })
		combo, is := isStraight(intCards)
		assert.Equal(t, true, is)
		assert.Equal(t, Straight, combo)
	}
}

func TestIsTwoPair(t *testing.T) {
	cards, err := typing.FromStrings([]string{"9D", "9S", "7C", "5S", "5D"})
	priority := getPriority(cards[1].ToInt(), cards[3].ToInt())
	intCards := typing.ToInts(cards)
	assert.NoError(t, err)
	combo, is := isTwoPair(intCards)
	assert.Equal(t, true, is)
	assert.Equal(t, TwoPair+priority, combo)
}

func TestCardSort(t *testing.T) {
	cards := []int{1, 2, 3, 7, 4, 5, 0}
	sortCards(cards)
	assert.Equal(t, []int{7, 5, 4, 3, 2, 1, 0}, cards)
}

func TestCompareCombo(t *testing.T) {
	cardsA := []int{31, 20, 17, 5, 1}
	cardsB := []int{29, 21, 19, 6, 3}
	res := compareCombo(0, 0, cardsA, cardsB)
	assert.Equal(t, 0, res)
}
