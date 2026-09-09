package contribution

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// rebind recomputes the public inputs (commitment and share digests, the
// challenge and the BRLC commitment) from pub's public vectors and copies them
// into w, exactly as BuildWitness does, so that an adversarial witness fails
// only the check a test targets and not the transcript binding.
func rebind(c *qt.C, w *ContributionCircuit, pub *PublicInputs) {
	keyDigests := make([]*big.Int, MaxKeys)
	for j := range MaxKeys {
		padded, err := ccommon.PadPoints(pub.Commitments[j], MaxCoefficients)
		c.Assert(err, qt.IsNil)
		keyDigests[j], err = ccommon.CommitmentKeyDigestNative(padded)
		c.Assert(err, qt.IsNil)
	}
	commitmentHash, err := ccommon.CommitmentsHashNative(pub.RoundHash, pub.ContributorIndex, pub.Threshold, keyDigests)
	c.Assert(err, qt.IsNil)
	recipientIndexes, err := ccommon.PadBigInts(pub.RecipientIndexes, MaxRecipients)
	c.Assert(err, qt.IsNil)
	paddedKeys, err := ccommon.PadPoints(pub.RecipientKeys, MaxRecipients)
	c.Assert(err, qt.IsNil)
	transcript, err := pub.Transcript()
	c.Assert(err, qt.IsNil)
	ephemerals, err := ccommon.PadPoints(transcript.Ephemerals, MaxRecipients)
	c.Assert(err, qt.IsNil)
	rowDigests := make([]*big.Int, MaxRecipients)
	for i := range MaxRecipients {
		rowShares := make([]*big.Int, MaxKeys)
		for j := range MaxKeys {
			rowShares[j] = big.NewInt(0)
			if i < len(transcript.MaskedShares[j]) {
				rowShares[j] = transcript.MaskedShares[j][i]
			}
		}
		rowDigests[i], err = ccommon.EncryptedShareRowDigestNative(recipientIndexes[i], paddedKeys[i], ephemerals[i], rowShares)
		c.Assert(err, qt.IsNil)
	}
	shareHash, err := ccommon.EncryptedSharesHashNative(pub.RoundHash, pub.ContributorIndex, pub.CommitteeSize, rowDigests)
	c.Assert(err, qt.IsNil)
	words, err := pub.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	anchor, err := ccommon.ChallengeAnchor(words, commitmentHash, shareHash)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(pub.RoundHash, TranscriptDomain, anchor)
	c.Assert(err, qt.IsNil)
	commitment, err := ccommon.BRLCNative(challenge, words...)
	c.Assert(err, qt.IsNil)
	w.CommitmentHash, w.ShareHash, w.Challenge, w.TranscriptCommitment = commitmentHash, shareHash, challenge, commitment
}

// TestRebindReproducesBuildWitness pins the helper to BuildWitness: on an
// untouched witness it must reproduce the same public inputs.
func TestRebindReproducesBuildWitness(t *testing.T) {
	c := qt.New(t)
	witness, pub, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	rebound := *witness
	rebind(c, &rebound, pub)
	c.Assert(rebound.CommitmentHash.(*big.Int).Cmp(pub.CommitmentHash), qt.Equals, 0)
	c.Assert(rebound.ShareHash.(*big.Int).Cmp(pub.ShareHash), qt.Equals, 0)
	c.Assert(rebound.Challenge.(*big.Int).Cmp(pub.Challenge), qt.Equals, 0)
	c.Assert(rebound.TranscriptCommitment.(*big.Int).Cmp(pub.TranscriptCommitment), qt.Equals, 0)
}

// TestContributionCircuitRejectsOversizedShares is the regression for the
// canonical range check on shares. The tampered share is honest + q: the same
// group element (q·G = O), so the Feldman equation still holds; its masked
// ciphertext is shifted by q as well, so the mask equation still holds; and
// the digests and BRLC commitment are recomputed over the shifted transcript,
// so the binding checks still hold. Only the range check can reject it, and a
// refactor that dropped the range check would make this witness satisfy the
// circuit and fail the test.
func TestContributionCircuitRejectsOversizedShares(t *testing.T) {
	c := qt.New(t)

	witness, pub, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	q := group.ScalarField()
	p := group.BaseField()
	honest, ok := witness.Shares[0][0].(*big.Int)
	c.Assert(ok, qt.IsTrue)
	tampered := new(big.Int).Add(honest, q)
	c.Assert(tampered.Cmp(p) < 0, qt.IsTrue, qt.Commentf("honest + q must stay representable in the native field"))
	witness.Shares[0][0] = tampered

	masked, ok := witness.MaskedShares[0][0].(*big.Int)
	c.Assert(ok, qt.IsTrue)
	shifted := new(big.Int).Add(masked, q)
	shifted.Mod(shifted, p)
	witness.MaskedShares[0][0] = shifted
	pub.EncryptedShares[0][0].Ciphertext = shifted
	rebind(c, witness, pub)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitRejectsOversizedConstantTerms exercises the range
// check on the one coefficient that is still a witness: the constant term
// whose commitment pins the pool-key contribution. Higher coefficients are
// implied by the shares and carry cofactor preimages instead.
func TestContributionCircuitRejectsOversizedConstantTerms(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	rbjj := group.ScalarField()
	// Make key 0's constant term equal to r_bjj exactly (out of range by 1).
	asn.Coefficients[0][0] = new(big.Int).Set(rbjj)

	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitRejectsTorsionCommitments is the regression for the
// cofactor preimages: adding the order-2 point T = (0, p−1) to C_1 and C_2
// leaves every Feldman check intact (x·T + x²·T = x(x+1)·T = O for every
// integer x) but the perturbed commitments are outside the prime subgroup
// and have no preimage under multiplication by 8, so the proof must fail.
func TestContributionCircuitRejectsTorsionCommitments(t *testing.T) {
	c := qt.New(t)
	witness, pub, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	torsion := group.NewPoint().SetPoint(big.NewInt(0), new(big.Int).Sub(group.BaseField(), big.NewInt(1)))
	for _, m := range []int{1, 2} {
		point, err := group.Decode(pub.Commitments[0][m])
		c.Assert(err, qt.IsNil)
		point.Add(point, torsion)
		encoded := group.Encode(point)
		c.Assert(group.IsOnCurve(encoded.X, encoded.Y), qt.IsTrue)
		c.Assert(group.IsInPrimeSubgroup(encoded.X, encoded.Y), qt.IsFalse)
		pub.Commitments[0][m] = encoded
		witness.Commitments[0][m] = ccommon.CircuitPoint(encoded)
	}
	// The digests and the BRLC commitment are recomputed over the perturbed
	// commitments, so the transcript binding holds; the shares are unchanged
	// and still satisfy the Feldman equation at every position; only the
	// cofactor certificate (8·Q = C has no solution off the prime subgroup)
	// can reject the witness.
	rebind(c, witness, pub)
	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}
