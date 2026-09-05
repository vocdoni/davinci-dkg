package battery

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// registerOperator registers the actor's deterministic BabyJubJub key in the
// registry (the same derivation the node binary and the harness use) and
// returns the mined tx.
func (f *Fleet) registerOperator(ctx context.Context, a *actor) (txOutcome, error) {
	secret, err := nodeSecret(a.PrivKey)
	if err != nil {
		return txOutcome{}, err
	}
	pub := helpers.ScalarBasePoint(secret)
	_, _, proof, err := schnorr.ProveOperatorRegister(secret, a.Address())
	if err != nil {
		return txOutcome{}, fmt.Errorf("operator schnorr proof: %w", err)
	}
	return f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Registry.RegisterKey(auth, pub.X, pub.Y, proof.Ax, proof.Ay, proof.Z)
	})
}

func (f *Fleet) claimSlot(ctx context.Context, a *actor, epochID [12]byte) (txOutcome, error) {
	return f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Manager.ClaimSlot(auth, epochID)
	})
}

// awaitBoundary makes sure the next createEpoch is at least `lead` blocks
// away (so operator registrations mined now land before it) and returns
// the current nonce plus the boundary block.
func (f *Fleet) awaitBoundary(ctx context.Context, t *testing.T, lead uint64) (uint64, uint64, error) {
	t.Helper()
	for {
		nonce, err := f.epochNonce(ctx)
		if err != nil {
			return 0, 0, err
		}
		next, err := f.nextEpochStart(ctx)
		if err != nil {
			return 0, 0, err
		}
		head, err := f.head(ctx)
		if err != nil {
			return 0, 0, err
		}
		if head+lead <= next {
			return nonce, next, nil
		}
		t.Logf("boundary at %d is only %d blocks away (need %d) — waiting for the next epoch to be created", next, next-head, lead)
		if _, err := f.waitNonceAbove(ctx, nonce); err != nil {
			return 0, 0, err
		}
	}
}

// joinedEpoch is the epoch the battery's own operators claimed slots in.
type joinedEpoch struct {
	ID   [12]byte
	View web3.EpochView
}

// enrollOperators registers every actor as an operator ahead of the next
// epoch boundary, so they sit in the registry snapshot createEpoch takes.
func enrollOperators(ctx context.Context, t *testing.T, f *Fleet, operators []*actor) (uint64, uint64) {
	t.Helper()
	nonce, boundary, err := f.awaitBoundary(ctx, t, 6)
	if err != nil {
		t.Fatal(err)
	}
	for _, op := range operators {
		out, err := f.registerOperator(ctx, op)
		if !expectOK(t, op.Label+"/registerKey", "registerKey", out, err, "") {
			t.FailNow()
		}
	}
	return nonce, boundary
}

