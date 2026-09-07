package node

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/consensys/gnark/frontend"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/log"
)

// circuitKind names one of the four pinned circuits a node proves with.
type circuitKind int

const (
	circuitContribution circuitKind = iota
	circuitFinalize
	circuitPartialDecrypt
	circuitCombine
)

// circuitSpecs describes each circuit: its pinned artifacts, a fresh instance
// for compilation, and whether its runtime stays resident. Only the two small
// decryption circuits do: the contribution and finalize proving keys (800 MB
// and 446 MB on disk, several GB decoded) are loaded for one proof and
// dropped right after, so a coordinator at rest costs about as much memory
// as a warden, which never loads them at all.
var circuitSpecs = [...]struct {
	name      string
	artifacts *circuits.CircuitArtifacts
	circuit   func() frontend.Circuit
	resident  bool
}{
	circuitContribution: {
		name:      "contribution",
		artifacts: contribution.Artifacts,
		circuit:   func() frontend.Circuit { return &contribution.ContributionCircuit{} },
	},
	circuitFinalize: {
		name:      "finalize",
		artifacts: finalize.Artifacts,
		circuit:   func() frontend.Circuit { return &finalize.FinalizeCircuit{} },
	},
	circuitPartialDecrypt: {
		name:      "partialdecrypt",
		artifacts: partialdecrypt.Artifacts,
		circuit:   func() frontend.Circuit { return &partialdecrypt.PartialDecryptCircuit{} },
		resident:  true,
	},
	circuitCombine: {
		name:      "decryptcombine",
		artifacts: decryptcombine.Artifacts,
		circuit:   func() frontend.Circuit { return &decryptcombine.DecryptCombineCircuit{} },
		resident:  true,
	},
}

// runtimeCache is process-wide: the daemon runs one node, integration tests
// run several in one process and share each decoded key. Loading holds the
// lock (seconds for a heavy key), so a concurrent prover waits instead of
// decoding a second copy. Releasing a runtime another node of the process is
// still proving with is harmless: that prover keeps its pointer and the next
// caller reloads.
var runtimeCache struct {
	sync.Mutex
	rt [len(circuitSpecs)]*circuits.CircuitRuntime
	// verified marks circuits whose compiled form already matched the pinned
	// hash once (LoadPinned): later reloads read the hash-checked artifacts
	// without recompiling, 8 s instead of 19 s for the contribution circuit.
	verified [len(circuitSpecs)]bool
}

// circuitRuntime returns the pinned runtime of kind, loading it on first use.
func circuitRuntime(ctx context.Context, kind circuitKind) (*circuits.CircuitRuntime, error) {
	runtimeCache.Lock()
	defer runtimeCache.Unlock()
	if rt := runtimeCache.rt[kind]; rt != nil {
		return rt, nil
	}
	spec := circuitSpecs[kind]
	var rt *circuits.CircuitRuntime
	var err error
	if runtimeCache.verified[kind] {
		rt, err = spec.artifacts.LoadOrDownload(ctx)
	} else {
		rt, err = spec.artifacts.LoadPinned(ctx, spec.circuit())
	}
	if err != nil {
		return nil, fmt.Errorf("%s circuit: %w", spec.name, err)
	}
	log.Infow("circuit artifacts loaded", "circuit", spec.name, "constraints", rt.ConstraintSystem().GetNbConstraints())
	runtimeCache.rt[kind] = rt
	runtimeCache.verified[kind] = true
	return rt, nil
}

// releaseRuntime drops a non-resident runtime and hands its memory back to
// the OS right away instead of waiting for the scavenger.
func releaseRuntime(kind circuitKind) {
	if circuitSpecs[kind].resident {
		return
	}
	runtimeCache.Lock()
	runtimeCache.rt[kind] = nil
	runtimeCache.Unlock()
	debug.FreeOSMemory()
}

// preloadRuntimes loads every runtime the role will ever prove with once at
// startup, so a missing artifact or a hash that does not match the release
// fails the process immediately rather than at the first deadline, then
// drops the non-resident ones again.
func preloadRuntimes(ctx context.Context, warden bool) error {
	for kind, spec := range circuitSpecs {
		if warden && !spec.resident {
			continue
		}
		if _, err := circuitRuntime(ctx, circuitKind(kind)); err != nil {
			return err
		}
		releaseRuntime(circuitKind(kind))
	}
	return nil
}
