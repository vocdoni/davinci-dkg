package battery

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// releaseMode is how an organizer treats one ciphertext's share.
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
	plaintextBits  uint
}

func loadSwarmConfig() swarmConfig {
	return swarmConfig{
		organizers:     envInt("BATTERY_ORGANIZERS", 8),
		perOrganizer:   envInt("BATTERY_CIPHERTEXTS", 6),
		delayBlocks:    envUint64("BATTERY_SHARE_DELAY_BLOCKS", 6),
		withheldWait:   envUint64("BATTERY_WITHHELD_WAIT_BLOCKS", 40),
		combineWait:    envUint64("BATTERY_COMBINE_WAIT_BLOCKS", 240),
		minServiceLeft: envUint64("BATTERY_MIN_SERVICE_BLOCKS", 90),
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
	ShareBlock  uint64
	ShareGas    uint64
	Combine     *combineEvent
	Partials    int
	Pass        bool
	Note        string
}

func (c *ctTrack) label() string { return fmt.Sprintf("%s/ct%d(%s)", c.Organizer, c.Idx, c.Mode) }

// TestOrganizerSwarm runs K organizers concurrently inside one Live epoch,
// each registering an application and submitting M ciphertexts whose
// organizer shares are released immediately, after a delay, or withheld.
// Every released ciphertext must combine to its plaintext; withheld ones
// must not combine. Per-ciphertext latency, partial counts and gas are
// reported; an epoch boundary crossed mid-run is reported, not fatal.
func TestOrganizerSwarm(t *testing.T) {
	f := requireFleet(t)
	ctx, cancel := testContext(t)
	defer cancel()
	cfg := loadSwarmConfig()

	epochID, epoch, err := f.waitLiveEpoch(ctx, t, cfg.minServiceLeft)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("swarm: epoch %x t=%d n=%d serviceEnd=%d organizers=%d ciphertexts/organizer=%d",
		epochID, epoch.Policy.Threshold, epoch.Policy.CommitteeSize, f.serviceEnd(epoch), cfg.organizers, cfg.perOrganizer)

	actors, err := f.newActors(ctx, "org", cfg.organizers)
	if err != nil {
		t.Fatal(err)
	}

	started := time.Now()
	tracks := make([][]*ctTrack, cfg.organizers)
	var wg sync.WaitGroup
	for i, a := range actors {
		wg.Add(1)
		go func(i int, a *actor) {
			defer wg.Done()
			tracks[i] = runOrganizer(ctx, t, f, a, epochID, epoch, cfg)
		}(i, a)
	}
	wg.Wait()

	var all []*ctTrack
	for _, ts := range tracks {
		all = append(all, ts...)
	}
	summarizeSwarm(ctx, t, f, epoch, all, started)
}

// runOrganizer is one organizer's whole life: register, submit M
// ciphertexts, release shares per plan, then judge every ciphertext.
func runOrganizer(
	ctx context.Context, t *testing.T, f *Fleet, a *actor, epochID [12]byte, epoch web3.EpochView, cfg swarmConfig,
) []*ctTrack {
	t.Helper()
	app, out, err := f.registerApplication(ctx, a, epochID, golangtypes.DKGTypesAppPolicy{})
	if !expectOK(t, a.Label+"/register", "registerApplication", out, err, "") {
		return nil
	}
	tracks := submitSwarmCiphertexts(ctx, t, f, a, app, cfg)
	releaseSwarmShares(ctx, t, f, app, tracks, cfg)
	for _, tr := range tracks {
		judgeCiphertext(ctx, t, f, epoch, tr, cfg)
	}
	return tracks
}

func submitSwarmCiphertexts(
	ctx context.Context, t *testing.T, f *Fleet, a *actor, app *application, cfg swarmConfig,
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
			Organizer: a.Label, Epoch: app.Epoch, Aid: app.Aid, Idx: idx, Mode: planFor(j), Plaintext: m,
			C1: c1, C2: c2, SubmitBlock: out.Block, SubmitGas: out.Gas,
		})
	}
	return tracks
}

// planFor spreads the release modes: every 4th ciphertext is withheld
// (~25%), the rest alternate between immediate and delayed.
func planFor(j int) releaseMode {
	switch {
	case j%4 == 3:
		return releaseWithheld
	case j%2 == 0:
		return releaseImmediately
	default:
		return releaseDelayed
	}
}

