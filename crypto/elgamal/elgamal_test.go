package elgamal

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestEncryptDecryptsWithSecret(t *testing.T) {
	c := qt.New(t)
	sk := big.NewInt(123456789)
	pk := group.NewPoint()
	pk.ScalarBaseMult(sk)
	c1, c2, err := Encrypt(group.Encode(pk), big.NewInt(42))
	c.Assert(err, qt.IsNil)
	// m·G = C2 − sk·C1
	c1p, _ := group.Decode(c1)
	c2p, _ := group.Decode(c2)
	skC1 := group.NewPoint()
	skC1.ScalarMult(c1p, sk)
	neg := group.NewPoint()
	neg.Neg(skC1)
	mG := group.NewPoint()
	mG.Add(c2p, neg)
	want := group.NewPoint()
	want.ScalarBaseMult(big.NewInt(42))
	c.Assert(samePoint(group.Encode(mG), group.Encode(want)), qt.IsTrue)
}

func samePoint(a, b types.CurvePoint) bool {
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

func TestApplicationKeyModes(t *testing.T) {
	c := qt.New(t)
	ep := group.NewPoint()
	ep.ScalarBaseMult(big.NewInt(7))
	org := group.NewPoint()
	org.ScalarBaseMult(big.NewInt(5))
	// mode 0: (7 + 3)·G
	k0, err := ApplicationKey(group.Encode(ep), 0, big.NewInt(3), types.CurvePoint{})
	c.Assert(err, qt.IsNil)
	want := group.NewPoint()
	want.ScalarBaseMult(big.NewInt(10))
	c.Assert(samePoint(k0, group.Encode(want)), qt.IsTrue)
	// mode 1: (7 + 5)·G
	k1, err := ApplicationKey(group.Encode(ep), 1, nil, group.Encode(org))
	c.Assert(err, qt.IsNil)
	want.ScalarBaseMult(big.NewInt(12))
	c.Assert(samePoint(k1, group.Encode(want)), qt.IsTrue)
}

func TestPoKBindsRandomnessToCiphertextAndContext(t *testing.T) {
	c := qt.New(t)
	pk := group.NewPoint()
	pk.ScalarBaseMult(big.NewInt(99))
	eid := [12]byte{1}
	aid := [32]byte{2}
	c1, c2, pok, err := EncryptWithProof(eid, aid, group.Encode(pk), big.NewInt(7))
	c.Assert(err, qt.IsNil)
	c.Assert(VerifyPoK(eid, aid, c1, c2, pok), qt.IsTrue)
	// Same C1 under another application (the oracle attack) must not verify.
	c.Assert(VerifyPoK(eid, [32]byte{3}, c1, c2, pok), qt.IsFalse)
	// A re-randomised C1' = C1 + G with the old proof must not verify either.
	c1p, _ := group.Decode(c1)
	g := group.NewPoint()
	g.ScalarBaseMult(big.NewInt(1))
	shifted := group.NewPoint()
	shifted.Add(c1p, g)
	c.Assert(VerifyPoK(eid, aid, group.Encode(shifted), c2, pok), qt.IsFalse)
	// Tampered C2 must not verify.
	c.Assert(VerifyPoK(eid, aid, c1, c1, pok), qt.IsFalse)
}
