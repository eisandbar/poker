package typing

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OddsTestSuite struct {
	suite.Suite
}

var testOdds = Odds{
	Win:  124,
	Draw: 81,
	Lose: 44,
}

const tolerance = 0.01

func (suite *OddsTestSuite) TestNormalizeSum() {
	odds := testOdds
	odds.Normalize()
	sum := odds.Win + odds.Draw + odds.Lose
	suite.InDelta(1, sum, tolerance)
}

func (suite *OddsTestSuite) TestNormalizeWinLose() {
	odds := testOdds
	winLose := odds.Win / odds.Lose
	odds.Normalize()
	suite.InDelta(winLose, odds.Win/odds.Lose, tolerance)
}

func (suite *OddsTestSuite) TestNormalizeWinDraw() {
	odds := testOdds
	winLose := odds.Win / odds.Draw
	odds.Normalize()
	suite.InDelta(winLose, odds.Win/odds.Draw, tolerance)
}

func TestOddsTestSuite(t *testing.T) {
	suite.Run(t, new(OddsTestSuite))
}
