package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/finalizer"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestFinalizerEventLogPath exercises finalizer.BuildAndSubmit end-to-end
// using the new ContributionSubmitted event-log scan that replaced the prior
// 2000-block serial BlockByNumber walk.
//
// The previous implementation could spend 5–10 minutes on a public RPC
// while the node was inside the auto-finalize stagger window — the
// observable symptom was rounds that should finalize within ~3 blocks
// taking ~50 blocks instead. This test guards the fast path: if a future
// change accidentally reverts to a per-block scan it will still pass
// against Anvil (no observable latency locally), but a regression that
// breaks the event filter / calldata parse will fail here.
func TestFinalizerEventLogPath(t *testing.T) {
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
		FinalizeNotBeforeBlock:    head + 51,
		DisclosureAllowed:         false,
	}

	roundID, err := helpers.CreateRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, roundID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, roundID), qt.IsNil)

	round, err := helpers.WaitRoundStatus(ctx, services, roundID, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Policy.CommitteeSize, qt.Equals, uint16(3))

	committee := []common.Address{
		services.TxManager.Address(),
		actor1.Address(),
		actor2.Address(),
	}

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
		helpers.SubmitContributionAs(ctx, &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
			roundID, 1, submission0.CommitmentsHash, submission0.EncryptedSharesHash, submission0.Commitment0X, submission0.Commitment0Y, submission0.Transcript, submission0.Proof, submission0.Input),
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

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, roundID), qt.IsNil)

	res, err := finalizer.BuildAndSubmit(ctx, services.Contracts, services.Manager, services.TxManager, roundID, 2, 3, committee)
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.IsNotNil)
	c.Assert(len(res.ShareCommitments), qt.Equals, 3)

	finalized, err := services.Contracts.GetRound(ctx, roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(finalized.Status, qt.Equals, uint8(3)) // Finalized
}
