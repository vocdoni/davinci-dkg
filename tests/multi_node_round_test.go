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

	policy := types.EpochPolicy{
		Threshold:                 2,
		CommitteeSize:             3,
		MinValidContributions:     2,
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		FinalizeNotBeforeBlock:    head + 51,
	}

	epochID, err := helpers.CreateEpoch(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	// Lottery flow: advance past seedBlock so blockhash is available, then have
	// each registered actor self-claim a slot. The committee fills first-come
	// first-served; there is no organizer-driven SelectParticipants step.
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, epochID), qt.IsNil)

	committee := []common.Address{services.TxManager.Address(), actor1.Address(), actor2.Address()}
	_ = committee
	epoch, err := helpers.WaitEpochPhase(ctx, services, epochID, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Policy.CommitteeSize, qt.Equals, uint16(3))

	contributions := [][]*big.Int{
		{big.NewInt(3), big.NewInt(1)},
		{big.NewInt(5), big.NewInt(2)},
		{big.NewInt(7), big.NewInt(4)},
	}
	recipientIndexes := []uint16{1, 2, 3}

	submission0, err := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, 1, contributions[0], recipientIndexes)
	c.Assert(err, qt.IsNil)
	submission1, err := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, 2, contributions[1], recipientIndexes)
	c.Assert(err, qt.IsNil)
	submission2, err := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, 3, contributions[2], recipientIndexes)
	c.Assert(err, qt.IsNil)

	c.Assert(
		helpers.SubmitContributionAs(ctx, &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}, epochID, 1, submission0.CommitmentsHash, submission0.EncryptedSharesHash, submission0.Transcript, submission0.Proof, submission0.Input),
		qt.IsNil,
	)
	c.Assert(
		helpers.SubmitContributionAs(ctx, actor1, epochID, 2, submission1.CommitmentsHash, submission1.EncryptedSharesHash, submission1.Transcript, submission1.Proof, submission1.Input),
		qt.IsNil,
	)
	c.Assert(
		helpers.SubmitContributionAs(ctx, actor2, epochID, 3, submission2.CommitmentsHash, submission2.EncryptedSharesHash, submission2.Transcript, submission2.Proof, submission2.Input),
		qt.IsNil,
	)

	finalizeOutput, err := helpers.BuildFinalizeEpochOutput(ctx, epochID, 2, 3, recipientIndexes, contributions)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeEpoch(
		auth,
		epochID,
		finalizeOutput.AggregateCommitmentsHash,
		finalizeOutput.CollectivePublicKeyHash,
		finalizeOutput.ShareCommitmentHash,
		finalizeOutput.Transcript,
		finalizeOutput.Proof,
		finalizeOutput.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	epoch, err = helpers.WaitEpochPhase(ctx, services, epochID, 3)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.ContributionCount, qt.Equals, uint16(3))

	recoveredShares, err := helpers.RecoverParticipantShares(contributions, recipientIndexes)
	c.Assert(err, qt.IsNil)
	c.Assert(len(recoveredShares), qt.Equals, 3)

	partial0, err := helpers.BuildPartialDecryptionSubmission(ctx, epochID, 1, 1, big.NewInt(9), recoveredShares[0], big.NewInt(11))
	c.Assert(err, qt.IsNil)
	partial1, err := helpers.BuildPartialDecryptionSubmission(ctx, epochID, 1, 2, big.NewInt(9), recoveredShares[1], big.NewInt(13))
	c.Assert(err, qt.IsNil)

	combineOutput, err := helpers.BuildDecryptCombineOutput(
		ctx,
		epochID,
		1, // ciphertextIndex
		2,
		big.NewInt(9),
		[]uint16{1, 2},
		[]types.CurvePoint{partial0.Delta, partial1.Delta},
		big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	// submitCiphertext must precede submitPartialDecryption so the
	// proof's pi[5..6] can be bound against the on-chain C1.
	c.Assert(helpers.SubmitCiphertextAs(ctx,
		&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
		epochID, 1,
		combineOutput.CiphertextC1.X, combineOutput.CiphertextC1.Y,
		combineOutput.CiphertextC2.X, combineOutput.CiphertextC2.Y,
	), qt.IsNil)

	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}, epochID, 1, 1, combineOutput.CiphertextC1, combineOutput.CiphertextC2, partial0.DeltaHash, partial0.Proof, partial0.Input), qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, actor1, epochID, 2, 1, combineOutput.CiphertextC1, combineOutput.CiphertextC2, partial1.DeltaHash, partial1.Proof, partial1.Input), qt.IsNil)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.Manager.CombineDecryption(
		auth,
		epochID,
		[32]byte{}, // legacy per-epoch path: zero aid
		1,
		combineOutput.CombineHash,
		combineOutput.Plaintext,
		combineOutput.Transcript,
		combineOutput.Proof,
		combineOutput.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
}
