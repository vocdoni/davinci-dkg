package node

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/finalizer"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// refusingContribChain is fakeContribChain behind an endpoint that refuses
// eth_getLogs ranges starting more than window blocks behind the head with
// publicnode's HTTP 403, as the pool of a node parked on it did.
type refusingContribChain struct {
	*fakeContribChain
	window  uint64
	refused int
}

func (f *refusingContribChain) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	if q.FromBlock.Uint64()+f.window < f.head {
		f.refused++
		return nil, rpc.HTTPError{
			StatusCode: http.StatusForbidden, Status: "403 Forbidden",
			Body: []byte(`{"code":-32602,"message":"Archive requests require a personal token."}`),
		}
	}
	return f.fakeContribChain.FilterLogs(ctx, q)
}

const (
	liveContributedAt = 150    // block the fixture's contributions landed in
	liveLogWindow     = 10_000 // how far back the fixture's endpoint serves logs
)

// dealerKey signs the fixture's contribution transactions (a JSON-RPC
// client refuses unsigned ones).
var dealerKey, _ = ethcrypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")

// liveEpoch is a Live epoch of a 3-member committee (t = 2) whose dealers
// contributed at liveContributedAt, served with its logs still recent, and
// the member myIdx whose shares come from it.
type liveEpoch struct {
	id       [12]byte
	epoch    epochView
	selected []common.Address
	chain    *refusingContribChain
	manager  *gtypes.DKGManager
	genuine  map[common.Address][]byte
	words    map[uint16][]*big.Int
	pis      map[uint16]*contribution.PublicInputs
	coeffs   map[uint16][][]*big.Int
	layout   contribution.Layout
	secret   *big.Int
	myIdx    uint16
}

func newLiveEpoch(t *testing.T, id [12]byte, myIdx uint16, dealers ...uint16) *liveEpoch {
	t.Helper()
	layout, err := contribution.NewLayout(2, 3)
	qt.Assert(t, err, qt.IsNil)
	l := &liveEpoch{
		id: id,
		epoch: epochView{
			Status: epochLive, Policy: web3.EpochPolicy{Threshold: 2, CommitteeSize: 3},
			SeedBlock: 100, ContributionCount: uint16(len(dealers)),
		},
		selected: []common.Address{
			common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2)), common.BigToAddress(big.NewInt(3)),
		},
		chain: &refusingContribChain{
			fakeContribChain: &fakeContribChain{
				head:    liveContributedAt + 50,
				txs:     map[common.Hash]*ethtypes.Transaction{},
				records: map[common.Address]gtypes.DKGTypesContributionRecord{},
			},
			window: liveLogWindow,
		},
		genuine: map[common.Address][]byte{},
		words:   map[uint16][]*big.Int{},
		pis:     map[uint16]*contribution.PublicInputs{},
		coeffs:  map[uint16][][]*big.Int{},
		layout:  layout,
		secret:  big.NewInt(424242),
		myIdx:   myIdx,
	}
	for _, dealerIdx := range dealers {
		dealer := l.selected[dealerIdx-1]
		pi, words, coeffs := dealtContribution(t, id, layout, dealerIdx, myIdx, l.secret)
		l.pis[dealerIdx], l.words[dealerIdx], l.coeffs[dealerIdx] = pi, words, coeffs
		l.genuine[dealer] = packContribution(t, id, dealerIdx, pi, words)
		tx, err := ethtypes.SignTx(ethtypes.NewTx(&ethtypes.LegacyTx{Nonce: uint64(dealerIdx), Data: l.genuine[dealer]}),
			ethtypes.HomesteadSigner{}, dealerKey)
		qt.Assert(t, err, qt.IsNil)
		l.chain.txs[tx.Hash()] = tx
		l.chain.records[dealer] = gtypes.DKGTypesContributionRecord{
			Contributor: dealer, ContributorIndex: dealerIdx, CommitmentsHash: common.BigToHash(pi.CommitmentHash), Accepted: true,
		}
		l.chain.logs = append(l.chain.logs, contributionLog(t, id, dealer, liveContributedAt, tx))
	}
	l.manager, err = gtypes.NewDKGManager(common.Address{}, l.chain)
	qt.Assert(t, err, qt.IsNil)
	return l
}

