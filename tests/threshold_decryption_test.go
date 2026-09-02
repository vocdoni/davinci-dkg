package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
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

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}
	coefficients := []*big.Int{big.NewInt(11)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)
	self := selfActor()
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, result.EpochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	partial, err := helpers.BuildPartialDecryptionSubmission(ctx, result.EpochID, aid, 1, 1, big.NewInt(9), coefficients[0], big.NewInt(5))
	c.Assert(err, qt.IsNil)

	combine, err := helpers.BuildDecryptCombineOutput(ctx, result.EpochID, aid, 1, 1, big.NewInt(9), skOrg,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3))
	c.Assert(err, qt.IsNil)

	// submitCiphertext must precede submitPartialDecryption: the
	// partial-decrypt verifier binds pi[4..5] to the on-chain C1.
	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, result.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))

	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, result.EpochID, aid, 1, assignedIdx,
		combine.CiphertextC1, combine.CiphertextC2, partial.DeltaHash, partial.Proof, partial.Input), qt.IsNil)

	// Without the organizer share the combine reverts with
	// OrganizerShareMissing: t partials alone are not enough.
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	if tx, err := services.Manager.CombineDecryption(
		auth, result.EpochID, aid, assignedIdx,
		combine.CombineHash, combine.Plaintext, combine.Transcript, combine.Proof, combine.Input,
	); err == nil {
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNotNil,
			qt.Commentf("combine must revert while the organizer share is missing"))
	} else {
		c.Assert(err.Error(), qt.Contains, "execution reverted")
	}

	c.Assert(helpers.PostOrganizerShare(ctx, self, services.AppManager, result.EpochID, aid, assignedIdx, combine), qt.IsNil)

	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CombineDecryption(
		auth,
		result.EpochID,
		aid,
		assignedIdx,
		combine.CombineHash,
		combine.Plaintext,
		combine.Transcript,
		combine.Proof,
		combine.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, result.EpochID, aid, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Completed, qt.IsTrue)
	c.Assert(record.Plaintext.String(), qt.Equals, "3")
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

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}
	coefficients := []*big.Int{big.NewInt(11)}

	result, err := helpers.CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	c.Assert(err, qt.IsNil)

	baseValues := []*big.Int{big.NewInt(9), big.NewInt(13)}
	plaintexts := []*big.Int{big.NewInt(3), big.NewInt(5)}

	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)
	self := selfActor()
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, result.EpochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	for i := range baseValues {
		ciphertextIndex := uint16(i + 1)

		partial, err := helpers.BuildPartialDecryptionSubmission(
			ctx,
			result.EpochID,
			aid,
			ciphertextIndex,
			1,
			baseValues[i],
			coefficients[0],
			big.NewInt(int64(5+i)),
		)
		c.Assert(err, qt.IsNil)

		combine, err := helpers.BuildDecryptCombineOutput(
			ctx,
			result.EpochID,
			aid,
			ciphertextIndex,
			1,
			baseValues[i],
			skOrg,
			[]uint16{1},
			[]types.CurvePoint{partial.Delta},
			plaintexts[i],
		)
		c.Assert(err, qt.IsNil)

		// submitCiphertext must precede submitPartialDecryption.
		assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, result.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
		c.Assert(err, qt.IsNil)
		c.Assert(assignedIdx, qt.Equals, ciphertextIndex)

		c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, result.EpochID, aid, 1, ciphertextIndex,
			combine.CiphertextC1, combine.CiphertextC2, partial.DeltaHash, partial.Proof, partial.Input), qt.IsNil)

		c.Assert(helpers.PostOrganizerShare(ctx, self, services.AppManager, result.EpochID, aid, ciphertextIndex, combine), qt.IsNil)

		auth, err := services.TxManager.NewTransactOpts(ctx)
		c.Assert(err, qt.IsNil)
		tx, err := services.Manager.CombineDecryption(
			auth,
			result.EpochID,
			aid,
			ciphertextIndex,
			combine.CombineHash,
			combine.Plaintext,
			combine.Transcript,
			combine.Proof,
			combine.Input,
		)
		c.Assert(err, qt.IsNil)
		c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

		record, err := helpers.WaitCombinedDecryption(ctx, services, result.EpochID, aid, ciphertextIndex)
		c.Assert(err, qt.IsNil)
		c.Assert(record.Completed, qt.IsTrue)
		c.Assert(record.Plaintext.String(), qt.Equals, plaintexts[i].String())
	}
}
