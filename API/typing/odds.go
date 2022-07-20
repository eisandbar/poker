package typing

type Odds struct {
	win  float64
	draw float64
	lose float64
}

func (odds *Odds) Normalize() {
	sum := odds.win + odds.draw + odds.lose
	odds.win, odds.draw, odds.lose = odds.win/sum, odds.draw/sum, odds.lose/sum
}