// node is a freshly started process of member myIdx: no private share and
// none of its own polynomials in memory, only the datadir's cache.
func (l *liveEpoch) node(cacheDir string) *Node {
	return &Node{
		address: l.selected[l.myIdx-1], bjjSecret: l.secret, manager: l.manager,
		privateShares: map[poolSlot]*big.Int{}, ownContribs: map[[12]byte]*savedContrib{},
		selectedCache: map[[12]byte][]common.Address{l.id: l.selected},
		terminal:      map[[12]byte]bool{},
		retain:        map[[12]byte]retention{},
		keepRetry:     map[[12]byte]*serviceBackoff{},
		contribCache:  &contributionCache{dir: cacheDir},
	}
}

func (l *liveEpoch) retention() retention {
	return retention{epoch: l.epoch, selected: l.selected, myIdx: l.myIdx}
}

// logReads is every eth_getLogs the fixture's endpoint was sent, served or refused.
func (l *liveEpoch) logReads() int { return len(l.chain.queries) + l.chain.refused }

// share is member myIdx's expected share of pool key key: Σ_c f_{c,key}(myIdx).
func (l *liveEpoch) share(t *testing.T, key uint8) *big.Int {
	t.Helper()
	want := new(big.Int)
	for _, coeffs := range l.coeffs {
		s, err := ccommon.EvaluatePolynomialNative(coeffs[key], big.NewInt(int64(l.myIdx)))
		qt.Assert(t, err, qt.IsNil)
		want.Add(want, s)
	}
	return want.Mod(want, group.ScalarField())
}

// A member builds its share of a pool key the first time an application
// claims the key, often a day after the epoch went Live, when the
// ContributionSubmitted logs that locate the dealers' calldata are older
// than many endpoints serve. Every accepted dealer's calldata, the node's own
// included, is kept on disk while the logs are recent (keepContributions),
// so that later build, even after a restart that lost the node's own
// polynomials, sends no eth_getLogs at all. Keeping checks every pool key's
// share, so an entry wrong for any key is replaced while the logs are recent.
func TestShareBuildsFromKeptContributionsAfterTheLogWindow(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	const key = uint8(3)
	l := newLiveEpoch(t, [12]byte{7}, 2, 1, 2, 3)
	// Dealer 1's calldata with the last key's share to us flipped.
	flipped := slices.Clone(l.words[1])
	q := l.layout.MaskedShareOffset(contribution.MaxKeys-1, int(l.myIdx)-1)
	flipped[q] = new(big.Int).Xor(flipped[q], big.NewInt(1))
	poisoned := packContribution(t, l.id, 1, l.pis[1], flipped)

	dir := filepath.Join(t.TempDir(), "contributions")
	n := l.node(dir)
	n.contribCache.Put(l.id, l.selected[0], poisoned)

	_, done, err := n.keepContributions(ctx, l.chain, l.id, l.retention(), len(l.selected))
	c.Assert(err, qt.IsNil)
	c.Assert(done, qt.IsTrue)
	for _, dealer := range l.selected {
		got, ok := n.contribCache.Get(l.id, dealer)
		c.Assert(ok, qt.IsTrue, qt.Commentf("dealer %s not kept", dealer))
		c.Assert(bytes.Equal(got, l.genuine[dealer]), qt.IsTrue, qt.Commentf("dealer %s", dealer))
	}

	// A day later no endpoint serves those logs any more.
	l.chain.head = liveContributedAt + 5*liveLogWindow
	_, err = finalizer.ContributionCalldata(ctx, l.chain, l.manager, l.id, l.selected[0], l.epoch.SeedBlock)
	c.Assert(err, qt.ErrorMatches, ".*403 Forbidden.*")
	reads := l.logReads()

	share, err := l.node(dir).buildPrivateShare(ctx, l.chain, l.id, key, l.myIdx, l.selected, l.epoch)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(l.share(t, key)), qt.Equals, 0)
	c.Assert(l.logReads(), qt.Equals, reads, qt.Commentf("the build must send no eth_getLogs"))

	// Keeping them again (a restart while the epoch is Live) checks the
	// cache and sends nothing either.
	fetched, done, err := l.node(dir).keepContributions(ctx, l.chain, l.id, l.retention(), 0)
	c.Assert(err, qt.IsNil)
	c.Assert(done, qt.IsTrue)
	c.Assert(fetched, qt.Equals, 0)
	c.Assert(l.logReads(), qt.Equals, reads)

	// Without the kept calldata the build fails on the refused range: the incident.
	_, err = l.node(filepath.Join(t.TempDir(), "contributions")).buildPrivateShare(ctx, l.chain, l.id, key, l.myIdx, l.selected, l.epoch)
	c.Assert(err, qt.ErrorMatches, `recover share from .*: filter ContributionSubmitted .*403 Forbidden.*`)
}

