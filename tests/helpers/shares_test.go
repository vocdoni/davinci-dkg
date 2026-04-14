package helpers

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRecoverParticipantShares(t *testing.T) {
	c := qt.New(t)

	shares, err := RecoverParticipantShares(
		[][]*big.Int{
			{big.NewInt(3), big.NewInt(1)},
			{big.NewInt(5), big.NewInt(2)},
		},
		[]uint16{1, 2},
	)
	c.Assert(err, qt.IsNil)
	c.Assert(len(shares), qt.Equals, 2)
	c.Assert(shares[0].Cmp(big.NewInt(11)), qt.Equals, 0)
	c.Assert(shares[1].Cmp(big.NewInt(14)), qt.Equals, 0)
}
