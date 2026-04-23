package tests

// TestGasProfilesMultiNode runs a complete DKG round for each committee size
// defined in benchSizes and logs the per-call gas cost of every protocol phase.
// This is the data source for BENCHMARKS.md.
//
// Run with:
//
//	RUN_INTEGRATION_TESTS=true go test -v -run TestGasProfilesMultiNode \
//	  -timeout 60m ./tests/...
//
// For MaxN=32, first change circuits/common/sizes.go and run `make circuits`.

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// benchSizes lists the committee sizes to benchmark.
// Threshold = ceil(2n/3).
// Adjust to match the current MaxN; sizes exceeding MaxN will be skipped.
var benchSizes = []int{4, 8, 12, 16, 20, 24, 28, 32}

func TestGasProfilesMultiNode(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	if !helpers.IsBenchmarkEnabled() {
		t.Skip("benchmark disabled — set RUN_BENCHMARKS=true to run the multi-size gas sweep")
	}

	maxN := contribution.MaxRecipients // compile-time MaxN

	type row struct {
		n                     int
		t                     int
		createGas             uint64
		claimSlotGas          uint64 // average per slot
		submitContributionGas uint64 // last submitted (warm storage)
		finalizeGas           uint64
		partialDecryptGas     uint64
		combineGas            uint64
	}

	var results []row

	for _, n := range benchSizes {
		if n > maxN {
			t.Logf("skipping n=%d (MaxN=%d)", n, maxN)
			continue
		}
		threshold := (2*n + 2) / 3 // ceil(2n/3)
		t.Logf("=== n=%d t=%d ===", n, threshold)

		r := benchmarkGasForN(t, n, threshold)
		results = append(results, row{
			n:                     n,
			t:                     threshold,
			createGas:             r.createGas,
			claimSlotGas:          r.claimSlotGas,
			submitContributionGas: r.submitContributionGas,
			finalizeGas:           r.finalizeGas,
			partialDecryptGas:     r.partialDecryptGas,
			combineGas:            r.combineGas,
		})
	}

	// Print a markdown table for easy copy-paste into BENCHMARKS.md.
	t.Log("\n\n=== GAS PROFILE RESULTS (MaxN=" + fmt.Sprintf("%d", maxN) + ") ===")
	t.Log("| n | t | createRound | claimSlot (avg) | submitContribution | finalizeRound | submitPartialDecryption | combineDecryption |")
	t.Log("|---|---|---|---|---|---|---|---|")
	for _, r := range results {
		t.Logf("| %d | %d | %d | %d | %d | %d | %d | %d |",
			r.n, r.t, r.createGas, r.claimSlotGas,
			r.submitContributionGas, r.finalizeGas,
			r.partialDecryptGas, r.combineGas,
		)
	}

	// Also print the compact form used in BENCHMARKS.md:
	t.Log("\n=== Compact (submitContribution | finalizeRound | submitPartialDecryption | combineDecryption) ===")
	var sb strings.Builder
	for _, r := range results {
		fmt.Fprintf(&sb, "| %d | %d | %d | %d | %d | %d |\n",
			r.n, r.t,
			r.submitContributionGas, r.finalizeGas,
			r.partialDecryptGas, r.combineGas,
		)
	}
	t.Log(sb.String())
}

type gasProfileResult struct {
	createGas             uint64
	claimSlotGas          uint64 // average across all claimSlot calls
	submitContributionGas uint64 // last contribution call (fully warm)
	finalizeGas           uint64
	partialDecryptGas     uint64
	combineGas            uint64
}

