package tests

import (
	"context"
	"errors"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/finalizer"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestFinalizerEventLogPath exercises finalizer.BuildAndSubmitFinalize
// end-to-end over the ContributionSubmitted event-log scan that replaced the
// prior 2000-block serial BlockByNumber walk.
//
// The previous implementation could spend 5–10 minutes on a public RPC
// while the node was inside the auto-finalize stagger window — the
// observable symptom was epochs that should finalize within ~3 blocks
// taking ~50 blocks instead. This test guards the fast path: if a future
// change accidentally reverts to a per-block scan it will still pass
// against Anvil (no observable latency locally), but a regression that
// breaks the event filter / calldata parse will fail here.
//
// It also pins the pool-key reconstruction: the finalization is proven purely
// from the contributions' calldata, so a divergence between the compact
// transcript layout the node writes and the one it reads back shows up as a
// finalizeEpoch revert.
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

	policy := types.EpochPolicy{
		Threshold:                       2,
		CommitteeSize:                   3,
		MinValidContributions:           2,
		LotteryAlphaBps:                 helpers.DefaultLotteryAlphaBps,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}

	epochID, err := helpers.CreateEpoch(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, epochID), qt.IsNil)

	epoch, err := helpers.WaitEpochPhase(ctx, services, epochID, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Policy.CommitteeSize, qt.Equals, uint16(3))

	actors := []*helpers.TestActor{selfActor(), actor1, actor2}
	contributions := [][][]*big.Int{
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(5), big.NewInt(2)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(7), big.NewInt(4)}),
	}
	participantIndexes := []uint16{1, 2, 3}

	for i, actor := range actors {
		sub, subErr := helpers.BuildContributionSubmission(
			ctx, services, epochID, 2, 3, uint16(i+1), contributions[i], participantIndexes,
		)
		c.Assert(subErr, qt.IsNil)
		c.Assert(helpers.SubmitContributionAs(ctx, actor, epochID, uint16(i+1), sub), qt.IsNil)
	}

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)

	res, err := finalizer.BuildAndSubmitFinalize(ctx, services.Contracts, services.Manager, services.TxManager, epochID, nil)
	c.Assert(err, qt.IsNil)
	c.Assert(res, qt.IsNotNil)
	c.Assert(res.ParticipantIndexes, qt.DeepEquals, participantIndexes)
	c.Assert(res.PoolKeys, qt.HasLen, ccommon.MaxK)
	c.Assert(res.ShareCommitments, qt.HasLen, ccommon.MaxK)
	c.Assert(res.ShareCommitments[0], qt.HasLen, 3)
	finalized, err := services.Contracts.GetEpoch(ctx, epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(finalized.Status, qt.Equals, uint8(3)) // Live

	// The keys the finalizer reconstructed from calldata must be the ones the
	// local witness builder produces from the same coefficients, and the
	// contract must have stored exactly them.
	expected, err := helpers.BuildFinalizeSubmission(ctx, epochID, 2, 3, participantIndexes, contributions)
	c.Assert(err, qt.IsNil)
	for key := range uint8(ccommon.MaxK) {
		x, y, err := services.Manager.GetPoolKey(services.CallOpts(ctx), epochID, key)
		c.Assert(err, qt.IsNil)
		c.Assert(x.Cmp(expected.PoolKey(key).X), qt.Equals, 0, qt.Commentf("key %d", key))
		c.Assert(y.Cmp(expected.PoolKey(key).Y), qt.Equals, 0)
		c.Assert(res.PoolKeys[key].X.Cmp(x), qt.Equals, 0)

		root, err := services.Manager.GetPoolShareRoot(services.CallOpts(ctx), epochID, key)
		c.Assert(err, qt.IsNil)
		c.Assert(root, qt.Equals, expected.ShareTree(key).Root(), qt.Commentf("key %d", key))
		c.Assert(res.ShareRoots[key], qt.Equals, root)
	}

	// A second finalization is the benign race the node treats as a win for
	// someone else.
	_, err = finalizer.BuildAndSubmitFinalize(ctx, services.Contracts, services.Manager, services.TxManager, epochID, nil)
	c.Assert(errors.Is(err, finalizer.ErrAlreadyLive), qt.IsTrue, qt.Commentf("got %v", err))
}
