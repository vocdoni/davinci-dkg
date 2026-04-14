package service

import "github.com/vocdoni/davinci-dkg/types"

// PhaseCapabilities captures which worker actions are legal for a given phase.
type PhaseCapabilities struct {
	Contribution bool
	Finalize     bool
	Decrypt      bool
	Disclose     bool
}

// CapabilitiesForPhase centralizes the service-side phase gating rules.
//
// On-chain status mapping:
//
//	Readiness(1), Contribution(2), Finalized(3), Aborted(4), Completed(5)
//
// The Finalized on-chain status persists throughout both decryption and
// disclosure operations; the Go-only RoundPhaseDecryption/Disclosure phases
// are used for local state refinement and carry the same capabilities as
// RoundPhaseFinalized.
func CapabilitiesForPhase(phase types.RoundPhase, contributionCount int, minValidContributions uint16, disclosureAllowed bool) PhaseCapabilities {
	caps := PhaseCapabilities{}
	switch phase {
	case types.RoundPhaseContribution:
		caps.Contribution = true
		if contributionCount >= int(minValidContributions) {
			caps.Finalize = true
		}
	case types.RoundPhaseFinalized, types.RoundPhaseDecryption:
		caps.Decrypt = true
		if disclosureAllowed {
			caps.Disclose = true
		}
	case types.RoundPhaseDisclosure:
		caps.Disclose = disclosureAllowed
		// RoundPhaseAborted and RoundPhaseCompleted are terminal states.
		// No further capabilities are granted.
	}
	return caps
}
