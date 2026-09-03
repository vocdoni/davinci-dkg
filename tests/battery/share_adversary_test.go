package battery

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// slot is one submitted ciphertext of the adversary scenarios.
type slot struct {
	Idx         uint16
	Plaintext   *big.Int
	C1, C2      types.CurvePoint
	SubmitBlock uint64
	combine     *combineEvent // set once the slot was seen combined
}

// submitSlots encrypts and submits `count` random plaintexts under app.
func submitSlots(ctx context.Context, t *testing.T, f *Fleet, app *application, count int, step string) []slot {
	t.Helper()
	slots := make([]slot, 0, count)
	for i := range count {
		m, err := randomPlaintext(20)
		if err != nil {
			t.Fatal(err)
		}
		c1, c2, err := app.encrypt(m)
		if err != nil {
			t.Fatal(err)
		}
		idx, out, err := f.submitCiphertext(ctx, app.Organizer, app.Epoch, app.Aid, c1, c2)
		if !expectOK(t, fmt.Sprintf("%s/submit-%d", step, i+1), "submitCiphertext", out, err, fmt.Sprintf("idx=%d", idx)) {
			t.FailNow()
		}
		slots = append(slots, slot{Idx: idx, Plaintext: m, C1: c1, C2: c2, SubmitBlock: out.Block})
	}
	return slots
}

// assertNotCombined waits `wait` blocks after `since` and records that the
// slot has not been combined, with the partial count observed.
func assertNotCombined(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, since, wait uint64, step string,
) {
	t.Helper()
	if _, err := f.waitBlock(ctx, since+wait); err != nil {
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
	pass := !rec.Completed
	record(t, Result{
		Step: step, Kind: "measure", Pass: pass,
		Notes: fmt.Sprintf("combined=%v partials=%d after %d blocks", rec.Completed, len(partials), wait),
	})
	if !pass {
		t.Errorf("%s: slot %d combined (plaintext %s) although it must not", step, s.Idx, rec.Plaintext)
	}
}

// assertCombines waits for the combine of a slot and checks the plaintext.
func assertCombines(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, since, wait uint64, step string,
) *combineEvent {
	t.Helper()
	ev, ok, err := f.waitCombine(ctx, app.Epoch, app.Aid, s.Idx, scanFrom(epoch), since+wait)
	if err != nil {
		t.Fatal(err)
	}
	partials, perr := f.partials(ctx, app.Epoch, app.Aid, s.Idx, scanFrom(epoch))
	if perr != nil {
		t.Fatal(perr)
	}
	res := Result{Step: step, Kind: "combineDecryption(node)", Notes: fmt.Sprintf("partials=%d", len(partials))}
	switch {
	case !ok:
		res.Notes += fmt.Sprintf(" not combined within %d blocks", wait)
	case ev.Plaintext.Cmp(s.Plaintext) != 0:
		res.Notes += fmt.Sprintf(" plaintext mismatch want %s got %s", s.Plaintext, ev.Plaintext)
	default:
		res.Pass = true
		res.Tx = ev.Tx.Hex()
		res.Gas = ev.Gas
		res.Block = ev.Block
		res.LatencyBlocks = int64(ev.Block) - int64(since)
		res.LatencySeconds = f.blocksToSeconds(ctx, since, ev.Block)
		res.Notes += fmt.Sprintf(" plaintext=%s combiner=%s", ev.Plaintext, ev.Sender.Hex())
	}
	record(t, res)
	if !res.Pass {
		t.Errorf("%s: %s", step, res.Notes)
	}
	return ev
}

// TestShareAdversary attacks the organizer-share path of one application:
// a tampered Δ, a share replayed across ciphertexts, a share relayed by a
// stranger, shares for non-existent indexes and a re-submission after the
// plaintext landed. The contract accepts any well-formed words; the nodes
// must refuse to combine until the DLEQ verifies.
func TestShareAdversary(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	noCombineWait := envUint64("BATTERY_NO_COMBINE_WAIT_BLOCKS", 40)
	combineWait := envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240)

	epochID, epoch, err := f.waitLiveEpoch(ctx, t, envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90))
	if err != nil {
		t.Fatal(err)
	}
	org, err := f.newActor(ctx, "share-organizer")
	if err != nil {
		t.Fatal(err)
	}
	stranger, err := f.newActor(ctx, "share-stranger")
	if err != nil {
		t.Fatal(err)
	}
	app, out, err := f.registerApplication(ctx, org, epochID, golangtypes.DKGTypesAppPolicy{})
	if !expectOK(t, "register", "registerApplication", out, err, "") {
		t.FailNow()
	}
	slots := submitSlots(ctx, t, f, app, 4, "setup")

	shareTampered(ctx, t, f, epoch, app, slots[0], noCombineWait, combineWait)
	shareReplayed(ctx, t, f, epoch, app, slots[0], slots[1], noCombineWait, combineWait)
	shareRelayed(ctx, t, f, epoch, app, stranger, slots[2], combineWait)
	shareBadIndexes(ctx, t, f, app, slots[3])
	shareAfterCombine(ctx, t, f, epoch, app, slots[3], combineWait)
}

