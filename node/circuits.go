package node

import (
	"context"
	"fmt"
	"sync"

	"github.com/consensys/gnark/frontend"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/log"
)

// circuitRuntimes holds the four pinned circuit runtimes a node proves with.
// They are loaded once at startup so a missing artifact or a hash that does
// not match the release fails the process immediately, not at the first
// contribution deadline; loading also compiles each circuit and decodes its
// proving key, seconds of work that would otherwise repeat on every proof.
type circuitRuntimes struct {
	contribution   *circuits.CircuitRuntime
	finalize       *circuits.CircuitRuntime
	partialDecrypt *circuits.CircuitRuntime
	combine        *circuits.CircuitRuntime
}

// loadRuntimes is process-wide: every Node in the process (the daemon runs
// one, integration tests run several) shares one copy of each proving key.
var loadRuntimes = sync.OnceValues(func() (circuitRuntimes, error) {
	ctx := context.Background()
	load := func(name string, ca *circuits.CircuitArtifacts, circuit frontend.Circuit) (*circuits.CircuitRuntime, error) {
		rt, err := ca.LoadPinned(ctx, circuit)
		if err != nil {
			return nil, fmt.Errorf("%s circuit: %w", name, err)
		}
		log.Infow("circuit artifacts loaded", "circuit", name, "constraints", rt.ConstraintSystem().GetNbConstraints())
		return rt, nil
	}
	var rts circuitRuntimes
	var err error
	if rts.contribution, err = load("contribution", contribution.Artifacts, &contribution.ContributionCircuit{}); err != nil {
		return circuitRuntimes{}, err
	}
	if rts.finalize, err = load("finalize", finalize.Artifacts, &finalize.FinalizeCircuit{}); err != nil {
		return circuitRuntimes{}, err
	}
	if rts.partialDecrypt, err = load("partialdecrypt", partialdecrypt.Artifacts, &partialdecrypt.PartialDecryptCircuit{}); err != nil {
		return circuitRuntimes{}, err
	}
	if rts.combine, err = load("decryptcombine", decryptcombine.Artifacts, &decryptcombine.DecryptCombineCircuit{}); err != nil {
		return circuitRuntimes{}, err
	}
	return rts, nil
})
