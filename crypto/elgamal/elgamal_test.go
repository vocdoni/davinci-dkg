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

func TestApplicationKeyAddsOrganizerKey(t *testing.T) {
	c := qt.New(t)
	ep := group.NewPoint()
	ep.ScalarBaseMult(big.NewInt(7))
	org := group.NewPoint()
	org.ScalarBaseMult(big.NewInt(5))

	got, err := ApplicationKey(group.Encode(ep), group.Encode(org))
	c.Assert(err, qt.IsNil)
	want := group.NewPoint()
	want.ScalarBaseMult(big.NewInt(12))
	c.Assert(samePoint(got, group.Encode(want)), qt.IsTrue)

	_, err = ApplicationKey(types.CurvePoint{}, group.Encode(org))
	c.Assert(err, qt.Not(qt.IsNil))
	_, err = ApplicationKey(group.Encode(ep), types.CurvePoint{})
	c.Assert(err, qt.Not(qt.IsNil))
}

// Decrypting an application ciphertext needs both halves of PK_aid: the
// epoch secret alone recovers nothing.
func TestApplicationCiphertextNeedsBothSecrets(t *testing.T) {
	c := qt.New(t)
	skEp := big.NewInt(31337)
	skOrg := big.NewInt(4242)
	ep := group.NewPoint()
	ep.ScalarBaseMult(skEp)
	org := group.NewPoint()
	org.ScalarBaseMult(skOrg)
	pkAid, err := ApplicationKey(group.Encode(ep), group.Encode(org))
	c.Assert(err, qt.IsNil)

	m := big.NewInt(9)
	c1, c2, err := Encrypt(pkAid, m)
	c.Assert(err, qt.IsNil)
	c1p, err := group.Decode(c1)
	c.Assert(err, qt.IsNil)
	c2p, err := group.Decode(c2)
	c.Assert(err, qt.IsNil)

	subtract := func(terms ...*big.Int) types.CurvePoint {
		acc := group.NewPoint()
		acc.Set(c2p)
		for _, sk := range terms {
			scaled := group.NewPoint()
			scaled.ScalarMult(c1p, sk)
			neg := group.NewPoint()
			neg.Neg(scaled)
			acc.Add(acc, neg)
		}
		return group.Encode(acc)
	}

	want := group.NewPoint()
	want.ScalarBaseMult(m)
	c.Assert(samePoint(subtract(skEp, skOrg), group.Encode(want)), qt.IsTrue)
	c.Assert(samePoint(subtract(skEp), group.Encode(want)), qt.IsFalse)
	c.Assert(samePoint(subtract(skOrg), group.Encode(want)), qt.IsFalse)
}
