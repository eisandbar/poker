package internal

var checkers = []func(a []int) (int, bool){isStraightFlush, isQuads, isFullHouse, isFlush, isStraight, isSet, isTwoPair, isPair}

func isPair(cards []int) (int, bool) {
	for i := 1; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 {
			return Pair + (cards[i] / 4), true
		}
	}
	return 0, false
}

func isTwoPair(cards []int) (int, bool) {
	pairs := 0
	for i := 1; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 {
			pairs++
			i++
		}
	}

	if pairs == 2 {
		// if there are 2 pairs, the second and fourth card are always part of them
		return TwoPair + (13*(cards[1]/4+1) + (cards[3] / 4)), true
	}
	return 0, false
}

func isSet(cards []int) (int, bool) {
	for i := 2; i < len(cards); i++ {
		if cards[i]/4 == cards[i-1]/4 && cards[i]/4 == cards[i-2]/4 {
			return Set + (cards[i] / 4), true
		}
	}
	return 0, false
}

func isFullHouse(cards []int) (int, bool) {
	if len(cards) != 5 {
		return 0, false
	}
	if cards[0]/4 == cards[1]/4 && cards[3]/4 == cards[4]/4 && (cards[2]/4 == cards[0]/4 || cards[2]/4 == cards[4]/4) {
		if cards[2]/4 == cards[0]/4 {
			return FullHouse + (13*(cards[2]/4+1) + (cards[4] / 4)), true
		} else {

			return FullHouse + (13*(cards[2]/4+1) + (cards[0] / 4)), true
		}
	}
	return 0, false
}

func isStraight(cards []int) (int, bool) {
	if len(cards) != 5 {
		return 0, false
	}
	if cards[0]/4-cards[4]/4 == 4 || cards[0]/4-cards[1]/4 == 8 {
		for i := 1; i < len(cards); i++ {
			if cards[i]/4 == cards[i-1]/4 {
				return 0, false
			}
		}
		return Straight, true
	}
	return 0, false
}

func isFlush(cards []int) (int, bool) {
	if len(cards) != 5 {
		return 0, false
	}
	suit := cards[0] % 4
	for _, card := range cards {
		if card%4 != suit {
			return 0, false
		}
	}
	return Flush, true
}

func isStraightFlush(cards []int) (int, bool) {
	_, flush := isFlush(cards)
	_, straight := isStraight(cards)
	if flush && straight {
		return StraightFlush, true
	}
	return 0, false
}

func isQuads(cards []int) (int, bool) {
	if len(cards) != 5 {
		return 0, false
	}
	if cards[1]/4 == cards[2]/4 && cards[1]/4 == cards[3]/4 && (cards[0]/4 == cards[1]/4 || cards[4]/4 == cards[1]/4) {
		return Quads + (cards[1] / 4), true
	}
	return 0, false
}
