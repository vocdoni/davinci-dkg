package prover

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	qt "github.com/frankban/quicktest"
)

type proverTestCircuit struct {
	X frontend.Variable `gnark:",public"`
	Y frontend.Variable
}

func (c *proverTestCircuit) Define(api frontend.API) error {
	api.AssertIsEqual(c.X, c.Y)
	return nil
}

func TestSetupAndProve(t *testing.T) {
	c := qt.New(t)
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &proverTestCircuit{})
	c.Assert(err, qt.IsNil)

	pk, vk, err := Setup(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(pk, qt.Not(qt.IsNil))
	c.Assert(vk, qt.Not(qt.IsNil))

	proof, err := Prove(ecc.BN254, ccs, pk, &proverTestCircuit{X: 7, Y: 7})
	c.Assert(err, qt.IsNil)
	c.Assert(proof, qt.Not(qt.IsNil))
}

func TestSetProver(t *testing.T) {
	c := qt.New(t)
	original := prover
	defer func() { prover = original }()

	called := false
	SetProver(func(
		curve ecc.ID,
		ccs constraint.ConstraintSystem,
		pk groth16.ProvingKey,
		assignment frontend.Circuit,
		opts ...backend.ProverOption,
	) (groth16.Proof, error) {
		called = true
		return nil, nil
	})

	_, err := Prove(ecc.BN254, nil, nil, &proverTestCircuit{})
	c.Assert(err, qt.IsNil)
	c.Assert(called, qt.IsTrue)
}
