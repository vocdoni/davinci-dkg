package partialdecrypt

import (
	"fmt"

	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/vocdoni/davinci-dkg/log"
)

// Compile compiles the partial decrypt circuit definition.
func Compile() (constraint.ConstraintSystem, error) {
	log.Infow("compiling circuit definition", "circuit", Artifacts.Name())
	ccs, err := frontend.Compile(Artifacts.Curve().ScalarField(), r1cs.NewBuilder, &PartialDecryptCircuit{})
	if err != nil {
		return nil, fmt.Errorf("compile partial decrypt circuit: %w", err)
	}
	return ccs, nil
}
