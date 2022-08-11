package typing

type Odds struct {
	Win  float64
	Draw float64
	Lose float64
}

func (odds *Odds) Normalize() {
	sum := odds.Win + odds.Draw + odds.Lose
	odds.Win, odds.Draw, odds.Lose = odds.Win/sum, odds.Draw/sum, odds.Lose/sum
}
