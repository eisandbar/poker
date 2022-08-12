package internal

import (
	"eisandbar/poker/util"
	"fmt"
)

// Given player hands and board cards, find the strongest combination for each player and decide who wins
// For each board, a player has 31 possible combinations

const (
	High int = iota
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
		return 0, fmt.Errorf("Wrong number of cards, %w", util.BadInput)
	}

	if checkDuplicates([][]int{player, opponent, board}) {
		return 0, fmt.Errorf("Duplicate cards detected, %w", util.BadInput)
	}

	playerCombo, playerPriority := bestHand(player, board)
	opponentCombo, opponentPriority := bestHand(opponent, board)
	return compareCombo(playerCombo, opponentCombo, playerPriority, opponentPriority)
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

	combo, priority := comboValue(board)
	for _, card := range hand {
		for i, _ := range board {
			temp := make([]int, 5)
			copy(temp, board)
			temp[i] = card
			combo, priority = updateCombo(combo, priority, temp)
		}
	}
	for i := range board {
		for j := i + 1; j < len(board); j++ {

			temp := make([]int, 5)
			copy(temp, board)
			temp[i] = hand[0]
			temp[j] = hand[1]
			combo, priority = updateCombo(combo, priority, temp)
		}
	}
	return combo, priority
}

func updateCombo(combo int, priority, cards []int) (int, []int) {
	newCombo, newPriority := comboValue(cards)
	better, err := compareCombo(newCombo, combo, newPriority, priority)
	if err == nil && better == 1 {
		return newCombo, newPriority
	}
	return combo, priority
}

// Given 5 cards calculates their combo value
func comboValue(cards []int) (int, []int) {
	sortCards(cards)
	values := getValues(cards)
	for _, checker := range checkers {
		if combo, priority, is := checker(cards, values); is {
			return combo, append(priority, values...)
		}
	}

	return 0, values
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
func compareCombo(comboA, comboB int, priorityA, priorityB []int) (int, error) {
	// Comparing the combo
	if comboA > comboB {
		return 1, nil
	}
	if comboA < comboB {
		return -1, nil
	}

	// comparing the kickers
	if len(priorityA) != len(priorityB) {
		return 0, util.BadInput
	}
	for i, value := range priorityA {
		if value > priorityB[i] {
			return 1, nil
		}
		if value < priorityB[i] {
			return -1, nil
		}
	}

	// both hands are equal in value
	return 0, nil
}