// Retention runs after the tick's other work, fetches at most
// keepFetchesPerTick dealers per tick and resumes on the next one, and does
// nothing at all while an epoch may still need a slot claim, a contribution
// or a finalization, however slow history reads are.
func TestRetentionYieldsToKeyGenerationAndKeepsOneDealerPerTick(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4}, 2, 1, 2, 3)
	n := l.node(filepath.Join(t.TempDir(), "contributions"))
	n.retain[l.id] = l.retention()
	head := l.chain.head

	assembling := [12]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 5}
	active := &tickCtx{head: head, epochs: map[[12]byte]epochView{assembling: {
		Status: epochKeyAssembly,
		Policy: web3.EpochPolicy{CommitteeSize: 3, LiveNotBeforeBlock: head + 10},
	}}}
	n.retainContributions(ctx, active, l.chain)
	c.Assert(l.logReads(), qt.Equals, 0, qt.Commentf("retention must yield to an assembling epoch"))
	c.Assert(n.retain, qt.HasLen, 1)

	selecting := &tickCtx{head: head, epochs: map[[12]byte]epochView{assembling: {
		Status: epochCommitteeSelection, Policy: web3.EpochPolicy{CommitteeSelectionDeadlineBlock: head + 1},
	}}}
	n.retainContributions(ctx, selecting, l.chain)
	c.Assert(l.logReads(), qt.Equals, 0, qt.Commentf("retention must yield to a selecting epoch"))

	// Past the rotation's last slot (or for an epoch this node is done
	// with) the step works again, one dealer per tick.
	stale := &tickCtx{head: head, epochs: map[[12]byte]epochView{assembling: {
		Status: epochKeyAssembly, Policy: web3.EpochPolicy{CommitteeSize: 3, LiveNotBeforeBlock: head - 100},
	}}}
	for tick := 1; tick <= 3; tick++ {
		n.retainContributions(ctx, stale, l.chain)
		kept := 0
		for _, dealer := range l.selected {
			if _, ok := n.contribCache.Get(l.id, dealer); ok {
				kept++
			}
		}
		c.Assert(kept, qt.Equals, tick, qt.Commentf("tick %d", tick))
	}
	c.Assert(n.retain, qt.HasLen, 0, qt.Commentf("a fully kept epoch leaves the queue"))
	reads := l.logReads()
	n.retainContributions(ctx, &tickCtx{head: head}, l.chain)
	c.Assert(l.logReads(), qt.Equals, reads)
}

