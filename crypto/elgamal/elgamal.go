// Package elgamal implements exponential ElGamal on BabyJubJub as used by
// applications encrypting to a DKG key: (C1, C2) = (r·G, m·G + r·PK).
package elgamal

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// Encrypt encrypts the scalar m under pk with fresh randomness.
//
// There is deliberately no proof of knowledge of r: the submitter of an
// aggregated tally does not know the randomness of the aggregate, so a PoK
// would be incompatible with homomorphic aggregation. Cross-application
// replay is instead prevented by the per-application organizer key — a
// ciphertext copied into another application decrypts to sk_ep·C1, which is
// useless without that application's sk_org·C1.
func Encrypt(pk types.CurvePoint, m *big.Int) (c1, c2 types.CurvePoint, err error) {
	r, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return c1, c2, err
	}
	return encryptWithRandomness(pk, m, r)
}

// ApplicationKey derives the application encryption key from the epoch key
// and the application's organizer key: PK_aid = PK_ep + PK_org. Decryption
// therefore needs both the committee threshold (for sk_ep·C1) and the
// organizer's share (sk_org·C1).
func ApplicationKey(pkEp, pkOrg types.CurvePoint) (types.CurvePoint, error) {
	base, err := group.Decode(pkEp)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode epoch key: %w", err)
	}
	term, err := group.Decode(pkOrg)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode organizer key: %w", err)
	}
	sum := group.NewPoint()
	sum.Add(base, term)
	return group.Encode(sum), nil
}

func encryptWithRandomness(pk types.CurvePoint, m, r *big.Int) (c1, c2 types.CurvePoint, err error) {
	pkPoint, err := group.Decode(pk)
	if err != nil {
		return c1, c2, fmt.Errorf("decode pk: %w", err)
	}
	c1Point := group.NewPoint()
	c1Point.ScalarBaseMult(r)
	mG := group.NewPoint()
	mG.ScalarBaseMult(m)
	rPK := group.NewPoint()
	rPK.ScalarMult(pkPoint, r)
	c2Point := group.NewPoint()
	c2Point.Add(mG, rPK)
	return group.Encode(c1Point), group.Encode(c2Point), nil
}
