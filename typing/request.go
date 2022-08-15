package typing

import (
	"github.com/eisandbar/poker/util"
)

type Request struct {
	PlayerHand    []string `json:"player_hand"`
	OpponentHand  []string `json:"opponent_hand"`
	OpponentRange []string `json:"opponent_range"`
	BoardCards    []string `json:"board_cards"`
}

type CardRequest struct {
	PlayerHand    []Card
	OpponentHands [][]Card
	BoardCards    []Card
}

func (req Request) ConvertsRequest() (CardRequest, error) {
	// Function checks that request fits our requirements
	// PlayerHand is always length 2
	// OpponentHand is length 2 or OpponentRange isn't empty
	// BoardCards length is <= 5, and if there are board cards, there are at least 3 (flop)
	// All cards and ranges are correct
	// Returns the request converted to CardRequest
	if len(req.PlayerHand) != 2 {
		return CardRequest{}, util.BadInput
	}
	if !(len(req.OpponentHand) == 2 && len(req.OpponentRange) == 0) && !(len(req.OpponentHand) == 0 && len(req.OpponentRange) != 0) {
		return CardRequest{}, util.BadInput
	}
	if len(req.BoardCards) > 5 || (len(req.BoardCards) != 0 && len(req.BoardCards) < 3) {
		return CardRequest{}, util.BadInput
	}

	seenCards := make([]bool, 52)
	cardReq := CardRequest{}

	for _, cardString := range req.PlayerHand {
		card, err := NewCard(cardString)
		if err != nil || seenCards[card.ToInt()] {
			return CardRequest{}, util.BadInput
		}
		seenCards[card.ToInt()] = true
		cardReq.PlayerHand = append(cardReq.PlayerHand, card)
	}

	for _, cardString := range req.BoardCards {
		card, err := NewCard(cardString)
		if err != nil || seenCards[card.ToInt()] {
			return CardRequest{}, util.BadInput
		}
		seenCards[card.ToInt()] = true
		cardReq.BoardCards = append(cardReq.BoardCards, card)
	}

	if len(req.OpponentHand) == 2 {
		hand := make([]Card, 2)
		for i, cardString := range req.OpponentHand {
			card, err := NewCard(cardString)
			if err != nil || seenCards[card.ToInt()] {
				return CardRequest{}, util.BadInput
			}
			seenCards[card.ToInt()] = true
			hand[i] = card
		}
		cardReq.OpponentHands = append(cardReq.OpponentHands, hand)

	} else {

	}

	return cardReq, nil
}
