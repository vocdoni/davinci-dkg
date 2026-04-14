package prover

import (
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	gpugroth16 "github.com/consensys/gnark/backend/accelerated/icicle/groth16"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/vocdoni/davinci-dkg/log"
)

// Setup wraps groth16.Setup and keeps the same runtime boundary used by
// davinci-node.
func Setup(ccs constraint.ConstraintSystem) (groth16.ProvingKey, groth16.VerifyingKey, error) {
	start := time.Now()
	log.Debugw("generating circuit keys", "gpu", UseGPUProver, "constraints", ccs.GetNbConstraints())
	if UseGPUProver {
		pk, vk, err := gpugroth16.Setup(ccs)
		if err == nil {
			log.Debugw("circuit keys setup done",
				"gpu", UseGPUProver,
				"constraints", ccs.GetNbConstraints(),
				"elapsed", time.Since(start).String(),
			)
		}
		return pk, vk, err
	}
	pk, vk, err := groth16.Setup(ccs)
	if err == nil {
		log.Debugw("circuit keys setup done",
			"gpu", UseGPUProver,
			"constraints", ccs.GetNbConstraints(),
			"elapsed", time.Since(start).String(),
		)
	}
	return pk, vk, err
}

// NewProvingKey returns an empty proving key compatible with the active backend.
func NewProvingKey(curve ecc.ID) groth16.ProvingKey {
	if UseGPUProver {
		return gpugroth16.NewProvingKey(curve)
	}
	return groth16.NewProvingKey(curve)
}
