package service

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

type RoundReader interface {
	GetRound(ctx context.Context, roundID [12]byte) (web3.RoundView, error)
	SelectedParticipants(ctx context.Context, roundID [12]byte) ([]common.Address, error)
}

type RoundMonitor struct {
	contracts RoundReader
	storage   *storage.Storage
}

func NewRoundMonitor(contracts RoundReader, st *storage.Storage) *RoundMonitor {
	if st == nil {
		st = storage.New()
	}
	return &RoundMonitor{
		contracts: contracts,
		storage:   st,
	}
}

func (m *RoundMonitor) SyncRound(ctx context.Context, roundID [12]byte) error {
	roundView, err := m.contracts.GetRound(ctx, roundID)
	if err != nil {
		return err
	}
	selected, err := m.contracts.SelectedParticipants(ctx, roundID)
	if err != nil {
		return err
	}

	round := types.Round{
		ID:        strings.TrimRight(string(roundID[:]), "\x00"),
		Organizer: roundView.Organizer,
		Policy: types.RoundPolicy{
			Threshold:                 roundView.Policy.Threshold,
			CommitteeSize:             roundView.Policy.CommitteeSize,
			MinValidContributions:     roundView.Policy.MinValidContributions,
			LotteryAlphaBps:           roundView.Policy.LotteryAlphaBps,
			SeedDelay:                 roundView.Policy.SeedDelay,
			RegistrationDeadlineBlock: roundView.Policy.RegistrationDeadlineBlock,
			ContributionDeadlineBlock: roundView.Policy.ContributionDeadlineBlock,
			DisclosureAllowed:         roundView.Policy.DisclosureAllowed,
		},
		Phase:                mapRoundPhase(roundView.Status),
		SelectedParticipants: selected,
	}
	return m.storage.UpsertRound(round)
}

// mapRoundPhase converts an on-chain DKGTypes.RoundStatus uint8 to the
// Go-side RoundPhase constant. The on-chain enum is:
//
//	None=0, Readiness=1, Contribution=2, Finalized=3, Aborted=4, Completed=5
//
// Note: RoundPhaseDecryption and RoundPhaseDisclosure are Go-side phases used
// for local state tracking; they are never mapped from chain status (the chain
// keeps the Finalized status throughout both decryption and disclosure).
func mapRoundPhase(status uint8) types.RoundPhase {
	switch status {
	case 1:
		return types.RoundPhaseRegistration
	case 2:
		return types.RoundPhaseContribution
	case 3:
		return types.RoundPhaseFinalized
	case 4:
		return types.RoundPhaseAborted
	case 5:
		return types.RoundPhaseCompleted
	default:
		return types.RoundPhaseUnknown
	}
}
