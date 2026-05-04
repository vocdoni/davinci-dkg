package service

import "github.com/vocdoni/davinci-dkg/types"

// PhaseCapabilities captures which worker actions are legal for a given phase.
type PhaseCapabilities struct {
	Contribution bool
	Finalize     bool
	Decrypt      bool
}

// CapabilitiesForPhase centralizes the service-side phase gating rules.
//
// On-chain status mapping:
//
//	Registration(1), Contribution(2), Finalized(3), Aborted(4), Completed(5)
//
// The Finalized on-chain status persists throughout decryption operations;
// the Go-only EpochPhaseDecryption phase is used for local state
// refinement and carries the same capabilities as EpochPhaseLive.
func CapabilitiesForPhase(phase types.EpochPhase, contributionCount int, minValidContributions uint16) PhaseCapabilities {
	caps := PhaseCapabilities{}
	switch phase {
	case types.EpochPhaseKeyAssembly:
		caps.Contribution = true
		if contributionCount >= int(minValidContributions) {
			caps.Finalize = true
		}
	case types.EpochPhaseLive, types.EpochPhaseDecryption:
		caps.Decrypt = true
		// EpochPhaseAborted and EpochPhaseCompleted are terminal states.
		// No further capabilities are granted.
	}
	return caps
}
