package typing

import (
	"github.com/eisandbar/poker/util"

	"testing"

	"github.com/stretchr/testify/suite"
)

type RequestTestSuite struct {
	suite.Suite
}

var testRequest = Request{ // When changing the struct for tests change the whole slice, not just an item
	PlayerHand:    []string{"AH", "QH"},
	OpponentHand:  []string{"AS", "KS"},
	OpponentRange: []string{},
	BoardCards:    []string{"4S", "5D", "JD"},
}

func (suite *RequestTestSuite) TestConvertEmpty() {
	req := Request{}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertMissingPlayerHand() {
	req := testRequest
	req.PlayerHand = []string{}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertMissingOpponentHand() {
	req := testRequest
	req.OpponentHand = []string{}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertOpponentHandAndRange() {
	req := testRequest
	req.OpponentRange = []string{"AKS", "AKO", "AA"}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertWrongCardNumber() {
	req := testRequest
	req.PlayerHand = append(req.PlayerHand, "JH")
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertBadString() {
	req := testRequest
	req.PlayerHand = []string{"AK", "QH"}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertPlayerCardOverlap() {
	req := testRequest
	req.PlayerHand = []string{"AH", "AH"}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertOpponentCardOverlap() {
	req := testRequest
	req.OpponentHand = []string{"AH", "KS"}
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertBoardCardOverlap() {
	req := testRequest
	req.BoardCards = append(req.BoardCards, "AH")
	cardReq, err := req.ConvertsRequest()
	suite.Equal(CardRequest{}, cardReq)
	suite.Equal(util.BadInput, err)
}

func (suite *RequestTestSuite) TestConvertGood() {
	req := testRequest
	suite.T().Log(req)
	cardReq, err := req.ConvertsRequest()
	suite.Equal("AH", cardReq.PlayerHand[0].Value())
	suite.Equal("QH", cardReq.PlayerHand[1].Value())
	suite.Equal("AS", cardReq.OpponentHands[0][0].Value())
	suite.Equal("KS", cardReq.OpponentHands[0][1].Value())
	suite.Equal("JD", cardReq.BoardCards[2].Value())
	suite.Equal(nil, err)
}

func TestRequestTestSuite(t *testing.T) {
	suite.Run(t, new(RequestTestSuite))
}
