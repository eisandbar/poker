package internal

var checkers = []func(a, b []int) (int, []int, bool){isStraightFlush, isQuads, isFullHouse, isFlush, isStraight, isSet, isTwoPair, isPair}

func getValues(cards []int) []int {
	temp := make([]int, len(cards))
	for i, card := range cards {
		temp[i] = card / 4
	}
	return temp
}

func isPair(cards, values []int) (int, []int, bool) {
	for i := 1; i < len(values); i++ {
		if values[i] == values[i-1] {
			return Pair, []int{values[i]}, true
		}
	}
	return 0, nil, false
}

func isTwoPair(cards, values []int) (int, []int, bool) {
	pairs := []int{}
	for i := 1; i < len(values); i++ {
		if values[i] == values[i-1] {
			pairs = append(pairs, values[i])
			i++
		}
	}

	if len(pairs) == 2 {
		if pairs[1] > pairs[0] {
			pairs[0], pairs[1] = pairs[1], pairs[0]
		}
		return TwoPair, pairs, true
	}
	return 0, nil, false
}

func isSet(cards, values []int) (int, []int, bool) {
	for i := 2; i < len(values); i++ {
		if values[i] == values[i-1] && values[i] == values[i-2] {
			return Set, []int{values[i]}, true
		}
	}
	return 0, nil, false
}

func isFullHouse(cards, values []int) (int, []int, bool) {
	if len(values) != 5 {
		return 0, nil, false
	}
	if values[0] == values[1] && values[3] == values[4] && (values[2] == values[0] || values[2] == values[4]) {
		if values[2] == values[0] {
			return FullHouse, []int{values[2], values[4]}, true
		} else {

			return FullHouse, []int{values[2], values[0]}, true
		}
	}
	return 0, nil, false
}

func isStraight(cards, values []int) (int, []int, bool) {
	if len(values) != 5 {
		return 0, nil, false
	}
	if _, _, is := isPair(cards, values); is {
		return 0, nil, false
	}
	if values[0]-values[4] == 4 || values[0]-values[1] == 8 {
		return Straight, nil, true
	}
	return 0, nil, false
}

func isFlush(cards, values []int) (int, []int, bool) {
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

func isStraightFlush(cards, values []int) (int, []int, bool) {
	_, _, flush := isFlush(cards, values)
	_, _, straight := isStraight(cards, values)
	if flush && straight {
		return StraightFlush, nil, true
	}
	return 0, nil, false
}

func isQuads(cards, values []int) (int, []int, bool) {
	if len(values) != 5 {
		return 0, nil, false
	}
	if values[1] == values[2] && values[1] == values[3] && (values[0] == values[1] || values[4] == values[1]) {
		return Quads, []int{values[1]}, true
	}
	return 0, nil, false
}
