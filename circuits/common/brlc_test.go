package common

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBRLCNative(t *testing.T) {
	c := qt.New(t)

	commitment, err := BRLCNative(big.NewInt(5), big.NewInt(2), big.NewInt(3), big.NewInt(7))
	c.Assert(err, qt.IsNil)
	c.Assert(commitment.Cmp(big.NewInt(960)), qt.Equals, 0)
}

func TestBRLCNativeRejectsNilValue(t *testing.T) {
	c := qt.New(t)

	_, err := BRLCNative(big.NewInt(5), big.NewInt(2), nil)
	c.Assert(err, qt.ErrorMatches, "value 1 is nil")
}
