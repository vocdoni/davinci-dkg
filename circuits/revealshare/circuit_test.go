package revealshare

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func testAssignment() Assignment {
	return Assignment{
		RoundHash:          big.NewInt(6666),
		Threshold:          3,
		ParticipantIndexes: []uint16{1, 2, 3},
		RevealedShares:     []*big.Int{big.NewInt(6), big.NewInt(15), big.NewInt(28)},
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.DisclosureHash, qt.Not(qt.IsNil))
}

func TestRevealShareCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &RevealShareCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestRevealShareCircuitRejectsTamperedSecret(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.ReconstructedSecret = big.NewInt(999999)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &RevealShareCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestRevealShareArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
