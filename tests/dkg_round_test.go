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

	poolKeyHash, err := services.Manager.GetPoolKeyVerifierVKeyHash(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	c.Assert(common.Hash(poolKeyHash), qt.Not(qt.Equals), common.Hash{})
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

	// finalizeEpoch carries no proof; the key material only lands with the
	// per-key activations.
	status, err := services.Manager.GetPoolStatus(services.CallOpts(ctx), result.EpochID)
	c.Assert(err, qt.IsNil)
	c.Assert(status.Activated, qt.Equals, uint8(1), qt.Commentf("pool key 0 is activated"))
	c.Assert(status.NextIndex, qt.Equals, uint8(0), qt.Commentf("nothing claimed it yet"))

	x, y, err := services.Manager.GetPoolKey(services.CallOpts(ctx), result.EpochID, 0)
	c.Assert(err, qt.IsNil)
	activation := result.Activation(0)
	c.Assert(x.Cmp(activation.PoolKey.X), qt.Equals, 0)
	c.Assert(y.Cmp(activation.PoolKey.Y), qt.Equals, 0)

	// Key 1 was never proven: reading it reverts with PoolKeyNotActive.
	_, _, err = services.Manager.GetPoolKey(services.CallOpts(ctx), result.EpochID, 1)
	c.Assert(err, qt.IsNotNil)
}