// waitNewEpoch blocks until the epoch after `nonce` exists.
func waitNewEpoch(ctx context.Context, t *testing.T, f *Fleet, nonce uint64) joinedEpoch {
	t.Helper()
	next, err := f.waitNonceAbove(ctx, nonce)
	if err != nil {
		t.Fatal(err)
	}
	id := f.epochID(next)
	view, err := f.epoch(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("joined epoch %x: start=%d seed=%d t=%d n=%d mMin=%d alpha=%d",
		id, view.StartBlock, view.SeedBlock, view.Policy.Threshold, view.Policy.CommitteeSize,
		view.Policy.MinValidContributions, view.Policy.LotteryAlphaBps)
	return joinedEpoch{ID: id, View: view}
}

// claimAtSeed waits for the seed block to pass and claims immediately; the
// real nodes poll every 5 s so a tight poller normally lands first.
// lostLottery is claimPhase's answer when the member was not admitted by the
// lottery in this epoch (a correct outcome the caller retries).
const lostLottery uint16 = 0xffff

func claimAtSeed(ctx context.Context, t *testing.T, f *Fleet, a *actor, ep joinedEpoch) (txOutcome, error) {
	t.Helper()
	if _, err := f.waitBlock(ctx, ep.View.SeedBlock+1); err != nil {
		t.Fatal(err)
	}
	return f.claimSlot(ctx, a, ep.ID)
}

// memberIndex resolves the 1-based committee index of an address.
func memberIndex(committee []common.Address, addr common.Address) uint16 {
	for i, member := range committee {
		if member == addr {
			return uint16(i + 1)
		}
	}
	return 0
}

// TestCommitteeAdversary registers a fresh operator ahead of an epoch,
// claims a slot in the real lottery and then attacks every committee-side
// entry point while interoperating with the real nodes: duplicate / late /
// unregistered claims, foreign and malformed contributions, an early
// finalize, an abort of a healthy epoch, out-of-policy createEpoch calls,
// then — as a genuine member of the Live committee — a real partial
// decryption, its duplicate, a broken one and a combine after the nodes'.
func TestCommitteeAdversary(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	combineWait := envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240)

	member, err := f.newActor(ctx, "adversary-member")
	if err != nil {
		t.Fatal(err)
	}
	late, err := f.newActor(ctx, "late-operator")
	if err != nil {
		t.Fatal(err)
	}
	stranger, err := f.newActor(ctx, "unregistered")
	if err != nil {
		t.Fatal(err)
	}
	organizer, err := f.newActor(ctx, "committee-organizer")
	if err != nil {
		t.Fatal(err)
	}

	nonce, boundary := enrollOperators(ctx, t, f, []*actor{member})
	if _, err := f.waitBlock(ctx, boundary); err != nil {
		t.Fatal(err)
	}
	createEpochProbes(ctx, t, f, stranger)
	// The member is a fresh operator in the real lottery: with N registered
	// operators its admission probability is min(1, α·n/N), so a lost lottery
	// is a correct outcome, not a failure. Try up to three epochs.
	var ep joinedEpoch
	var myIdx uint16
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			// The late operator of the previous attempt is registered by now.
			if late, err = f.newActor(ctx, "late-operator"); err != nil {
				t.Fatal(err)
			}
		}
		ep = waitNewEpoch(ctx, t, f, nonce)
		myIdx = claimPhase(ctx, t, f, ep, member, late, stranger)
		if myIdx != lostLottery {
			break
		}
		nonce = binary.BigEndian.Uint64(ep.ID[4:])
	}
	if myIdx == 0 || myIdx == lostLottery {
		return
	}
	own := contributionPhase(ctx, t, f, ep, member, stranger, myIdx)
	if own == nil {
		return
	}
	liveView, err := f.waitStatus(ctx, ep.ID, statusLive)
	if err != nil {
		t.Fatal(err)
	}
	record(t, Result{
		Step: "epoch-live", Kind: "measure", Pass: true, Block: liveView.Policy.LiveNotBeforeBlock,
		Notes: fmt.Sprintf("contributions=%d/%d", liveView.ContributionCount, liveView.Policy.CommitteeSize),
	})
	ep.View = liveView
	decryptionPhase(ctx, t, f, ep, member, organizer, myIdx, own, combineWait)
}

