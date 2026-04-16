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
	roundHash := big.NewInt(0xDEADBEEF)
	idx := big.NewInt(7)

	pubKey := group.NewPoint()
	pubKey.ScalarBaseMult(secret)

	base := group.NewPoint()
	base.ScalarBaseMult(big.NewInt(17))

	target := group.NewPoint()
	target.ScalarMult(base, secret)

	proof, err := Prove(secret, roundHash, idx, group.Encode(pubKey), group.Encode(base), group.Encode(target))
	c.Assert(err, qt.IsNil)
	c.Assert(Verify(*proof, roundHash, idx, group.Encode(pubKey), group.Encode(base), group.Encode(target)), qt.IsNil)

	// Tampering the round hash must reject the proof.
	wrongRound := big.NewInt(0xCAFEBABE)
	c.Assert(
		Verify(*proof, wrongRound, idx, group.Encode(pubKey), group.Encode(base), group.Encode(target)),
		qt.Not(qt.IsNil),
	)
	// Tampering the participant index must reject the proof.
	c.Assert(
		Verify(*proof, roundHash, big.NewInt(8), group.Encode(pubKey), group.Encode(base), group.Encode(target)),
		qt.Not(qt.IsNil),
	)
}

func TestVerifyRejectsTamperedTarget(t *testing.T) {
	c := qt.New(t)

	secret := big.NewInt(23)
	roundHash := big.NewInt(0xDEADBEEF)
	idx := big.NewInt(7)

	pubKey := group.NewPoint()
	pubKey.ScalarBaseMult(secret)

	base := group.NewPoint()
	base.ScalarBaseMult(big.NewInt(17))

	target := group.NewPoint()
	target.ScalarMult(base, secret)

	proof, err := Prove(secret, roundHash, idx, group.Encode(pubKey), group.Encode(base), group.Encode(target))
	c.Assert(err, qt.IsNil)

	tampered := group.NewPoint()
	tampered.ScalarBaseMult(big.NewInt(19))

	err = Verify(*proof, roundHash, idx, group.Encode(pubKey), group.Encode(base), group.Encode(tampered))
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "invalid dleq proof")
}
