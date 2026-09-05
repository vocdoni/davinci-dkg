package battery

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// releaseMode is how an organizer treats its application's secret: an
// automatic registration (nothing to release), a locked one revealed after a
// delay, or a locked one whose secret is never published.
type releaseMode int

const (
	releaseImmediately releaseMode = iota
	releaseDelayed
	releaseWithheld
)

func (m releaseMode) String() string {
	switch m {
	case releaseImmediately:
		return "immediate"
	case releaseDelayed:
		return "delayed"
	default:
		return "withheld"
	}
}

// swarmConfig is the knob set of TestOrganizerSwarm, all env-overridable.
type swarmConfig struct {
	organizers     int
	perOrganizer   int
	delayBlocks    uint64
	withheldWait   uint64
	combineWait    uint64
	minServiceLeft uint64
	activationWait uint64
	plaintextBits  uint
}

// activationPlateau is how many blocks the ready-key count may stand still
// before a wave settles for what the fleet activated: the nodes activate one
// key per tick (a few blocks), so a longer stall means ActivateAhead is hit.
const activationPlateau = 15

func loadSwarmConfig() swarmConfig {
	return swarmConfig{
		// Six organizers plus the reveal adversary's two applications fill
		// exactly one epoch's pool of MaxK keys; a larger swarm spills into
		// the next epoch, wave by wave.
		organizers:     envInt("BATTERY_ORGANIZERS", 6),
		perOrganizer:   envInt("BATTERY_CIPHERTEXTS", 6),
		delayBlocks:    envUint64("BATTERY_REVEAL_DELAY_BLOCKS", 6),
		withheldWait:   envUint64("BATTERY_WITHHELD_WAIT_BLOCKS", 40),
		combineWait:    envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240),
		minServiceLeft: envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90),
		activationWait: activationWait(),
		plaintextBits:  20,
	}
}

// ctTrack follows one ciphertext from submission to its final verdict.
type ctTrack struct {
	Organizer   string
	Epoch       [12]byte
	Aid         [32]byte
	Idx         uint16
	Mode        releaseMode
	Plaintext   *big.Int
	C1, C2      types.CurvePoint
	SubmitBlock uint64
	SubmitGas   uint64
	RevealBlock uint64
	RevealGas   uint64
	Combine     *combineEvent
	Partials    int
	PartialGas  uint64 // average gas of the partials on chain
	Threshold   int    // t of the ciphertext's epoch
	Committee   int    // n of the ciphertext's epoch
	ServiceEnd  uint64 // cadence boundary of the ciphertext's epoch
	Pass        bool
	Note        string
}

func (c *ctTrack) label() string { return fmt.Sprintf("%s/ct%d(%s)", c.Organizer, c.Idx, c.Mode) }

// TestOrganizerSwarm runs K organizers concurrently, each registering an
// application — automatic, locked-and-revealed after a delay, or
// locked-and-withheld — and submitting M ciphertexts under it. Every
// registration claims one of its epoch's MaxK pool keys, so the swarm runs
// in waves: as many organizers as the newest Live epoch has keys left (and
// the nodes have activated), the rest in the next epoch once the nodes
// create it. Every unlocked application's ciphertexts must combine to their
// plaintexts; the withheld ones must not combine at all. Per-ciphertext
// latency, partial counts and gas are reported; an epoch boundary crossed
// mid-run is reported, not fatal.
func TestOrganizerSwarm(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	cfg := loadSwarmConfig()

	actors, err := f.newActors(ctx, "org", cfg.organizers)
	if err != nil {
		t.Fatal(err)
	}

	started := time.Now()
	tracks := make([][]*ctTrack, cfg.organizers)
	var wg sync.WaitGroup
	for first := 0; first < cfg.organizers; {
		// A wave's organizers keep running while the next wave waits for
		// its epoch, so the fleet is never idle between waves.
		epochID, epoch, size, err := f.swarmWave(ctx, t, cfg, cfg.organizers-first)
		if err != nil {
			t.Errorf("swarm: no epoch for organizers %d..%d: %v", first, cfg.organizers-1, err)
			break
		}
		t.Logf("swarm: wave of %d organizers (%d..%d) in epoch %x t=%d n=%d serviceEnd=%d ciphertexts/organizer=%d",
			size, first, first+size-1, epochID, epoch.Policy.Threshold, epoch.Policy.CommitteeSize,
			f.serviceEnd(epoch), cfg.perOrganizer)
		for i := first; i < first+size; i++ {
			wg.Add(1)
			go func(i int, a *actor) {
				defer wg.Done()
				tracks[i] = runOrganizer(ctx, t, f, a, epochID, epoch, planFor(i), cfg)
			}(i, actors[i])
		}
		first += size
	}
	wg.Wait()

	var all []*ctTrack
	for _, ts := range tracks {
		all = append(all, ts...)
	}
	summarizeSwarm(ctx, t, f, all, started)
}

