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

func TestGasProfiles(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	if !helpers.IsBenchmarkEnabled() {
		t.Skip("benchmark disabled — set RUN_BENCHMARKS=true to run gas-profile tests")
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
		LotteryAlphaBps:                 helpers.DefaultLotteryAlphaBps,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}

	epochID, createGas := createEpochForGasProfile(t, ctx, policy)
	// At committee size 1 this claimSlot call pays every one-time cost the
	// lottery can incur (seed resolve + committee snapshot + lottery check),
	// so the number is higher than the per-node amortised claimSlot cost in
	// BENCHMARKS.md — don't compare them directly.
	claimSlotGas := claimSlotForGasProfile(t, ctx, epochID, policy)
	contributionGas := submitContributionForGasProfile(t, ctx, epochID)
	finalizeGas := finalizeForGasProfile(t, ctx, epochID)
	partialDecryptGas, combineGas := decryptGasProfile(t, ctx, epochID)

	t.Logf(
		"gas profile create=%d claimSlot=%d contribution=%d finalize=%d partial_decrypt=%d combine=%d",
		createGas,
		claimSlotGas,
		contributionGas,
		finalizeGas,
		partialDecryptGas,
		combineGas,
	)

	// Generous ceilings — this is a benchmark, not a regression gate. BENCHMARKS.md
	// is the authoritative reference; these assertions only catch gross regressions
	// (e.g. an accidental O(N²) loop). Keep them loose so they tolerate normal
	// solc / contract tweaks without wasting CI cycles on a known-benchmark file.
	c.Assert(createGas < uint64(300_000), qt.IsTrue)
	c.Assert(claimSlotGas < uint64(250_000), qt.IsTrue)
	c.Assert(contributionGas < uint64(650_000), qt.IsTrue)
	c.Assert(finalizeGas < uint64(1_200_000), qt.IsTrue)
	c.Assert(partialDecryptGas < uint64(500_000), qt.IsTrue)
	c.Assert(combineGas < uint64(500_000), qt.IsTrue)
}

func createEpochForGasProfile(t *testing.T, ctx context.Context, policy types.EpochPolicy) ([12]byte, uint64) {
	t.Helper()
	c := qt.New(t)

	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	currentNonce, err := services.Manager.EpochNonce(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CreateEpoch(
		auth,
		policy.Threshold,
		policy.CommitteeSize,
		policy.MinValidContributions,
		policy.LotteryAlphaBps,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return helpers.ComputeRoundID(prefix, currentNonce+1), receipt.GasUsed
}

// claimSlotForGasProfile advances past the seed block and measures the gas for
// a single ClaimSlot call. At committee size 1 the only claimer pays every
// one-time cost: seed resolve (blockhash lookup), lottery check, committee
// snapshot. The BENCHMARKS.md averaged figure is lower because later claimers
// share those costs — do not compare this number directly.
func claimSlotForGasProfile(t *testing.T, ctx context.Context, epochID [12]byte, policy types.EpochPolicy) uint64 {
	t.Helper()
	c := qt.New(t)

	// Advance past the seed block so the lottery blockhash is available.
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.ClaimSlot(auth, epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func submitContributionForGasProfile(t *testing.T, ctx context.Context, epochID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	submission, err := helpers.BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, []*big.Int{big.NewInt(7)}, []uint16{1})
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitContribution(
		auth,
		epochID,
		1,
		submission.CommitmentsHash,
		submission.EncryptedSharesHash, submission.Transcript,
		submission.Proof,
		submission.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func finalizeForGasProfile(t *testing.T, ctx context.Context, epochID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	output, err := helpers.BuildFinalizeEpochOutput(ctx, epochID, 1, 1, []uint16{1}, [][]*big.Int{{big.NewInt(7)}})
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeEpoch(
		auth,
		epochID,
		output.AggregateCommitmentsHash,
		output.CollectivePublicKeyHash,
		output.ShareCommitmentHash,
		output.Transcript,
		output.Proof,
		output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

// decryptGasProfile registers an application, submits one ciphertext and
// drives the partial decryption + organizer share + combine for it, returning
// the partial-decrypt and combine gas.
//
// The three calls have to happen in this order: the ciphertext binds both
// proofs, and combineDecryption reverts with OrganizerShareMissing until the
// organizer has published Δ = sk_org·C1.
func decryptGasProfile(t *testing.T, ctx context.Context, epochID [12]byte) (partialGas, combineGas uint64) {
	t.Helper()
	c := qt.New(t)

	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)
	self := selfActor()
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, epochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	const ciphertextBase = 9
	partial, err := helpers.BuildPartialDecryptionSubmission(
		ctx, epochID, aid, 1, 1, big.NewInt(ciphertextBase), big.NewInt(7), big.NewInt(5),
	)
	c.Assert(err, qt.IsNil)

	output, err := helpers.BuildDecryptCombineOutput(
		ctx, epochID, aid, 1 /* ciphertextIndex */, 1,
		big.NewInt(ciphertextBase), skOrg,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, epochID, aid, output.CiphertextC1, output.CiphertextC2)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))

	partial.C2 = output.CiphertextC2
	partialGas, err = helpers.SubmitPartialDecryptionMeasured(ctx, services, self, epochID, aid, 1, assignedIdx, partial)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.PostOrganizerShare(ctx, self, services.AppManager, epochID, aid, assignedIdx, output), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CombineDecryption(
		auth, epochID, aid, assignedIdx,
		output.CombineHash, output.Plaintext,
		output.Transcript, output.Proof, output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return partialGas, receipt.GasUsed
}