// A failed attempt keeps the epoch queued and backs it off: the next tick
// is skipped without a request, and a later one finishes the work.
func TestRetentionBacksOffAfterAFailure(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{9}, 1, 2)
	l.chain.head = liveContributedAt + 5*liveLogWindow // a restart long after Live, empty cache
	n := l.node(filepath.Join(t.TempDir(), "contributions"))
	n.retain[l.id] = l.retention()
	tc := &tickCtx{head: l.chain.head}

	n.retainContributions(ctx, tc, l.chain)
	c.Assert(n.retain, qt.HasLen, 1)
	c.Assert(n.keepRetry[l.id], qt.Not(qt.IsNil))
	reads := l.logReads()
	n.retainContributions(ctx, tc, l.chain)
	c.Assert(l.logReads(), qt.Equals, reads, qt.Commentf("a backed-off attempt sends nothing"))

	l.chain.head = liveContributedAt + 50 // an endpoint that serves the range
	n.retainContributions(ctx, tc, l.chain)
	c.Assert(n.retain, qt.HasLen, 0)
	c.Assert(n.keepRetry, qt.HasLen, 0)
	_, ok := n.contribCache.Get(l.id, l.selected[1])
	c.Assert(ok, qt.IsTrue)
}

// fakeCommittee answers selectedParticipants with a fixed list, as an
// endpoint that is up to date (the full committee) or one slot claim behind.
type fakeCommittee struct{ members []common.Address }

func (f fakeCommittee) SelectedParticipants(context.Context, [12]byte) ([]common.Address, error) {
	return slices.Clone(f.members), nil
}

// A fallback one slot claim behind the chain answers selectedParticipants
// with a committee short of its last member, here this node. That answer is
// neither cached nor read as "not a member" (which would finish the epoch
// and drop its ciphertexts for good): it is an error to retry, and the full
// committee from an endpoint that caught up is used once it arrives.
func TestIncompleteCommitteeIsRetriedNotCached(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{3}, 3, 1, 2, 3)
	n := l.node(filepath.Join(t.TempDir(), "contributions"))
	delete(n.selectedCache, l.id)
	behind := fakeCommittee{members: l.selected[:2]}

	_, err := n.selected(ctx, behind, l.id, l.epoch.Policy.CommitteeSize)
	c.Assert(err, qt.ErrorMatches, `committee has 2 of its 3 members.*`)
	c.Assert(n.selectedCache, qt.HasLen, 0)
	c.Assert(n.queueRetention(ctx, behind, l.id, l.epoch), qt.IsNotNil)
	c.Assert(n.retain, qt.HasLen, 0)

	full := fakeCommittee{members: l.selected}
	got, err := n.selected(ctx, full, l.id, l.epoch.Policy.CommitteeSize)
	c.Assert(err, qt.IsNil)
	c.Assert(got, qt.DeepEquals, l.selected)
	c.Assert(n.queueRetention(ctx, full, l.id, l.epoch), qt.IsNil)
	c.Assert(n.retain[l.id].myIdx, qt.Equals, uint16(3))
}

