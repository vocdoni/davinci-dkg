package decryptcombine

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// The organizer's share is the only decryption input that is not backed by
// a Groth16 proof of its own — the contract stores it verbatim. Its
// soundness therefore rests entirely on the Chaum-Pedersen relation this
// circuit checks. Each case below forges one word of the share while
// keeping every digest self-consistent (the assignment is re-built, so
// CombineHash, ρ and the BRLC all match the forged words); the only
// constraint left to catch the forgery is the DLEQ itself.
func TestDecryptCombineRejectsForgedOrganizerShare(t *testing.T) {
	c := qt.New(t)
	assert := test.NewAssert(t)
	honest := testAssignment()

	t.Run("honest", func(t *testing.T) {
		witness, _, err := BuildWitness(honest)
		c.Assert(err, qt.IsNil)
		assert.SolvingSucceeded(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	// Δ' = Δ + G with C2' = C2 + G keeps the decryption identity intact, so
	// only z·C1 == A2 + e·Δ can reject it. Without the DLEQ a combiner could
	// pick any Δ and shift the recovered plaintext at will.
	t.Run("shifted-delta", func(t *testing.T) {
		forged := honest
		forged.DeltaOrg = addGenerator(honest.DeltaOrg)
		forged.CiphertextC2 = addGenerator(honest.CiphertextC2)
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	t.Run("tampered-z", func(t *testing.T) {
		forged := honest
		forged.OrganizerProof.Response = new(big.Int).Add(honest.OrganizerProof.Response, big.NewInt(1))
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	t.Run("wrong-organizer-key", func(t *testing.T) {
		forged := honest
		forged.OrganizerPK = addGenerator(honest.OrganizerPK)
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	t.Run("swapped-commitments", func(t *testing.T) {
		forged := honest
		forged.OrganizerProof.A1, forged.OrganizerProof.A2 = honest.OrganizerProof.A2, honest.OrganizerProof.A1
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})
}

// `e` is a transcript word, not something the circuit recomputes (keccak in
// a SNARK is far too expensive). The contract pins it to the keccak of the
// calldata; here we only need the algebraic half: a witness whose `e` is not
// the one the two verification equations were built for must not solve.
func TestDecryptCombineRejectsTamperedChallengeWord(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	e, ok := witness.OrganizerE.(*big.Int)
	c.Assert(ok, qt.IsTrue)
	witness.OrganizerE = new(big.Int).Add(e, big.NewInt(1))

	assert := test.NewAssert(t)
	assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// Both organizer scalars are range-bound to [0, r_bjj): a non-canonical
// z' = z + r_bjj satisfies the group equations but would give the same share
// two distinct on-chain encodings.
func TestDecryptCombineRejectsNonCanonicalOrganizerScalars(t *testing.T) {
	c := qt.New(t)
	order := group.ScalarField()
	assert := test.NewAssert(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	z, ok := witness.OrganizerZ.(*big.Int)
	c.Assert(ok, qt.IsTrue)
	tampered := *witness
	tampered.OrganizerZ = new(big.Int).Add(z, order)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &tampered, test.WithCurves(ecc.BN254))

	e, ok := witness.OrganizerE.(*big.Int)
	c.Assert(ok, qt.IsTrue)
	tampered = *witness
	tampered.OrganizerE = new(big.Int).Add(e, order)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &tampered, test.WithCurves(ecc.BN254))
}

// The organizer share is a mandatory addend of the decryption identity:
// C2 = m·G + Σ λ_k δ_k + Δ_org. A ciphertext assembled without it (the old
// public-derivation path) must no longer combine.
func TestDecryptCombineRequiresOrganizerShareInC2(t *testing.T) {
	c := qt.New(t)

	honest := testAssignment()
	c2Point, err := group.Decode(honest.CiphertextC2)
	c.Assert(err, qt.IsNil)
	deltaOrgPoint, err := group.Decode(honest.DeltaOrg)
	c.Assert(err, qt.IsNil)
	negDeltaOrg := group.NewPoint()
	negDeltaOrg.Neg(deltaOrgPoint)
	withoutShare := group.NewPoint()
	withoutShare.Add(c2Point, negDeltaOrg)

	forged := honest
	forged.CiphertextC2 = group.Encode(withoutShare)
	witness, _, err := BuildWitness(forged)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// Malformed organizer inputs must be caught at witness-build time rather
// than deep inside the solver.
func TestBuildWitnessRejectsMalformedOrganizerInputs(t *testing.T) {
	c := qt.New(t)

	base := testAssignment()

	missingProof := base
	missingProof.OrganizerProof.Response = nil
	_, _, err := BuildWitness(missingProof)
	c.Assert(err, qt.Not(qt.IsNil))

	nonCanonicalZ := base
	nonCanonicalZ.OrganizerProof.Response = new(big.Int).Add(base.OrganizerProof.Response, group.ScalarField())
	_, _, err = BuildWitness(nonCanonicalZ)
	c.Assert(err, qt.Not(qt.IsNil))

	oversizedEid := base
	oversizedEid.RoundHash = new(big.Int).Lsh(big.NewInt(1), 96)
	_, _, err = BuildWitness(oversizedEid)
	c.Assert(err, qt.Not(qt.IsNil))

	oversizedCtIdx := base
	oversizedCtIdx.CtIdx = big.NewInt(1 << 16)
	_, _, err = BuildWitness(oversizedCtIdx)
	c.Assert(err, qt.Not(qt.IsNil))
}
