package decryptcombine

import (
	"context"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// automaticAssignment builds the same 1-of-1 combine as testAssignment for
// an automatic application: no organizer key at all, so PK_org is the
// identity, sk_org = 0 and Δ_org drops out of C2.
func automaticAssignment() Assignment {
	asn := testAssignment()
	c2Point := mustDecode(asn.CiphertextC2)
	negDeltaOrg := group.NewPoint()
	negDeltaOrg.Neg(mustDecode(organizerShare(asn.CiphertextC1)))
	c2Point.Add(c2Point, negDeltaOrg)

	asn.OrganizerPK = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	asn.OrganizerSecret = big.NewInt(0)
	asn.CiphertextC2 = group.Encode(c2Point)
	return asn
}

// The organizer half of the combine is now a knowledge statement: the
// prover must exhibit sk_org with PK_org = sk_org·G, and Δ_org is derived
// from it in circuit rather than read from calldata. A secret that is not
// the discrete log of the application's registered key must not solve —
// otherwise a combiner could invent any Δ_org and shift the plaintext.
func TestDecryptCombineRejectsWrongOrganizerSecret(t *testing.T) {
	c := qt.New(t)
	assert := test.NewAssert(t)
	honest := testAssignment()

	t.Run("honest", func(t *testing.T) {
		witness, _, err := BuildWitness(honest)
		c.Assert(err, qt.IsNil)
		assert.SolvingSucceeded(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	// sk' = sk + 1 with the registered PK_org: PK_org != sk'·G.
	t.Run("shifted-secret", func(t *testing.T) {
		forged := honest
		forged.OrganizerSecret = new(big.Int).Add(honest.OrganizerSecret, big.NewInt(1))
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	// A secret whose key is a different (well-formed) point: the whole
	// assignment is rebuilt so CombineHash, ρ and the BRLC are all
	// self-consistent, leaving PK_org == sk·G as the only check that can
	// reject it.
	t.Run("other-organizer-key", func(t *testing.T) {
		other := big.NewInt(7)
		pk := group.NewPoint()
		pk.ScalarBaseMult(other)
		forged := honest
		forged.OrganizerPK = group.Encode(pk)
		witness, _, err := BuildWitness(forged)
		c.Assert(err, qt.IsNil)
		assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})
}

// Automatic applications register the identity key and a zero secret. Both
// in-circuit multiplications must yield the identity for that witness:
// FixedBaseMul short-circuits every zero nibble, and gnark's fake-GLV
// ScalarMul has an explicit zero-scalar branch in its half-GCD hint. The
// full prove/verify (not just SolvingSucceeded) is what pins that the
// gadgets really do handle the degenerate case.
func TestDecryptCombineAcceptsAutomaticApplication(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(automaticAssignment())
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &DecryptCombineCircuit{})
	c.Assert(err, qt.IsNil)
	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(runtime.Verify(proof, publicInputs.PublicWitness()), qt.IsNil)
}

// The mirror image: a zero secret only opens the identity key. An
// organizer-locked application whose PK_org is a real point cannot be
// combined by claiming sk_org = 0, even when C2 was assembled without any
// organizer share (so the decryption identity holds with Δ_org = O).
func TestDecryptCombineRejectsZeroSecretForRealOrganizerKey(t *testing.T) {
	c := qt.New(t)

	forged := automaticAssignment()
	forged.OrganizerPK = organizerPK()
	witness, _, err := BuildWitness(forged)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// The organizer secret is range-bound to [0, r_bjj): sk' = sk + r_bjj
// multiplies to the same two points but is a second witness for the same
// statement.
func TestDecryptCombineRejectsNonCanonicalOrganizerSecret(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	secret, ok := witness.OrganizerSecret.(*big.Int)
	c.Assert(ok, qt.IsTrue)
	tampered := *witness
	tampered.OrganizerSecret = new(big.Int).Add(secret, group.ScalarField())

	assert := test.NewAssert(t)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &tampered, test.WithCurves(ecc.BN254))
}

// The organizer share is a mandatory addend of the decryption identity:
// C2 = m·G + Σ λ_k δ_k + Δ_org. A ciphertext under a real organizer key
// that was assembled without it must not combine.
func TestDecryptCombineRequiresOrganizerShareInC2(t *testing.T) {
	c := qt.New(t)

	forged := automaticAssignment()
	forged.OrganizerPK = organizerPK()
	forged.OrganizerSecret = new(big.Int).Set(testSkOrg)
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

	missingSecret := base
	missingSecret.OrganizerSecret = nil
	_, _, err := BuildWitness(missingSecret)
	c.Assert(err, qt.Not(qt.IsNil))

	nonCanonicalSecret := base
	nonCanonicalSecret.OrganizerSecret = new(big.Int).Add(base.OrganizerSecret, group.ScalarField())
	_, _, err = BuildWitness(nonCanonicalSecret)
	c.Assert(err, qt.Not(qt.IsNil))

	missingKey := base
	missingKey.OrganizerPK = types.CurvePoint{}
	_, _, err = BuildWitness(missingKey)
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
