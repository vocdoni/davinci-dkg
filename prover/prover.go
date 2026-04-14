package prover

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
)

// Prove runs the configured proving backend for the provided assignment.
func Prove(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	assignment frontend.Circuit,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	return prover(curve, ccs, pk, assignment, opts...)
}

func defaultProver(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	assignment frontend.Circuit,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	return CPUProver(curve, ccs, pk, assignment, opts...)
}

// CPUProver is the standard CPU implementation.
func CPUProver(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	assignment frontend.Circuit,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	w, err := frontend.NewWitness(assignment, curve.ScalarField())
	if err != nil {
		return nil, fmt.Errorf("failed to create witness: %w", err)
	}
	return groth16.Prove(ccs, pk, w, opts...)
}

// GPUProver is not available in the default build.
func GPUProver(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	assignment frontend.Circuit,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	panic("GPU prover not supported in this build")
}

// ProveWithWitness proves using an already-materialized witness.
func ProveWithWitness(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	w witness.Witness,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	return CPUProverWithWitness(curve, ccs, pk, w, opts...)
}

// CPUProverWithWitness proves with the CPU backend and a prebuilt witness.
func CPUProverWithWitness(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	w witness.Witness,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	return groth16.Prove(ccs, pk, w, opts...)
}

// GPUProverWithWitness is not available in the default build.
func GPUProverWithWitness(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	w witness.Witness,
	opts ...backend.ProverOption,
) (groth16.Proof, error) {
	panic("GPU prover not supported in this build")
}
