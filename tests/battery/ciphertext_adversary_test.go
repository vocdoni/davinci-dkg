package battery

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// ciphertextAdversaryKeys is how many pool keys the scenario claims in its
// epoch: application A, the three policy probes (cap, future window, past
// window), the honest neighbour every poison is measured against, and
// application B of the cross-application copy.
const ciphertextAdversaryKeys = 6

// TestCiphertextAdversary attacks submitCiphertext and the nodes' ciphertext
// validation: policy reverts (submitter, aid, cap, window), malformed points,
// then the three ciphertexts the contract accepts by design but the nodes
// must handle safely — a cofactor-subgroup point (no partial may ever be
// published), an undecryptable C2 (every combiner burns a BSGS to 2^50) and
// a verbatim copy of another application's ciphertext. For each poison the
// latency of the honest ciphertexts submitted right after it is measured
// against the one submitted right before and reported explicitly.
func TestCiphertextAdversary(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	combineWait := envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240)
	observe := envUint64("BATTERY_POISON_OBSERVE_BLOCKS", 45)

	epochID, epoch, err := f.waitLiveEpoch(ctx, t, envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90), ciphertextAdversaryKeys)
	if err != nil {
		t.Fatal(err)
	}
	orgA, err := f.newActor(ctx, "ct-organizer-a")
	if err != nil {
		t.Fatal(err)
	}
	orgB, err := f.newActor(ctx, "ct-organizer-b")
	if err != nil {
		t.Fatal(err)
	}
	stranger, err := f.newActor(ctx, "ct-stranger")
	if err != nil {
		t.Fatal(err)
	}
	appA, out, err := f.registerApplication(ctx, orgA, epochID, automaticPolicy())
	if !expectOK(t, "register-a", "registerApplication", out, err, "") {
		t.FailNow()
	}
	// Neighbour probes live in a fresh application: an undecryptable
	// ciphertext taints its own application for the epoch (the nodes stop
	// serving it by design), so the same aid cannot measure fleet latency.
	// One neighbour serves every poison — it only ever sees honest
	// ciphertexts, and each registration costs one of the epoch's pool keys.
	neighbour, out, err := f.registerApplication(ctx, orgA, epochID, automaticPolicy())
	if !expectOK(t, "register-neighbour", "registerApplication", out, err, "") {
		t.FailNow()
	}

	ciphertextPolicyReverts(ctx, t, f, epochID, appA, stranger)
	ciphertextPointReverts(ctx, t, f, appA)

	baseline := honestRoundTrip(ctx, t, f, epoch, appA, combineWait, "baseline")

	// Cofactor point: accepted on chain, must never receive a partial.
	torsionC1, torsionC2, err := torsionCiphertext(appA)
	if err != nil {
		t.Fatal(err)
	}
	dosProbe(ctx, t, f, epoch, appA, neighbour, "torsion", torsionC1, torsionC2, combineWait, observe, false)

	// Undecryptable C2: valid partials, valid share, dlog out of range.
	c1, _, err := appA.encrypt(big.NewInt(13))
	if err != nil {
		t.Fatal(err)
	}
	randomC2, err := randomSubgroupPoint()
	if err != nil {
		t.Fatal(err)
	}
	dosProbe(ctx, t, f, epoch, appA, neighbour, "random-c2", c1, randomC2, combineWait, observe, true)

	// Verbatim copy of A's baseline into B. A and B hold different pool
	// keys, so B's partials decrypt it to m·G + r·(P_A − P_B): garbage that
	// never combines.
	appB, out, err := f.registerApplication(ctx, orgB, epochID, automaticPolicy())
	if !expectOK(t, "cross-app/register-b", "registerApplication", out, err, "") {
		return
	}
	dosProbe(ctx, t, f, epoch, appB, neighbour, "cross-app", baseline.C1, baseline.C2, combineWait, observe, true)
}