// An endpoint behind the chain answers getContribution for a dealer the Live
// epoch already counts as accepted (ContributionCount 3) with a record not
// yet accepted. Neither retention nor a share build may take the two
// accepted records it shows for the accepted set: retention would finish
// with a dealer missing, and a share summed over a subset is wrong, so every
// partial made with it reverts. Both fail, to be retried, and keep nothing.
func TestIncompleteAcceptedSetIsRetried(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{4}, 2, 1, 2, 3)
	late := l.selected[2]
	accepted := l.chain.records[late]
	stale := accepted
	stale.Accepted = false
	l.chain.records[late] = stale
	n := l.node(filepath.Join(t.TempDir(), "contributions"))
	n.retain[l.id] = l.retention()

	_, done, err := n.keepContributions(ctx, l.chain, l.id, l.retention(), len(l.selected))
	c.Assert(err, qt.ErrorMatches, `2 accepted contributions, the epoch has 3.*`)
	c.Assert(done, qt.IsFalse)
	n.retainContributions(ctx, &tickCtx{head: l.chain.head}, l.chain)
	c.Assert(n.retain, qt.HasLen, 1, qt.Commentf("an incomplete set must leave the epoch queued"))

	_, err = n.buildPrivateShare(ctx, l.chain, l.id, 0, l.myIdx, l.selected, l.epoch)
	c.Assert(err, qt.ErrorMatches, `2 accepted contributions, the epoch has 3.*`)
	_, err = n.buildPrivateShare(ctx, l.chain, l.id, 0, l.myIdx, l.selected[:2], l.epoch)
	c.Assert(err, qt.ErrorMatches, `committee has 2 of its 3 members.*`)
	c.Assert(n.privateShares, qt.HasLen, 0, qt.Commentf("no share may be cached from a subset"))

	// Once the endpoint has caught up, both work.
	l.chain.records[late] = accepted
	share, err := n.buildPrivateShare(ctx, l.chain, l.id, 0, l.myIdx, l.selected, l.epoch)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(l.share(t, 0)), qt.Equals, 0)
	_, done, err = n.keepContributions(ctx, l.chain, l.id, l.retention(), len(l.selected))
	c.Assert(err, qt.IsNil)
	c.Assert(done, qt.IsTrue)
}

// A cache write that fails (disk full, permissions, a file where the epoch's
// directory should be) is not retention: the epoch stays queued under the
// retry backoff and is kept once storage works again. A share build only
// uses the cache, so the same failure does not stop it.
func TestRetentionStaysQueuedUntilTheCacheWriteSucceeds(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{5}, 1, 2)
	dir := filepath.Join(t.TempDir(), "contributions")
	c.Assert(os.MkdirAll(dir, 0o700), qt.IsNil)
	obstruction := filepath.Join(dir, hex.EncodeToString(l.id[:]))
	c.Assert(os.WriteFile(obstruction, []byte("not a directory"), 0o600), qt.IsNil)
	n := l.node(dir)
	n.retain[l.id] = l.retention()
	tc := &tickCtx{head: l.chain.head}

	n.retainContributions(ctx, tc, l.chain)
	c.Assert(n.retain, qt.HasLen, 1, qt.Commentf("a failed write must leave the epoch queued"))
	c.Assert(n.keepRetry[l.id], qt.Not(qt.IsNil))
	share, err := n.buildPrivateShare(ctx, l.chain, l.id, 0, l.myIdx, l.selected, l.epoch)
	c.Assert(err, qt.IsNil, qt.Commentf("a share build does not depend on the cache write"))
	c.Assert(share.Cmp(l.share(t, 0)), qt.Equals, 0)

	c.Assert(os.Remove(obstruction), qt.IsNil)
	for range 3 { // the backoff skips a tick
		n.retainContributions(ctx, tc, l.chain)
	}
	c.Assert(n.retain, qt.HasLen, 0)
	_, ok := n.contribCache.Get(l.id, l.selected[1])
	c.Assert(ok, qt.IsTrue)
}

// countingChain counts the eth_calls a retention pass sends (the contribution
// record reads) and can fail them, or hold every eth_getLogs until the
// caller gives up, as a primary that times out on old ranges does.
type countingChain struct {
	*refusingContribChain
	calls     atomic.Int32
	failCalls bool
	holdLogs  bool
}

func (f *countingChain) CallContract(ctx context.Context, call ethereum.CallMsg, block *big.Int) ([]byte, error) {
	f.calls.Add(1)
	if f.failCalls {
		return nil, errors.New("read tcp 10.0.0.2:41234->10.0.0.1:443: read: connection reset by peer")
	}
	return f.refusingContribChain.CallContract(ctx, call, block)
}

func (f *countingChain) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	if f.holdLogs {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
	return f.refusingContribChain.FilterLogs(ctx, q)
}