// swarmWave picks the epoch of the next wave and sizes it: the newest Live
// epoch with enough service blocks left and at least one unclaimed pool
// key, then as many organizers as that epoch has keys left — MaxK at most —
// and the nodes activate within the budget. A fleet that activates fewer
// keys ahead than that (DAVINCI_DKG_ACTIVATE_AHEAD below MaxK) gets a
// smaller wave and a log line rather than a wave of PoolKeyNotActive reverts.
func (f *Fleet) swarmWave(
	ctx context.Context, t *testing.T, cfg swarmConfig, remaining int,
) ([12]byte, web3.EpochView, int, error) {
	t.Helper()
	for {
		epochID, epoch, err := f.waitLiveEpoch(ctx, t, cfg.minServiceLeft, 1)
		if err != nil {
			return epochID, epoch, 0, err
		}
		next, err := f.poolStatus(ctx, epochID)
		if err != nil {
			return epochID, epoch, 0, err
		}
		unclaimed, _ := poolCounts(next)
		size := min(remaining, unclaimed, ccommon.MaxK)
		ready, err := f.waitPoolKeys(ctx, epochID, size, cfg.activationWait, activationPlateau)
		switch {
		case err == nil:
			return epochID, epoch, size, nil
		case errors.Is(err, errPoolSlow) && ready > 0:
			t.Logf("swarm: %v — shrinking the wave to %d organizers", err, ready)
			return epochID, epoch, ready, nil
		case errors.Is(err, errPoolShort):
			t.Logf("swarm: %v — picking the epoch again", err) // claimed out under us
		default:
			return epochID, epoch, 0, err
		}
	}
}

// runOrganizer is one organizer's whole life: register in the mode it drew,
// submit M ciphertexts, reveal its secret per plan, then judge every
// ciphertext.
func runOrganizer(
	ctx context.Context, t *testing.T, f *Fleet, a *actor, epochID [12]byte, epoch web3.EpochView,
	mode releaseMode, cfg swarmConfig,
) []*ctTrack {
	t.Helper()
	policy := automaticPolicy()
	if mode != releaseImmediately {
		policy = golangtypes.DKGTypesAppPolicy{} // organizer-locked
	}
	app, out, err := f.registerApplication(ctx, a, epochID, policy)
	if !expectOK(t, a.Label+"/register", "registerApplication", out, err, mode.String()) {
		return nil
	}
	tracks := submitSwarmCiphertexts(ctx, t, f, a, app, mode, cfg)
	for _, tr := range tracks {
		tr.Threshold = int(epoch.Policy.Threshold)
		tr.Committee = int(epoch.Policy.CommitteeSize)
		tr.ServiceEnd = f.serviceEnd(epoch)
	}
	revealSwarmSecret(ctx, t, f, app, tracks, mode, cfg)
	for _, tr := range tracks {
		judgeCiphertext(ctx, t, f, epoch, tr, cfg)
	}
	return tracks
}

func submitSwarmCiphertexts(
	ctx context.Context, t *testing.T, f *Fleet, a *actor, app *application, mode releaseMode, cfg swarmConfig,
) []*ctTrack {
	t.Helper()
	tracks := make([]*ctTrack, 0, cfg.perOrganizer)
	for j := range cfg.perOrganizer {
		m, err := randomPlaintext(cfg.plaintextBits)
		if err != nil {
			t.Error(err)
			return tracks
		}
		c1, c2, err := app.encrypt(m)
		if err != nil {
			t.Error(err)
			return tracks
		}
		idx, out, err := f.submitCiphertext(ctx, a, app.Epoch, app.Aid, c1, c2)
		step := fmt.Sprintf("%s/ct%d/submit", a.Label, j+1)
		if !expectOK(t, step, "submitCiphertext", out, err, fmt.Sprintf("idx=%d m=%s", idx, m)) {
			continue
		}
		tracks = append(tracks, &ctTrack{
			Organizer: a.Label, Epoch: app.Epoch, Aid: app.Aid, Idx: idx, Mode: mode, Plaintext: m,
			C1: c1, C2: c2, SubmitBlock: out.Block, SubmitGas: out.Gas,
		})
	}
	return tracks
}

