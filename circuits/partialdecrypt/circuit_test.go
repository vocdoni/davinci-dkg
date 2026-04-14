package partialdecrypt

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

func testAssignment() Assignment {
	base := group.Generator()
	base.ScalarBaseMult(big.NewInt(5))
	return Assignment{
		RoundHash:        big.NewInt(4444),
		ParticipantIndex: 2,
		Base:             group.Encode(base),
		Secret:           big.NewInt(7),
		Nonce:            big.NewInt(11),
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.Response, qt.Not(qt.IsNil))
}

func TestPartialDecryptCircuitProveAndVerify(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &PartialDecryptCircuit{})
	c.Assert(err, qt.IsNil)

	proof, err := runtime.ProveAndVerify(witness)
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))

	err = runtime.Verify(proof, publicInputs.PublicWitness())
	c.Assert(err, qt.IsNil)
}

func TestPartialDecryptCircuitRejectsTamperedDelta(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	witness.Delta = ccommon.CircuitPoint(types.CurvePoint{
		X: group.Encode(group.Generator()).X,
		Y: group.Encode(group.Generator()).Y,
	})

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &PartialDecryptCircuit{})
	c.Assert(err, qt.IsNil)

	_, err = runtime.ProveAndVerify(witness)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestPartialDecryptArtifactsMatchCompiledCircuit(t *testing.T) {
	c := qt.New(t)

	ccs, err := Compile()
	c.Assert(err, qt.IsNil)

	matches, err := Artifacts.Matches(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(matches, qt.IsTrue)
}
