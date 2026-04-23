package service

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

func TestRoundMonitorSyncRound(t *testing.T) {
	c := qt.New(t)

	contracts := &mockContracts{
		round: web3.RoundView{
			Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			Policy: web3.RoundPolicy{
				Threshold:                 2,
				CommitteeSize:             2,
				MinValidContributions:     2,
				RegistrationDeadlineBlock: 10,
				ContributionDeadlineBlock: 20,
				FinalizeNotBeforeBlock:    21,
			},
			Status: 2,
		},
		selected: []common.Address{
			common.HexToAddress("0x2000000000000000000000000000000000000002"),
			common.HexToAddress("0x3000000000000000000000000000000000000003"),
		},
	}
	st := storage.New()
	monitor := NewRoundMonitor(contracts, st)

	var roundID [12]byte
	copy(roundID[:], []byte("round-1"))

	err := monitor.SyncRound(context.Background(), roundID)
	c.Assert(err, qt.IsNil)

	round, err := st.Round("round-1")
	c.Assert(err, qt.IsNil)
	c.Assert(round.Organizer, qt.Equals, contracts.round.Organizer)
	c.Assert(round.Phase, qt.Equals, types.RoundPhaseContribution)
	c.Assert(round.SelectedParticipants, qt.DeepEquals, contracts.selected)
}

func TestRoundMonitorSyncRoundUpdatesExistingSnapshot(t *testing.T) {
	c := qt.New(t)

	contracts := &mockContracts{
		round: web3.RoundView{
			Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			Policy: web3.RoundPolicy{
				Threshold:                 2,
				CommitteeSize:             2,
				MinValidContributions:     2,
				RegistrationDeadlineBlock: 10,
				ContributionDeadlineBlock: 20,
				FinalizeNotBeforeBlock:    21,
			},
			Status: 1,
		},
		selected: []common.Address{
			common.HexToAddress("0x2000000000000000000000000000000000000002"),
		},
	}
	st := storage.New()
	monitor := NewRoundMonitor(contracts, st)

	var roundID [12]byte
	copy(roundID[:], []byte("round-1"))

	c.Assert(monitor.SyncRound(context.Background(), roundID), qt.IsNil)

	contracts.round.Status = 3
	contracts.selected = []common.Address{
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
		common.HexToAddress("0x3000000000000000000000000000000000000003"),
	}

	c.Assert(monitor.SyncRound(context.Background(), roundID), qt.IsNil)

	round, err := st.Round("round-1")
	c.Assert(err, qt.IsNil)
	c.Assert(round.Phase, qt.Equals, types.RoundPhaseFinalized)
	c.Assert(round.SelectedParticipants, qt.DeepEquals, contracts.selected)
}

type mockContracts struct {
	round    web3.RoundView
	selected []common.Address
}

func TestMapRoundPhase(t *testing.T) {
	c := qt.New(t)

	// Solidity DKGTypes.RoundStatus: None=0, Readiness=1, Contribution=2, Finalized=3, Aborted=4, Completed=5
	c.Assert(mapRoundPhase(0), qt.Equals, types.RoundPhaseUnknown)
	c.Assert(mapRoundPhase(1), qt.Equals, types.RoundPhaseRegistration)
	c.Assert(mapRoundPhase(2), qt.Equals, types.RoundPhaseContribution)
	c.Assert(mapRoundPhase(3), qt.Equals, types.RoundPhaseFinalized)
	c.Assert(mapRoundPhase(4), qt.Equals, types.RoundPhaseAborted)
	c.Assert(mapRoundPhase(5), qt.Equals, types.RoundPhaseCompleted)
	c.Assert(mapRoundPhase(99), qt.Equals, types.RoundPhaseUnknown)
}

func (m *mockContracts) GetRound(_ context.Context, _ [12]byte) (web3.RoundView, error) {
	return m.round, nil
}

func (m *mockContracts) SelectedParticipants(_ context.Context, _ [12]byte) ([]common.Address, error) {
	return append([]common.Address(nil), m.selected...), nil
}
