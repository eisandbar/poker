package internal

import (
	"eisandbar/poker/typing"
)

func CalculateOdds(req typing.CardRequest) typing.Odds {
	var odds typing.Odds

	for _, hand := range req.OpponentHands {
		odds = merge(odds, recursive(req.PlayerHand, hand, req.BoardCards, 53))
	}
	return odds
}

func recursive(player, opponent, board []typing.Card, prev int) typing.Odds {

	var odds typing.Odds
	if len(board) < 5 {
		for i := prev - 1; i >= 0; i-- {
			temp, err := typing.CardFromInt(i)
			if err != nil {
				continue
			}
			odds = merge(odds, recursive(player, opponent, append(board, temp), i))
		}
	} else {

		res, err := AssessWin(player, opponent, board)
		if err != nil {
			return odds
		}
		switch res {
		case 1:
			odds.Win++
		case -1:
			odds.Lose++
		case 0:
			odds.Draw++
		}
	}
	return odds
}

func merge(a, b typing.Odds) typing.Odds {
	var odds typing.Odds
	odds.Win = a.Win + b.Win
	odds.Lose = a.Lose + b.Lose
	odds.Draw = a.Draw + b.Draw
	return odds
}
