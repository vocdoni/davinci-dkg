package battery

import (
	"context"
	"fmt"
	"math/big"
	"testing"

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

// TestRevealAdversary attacks the one-shot organizer reveal: a wrong secret,
// the sealed window before the right one lands, a stranger relaying it (the
// call is permissionless), a second reveal, and a reveal aimed at an
// automatic application. The fleet must combine nothing while the
// application is locked and everything — including ciphertexts submitted
// afterwards — once the secret is out.
func TestRevealAdversary(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	noCombineWait := envUint64("BATTERY_NO_COMBINE_WAIT_BLOCKS", 40)
	combineWait := envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240)

	// Two pool keys: the locked application and the automatic one.
	epochID, epoch, err := f.waitLiveEpoch(ctx, t, envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90), 2)
	if err != nil {
		t.Fatal(err)
	}
	org, err := f.newActor(ctx, "reveal-organizer")
	if err != nil {
		t.Fatal(err)
	}
	stranger, err := f.newActor(ctx, "reveal-stranger")
	if err != nil {
		t.Fatal(err)
	}
	app, out, err := f.registerApplication(ctx, org, epochID, golangtypes.DKGTypesAppPolicy{})
	if !expectOK(t, "register", "registerApplication", out, err, "organizer-locked") {
		t.FailNow()
	}
	slots := submitSlots(ctx, t, f, app, 2, "setup")

	revealWrongSecret(ctx, t, f, epoch, app, slots[0], noCombineWait)
	revealRelayed(ctx, t, f, epoch, app, stranger, slots, combineWait)
	revealTwice(ctx, t, f, app)
	revealAfterTheFact(ctx, t, f, epoch, app, combineWait)
	revealAutomatic(ctx, t, f, epochID, org)
}

// (a) A secret whose sk·G is not the registered PK_org is refused on chain,
// and the fleet stays unable to combine.
func revealWrongSecret(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, s slot, noWait uint64,
) {
	t.Helper()
	wrong, err := randomScalar()
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.revealSecret(ctx, app.Organizer, app, wrong)
	expectRevert(t, "reveal/wrong-secret", err, "InvalidOrganizerSecret")
	_, err = f.revealSecret(ctx, app.Organizer, app, big.NewInt(0))
	expectRevert(t, "reveal/zero-secret", err, "InvalidOrganizerSecret")
	assertNotCombined(ctx, t, f, epoch, app, s, s.SubmitBlock, noWait, "reveal/sealed")
}

// (b) The right secret posted by an unrelated address: the reveal is
// permissionless, and every ciphertext of the application unblocks at once.
func revealRelayed(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application,
	stranger *actor, slots []slot, wait uint64,
) {
	t.Helper()
	out, err := f.revealSecret(ctx, stranger, app, app.SkOrg)
	if !expectOK(t, "reveal/relayed", "revealOrganizerSecret", out, err, "posted by "+stranger.Label) {
		return
	}
	record, err := f.Services.AppManager.GetApplication(f.callOpts(ctx), app.Epoch, app.Aid)
	if err != nil {
		t.Fatal(err)
	}
	if record.OrganizerSecret.Cmp(app.SkOrg) != 0 {
		t.Fatalf("stored organizer secret %s does not match the revealed one", record.OrganizerSecret)
	}
	for _, s := range slots {
		assertCombines(ctx, t, f, epoch, app, s, out.Block, wait, fmt.Sprintf("reveal/combine-%d", s.Idx))
	}
}

// (c) The reveal is one-shot: a second call reverts whatever it carries.
func revealTwice(ctx context.Context, t *testing.T, f *Fleet, app *application) {
	t.Helper()
	_, err := f.revealSecret(ctx, app.Organizer, app, app.SkOrg)
	expectRevert(t, "reveal/second", err, "AlreadyRevealed")
	other, err := randomScalar()
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.revealSecret(ctx, app.Organizer, app, other)
	expectRevert(t, "reveal/second-other-secret", err, "AlreadyRevealed", "InvalidOrganizerSecret")
}

