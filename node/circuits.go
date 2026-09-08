package node

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"

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
// decryption circuits do: the contribution and finalize proving keys (243 MB
// and 436 MB on disk, a few GB decoded) are loaded for one proof and dropped
// right after, so a node at rest holds a few hundred MB instead of the 5 GB
// of keys v0.5 kept resident.
var circuitSpecs = [...]struct {
	name      string
	artifacts *circuits.CircuitArtifacts
	resident  bool
}{
	circuitContribution: {
		name:      "contribution",
		artifacts: contribution.Artifacts,
	},
	circuitFinalize: {
		name:      "finalize",
		artifacts: finalize.Artifacts,
	},
	circuitPartialDecrypt: {
		name:      "partialdecrypt",
		artifacts: partialdecrypt.Artifacts,
		resident:  true,
	},
	circuitCombine: {
		name:      "decryptcombine",
		artifacts: decryptcombine.Artifacts,
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
}

// circuitRuntime returns the pinned runtime of kind, loading it on first use.
func circuitRuntime(ctx context.Context, kind circuitKind) (*circuits.CircuitRuntime, error) {
	runtimeCache.Lock()
	defer runtimeCache.Unlock()
	if rt := runtimeCache.rt[kind]; rt != nil {
		return rt, nil
	}
	spec := circuitSpecs[kind]
	rt, err := spec.artifacts.LoadPinned(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s circuit: %w", spec.name, err)
	}
	log.Infow("circuit artifacts loaded", "circuit", spec.name, "constraints", rt.ConstraintSystem().GetNbConstraints())
	runtimeCache.rt[kind] = rt
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

// preloadRuntimes is the startup fail-fast: every artifact of the four
// circuits is downloaded if missing and stream-verified against its pinned
// hash (no decoding, no compilation, a few hundred MB of reads), so a bad
// release fails the process now rather than at the first deadline; then only
// the two resident decryption runtimes are decoded.
func preloadRuntimes(ctx context.Context) error {
	for _, spec := range circuitSpecs {
		if err := spec.artifacts.EnsureCached(ctx); err != nil {
			return err
		}
		log.Infow("circuit artifacts verified", "circuit", spec.name)
	}
	for kind, spec := range circuitSpecs {
		if !spec.resident {
			continue
		}
		if _, err := circuitRuntime(ctx, circuitKind(kind)); err != nil {
			return err
		}
	}
	return nil
}
