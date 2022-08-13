package internal

import (
	"eisandbar/poker/util"
)

// Given player hands and board cards, find the strongest combination for each player and decide who wins
const (
	High int = iota * 1000
	Pair
	TwoPair
	Set
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
)

// Given player and opponent hands and all boards cards assess who wins
func AssessWin(player, opponent, board []int) (int, error) {
	if len(board) != 5 || len(player) != 2 || len(opponent) != 2 {
		return 0, util.BadInput
	}

	if checkDuplicates([][]int{player, opponent, board}) {
		return 0, util.BadInput
	}

	playerCombo, playerCards := bestHand(player, board)
	opponentCombo, opponentCards := bestHand(opponent, board)
	return compareCombo(playerCombo, opponentCombo, playerCards, opponentCards)
}

func checkDuplicates(cardGroups [][]int) bool {
	seen := make([]bool, 52)
	for _, cards := range cardGroups {
		for _, card := range cards {
			if seen[card] {
				return true
			}
			seen[card] = true
		}
	}
	return false
}

func bestHand(hand, board []int) (int, []int) {
	cards := []int(board)
	combo := comboValue(board)
	for _, card := range hand {
		for i, _ := range board {
			temp := make([]int, 5)
			copy(temp, board)
			temp[i] = card
			combo, cards = updateCombo(combo, cards, temp)
		}
	}
	for i := range board {
		for j := i + 1; j < len(board); j++ {

			temp := make([]int, 5)
			copy(temp, board)
			temp[i] = hand[0]
			temp[j] = hand[1]
			combo, cards = updateCombo(combo, cards, temp)
		}
	}
	return combo, cards
}

func updateCombo(combo int, oldCards, cards []int) (int, []int) {
	newCombo := comboValue(cards)
	better, err := compareCombo(newCombo, combo, oldCards, cards)
	if err == nil && better == 1 {
		return newCombo, cards
	}
	return combo, oldCards
}

// Given 5 cards calculates their combo value
func comboValue(cards []int) int {
	sortCards(cards)
	for _, checker := range checkers {
		if combo, is := checker(cards); is {
			return combo
		}
	}

	return 0
}

// sorts cards in decreasing order
func sortCards(cards []int) {
	n := len(cards)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if cards[j] > cards[i] {
				cards[i], cards[j] = cards[j], cards[i]
			}
		}
	}
}

// Given two combo values, calculate which one is stronger
// 1 means comboA is stronger, -1 meas comboB is stronger and 0 means a draw
func compareCombo(comboA, comboB int, cardsA, cardsB []int) (int, error) {
	// Comparing the combo
	if comboA > comboB {
		return 1, nil
	}
	if comboA < comboB {
		return -1, nil
	}

	for i := 0; i < 5; i++ {
		if cardsA[i]/4 > cardsB[i]/4 {
			return 1, nil
		}
		if cardsA[i]/4 < cardsB[i]/4 {
			return -1, nil
		}
	}

	// both hands are equal in value
	return 0, nil
}
