package internal

import (
	"eisandbar/poker/typing"
)

func CalculateOdds(req typing.CardRequest) typing.Odds {

	player, board := typing.ToInts(req.PlayerHand), typing.ToInts(req.BoardCards)
	var odds typing.Odds
	oddsChan := make(chan typing.Odds)
	defer close(oddsChan)
	for _, opponentHand := range req.OpponentHands {
		// go func(hand []typing.Card) {

		opponent := typing.ToInts(opponentHand)
		// 	oddsChan <- recursive(player, opponent, board, typing.CardsInDeck)
		// }(opponentHand)
		// }
		// for i := 0; i < len(req.OpponentHands); i++ {
		// 	newOdds := <-oddsChan
		newOdds := recursive(player, opponent, board, typing.CardsInDeck)
		odds = merge(odds, newOdds)
	}
	return odds
}

func recursive(player, opponent, board []int, prev int) typing.Odds {
	var odds typing.Odds
	if len(board) < 5 { // fill board
		for i := prev - 1; i >= 0 && i < typing.CardsInDeck; i-- {
			temp := make([]int, len(board), len(board)+1)
			copy(temp, board)

			newOdds := recursive(player, opponent, append(temp, i), i)
			odds = merge(odds, newOdds)
		}
	} else { // calculate winner

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