// A retention pass is bounded in elapsed time, record reads and fetches
// included; it is charged per attempt, so queued epochs whose reads keep
// failing cost one attempt per tick, not one each; and it does not run in a
// tick whose epoch discovery failed, which the next tick retries first.
func TestRetentionPassIsBoundedAndChargedPerAttempt(t *testing.T) {
	ctx := context.Background()
	newQueue := func(t *testing.T) (*Node, *countingChain, *liveEpoch) {
		l := newLiveEpoch(t, [12]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}, 2, 1, 2, 3)
		chain := &countingChain{refusingContribChain: l.chain}
		m, err := gtypes.NewDKGManager(common.Address{}, chain)
		qt.Assert(t, err, qt.IsNil)
		l.manager = m
		n := l.node(filepath.Join(t.TempDir(), "contributions"))
		n.keepBudget = 100 * time.Millisecond
		for nonce := byte(1); nonce <= 3; nonce++ {
			id := [12]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, nonce}
			n.retain[id] = l.retention()
			n.selectedCache[id] = l.selected
		}
		return n, chain, l
	}

	t.Run("failed discovery", func(t *testing.T) {
		c := qt.New(t)
		n, chain, _ := newQueue(t)
		n.retainContributions(ctx, &tickCtx{head: 200, failed: true}, chain)
		c.Assert(chain.calls.Load(), qt.Equals, int32(0), qt.Commentf("no retention in a tick whose discovery failed"))
	})
	t.Run("failing record reads", func(t *testing.T) {
		c := qt.New(t)
		n, chain, _ := newQueue(t)
		chain.failCalls = true
		n.retainContributions(ctx, &tickCtx{head: 200}, chain)
		c.Assert(chain.calls.Load(), qt.Equals, int32(3), qt.Commentf("one epoch's record reads (3 members) per pass"))
		c.Assert(n.retain, qt.HasLen, 3)
	})
	t.Run("slow history", func(t *testing.T) {
		c := qt.New(t)
		n, chain, _ := newQueue(t)
		chain.holdLogs = true
		start := time.Now()
		n.retainContributions(ctx, &tickCtx{head: 200}, chain)
		c.Assert(time.Since(start) < 2*time.Second, qt.IsTrue, qt.Commentf("the pass took %s", time.Since(start)))
		c.Assert(n.retain, qt.HasLen, 3)
	})
}

