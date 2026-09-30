package types

import "fmt"

// EpochPhase identifies the current lifecycle phase of a DKG epoch.
type EpochPhase uint8

const (
	EpochPhaseUnknown EpochPhase = iota
	EpochPhaseCommitteeSelection
	EpochPhaseKeyAssembly
	EpochPhaseLive
	EpochPhaseDecryption
	EpochPhaseAborted
	EpochPhaseCompleted
)

func (p EpochPhase) String() string {
	switch p {
	case EpochPhaseCommitteeSelection:
		return "registration"
	case EpochPhaseKeyAssembly:
		return "contribution"
	case EpochPhaseLive:
		return "finalized"
	case EpochPhaseDecryption:
		return "decryption"
	case EpochPhaseAborted:
		return "aborted"
	case EpochPhaseCompleted:
		return "completed"
	default:
		return "unknown"
	}
}

// EpochPolicy configures one DKG epoch. The phase deadline blocks are derived
// on chain at createEpoch from the manager's deadline offsets, so
// callers only supply the four policy fields createEpoch takes.
//
// The deadline-block fields mirror the on-chain EpochPolicy struct, which
// createEpoch populates for phase checks. Callers
// constructing a fresh policy for createEpoch may leave them zero; helpers
// fill the per-phase blocks from the on-chain `getEpoch` view after
// creation.
type EpochPolicy struct {
	Threshold                       uint16
	CommitteeSize                   uint16
	MinValidContributions           uint16
	LotteryAlphaBps                 uint16
	CommitteeSelectionDeadlineBlock uint64
	KeyAssemblyDeadlineBlock        uint64
	LiveNotBeforeBlock              uint64
}

// Validate checks that the policy is internally coherent.
func (p EpochPolicy) Validate() error {
	if p.Threshold == 0 || p.CommitteeSize == 0 {
		return fmt.Errorf("threshold and committee size must be non-zero")
	}
	if p.Threshold > p.CommitteeSize {
		return fmt.Errorf("threshold cannot exceed committee size")
	}
	if p.MinValidContributions == 0 || p.MinValidContributions > p.CommitteeSize {
		return fmt.Errorf("min valid contributions out of range")
	}
	if p.LotteryAlphaBps < 10000 {
		return fmt.Errorf("lottery alpha must be at least 1.0 (10000 bps)")
	}
	// Phase deadlines are derived on chain, so there is nothing to validate
	// about them here.
	return nil
}
