package poolkey

import (
	"context"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

const (
	testThreshold = 3
	testCommittee = 3
	testKeyIndex  = 2
)

// testCoefficients builds one contributor's MaxK polynomials deterministically,
// distinct per contributor and per pool key.
func testCoefficients(contributorIndex uint16, threshold int) [][]*big.Int {
	sets := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		sets[j] = make([]*big.Int, threshold)
		for m := range threshold {
			sets[j][m] = big.NewInt(int64(int(contributorIndex)*1000 + (j+1)*10 + m + 1))
		}
	}
	return sets
}

// testCommitments turns one contributor's polynomials into the Feldman
// commitment points the activation prover reads from the contribution calldata.
func testCommitments(coefficients [][]*big.Int) [][]types.CurvePoint {
	commitments := make([][]types.CurvePoint, MaxKeys)
	for j := range MaxKeys {
		commitments[j] = make([]types.CurvePoint, len(coefficients[j]))
		for m, coefficient := range coefficients[j] {
			point := group.NewPoint()
			point.ScalarBaseMult(coefficient)
			commitments[j][m] = group.Encode(point)
		}
	}
	return commitments
}

// assignmentFor builds an activation over the given accepted members of a
// committee of `committee`, every member dealing testCoefficients.
func assignmentFor(threshold, committee uint16, accepted []uint16) Assignment {
	commitments := make([][][]types.CurvePoint, len(accepted))
	for i, index := range accepted {
		commitments[i] = testCommitments(testCoefficients(index, int(threshold)))
	}
	return Assignment{
		RoundHash:          big.NewInt(2222),
		Threshold:          threshold,
		CommitteeSize:      committee,
		KeyIndex:           testKeyIndex,
		ParticipantIndexes: accepted,
		Commitments:        commitments,
	}
}

func testAssignment() Assignment {
	indexes := make([]uint16, testCommittee)
	for i := range testCommittee {
		indexes[i] = uint16(i + 1)
	}
	return assignmentFor(testThreshold, testCommittee, indexes)
}

// expectedShareCommitment evaluates Σ_m p^m · Ā[m] natively from the
// accepted contributors' coefficients, independently of the witness builder.
func expectedShareCommitment(accepted []uint16, member int64) types.CurvePoint {
	sum := big.NewInt(0)
	for _, index := range accepted {
		coefficients := testCoefficients(index, testThreshold)[testKeyIndex]
		share, err := ccommon.EvaluatePolynomialNative(coefficients, big.NewInt(member))
		if err != nil {
			panic(err)
		}
		sum.Add(sum, share)
	}
	point := group.NewPoint()
	point.ScalarBaseMult(sum)
	return group.Encode(point)
}

func assertPointEqual(c *qt.C, got, want types.CurvePoint) {
	c.Helper()
	c.Assert(got.X.Cmp(want.X), qt.Equals, 0)
	c.Assert(got.Y.Cmp(want.Y), qt.Equals, 0)
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	witness, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))

	// P_j = Σ_i A_{i,j,0}: the pool key is the sum of the accepted
	// contributors' constant-term commitments for the activated key.
	expected := group.NewPoint()
	expected.SetZero()
	for i := range testCommittee {
		point := group.NewPoint()
		point.ScalarBaseMult(testCoefficients(uint16(i+1), testThreshold)[testKeyIndex][0])
		expected.Add(expected, point)
	}
	encoded := group.Encode(expected)
	assertPointEqual(c, publicInputs.PoolKey, encoded)
	assertPointEqual(c, publicInputs.AggregateCommitments[0], encoded)

	// D_p = Σ_m p^m Ā[m] for every member, the identity beyond the committee.
	c.Assert(publicInputs.ShareCommitments, qt.HasLen, MaxParticipants)
	for i := range testCommittee {
		assertPointEqual(c, publicInputs.ShareCommitments[i], expectedShareCommitment(assignment.ParticipantIndexes, int64(i+1)))
	}
	for i := testCommittee; i < MaxParticipants; i++ {
		assertPointEqual(c, publicInputs.ShareCommitments[i], ccommon.IdentityCurvePoint())
	}
}

