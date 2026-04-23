package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestSecretDisclosureHappyPath(t *testing.T) {
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
		FinalizeNotBeforeBlock:    head + 51,
		DisclosureAllowed:         true,
	}
	coefficients := []*big.Int{big.NewInt(13)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	submission, err := helpers.BuildRevealShareSubmission(
		ctx,
		result.RoundID,
		1,
		coefficients[0],
		result.ShareCommitments[0],
	)
	c.Assert(err, qt.IsNil)

	disclosure, err := helpers.BuildRevealShareOutput(ctx, result.RoundID, 1, []uint16{1}, []*big.Int{coefficients[0]})
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitRevealedShare(auth, result.RoundID, 1, submission.ShareValue, submission.Proof, submission.Input)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.Manager.ReconstructSecret(
		auth,
		result.RoundID,
		disclosure.DisclosureHash,
		disclosure.ReconstructedSecretHash,
		disclosure.Transcript,
		disclosure.Proof,
		disclosure.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	record, err := helpers.WaitRevealedShare(ctx, services, result.RoundID, services.TxManager.Address())
	c.Assert(err, qt.IsNil)
	c.Assert(record.Accepted, qt.IsTrue)
	// shareHash is no longer persisted in storage; it lives in the RevealedShareSubmitted event.

	round, err := services.Contracts.GetRound(ctx, result.RoundID)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Status, qt.Equals, uint8(5))
	// reconstructedSecretHash is no longer persisted; it lives in the
	// SecretReconstructed event log.
}
