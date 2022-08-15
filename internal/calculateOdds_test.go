package internal_test

import (
	"testing"

	"github.com/eisandbar/poker/internal"
	"github.com/eisandbar/poker/typing"

	"github.com/eisandbar/poker/util"

	"github.com/stretchr/testify/assert"
)

func TestCalculateOdds(t *testing.T) {
	testData := []struct {
		name string
		req  typing.Request
		odds typing.Odds
	}{
		{
			name: "Draw",
			req: typing.Request{
				PlayerHand:   []string{"KH", "KS"},
				OpponentHand: []string{"KC", "KD"},
				BoardCards:   []string{"JH", "TC", "4S"}},
			odds: typing.Odds{Draw: 1},
		},
		{
			name: "Aces vs Kings",
			req: typing.Request{
				PlayerHand:   []string{"AH", "AS"},
				OpponentHand: []string{"KC", "KD"},
				BoardCards:   []string{}},
			odds: typing.Odds{Win: .8106, Draw: .0038, Lose: .1855},
		},
		{
			name: "Aces vs AK",
			req: typing.Request{
				PlayerHand:   []string{"AH", "AS"},
				OpponentHand: []string{"AC", "KD"},
				BoardCards:   []string{}},
			odds: typing.Odds{Win: .9195, Draw: .0125, Lose: .0680},
		},
		{
			name: "J9 vs Q6",
			req: typing.Request{
				PlayerHand:   []string{"JH", "9H"},
				OpponentHand: []string{"QC", "6C"},
				BoardCards:   []string{}},
			odds: typing.Odds{Win: .4450, Draw: .0086, Lose: .5461},
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			cardReq, err := tt.req.ConvertsRequest()
			assert.NoError(t, err)
			got := internal.CalculateOdds(cardReq)

			t.Logf("%+v", got)
			got.Normalize()
			assert.InDelta(t, tt.odds.Win, got.Win, util.Accuracy)
			assert.InDelta(t, tt.odds.Lose, got.Lose, util.Accuracy)
			assert.InDelta(t, tt.odds.Draw, got.Draw, util.Accuracy)

		})
	}
}

func BenchmarkCalculateOdds(b *testing.B) {
	hands := [][]string{
		{"KC", "KD"},
		{"KC", "KD"},
		{"KC", "KD"},
		{"KC", "KD"},
		{"KC", "KD"},
		{"KC", "KD"},
	}
	playerHand, _ := typing.FromStrings([]string{"AH", "AS"})
	opponentHands := make([][]typing.Card, 0)
	for _, hand := range hands {
		temp, _ := typing.FromStrings(hand)
		opponentHands = append(opponentHands, temp)
	}
	b.ResetTimer()
	req := typing.CardRequest{
		PlayerHand:    playerHand,
		OpponentHands: opponentHands,
	}
	for i := 0; i < b.N; i++ {
		internal.CalculateOdds(req)
	}
}
