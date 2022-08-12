package internal_test

import (
	"eisandbar/poker/internal"
	"eisandbar/poker/typing"
	"eisandbar/poker/util"
	"testing"

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
			odds: typing.Odds{},
		},
		{
			name: "Aces vs Kings",
			req: typing.Request{
				PlayerHand:   []string{"AH", "AS"},
				OpponentHand: []string{"KC", "KD"},
				BoardCards:   []string{}},
			odds: typing.Odds{.81, 0, .185},
		},
		{
			name: "Aces vs AK",
			req: typing.Request{
				PlayerHand:   []string{"AH", "AS"},
				OpponentHand: []string{"AC", "KH"},
				BoardCards:   []string{}},
			odds: typing.Odds{.92, 0, .058},
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

		})
	}
}
