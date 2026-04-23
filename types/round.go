package types

import "fmt"

// RoundPhase identifies the current lifecycle phase of a DKG round.
type RoundPhase uint8

const (
	RoundPhaseUnknown RoundPhase = iota
	RoundPhaseRegistration
	RoundPhaseContribution
	RoundPhaseFinalized
	RoundPhaseDecryption
	RoundPhaseDisclosure
	RoundPhaseAborted
	RoundPhaseCompleted
)

func (p RoundPhase) String() string {
	switch p {
	case RoundPhaseRegistration:
		return "registration"
	case RoundPhaseContribution:
		return "contribution"
	case RoundPhaseFinalized:
		return "finalized"
	case RoundPhaseDecryption:
		return "decryption"
	case RoundPhaseDisclosure:
		return "disclosure"
	case RoundPhaseAborted:
		return "aborted"
	case RoundPhaseCompleted:
		return "completed"
	default:
		return "unknown"
	}
}

// RoundPolicy configures the thresholds and phase windows for one DKG round.
type RoundPolicy struct {
	Threshold                 uint16
	CommitteeSize             uint16
	MinValidContributions     uint16
	LotteryAlphaBps           uint16
	SeedDelay                 uint16
	RegistrationDeadlineBlock uint64
	ContributionDeadlineBlock uint64
	// FinalizeNotBeforeBlock is the earliest block at which finalizeRound can
	// succeed. Must be strictly greater than ContributionDeadlineBlock; allows
	// every selected participant time to submit before the contribution set is
	// frozen.
	FinalizeNotBeforeBlock uint64
	DisclosureAllowed      bool
	DecryptionPolicy       DecryptionPolicy
}

// DecryptionPolicy mirrors the on-chain DKGTypes.DecryptionPolicy struct and
// gates who may call submitCiphertext for a round. All checks AND together;
// a zero-valued field is a no-op for that check.
type DecryptionPolicy struct {
	OwnerOnly          bool
	MaxDecryptions     uint16
	NotBeforeBlock     uint64
	NotBeforeTimestamp uint64
	NotAfterBlock      uint64
	NotAfterTimestamp  uint64
}

// Validate checks that the policy is internally coherent.
func (p RoundPolicy) Validate() error {
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
	if p.SeedDelay == 0 || p.SeedDelay > 256 {
		return fmt.Errorf("seed delay must be in (0, 256]")
	}
	if p.RegistrationDeadlineBlock == 0 || p.ContributionDeadlineBlock == 0 {
		return fmt.Errorf("deadline blocks must be non-zero")
	}
	if p.ContributionDeadlineBlock <= p.RegistrationDeadlineBlock {
		return fmt.Errorf("contribution deadline must be after registration deadline")
	}
	if p.FinalizeNotBeforeBlock <= p.ContributionDeadlineBlock {
		return fmt.Errorf("finalize-not-before block must be after contribution deadline")
	}
	return nil
}
