package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestCommitteeRoundHappyPath(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	actor1, err := services.Actor(1)
	c.Assert(err, qt.IsNil)
	actor2, err := services.Actor(2)
	c.Assert(err, qt.IsNil)

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.RoundPolicy{
		Threshold:                 2,
		CommitteeSize:             3,
		MinValidContributions:     2,
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		SeedDelay:                 helpers.DefaultSeedDelay,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         true,
	}

	roundID, err := helpers.CreateRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	// Lottery flow: advance past seedBlock so blockhash is available, then have
	// each registered actor self-claim a slot. The committee fills first-come
	// first-served; there is no organizer-driven SelectParticipants step.
	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, roundID), qt.IsNil)

	committee := []common.Address{services.TxManager.Address(), actor1.Address(), actor2.Address()}
	_ = committee
	round, err := helpers.WaitRoundStatus(ctx, services, roundID, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Policy.CommitteeSize, qt.Equals, uint16(3))

	contributions := [][]*big.Int{
		{big.NewInt(3), big.NewInt(1)},
		{big.NewInt(5), big.NewInt(2)},
		{big.NewInt(7), big.NewInt(4)},
	}
	recipientIndexes := []uint16{1, 2, 3}

	submission0, err := helpers.BuildContributionSubmission(ctx, services, roundID, 2, 3, 1, contributions[0], recipientIndexes)
	c.Assert(err, qt.IsNil)
	submission1, err := helpers.BuildContributionSubmission(ctx, services, roundID, 2, 3, 2, contributions[1], recipientIndexes)
	c.Assert(err, qt.IsNil)
	submission2, err := helpers.BuildContributionSubmission(ctx, services, roundID, 2, 3, 3, contributions[2], recipientIndexes)
	c.Assert(err, qt.IsNil)

	c.Assert(
		helpers.SubmitContributionAs(ctx, &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}, roundID, 1, submission0.CommitmentsHash, submission0.EncryptedSharesHash, submission0.Commitment0X, submission0.Commitment0Y, submission0.Transcript, submission0.Proof, submission0.Input),
		qt.IsNil,
	)
	c.Assert(
		helpers.SubmitContributionAs(ctx, actor1, roundID, 2, submission1.CommitmentsHash, submission1.EncryptedSharesHash, submission1.Commitment0X, submission1.Commitment0Y, submission1.Transcript, submission1.Proof, submission1.Input),
		qt.IsNil,
	)
	c.Assert(
		helpers.SubmitContributionAs(ctx, actor2, roundID, 3, submission2.CommitmentsHash, submission2.EncryptedSharesHash, submission2.Commitment0X, submission2.Commitment0Y, submission2.Transcript, submission2.Proof, submission2.Input),
		qt.IsNil,
	)

	finalizeOutput, err := helpers.BuildFinalizeRoundOutput(ctx, roundID, 2, 3, recipientIndexes, contributions)
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeRound(
		auth,
		roundID,
		finalizeOutput.AggregateCommitmentsHash,
		finalizeOutput.CollectivePublicKeyHash,
		finalizeOutput.ShareCommitmentHash,
		finalizeOutput.Transcript,
		finalizeOutput.Proof,
		finalizeOutput.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	round, err = helpers.WaitRoundStatus(ctx, services, roundID, 3)
	c.Assert(err, qt.IsNil)
	c.Assert(round.ContributionCount, qt.Equals, uint16(3))

	recoveredShares, err := helpers.RecoverParticipantShares(contributions, recipientIndexes)
	c.Assert(err, qt.IsNil)
	c.Assert(len(recoveredShares), qt.Equals, 3)

	partial0, err := helpers.BuildPartialDecryptionSubmission(ctx, roundID, 1, big.NewInt(9), recoveredShares[0], big.NewInt(11))
	c.Assert(err, qt.IsNil)
	partial1, err := helpers.BuildPartialDecryptionSubmission(ctx, roundID, 2, big.NewInt(9), recoveredShares[1], big.NewInt(13))
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}, roundID, 1, 1, partial0.DeltaHash, partial0.Proof, partial0.Input), qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, actor1, roundID, 2, 1, partial1.DeltaHash, partial1.Proof, partial1.Input), qt.IsNil)

	combineOutput, err := helpers.BuildDecryptCombineOutput(
		ctx,
		roundID,
		2,
		big.NewInt(9),
		[]uint16{1, 2},
		[]types.CurvePoint{partial0.Delta, partial1.Delta},
		big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.Manager.CombineDecryption(
		auth,
		roundID,
		1,
		combineOutput.CombineHash,
		combineOutput.PlaintextHash,
		combineOutput.Transcript,
		combineOutput.Proof,
		combineOutput.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	submission0Reveal, err := helpers.BuildRevealShareSubmission(ctx, roundID, 1, recoveredShares[0], finalizeOutput.ShareCommitments[0])
	c.Assert(err, qt.IsNil)
	submission1Reveal, err := helpers.BuildRevealShareSubmission(ctx, roundID, 2, recoveredShares[1], finalizeOutput.ShareCommitments[1])
	c.Assert(err, qt.IsNil)
	disclosureOutput, err := helpers.BuildRevealShareOutput(ctx, roundID, 2, []uint16{1, 2}, []*big.Int{recoveredShares[0], recoveredShares[1]})
	c.Assert(err, qt.IsNil)

	c.Assert(
		helpers.SubmitRevealedShareAs(
			ctx,
			&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
			roundID,
			1,
			submission0Reveal.ShareValue,
			submission0Reveal.Proof,
			submission0Reveal.Input,
		),
		qt.IsNil,
	)
	c.Assert(
		helpers.SubmitRevealedShareAs(ctx, actor1, roundID, 2, submission1Reveal.ShareValue, submission1Reveal.Proof, submission1Reveal.Input),
		qt.IsNil,
	)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.Manager.ReconstructSecret(
		auth,
		roundID,
		disclosureOutput.DisclosureHash,
		disclosureOutput.ReconstructedSecretHash,
		disclosureOutput.Transcript,
		disclosureOutput.Proof,
		disclosureOutput.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	round, err = services.Contracts.GetRound(ctx, roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Status, qt.Equals, uint8(5))
}
