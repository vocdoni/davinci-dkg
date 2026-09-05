package finalize

import (
	"math/big"
	"testing"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

const (
	testThreshold = 2
	testCommittee = 3
)

// testCoefficients builds one dealer's MaxK polynomials deterministically,
// distinct per dealer and per pool key.
func testCoefficients(dealerIndex uint16, threshold int) [][]*big.Int {
	sets := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		sets[j] = make([]*big.Int, threshold)
		for m := range threshold {
			sets[j][m] = big.NewInt(int64(int(dealerIndex)*1000 + (j+1)*10 + m + 1))
		}
	}
	return sets
}

// testCommitments turns one dealer's polynomials into the Feldman commitment
// points the finalizer reads from the contribution calldata.
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

// assignmentFor builds a finalization over the given accepted dealers of a
// committee of `committee`, every dealer dealing testCoefficients.
func assignmentFor(threshold, committee uint16, accepted []uint16) Assignment {
	commitments := make([][][]types.CurvePoint, len(accepted))
	for i, index := range accepted {
		commitments[i] = testCommitments(testCoefficients(index, int(threshold)))
	}
	return Assignment{
		RoundHash:          big.NewInt(2222),
		Threshold:          threshold,
		CommitteeSize:      committee,
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

// expectedShareCommitment evaluates Σ_m p^m · Ā[j][m] natively from the
// accepted dealers' coefficients, independently of the witness builder.
func expectedShareCommitment(accepted []uint16, threshold int, key int, member int64) types.CurvePoint {
	sum := big.NewInt(0)
	for _, index := range accepted {
		coefficients := testCoefficients(index, threshold)[key]
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

func expectedPoolKey(accepted []uint16, threshold int, key int) types.CurvePoint {
	sum := group.NewPoint()
	sum.SetZero()
	for _, index := range accepted {
		point := group.NewPoint()
		point.ScalarBaseMult(testCoefficients(index, threshold)[key][0])
		sum.Add(sum, point)
	}
	return group.Encode(sum)
}

func assertPointEqual(c *qt.C, got, want types.CurvePoint) {
	c.Helper()
	c.Assert(got.X.Cmp(want.X), qt.Equals, 0)
	c.Assert(got.Y.Cmp(want.Y), qt.Equals, 0)
}

func isIdentity(p types.CurvePoint) bool {
	return p.X.Sign() == 0 && p.Y.Cmp(big.NewInt(1)) == 0
}

// Every key's P_j is the sum of the accepted constant-term commitments and
// every D_j,i the aggregate polynomial at i+1, identity beyond the committee.
func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	witness, pi, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(pi.PoolKeys, qt.HasLen, MaxKeys)
	c.Assert(pi.ShareCommitments, qt.HasLen, MaxKeys)
	for j := range MaxKeys {
		assertPointEqual(c, pi.PoolKeys[j], expectedPoolKey(assignment.ParticipantIndexes, testThreshold, j))
		assertPointEqual(c, pi.AggregateCommitments[j][0], pi.PoolKeys[j])
		c.Assert(pi.ShareCommitments[j], qt.HasLen, MaxParticipants)
		for i := range testCommittee {
			assertPointEqual(c, pi.ShareCommitments[j][i],
				expectedShareCommitment(assignment.ParticipantIndexes, testThreshold, j, int64(i+1)))
		}
		for i := testCommittee; i < MaxParticipants; i++ {
			c.Assert(isIdentity(pi.ShareCommitments[j][i]), qt.IsTrue)
		}
		for m := testThreshold; m < MaxCoefficients; m++ {
			c.Assert(isIdentity(pi.AggregateCommitments[j][m]), qt.IsTrue)
		}
	}
	// Keys differ from each other: independent polynomials per key.
	c.Assert(pi.PoolKeys[0].X.Cmp(pi.PoolKeys[MaxKeys-1].X), qt.Not(qt.Equals), 0)
}

// The transcript is L_F words in the §7 order; the digest is the three-level
// Poseidon over exactly those words; the anchor folds the digest in before
// the calldata keccak; the public inputs are the 7 contract words.
func TestTranscriptLayoutAndDigest(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, pi, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	c.Assert(TranscriptWords, qt.Equals, 2*MaxParticipants+MaxKeys*(2+2*MaxParticipants))
	c.Assert(TranscriptWords, qt.Equals, 1120)
	c.Assert(KeyWords, qt.Equals, 2+2*MaxParticipants)
	c.Assert(KeyOffset(MaxKeys), qt.Equals, TranscriptWords)

	transcript, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	c.Assert(transcript, qt.HasLen, TranscriptWords)
	for d := range MaxParticipants {
		wantIndex, wantHash := int64(0), big.NewInt(0)
		if d < testCommittee {
			wantIndex = int64(assignment.ParticipantIndexes[d])
			wantHash = pi.ContributionHashes[d]
		}
		c.Assert(transcript[IndexesStart+d].Int64(), qt.Equals, wantIndex)
		c.Assert(transcript[HashesStart+d].Cmp(wantHash), qt.Equals, 0)
	}
	for j := range MaxKeys {
		c.Assert(transcript[PoolKeyOffset(j)].Cmp(pi.PoolKeys[j].X), qt.Equals, 0)
		c.Assert(transcript[PoolKeyOffset(j)+1].Cmp(pi.PoolKeys[j].Y), qt.Equals, 0)
		for i := range MaxParticipants {
			c.Assert(transcript[ShareCommitmentOffset(j, i)].Cmp(pi.ShareCommitments[j][i].X), qt.Equals, 0)
			c.Assert(transcript[ShareCommitmentOffset(j, i)+1].Cmp(pi.ShareCommitments[j][i].Y), qt.Equals, 0)
		}
	}

	// Digest: R over the rows, B_j per key block, T over everything.
	parts, err := TranscriptDigestParts(pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount, transcript)
	c.Assert(err, qt.IsNil)
	rowsInputs := append([]*big.Int{big.NewInt(0)}, transcript[:KeysStart]...)
	rows, err := ccommon.MultiHashNative(rowsInputs...)
	c.Assert(err, qt.IsNil)
	c.Assert(parts.Rows.Cmp(rows), qt.Equals, 0)
	c.Assert(len(rowsInputs), qt.Equals, 65)
	outer := []*big.Int{
		big.NewInt(2), pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount,
		big.NewInt(MaxKeys), big.NewInt(TranscriptWords), rows,
	}
	for j := range MaxKeys {
		keyInputs := append([]*big.Int{big.NewInt(1), big.NewInt(int64(j))}, transcript[KeyOffset(j):KeyOffset(j+1)]...)
		c.Assert(len(keyInputs), qt.Equals, 68)
		block, err := ccommon.MultiHashNative(keyInputs...)
		c.Assert(err, qt.IsNil)
		c.Assert(parts.Keys[j].Cmp(block), qt.Equals, 0)
		outer = append(outer, block)
	}
	c.Assert(len(outer), qt.Equals, 24)
	digest, err := ccommon.MultiHashNative(outer...)
	c.Assert(err, qt.IsNil)
	c.Assert(parts.Digest.Cmp(digest), qt.Equals, 0)
	c.Assert(pi.TranscriptDigest.Cmp(digest), qt.Equals, 0)

	anchor, err := ccommon.ChallengeAnchor(transcript, digest)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(assignment.RoundHash, TranscriptDomain, anchor)
	c.Assert(err, qt.IsNil)
	c.Assert(challenge.Cmp(pi.Challenge), qt.Equals, 0)
	commitment, err := pi.BRLCCommitment(pi.Challenge)
	c.Assert(err, qt.IsNil)
	c.Assert(commitment.Cmp(pi.TranscriptCommitment), qt.Equals, 0)

	scalars := pi.Scalars()
	c.Assert(scalars, qt.HasLen, PublicInputWords)
	c.Assert(scalars[0].Cmp(assignment.RoundHash), qt.Equals, 0)
	c.Assert(scalars[1].Int64(), qt.Equals, int64(testThreshold))
	c.Assert(scalars[2].Int64(), qt.Equals, int64(testCommittee))
	c.Assert(scalars[3].Int64(), qt.Equals, int64(testCommittee))
	c.Assert(scalars[4].Cmp(pi.TranscriptDigest), qt.Equals, 0)
	c.Assert(scalars[5].Cmp(pi.Challenge), qt.Equals, 0)
	c.Assert(scalars[6].Cmp(pi.TranscriptCommitment), qt.Equals, 0)

	// Any other word vector or any other count is a different digest.
	altered := append([]*big.Int{}, transcript...)
	altered[PoolKeyOffset(5)] = new(big.Int).Add(altered[PoolKeyOffset(5)], big.NewInt(1))
	other, err := TranscriptDigestNative(pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount, altered)
	c.Assert(err, qt.IsNil)
	c.Assert(other.Cmp(digest), qt.Not(qt.Equals), 0)
	other, err = TranscriptDigestNative(pi.RoundHash, pi.Threshold, pi.CommitteeSize, big.NewInt(2), transcript)
	c.Assert(err, qt.IsNil)
	c.Assert(other.Cmp(digest), qt.Not(qt.Equals), 0)
	_, err = TranscriptDigestNative(pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount, transcript[1:])
	c.Assert(err, qt.Not(qt.IsNil))
}

// The finalization proof only works if it reproduces exactly the hash the
// contribution circuit committed to on chain, so pin the two together; a
// stored hash that disagrees is refused by the builder.
func TestDigestMatchesTheContributionCircuit(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, pi, err := BuildWitness(assignment)
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

	stored := make([]*big.Int, testCommittee)
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
		c.Assert(contributionInputs.CommitmentHash.Cmp(pi.ContributionHashes[i]), qt.Equals, 0)
		stored[i] = contributionInputs.CommitmentHash

		// The finalizer reads the commitments back from the compact calldata.
		layout, err := contributionInputs.Layout()
		c.Assert(err, qt.IsNil)
		words, err := contributionInputs.TranscriptScalars()
		c.Assert(err, qt.IsNil)
		decoded, err := layout.Decode(words)
		c.Assert(err, qt.IsNil)
		for j := range MaxKeys {
			for m := range testThreshold {
				assertPointEqual(c, decoded.Commitments[j][m], assignment.Commitments[i][j][m])
			}
		}
	}

	assignment.ContributionHashes = stored
	_, _, err = BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	assignment.ContributionHashes[1] = big.NewInt(424242)
	_, _, err = BuildWitness(assignment)
	c.Assert(err, qt.ErrorMatches, "dealer 1 .* does not match the stored .*")
}

// The share commitments the proof publishes are exactly what the contract
// hashes into the Merkle root of every key: one leaf per committee member at
// slot p−1, EmptyLeaf beyond, and a path a partial decryption can climb.
func TestTranscriptShareRootsPerKey(t *testing.T) {
	c := qt.New(t)

	_, pi, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	roots, err := pi.ShareRoots()
	c.Assert(err, qt.IsNil)

	seen := make(map[[32]byte]struct{}, MaxKeys)
	for j := range MaxKeys {
		leaves, err := pi.ShareLeaves(j)
		c.Assert(err, qt.IsNil)
		c.Assert(ccommon.MerkleRoot(leaves), qt.Equals, roots[j])
		for i := testCommittee; i < MaxParticipants; i++ {
			c.Assert(leaves[i], qt.Equals, ccommon.EmptyLeaf)
		}
		_, duplicate := seen[roots[j]]
		c.Assert(duplicate, qt.IsFalse, qt.Commentf("key %d shares a root", j))
		seen[roots[j]] = struct{}{}

		for member := 1; member <= testCommittee; member++ {
			leafIndex := member - 1
			path, err := ccommon.MerklePath(leaves, leafIndex)
			c.Assert(err, qt.IsNil)
			node, err := ccommon.ShareCommitmentLeaf(pi.ShareCommitments[j][leafIndex])
			c.Assert(err, qt.IsNil)
			for depth, sibling := range path {
				if leafIndex>>depth&1 == 0 {
					node = ccommon.MerkleNode(node, sibling)
					continue
				}
				node = ccommon.MerkleNode(sibling, node)
			}
			c.Assert(node, qt.Equals, roots[j])
		}
	}
	_, err = pi.ShareLeaves(MaxKeys)
	c.Assert(err, qt.Not(qt.IsNil))
}

// A member that did not contribute still holds a share of every accepted
// polynomial, so its D_j,p is published like everyone else's; accepted rows
// may come in any order. Committee of 4, member 3 silent, rows descending.
func TestTranscriptNonContiguousAcceptedRows(t *testing.T) {
	c := qt.New(t)

	accepted := []uint16{4, 1, 2}
	assignment := assignmentFor(testThreshold, 4, accepted)
	_, pi, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)
	c.Assert(pi.AcceptedCount.Int64(), qt.Equals, int64(3))
	c.Assert(pi.CommitteeSize.Int64(), qt.Equals, int64(4))

	transcript, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	c.Assert(transcript[IndexesStart].Int64(), qt.Equals, int64(4))
	c.Assert(transcript[IndexesStart+2].Int64(), qt.Equals, int64(2))
	c.Assert(transcript[IndexesStart+3].Int64(), qt.Equals, int64(0))
	c.Assert(transcript[HashesStart+3].Sign(), qt.Equals, 0)

	for j := range MaxKeys {
		for member := int64(1); member <= 4; member++ {
			assertPointEqual(c, pi.ShareCommitments[j][member-1], expectedShareCommitment(accepted, testThreshold, j, member))
		}
		silent := pi.ShareCommitments[j][2]
		c.Assert(isIdentity(silent), qt.IsFalse, qt.Commentf("key %d: D_3 is the identity", j))
		c.Assert(transcript[ShareCommitmentOffset(j, 2)].Cmp(silent.X), qt.Equals, 0)
		c.Assert(isIdentity(pi.ShareCommitments[j][4]), qt.IsTrue)
	}
}

// The circuit accepts the honest witness (tiny t, n through the test engine)
// and refuses a tampered hash, aggregate, share commitment or digest.
func TestFoldFinalizeCircuitSolves(t *testing.T) {
	c := qt.New(t)
	field := ecc.BN254.ScalarField()

	witness, pi, err := BuildWitness(assignmentFor(testThreshold, 4, []uint16{1, 2, 4}))
	c.Assert(err, qt.IsNil)

	started := time.Now()
	c.Assert(test.IsSolved(&FinalizeCircuit{}, witness, field), qt.IsNil)
	t.Logf("test-engine solve of the finalize circuit: %s", time.Since(started))

	// A dealer's hash is the only thing binding its commitments to the epoch.
	wrong := *witness
	wrong.ContributionHashes[1] = big.NewInt(424242)
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// The aggregate of key 7 is pinned to the sum of the accepted rows.
	wrong = *witness
	wrong.AggregateCommitments[7][0] = ccommon.CircuitPoint(pi.PoolKeys[8])
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// D_j,p is a function of the aggregate: another member's (valid, on-curve)
	// commitment must not be provable at its slot.
	wrong = *witness
	wrong.ShareCommitments[15][1] = ccommon.CircuitPoint(pi.ShareCommitments[15][0])
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// An active row must name a member in [1, n].
	wrong = *witness
	wrong.ParticipantIndexes[0] = big.NewInt(5)
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))
	wrong = *witness
	wrong.ParticipantIndexes[0] = big.NewInt(0)
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// The digest pins the witness transcript before ρ exists: a digest over
	// the honest transcript with one word changed must not be provable.
	transcript, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	altered := append([]*big.Int{}, transcript...)
	altered[PoolKeyOffset(0)] = new(big.Int).Add(altered[PoolKeyOffset(0)], big.NewInt(1))
	digest, err := TranscriptDigestNative(pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount, altered)
	c.Assert(err, qt.IsNil)
	wrong = *witness
	wrong.TranscriptDigest = digest
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// Count bounds: a < t and t = 0 are rejected.
	wrong = *witness
	wrong.AcceptedCount = big.NewInt(1)
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))
	wrong = *witness
	wrong.Threshold = big.NewInt(0)
	c.Assert(test.IsSolved(&FinalizeCircuit{}, &wrong, field), qt.Not(qt.IsNil))
}

