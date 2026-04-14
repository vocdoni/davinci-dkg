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

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         false,
	}
	coefficients := []*big.Int{big.NewInt(7)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)
	c.Assert(result.RoundID, qt.Not(qt.Equals), [12]byte{})
	c.Assert(result.Round.Status, qt.Equals, uint8(3))
	c.Assert(result.Round.ContributionCount, qt.Equals, uint16(1))
	// AggregateCommitmentsHash / CollectivePublicKeyHash are no longer persisted
	// in storage; they live in the RoundFinalized event.
}
