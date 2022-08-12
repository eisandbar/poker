package internal

var checkers = []func(a []int) (int, []int, bool){isStraightFlush, isQuads, isFullHouse, isFlush, isStraight, isSet, isTwoPair, isPair}

func isPair(cards []int) (int, []int, bool) {
	for i := 1; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 {
			return Pair, []int{cards[i]}, true
		}
	}
	return 0, nil, false
}

func isTwoPair(cards []int) (int, []int, bool) {
	pairs := []int{}
	for i := 1; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 {
			pairs = append(pairs, cards[i])
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

func isSet(cards []int) (int, []int, bool) {
	for i := 2; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 && cards[i]/4 == cards[i-2]/4 {
			return Set, []int{cards[i]}, true
		}
	}
	return 0, nil, false
}

func isFullHouse(cards []int) (int, []int, bool) {
	if len(cards) != 5 {
		return 0, nil, false
	}
	if cards[0]/4 == cards[1]/4 && cards[3]/4 == cards[4]/4 && (cards[2]/4 == cards[0]/4 || cards[2]/4 == cards[4]/4) {
		if cards[2]/4 == cards[0]/4 {
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
	if cards[0]/4-cards[4]/4 == 4 || cards[0]/4-cards[1]/4 == 8 {
		if _, _, is := isPair(cards); is {
			return 0, nil, false
		}
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
	if cards[1]/4 == cards[2]/4 && cards[1]/4 == cards[3]/4 && (cards[0]/4 == cards[1]/4 || cards[4]/4 == cards[1]/4) {
		return Quads, []int{cards[1]}, true
	}
	return 0, nil, false
}