// planFor spreads the modes across the organizers: every 4th one withholds
// its secret (~25%), the rest alternate between an automatic application and
// a locked one revealed after a delay.
func planFor(i int) releaseMode {
	switch {
	case i%4 == 3:
		return releaseWithheld
	case i%2 == 0:
		return releaseImmediately
	default:
		return releaseDelayed
	}
}

// revealSwarmSecret publishes the application's organizer secret according
// to its mode. The reveal is one-shot and per application, so every
// ciphertext of the organizer unblocks at the same block.
func revealSwarmSecret(
	ctx context.Context, t *testing.T, f *Fleet, app *application, tracks []*ctTrack, mode releaseMode, cfg swarmConfig,
) {
	t.Helper()
	switch mode {
	case releaseImmediately:
		// Automatic application: nothing to reveal, the committee owns it
		// from the moment each ciphertext lands.
		for _, tr := range tracks {
			tr.RevealBlock = tr.SubmitBlock
		}
	case releaseWithheld:
		return
	case releaseDelayed:
		if len(tracks) == 0 {
			return
		}
		if _, err := f.waitBlock(ctx, tracks[len(tracks)-1].SubmitBlock+cfg.delayBlocks); err != nil {
			t.Error(err)
			return
		}
		out, err := f.releaseSecret(ctx, app)
		if !expectOK(t, app.Organizer.Label+"/reveal", "revealOrganizerSecret", out, err, "") {
			for _, tr := range tracks {
				tr.Note = "reveal failed: " + err.Error()
			}
			return
		}
		for _, tr := range tracks {
			tr.RevealBlock = out.Block
			tr.RevealGas = out.Gas
		}
	}
}

// judgeCiphertext waits for the verdict of one ciphertext and records it.
func judgeCiphertext(ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, tr *ctTrack, cfg swarmConfig) {
	t.Helper()
	from := scanFrom(epoch)
	if tr.Mode == releaseWithheld {
		judgeWithheld(ctx, t, f, from, tr, cfg)
		return
	}
	if tr.RevealBlock == 0 {
		recordTrack(ctx, t, f, tr, "organizer secret never landed (reveal failed)")
		return
	}
	ev, ok, err := f.waitCombine(ctx, tr.Epoch, tr.Aid, tr.Idx, from, tr.RevealBlock+cfg.combineWait)
	tr.Combine = ev
	tr.Partials, tr.PartialGas = partialStats(ctx, f, tr, from)
	switch {
	case err != nil:
		tr.Note = "wait combine: " + err.Error()
	case !ok:
		tr.Note = fmt.Sprintf("not combined within %d blocks of the reveal", cfg.combineWait)
	case ev.Plaintext.Cmp(tr.Plaintext) != 0:
		tr.Note = fmt.Sprintf("plaintext mismatch: want %s got %s", tr.Plaintext, ev.Plaintext)
	default:
		tr.Pass = true
	}
	recordTrack(ctx, t, f, tr, tr.Note)
}

func judgeWithheld(ctx context.Context, t *testing.T, f *Fleet, from uint64, tr *ctTrack, cfg swarmConfig) {
	t.Helper()
	if _, err := f.waitBlock(ctx, tr.SubmitBlock+cfg.withheldWait); err != nil {
		t.Error(err)
		return
	}
	rec, err := f.Services.Contracts.GetCombinedDecryption(ctx, tr.Epoch, tr.Aid, tr.Idx)
	tr.Partials, tr.PartialGas = partialStats(ctx, f, tr, from)
	switch {
	case err != nil:
		tr.Note = "read combine record: " + err.Error()
	case rec.Completed:
		tr.Note = fmt.Sprintf("COMBINED WITHOUT THE ORGANIZER SECRET (plaintext %s)", rec.Plaintext)
	default:
		tr.Pass = true
		tr.Note = fmt.Sprintf("not combined after %d blocks, %d partials on chain", cfg.withheldWait, tr.Partials)
	}
	recordTrack(ctx, t, f, tr, tr.Note)
}

// partialStats counts the partials on chain for a slot and averages their
// gas from the receipts; -1 partials when the scan fails, 0 gas when a
// receipt is missing.
func partialStats(ctx context.Context, f *Fleet, tr *ctTrack, from uint64) (int, uint64) {
	evs, err := f.partials(ctx, tr.Epoch, tr.Aid, tr.Idx, from)
	if err != nil {
		return -1, 0
	}
	if len(evs) == 0 {
		return 0, 0
	}
	var gas uint64
	for _, ev := range evs {
		receipt, err := f.Services.Contracts.Client().TransactionReceipt(ctx, ev.Tx)
		if err != nil {
			return len(evs), 0
		}
		gas += receipt.GasUsed
	}
	return len(evs), gas / uint64(len(evs))
}

