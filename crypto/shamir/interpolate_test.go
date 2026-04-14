package shamir

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestInterpolateConstant(t *testing.T) {
	c := qt.New(t)

	modulus := big.NewInt(97)
	poly, err := NewPolynomial([]*big.Int{big.NewInt(42), big.NewInt(5), big.NewInt(9)}, modulus)
	c.Assert(err, qt.IsNil)

	secret, err := InterpolateConstant([]Share{
		{Index: 1, Value: poly.Evaluate(big.NewInt(1))},
		{Index: 2, Value: poly.Evaluate(big.NewInt(2))},
		{Index: 3, Value: poly.Evaluate(big.NewInt(3))},
	}, modulus)
	c.Assert(err, qt.IsNil)
	c.Assert(secret.Cmp(big.NewInt(42)), qt.Equals, 0)
}

func TestInterpolateConstantRejectsDuplicateIndex(t *testing.T) {
	c := qt.New(t)

	_, err := InterpolateConstant([]Share{
		{Index: 1, Value: big.NewInt(10)},
		{Index: 1, Value: big.NewInt(20)},
	}, big.NewInt(97))

	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "duplicate index")
}
