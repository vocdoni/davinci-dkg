package prover

import (
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/vocdoni/davinci-dkg/log"
)

// UseGPUProver indicates whether to use the GPU prover when the binary is built
// with GPU support. The current repository ships CPU-only builds by default.
var UseGPUProver = false

func init() {
	if os.Getenv("GPU_PROVER") == "true" ||
		os.Getenv("GPU_PROVER") == "y" ||
		os.Getenv("GPU_PROVER") == "1" ||
		os.Getenv("GPU_PROVER") == "yes" {
		UseGPUProver = true
	}
	log.Infow("GPU prover usage", "enabled", UseGPUProver)
}

// ProverFunc defines the assignment-based proving entrypoint.
type ProverFunc func(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	assignment frontend.Circuit,
	opts ...backend.ProverOption,
) (groth16.Proof, error)

// ProverWithWitnessFunc defines the witness-based proving entrypoint.
type ProverWithWitnessFunc func(
	curve ecc.ID,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	w witness.Witness,
	opts ...backend.ProverOption,
) (groth16.Proof, error)

var prover ProverFunc = defaultProver

// SetProver replaces the default assignment-based prover.
func SetProver(p ProverFunc) {
	prover = p
}
