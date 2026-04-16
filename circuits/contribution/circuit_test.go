package contribution

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/types"
)

func testAssignment() Assignment {
	key1 := group.NewPoint()
	key1.ScalarBaseMult(big.NewInt(13))
	key2 := group.NewPoint()
	key2.ScalarBaseMult(big.NewInt(17))
	key3 := group.NewPoint()
	key3.ScalarBaseMult(big.NewInt(19))
	key4 := group.NewPoint()
	key4.ScalarBaseMult(big.NewInt(23))

	return Assignment{
		RoundHash:        big.NewInt(12345),
		Threshold:        3,
		CommitteeSize:    4,
		ContributorIndex: 1,
		Coefficients: []*big.Int{
			big.NewInt(11),
			big.NewInt(7),
			big.NewInt(3),
		},
		RecipientIndexes: []uint16{1, 2, 3, 4},
		RecipientKeys: []types.NodeKey{
			{PubX: group.Encode(key1).X, PubY: group.Encode(key1).Y},
			{PubX: group.Encode(key2).X, PubY: group.Encode(key2).Y},
			{PubX: group.Encode(key3).X, PubY: group.Encode(key3).Y},
			{PubX: group.Encode(key4).X, PubY: group.Encode(key4).Y},
		},
		EncryptionNonces: []*big.Int{
			big.NewInt(31),
			big.NewInt(37),
			big.NewInt(41),
			big.NewInt(43),
		},
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.CommitmentHash, qt.Not(qt.IsNil))
	c.Assert(publicInputs.ShareHash, qt.Not(qt.IsNil))
}

func TestBuildWitnessIncludesCommitmentsAndEncryptedShares(t *testing.T) {
	c := qt.New(t)

	assignment := testAssignment()
	_, publicInputs, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	c.Assert(len(publicInputs.Commitments), qt.Equals, len(assignment.Coefficients))
	c.Assert(len(publicInputs.EncryptedShares), qt.Equals, len(assignment.RecipientIndexes))

	expectedCommitment := group.NewPoint()
	expectedCommitment.ScalarBaseMult(assignment.Coefficients[0])
	expectedCommitmentEncoded := group.Encode(expectedCommitment)
	c.Assert(publicInputs.Commitments[0].X.Cmp(expectedCommitmentEncoded.X), qt.Equals, 0)
	c.Assert(publicInputs.Commitments[0].Y.Cmp(expectedCommitmentEncoded.Y), qt.Equals, 0)

	expectedCiphertext, err := shareenc.EncryptShareWithNonceRoundHash(
		assignment.RoundHash,
		assignment.ContributorIndex,
		assignment.RecipientIndexes[0],
		publicInputs.Shares[0],
		assignment.RecipientKeys[0],
		assignment.EncryptionNonces[0],
	)
	c.Assert(err, qt.IsNil)
	c.Assert(publicInputs.EncryptedShares[0].Ephemeral.X.Cmp(expectedCiphertext.Ephemeral.X), qt.Equals, 0)
	c.Assert(publicInputs.EncryptedShares[0].Ephemeral.Y.Cmp(expectedCiphertext.Ephemeral.Y), qt.Equals, 0)
	c.Assert(publicInputs.EncryptedShares[0].Ciphertext.Cmp(expectedCiphertext.MaskedShare), qt.Equals, 0)
}

func TestContributionCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestContributionCircuitRejectsTamperedShare(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.Shares[0] = big.NewInt(999999)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestContributionArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}