// createEpochProbes fires out-of-policy createEpoch calls at the cadence
// boundary, racing the nodes' valid one. Each is only simulated (gas
// estimation reverts before anything is signed).
func createEpochProbes(ctx context.Context, t *testing.T, f *Fleet, a *actor) {
	t.Helper()
	type probe struct {
		step             string
		th, n, mMin, bps uint16
	}
	probes := []probe{
		{"createEpoch/n-above-maxN", 17, f.MaxN + 1, 20, 15000},
		{"createEpoch/alpha-below-one", 16, 24, 20, 9999},
		{"createEpoch/minValid-below-threshold", 16, 24, 15, 15000},
	}
	if f.MinThreshold > 1 {
		probes = append(probes, probe{"createEpoch/threshold-below-floor", f.MinThreshold - 1, 24, 20, 15000})
	}
	if f.MinCommitteeSize > 1 {
		probes = append(probes, probe{"createEpoch/n-below-floor", 1, f.MinCommitteeSize - 1, 1, 15000})
	}
	if f.MaxAlphaBps < 65535 {
		probes = append(probes, probe{"createEpoch/alpha-above-cap", 16, 24, 20, f.MaxAlphaBps + 1})
	} else {
		record(t, Result{
			Step: "createEpoch/alpha-above-cap", Kind: "revert", Pass: true,
			Notes: "not testable: MAX_LOTTERY_ALPHA_BPS is 65535 (uint16 max) on this deployment",
		})
	}
	var wg sync.WaitGroup
	for _, p := range probes {
		wg.Add(1)
		go func(p probe) {
			defer wg.Done()
			_, err := f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
				return f.Services.Manager.CreateEpoch(auth, p.th, p.n, p.mMin, p.bps)
			})
			// InvalidPhase means a node's valid createEpoch landed first and
			// closed the cadence gate before the policy was even looked at.
			expectRevert(t, p.step, err, "InvalidPolicy", "InvalidPhase")
		}(p)
	}
	wg.Wait()
}

// claimPhase runs the lottery-side checks and returns the member's index.
func claimPhase(ctx context.Context, t *testing.T, f *Fleet, ep joinedEpoch, member, late, stranger *actor) uint16 {
	t.Helper()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// Registered after createEpoch: outside the lottery snapshot. The
		// registration lands at startBlock+1 or later, the claim right after;
		// the committee (24 of 33) usually fills a block or two later.
		defer wg.Done()
		out, err := f.registerOperator(ctx, late)
		if !expectOK(t, "late/registerKey", "registerKey", out, err, fmt.Sprintf("epoch start=%d", ep.View.StartBlock)) {
			return
		}
		_, err = claimAtSeed(ctx, t, f, late, ep)
		expectRevertOrInconclusive(t, "claim/registered-after-createEpoch", err, "NotInSnapshot", "SlotsFull", "InvalidPhase")
	}()
	go func() {
		defer wg.Done()
		_, err := claimAtSeed(ctx, t, f, stranger, ep)
		expectRevertOrInconclusive(t, "claim/unregistered", err, "NotRegistered", "SlotsFull", "InvalidPhase")
	}()
	out, err := claimAtSeed(ctx, t, f, member, ep)
	wg.Wait()
	if err != nil {
		if name, ok := revertName(err); ok && name == "NotEligible" {
			record(t, Result{
				Step: "claim/member", Kind: "claimSlot", Pass: true,
				Notes: "lost the lottery (admission is min(1, α·n/N) per epoch); trying the next epoch",
			})
			t.Logf("claim/member: not eligible in epoch %x, trying the next one", ep.ID)
			return lostLottery
		}
	}
	if !expectOK(t, "claim/member", "claimSlot", out, err, "fresh operator in the real lottery") {
		return 0
	}
	_, err = f.claimSlot(ctx, member, ep.ID)
	expectRevert(t, "claim/duplicate", err, "AlreadyClaimed", "InvalidPhase", "SlotsFull")

	view, err := f.waitStatus(ctx, ep.ID, statusKeyAssembly)
	if err != nil {
		t.Fatal(err)
	}
	committee, err := f.committee(ctx, ep.ID)
	if err != nil {
		t.Fatal(err)
	}
	idx := memberIndex(committee, member.Address())
	record(t, Result{
		Step: "claim/committee-filled", Kind: "measure", Pass: idx > 0, Block: view.SeedBlock,
		Notes: fmt.Sprintf("committee=%d memberIndex=%d", len(committee), idx),
	})
	if idx == 0 {
		t.Errorf("member %s is not in the committee", member.Address().Hex())
	}
	return idx
}

