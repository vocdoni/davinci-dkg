package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestContractsSmoke(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	contributionHash, err := services.Contracts.GetContributionVerifierVKeyHash(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(contributionHash, qt.Not(qt.Equals), common.Hash{})

	partialHash, err := services.Contracts.GetPartialDecryptVerifierVKeyHash(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(partialHash, qt.Not(qt.Equals), common.Hash{})

	finalizeHash, err := services.Manager.GetFinalizeVerifierVKeyHash(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	c.Assert(common.Hash(finalizeHash), qt.Not(qt.Equals), common.Hash{})
}

func TestDKGRoundHappyPath(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}
	coefficients := []*big.Int{big.NewInt(7)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)
	c.Assert(result.EpochID, qt.Not(qt.Equals), [12]byte{})
	c.Assert(result.Epoch.Status, qt.Equals, uint8(3)) // Live
	c.Assert(result.Epoch.ContributionCount, qt.Equals, uint16(1))

	// finalizeEpoch carries the batched proof: every pool key and share root
	// is stored the moment the epoch is Live, nothing is claimed yet.
	next, err := services.Manager.GetPoolStatus(services.CallOpts(ctx), result.EpochID)
	c.Assert(err, qt.IsNil)
	c.Assert(next, qt.Equals, uint8(0), qt.Commentf("nothing claimed yet"))

	for key := range uint8(ccommon.MaxK) {
		x, y, err := services.Manager.GetPoolKey(services.CallOpts(ctx), result.EpochID, key)
		c.Assert(err, qt.IsNil, qt.Commentf("key %d", key))
		c.Assert(x.Cmp(result.PoolKey(key).X), qt.Equals, 0, qt.Commentf("key %d", key))
		c.Assert(y.Cmp(result.PoolKey(key).Y), qt.Equals, 0)
		root, err := services.Manager.GetPoolShareRoot(services.CallOpts(ctx), result.EpochID, key)
		c.Assert(err, qt.IsNil)
		c.Assert(root, qt.Equals, result.Shares(key).Root(), qt.Commentf("key %d", key))
	}
	// The pool ends at MaxK.
	_, _, err = services.Manager.GetPoolKey(services.CallOpts(ctx), result.EpochID, uint8(ccommon.MaxK))
	c.Assert(err, qt.IsNotNil)
}