// (a) Δ replaced by a random point, DLEQ words otherwise well-formed: the
// contract stores it, no node may combine; the honest share then overwrites
// it and the combine lands.
func shareTampered(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, noWait, wait uint64,
) {
	t.Helper()
	_, proof, err := app.share(s.Idx, s.C1)
	if err != nil {
		t.Fatal(err)
	}
	badDelta, err := randomSubgroupPoint()
	if err != nil {
		t.Fatal(err)
	}
	out, err := f.submitShareWords(ctx, org(app), app.Epoch, app.Aid, s.Idx, s.C1, s.C2, badDelta, proof)
	if !expectOK(t, "tampered/submit", "submitOrganizerShare", out, err, "Δ replaced by a random point") {
		return
	}
	assertNotCombined(ctx, t, f, epoch, app, s, out.Block, noWait, "tampered/no-combine")
	_, _, fix, err := f.releaseShare(ctx, app, s.Idx, s.C1, s.C2)
	if !expectOK(t, "tampered/overwrite", "submitOrganizerShare", fix, err, "honest share overwrites the tampered one") {
		return
	}
	assertCombines(ctx, t, f, epoch, app, s, fix.Block, wait, "tampered/recovery")
}

func org(app *application) *actor { return app.Organizer }

// (b) The words of slot a replayed for slot b: with a's ciphertext the
// contract rejects the hash binding; with b's ciphertext it accepts the
// words but the DLEQ is for a's C1, so no node combines until fixed.
func shareReplayed(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, a, b slot, noWait, wait uint64,
) {
	t.Helper()
	deltaA, proofA, err := app.share(a.Idx, a.C1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.submitShareWords(ctx, org(app), app.Epoch, app.Aid, b.Idx, a.C1, a.C2, deltaA, proofA)
	expectRevert(t, "replay/wrong-ciphertext-binding", err, "InvalidProofInput")

	out, err := f.submitShareWords(ctx, org(app), app.Epoch, app.Aid, b.Idx, b.C1, b.C2, deltaA, proofA)
	if !expectOK(t, "replay/submit", "submitOrganizerShare", out, err, "slot-a words under slot b's ciphertext") {
		return
	}
	assertNotCombined(ctx, t, f, epoch, app, b, out.Block, noWait, "replay/no-combine")
	_, _, fix, err := f.releaseShare(ctx, app, b.Idx, b.C1, b.C2)
	if !expectOK(t, "replay/overwrite", "submitOrganizerShare", fix, err, "") {
		return
	}
	assertCombines(ctx, t, f, epoch, app, b, fix.Block, wait, "replay/recovery")
}

// (c) Correct words posted from an unrelated address: the share path is
// permissionless, so the combine must land.
func shareRelayed(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, stranger *actor, s slot, wait uint64,
) {
	t.Helper()
	delta, proof, err := app.share(s.Idx, s.C1)
	if err != nil {
		t.Fatal(err)
	}
	out, err := f.submitShareWords(ctx, stranger, app.Epoch, app.Aid, s.Idx, s.C1, s.C2, delta, proof)
	if !expectOK(t, "relay/submit", "submitOrganizerShare", out, err, "posted by "+stranger.Label) {
		return
	}
	assertCombines(ctx, t, f, epoch, app, s, out.Block, wait, "relay/combine")
}

// (d) Indexes that do not exist, and an unregistered aid.
func shareBadIndexes(ctx context.Context, t *testing.T, f *Fleet, app *application, s slot) {
	t.Helper()
	delta, proof, err := app.share(s.Idx, s.C1)
	if err != nil {
		t.Fatal(err)
	}
	count, err := f.Services.Manager.CiphertextCount(f.callOpts(ctx), app.Epoch, app.Aid)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		step string
		idx  uint16
		aid  [32]byte
		want []string
	}{
		{"bad-index/unsubmitted", count + 5, app.Aid, []string{"CiphertextNotSubmitted"}},
		{"bad-index/zero", 0, app.Aid, []string{"InvalidCiphertext"}},
		{"bad-index/above-cap", 257, app.Aid, []string{"InvalidCiphertext"}},
	}
	for _, c := range cases {
		_, err := f.submitShareWords(ctx, org(app), app.Epoch, c.aid, c.idx, s.C1, s.C2, delta, proof)
		expectRevert(t, c.step, err, c.want...)
	}
	unknown, err := randomAid()
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.submitShareWords(ctx, org(app), app.Epoch, unknown, s.Idx, s.C1, s.C2, delta, proof)
	expectRevert(t, "bad-index/unregistered-aid", err, "InvalidApplication")
}

// (e) Once the plaintext is on chain the share is moot and re-submission
// must revert with AlreadyCombined.
func shareAfterCombine(ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, wait uint64) {
	t.Helper()
	delta, proof, out, err := f.releaseShare(ctx, app, s.Idx, s.C1, s.C2)
	if !expectOK(t, "after-combine/release", "submitOrganizerShare", out, err, "") {
		return
	}
	if ev := assertCombines(ctx, t, f, epoch, app, s, out.Block, wait, "after-combine/combine"); ev == nil {
		return
	}
	_, err = f.submitShareWords(ctx, org(app), app.Epoch, app.Aid, s.Idx, s.C1, s.C2, delta, proof)
	expectRevert(t, "after-combine/resubmit", err, "AlreadyCombined")
	fresh, freshProof, err := app.share(s.Idx, s.C1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.submitShareWords(ctx, org(app), app.Epoch, app.Aid, s.Idx, s.C1, s.C2, fresh, dleq.Proof{
		A1: freshProof.A1, A2: freshProof.A2, Response: freshProof.Response,
	})
	expectRevert(t, "after-combine/resubmit-fresh-nonce", err, "AlreadyCombined")
}