// The public-input order is the contract's ABI: pi[0..7] must stay exactly
// as documented or DKGManager reads the wrong words. The digest is the
// Poseidon hash of (eid, keyIndex, transcript) and the anchor folds it in
// before the transcript's keccak.
func TestPublicInputLayoutAndDigest(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	scalars := publicInputs.Scalars()
	c.Assert(scalars, qt.HasLen, 8)
	c.Assert(scalars[0].Cmp(assignment.RoundHash), qt.Equals, 0)
	c.Assert(scalars[1].Int64(), qt.Equals, int64(testThreshold))
	c.Assert(scalars[2].Int64(), qt.Equals, int64(testCommittee))
	c.Assert(scalars[3].Int64(), qt.Equals, int64(testCommittee))
	c.Assert(scalars[4].Int64(), qt.Equals, int64(testKeyIndex))
	c.Assert(scalars[5].Cmp(publicInputs.TranscriptDigest), qt.Equals, 0)
	c.Assert(scalars[6].Cmp(publicInputs.Challenge), qt.Equals, 0)
	c.Assert(scalars[7].Cmp(publicInputs.TranscriptCommitment), qt.Equals, 0)

	transcript, err := publicInputs.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	c.Assert(transcript, qt.HasLen, TranscriptWords)

	digestInputs := append([]*big.Int{assignment.RoundHash, big.NewInt(testKeyIndex)}, transcript...)
	digest, err := ccommon.MultiHashNative(digestInputs...)
	c.Assert(err, qt.IsNil)
	c.Assert(digest.Cmp(publicInputs.TranscriptDigest), qt.Equals, 0)

	anchor, err := ccommon.ChallengeAnchor(transcript, digest)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(assignment.RoundHash, poolKeyTranscriptDomain, anchor)
	c.Assert(err, qt.IsNil)
	c.Assert(challenge.Cmp(publicInputs.Challenge), qt.Equals, 0)

	commitment, err := publicInputs.BRLCCommitment(publicInputs.Challenge)
	c.Assert(err, qt.IsNil)
	c.Assert(commitment.Cmp(publicInputs.TranscriptCommitment), qt.Equals, 0)
}

// The activation proof only works if it reproduces exactly the hash the
// contribution circuit committed to on chain, so pin the two together.
func TestContributionHashesMatchTheContributionCircuit(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	recipientKeys := make([]types.NodeKey, testCommittee)
	recipientIndexes := make([]uint16, testCommittee)
	nonces := make([]*big.Int, testCommittee)
	for i := range testCommittee {
		point := group.NewPoint()
		point.ScalarBaseMult(big.NewInt(int64(i*100 + 13)))
		encoded := group.Encode(point)
		recipientKeys[i] = types.NodeKey{PubX: encoded.X, PubY: encoded.Y}
		recipientIndexes[i] = uint16(i + 1)
		nonces[i] = big.NewInt(int64(1000 + i))
	}

	for i := range testCommittee {
		_, contributionInputs, err := contribution.BuildWitness(contribution.Assignment{
			RoundHash:        assignment.RoundHash,
			Threshold:        testThreshold,
			CommitteeSize:    testCommittee,
			ContributorIndex: assignment.ParticipantIndexes[i],
			Coefficients:     testCoefficients(assignment.ParticipantIndexes[i], testThreshold),
			RecipientIndexes: recipientIndexes,
			RecipientKeys:    recipientKeys,
			EncryptionNonces: nonces,
		})
		c.Assert(err, qt.IsNil)
		c.Assert(contributionInputs.CommitmentHash.Cmp(publicInputs.ContributionHashes[i]), qt.Equals, 0)
	}
}

// The share commitments the proof publishes are exactly what the contract
// hashes into the Merkle root a partial decryption later proves against:
// one leaf per committee member at slot p−1.
func TestShareCommitmentsBuildAMerkleRoot(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	members := make([]uint16, testCommittee)
	for i := range members {
		members[i] = uint16(i + 1)
	}
	leaves, err := ccommon.ShareCommitmentLeaves(members, publicInputs.ShareCommitments[:testCommittee])
	c.Assert(err, qt.IsNil)
	root := ccommon.MerkleRoot(leaves)
	c.Assert(root == [32]byte{}, qt.IsFalse)
	for i := testCommittee; i < MaxParticipants; i++ {
		c.Assert(leaves[i], qt.Equals, ccommon.EmptyLeaf)
	}

	// Climb the path the way submitPartialDecryption will.
	for _, member := range members {
		leafIndex := int(member) - 1
		path, err := ccommon.MerklePath(leaves, leafIndex)
		c.Assert(err, qt.IsNil)
		node, err := ccommon.ShareCommitmentLeaf(publicInputs.ShareCommitments[leafIndex])
		c.Assert(err, qt.IsNil)
		for depth, sibling := range path {
			if leafIndex>>depth&1 == 0 {
				node = ccommon.MerkleNode(node, sibling)
				continue
			}
			node = ccommon.MerkleNode(sibling, node)
		}
		c.Assert(node, qt.Equals, root)
	}
}

