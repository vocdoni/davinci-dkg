package battery

import (
	"testing"

	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// TestFleetStatus is a read-only smoke test: it prints the deployment
// immutables, the registry size and the newest epoch so a battery run is
// self-describing, and checks the torsion-point construction the ciphertext
// adversary relies on.
func TestFleetStatus(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()

	head, err := f.head(ctx)
	if err != nil {
		t.Fatal(err)
	}
	active, err := f.Services.Registry.ActiveCount(f.callOpts(ctx))
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := f.epochNonce(ctx)
	if err != nil {
		t.Fatal(err)
	}
	nextStart, err := f.nextEpochStart(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("head=%d activeOperators=%d epochNonce=%d nextEpochStart=%d", head, active, nonce, nextStart)
	t.Logf("immutables: epochDuration=%d minThreshold=%d minCommitteeSize=%d maxAlphaBps=%d prefix=%#x",
		f.EpochDuration, f.MinThreshold, f.MinCommitteeSize, f.MaxAlphaBps, f.Prefix)
	t.Logf("contracts: manager=%s registry=%s appManager=%s",
		f.Services.Addresses.Manager.Hex(), f.Services.Addresses.Registry.Hex(), f.Services.Addresses.AppManager.Hex())

	for n := nonce; n > 0 && n+2 > nonce; n-- {
		id := f.epochID(n)
		e, err := f.epoch(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		committee, err := f.committee(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("epoch %x: status=%d t=%d n=%d mMin=%d alpha=%d start=%d seed=%d csDeadline=%d kaDeadline=%d "+
			"liveNotBefore=%d serviceEnd=%d claimed=%d contributions=%d ciphertexts=%d committee=%d",
			id, e.Status, e.Policy.Threshold, e.Policy.CommitteeSize, e.Policy.MinValidContributions,
			e.Policy.LotteryAlphaBps, e.StartBlock, e.SeedBlock, e.Policy.CommitteeSelectionDeadlineBlock,
			e.Policy.KeyAssemblyDeadlineBlock, e.Policy.LiveNotBeforeBlock, f.serviceEnd(e),
			e.ClaimedCount, e.ContributionCount, e.CiphertextCount, len(committee))
	}

	tp := torsionPoint()
	if !group.IsOnCurve(tp.X, tp.Y) || group.IsInPrimeSubgroup(tp.X, tp.Y) || group.IsIdentity(tp.X, tp.Y) {
		t.Fatalf("torsion point (0, -1) must be on-curve, non-identity and outside the prime subgroup")
	}
	honest, err := randomSubgroupPoint()
	if err != nil {
		t.Fatal(err)
	}
	mixed, err := addPoints(honest, tp)
	if err != nil {
		t.Fatal(err)
	}
	if !group.IsOnCurve(mixed.X, mixed.Y) || group.IsInPrimeSubgroup(mixed.X, mixed.Y) {
		t.Fatalf("honest + torsion must be on-curve and outside the prime subgroup")
	}
	record(t, Result{
		Step: "torsion-construction", Kind: "measure", Pass: true,
		Notes: "(0,-1) and P+(0,-1) are on-curve, outside the prime-order subgroup",
	})
}
