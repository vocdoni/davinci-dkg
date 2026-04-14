package shamir

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestPolynomialEvaluate(t *testing.T) {
	c := qt.New(t)

	modulus := big.NewInt(97)
	poly, err := NewPolynomial([]*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}, modulus)
	c.Assert(err, qt.IsNil)

	value := poly.Evaluate(big.NewInt(4))

	c.Assert(value.Cmp(big.NewInt(49)), qt.Equals, 0)
}

func TestPolynomialValidateDegree(t *testing.T) {
	c := qt.New(t)

	poly, err := NewPolynomial([]*big.Int{big.NewInt(5), big.NewInt(3), big.NewInt(2)}, big.NewInt(97))
	c.Assert(err, qt.IsNil)

	c.Assert(poly.ValidateDegree(3), qt.IsNil)
	c.Assert(poly.ValidateDegree(2), qt.Not(qt.IsNil))
}