// (d) A ciphertext submitted after the reveal needs no further organizer
// action: the committee owns the application from then on.
func revealAfterTheFact(
	ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, app *application, wait uint64,
) {
	t.Helper()
	s := submitSlots(ctx, t, f, app, 1, "reveal/post")[0]
	assertCombines(ctx, t, f, epoch, app, s, s.SubmitBlock, wait, "reveal/post-combine")
}

// (e) An automatic application has no secret to reveal, and says so.
func revealAutomatic(ctx context.Context, t *testing.T, f *Fleet, epochID [12]byte, org *actor) {
	t.Helper()
	auto, out, err := f.registerApplication(ctx, org, epochID, automaticPolicy())
	if !expectOK(t, "reveal/automatic-register", "registerApplication", out, err, "") {
		return
	}
	secret, err := randomScalar()
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.revealSecret(ctx, org, auto, secret)
	expectRevert(t, "reveal/automatic", err, "AlreadyRevealed")
}

// TestCrossApplicationAdversary copies one application's ciphertext into
// another of the same epoch. Each application holds its own committee key,
// so the copy decrypts under the wrong key and is worthless: the fleet's
// partials for it never add up and no combine ever lands. This is the
// property pool keys exist for.
func TestCrossApplicationAdversary(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	noCombineWait := envUint64("BATTERY_NO_COMBINE_WAIT_BLOCKS", 40)
	combineWait := envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240)

	// Two pool keys: application A and application B.
	epochID, epoch, err := f.waitLiveEpoch(ctx, t, envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90), 2)
	if err != nil {
		t.Fatal(err)
	}
	orgA, err := f.newActor(ctx, "xapp-organizer-a")
	if err != nil {
		t.Fatal(err)
	}
	orgB, err := f.newActor(ctx, "xapp-organizer-b")
	if err != nil {
		t.Fatal(err)
	}
	appA, out, err := f.registerApplication(ctx, orgA, epochID, automaticPolicy())
	if !expectOK(t, "xapp/register-a", "registerApplication", out, err, "") {
		t.FailNow()
	}
	appB, out, err := f.registerApplication(ctx, orgB, epochID, automaticPolicy())
	if !expectOK(t, "xapp/register-b", "registerApplication", out, err, "") {
		t.FailNow()
	}
	samePool := appA.PoolIndex == appB.PoolIndex
	record(t, Result{
		Step: "xapp/distinct-pool-keys", Kind: "measure", Pass: !samePool,
		Notes: fmt.Sprintf("poolIndex a=%d b=%d", appA.PoolIndex, appB.PoolIndex),
	})
	if samePool {
		t.Fatalf("applications must claim distinct pool keys")
	}

	// A genuine ciphertext of A, combined by the fleet.
	honest := submitSlots(ctx, t, f, appA, 1, "xapp/honest")[0]
	assertCombines(ctx, t, f, epoch, appA, honest, honest.SubmitBlock, combineWait, "xapp/honest-combine")

	// The very same words under B.
	idx, out, err := f.submitCiphertext(ctx, appB.Organizer, epochID, appB.Aid, honest.C1, honest.C2)
	if !expectOK(t, "xapp/copy", "submitCiphertext", out, err, "A's ciphertext replayed into B") {
		return
	}
	copied := slot{Idx: idx, Plaintext: honest.Plaintext, C1: honest.C1, C2: honest.C2, SubmitBlock: out.Block}
	assertNotCombined(ctx, t, f, epoch, appB, copied, out.Block, noCombineWait, "xapp/copy-no-combine")

	// If the nodes ever do combine it, the value must not be A's plaintext.
	rec, err := f.Services.Contracts.GetCombinedDecryption(ctx, epochID, appB.Aid, idx)
	if err != nil {
		t.Fatal(err)
	}
	leaked := rec.Completed && rec.Plaintext.Cmp(honest.Plaintext) == 0
	record(t, Result{
		Step: "xapp/no-oracle", Kind: "measure", Pass: !leaked,
		Notes: fmt.Sprintf("combined=%v plaintextLeaked=%v", rec.Completed, leaked),
	})
	if leaked {
		t.Fatalf("application B decrypted application A's ciphertext")
	}
}