// recordTrack emits the per-ciphertext report row, preceded by one row for
// the nodes' partials so their gas lands in the digest.
func recordTrack(ctx context.Context, t *testing.T, f *Fleet, tr *ctTrack, note string) {
	t.Helper()
	if tr.Partials > 0 && tr.PartialGas > 0 {
		record(t, Result{
			Step: tr.label() + "/partials", Kind: "submitPartialDecryption(node)", Gas: tr.PartialGas,
			Block: tr.SubmitBlock, Pass: true, Notes: fmt.Sprintf("partials=%d avgGas=%d", tr.Partials, tr.PartialGas),
		})
	}
	res := Result{
		Step: tr.label(), Kind: "ciphertext", Block: tr.SubmitBlock, Pass: tr.Pass,
		Notes: fmt.Sprintf("partials=%d submitGas=%d revealGas=%d %s", tr.Partials, tr.SubmitGas, tr.RevealGas, note),
	}
	if tr.Combine != nil && tr.Combine.Block > 0 {
		res.Kind = "combineDecryption(node)"
		res.Tx = tr.Combine.Tx.Hex()
		res.Gas = tr.Combine.Gas
		res.LatencyBlocks = int64(tr.Combine.Block) - int64(tr.SubmitBlock)
		res.LatencySeconds = f.blocksToSeconds(ctx, tr.SubmitBlock, tr.Combine.Block)
		res.Notes += fmt.Sprintf(" combineBlock=%d combiner=%s", tr.Combine.Block, tr.Combine.Sender.Hex())
	}
	record(t, res)
	if !tr.Pass {
		t.Errorf("%s: %s", tr.label(), note)
	}
}

// summarizeSwarm aggregates the per-ciphertext rows into throughput and
// latency figures and flags anomalies (partials ≠ t, boundary crossings).
// Tracks carry their own epoch's t, n and boundary, since a swarm that
// spilled over the pool of one epoch spans several.
func summarizeSwarm(ctx context.Context, t *testing.T, f *Fleet, all []*ctTrack, started time.Time) {
	t.Helper()
	var released, combined, withheld, withheldOK, overT, overN, crossed int
	var firstSubmit, lastCombine uint64
	var latBlocks int64
	var latSeconds float64
	epochs := map[[12]byte]struct{}{}
	for _, tr := range all {
		epochs[tr.Epoch] = struct{}{}
		if firstSubmit == 0 || tr.SubmitBlock < firstSubmit {
			firstSubmit = tr.SubmitBlock
		}
		if tr.Partials > tr.Threshold {
			overT++
		}
		if tr.Partials > tr.Committee {
			overN++
		}
		if tr.Mode == releaseWithheld {
			withheld++
			if tr.Pass {
				withheldOK++
			}
			continue
		}
		released++
		if tr.Combine == nil || tr.Combine.Block == 0 {
			continue
		}
		combined++
		lastCombine = max(lastCombine, tr.Combine.Block)
		latBlocks += int64(tr.Combine.Block) - int64(tr.SubmitBlock)
		latSeconds += f.blocksToSeconds(ctx, tr.SubmitBlock, tr.Combine.Block)
		if tr.Combine.Block >= tr.ServiceEnd {
			crossed++
		}
	}
	note := fmt.Sprintf("epochs=%d released=%d combined=%d withheld=%d withheldNotCombined=%d partials>t=%d partials>n=%d "+
		"combinedAfterEpochBoundary=%d wall=%.0fs", len(epochs), released, combined, withheld, withheldOK, overT, overN,
		crossed, time.Since(started).Seconds())
	res := Result{
		Step: "summary", Kind: "measure", Notes: note,
		Pass: combined == released && withheldOK == withheld && overN == 0,
	}
	if combined > 0 {
		span := f.blocksToSeconds(ctx, firstSubmit, lastCombine)
		res.LatencyBlocks = latBlocks / int64(combined)
		res.LatencySeconds = latSeconds / float64(combined)
		res.Notes += fmt.Sprintf(" avgLatency=%db/%.1fs throughput=%.3f ct/s (%d ct over %d blocks, %.0fs)",
			res.LatencyBlocks, res.LatencySeconds, float64(combined)/max(span, 1), combined, lastCombine-firstSubmit, span)
	}
	record(t, res)
	if overN > 0 {
		t.Errorf("swarm: %d ciphertexts collected more partials than the committee size", overN)
	}
}
