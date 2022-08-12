package internal

import (
	"eisandbar/poker/typing"
	"eisandbar/poker/util"
	"fmt"
	"sort"
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
func AssessWin(playerHand, opponentHand, boardCards []typing.Card) (int, error) {
	if len(boardCards) != 5 || len(playerHand) != 2 || len(opponentHand) != 2 {
		return 0, util.BadInput
	}

	player, opponent, board := typing.ToInts(playerHand), typing.ToInts(opponentHand), typing.ToInts(boardCards)

	if checkDuplicates(append(append(player, opponent...), board...)) {
		return 0, fmt.Errorf("Duplicate cards detected, %w", util.BadInput)
	}

	playerCombo, playerPriority := bestHand(player, board)
	opponentCombo, opponentPriority := bestHand(opponent, board)
	return compareCombo(playerCombo, opponentCombo, playerPriority, opponentPriority)
}

func checkDuplicates(input []int) bool {
	sort.Ints(input)
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			return true
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
	sort.Slice(cards, func(i, j int) bool { return cards[i] > cards[j] })
	for _, checker := range checkers {
		if combo, priority, is := checker(cards); is {
			return combo, append(priority, getValues(cards)...)
		}
	}

	return 0, cards
}

// Given two combo values, calculate which one is stronger
// 1 means comboA is stronger, -1 meas comboB is stronger and 0 means a draw
func compareCombo(comboA, comboB int, priorityA, priorityB []int) (int, error) {
	if comboA > comboB {
		return 1, nil
	}
	if comboA < comboB {
		return -1, nil
	}
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
	return 0, nil
}
