package typing_test

import (
	"eisandbar/poker/typing"
	"eisandbar/poker/util"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	testData := []struct {
		name, card string
		err        error
	}{
		{"Empty", "", util.BadInput},
	}
	for _, tt := range testData {

		t.Run(tt.name, func(t *testing.T) {
			card, err := typing.NewCard(tt.card)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.card, card.Value())

		})
	}

}
