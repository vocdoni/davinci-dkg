package helpers

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/elgamal"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

func WaitCombinedDecryption(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	ciphertextIndex uint16,
) (web3.CombinedDecryptionView, error) {
	var record web3.CombinedDecryptionView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		record, fetchErr = services.Contracts.GetCombinedDecryption(ctx, epochID, [32]byte{}, ciphertextIndex)
		return fetchErr == nil && record.Completed
	})
	if err != nil {
		return web3.CombinedDecryptionView{}, err
	}
	return record, nil
}

// EncryptScalar is textbook ElGamal on BabyJubJub: (C1, C2) = (r·G, m·G + r·PK).
// It is what an application encrypts with before calling submitCiphertext.
func EncryptScalar(pk types.CurvePoint, m *big.Int) (c1, c2 types.CurvePoint, err error) {
	r, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return c1, c2, err
	}
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

// ProveCiphertext returns the Schnorr proof of knowledge of r for a test
// ciphertext with C1 = r·G, bound to (epochID, aid = 0). Panics on error:
// test-only helper.
func ProveCiphertext(epochID [12]byte, c1, c2 types.CurvePoint, r *big.Int) elgamal.PoK {
	pok, err := elgamal.ProveKnowledge(epochID, [32]byte{}, c1, c2, r)
	if err != nil {
		panic(err)
	}
	return pok
}

// AddPoints returns a + b on BabyJubJub (used to derive PK_aid = PK_ep + X).
func AddPoints(a, b types.CurvePoint) (types.CurvePoint, error) {
	pa, err := group.Decode(a)
	if err != nil {
		return types.CurvePoint{}, err
	}
	pb, err := group.Decode(b)
	if err != nil {
		return types.CurvePoint{}, err
	}
	sum := group.NewPoint()
	sum.Add(pa, pb)
	return group.Encode(sum), nil
}

// ScalarBasePoint returns s·G.
func ScalarBasePoint(s *big.Int) types.CurvePoint {
	p := group.NewPoint()
	p.ScalarBaseMult(s)
	return group.Encode(p)
}
