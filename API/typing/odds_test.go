package typing

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OddsTestSuite struct {
	suite.Suite
}

var testOdds = Odds{
	win:  124,
	draw: 81,
	lose: 44,
}

const tolerance = 0.01

func (suite *OddsTestSuite) TestNormalizeSum() {
	odds := testOdds
	odds.Normalize()
	sum := odds.win + odds.draw + odds.lose
	suite.InDelta(1, sum, tolerance)
}

func (suite *OddsTestSuite) TestNormalizeWinLose() {
	odds := testOdds
	winLose := odds.win / odds.lose
	odds.Normalize()
	suite.InDelta(winLose, odds.win/odds.lose, tolerance)
}

func (suite *OddsTestSuite) TestNormalizeWinDraw() {
	odds := testOdds
	winLose := odds.win / odds.draw
	odds.Normalize()
	suite.InDelta(winLose, odds.win/odds.draw, tolerance)
}

func TestOddsTestSuite(t *testing.T) {
	suite.Run(t, new(OddsTestSuite))
}
