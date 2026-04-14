package revealsubmit

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

func testAssignment() Assignment {
	share := big.NewInt(13)
	point := group.Generator()
	point.ScalarBaseMult(share)
	return Assignment{
		RoundHash:        big.NewInt(7777),
		ParticipantIndex: 1,
		ShareValue:       share,
		ShareCommitment:  group.Encode(point),
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
}

func TestRevealSubmitCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &RevealSubmitCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestRevealSubmitRejectsTamperedShare(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.ShareValue = big.NewInt(99)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &RevealSubmitCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestRevealSubmitArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
