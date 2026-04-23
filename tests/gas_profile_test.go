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

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		SeedDelay:                 helpers.DefaultSeedDelay,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		FinalizeNotBeforeBlock:    head + 51,
		DisclosureAllowed:         true,
	}

	roundID, createGas := createRoundForGasProfile(t, ctx, policy)
	// At committee size 1 this claimSlot call pays every one-time cost the
	// lottery can incur (seed resolve + committee snapshot + lottery check),
	// so the number is higher than the per-node amortised claimSlot cost in
	// BENCHMARKS.md — don't compare them directly.
	claimSlotGas := claimSlotForGasProfile(t, ctx, roundID, policy)
	contributionGas := submitContributionForGasProfile(t, ctx, roundID)
	finalizeGas := finalizeForGasProfile(t, ctx, roundID)
	partialDecryptGas := submitPartialDecryptForGasProfile(t, ctx, roundID)
	combineGas := combineDecryptionForGasProfile(t, ctx, roundID)
	revealGas := revealShareForGasProfile(t, ctx, roundID)
	reconstructGas := reconstructSecretForGasProfile(t, ctx, roundID)

	t.Logf(
		"gas profile create=%d claimSlot=%d contribution=%d finalize=%d partial_decrypt=%d combine=%d reveal=%d reconstruct=%d",
		createGas,
		claimSlotGas,
		contributionGas,
		finalizeGas,
		partialDecryptGas,
		combineGas,
		revealGas,
		reconstructGas,
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
	c.Assert(revealGas < uint64(400_000), qt.IsTrue)
	c.Assert(reconstructGas < uint64(400_000), qt.IsTrue)
}

func createRoundForGasProfile(t *testing.T, ctx context.Context, policy types.RoundPolicy) ([12]byte, uint64) {
	t.Helper()
	c := qt.New(t)

	prefix, err := services.Manager.ROUNDPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	currentNonce, err := services.Manager.RoundNonce(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CreateRound(
		auth,
		policy.Threshold,
		policy.CommitteeSize,
		policy.MinValidContributions,
		policy.LotteryAlphaBps,
		policy.SeedDelay,
		policy.RegistrationDeadlineBlock,
		policy.ContributionDeadlineBlock,
		policy.FinalizeNotBeforeBlock,
		policy.DisclosureAllowed,
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
func claimSlotForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte, policy types.RoundPolicy) uint64 {
	t.Helper()
	c := qt.New(t)

	// Advance past the seed block so the lottery blockhash is available.
	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.ClaimSlot(auth, roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func submitContributionForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	submission, err := helpers.BuildContributionSubmission(ctx, services, roundID, 1, 1, 1, []*big.Int{big.NewInt(7)}, []uint16{1})
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitContribution(
		auth,
		roundID,
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

func finalizeForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	output, err := helpers.BuildFinalizeRoundOutput(ctx, roundID, 1, 1, []uint16{1}, [][]*big.Int{{big.NewInt(7)}})
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, roundID), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeRound(
		auth,
		roundID,
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

func submitPartialDecryptForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	output, err := helpers.BuildPartialDecryptionSubmission(ctx, roundID, 1, big.NewInt(9), big.NewInt(7), big.NewInt(5))
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitPartialDecryption(auth, roundID, 1, 1, output.DeltaHash, output.Proof, output.Input)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func combineDecryptionForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	deltaPoint := group.Generator()
	deltaPoint.ScalarBaseMult(big.NewInt(63))
	output, err := helpers.BuildDecryptCombineOutput(
		ctx,
		roundID,
		1,
		big.NewInt(9),
		[]uint16{1},
		[]types.CurvePoint{group.Encode(deltaPoint)},
		big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.SubmitCiphertextAs(ctx,
		&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
		roundID, 1,
		output.CiphertextC1.X, output.CiphertextC1.Y,
		output.CiphertextC2.X, output.CiphertextC2.Y,
	), qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CombineDecryption(
		auth,
		roundID,
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

func revealShareForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	finalizeOutput, err := helpers.BuildFinalizeRoundOutput(ctx, roundID, 1, 1, []uint16{1}, [][]*big.Int{{big.NewInt(7)}})
	c.Assert(err, qt.IsNil)
	output, err := helpers.BuildRevealShareSubmission(ctx, roundID, 1, big.NewInt(7), finalizeOutput.ShareCommitments[0])
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.SubmitRevealedShare(auth, roundID, 1, output.ShareValue, output.Proof, output.Input)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)

	return receipt.GasUsed
}

func reconstructSecretForGasProfile(t *testing.T, ctx context.Context, roundID [12]byte) uint64 {
	t.Helper()
	c := qt.New(t)

	output, err := helpers.BuildRevealShareOutput(ctx, roundID, 1, []uint16{1}, []*big.Int{big.NewInt(7)})
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.ReconstructSecret(
		auth,
		roundID,
		output.DisclosureHash,
		output.ReconstructedSecretHash,
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
