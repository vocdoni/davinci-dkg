package storage

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/db/metadb"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestSaveAndGetRound(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseRegistration,
	}

	err := st.SaveRound(round)
	c.Assert(err, qt.IsNil)

	got, err := st.Round(round.ID)
	c.Assert(err, qt.IsNil)
	c.Assert(got.ID, qt.Equals, round.ID)
	c.Assert(got.Organizer, qt.Equals, round.Organizer)
	c.Assert(got.Phase, qt.Equals, types.RoundPhaseRegistration)
}

func TestMarkReadyTracksDistinctParticipants(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseRegistration,
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	operator := common.HexToAddress("0x2000000000000000000000000000000000000002")

	c.Assert(st.MarkReady(round.ID, operator), qt.IsNil)
	c.Assert(st.ReadyCount(round.ID), qt.Equals, 1)

	err := st.MarkReady(round.ID, operator)
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "already marked ready")
}

func TestSetSelectedParticipants(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseRegistration,
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	participants := []common.Address{
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
		common.HexToAddress("0x3000000000000000000000000000000000000003"),
	}

	err := st.SetSelectedParticipants(round.ID, participants)
	c.Assert(err, qt.IsNil)

	got, err := st.Round(round.ID)
	c.Assert(err, qt.IsNil)
	c.Assert(got.Phase, qt.Equals, types.RoundPhaseContribution)
	c.Assert(got.SelectedParticipants, qt.DeepEquals, participants)
}

func TestSaveContribution(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseContribution,
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	contribution := types.Contribution{
		RoundID:          round.ID,
		Contributor:      common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributorIndex: 1,
		Commitments:      []types.CurvePoint{{X: big.NewInt(1), Y: big.NewInt(2)}},
		EncryptedShares: []types.EncryptedShare{{
			Recipient:      common.HexToAddress("0x3000000000000000000000000000000000000003"),
			RecipientIndex: 2,
			Ephemeral:      types.CurvePoint{X: big.NewInt(3), Y: big.NewInt(4)},
			Ciphertext:     big.NewInt(5),
		}},
		Proof: []byte{0x01},
	}

	c.Assert(st.SaveContribution(contribution), qt.IsNil)

	got, err := st.Contribution(round.ID, contribution.Contributor)
	c.Assert(err, qt.IsNil)
	c.Assert(got.ContributorIndex, qt.Equals, uint16(1))
	c.Assert(len(st.Contributions(round.ID)), qt.Equals, 1)
}

func TestSavePartialDecryptionAndRevealedShare(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseDecryption,
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	decryption := types.PartialDecryption{
		RoundID:          round.ID,
		Participant:      common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ParticipantIndex: 1,
		CiphertextIndex:  1,
		Delta:            types.CurvePoint{X: big.NewInt(8), Y: big.NewInt(9)},
		Proof:            []byte{0x02},
	}
	c.Assert(st.SavePartialDecryption(decryption), qt.IsNil)
	c.Assert(len(st.PartialDecryptions(round.ID)), qt.Equals, 1)

	share := types.RevealedShare{
		RoundID:          round.ID,
		Participant:      common.HexToAddress("0x3000000000000000000000000000000000000003"),
		ParticipantIndex: 2,
		Share:            big.NewInt(12),
	}
	c.Assert(st.SaveRevealedShare(share), qt.IsNil)
	c.Assert(len(st.RevealedShares(round.ID)), qt.Equals, 1)
}

func TestSavePartialDecryption_AllowsDistinctCiphertextsForSameParticipant(t *testing.T) {
	c := qt.New(t)

	st := New()
	round := types.Round{
		ID: "round-2",
		Policy: types.RoundPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	decryption1 := types.PartialDecryption{
		RoundID:          round.ID,
		Participant:      common.HexToAddress("0x1000000000000000000000000000000000000001"),
		ParticipantIndex: 1,
		CiphertextIndex:  1,
		Delta:            types.CurvePoint{X: big.NewInt(10), Y: big.NewInt(20)},
		Proof:            []byte{1},
	}
	decryption2 := types.PartialDecryption{
		RoundID:          round.ID,
		Participant:      decryption1.Participant,
		ParticipantIndex: 1,
		CiphertextIndex:  2,
		Delta:            types.CurvePoint{X: big.NewInt(30), Y: big.NewInt(40)},
		Proof:            []byte{2},
	}

	c.Assert(st.SavePartialDecryption(decryption1), qt.IsNil)
	c.Assert(st.SavePartialDecryption(decryption2), qt.IsNil)

	decryptions := st.PartialDecryptions(round.ID)
	c.Assert(decryptions, qt.HasLen, 2)
}

func TestPebbleBackedStorageRoundAndContribution(t *testing.T) {
	c := qt.New(t)

	database := metadb.NewTest(t)
	st := NewWithDB(database)

	round := types.Round{
		ID:        "round-1",
		Organizer: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Policy: types.RoundPolicy{
			Threshold:                 2,
			CommitteeSize:             2,
			MinValidContributions:     2,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		},
		Phase: types.RoundPhaseContribution,
	}
	c.Assert(st.SaveRound(round), qt.IsNil)

	storedRound, err := st.Round(round.ID)
	c.Assert(err, qt.IsNil)
	c.Assert(storedRound.ID, qt.Equals, round.ID)

	contribution := types.Contribution{
		RoundID:          round.ID,
		Contributor:      common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributorIndex: 1,
		Commitments:      []types.CurvePoint{{X: big.NewInt(1), Y: big.NewInt(2)}},
		EncryptedShares: []types.EncryptedShare{{
			Recipient:      common.HexToAddress("0x3000000000000000000000000000000000000003"),
			RecipientIndex: 2,
			Ephemeral:      types.CurvePoint{X: big.NewInt(3), Y: big.NewInt(4)},
			Ciphertext:     big.NewInt(5),
		}},
		Proof: []byte{0x01},
	}
	c.Assert(st.SaveContribution(contribution), qt.IsNil)
	c.Assert(len(st.Contributions(round.ID)), qt.Equals, 1)
}
