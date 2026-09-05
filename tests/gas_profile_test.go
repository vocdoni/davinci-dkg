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

// gasProfile is one full pass over every user-facing transaction of the
// protocol. These are the numbers published in BENCHMARKS.md.
type gasProfile struct {
	createEpoch        uint64
	claimSlot          uint64
	submitContribution uint64
	finalizeEpoch      uint64
	registerLocked     uint64
	registerAutomatic  uint64
	revealSecret       uint64
	submitCiphertext   uint64
	partialDecrypt     uint64
	combine            uint64
}

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

	epochID, profile := singleNodeGasProfile(t, ctx, policy)

	t.Logf(
		"gas profile epoch=%s create=%d claimSlot=%d contribution=%d finalize=%d "+
			"registerLocked=%d registerAutomatic=%d reveal=%d ciphertext=%d partial_decrypt=%d combine=%d",
		helpers.RoundIDToString(epochID),
		profile.createEpoch, profile.claimSlot, profile.submitContribution, profile.finalizeEpoch,
		profile.registerLocked, profile.registerAutomatic, profile.revealSecret,
		profile.submitCiphertext, profile.partialDecrypt, profile.combine,
	)

	// Generous ceilings — this is a benchmark, not a regression gate. BENCHMARKS.md
	// is the authoritative reference; these assertions only catch gross regressions
	// (e.g. an accidental O(N²) loop). Keep them loose so they tolerate normal
	// solc / contract tweaks without wasting CI cycles on a known-benchmark file.
	c.Assert(profile.createEpoch < 300_000, qt.IsTrue)
	// At committee size 1 this claimSlot call pays every one-time cost the
	// lottery can incur (seed resolve + committee snapshot + lottery check),
	// so the number is higher than the per-node amortised claimSlot cost in
	// BENCHMARKS.md — don't compare them directly.
	c.Assert(profile.claimSlot < 250_000, qt.IsTrue)
	c.Assert(profile.submitContribution < 3_000_000, qt.IsTrue)
	// finalizeEpoch now verifies the batched proof and stores 16 keys and 16
	// roots (docs/pool-keys-v4.md §13 estimates 2.1–2.8 M).
	c.Assert(profile.finalizeEpoch < 4_000_000, qt.IsTrue)
	c.Assert(profile.registerLocked < 700_000, qt.IsTrue)
	c.Assert(profile.registerAutomatic < 400_000, qt.IsTrue)
	c.Assert(profile.revealSecret < 300_000, qt.IsTrue)
	c.Assert(profile.submitCiphertext < 250_000, qt.IsTrue)
	c.Assert(profile.partialDecrypt < 600_000, qt.IsTrue)
	c.Assert(profile.combine < 500_000, qt.IsTrue)
}

// singleNodeGasProfile drives one committee-of-one epoch through every
// transaction and returns their gas usage.
func singleNodeGasProfile(t *testing.T, ctx context.Context, policy types.EpochPolicy) ([12]byte, gasProfile) {
	t.Helper()
	c := qt.New(t)
	self := selfActor()
	var profile gasProfile

	epochID, createGas := createRoundMeasured(t, ctx, policy)
	profile.createEpoch = createGas

	// Advance past the seed block so the lottery blockhash is available.
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	claimGas, err := helpers.ClaimSlotMeasured(ctx, services, self, epochID)
	c.Assert(err, qt.IsNil)
	profile.claimSlot = claimGas

	contributions := [][][]*big.Int{helpers.DealPoolCoefficients([]*big.Int{big.NewInt(7)})}
	submission, err := helpers.BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, contributions[0], []uint16{1})
	c.Assert(err, qt.IsNil)
	profile.submitContribution, err = helpers.SubmitContributionMeasured(ctx, services, self, epochID, 1, submission)
	c.Assert(err, qt.IsNil)

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	finalization, err := helpers.BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, contributions)
	c.Assert(err, qt.IsNil)
	profile.finalizeEpoch, err = helpers.FinalizeEpochMeasured(ctx, services, self, epochID, finalization)
	c.Assert(err, qt.IsNil)

	round := &helpers.FinalizedRoundResult{
		EpochID:            epochID,
		Threshold:          1,
		CommitteeSize:      1,
		ParticipantIndexes: []uint16{1},
		Contributions:      contributions,
		Finalize:           finalization,
	}

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

	// The decryption pass runs on the automatic application, which claimed
	// the second pool key.
	const secondKey uint8 = 1
	share := poolShare(c, round, secondKey, 1)
	const ciphertextBase = 9
	partial, err := helpers.BuildPartialDecryptionSubmission(
		ctx, epochID, autoAid, 1, 1, big.NewInt(ciphertextBase), share, big.NewInt(5), round.Shares(secondKey),
	)
	c.Assert(err, qt.IsNil)
	output, err := helpers.BuildDecryptCombineOutput(
		ctx, epochID, autoAid, 1, 1, big.NewInt(ciphertextBase), nil,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	assignedIdx, ciphertextGas, err := helpers.SubmitCiphertextMeasured(
		ctx, self, epochID, autoAid, output.CiphertextC1, output.CiphertextC2,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))
	profile.submitCiphertext = ciphertextGas

	partial.C2 = output.CiphertextC2
	profile.partialDecrypt, err = helpers.SubmitPartialDecryptionMeasured(
		ctx, services, self, epochID, autoAid, 1, assignedIdx, partial,
	)
	c.Assert(err, qt.IsNil)

	profile.combine, err = helpers.CombineDecryptionMeasured(ctx, services, self, epochID, autoAid, assignedIdx, output)
	c.Assert(err, qt.IsNil)

	return epochID, profile
}
