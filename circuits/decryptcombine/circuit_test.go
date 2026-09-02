package decryptcombine

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// Fixed context for the fixtures below. RoundHash must fit in 12 bytes and
// Aid in 32, because both feed the keccak organizer-share transcript.
var (
	testRoundHash = big.NewInt(5555)
	testAid       = big.NewInt(0xAA)
	testCtIdx     = big.NewInt(3)
	testSkOrg     = big.NewInt(4242)
)

// testAssignment builds a 1-of-1 combine over a real organizer share:
//
//	C1 = 5·G, Δ_org = sk_org·C1, δ_1 = 14·G, m = 3
//	C2 = m·G + λ_1·δ_1 + Δ_org  (λ_1 = 1 for a single share at x = 1)
func testAssignment() Assignment {
	c1Point := group.NewPoint()
	c1Point.ScalarBaseMult(big.NewInt(5))
	c1 := group.Encode(c1Point)

	deltaOrg, proof := mustOrganizerShare(c1)
	deltaOrgPoint, err := group.Decode(deltaOrg)
	if err != nil {
		panic(err)
	}

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
		DeltaOrg:           deltaOrg,
		OrganizerPK:        organizerPK(),
		OrganizerProof:     proof,
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

// mustOrganizerShare produces the organizer's Δ and DLEQ for the fixed test
// context, exactly the way an organizer's browser would.
func mustOrganizerShare(c1 types.CurvePoint) (types.CurvePoint, dleq.Proof) {
	var eid [12]byte
	testRoundHash.FillBytes(eid[:])
	var aid [32]byte
	testAid.FillBytes(aid[:])
	delta, proof, err := dleq.ProveOrganizerShare(eid, aid, uint16(testCtIdx.Uint64()), testSkOrg, c1)
	if err != nil {
		panic(err)
	}
	return delta, proof
}

// addGenerator returns p + G, used to build points that are well-formed but
// not what the DLEQ was proved over.
func addGenerator(p types.CurvePoint) types.CurvePoint {
	point, err := group.Decode(p)
	if err != nil {
		panic(err)
	}
	g := group.NewPoint()
	g.ScalarBaseMult(big.NewInt(1))
	sum := group.NewPoint()
	sum.Add(point, g)
	return group.Encode(sum)
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.PlaintextHash, qt.Not(qt.IsNil))
}

// The public-input order is the contract's ABI: pi[0..10] must stay exactly
// as documented or DKGManager reads the wrong words.
func TestPublicInputAndTranscriptLayout(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	_, publicInputs, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	scalars := publicInputs.Scalars()
	c.Assert(scalars, qt.HasLen, 11)
	c.Assert(scalars[0].Cmp(testRoundHash), qt.Equals, 0)
	c.Assert(scalars[1].Cmp(testAid), qt.Equals, 0)
	c.Assert(scalars[2].Cmp(testCtIdx), qt.Equals, 0)
	c.Assert(scalars[3].Cmp(asn.DeltaOrg.X), qt.Equals, 0)
	c.Assert(scalars[4].Cmp(asn.DeltaOrg.Y), qt.Equals, 0)
	c.Assert(scalars[5].Int64(), qt.Equals, int64(1)) // threshold
	c.Assert(scalars[6].Int64(), qt.Equals, int64(1)) // share count
	c.Assert(scalars[7].Cmp(publicInputs.CombineHash), qt.Equals, 0)
	c.Assert(scalars[8].Cmp(publicInputs.PlaintextHash), qt.Equals, 0)
	c.Assert(scalars[9].Cmp(publicInputs.Challenge), qt.Equals, 0)
	c.Assert(scalars[10].Cmp(publicInputs.TranscriptCommitment), qt.Equals, 0)

	transcript := publicInputs.TranscriptScalars()
	c.Assert(transcript, qt.HasLen, TranscriptWords)
	c.Assert(transcript, qt.HasLen, 12+3*MaxShares)
	want := []*big.Int{
		asn.CiphertextC1.X, asn.CiphertextC1.Y, asn.CiphertextC2.X, asn.CiphertextC2.Y,
		asn.OrganizerPK.X, asn.OrganizerPK.Y,
		asn.OrganizerProof.A1.X, asn.OrganizerProof.A1.Y,
		asn.OrganizerProof.A2.X, asn.OrganizerProof.A2.Y,
		asn.OrganizerProof.Response,
	}
	for i, value := range want {
		c.Assert(transcript[i].Cmp(value), qt.Equals, 0, qt.Commentf("transcript word %d", i))
	}
	// w[11] is the keccak challenge the contract recomputes from calldata.
	var eid [12]byte
	testRoundHash.FillBytes(eid[:])
	var aid [32]byte
	testAid.FillBytes(aid[:])
	wantE := dleq.OrganizerShareChallenge(
		eid, aid, uint16(testCtIdx.Uint64()),
		asn.OrganizerPK, asn.CiphertextC1, asn.DeltaOrg,
		asn.OrganizerProof.A1, asn.OrganizerProof.A2,
	)
	c.Assert(transcript[11].Cmp(wantE), qt.Equals, 0)
	c.Assert(publicInputs.OrganizerE.Cmp(wantE), qt.Equals, 0)
	// Indexes then partials.
	c.Assert(transcript[12].Int64(), qt.Equals, int64(1))
	c.Assert(transcript[12+MaxShares].Cmp(asn.PartialDecryptions[0].X), qt.Equals, 0)
	c.Assert(transcript[12+MaxShares+1].Cmp(asn.PartialDecryptions[0].Y), qt.Equals, 0)
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