// contributionPhase submits the member's real contribution amid foreign,
// malformed, duplicate and mistimed calls.
func contributionPhase(
	ctx context.Context, t *testing.T, f *Fleet, ep joinedEpoch, member, stranger *actor, myIdx uint16,
) *ownContribution {
	t.Helper()
	th, n := ep.View.Policy.Threshold, ep.View.Policy.CommitteeSize
	base, err := randomScalars(int(th))
	if err != nil {
		t.Fatal(err)
	}
	// One contribution deals the epoch's whole pool: MaxK polynomials under
	// one ephemeral per recipient.
	coeffs := helpers.DealPoolCoefficients(base)
	recipients := make([]uint16, n)
	for i := range recipients {
		recipients[i] = uint16(i + 1)
	}
	sub, err := helpers.BuildContributionSubmission(ctx, f.Services, ep.ID, th, n, myIdx, coeffs, recipients)
	if err != nil {
		t.Fatalf("build contribution: %v", err)
	}
	contribute := func(a *actor, idx uint16, proof []byte) (txOutcome, error) {
		return f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return f.Services.Manager.SubmitContribution(auth, ep.ID, idx,
				sub.CommitmentsHash, sub.EncryptedSharesHash, sub.Transcript, proof, sub.Input)
		})
	}
	_, err = contribute(stranger, myIdx, sub.Proof)
	expectRevert(t, "contribution/non-member", err, "NotSelectedParticipant")
	_, err = contribute(member, myIdx, sub.Proof[:len(sub.Proof)-32])
	expectRevert(t, "contribution/malformed-proof", err, "InvalidProofEncoding", "ProofInvalid", "InvalidProofInput")
	out, err := contribute(member, myIdx, sub.Proof)
	if !expectOK(t, "contribution/member", "submitContribution", out, err, fmt.Sprintf("index=%d", myIdx)) {
		return nil
	}
	_, err = contribute(member, myIdx, sub.Proof)
	expectRevert(t, "contribution/duplicate", err, "AlreadyContributed")

	// Before the gate the contract refuses the call on phase alone, so no
	// proof is needed to probe it: empty arguments never reach the checks.
	head, _ := f.head(ctx)
	_, err = f.send(ctx, member, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Manager.FinalizeEpoch(auth, ep.ID, [32]byte{}, nil, nil, nil)
	})
	expectRevert(t, fmt.Sprintf("finalize/before-gate(head=%d,gate=%d)", head, ep.View.Policy.LiveNotBeforeBlock),
		err, "InvalidPhase", "AlreadyLive")
	_, err = f.send(ctx, member, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Manager.AbortEpoch(auth, ep.ID)
	})
	expectRevert(t, "abort/healthy-epoch", err, "InvalidPhase")
	return &ownContribution{Index: myIdx, Coefficients: coeffs}
}

