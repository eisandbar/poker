package typing

import (
	"eisandbar/poker/util"

	"testing"

	"github.com/stretchr/testify/suite"
)

type CardTestSuite struct {
	suite.Suite
}

var testCard = Card{
	Value: 10, // Queen
	Suit:  2,  // Hearts
}

func (suite *CardTestSuite) TestToIntEmpty() {
	card := Card{}
	suite.Equal(0, card.ToInt())
}
func (suite *CardTestSuite) TestToInt() {
	card := testCard
	suite.Equal(36, card.ToInt())
}
func (suite *CardTestSuite) TestToStringEmpty() {
	card := Card{}
	suite.Equal("2D", card.ToString())
}

func (suite *CardTestSuite) TestToString() {
	card := testCard
	suite.Equal("QH", card.ToString())
}

func (suite *CardTestSuite) TestFromIntZero() {
	card, err := CardFromInt(0)
	suite.Equal(Card{}, card)
	suite.Equal(nil, err)
}
func (suite *CardTestSuite) TestFromInt() {
	card, err := CardFromInt(36)
	suite.Equal(testCard, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromIntBadInput() {
	_, err := CardFromInt(100)
	suite.Equal(util.BadInput, err)
}

func (suite *CardTestSuite) TestFromStringZero() {
	card, err := CardFromString("2D")
	suite.Equal(Card{}, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromString() {
	card, err := CardFromString("QH")
	suite.Equal(testCard, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromStringBadInput() {
	card, err := CardFromString("1D")
	suite.Equal(Card{}, card)
	suite.Equal(util.BadInput, err)
}

func TestCardTestSuite(t *testing.T) {
	suite.Run(t, new(CardTestSuite))
}
