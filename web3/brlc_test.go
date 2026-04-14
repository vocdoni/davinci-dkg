package web3

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestDeriveBRLCChallengeDeterministic(t *testing.T) {
	c := qt.New(t)

	var roundID [12]byte
	copy(roundID[:], []byte("round-000001"))
	anchor := common.HexToHash("0x1234")

	first := DeriveBRLCChallenge(roundID, "contribution", anchor)
	second := DeriveBRLCChallenge(roundID, "contribution", anchor)

	c.Assert(first.Cmp(second), qt.Equals, 0)
}

func TestBRLCCommit(t *testing.T) {
	c := qt.New(t)

	commitment, err := BRLCCommit(big.NewInt(5), big.NewInt(2), big.NewInt(3), big.NewInt(7))
	c.Assert(err, qt.IsNil)
	c.Assert(commitment.Cmp(big.NewInt(960)), qt.Equals, 0)
}
