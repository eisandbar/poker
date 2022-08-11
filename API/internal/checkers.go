package internal

import "sort"

var checkers = []func(card []int) (int, []int, bool){isStraightFlush, isQuads, isFullHouse, isFlush, isStraight, isSet, isTwoPair, isPair}

func getValues(cards []int) []int {
	temp := make([]int, len(cards))
	for i, card := range cards {
		temp[i] = card / 4
	}
	return temp
}

func isPair(cards []int) (int, []int, bool) {
	cards = getValues(cards)
	for i := 1; i < len(cards); i++ {
		if cards[i] == cards[i-1] {
			return Pair, []int{cards[i]}, true
		}
	}
	return 0, nil, false
}

func isTwoPair(cards []int) (int, []int, bool) {
	cards = getValues(cards)
	pairs := []int{}
	for i := 1; i < len(cards); i++ {
		if cards[i] == cards[i-1] {
			pairs = append(pairs, cards[i])
			i++
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i] > pairs[j] })
	if len(pairs) == 2 {
		return TwoPair, pairs, true
	}
	return 0, nil, false
}

func isSet(cards []int) (int, []int, bool) {
	cards = getValues(cards)
	for i := 2; i < len(cards); i++ {
		if cards[i] == cards[i-1] && cards[i] == cards[i-2] {
			return Set, []int{cards[i]}, true
		}
	}
	return 0, nil, false
}

func isFullHouse(cards []int) (int, []int, bool) {
	if len(cards) != 5 {
		return 0, nil, false
	}
	cards = getValues(cards)
	if cards[0] == cards[1] && cards[3] == cards[4] && (cards[2] == cards[0] || cards[2] == cards[4]) {
		if cards[2] == cards[0] {
			return FullHouse, []int{cards[2], cards[4]}, true
		} else {

			return FullHouse, []int{cards[2], cards[0]}, true
		}
	}
	return 0, nil, false
}

func isStraight(cards []int) (int, []int, bool) {
	if len(cards) != 5 {
		return 0, nil, false
	}
	if _, _, is := isPair(cards); is {
		return 0, nil, false
	}
	cards = getValues(cards)
	if cards[0]-cards[4] == 4 || cards[0]-cards[1] == 8 {
		return Straight, nil, true
	}
	return 0, nil, false
}

func isFlush(cards []int) (int, []int, bool) {
	if len(cards) != 5 {
		return 0, nil, false
	}
	suit := cards[0] % 4
	for _, card := range cards {
		if card%4 != suit {
			return 0, nil, false
		}
	}
	return Flush, nil, true
}

func isStraightFlush(cards []int) (int, []int, bool) {
	_, _, flush := isFlush(cards)
	_, _, straight := isStraight(cards)
	if flush && straight {
		return StraightFlush, nil, true
	}
	return 0, nil, false
}

func isQuads(cards []int) (int, []int, bool) {
	if len(cards) != 5 {
		return 0, nil, false
	}
	cards = getValues(cards)
	if cards[1] == cards[2] && cards[1] == cards[3] && (cards[0] == cards[1] || cards[4] == cards[1]) {
		return Quads, []int{cards[1]}, true
	}
	return 0, nil, false
}
