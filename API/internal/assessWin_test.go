package internal_test

import (
	"eisandbar/poker/internal"
	"eisandbar/poker/typing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssessWin(t *testing.T) {
	playerHand, _ := typing.FromStrings([]string{"AH", "KH"})
	opponentHand, _ := typing.FromStrings([]string{"QH", "9D"})
	testBoards := []struct {
		name       string
		boardCards []string
		res        int
	}{
		{"Ace High", []string{"JC", "TC", "5C", "2S", "3D"}, 1},
		{"Ace Pair", []string{"AC", "TC", "5C", "2S", "3D"}, 1},
		{"Nine Set", []string{"9C", "9H", "AC", "KS", "3D"}, -1},
		{"Queen Pair", []string{"QC", "TC", "5C", "2S", "3D"}, -1},
		{"King Pair", []string{"QC", "KC", "5C", "2S", "3D"}, 1},
		{"Two Pair", []string{"QC", "KC", "9C", "2S", "3D"}, -1},
		{"Straight", []string{"JC", "TC", "AC", "8S", "3D"}, -1},
		{"Flush", []string{"JC", "TC", "8H", "2H", "3H"}, 1},
		{"Better Kicker", []string{"JC", "JS", "JD", "JH", "2H"}, 1},
		{"Draw", []string{"JC", "JS", "JD", "JH", "AS"}, 0},
	}

	for _, tt := range testBoards {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := typing.FromStrings(tt.boardCards)

			got, err := internal.AssessWin(playerHand, opponentHand, board)
			assert.NoError(t, err)
			assert.Equal(t, tt.res, got)
		})
	}
}

func TestAssessWinComplex(t *testing.T) {
	playerHand, _ := typing.FromStrings([]string{"TH", "KH"})
	opponentHand, _ := typing.FromStrings([]string{"TS", "QH"})
	testBoards := []struct {
		name       string
		boardCards []string
		res        int
	}{
		{"Pair Better Kicker", []string{"JC", "JD", "5C", "2S", "3D"}, 1},
		{"Two Pair Better Kicker", []string{"JC", "JD", "TC", "2S", "3D"}, 1},
		{"Better Pairs", []string{"9C", "9H", "TC", "KS", "AD"}, 1},
		{"Same FullHouse", []string{"JC", "JD", "JS", "TC", "3D"}, 0},
		{"Same FullHouse", []string{"JC", "JD", "JS", "AC", "AD"}, 0},
		{"Better FullHouse", []string{"QC", "QS", "KC", "KS", "3D"}, 1},
		{"Better FullHouse Queens", []string{"QC", "QS", "TC", "TD", "3D"}, -1},
	}

	for _, tt := range testBoards {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := typing.FromStrings(tt.boardCards)
			got, err := internal.AssessWin(playerHand, opponentHand, board)
			assert.NoError(t, err)
			assert.Equal(t, tt.res, got)
		})
	}
}

func TestAssessWinBadInput(t *testing.T) {
	playerHand, _ := typing.FromStrings([]string{"TH", "KH"})
	opponentHand, _ := typing.FromStrings([]string{"TS", "QH"})
	testBoards := []struct {
		name       string
		boardCards []string
	}{
		{"Repeating Ten", []string{"TH", "JD", "5C", "2S", "3D"}},
		{"Repeating Queen", []string{"JC", "QH", "TC", "2S", "3D"}},
		{"Repeating Board", []string{"9C", "9C", "TC", "KS", "AD"}},
		{"Bad Card", []string{"JC", "FD", "JS", "TC", "3D"}},
	}

	for _, tt := range testBoards {
		t.Run(tt.name, func(t *testing.T) {

			board, _ := typing.FromStrings(tt.boardCards)
			_, err := internal.AssessWin(playerHand, opponentHand, board)
			assert.Error(t, err)
		})
	}
}