// TestContributionCircuitSolvingLargeCoefficients verifies that the contribution
// circuit is satisfiable when coefficients are near the BabyJubJub subgroup order
// (r_bjj). Previously, EvaluatePolynomialNative used the BN254 scalar field as
// the modulus; polynomial evaluations at x ≥ 2 could produce shares ≥ r_bjj,
// breaking the AddModSubgroupOrder constraint (carry must be in {0,1}).
//
// Uses n=4, t=3 so the test is fast even at MaxN=32.
func TestContributionCircuitSolvingLargeCoefficients(t *testing.T) {
	const n, threshold = 4, 3
	c := qt.New(t)

	// Coefficients near r_bjj — these would have caused share overflow under the old code.
	rbjj := group.ScalarField()
	bigCoeff := new(big.Int).Sub(rbjj, big.NewInt(1)) // r_bjj - 1

	recipientKeys := make([]types.NodeKey, n)
	recipientIndexes := make([]uint16, n)
	for j := range n {
		p := group.NewPoint()
		p.ScalarBaseMult(big.NewInt(int64(j*100 + 13)))
		enc := group.Encode(p)
		recipientKeys[j] = types.NodeKey{PubX: enc.X, PubY: enc.Y}
		recipientIndexes[j] = uint16(j + 1)
	}
	nonces := []*big.Int{big.NewInt(1001), big.NewInt(1002), big.NewInt(1003), big.NewInt(1004)}

	for i := range n {
		t.Run("contributor_"+big.NewInt(int64(i+1)).String(), func(t *testing.T) {
			coefficients := make([]*big.Int, threshold)
			for k := range threshold {
				// Alternate between near-max and mid-range values.
				if k%2 == 0 {
					coefficients[k] = new(big.Int).Set(bigCoeff)
				} else {
					coefficients[k] = new(big.Int).Rsh(rbjj, 1) // r_bjj / 2
				}
			}

			assignment := Assignment{
				RoundHash:        big.NewInt(99999),
				Threshold:        uint16(threshold),
				CommitteeSize:    uint16(n),
				ContributorIndex: uint16(i + 1),
				Coefficients:     coefficients,
				RecipientIndexes: recipientIndexes,
				RecipientKeys:    recipientKeys,
				EncryptionNonces: nonces,
			}

			witness, _, err := BuildWitness(assignment)
			c.Assert(err, qt.IsNil)

			assert := test.NewAssert(t)
			assert.SolvingSucceeded(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
		})
	}
}

// TestContributionCircuitSolvingN12T8Contributor9 checks that the contribution
// circuit constraints are satisfiable for n=12, t=8, contributor 9 using the
// same small coefficients that caused "pairing doesn't match" in the gas
// benchmark (coefficients[8][k] = (8+1)*10 + k + 1 = 91..98).
//
// Run with MaxN=32:
//
//	go test -v -run TestContributionCircuitSolvingN12T8Contributor9 ./circuits/contribution/...
func TestContributionCircuitSolvingN12T8Contributor9(t *testing.T) {
	const n, threshold, contributorI = 12, 8, 9 // i=8 zero-based → index 9

	if MaxRecipients < n {
		t.Skipf("MaxN=%d < n=%d; recompile with MaxN=32 to run this test", MaxRecipients, n)
	}

	c := qt.New(t)

	recipientKeys := make([]types.NodeKey, n)
	recipientIndexes := make([]uint16, n)
	for j := range n {
		p := group.NewPoint()
		p.ScalarBaseMult(big.NewInt(int64(j*100 + 13)))
		enc := group.Encode(p)
		recipientKeys[j] = types.NodeKey{PubX: enc.X, PubY: enc.Y}
		recipientIndexes[j] = uint16(j + 1)
	}
	nonces := make([]*big.Int, n)
	for j := range n {
		nonces[j] = big.NewInt(int64(1000 + j + 1))
	}

	// Exact coefficients from the gas benchmark that triggered the failure.
	coefficients := make([]*big.Int, threshold)
	for k := range threshold {
		coefficients[k] = big.NewInt(int64(contributorI*10 + k + 1))
	}

	assignment := Assignment{
		RoundHash:        big.NewInt(12345),
		Threshold:        uint16(threshold),
		CommitteeSize:    uint16(n),
		ContributorIndex: uint16(contributorI),
		Coefficients:     coefficients,
		RecipientIndexes: recipientIndexes,
		RecipientKeys:    recipientKeys,
		EncryptionNonces: nonces,
	}

	witness, _, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitProveAndVerifyN12AllContributors runs a full Groth16
// prove+verify cycle for all 12 contributors at n=12, t=8 — including
// contributor 9 that failed with "pairing doesn't match" in the gas benchmark.
// Uses the same coefficient pattern as the benchmark: coefficients[i][k] = (i+1)*10 + k + 1.
//
// Run with MaxN=32:
//
//	go test -v -run TestContributionCircuitProveAndVerifyN12AllContributors ./circuits/contribution/...
func TestContributionCircuitProveAndVerifyN12AllContributors(t *testing.T) {
	const n, threshold = 12, 8

	if MaxRecipients < n {
		t.Skipf("MaxN=%d < n=%d; recompile with MaxN=32 to run this test", MaxRecipients, n)
	}

	c := qt.New(t)

	recipientKeys := make([]types.NodeKey, n)
	recipientIndexes := make([]uint16, n)
	for j := range n {
		p := group.NewPoint()
		p.ScalarBaseMult(big.NewInt(int64(j*100 + 13)))
		enc := group.Encode(p)
		recipientKeys[j] = types.NodeKey{PubX: enc.X, PubY: enc.Y}
		recipientIndexes[j] = uint16(j + 1)
	}
	nonces := make([]*big.Int, n)
	for j := range n {
		nonces[j] = big.NewInt(int64(1000 + j + 1))
	}

	// Realistic-looking roundHash (simulate a 12-byte blockchain round ID).
	roundHash := new(big.Int).SetBytes([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00, 0x00, 0x03})

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	c.Assert(err, qt.IsNil)

	for i := range n {
		t.Run(fmt.Sprintf("contributor_%d", i+1), func(t *testing.T) {
			c := qt.New(t)
			coefficients := make([]*big.Int, threshold)
			for k := range threshold {
				coefficients[k] = big.NewInt(int64((i+1)*10 + k + 1))
			}
			assignment := Assignment{
				RoundHash:        roundHash,
				Threshold:        uint16(threshold),
				CommitteeSize:    uint16(n),
				ContributorIndex: uint16(i + 1),
				Coefficients:     coefficients,
				RecipientIndexes: recipientIndexes,
				RecipientKeys:    recipientKeys,
				EncryptionNonces: nonces,
			}
			witness, publicInputs, err := BuildWitness(assignment)
			c.Assert(err, qt.IsNil)

			proof, err := runtime.ProveAndVerify(witness)
			c.Assert(err, qt.IsNil)
			c.Assert(proof, qt.Not(qt.IsNil))

			err = runtime.Verify(proof, publicInputs.PublicWitness())
			c.Assert(err, qt.IsNil)
		})
	}
}

// TestContributionCircuitSolvingN20T14 checks that the contribution circuit
// constraints are satisfiable for n=20, t=14, contributor i=6 using gnark's
// fast constraint solver (no proving key needed). This mirrors the exact case
// that failed with "points in the proof are not in the correct subgroup"
// during the MaxN=32 benchmark run.
//
// Run with MaxN=32 (circuits/common/sizes.go: const MaxN = 32):
//
//	go test -run TestContributionCircuitSolvingN20T14 ./circuits/contribution/...
func TestContributionCircuitSolvingN20T14(t *testing.T) {
	const n, threshold = 20, 14

	if MaxRecipients < n {
		t.Skipf("MaxN=%d < n=%d; recompile with MaxN=32 to run this test", MaxRecipients, n)
	}

	c := qt.New(t)

	// Build recipient keys deterministically (key_j = (j*100+13)*G).
	recipientKeys := make([]types.NodeKey, n)
	recipientIndexes := make([]uint16, n)
	for j := range n {
		p := group.NewPoint()
		p.ScalarBaseMult(big.NewInt(int64(j*100 + 13)))
		enc := group.Encode(p)
		recipientKeys[j] = types.NodeKey{PubX: enc.X, PubY: enc.Y}
		recipientIndexes[j] = uint16(j + 1)
	}

	nonces := make([]*big.Int, n)
	for j := range n {
		nonces[j] = big.NewInt(int64(1000 + j + 1))
	}

	// Test every contributor, starting from i=6 (the one that failed).
	for startI := 6; startI < n; startI++ {
		i := startI
		t.Run("contributor_"+big.NewInt(int64(i+1)).String(), func(t *testing.T) {
			coefficients := make([]*big.Int, threshold)
			for k := range threshold {
				coefficients[k] = big.NewInt(int64((i+1)*10 + k + 1))
			}

			assignment := Assignment{
				RoundHash:        big.NewInt(12345),
				Threshold:        uint16(threshold),
				CommitteeSize:    uint16(n),
				ContributorIndex: uint16(i + 1),
				Coefficients:     coefficients,
				RecipientIndexes: recipientIndexes,
				RecipientKeys:    recipientKeys,
				EncryptionNonces: nonces,
			}

			witness, _, err := BuildWitness(assignment)
			c.Assert(err, qt.IsNil)

			assert := test.NewAssert(t)
			assert.SolvingSucceeded(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
		})
	}
}