func TestValidateRejectsBadAssignment(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	assignment.ParticipantIndexes[1] = assignment.ParticipantIndexes[0]
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))

	assignment = testAssignment()
	assignment.Commitments[0] = assignment.Commitments[0][:MaxKeys-1]
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))

	assignment = testAssignment()
	assignment.ContributionHashes = []*big.Int{big.NewInt(1)}
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))

	assignment = testAssignment()
	assignment.ParticipantIndexes = assignment.ParticipantIndexes[:1]
	assignment.Commitments = assignment.Commitments[:1]
	c.Assert(assignment.Validate(), qt.Not(qt.IsNil))
}

// Compile-only: logs the R1CS size of the finalization circuit. Skipped
// under -short; compiling the full-size circuit takes minutes.
func TestCompileFinalizeConstraintCount(t *testing.T) {
	if testing.Short() {
		t.Skip("compile-only constraint count skipped under -short")
	}
	started := time.Now()
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &FinalizeCircuit{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("finalize circuit (MaxN=%d, MaxK=%d): %d constraints, %d public inputs, compiled in %s",
		MaxParticipants, MaxKeys, ccs.GetNbConstraints(), ccs.GetNbPublicVariables()-1, time.Since(started))
}

// Release gate, not a unit test: passes only once `make circuits` re-pinned
// the v4 artifacts in config/circuit_artifacts.go.
func TestFinalizeArtifactsMatchCompiledCircuit(t *testing.T) {
	if testing.Short() {
		t.Skip("artifact pin check skipped under -short")
	}
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
