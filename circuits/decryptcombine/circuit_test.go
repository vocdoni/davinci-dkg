package decryptcombine

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

func testAssignment() Assignment {
	c1Point := group.Generator()
	c1Point.ScalarBaseMult(big.NewInt(5))
	delta0Point := group.Generator()
	delta0Point.ScalarBaseMult(big.NewInt(14))
	messagePoint := group.Generator()
	messagePoint.ScalarBaseMult(big.NewInt(3))
	c2Point := group.NewPoint()
	c2Point.Set(messagePoint)
	c2Point.Add(c2Point, delta0Point)

	return Assignment{
		RoundHash:          big.NewInt(5555),
		Threshold:          1,
		CiphertextC1:       group.Encode(c1Point),
		CiphertextC2:       group.Encode(c2Point),
		ParticipantIndexes: []uint16{1},
		PartialDecryptions: []types.CurvePoint{group.Encode(delta0Point)},
		Plaintext:          big.NewInt(3),
	}
}

func TestBuildWitness(t *testing.T) {
	c := qt.New(t)

	witness, publicInputs, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)
	c.Assert(witness, qt.Not(qt.IsNil))
	c.Assert(publicInputs, qt.Not(qt.IsNil))
	c.Assert(publicInputs.PlaintextHash, qt.Not(qt.IsNil))
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
