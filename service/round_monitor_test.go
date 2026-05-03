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

func TestEpochMonitorSyncEpoch(t *testing.T) {
	c := qt.New(t)

	contracts := &mockContracts{
		epoch: web3.EpochView{
			Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			Policy: web3.EpochPolicy{
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
	monitor := NewEpochMonitor(contracts, st)

	var epochID [12]byte
	copy(epochID[:], []byte("epoch-1"))

	err := monitor.SyncEpoch(context.Background(), epochID)
	c.Assert(err, qt.IsNil)

	epoch, err := st.Epoch("epoch-1")
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Organizer, qt.Equals, contracts.epoch.Organizer)
	c.Assert(epoch.Phase, qt.Equals, types.EpochPhaseContribution)
	c.Assert(epoch.SelectedParticipants, qt.DeepEquals, contracts.selected)
}

func TestEpochMonitorSyncEpochUpdatesExistingSnapshot(t *testing.T) {
	c := qt.New(t)

	contracts := &mockContracts{
		epoch: web3.EpochView{
			Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
			Policy: web3.EpochPolicy{
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
	monitor := NewEpochMonitor(contracts, st)

	var epochID [12]byte
	copy(epochID[:], []byte("epoch-1"))

	c.Assert(monitor.SyncEpoch(context.Background(), epochID), qt.IsNil)

	contracts.epoch.Status = 3
	contracts.selected = []common.Address{
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
		common.HexToAddress("0x3000000000000000000000000000000000000003"),
	}

	c.Assert(monitor.SyncEpoch(context.Background(), epochID), qt.IsNil)

	epoch, err := st.Epoch("epoch-1")
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Phase, qt.Equals, types.EpochPhaseFinalized)
	c.Assert(epoch.SelectedParticipants, qt.DeepEquals, contracts.selected)
}

type mockContracts struct {
	epoch    web3.EpochView
	selected []common.Address
}

func TestMapEpochPhase(t *testing.T) {
	c := qt.New(t)

	// Solidity DKGTypes.EpochPhase: None=0, Readiness=1, Contribution=2, Finalized=3, Aborted=4, Completed=5
	c.Assert(mapEpochPhase(0), qt.Equals, types.EpochPhaseUnknown)
	c.Assert(mapEpochPhase(1), qt.Equals, types.EpochPhaseRegistration)
	c.Assert(mapEpochPhase(2), qt.Equals, types.EpochPhaseContribution)
	c.Assert(mapEpochPhase(3), qt.Equals, types.EpochPhaseFinalized)
	c.Assert(mapEpochPhase(4), qt.Equals, types.EpochPhaseAborted)
	c.Assert(mapEpochPhase(5), qt.Equals, types.EpochPhaseCompleted)
	c.Assert(mapEpochPhase(99), qt.Equals, types.EpochPhaseUnknown)
}

func (m *mockContracts) GetEpoch(_ context.Context, _ [12]byte) (web3.EpochView, error) {
	return m.epoch, nil
}

func (m *mockContracts) SelectedParticipants(_ context.Context, _ [12]byte) ([]common.Address, error) {
	return append([]common.Address(nil), m.selected...), nil
}