// ciphertextPolicyReverts covers the DKGAppManager policy gate.
func ciphertextPolicyReverts(ctx context.Context, t *testing.T, f *Fleet, epochID [12]byte, appA *application, stranger *actor) {
	t.Helper()
	c1, c2, err := appA.encrypt(big.NewInt(7))
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = f.submitCiphertext(ctx, stranger, epochID, appA.Aid, c1, c2)
	expectRevert(t, "policy/not-authorized-submitter", err, "NotOwner")

	unknown, err := randomAid()
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = f.submitCiphertext(ctx, appA.Organizer, epochID, unknown, c1, c2)
	expectRevert(t, "policy/unregistered-aid", err, "InvalidApplication")

	capped, out, err := f.registerApplication(ctx, appA.Organizer, epochID, automaticPolicyWith(golangtypes.DKGTypesAppPolicy{MaxCiphertexts: 1}))
	if expectOK(t, "policy/register-capped", "registerApplication", out, err, "maxCiphertexts=1") {
		_, out, err := f.submitCiphertext(ctx, capped.Organizer, epochID, capped.Aid, c1, c2)
		expectOK(t, "policy/capped-first", "submitCiphertext", out, err, "")
		_, _, err = f.submitCiphertext(ctx, capped.Organizer, epochID, capped.Aid, c1, c2)
		expectRevert(t, "policy/capped-second", err, "DecryptionLimitReached")
	}

	head, err := f.head(ctx)
	if err != nil {
		t.Fatal(err)
	}
	futurePolicy := automaticPolicyWith(golangtypes.DKGTypesAppPolicy{NotBeforeBlock: head + 100_000})
	future, out, err := f.registerApplication(ctx, appA.Organizer, epochID, futurePolicy)
	if expectOK(t, "policy/register-future-window", "registerApplication", out, err, "") {
		_, _, err = f.submitCiphertext(ctx, future.Organizer, epochID, future.Aid, c1, c2)
		expectRevert(t, "policy/window-not-yet-open", err, "DecryptionNotYetAllowed")
	}
	past, out, err := f.registerApplication(ctx, appA.Organizer, epochID, automaticPolicyWith(golangtypes.DKGTypesAppPolicy{NotAfterBlock: 1}))
	if expectOK(t, "policy/register-past-window", "registerApplication", out, err, "") {
		_, _, err = f.submitCiphertext(ctx, past.Organizer, epochID, past.Aid, c1, c2)
		expectRevert(t, "policy/window-expired", err, "DecryptionExpired")
	}
}

// ciphertextPointReverts covers the on-chain point checks: canonical,
// on-curve, non-identity. Every rejection is InvalidCiphertext.
func ciphertextPointReverts(ctx context.Context, t *testing.T, f *Fleet, app *application) {
	t.Helper()
	honest, err := randomSubgroupPoint()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		step   string
		c1, c2 types.CurvePoint
	}{
		{"point/off-curve-c1", types.CurvePoint{X: big.NewInt(1), Y: big.NewInt(2)}, honest},
		{"point/off-curve-c2", honest, types.CurvePoint{X: big.NewInt(3), Y: big.NewInt(4)}},
		{"point/identity-c1", types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}, honest},
		{"point/non-canonical-c1", types.CurvePoint{X: group.BaseField(), Y: honest.Y}, honest},
	}
	for _, c := range cases {
		_, _, err := f.submitCiphertext(ctx, app.Organizer, app.Epoch, app.Aid, c.c1, c.c2)
		expectRevert(t, c.step, err, "InvalidCiphertext")
	}
}

// torsionCiphertext is an honest ciphertext whose C1 was pushed out of the
// prime subgroup by adding the order-2 point (0, -1).
func torsionCiphertext(app *application) (types.CurvePoint, types.CurvePoint, error) {
	c1, c2, err := app.encrypt(big.NewInt(11))
	if err != nil {
		return c1, c2, err
	}
	mixed, err := addPoints(c1, torsionPoint())
	return mixed, c2, err
}

// honestRoundTrip submits one honest ciphertext and waits for the fleet to
// combine it. The application is automatic, so there is nothing for the
// organizer to release. Returns the slot (with its combine event attached
// through the report).
func honestRoundTrip(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, wait uint64, step string,
) slot {
	t.Helper()
	s := submitSlots(ctx, t, f, app, 1, step)[0]
	ev := assertCombines(ctx, t, f, epoch, app, s, s.SubmitBlock, wait, step+"/combine")
	if ev != nil {
		s.combine = ev
	}
	return s
}