func benchmarkGasForN(t *testing.T, n, threshold int) gasProfileResult {
	t.Helper()
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	// ── 0. ensure all n actor BJJ keys are registered ──────────────────────
	actors := make([]*helpers.TestActor, n)
	for i := range n {
		a, err := services.Actor(i)
		c.Assert(err, qt.IsNil)
		actors[i] = a
		c.Assert(helpers.EnsureNodeKeyRegistered(ctx, services, a), qt.IsNil)
	}

	// ── 1. createRound ──────────────────────────────────────────────────────
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	policy := types.RoundPolicy{
		Threshold:                 uint16(threshold),
		CommitteeSize:             uint16(n),
		MinValidContributions:     uint16(threshold),
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		SeedDelay:                 helpers.DefaultSeedDelay,
		RegistrationDeadlineBlock: head + 50,
		ContributionDeadlineBlock: head + 200,
		FinalizeNotBeforeBlock:    head + 201,
		DisclosureAllowed:         false,
	}
	roundID, createGas := createRoundMeasured(t, ctx, policy)

	// ── 2. claimSlot for all n actors ───────────────────────────────────────
	c.Assert(helpers.MineBlocks(ctx, services, uint64(policy.SeedDelay)+1), qt.IsNil)
	var totalClaimGas uint64
	for _, actor := range actors {
		gas, err := helpers.ClaimSlotMeasured(ctx, services, actor, roundID)
		c.Assert(err, qt.IsNil)
		totalClaimGas += gas
	}
	avgClaimGas := totalClaimGas / uint64(n)

	// ── 3. build and submit n contributions ─────────────────────────────────
	recipientIndexes := make([]uint16, n)
	for i := range n {
		recipientIndexes[i] = uint16(i + 1)
	}
	coefficients := make([][]*big.Int, n)
	for i := range n {
		coefficients[i] = make([]*big.Int, threshold)
		for k := range threshold {
			coefficients[i][k] = big.NewInt(int64((i+1)*10 + k + 1))
		}
	}

	var lastContribGas uint64
	for i, actor := range actors {
		sub, err := helpers.BuildContributionSubmission(
			ctx, services, roundID,
			uint16(threshold), uint16(n), uint16(i+1),
			coefficients[i], recipientIndexes,
		)
		c.Assert(err, qt.IsNil)
		gas, err := helpers.SubmitContributionMeasured(ctx, services, actor, roundID, uint16(i+1), sub)
		c.Assert(err, qt.IsNil)
		lastContribGas = gas // keep the last one (warmest storage)
	}

	// ── 4. finalizeRound ────────────────────────────────────────────────────
	output, err := helpers.BuildFinalizeRoundOutput(ctx, roundID,
		uint16(threshold), uint16(n), recipientIndexes, coefficients)
	c.Assert(err, qt.IsNil)
	finalizeGas := finalizeRoundMeasured(t, ctx, roundID, output)

	// ── 5. submitPartialDecryption (first t actors) ─────────────────────────
	recoveredShares, err := helpers.RecoverParticipantShares(coefficients, recipientIndexes)
	c.Assert(err, qt.IsNil)

	ciphertextBase := big.NewInt(42) // arbitrary C1 scalar
	var lastPartialGas uint64
	partialActors := actors[:threshold]
	partials := make([]struct {
		delta types.CurvePoint
		idx   uint16
	}, threshold)
	for i, actor := range partialActors {
		partial, err := helpers.BuildPartialDecryptionSubmission(
			ctx, roundID, uint16(i+1), ciphertextBase, recoveredShares[i], big.NewInt(int64(i+100)),
		)
		c.Assert(err, qt.IsNil)
		gas, err := helpers.SubmitPartialDecryptionMeasured(ctx, services, actor, roundID, uint16(i+1), 1, partial)
		c.Assert(err, qt.IsNil)
		lastPartialGas = gas
		partials[i].delta = partial.Delta
		partials[i].idx = uint16(i + 1)
	}

	// ── 6. combineDecryption ────────────────────────────────────────────────
	idxs := make([]uint16, threshold)
	deltas := make([]types.CurvePoint, threshold)
	for i := range threshold {
		idxs[i] = partials[i].idx
		deltas[i] = partials[i].delta
	}

	// BuildDecryptCombineOutput constructs c2 = plaintext*G + Lagrange(deltas)
	// and proves consistency; any consistent plaintext scalar works here.
	plaintextScalar := big.NewInt(99)

	combineOut, err := helpers.BuildDecryptCombineOutput(
		ctx, roundID, uint16(threshold),
		ciphertextBase, idxs, deltas, plaintextScalar,
	)
	c.Assert(err, qt.IsNil)

	// The combine tx is now bound to an on-chain ciphertext; submit it first.
	c.Assert(helpers.SubmitCiphertextAs(ctx,
		&helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager},
		roundID, 1,
		combineOut.CiphertextC1.X, combineOut.CiphertextC1.Y,
		combineOut.CiphertextC2.X, combineOut.CiphertextC2.Y,
	), qt.IsNil)

	combineGas := combineMeasured(t, ctx, roundID, combineOut)

	t.Logf("n=%d t=%d create=%d claim_avg=%d contrib=%d finalize=%d pdecrypt=%d combine=%d",
		n, threshold, createGas, avgClaimGas, lastContribGas, finalizeGas, lastPartialGas, combineGas)

	return gasProfileResult{
		createGas:             createGas,
		claimSlotGas:          avgClaimGas,
		submitContributionGas: lastContribGas,
		finalizeGas:           finalizeGas,
		partialDecryptGas:     lastPartialGas,
		combineGas:            combineGas,
	}
}

// ── measurement helpers ──────────────────────────────────────────────────────

func createRoundMeasured(t *testing.T, ctx context.Context, policy types.RoundPolicy) ([12]byte, uint64) {
	t.Helper()
	c := qt.New(t)

	prefix, err := services.Manager.ROUNDPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	nonce, err := services.Manager.RoundNonce(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CreateRound(auth,
		policy.Threshold, policy.CommitteeSize, policy.MinValidContributions,
		policy.LotteryAlphaBps, policy.SeedDelay,
		policy.RegistrationDeadlineBlock, policy.ContributionDeadlineBlock,
		policy.FinalizeNotBeforeBlock,
		policy.DisclosureAllowed,
		helpers.ZeroDecryptionPolicy(),
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)
	return helpers.ComputeRoundID(prefix, nonce+1), receipt.GasUsed
}

func finalizeRoundMeasured(t *testing.T, ctx context.Context, roundID [12]byte, output *helpers.FinalizeRoundOutput) uint64 {
	t.Helper()
	c := qt.New(t)
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, roundID), qt.IsNil)
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.FinalizeRound(auth, roundID,
		output.AggregateCommitmentsHash, output.CollectivePublicKeyHash, output.ShareCommitmentHash,
		output.Transcript, output.Proof, output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)
	return receipt.GasUsed
}

func combineMeasured(t *testing.T, ctx context.Context, roundID [12]byte, output *helpers.DecryptCombineOutput) uint64 {
	t.Helper()
	c := qt.New(t)
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CombineDecryption(auth, roundID, 1,
		output.CombineHash, output.Plaintext,
		output.Transcript, output.Proof, output.Input,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)
	return receipt.GasUsed
}
