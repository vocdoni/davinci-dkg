package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestThresholdDecryptionHappyPath(t *testing.T) {
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
	coefficients := []*big.Int{big.NewInt(11)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	partial, err := helpers.BuildPartialDecryptionSubmission(ctx, result.RoundID, 1, big.NewInt(9), coefficients[0], big.NewInt(5))
	c.Assert(err, qt.IsNil)

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
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	combine, err := helpers.BuildDecryptCombineOutput(ctx, result.RoundID, 1, big.NewInt(9), []uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3))
	c.Assert(err, qt.IsNil)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.Manager.CombineDecryption(
		auth,
		result.RoundID,
		1,
		combine.CombineHash,
		combine.PlaintextHash,
		combine.Transcript,
		combine.Proof,
		combine.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, result.RoundID, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Completed, qt.IsTrue)
	// combineHash + plaintextHash are no longer persisted in storage; they live in the DecryptionCombined event.
}

func TestThresholdDecryptionSupportsMultipleCiphertextsPerRound(t *testing.T) {
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
	coefficients := []*big.Int{big.NewInt(11)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	baseValues := []*big.Int{big.NewInt(9), big.NewInt(13)}
	plaintexts := []*big.Int{big.NewInt(3), big.NewInt(5)}

	for i := range baseValues {
		ciphertextIndex := uint16(i + 1)

		partial, err := helpers.BuildPartialDecryptionSubmission(
			ctx,
			result.RoundID,
			1,
			baseValues[i],
			coefficients[0],
			big.NewInt(int64(5+i)),
		)
		c.Assert(err, qt.IsNil)

		auth, err := services.TxManager.NewTransactOpts(ctx)
		c.Assert(err, qt.IsNil)
		tx, err := services.Manager.SubmitPartialDecryption(
			auth,
			result.RoundID,
			1,
			ciphertextIndex,
			partial.DeltaHash,
			partial.Proof,
			partial.Input,
		)
		c.Assert(err, qt.IsNil)
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

		combine, err := helpers.BuildDecryptCombineOutput(
			ctx,
			result.RoundID,
			1,
			baseValues[i],
			[]uint16{1},
			[]types.CurvePoint{partial.Delta},
			plaintexts[i],
		)
		c.Assert(err, qt.IsNil)

		auth, err = services.TxManager.NewTransactOpts(ctx)
		c.Assert(err, qt.IsNil)
		tx, err = services.Manager.CombineDecryption(
			auth,
			result.RoundID,
			ciphertextIndex,
			combine.CombineHash,
			combine.PlaintextHash,
			combine.Transcript,
			combine.Proof,
			combine.Input,
		)
		c.Assert(err, qt.IsNil)
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

		record, err := helpers.WaitCombinedDecryption(ctx, services, result.RoundID, ciphertextIndex)
		c.Assert(err, qt.IsNil)
		c.Assert(record.Completed, qt.IsTrue)
		// combineHash + plaintextHash are no longer persisted in storage; they live in the DecryptionCombined event.
	}
}
