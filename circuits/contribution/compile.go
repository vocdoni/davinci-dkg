package contribution

import (
	"fmt"

	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/vocdoni/davinci-dkg/log"
)

// Compile compiles the contribution circuit definition.
func Compile() (constraint.ConstraintSystem, error) {
	log.Infow("compiling circuit definition", "circuit", Artifacts.Name())
	ccs, err := frontend.Compile(Artifacts.Curve().ScalarField(), r1cs.NewBuilder, &ContributionCircuit{})
	if err != nil {
		return nil, fmt.Errorf("compile contribution circuit: %w", err)
	}
	return ccs, nil
}
