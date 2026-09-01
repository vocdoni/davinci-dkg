package node

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// A share-encryption nonce, a polynomial coefficient or a DLEQ witness that is
// small or predictable lets anyone unmask shares from calldata. Every secret
// scalar the node draws must be a fresh uniform element of the BabyJubJub
// scalar field.
func TestRandomScalarsAreDistinctFullRangeAndBelowOrder(t *testing.T) {
	c := qt.New(t)
	xs, err := randomScalars(8)
	c.Assert(err, qt.IsNil)
	c.Assert(xs, qt.HasLen, 8)

	order := group.ScalarField()
	big128 := new(big.Int).Lsh(big.NewInt(1), 128)
	seen := map[string]bool{}
	wide := 0
	for _, x := range xs {
		c.Assert(x.Sign(), qt.Equals, 1)
		c.Assert(x.Cmp(order), qt.Equals, -1)
		c.Assert(seen[x.String()], qt.IsFalse)
		seen[x.String()] = true
		if x.Cmp(big128) > 0 {
			wide++
		}
	}
	// 8 uniform draws below 2^128 out of a ~2^251 field is a 2^-984 event.
	c.Assert(wide, qt.Equals, 8)
}
