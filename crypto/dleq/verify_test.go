package dleq

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

func TestProveVerify(t *testing.T) {
	c := qt.New(t)

	secret := big.NewInt(23)

	pubKey := group.NewPoint()
	pubKey.ScalarBaseMult(secret)

	base := group.NewPoint()
	base.ScalarBaseMult(big.NewInt(17))

	target := group.NewPoint()
	target.ScalarMult(base, secret)

	proof, err := Prove(secret, group.Encode(pubKey), group.Encode(base), group.Encode(target))
	c.Assert(err, qt.IsNil)
	c.Assert(Verify(*proof, group.Encode(pubKey), group.Encode(base), group.Encode(target)), qt.IsNil)
}

func TestVerifyRejectsTamperedTarget(t *testing.T) {
	c := qt.New(t)

	secret := big.NewInt(23)

	pubKey := group.NewPoint()
	pubKey.ScalarBaseMult(secret)

	base := group.NewPoint()
	base.ScalarBaseMult(big.NewInt(17))

	target := group.NewPoint()
	target.ScalarMult(base, secret)

	proof, err := Prove(secret, group.Encode(pubKey), group.Encode(base), group.Encode(target))
	c.Assert(err, qt.IsNil)

	tampered := group.NewPoint()
	tampered.ScalarBaseMult(big.NewInt(19))

	err = Verify(*proof, group.Encode(pubKey), group.Encode(base), group.Encode(tampered))
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "invalid dleq proof")
}