// decryptionPhase acts as a genuine committee member of the Live epoch: it
// registers an automatic application, recovers its own share of that
// application's pool key from the contributions' calldata, and posts a
// partial with the Merkle path against the key's share root.
func decryptionPhase(
	ctx context.Context, t *testing.T, f *Fleet, ep joinedEpoch, member, organizer *actor,
	myIdx uint16, own *ownContribution, combineWait uint64,
) {
	t.Helper()
	rec, err := f.Services.Manager.GetContribution(f.callOpts(ctx), ep.ID, member.Address())
	if err != nil || !rec.Accepted {
		t.Fatalf("member contribution not accepted on chain (err=%v)", err)
	}
	secret, err := nodeSecret(member.PrivKey)
	if err != nil {
		t.Fatal(err)
	}
	committee, err := f.committee(ctx, ep.ID)
	if err != nil {
		t.Fatal(err)
	}

	app, out, err := f.registerApplication(ctx, organizer, ep.ID, automaticPolicy())
	if !expectOK(t, "app/register", "registerApplication", out, err, "") {
		return
	}

	share, accepted, err := recoverPrivateShare(
		ctx, f, ep.ID, myIdx, app.PoolIndex, committee, secret, own, scanFrom(ep.View),
	)
	if err != nil {
		t.Fatalf("recover private share: %v", err)
	}
	match, err := shareCommitmentMatches(ctx, f, ep.ID, app.PoolIndex, myIdx, share, scanFrom(ep.View))
	if err != nil {
		t.Fatal(err)
	}
	record(t, Result{
		Step: "share/recovered-from-calldata", Kind: "measure", Pass: match,
		Notes: fmt.Sprintf("accepted contributions=%d, d_i·G matches the finalization's share commitment=%v", accepted, match),
	})
	if !match {
		t.Fatalf("recovered share does not match the share commitment the finalization published")
	}
	tree, err := poolShareTree(ctx, f, ep.ID, app.PoolIndex, scanFrom(ep.View))
	if err != nil {
		t.Fatalf("rebuild share tree: %v", err)
	}

	s := submitSlots(ctx, t, f, app, 1, "app")[0]
	nonce, err := randomScalar()
	if err != nil {
		t.Fatal(err)
	}
	partial, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, ep.ID, app.Aid, s.Idx, myIdx, s.C1, s.C2, share, nonce, tree,
	)
	if err != nil {
		t.Fatalf("build partial decryption: %v", err)
	}
	submitPartial := func(proof []byte, path [][32]byte) (txOutcome, error) {
		return f.send(ctx, member, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return f.Services.Manager.SubmitPartialDecryption(auth, ep.ID, app.Aid, myIdx, s.Idx,
				s.C1.X, s.C1.Y, s.C2.X, s.C2.Y, partial.DeltaHash, proof, partial.Input, path)
		})
	}
	_, err = submitPartial(partial.Proof[:len(partial.Proof)-32], partial.ShareProof)
	expectRevert(t, "partial/broken-proof", err, "InvalidProofEncoding", "ProofInvalid", "InvalidProofInput")
	_, err = submitPartial(partial.Proof, make([][32]byte, len(partial.ShareProof)))
	expectRevert(t, "partial/broken-share-path", err, "InvalidProofInput", "InvalidShareProof")
	out, err = submitPartial(partial.Proof, partial.ShareProof)
	expectOK(t, "partial/member", "submitPartialDecryption", out, err, fmt.Sprintf("index=%d", myIdx))
	_, err = submitPartial(partial.Proof, partial.ShareProof)
	expectRevert(t, "partial/duplicate", err, "AlreadyPartiallyDecrypted")

	ev := assertCombines(ctx, t, f, ep.View, app, s, s.SubmitBlock, combineWait, "app/combine-by-nodes")
	if ev == nil {
		return
	}
	combineAfterNodes(ctx, t, f, ep, app, s, member)
}

// combineAfterNodes builds a genuine combine proof from the partials on
// chain and submits it after the nodes' combine landed: AlreadyCombined.
func combineAfterNodes(
	ctx context.Context, t *testing.T, f *Fleet, ep joinedEpoch, app *application, s slot, member *actor,
) {
	t.Helper()
	th := ep.View.Policy.Threshold
	partials, err := f.partials(ctx, ep.ID, app.Aid, s.Idx, scanFrom(ep.View))
	if err != nil {
		t.Fatal(err)
	}
	if len(partials) < int(th) {
		t.Fatalf("only %d partials on chain, need %d", len(partials), th)
	}
	idxs := make([]uint16, th)
	deltas := make([]types.CurvePoint, th)
	for i := range th {
		idxs[i] = partials[i].Index
		deltas[i] = partials[i].Delta
	}
	combine, err := helpers.BuildDecryptCombineOutputFromCiphertext(ctx, ep.ID, app.Aid, s.Idx, th,
		s.C1, s.C2, app.PKOrg, app.SkOrg, idxs, deltas, s.Plaintext)
	if err != nil {
		t.Fatalf("build combine: %v", err)
	}
	_, err = f.send(ctx, member, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Manager.CombineDecryption(auth, ep.ID, app.Aid, s.Idx,
			combine.CombineHash, combine.Plaintext, combine.Transcript, combine.Proof, combine.Input)
	})
	expectRevert(t, "combine/after-nodes", err, "AlreadyCombined")
	record(t, Result{
		Step: "combine/proof-built", Kind: "measure", Pass: true,
		Notes: fmt.Sprintf("genuine combine proof over %d on-chain partials built by the member", th),
	})
}
