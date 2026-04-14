package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestNodeKeyValidate(t *testing.T) {
	c := qt.New(t)

	c.Run("accepts a well formed key", func(c *qt.C) {
		key := NodeKey{
			Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			PubX:     big.NewInt(11),
			PubY:     big.NewInt(22),
		}

		err := key.Validate()

		c.Assert(err, qt.IsNil)
	})

	c.Run("rejects zero operator", func(c *qt.C) {
		key := NodeKey{
			PubX: big.NewInt(11),
			PubY: big.NewInt(22),
		}

		err := key.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "operator")
	})

	c.Run("rejects nil coordinates", func(c *qt.C) {
		key := NodeKey{
			Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			PubX:     nil,
			PubY:     big.NewInt(22),
		}

		err := key.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "public key")
	})
}
