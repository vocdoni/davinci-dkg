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
	GetEpoch(ctx context.Context, epochID [12]byte) (web3.EpochView, error)
	SelectedParticipants(ctx context.Context, epochID [12]byte) ([]common.Address, error)
}

type EpochMonitor struct {
	contracts RoundReader
	storage   *storage.Storage
}

func NewEpochMonitor(contracts RoundReader, st *storage.Storage) *EpochMonitor {
	if st == nil {
		st = storage.New()
	}
	return &EpochMonitor{
		contracts: contracts,
		storage:   st,
	}
}

func (m *EpochMonitor) SyncEpoch(ctx context.Context, epochID [12]byte) error {
	roundView, err := m.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return err
	}
	selected, err := m.contracts.SelectedParticipants(ctx, epochID)
	if err != nil {
		return err
	}

	epoch := types.Epoch{
		ID:        strings.TrimRight(string(epochID[:]), "\x00"),
		Organizer: roundView.Organizer,
		Policy: types.EpochPolicy{
			Threshold:                 roundView.Policy.Threshold,
			CommitteeSize:             roundView.Policy.CommitteeSize,
			MinValidContributions:     roundView.Policy.MinValidContributions,
			LotteryAlphaBps:           roundView.Policy.LotteryAlphaBps,
			RegistrationDeadlineBlock: roundView.Policy.RegistrationDeadlineBlock,
			ContributionDeadlineBlock: roundView.Policy.ContributionDeadlineBlock,
			FinalizeNotBeforeBlock:    roundView.Policy.FinalizeNotBeforeBlock,
		},
		Phase:                mapEpochPhase(roundView.Status),
		SelectedParticipants: selected,
	}
	return m.storage.UpsertEpoch(epoch)
}

// mapEpochPhase converts an on-chain DKGTypes.EpochPhase uint8 to the
// Go-side EpochPhase constant. The on-chain enum is:
//
//	None=0, Readiness=1, Contribution=2, Finalized=3, Aborted=4, Completed=5
//
// Note: EpochPhaseDecryption is a Go-side phase used for local state
// tracking; it is never mapped from chain status (the chain keeps the
// Finalized status throughout the decryption pipeline).
func mapEpochPhase(status uint8) types.EpochPhase {
	switch status {
	case 1:
		return types.EpochPhaseRegistration
	case 2:
		return types.EpochPhaseContribution
	case 3:
		return types.EpochPhaseFinalized
	case 4:
		return types.EpochPhaseAborted
	case 5:
		return types.EpochPhaseCompleted
	default:
		return types.EpochPhaseUnknown
	}
}
