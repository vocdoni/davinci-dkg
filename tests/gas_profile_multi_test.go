package tests

// TestGasProfilesMultiNode runs a complete DKG epoch for each committee size
// defined in benchSizes and logs the per-call gas cost of every protocol phase.
// This is the data source for BENCHMARKS.md.
//
// Run with:
//
//	RUN_INTEGRATION_TESTS=true RUN_BENCHMARKS=true go test -v -run TestGasProfilesMultiNode \
//	  -timeout 60m ./tests/...
//
// For MaxN=32, first change circuits/common/sizes.go and run `make circuits`.

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
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
	// Runs in its own process: it registers up to 32 operators, after which
	// the harness's single actor no longer wins every lottery, so
	// TestGasProfiles (RUN_BENCHMARKS) must not share the process with it.
	if os.Getenv("RUN_BENCHMARKS_MULTI") != "true" {
		t.Skip("benchmark disabled — set RUN_BENCHMARKS_MULTI=true to run the multi-size gas sweep")
	}

	maxN := contribution.MaxRecipients // compile-time MaxN

	type row struct {
		n       int
		t       int
		profile gasProfile
	}

	var results []row

	for _, n := range benchSizes {
		if n > maxN {
			t.Logf("skipping n=%d (MaxN=%d)", n, maxN)
			continue
		}
		threshold := (2*n + 2) / 3 // ceil(2n/3)
		t.Logf("=== n=%d t=%d ===", n, threshold)
		results = append(results, row{n: n, t: threshold, profile: benchmarkGasForN(t, n, threshold)})
	}

	// Print a markdown table for easy copy-paste into BENCHMARKS.md.
	t.Log("\n\n=== GAS PROFILE RESULTS (MaxN=" + fmt.Sprintf("%d", maxN) + ") ===")
	t.Log("| n | t | createEpoch | claimSlot (avg) | submitContribution | finalizeEpoch | activatePoolKey |" +
		" registerApplication (locked) | registerApplication (automatic) | revealOrganizerSecret |" +
		" submitCiphertext | submitPartialDecryption | combineDecryption |")
	t.Log("|---|---|---|---|---|---|---|---|---|---|---|---|---|")
	for _, r := range results {
		p := r.profile
		t.Logf(
			"| %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d | %d |",
			r.n, r.t, p.createEpoch, p.claimSlot, p.submitContribution, p.finalizeEpoch, p.activatePoolKey,
			p.registerLocked, p.registerAutomatic, p.revealSecret, p.submitCiphertext, p.partialDecrypt, p.combine,
		)
	}

	// Also print the compact form used in BENCHMARKS.md: the calls whose cost
	// actually moves with the committee size.
	t.Log("\n=== Compact (submitContribution | activatePoolKey | submitPartialDecryption | combineDecryption) ===")
	var sb strings.Builder
	for _, r := range results {
		fmt.Fprintf(
			&sb, "| %d | %d | %d | %d | %d | %d |\n",
			r.n, r.t, r.profile.submitContribution, r.profile.activatePoolKey,
			r.profile.partialDecrypt, r.profile.combine,
		)
	}
	t.Log(sb.String())
}

