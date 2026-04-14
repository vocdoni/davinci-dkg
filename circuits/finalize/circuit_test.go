package finalize

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func testAssignment() Assignment {
	return Assignment{
		RoundHash:          big.NewInt(2222),
		Threshold:          3,
		CommitteeSize:      3,
		ParticipantIndexes: []uint16{1, 2, 3},
		ContributionCoefficients: [][]*big.Int{
			{big.NewInt(10), big.NewInt(3), big.NewInt(1)},
			{big.NewInt(7), big.NewInt(2), big.NewInt(4)},
			{big.NewInt(9), big.NewInt(5), big.NewInt(6)},
		},
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.AggregateHash, qt.Not(qt.IsNil))
	c.Assert(publicInputs.ShareCommitmentHash, qt.Not(qt.IsNil))
}

func TestFinalizeCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &FinalizeCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestFinalizeCircuitRejectsTamperedAggregate(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.AggregateCommitments[0].X = big.NewInt(9999)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &FinalizeCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestFinalizeArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
