package decryptcombine

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// A prover who knows log_G(δ_k) for one accepted partial (the ciphertext's
// encryptor colluding with that share holder) can pick any plaintext M' and
// solve for a λ vector that still satisfies C2 = M'·G + Σ λ_k δ_k + Δ_org.
// The circuit must therefore pin λ to the canonical Lagrange vector of the
// qualifying set, not merely check the decryption identity.
func TestDecryptCombineRejectsNonCanonicalLagrangeVector(t *testing.T) {
	c := qt.New(t)
	q := group.ScalarField()
	mul := func(a, b *big.Int) *big.Int { return new(big.Int).Mod(new(big.Int).Mul(a, b), q) }

	// 2-of-2: shares d_1 = 14, d_2 = 20 at x = 1, 2 ⇒ f(0) = 2·14 − 20 = 8.
	r := big.NewInt(5) // encryption randomness, known to the forger
	d := []*big.Int{big.NewInt(14), big.NewInt(20)}
	m := big.NewInt(3)
	c1 := group.NewPoint()
	c1.ScalarBaseMult(r)
	deltas := make([]types.CurvePoint, 2)
	for k := range d {
		p := group.NewPoint()
		p.ScalarBaseMult(mul(d[k], r))
		deltas[k] = group.Encode(p)
	}
	deltaOrg, proof := mustOrganizerShare(group.Encode(c1))
	deltaOrgPoint, err := group.Decode(deltaOrg)
	c.Assert(err, qt.IsNil)
	// C2 = m·G + f(0)·r·G + Δ_org
	c2 := group.NewPoint()
	c2.ScalarBaseMult(new(big.Int).Add(m, mul(big.NewInt(8), r)))
	c2.Add(c2, deltaOrgPoint)

	honest, _, err := BuildWitness(Assignment{
		RoundHash:          new(big.Int).Set(testRoundHash),
		Aid:                new(big.Int).Set(testAid),
		CtIdx:              new(big.Int).Set(testCtIdx),
		DeltaOrg:           deltaOrg,
		OrganizerPK:        organizerPK(),
		OrganizerProof:     proof,
		Threshold:          2,
		CiphertextC1:       group.Encode(c1),
		CiphertextC2:       group.Encode(c2),
		ParticipantIndexes: []uint16{1, 2},
		PartialDecryptions: deltas,
		Plaintext:          m,
	})
	c.Assert(err, qt.IsNil)
	assert := test.NewAssert(t)
	assert.CheckCircuit(&DecryptCombineCircuit{}, test.WithValidAssignment(honest), test.WithCurves(ecc.BN254))

	// Forge M' = m + 1 by shifting λ_1 by −1/(d_1·r): Σ λ'_k δ_k = Σ λ_k δ_k − G.
	forged := *honest
	forged.Plaintext = new(big.Int).Add(m, big.NewInt(1))
	forged.PlaintextHash = forged.Plaintext
	lambda1, ok := honest.LagrangeCoefficients[0].(*big.Int)
	c.Assert(ok, qt.IsTrue)
	inv := new(big.Int).ModInverse(mul(d[0], r), q)
	forged.LagrangeCoefficients[0] = new(big.Int).Mod(new(big.Int).Sub(lambda1, inv), q)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &forged, test.WithCurves(ecc.BN254))
}
