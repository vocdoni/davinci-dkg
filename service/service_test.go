package service

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestContributionPlannerPendingContribution(t *testing.T) {
	c := qt.New(t)

	operator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	other := common.HexToAddress("0x2000000000000000000000000000000000000002")
	st := storage.New()
	c.Assert(st.SaveEpoch(types.Epoch{
		ID:        "epoch-1",
		Organizer: operator,
		Policy: types.EpochPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase:                types.EpochPhaseContribution,
		SelectedParticipants: []common.Address{operator, other},
	}), qt.IsNil)

	planner := NewContributor(operator, st)
	task, err := planner.PendingContribution("epoch-1")

	c.Assert(err, qt.IsNil)
	c.Assert(task, qt.Not(qt.IsNil))
	c.Assert(task.EpochID, qt.Equals, "epoch-1")
	c.Assert(task.ContributorIndex, qt.Equals, uint16(1))
}

func TestContributionPlannerSkipsWhenAlreadyContributed(t *testing.T) {
	c := qt.New(t)

	operator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	st := seededContributionRound(c, operator)
	c.Assert(st.SaveContribution(types.Contribution{
		EpochID:          "epoch-1",
		Contributor:      operator,
		ContributorIndex: 1,
		Commitments:      []types.CurvePoint{{X: one(), Y: one()}},
		EncryptedShares: []types.EncryptedShare{{
			Recipient:      operator,
			RecipientIndex: 1,
			Ephemeral:      types.CurvePoint{X: one(), Y: one()},
			Ciphertext:     one(),
		}},
		Proof: []byte{1},
	}), qt.IsNil)

	task, err := NewContributor(operator, st).PendingContribution("epoch-1")
	c.Assert(err, qt.IsNil)
	c.Assert(task, qt.IsNil)
}

func TestFinalizerPendingFinalize(t *testing.T) {
	c := qt.New(t)

	operator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	st := seededContributionRound(c, operator)
	saveContributionFor(c, st, "epoch-1", operator, 1)
	saveContributionFor(c, st, "epoch-1", common.HexToAddress("0x2000000000000000000000000000000000000002"), 2)

	task, err := NewFinalizer(st).PendingFinalize("epoch-1")
	c.Assert(err, qt.IsNil)
	c.Assert(task, qt.Not(qt.IsNil))
	c.Assert(task.EpochID, qt.Equals, "epoch-1")
	c.Assert(task.ContributionCount, qt.Equals, 2)
}

func TestDecryptorPendingPartialDecryption(t *testing.T) {
	c := qt.New(t)

	operator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	other := common.HexToAddress("0x2000000000000000000000000000000000000002")
	st := storage.New()
	c.Assert(st.SaveEpoch(types.Epoch{
		ID:        "epoch-1",
		Organizer: operator,
		Policy: types.EpochPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase:                types.EpochPhaseFinalized,
		SelectedParticipants: []common.Address{operator, other},
	}), qt.IsNil)

	task, err := NewDecryptor(operator, st).PendingPartialDecryption("epoch-1", 1)
	c.Assert(err, qt.IsNil)
	c.Assert(task, qt.Not(qt.IsNil))
	c.Assert(task.ParticipantIndex, qt.Equals, uint16(1))
	c.Assert(task.CiphertextIndex, qt.Equals, uint16(1))
}

func TestCapabilitiesForPhaseTerminalStates(t *testing.T) {
	c := qt.New(t)

	for _, phase := range []types.EpochPhase{types.EpochPhaseAborted, types.EpochPhaseCompleted, types.EpochPhaseUnknown} {
		caps := CapabilitiesForPhase(phase, 5, 3)
		c.Assert(caps.Contribution, qt.IsFalse, qt.Commentf("phase=%s should not allow contribution", phase))
		c.Assert(caps.Finalize, qt.IsFalse, qt.Commentf("phase=%s should not allow finalize", phase))
		c.Assert(caps.Decrypt, qt.IsFalse, qt.Commentf("phase=%s should not allow decrypt", phase))
	}
}

func TestDecryptorSkipsWhenRoundIsAborted(t *testing.T) {
	c := qt.New(t)

	operator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	st := storage.New()
	c.Assert(st.SaveEpoch(types.Epoch{
		ID:        "epoch-1",
		Organizer: operator,
		Policy: types.EpochPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase:                types.EpochPhaseAborted,
		SelectedParticipants: []common.Address{operator},
	}), qt.IsNil)

	task, err := NewDecryptor(operator, st).PendingPartialDecryption("epoch-1", 1)
	c.Assert(err, qt.IsNil)
	c.Assert(task, qt.IsNil)
}

func seededContributionRound(c *qt.C, operator common.Address) *storage.Storage {
	st := storage.New()
	c.Assert(st.SaveEpoch(types.Epoch{
		ID:        "epoch-1",
		Organizer: operator,
		Policy: types.EpochPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.EpochPhaseContribution,
		SelectedParticipants: []common.Address{
			operator,
			common.HexToAddress("0x2000000000000000000000000000000000000002"),
		},
	}), qt.IsNil)
	return st
}

func saveContributionFor(c *qt.C, st *storage.Storage, epochID string, operator common.Address, index uint16) {
	c.Assert(st.SaveContribution(types.Contribution{
		EpochID:          epochID,
		Contributor:      operator,
		ContributorIndex: index,
		Commitments:      []types.CurvePoint{{X: one(), Y: one()}},
		EncryptedShares: []types.EncryptedShare{{
			Recipient:      operator,
			RecipientIndex: index,
			Ephemeral:      types.CurvePoint{X: one(), Y: one()},
			Ciphertext:     one(),
		}},
		Proof: []byte{1},
	}), qt.IsNil)
}

func one() *big.Int {
	return big.NewInt(1)
}
