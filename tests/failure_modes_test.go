package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestContributionRejectsMalformedProof(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         false,
	}

	roundID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	submission, err := helpers.BuildContributionSubmission(ctx, services, roundID, 1, 1, 1, []*big.Int{big.NewInt(21)}, []uint16{1})
	c.Assert(err, qt.IsNil)
	submission.Proof = submission.Proof[:len(submission.Proof)-32]

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitContribution(
		auth,
		roundID,
		1,
		submission.CommitmentsHash,
		submission.EncryptedSharesHash,
		submission.Commitment0X,
		submission.Commitment0Y,
		submission.Transcript,
		submission.Proof,
		submission.Input,
	)
	if err == nil {
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNotNil)
	} else {
		c.Assert(err.Error(), qt.Contains, "execution reverted")
	}
}

func TestFinalizeRejectsMissingContribution(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         false,
	}

	roundID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	output, err := helpers.BuildFinalizeRoundOutput(ctx, roundID, 1, 1, []uint16{1}, [][]*big.Int{{big.NewInt(1)}})
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeRound(
		auth,
		roundID,
		output.AggregateCommitmentsHash,
		output.CollectivePublicKeyHash,
		output.ShareCommitmentHash,
		output.Transcript,
		output.Proof,
		output.Input,
	)
	if err == nil {
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNotNil)
	} else {
		c.Assert(err.Error(), qt.Contains, "execution reverted")
	}
}

func TestPartialDecryptRejectsMalformedProof(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         false,
	}
	coefficients := []*big.Int{big.NewInt(17)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	partial, err := helpers.BuildPartialDecryptionSubmission(ctx, result.RoundID, 1, big.NewInt(3), coefficients[0], big.NewInt(4))
	c.Assert(err, qt.IsNil)
	partial.Input = partial.Input[:len(partial.Input)-32]

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitPartialDecryption(
		auth,
		result.RoundID,
		1,
		1,
		partial.DeltaHash,
		partial.Proof,
		partial.Input,
	)
	if err == nil {
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNotNil)
	} else {
		c.Assert(err.Error(), qt.Contains, "execution reverted")
	}
}

func TestRoundCanFinalizeWithMissingContributorWhenPolicyPermits(t *testing.T) {
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
		DisclosureAllowed:         false,
	}

	roundID, err := helpers.CreateRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, roundID), qt.IsNil)

	submission0, err := helpers.BuildContributionSubmission(ctx, services, roundID, 2, 3, 1, []*big.Int{big.NewInt(3), big.NewInt(1)}, []uint16{1, 2, 3})
	c.Assert(err, qt.IsNil)
	submission1, err := helpers.BuildContributionSubmission(ctx, services, roundID, 2, 3, 2, []*big.Int{big.NewInt(5), big.NewInt(2)}, []uint16{1, 2, 3})
	c.Assert(err, qt.IsNil)

	selfActor := &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}
	c.Assert(helpers.SubmitContributionAs(ctx, selfActor, roundID, 1, submission0.CommitmentsHash, submission0.EncryptedSharesHash, submission0.Commitment0X, submission0.Commitment0Y, submission0.Transcript, submission0.Proof, submission0.Input), qt.IsNil)
	c.Assert(helpers.SubmitContributionAs(ctx, actor1, roundID, 2, submission1.CommitmentsHash, submission1.EncryptedSharesHash, submission1.Commitment0X, submission1.Commitment0Y, submission1.Transcript, submission1.Proof, submission1.Input), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	output, err := helpers.BuildFinalizeRoundOutput(
		ctx,
		roundID,
		2,
		3,
		[]uint16{1, 2},
		[][]*big.Int{{big.NewInt(3), big.NewInt(1)}, {big.NewInt(5), big.NewInt(2)}},
	)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeRound(
		auth,
		roundID,
		output.AggregateCommitmentsHash,
		output.CollectivePublicKeyHash,
		output.ShareCommitmentHash,
		output.Transcript,
		output.Proof,
		output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	round, err := services.Contracts.GetRound(ctx, roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Status, qt.Equals, uint8(3))
	c.Assert(round.ContributionCount, qt.Equals, uint16(2))
}
