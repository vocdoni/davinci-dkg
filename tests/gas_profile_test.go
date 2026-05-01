package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
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
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		SeedDelay:                 helpers.DefaultSeedDelay,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		FinalizeNotBeforeBlock:    head + 51,
	}

	epochID, createGas := createEpochForGasProfile(t, ctx, policy)
	// At committee size 1 this claimSlot call pays every one-time cost the
	// lottery can incur (seed resolve + committee snapshot + lottery check),
	// so the number is higher than the per-node amortised claimSlot cost in
	// BENCHMARKS.md — don't compare them directly.
	claimSlotGas := claimSlotForGasProfile(t, ctx, epochID, policy)
	contributionGas := submitContributionForGasProfile(t, ctx, epochID)
	finalizeGas := finalizeForGasProfile(t, ctx, epochID)
	partialDecryptGas := submitPartialDecryptForGasProfile(t, ctx, epochID)
	combineGas := combineDecryptionForGasProfile(t, ctx, epochID)

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
		policy.SeedDelay,
		policy.RegistrationDeadlineBlock,
		policy.ContributionDeadlineBlock,
		policy.FinalizeNotBeforeBlock,
		helpers.ZeroDecryptionPolicy(),
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
	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)

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
		submission.EncryptedSharesHash,
		submission.Commitment0X,
		submission.Commitment0Y,
		submission.Transcript,
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

func submitPartialDecryptForGasProfile(t *testing.T, ctx context.Context, epochID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	// CIRCUITS_AUDIT2 #5: submitPartialDecryption requires the
	// ciphertext to exist on-chain so it can bind C1 into the proof.
	// Submit a small canonical ciphertext at index 1 first.
	c1 := group.Generator()
	c1.ScalarBaseMult(big.NewInt(9))
	c2 := group.Generator()
	c2.ScalarBaseMult(big.NewInt(11))
	c1Enc, c2Enc := group.Encode(c1), group.Encode(c2)
	c.Assert(helpers.SubmitCiphertextAs(ctx,
		&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
		epochID, 1, c1Enc.X, c1Enc.Y, c2Enc.X, c2Enc.Y,
	), qt.IsNil)

	output, err := helpers.BuildPartialDecryptionSubmission(ctx, epochID, 1, big.NewInt(9), big.NewInt(7), big.NewInt(5))
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitPartialDecryption(
		auth, epochID, [32]byte{}, 1, 1,
		c1Enc.X, c1Enc.Y, c2Enc.X, c2Enc.Y,
		output.DeltaHash, output.Proof, output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func combineDecryptionForGasProfile(t *testing.T, ctx context.Context, epochID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	deltaPoint := group.Generator()
	deltaPoint.ScalarBaseMult(big.NewInt(63))
	output, err := helpers.BuildDecryptCombineOutput(
		ctx,
		epochID,
		1,
		big.NewInt(9),
		[]uint16{1},
		[]types.CurvePoint{group.Encode(deltaPoint)},
		big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.SubmitCiphertextAs(ctx,
		&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
		epochID, 1,
		output.CiphertextC1.X, output.CiphertextC1.Y,
		output.CiphertextC2.X, output.CiphertextC2.Y,
	), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CombineDecryption(
		auth,
		epochID,
		[32]byte{}, // legacy per-epoch path: zero aid
		1,
		output.CombineHash,
		output.Plaintext,
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

