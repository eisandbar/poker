package typing

import (
	"errors"
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
	card := Card{}
	err := card.FromInt(0)
	suite.Equal(Card{}, card)
	suite.Equal(nil, err)
}
func (suite *CardTestSuite) TestFromInt() {
	card := Card{}
	err := card.FromInt(36)
	suite.Equal(testCard, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromIntBadInput() {
	card := Card{}
	err := card.FromInt(100)
	suite.Equal(errors.New("Bad Input"), err)
}

func (suite *CardTestSuite) TestFromStringZero() {
	card := Card{}
	err := card.FromString("2D")
	suite.Equal(Card{}, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromString() {
	card := Card{}
	err := card.FromString("QH")
	suite.Equal(testCard, card)
	suite.Equal(nil, err)
}

func (suite *CardTestSuite) TestFromStringBadInput() {
	card := Card{}
	err := card.FromString("1D")
	suite.Equal(Card{}, card)
	suite.Equal(errors.New("Bad Input"), err)
}

func TestCardTestSuite(t *testing.T) {
	suite.Run(t, new(CardTestSuite))
}