func releaseSwarmShares(ctx context.Context, t *testing.T, f *Fleet, app *application, tracks []*ctTrack, cfg swarmConfig) {
	t.Helper()
	// Immediate releases first, then the delayed ones in submission order
	// (each waits for its own submitBlock + delay).
	for _, tr := range tracks {
		if tr.Mode == releaseImmediately {
			releaseTracked(ctx, t, f, app, tr)
		}
	}
	for _, tr := range tracks {
		if tr.Mode != releaseDelayed {
			continue
		}
		if _, err := f.waitBlock(ctx, tr.SubmitBlock+cfg.delayBlocks); err != nil {
			t.Error(err)
			return
		}
		releaseTracked(ctx, t, f, app, tr)
	}
}

func releaseTracked(ctx context.Context, t *testing.T, f *Fleet, app *application, tr *ctTrack) {
	t.Helper()
	_, _, out, err := f.releaseShare(ctx, app, tr.Idx, tr.C1, tr.C2)
	if expectOK(t, tr.label()+"/share", "submitOrganizerShare", out, err, "") {
		tr.ShareBlock = out.Block
		tr.ShareGas = out.Gas
	} else {
		tr.Note = "share submission failed: " + err.Error()
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
	if tr.ShareBlock == 0 {
		recordTrack(ctx, t, f, tr, "no share on chain (release failed)")
		return
	}
	ev, ok, err := f.waitCombine(ctx, tr.Epoch, tr.Aid, tr.Idx, from, tr.ShareBlock+cfg.combineWait)
	tr.Combine = ev
	tr.Partials = countPartials(ctx, f, tr, from)
	switch {
	case err != nil:
		tr.Note = "wait combine: " + err.Error()
	case !ok:
		tr.Note = fmt.Sprintf("not combined within %d blocks of the share", cfg.combineWait)
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
	tr.Partials = countPartials(ctx, f, tr, from)
	switch {
	case err != nil:
		tr.Note = "read combine record: " + err.Error()
	case rec.Completed:
		tr.Note = fmt.Sprintf("COMBINED WITHOUT ORGANIZER SHARE (plaintext %s)", rec.Plaintext)
	default:
		tr.Pass = true
		tr.Note = fmt.Sprintf("not combined after %d blocks, %d partials on chain", cfg.withheldWait, tr.Partials)
	}
	recordTrack(ctx, t, f, tr, tr.Note)
}

func countPartials(ctx context.Context, f *Fleet, tr *ctTrack, from uint64) int {
	evs, err := f.partials(ctx, tr.Epoch, tr.Aid, tr.Idx, from)
	if err != nil {
		return -1
	}
	return len(evs)
}

// recordTrack emits the per-ciphertext report row.
func recordTrack(ctx context.Context, t *testing.T, f *Fleet, tr *ctTrack, note string) {
	t.Helper()
	res := Result{
		Step: tr.label(), Kind: "ciphertext", Block: tr.SubmitBlock, Pass: tr.Pass,
		Notes: fmt.Sprintf("partials=%d submitGas=%d shareGas=%d %s", tr.Partials, tr.SubmitGas, tr.ShareGas, note),
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
func summarizeSwarm(ctx context.Context, t *testing.T, f *Fleet, epoch web3.EpochView, all []*ctTrack, started time.Time) {
	t.Helper()
	threshold := int(epoch.Policy.Threshold)
	committee := int(epoch.Policy.CommitteeSize)
	var released, combined, withheld, withheldOK, overT, overN, crossed int
	var firstSubmit, lastCombine uint64
	var latBlocks int64
	var latSeconds float64
	for _, tr := range all {
		if firstSubmit == 0 || tr.SubmitBlock < firstSubmit {
			firstSubmit = tr.SubmitBlock
		}
		if tr.Partials > threshold {
			overT++
		}
		if tr.Partials > committee {
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
		if tr.Combine.Block >= f.serviceEnd(epoch) {
			crossed++
		}
	}
	note := fmt.Sprintf("released=%d combined=%d withheld=%d withheldNotCombined=%d partials>t=%d partials>n=%d "+
		"combinedAfterEpochBoundary=%d wall=%.0fs", released, combined, withheld, withheldOK, overT, overN, crossed,
		time.Since(started).Seconds())
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
