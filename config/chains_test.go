package config

import (
	"math/big"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestChainByID(t *testing.T) {
	c := qt.New(t)

	op := ChainByID(10)
	c.Assert(op.BlockTime, qt.Equals, 2*time.Second)
	c.Assert(op.FallbackTipCap.Cmp(big.NewInt(1_000_000)), qt.Equals, 0)
	c.Assert(op.MaxFeePerGas.Cmp(big.NewInt(100_000_000)), qt.Equals, 0)

	gnosis := ChainByID(100)
	c.Assert(gnosis.BlockTime, qt.Equals, 5*time.Second)
	c.Assert(gnosis.FallbackTipCap, qt.IsNil)
	c.Assert(gnosis.MaxFeePerGas, qt.IsNil)

	for _, id := range []uint64{1, 11155111} {
		ch := ChainByID(id)
		c.Assert(ch.BlockTime, qt.Equals, 12*time.Second)
		c.Assert(ch.FallbackTipCap, qt.IsNil)
		c.Assert(ch.MaxFeePerGas, qt.IsNil)
	}

	c.Assert(ChainByID(31337), qt.DeepEquals, Chain{})
}

func TestChainByIDReturnsCopies(t *testing.T) {
	c := qt.New(t)
	a := ChainByID(10)
	a.FallbackTipCap.SetInt64(7)
	a.MaxFeePerGas.SetInt64(7)
	b := ChainByID(10)
	c.Assert(b.FallbackTipCap.Int64(), qt.Equals, int64(1_000_000))
	c.Assert(b.MaxFeePerGas.Int64(), qt.Equals, int64(100_000_000))
}
