package internal

import (
	"eisandbar/poker/typing"
)

func CalculateOdds(req typing.CardRequest) typing.Odds {

	player, board := typing.ToInts(req.PlayerHand), typing.ToInts(req.BoardCards)
	var odds typing.Odds
	oddsChan := make(chan typing.Odds)
	for _, hand := range req.OpponentHands {
		opponent := typing.ToInts(hand)
		go recursive(oddsChan, player, opponent, board, 52)
		newOdds := <-oddsChan
		odds = merge(odds, newOdds)
	}
	return odds
}

func recursive(oddsChan chan<- typing.Odds, player, opponent, board []int, prev int) {
	var odds typing.Odds
	if len(board) < 5 { // fill board
		newOddsChan := make(chan typing.Odds)
		for i := prev - 1; i >= 0 && i < 52; i-- {
			temp := make([]int, len(board), len(board)+1)
			copy(temp, board)

			go recursive(newOddsChan, player, opponent, append(temp, i), i)

		}
		for i := prev - 1; i >= 0; i-- {
			newOdds := <-newOddsChan
			odds = merge(odds, newOdds)
		}
	} else { // calculate winner

		res, err := AssessWin(player, opponent, board)
		if err != nil {
			oddsChan <- typing.Odds{}
			return
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
	oddsChan <- odds
}

func merge(a, b typing.Odds) typing.Odds {
	var odds typing.Odds
	odds.Win = a.Win + b.Win
	odds.Lose = a.Lose + b.Lose
	odds.Draw = a.Draw + b.Draw
	return odds
}
