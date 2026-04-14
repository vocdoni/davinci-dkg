package feldman

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/shamir"
)

func TestVerifyShare(t *testing.T) {
	c := qt.New(t)

	modulus := group.ScalarField()
	poly, err := shamir.NewPolynomial([]*big.Int{big.NewInt(7), big.NewInt(11)}, modulus)
	c.Assert(err, qt.IsNil)

	commitments, err := Commitments(poly.Coefficients)
	c.Assert(err, qt.IsNil)

	share := poly.Evaluate(big.NewInt(2))
	c.Assert(VerifyShare(commitments, 2, share), qt.IsNil)
}

func TestVerifyShareRejectsTamperedShare(t *testing.T) {
	c := qt.New(t)

	modulus := group.ScalarField()
	poly, err := shamir.NewPolynomial([]*big.Int{big.NewInt(7), big.NewInt(11)}, modulus)
	c.Assert(err, qt.IsNil)

	commitments, err := Commitments(poly.Coefficients)
	c.Assert(err, qt.IsNil)

	share := new(big.Int).Add(poly.Evaluate(big.NewInt(2)), big.NewInt(1))
	share.Mod(share, modulus)

	err = VerifyShare(commitments, 2, share)
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "does not match")
}