// A member that did not contribute still holds a share of every accepted
// polynomial, so its D_p is published (and provable) like everyone else's:
// committee of 4, member 3 silent.
func TestShareCommitmentForNonContributingMember(t *testing.T) {
	c := qt.New(t)

	accepted := []uint16{1, 2, 4}
	assignment := assignmentFor(testThreshold, 4, accepted)
	witness, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	c.Assert(publicInputs.AcceptedCount.Int64(), qt.Equals, int64(3))
	c.Assert(publicInputs.CommitteeSize.Int64(), qt.Equals, int64(4))

	for member := int64(1); member <= 4; member++ {
		assertPointEqual(c, publicInputs.ShareCommitments[member-1], expectedShareCommitment(accepted, member))
	}
	silent := publicInputs.ShareCommitments[2]
	c.Assert(silent.X.Sign() != 0 || silent.Y.Cmp(big.NewInt(1)) != 0, qt.IsTrue, qt.Commentf("D_3 is the identity"))
	assertPointEqual(c, publicInputs.ShareCommitments[4], ccommon.IdentityCurvePoint())

	// The transcript region [4N, 6N) is member-indexed, not row-indexed.
	transcript, err := publicInputs.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	c.Assert(transcript[4*MaxParticipants+2*2].Cmp(silent.X), qt.Equals, 0)
	c.Assert(transcript[4*MaxParticipants+2*2+1].Cmp(silent.Y), qt.Equals, 0)
	c.Assert(transcript[2].Int64(), qt.Equals, int64(4)) // third accepted row names member 4

	c.Assert(test.IsSolved(&PoolKeyCircuit{}, witness, ecc.BN254.ScalarField()), qt.IsNil)
}

func TestPoolKeyCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &PoolKeyCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

// A contributor's hash is the only thing binding its private commitments to
// the epoch, so a tampered one must not be provable.
func TestPoolKeyCircuitRejectsTamperedContributionHash(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.ContributionHashes[1] = big.NewInt(424242)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))
}

func TestPoolKeyCircuitRejectsTamperedAggregate(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.AggregateCommitments[0].X = big.NewInt(9999)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// D_p is a function of the aggregate; a published share commitment that is
// not that evaluation must not be provable, or a partial decryption could be
// verified against the wrong leaf.
func TestPoolKeyCircuitRejectsTamperedShareCommitment(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	// Swap in another member's (valid, on-curve) commitment.
	witness.ShareCommitments[1] = ccommon.CircuitPoint(publicInputs.ShareCommitments[0])

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// Swapping in another key's commitments must fail: the recomputed digest lands
// in the KeyIndex slot and no longer reproduces the stored hash.
func TestPoolKeyCircuitRejectsWrongKeyCommitments(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	for i := range assignment.Commitments {
		assignment.Commitments[i][testKeyIndex] = assignment.Commitments[i][testKeyIndex+1]
	}
	witness, _, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	// BuildWitness rehashed the tampered set, so restore the honest hashes the
	// contract would have stored.
	_, honest, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	for i := range MaxParticipants {
		witness.ContributionHashes[i] = honest.ContributionHashes[i]
	}

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// The one-hot key selector is what bounds KeyIndex: a value outside the pool
// selects no slot, so the selector sum is 0 and the proof fails.
func TestPoolKeyCircuitRejectsKeyIndexOutsidePool(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.KeyIndex = big.NewInt(MaxKeys)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// The digest pins the witness transcript before ρ exists: a digest computed
// over any other vector — here the honest transcript with one word changed,
// as a prover who re-derived ρ after editing calldata would present — must
// not be provable against the witness words.
func TestPoolKeyCircuitRejectsTranscriptWordChangedAfterDigest(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	transcript, err := publicInputs.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	altered := append([]*big.Int{}, transcript...)
	altered[2*MaxParticipants] = new(big.Int).Add(altered[2*MaxParticipants], big.NewInt(1)) // Ā[0].x
	digest, err := TranscriptDigestNative(publicInputs.RoundHash, publicInputs.KeyIndex, altered)
	c.Assert(err, qt.IsNil)
	c.Assert(digest.Cmp(publicInputs.TranscriptDigest), qt.Not(qt.Equals), 0)
	witness.TranscriptDigest = digest

	assert := test.NewAssert(t)
	assert.SolvingFailed(&PoolKeyCircuit{}, witness, test.WithCurves(ecc.BN254))

	// Same transcript under another key index is a different digest too.
	otherKey, err := TranscriptDigestNative(publicInputs.RoundHash, big.NewInt(testKeyIndex+1), transcript)
	c.Assert(err, qt.IsNil)
	c.Assert(otherKey.Cmp(publicInputs.TranscriptDigest), qt.Not(qt.Equals), 0)
}

func TestPoolKeyArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}

func TestValidateRejectsBadAssignment(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	assignment.KeyIndex = MaxKeys
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))

	assignment = testAssignment()
	assignment.ParticipantIndexes[1] = assignment.ParticipantIndexes[0]
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))

	assignment = testAssignment()
	assignment.Commitments[0] = assignment.Commitments[0][:MaxKeys-1]
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))
}
