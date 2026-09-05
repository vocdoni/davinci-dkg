package decryptcombine

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-node/crypto/ecc"
)

// Fixed context for the fixtures below. RoundHash must fit in 12 bytes and
// Aid in 32, because both are calldata-shaped on chain.
var (
	testRoundHash = big.NewInt(5555)
	testAid       = big.NewInt(0xAA)
	testCtIdx     = big.NewInt(3)
	testSkOrg     = big.NewInt(4242)
)

// testAssignment builds a 1-of-1 combine for an organizer-locked
// application whose secret has been revealed:
//
//	C1 = 5·G, Δ_org = sk_org·C1, δ_1 = 14·G, m = 3
//	C2 = m·G + λ_1·δ_1 + Δ_org  (λ_1 = 1 for a single share at x = 1)
func testAssignment() Assignment {
	c1Point := group.NewPoint()
	c1Point.ScalarBaseMult(big.NewInt(5))
	c1 := group.Encode(c1Point)

	deltaOrgPoint := mustDecode(organizerShare(c1))
	delta0Point := group.NewPoint()
	delta0Point.ScalarBaseMult(big.NewInt(14))
	messagePoint := group.NewPoint()
	messagePoint.ScalarBaseMult(big.NewInt(3))
	c2Point := group.NewPoint()
	c2Point.Set(messagePoint)
	c2Point.Add(c2Point, delta0Point)
	c2Point.Add(c2Point, deltaOrgPoint)

	return Assignment{
		RoundHash:          new(big.Int).Set(testRoundHash),
		Aid:                new(big.Int).Set(testAid),
		CtIdx:              new(big.Int).Set(testCtIdx),
		OrganizerPK:        organizerPK(),
		OrganizerSecret:    new(big.Int).Set(testSkOrg),
		Threshold:          1,
		CiphertextC1:       c1,
		CiphertextC2:       group.Encode(c2Point),
		ParticipantIndexes: []uint16{1},
		PartialDecryptions: []types.CurvePoint{group.Encode(delta0Point)},
		Plaintext:          big.NewInt(3),
	}
}

func organizerPK() types.CurvePoint {
	pk := group.NewPoint()
	pk.ScalarBaseMult(testSkOrg)
	return group.Encode(pk)
}

// organizerShare returns Δ_org = sk_org·C1 for the fixed test organizer,
// the value the circuit now derives from the secret instead of reading it
// from calldata.
func organizerShare(c1 types.CurvePoint) types.CurvePoint {
	delta := group.NewPoint()
	delta.ScalarMult(mustDecode(c1), testSkOrg)
	return group.Encode(delta)
}

func mustDecode(p types.CurvePoint) ecc.Point {
	point, err := group.Decode(p)
	if err != nil {
		panic(err)
	}
	return point
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.PlaintextHash, qt.Not(qt.IsNil))
}

// The public-input order is the contract's ABI: pi[0..8] must stay exactly
// as documented or DKGManager reads the wrong words.
func TestPublicInputAndTranscriptLayout(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	_, publicInputs, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	scalars := publicInputs.Scalars()
	c.Assert(scalars, qt.HasLen, 9)
	c.Assert(scalars[0].Cmp(testRoundHash), qt.Equals, 0)
	c.Assert(scalars[1].Cmp(testAid), qt.Equals, 0)
	c.Assert(scalars[2].Cmp(testCtIdx), qt.Equals, 0)
	c.Assert(scalars[3].Int64(), qt.Equals, int64(1)) // threshold
	c.Assert(scalars[4].Int64(), qt.Equals, int64(1)) // share count
	c.Assert(scalars[5].Cmp(publicInputs.CombineHash), qt.Equals, 0)
	c.Assert(scalars[6].Cmp(publicInputs.PlaintextHash), qt.Equals, 0)
	c.Assert(scalars[7].Cmp(publicInputs.Challenge), qt.Equals, 0)
	c.Assert(scalars[8].Cmp(publicInputs.TranscriptCommitment), qt.Equals, 0)

	transcript := publicInputs.TranscriptScalars()
	c.Assert(transcript, qt.HasLen, TranscriptWords)
	c.Assert(transcript, qt.HasLen, 6+3*MaxShares)
	want := []*big.Int{
		asn.CiphertextC1.X, asn.CiphertextC1.Y, asn.CiphertextC2.X, asn.CiphertextC2.Y,
		asn.OrganizerPK.X, asn.OrganizerPK.Y,
	}
	for i, value := range want {
		c.Assert(transcript[i].Cmp(value), qt.Equals, 0, qt.Commentf("transcript word %d", i))
	}
	// Indexes then partials.
	c.Assert(transcript[6].Int64(), qt.Equals, int64(1))
	c.Assert(transcript[6+MaxShares].Cmp(asn.PartialDecryptions[0].X), qt.Equals, 0)
	c.Assert(transcript[6+MaxShares+1].Cmp(asn.PartialDecryptions[0].Y), qt.Equals, 0)

	// BRLCCommitment is the contract's fold of that same vector.
	commitment, err := publicInputs.BRLCCommitment(publicInputs.Challenge)
	c.Assert(err, qt.IsNil)
	c.Assert(commitment.Cmp(publicInputs.TranscriptCommitment), qt.Equals, 0)
}

func TestDecryptCombineCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &DecryptCombineCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestDecryptCombineCircuitRejectsTamperedPlaintext(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.Plaintext = big.NewInt(123456)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &DecryptCombineCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestDecryptCombineArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