func benchmarkGasForN(t *testing.T, n, threshold int) gasProfile {
	t.Helper()
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	var profile gasProfile
	self := selfActor()

	// ── 0. ensure all n actor BJJ keys are registered ──────────────────────
	actors := make([]*helpers.TestActor, n)
	for i := range n {
		a, err := services.Actor(i)
		c.Assert(err, qt.IsNil)
		actors[i] = a
		c.Assert(helpers.EnsureNodeKeyRegistered(ctx, services, a), qt.IsNil)
	}

	// ── 1. createEpoch ──────────────────────────────────────────────────────
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	policy := types.EpochPolicy{
		Threshold:                       uint16(threshold),
		CommitteeSize:                   uint16(n),
		MinValidContributions:           uint16(threshold),
		LotteryAlphaBps:                 helpers.DefaultLotteryAlphaBps,
		CommitteeSelectionDeadlineBlock: head + 50,
		KeyAssemblyDeadlineBlock:        head + 200,
		LiveNotBeforeBlock:              head + 201,
	}
	epochID, createGas := createRoundMeasured(t, ctx, policy)
	profile.createEpoch = createGas

	// ── 2. claimSlot for all n actors ───────────────────────────────────────
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	var totalClaimGas uint64
	for _, actor := range actors {
		gas, claimErr := helpers.ClaimSlotMeasured(ctx, services, actor, epochID)
		c.Assert(claimErr, qt.IsNil)
		totalClaimGas += gas
	}
	profile.claimSlot = totalClaimGas / uint64(n)

	// ── 3. build and submit n contributions ─────────────────────────────────
	participantIndexes := make([]uint16, n)
	contributions := make([][][]*big.Int, n)
	for i := range n {
		participantIndexes[i] = uint16(i + 1)
		base := make([]*big.Int, threshold)
		for k := range threshold {
			base[k] = big.NewInt(int64((i+1)*10 + k + 1))
		}
		contributions[i] = helpers.DealPoolCoefficients(base)
	}

	for i, actor := range actors {
		sub, subErr := helpers.BuildContributionSubmission(
			ctx, services, epochID,
			uint16(threshold), uint16(n), uint16(i+1),
			contributions[i], participantIndexes,
		)
		c.Assert(subErr, qt.IsNil)
		gas, gasErr := helpers.SubmitContributionMeasured(ctx, services, actor, epochID, uint16(i+1), sub)
		c.Assert(gasErr, qt.IsNil)
		profile.submitContribution = gas // keep the last one (warmest storage)
	}

	// ── 4. finalizeEpoch (proof-less) ───────────────────────────────────────
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	profile.finalizeEpoch, err = helpers.FinalizeEpochMeasured(ctx, services, self, epochID)
	c.Assert(err, qt.IsNil)

	// ── 5. activatePoolKey, twice: the pool is claimed one key per app ──────
	round := &helpers.FinalizedRoundResult{
		EpochID:            epochID,
		Threshold:          uint16(threshold),
		CommitteeSize:      uint16(n),
		ParticipantIndexes: participantIndexes,
		Contributions:      contributions,
		Activations:        map[uint8]*helpers.PoolKeyActivation{},
	}
	first, err := helpers.BuildPoolKeyActivation(
		ctx, epochID, uint16(threshold), uint16(n), participantIndexes, contributions, 0,
	)
	c.Assert(err, qt.IsNil)
	profile.activatePoolKey, err = helpers.ActivatePoolKeyMeasured(ctx, services, self, epochID, first)
	c.Assert(err, qt.IsNil)
	round.Activations[0] = first
	second, err := helpers.ActivateRoundPoolKey(ctx, services, round, 1)
	c.Assert(err, qt.IsNil)

	// ── 6. registerApplication in both modes, plus the reveal ───────────────
	lockedAid, skOrg := randomAid(c), randomOrganizerSecret(c)
	profile.registerLocked, err = helpers.RegisterApplicationMeasured(
		ctx, services, self, services.AppManager, epochID, lockedAid, skOrg, golangtypes.DKGTypesAppPolicy{},
	)
	c.Assert(err, qt.IsNil)
	autoAid := randomAid(c)
	profile.registerAutomatic, err = helpers.RegisterApplicationMeasured(
		ctx, services, self, services.AppManager, epochID, autoAid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	)
	c.Assert(err, qt.IsNil)
	profile.revealSecret, err = helpers.RevealOrganizerSecretMeasured(
		ctx, services, self, services.AppManager, epochID, lockedAid, skOrg,
	)
	c.Assert(err, qt.IsNil)

	// ── 7. the decryption pass under the automatic application ──────────────
	shares, err := helpers.RecoverParticipantShares(contributions, second.KeyIndex, participantIndexes)
	c.Assert(err, qt.IsNil)

	ciphertextBase := big.NewInt(42) // arbitrary C1 scalar
	partials := make([]*helpers.PartialDecryptionSubmission, threshold)
	idxs := make([]uint16, threshold)
	deltas := make([]types.CurvePoint, threshold)
	for i := range threshold {
		partial, partialErr := helpers.BuildPartialDecryptionSubmission(
			ctx, epochID, autoAid, 1, uint16(i+1), ciphertextBase, shares[i], big.NewInt(int64(i+100)), second.Shares,
		)
		c.Assert(partialErr, qt.IsNil)
		partials[i] = partial
		idxs[i] = uint16(i + 1)
		deltas[i] = partial.Delta
	}

	// BuildDecryptCombineOutput constructs c2 = plaintext·G + Lagrange(deltas)
	// and proves consistency; any plaintext scalar works.
	plaintextScalar := big.NewInt(99)
	combineOut, err := helpers.BuildDecryptCombineOutput(
		ctx, epochID, autoAid, 1, uint16(threshold), ciphertextBase, nil, idxs, deltas, plaintextScalar,
	)
	c.Assert(err, qt.IsNil)

	// Both submitPartialDecryption and combineDecryption bind to the on-chain
	// ciphertext, so it has to land first.
	assignedIdx, ciphertextGas, err := helpers.SubmitCiphertextMeasured(
		ctx, self, epochID, autoAid, combineOut.CiphertextC1, combineOut.CiphertextC2,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))
	profile.submitCiphertext = ciphertextGas

	for i := range threshold {
		partials[i].C2 = combineOut.CiphertextC2
		gas, gasErr := helpers.SubmitPartialDecryptionMeasured(
			ctx, services, actors[i], epochID, autoAid, uint16(i+1), assignedIdx, partials[i],
		)
		c.Assert(gasErr, qt.IsNil)
		profile.partialDecrypt = gas
	}

	profile.combine, err = helpers.CombineDecryptionMeasured(ctx, services, self, epochID, autoAid, assignedIdx, combineOut)
	c.Assert(err, qt.IsNil)

	t.Logf("n=%d t=%d create=%d claim_avg=%d contrib=%d finalize=%d activate=%d "+
		"reg_locked=%d reg_auto=%d reveal=%d ciphertext=%d pdecrypt=%d combine=%d",
		n, threshold, profile.createEpoch, profile.claimSlot, profile.submitContribution, profile.finalizeEpoch,
		profile.activatePoolKey, profile.registerLocked, profile.registerAutomatic, profile.revealSecret,
		profile.submitCiphertext, profile.partialDecrypt, profile.combine)

	return profile
}

// createRoundMeasured creates an epoch and returns its id plus the gas used.
func createRoundMeasured(t *testing.T, ctx context.Context, policy types.EpochPolicy) ([12]byte, uint64) {
	t.Helper()
	c := qt.New(t)

	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	nonce, err := services.Manager.EpochNonce(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)

	// createEpoch is cadence-gated; earlier tests may have just created one.
	nextStart, err := services.Manager.NextEpochStartBlock(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	if head < nextStart {
		c.Assert(helpers.MineBlocks(ctx, services, nextStart-head), qt.IsNil)
	}

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CreateEpoch(
		auth,
		policy.Threshold, policy.CommitteeSize, policy.MinValidContributions,
		policy.LotteryAlphaBps,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	c.Assert(err, qt.IsNil)
	return helpers.ComputeRoundID(prefix, nonce+1), receipt.GasUsed
}
