package contribution

import (
	"context"
	"math/big"
	"testing"

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