// dosProbe is the before → poison → after → late sequence shared by the
// three poisons: an honest ciphertext right before (latency reference), the
// poison, an honest ciphertext right after in the untainted `neighbour`
// application, then — once the combiners have had `observe` blocks to burn
// on the poison — one more honest ciphertext. The poison's own partial /
// combine status is reported early and late.
func dosProbe(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app, neighbour *application,
	name string, c1, c2 types.CurvePoint, wait, observe uint64, partialsExpected bool,
) {
	t.Helper()
	before := honestRoundTrip(ctx, t, f, epoch, app, wait, name+"/before")

	idx, out, err := f.submitCiphertext(ctx, app.Organizer, app.Epoch, app.Aid, c1, c2)
	if !expectOK(t, name+"/poison/submit", "submitCiphertext", out, err, "accepted by the contract by design") {
		return
	}
	poison := slot{Idx: idx, C1: c1, C2: c2, SubmitBlock: out.Block}

	after := honestRoundTrip(ctx, t, f, epoch, neighbour, wait, name+"/after")
	latencyDelta(t, name+"/latency-after", before, after)
	poisonStatus(ctx, t, f, epoch, app, poison, name+"/poison/status-early", partialsExpected)

	if _, err := f.waitBlock(ctx, poison.SubmitBlock+observe); err != nil {
		t.Fatal(err)
	}
	late := honestRoundTrip(ctx, t, f, epoch, neighbour, wait, name+"/late")
	latencyDelta(t, name+"/latency-late", before, late)
	poisonStatus(ctx, t, f, epoch, app, poison, name+"/poison/status-late", partialsExpected)

	if partialsExpected {
		// Same application after the poison: the designed outcome is no
		// combine (tainted); a combine means the nodes had not failed the
		// search yet. Reported, not judged.
		s := submitSlots(ctx, t, f, app, 1, name+"/tainted")[0]
		taintWait := envUint64("BATTERY_NO_COMBINE_WAIT_BLOCKS", 40)
		if _, err := f.waitBlock(ctx, s.SubmitBlock+taintWait); err != nil {
			t.Fatal(err)
		}
		rec, err := f.Services.Contracts.GetCombinedDecryption(ctx, app.Epoch, app.Aid, s.Idx)
		if err != nil {
			t.Fatal(err)
		}
		partials, err := f.partials(ctx, app.Epoch, app.Aid, s.Idx, scanFrom(epoch))
		if err != nil {
			t.Fatal(err)
		}
		record(t, Result{
			Step: name + "/tainted/status", Kind: "measure", Pass: true, Block: s.SubmitBlock,
			Notes: fmt.Sprintf("same application after an undecryptable ciphertext: combined=%v partials=%d after %d blocks "+
				"(nodes taint the application for the epoch by design)", rec.Completed, len(partials), taintWait),
		})
	}
}

// latencyDelta records the submit→combine latency of an honest ciphertext
// next to the reference one, so a slowdown caused by a poison is a number in
// the report rather than a hidden failure.
func latencyDelta(t *testing.T, step string, ref, probe slot) {
	t.Helper()
	if ref.combine == nil || probe.combine == nil {
		record(t, Result{Step: step, Kind: "measure", Pass: false, Notes: "reference or probe ciphertext never combined"})
		return
	}
	pre := int64(ref.combine.Block) - int64(ref.SubmitBlock)
	post := int64(probe.combine.Block) - int64(probe.SubmitBlock)
	record(t, Result{
		Step: step, Kind: "measure", Pass: true, LatencyBlocks: post,
		Notes: fmt.Sprintf("honest latency before poison=%d blocks, this one=%d blocks (delta %+d)", pre, post, post-pre),
	})
}

// poisonStatus reports what the fleet did with a poisoned slot: it must
// never be combined, and for a cofactor point it must never receive a
// partial either.
func poisonStatus(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, step string, partialsExpected bool,
) {
	t.Helper()
	rec, err := f.Services.Contracts.GetCombinedDecryption(ctx, app.Epoch, app.Aid, s.Idx)
	if err != nil {
		t.Fatal(err)
	}
	partials, err := f.partials(ctx, app.Epoch, app.Aid, s.Idx, scanFrom(epoch))
	if err != nil {
		t.Fatal(err)
	}
	head, _ := f.head(ctx)
	pass := !rec.Completed && (partialsExpected || len(partials) == 0)
	notes := fmt.Sprintf("partials=%d combined=%v observedAt=%d (+%d blocks)",
		len(partials), rec.Completed, head, head-s.SubmitBlock)
	if rec.Completed {
		notes += " plaintext=" + rec.Plaintext.String()
	}
	if len(partials) > 0 {
		notes += fmt.Sprintf(" firstPartialBlock=%d lastPartialBlock=%d", partials[0].Block, partials[len(partials)-1].Block)
	}
	record(t, Result{Step: step, Kind: "measure", Pass: pass, Block: s.SubmitBlock, Notes: notes})
	if !pass {
		t.Errorf("%s: %s", step, notes)
	}
}
