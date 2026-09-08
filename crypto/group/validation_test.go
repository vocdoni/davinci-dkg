package group

import (
	"math/big"
	"testing"

	edbn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	qt "github.com/frankban/quicktest"
)

// The decoder never checked the curve equation, so IsOnCurve used to accept
// any canonical pair. It must evaluate the equation itself.
func TestIsOnCurveChecksTheEquation(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	gx, gy := new(big.Int), new(big.Int)
	params.Base.X.BigInt(gx)
	params.Base.Y.BigInt(gy)
	minusOne := new(big.Int).Sub(BaseField(), big.NewInt(1))

	c.Assert(IsOnCurve(gx, gy), qt.IsTrue, qt.Commentf("generator"))
	c.Assert(IsOnCurve(big.NewInt(0), big.NewInt(1)), qt.IsTrue, qt.Commentf("identity"))
	c.Assert(IsOnCurve(big.NewInt(0), minusOne), qt.IsTrue, qt.Commentf("order-2 point is on the curve"))
	c.Assert(IsInPrimeSubgroup(big.NewInt(0), minusOne), qt.IsFalse, qt.Commentf("but outside the prime subgroup"))
	c.Assert(IsInPrimeSubgroup(gx, gy), qt.IsTrue)

	for _, bad := range [][2]int64{{0, 0}, {1, 1}, {2, 3}, {1, 0}} {
		c.Assert(IsOnCurve(big.NewInt(bad[0]), big.NewInt(bad[1])), qt.IsFalse, qt.Commentf("(%d, %d)", bad[0], bad[1]))
		c.Assert(IsInPrimeSubgroup(big.NewInt(bad[0]), big.NewInt(bad[1])), qt.IsFalse)
	}
	c.Assert(IsOnCurve(BaseField(), big.NewInt(1)), qt.IsFalse, qt.Commentf("non-canonical x"))
}