// serveLiveEpoch serves a liveEpoch's chain over JSON-RPC, single requests
// and batches (eth_blockNumber, eth_call, eth_getLogs,
// eth_getTransactionByHash), for a real pool to point at. stallOld makes it
// hold every eth_getLogs starting more than liveLogWindow blocks behind the
// head until the client gives up, as a primary that times out on old ranges
// while it answers everything else promptly.
func serveLiveEpoch(t *testing.T, l *liveEpoch, stallOld bool) string {
	t.Helper()
	answer := func(r *http.Request, method string, params []json.RawMessage) (any, error) {
		ctx := r.Context()
		switch method {
		case "eth_blockNumber":
			return hexutil.EncodeUint64(l.chain.head), nil
		case "eth_call":
			var call struct {
				Input hexutil.Bytes `json:"input"`
				Data  hexutil.Bytes `json:"data"`
			}
			if err := json.Unmarshal(params[0], &call); err != nil {
				return nil, err
			}
			out, err := l.chain.CallContract(ctx, ethereum.CallMsg{Data: append(call.Data, call.Input...)}, nil)
			return hexutil.Bytes(out), err
		case "eth_getLogs":
			var filter struct {
				FromBlock hexutil.Uint64  `json:"fromBlock"`
				ToBlock   hexutil.Uint64  `json:"toBlock"`
				Topics    [][]common.Hash `json:"topics"`
			}
			if err := json.Unmarshal(params[0], &filter); err != nil {
				return nil, err
			}
			if stallOld && uint64(filter.FromBlock)+liveLogWindow < l.chain.head {
				<-ctx.Done()
				return nil, ctx.Err()
			}
			return l.chain.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: new(big.Int).SetUint64(uint64(filter.FromBlock)),
				ToBlock:   new(big.Int).SetUint64(uint64(filter.ToBlock)),
				Topics:    filter.Topics,
			})
		case "eth_getTransactionByHash":
			var hash common.Hash
			if err := json.Unmarshal(params[0], &hash); err != nil {
				return nil, err
			}
			tx, _, err := l.chain.TransactionByHash(ctx, hash)
			if err != nil {
				return nil, err
			}
			raw, err := tx.MarshalJSON()
			if err != nil {
				return nil, err
			}
			var fields map[string]any
			if err := json.Unmarshal(raw, &fields); err != nil {
				return nil, err
			}
			fields["blockNumber"] = hexutil.EncodeUint64(liveContributedAt)
			fields["blockHash"] = common.Hash{1}
			return fields, nil
		}
		return nil, fmt.Errorf("%s not served", method)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		type request struct {
			ID     json.RawMessage   `json:"id"`
			Method string            `json:"method"`
			Params []json.RawMessage `json:"params"`
		}
		batch := len(body) > 0 && body[0] == '['
		var reqs []request
		if batch {
			_ = json.Unmarshal(body, &reqs)
		} else {
			var one request
			_ = json.Unmarshal(body, &one)
			reqs = []request{one}
		}
		resps := make([]map[string]any, len(reqs))
		for i, req := range reqs {
			resps[i] = map[string]any{"jsonrpc": "2.0", "id": req.ID}
			if result, err := answer(r, req.Method, req.Params); err != nil {
				resps[i]["error"] = map[string]any{"code": -32000, "message": err.Error()}
			} else {
				resps[i]["result"] = result
			}
		}
		if r.Context().Err() != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if batch {
			_ = json.NewEncoder(w).Encode(resps)
			return
		}
		_ = json.NewEncoder(w).Encode(resps[0])
	}))
	t.Cleanup(srv.Close)
	return srv.URL
}

// A primary that answers everything promptly but stalls on old log ranges,
// beside a healthy secondary: each attempt of a retention pass gets a share
// of the pass's budget, a primary that runs out of it is skipped for the
// cooldown, and the cache fills within a few passes rather than never.
func TestRetentionMovesOffAPrimaryThatStallsOnOldLogs(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	l := newLiveEpoch(t, [12]byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 6}, 2, 1, 2, 3)
	l.chain.head = liveContributedAt + 5*liveLogWindow // the contributions are old
	l.chain.window = l.chain.head                      // nothing refuses: the primary stalls
	pool, err := web3.NewRPCPool([]string{serveLiveEpoch(t, l, true), serveLiveEpoch(t, l, false)})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := web3.NewPooledBackend(pool)
	l.manager, err = gtypes.NewDKGManager(common.Address{}, backend)
	c.Assert(err, qt.IsNil)
	n := l.node(filepath.Join(t.TempDir(), "contributions"))
	n.keepBudget = 5 * time.Second
	n.retain[l.id] = l.retention()

	start := time.Now()
	for pass := 0; pass < 6 && len(n.retain) > 0; pass++ {
		n.retainContributions(ctx, &tickCtx{head: l.chain.head}, backend)
	}
	c.Assert(n.retain, qt.HasLen, 0, qt.Commentf("retention must finish through the healthy secondary"))
	for _, dealer := range l.selected {
		got, ok := n.contribCache.Get(l.id, dealer)
		c.Assert(ok, qt.IsTrue)
		c.Assert(bytes.Equal(got, l.genuine[dealer]), qt.IsTrue)
	}
	c.Assert(time.Since(start) < 15*time.Second, qt.IsTrue, qt.Commentf("took %s", time.Since(start)))
}
