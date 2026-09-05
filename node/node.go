package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	mrand "math/rand"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/davinci-dkg/circuits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/crypto/feldman"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/finalizer"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// bjjKeyDomain must match tests/helpers/nodekeys.go so registry keys are consistent.
const bjjKeyDomain = "davinci-dkg/bjj-key/v1"

// EpochPhase values as stored on-chain (DKGTypes.EpochPhase).
const (
	epochCommitteeSelection uint8 = 1
	epochKeyAssembly        uint8 = 2
	epochLive               uint8 = 3
	epochAborted            uint8 = 4
	epochCompleted          uint8 = 5
)

type epochView = web3.EpochView

// savedContrib caches data from the node's own submitted contribution so it
// can compute the own-polynomial component of d_{j,i} offline. coefficients
// is indexed by pool key then coefficient: one polynomial per key.
type savedContrib struct {
	coefficients     [][]*big.Int
	recipientIndexes []uint16
	recipientKeys    []nodetypes.NodeKey
}

// poolSlot identifies one pool key of one epoch.
type poolSlot struct {
	epoch [12]byte
	key   uint8
}

// Node participates in every DKG epoch it can find on chain.
type Node struct {
	address   common.Address
	bjjSecret *big.Int

	contracts  *web3.Contracts
	manager    *gtypes.DKGManager
	appManager *gtypes.DKGAppManager
	registry   *gtypes.DKGRegistry
	txm        *txmanager.Manager
	runtimes   circuitRuntimes // the four pinned circuits, loaded once in New

	// per-epoch local state (key generation lifecycle)
	signaled      map[[12]byte]bool
	contributed   map[[12]byte]bool
	finalized     map[[12]byte]bool     // epochs whose finalization is settled (by us or someone else)
	terminal      map[[12]byte]bool     // epochs whose key-generation lifecycle needs no more work
	privateShares map[poolSlot]*big.Int // d_{j,i}, one per pool key we needed
	ownContribs   map[[12]byte]*savedContrib
	selectedCache map[[12]byte][]common.Address
	// finalizeRetry backs off a finalization attempt that failed for a
	// reason other than a lost race: the attempt costs a full Groth16 proof,
	// so a persistent fault (stale artifacts, a dealer whose calldata is
	// unreadable) must not burn one per tick, while a transient RPC hiccup
	// is retried within a tick or two.
	finalizeRetry map[[12]byte]*serviceBackoff
	// contribCache memoises validated submitContribution calldata under the
	// datadir so a rebuilt private share or finalization statement does not
	// rescan the event log for every dealer (see contribcache.go).
	contribCache *contributionCache

	// decryption service state (see decrypt.go)
	lookback    uint64
	lastCtScan  uint64
	ctSeq       uint64
	pending     map[ctKey]*ciphertext
	parked      map[ctKey]*parkedSlot // waiting for the organizer's reveal or the decryption window
	partialDone map[ctKey]bool
	served      map[ctKey]uint64 // finished slots → discovery block, until out of the re-scan window
	shareProofs map[poolSlot][][32]byte
	taints      map[taintKey]bool // applications (or submitters) that produced an undecryptable ciphertext
	backoff     map[ctKey]*serviceBackoff
	inflight    map[ctKey]inflightTx     // sent but unmined partial/combine per slot
	combineJobs map[ctKey]*combineResult // running (nil) or finished combine jobs, guarded by jobsMu
	jobsMu      sync.Mutex
	combineSem  chan struct{} // capacity 1: serialises the CPU-bound combine jobs
	critical    atomic.Int32  // > 0 while a contribution or finalization is in progress
	taintFile   string        // where taintedApps is persisted ("" disables)

	// auto-create-epoch state. autoCreateNextStart is the
	// nextEpochStartBlock() value the most recent attempt was scheduled
	// against; we skip re-scheduling for the same threshold so a single
	// jitter-delayed goroutine fires per cadence window.
	autoCreateNextStart uint64
	// autoCreateEarlyPool is the (epoch, poolNext) observation the most
	// recent early attempt was fired against, so one attempt is made per
	// registration that drains the pool (or per aborted epoch).
	autoCreateEarlyPool poolSlot
}

// New constructs a Node from the daemon config.
func New(cfg *Config) (*Node, error) {
	addrs := nodetypes.ContractAddresses{
		Manager: common.HexToAddress(cfg.resolvedManagerAddr()),
	}
	// web3.New() derives Registry and all verifier addresses from the manager's
	// public immutable fields when they are not supplied (zero address).
	c, err := web3.New(cfg.Web3.RPC, addrs)
	if err != nil {
		return nil, fmt.Errorf("web3 connect: %w", err)
	}
	txm, err := txmanager.New(c.Pool().Current, c.ChainID, cfg.PrivKey)
	if err != nil {
		return nil, fmt.Errorf("tx manager: %w", err)
	}
	txm.SetGasMultiplier(cfg.Web3.GasMultiplier)
	manager, err := gtypes.NewDKGManager(c.Addresses.Manager, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("manager binding: %w", err)
	}
	registry, err := gtypes.NewDKGRegistry(c.Addresses.Registry, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("registry binding: %w", err)
	}
	if c.Addresses.AppManager == (common.Address{}) {
		return nil, fmt.Errorf("DKGAppManager is not wired on manager %s", c.Addresses.Manager)
	}
	appManager, err := gtypes.NewDKGAppManager(c.Addresses.AppManager, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("app manager binding: %w", err)
	}

	bjjSecret, err := deriveBJJSecret(cfg.PrivKey)
	if err != nil {
		return nil, fmt.Errorf("derive bjj key: %w", err)
	}

	// Override artifact path from env if set.
	if d := os.Getenv("DAVINCI_DKG_ARTIFACTS_DIR"); d != "" {
		circuits.BaseDir = d
	}
	// Every proof this node will ever make needs the pinned release
	// artifacts; a missing file or a hash mismatch is fatal now rather than
	// at the first deadline.
	runtimes, err := loadRuntimes()
	if err != nil {
		return nil, fmt.Errorf("load circuit artifacts from %s: %w", circuits.BaseDir, err)
	}

	n := &Node{
		address:       txm.Address(),
		bjjSecret:     bjjSecret,
		contracts:     c,
		manager:       manager,
		appManager:    appManager,
		registry:      registry,
		txm:           txm,
		runtimes:      runtimes,
		signaled:      make(map[[12]byte]bool),
		contributed:   make(map[[12]byte]bool),
		finalized:     make(map[[12]byte]bool),
		terminal:      make(map[[12]byte]bool),
		privateShares: make(map[poolSlot]*big.Int),
		ownContribs:   make(map[[12]byte]*savedContrib),
		selectedCache: make(map[[12]byte][]common.Address),
		finalizeRetry: make(map[[12]byte]*serviceBackoff),
		contribCache:  &contributionCache{dir: contributionCacheDir(cfg.Datadir)},
		lookback:      cfg.DecryptLookbackBlocks,
		pending:       make(map[ctKey]*ciphertext),
		parked:        make(map[ctKey]*parkedSlot),
		partialDone:   make(map[ctKey]bool),
		served:        make(map[ctKey]uint64),
		shareProofs:   make(map[poolSlot][][32]byte),
		taints:        make(map[taintKey]bool),
		taintFile:     taintPath(cfg.Datadir),
		backoff:       make(map[ctKey]*serviceBackoff),
		inflight:      make(map[ctKey]inflightTx),
		combineJobs:   make(map[ctKey]*combineResult),
		combineSem:    make(chan struct{}, 1),
	}
	n.loadTaints()
	return n, nil
}

// deriveBJJSecret derives a BabyJubJub private scalar from an Ethereum private
// key using the same domain as tests/helpers/nodekeys.go.
//
// Derivation: poseidon2(keccak256(privKey || domain)[0:16], keccak256(...)[16:32])
// mod BJJ scalar field. Using keccak for pre-image binding and Poseidon for
// ZK-friendly output keeps the derivation compatible with in-circuit proofs.
func deriveBJJSecret(ethPrivKey string) (*big.Int, error) {
	preimage := append(common.FromHex(ethPrivKey), []byte(bjjKeyDomain)...)
	digest := ethcrypto.Keccak256(preimage)
	lo := new(big.Int).SetBytes(digest[:16])
	hi := new(big.Int).SetBytes(digest[16:])
	s, err := dkghash.HashFieldElements(lo, hi)
	if err != nil {
		return nil, fmt.Errorf("poseidon hash: %w", err)
	}
	s.Mod(s, group.ScalarField())
	if s.Sign() == 0 {
		s.SetInt64(1)
	}
	return s, nil
}

// LogStartupSnapshot emits a verbose banner describing the node's runtime
// configuration and the current on-chain state. Called once on startup so
// operators can verify at a glance that the node is pointed at the right
// network, knows the right contracts, and has found an active row.
func (n *Node) LogStartupSnapshot(ctx context.Context, cfg *Config) {
	log.Infow("==================== davinci-dkg-node startup ====================")

	// ── local configuration ──────────────────────────────────────────────
	log.Infow("config: node identity",
		"address", n.address,
		"datadir", cfg.Datadir)
	log.Infow("config: chain connection",
		"network", cfg.Web3.Network,
		"chainId", n.contracts.ChainID,
		"rpcHost", rpcHost(cfg.Web3.RPC[0]),
		"gasMultiplier", cfg.Web3.GasMultiplier)
	log.Infow("config: contracts",
		"registry", n.contracts.Addresses.Registry,
		"manager", cfg.ManagerAddr)
	log.Infow("config: participation",
		"pollInterval", cfg.PollInterval)

	// ── on-chain state ───────────────────────────────────────────────────
	callOpts := &bind.CallOpts{Context: ctx}

	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		log.Warnw("startup: failed to read chain head", "err", err)
	}

	prefix, err := n.manager.EPOCHPREFIX(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read EPOCH_PREFIX", "err", err)
	}
	epochNonce, err := n.manager.EpochNonce(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read epochNonce", "err", err)
	}
	log.Infow("chain: snapshot",
		"head", head,
		"roundPrefix", fmt.Sprintf("0x%08x", prefix),
		"epochNonce", epochNonce)

	nodeCount, err := n.registry.NodeCount(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read nodeCount", "err", err)
	}
	activeCount, err := n.registry.ActiveCount(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read activeCount", "err", err)
	}
	window, err := n.registry.INACTIVITYWINDOW(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read INACTIVITY_WINDOW", "err", err)
	}
	log.Infow("registry: snapshot",
		"nodeCount", nodeCount,
		"activeCount", activeCount,
		"inactivityWindow", window,
		"windowRemainingBlocks", window)

	// ── own registry row ─────────────────────────────────────────────────
	own, err := n.registry.GetNode(callOpts, n.address)
	if err != nil {
		log.Warnw("startup: failed to read own registry row", "err", err)
		log.Infow("==================================================================")
		return
	}
	statusLabel := "UNKNOWN"
	switch own.Status {
	case nodeStatusNone:
		statusLabel = "NONE"
	case nodeStatusActive:
		statusLabel = "ACTIVE"
	case nodeStatusInactive:
		statusLabel = "INACTIVE"
	}
	blocksSinceActive := uint64(0)
	if head > own.LastActiveBlock {
		blocksSinceActive = head - own.LastActiveBlock
	}
	log.Infow("self: registry row",
		"status", statusLabel,
		"lastActiveBlock", own.LastActiveBlock,
		"blocksSinceActive", blocksSinceActive,
		"pubX", own.PubX,
		"pubY", own.PubY)

	if own.Status == nodeStatusActive && window > 0 {
		deadline := own.LastActiveBlock + window
		var headroom int64
		if deadline >= head {
			headroom = int64(deadline - head)
		} else {
			headroom = -int64(head - deadline)
		}
		log.Infow("self: liveness budget",
			"reapDeadlineBlock", deadline,
			"blocksUntilReap", headroom)
	}

	// ── wallet funds ─────────────────────────────────────────────────────
	n.logFunds(ctx)

	log.Infow("==================================================================")
}

// logFunds queries the on-chain ETH balance and logs it alongside the
// accumulated gas cost tracked by the transaction manager since startup.
func (n *Node) logFunds(ctx context.Context) {
	balance, err := n.txm.Balance(ctx)
	if err != nil {
		log.Warnw("funds: failed to query balance", "address", n.address, "err", err)
		return
	}
	spent := n.txm.TotalGasSpent()
	log.Infow("funds: account",
		"address", n.address,
		"balance", formatETH(balance),
		"gasSpentThisSession", formatETH(spent))
}

// formatETH converts a wei amount to a human-readable ETH string.
func formatETH(wei *big.Int) string {
	if wei == nil {
		return "0.000000 ETH"
	}
	eth := new(big.Float).SetPrec(64).SetInt(wei)
	eth.Quo(eth, new(big.Float).SetPrec(64).SetFloat64(1e18))
	s, _ := eth.Float64()
	return fmt.Sprintf("%.6f ETH", s)
}

// bjjPublicKey returns (pubX, pubY) for this node's BabyJubJub key.
func (n *Node) bjjPublicKey() (*big.Int, *big.Int) {
	pub := group.NewPoint()
	pub.ScalarBaseMult(n.bjjSecret)
	enc := group.Encode(pub)
	return enc.X, enc.Y
}

// EpochPhase enum mirror (matches IDKGRegistry.NodeStatus in Solidity).
const (
	nodeStatusNone     uint8 = 0
	nodeStatusActive   uint8 = 1
	nodeStatusInactive uint8 = 2
)

// EnsureRegistered makes sure the node's BabyJubJub key is registered in
// DKGRegistry and that the row is in the ACTIVE state. It covers three cases:
//
//  1. brand-new node (status == NONE) → `registerKey`
//  2. already registered, key matches, ACTIVE → no-op
//  3. already registered but stale (wrong key or status == INACTIVE) →
//     `updateKey`, which rotates the key *and* auto-reactivates the row
func (n *Node) EnsureRegistered(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}
	existing, err := n.registry.GetNode(callOpts, n.address)
	if err != nil {
		return fmt.Errorf("get node: %w", err)
	}
	wantX, wantY := n.bjjPublicKey()

	// Happy fast-path: already registered, key matches, row is ACTIVE.
	if existing.Status == nodeStatusActive &&
		existing.PubX.Cmp(wantX) == 0 && existing.PubY.Cmp(wantY) == 0 {
		log.Infow("bjj key already registered and active",
			"address", n.address,
			"lastActiveBlock", existing.LastActiveBlock)
		return nil
	}

	// Build a Schnorr PoK over the operator's BJJ secret to satisfy the
	// registry's verification requirement (paper §5.1.1).
	_, _, schnorrProof, err := schnorr.ProveOperatorRegister(n.bjjSecret, n.address)
	if err != nil {
		return fmt.Errorf("schnorr proof: %w", err)
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for registration: %w", err)
	}
	var tx *ethtypes.Transaction
	switch existing.Status {
	case nodeStatusNone:
		log.Infow("registering bjj key on-chain (first time)",
			"address", n.address)
		tx, err = n.registry.RegisterKey(auth, wantX, wantY, schnorrProof.Ax, schnorrProof.Ay, schnorrProof.Z)
	case nodeStatusInactive:
		log.Warnw("node is INACTIVE on-chain, reactivating via updateKey",
			"address", n.address,
			"lastActiveBlock", existing.LastActiveBlock)
		tx, err = n.registry.UpdateKey(auth, wantX, wantY, schnorrProof.Ax, schnorrProof.Ay, schnorrProof.Z)
	default: // ACTIVE but stale key
		log.Infow("rotating bjj key on-chain",
			"address", n.address,
			"oldPubX", existing.PubX, "newPubX", wantX)
		tx, err = n.registry.UpdateKey(auth, wantX, wantY, schnorrProof.Ax, schnorrProof.Ay, schnorrProof.Z)
	}
	if err != nil {
		return fmt.Errorf("register/update key tx: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		return fmt.Errorf("wait register: %w", err)
	}
	log.Infow("bjj key registration confirmed", "address", n.address)
	return nil
}

// maintainLiveness runs on every tick and keeps the node's on-chain liveness
// row healthy without any operator action:
//
//  1. If we have drifted above the heartbeat trigger (80% of
//     INACTIVITY_WINDOW has elapsed since the last refresh) we call
//     heartbeat() proactively. The call is a single SSTORE (~5k gas).
//  2. If we have been reaped out-of-band (status flipped to INACTIVE)
//     — e.g. because the reaper ran before our first lucky epoch —
//     we call reactivate() to rejoin the active set.
//
// The method is tolerant of transient RPC errors: anything unexpected is
// logged at warn and the next tick retries.
func (n *Node) maintainLiveness(ctx context.Context) {
	callOpts := &bind.CallOpts{Context: ctx}
	node, err := n.registry.GetNode(callOpts, n.address)
	if err != nil {
		log.Warnw("liveness: getNode failed", "err", err)
		return
	}
	window, err := n.registry.INACTIVITYWINDOW(callOpts)
	if err != nil {
		log.Warnw("liveness: INACTIVITY_WINDOW read failed", "err", err)
		return
	}
	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		log.Warnw("liveness: blockNumber read failed", "err", err)
		return
	}

	// Case: we got reaped while running. Rejoin the active set.
	if node.Status == nodeStatusInactive {
		log.Warnw("liveness: node is INACTIVE on-chain, calling reactivate()",
			"address", n.address,
			"lastActiveBlock", node.LastActiveBlock,
			"head", head)
		if err := n.sendReactivate(ctx); err != nil {
			log.Warnw("liveness: reactivate failed", "err", err)
		}
		return
	}
	if node.Status != nodeStatusActive {
		// NONE status — not registered. EnsureRegistered handles this on
		// startup; if we get here something is very wrong.
		log.Warnw("liveness: node not registered on-chain",
			"address", n.address, "status", node.Status)
		return
	}

	// Case: we are ACTIVE but drifting. Refresh preemptively.
	// The heartbeat threshold is 80% of the window so we always leave a
	// generous safety margin against slow RPC, reorg variance, or a
	// temporarily stuck poller.
	elapsed := uint64(0)
	if head > node.LastActiveBlock {
		elapsed = head - node.LastActiveBlock
	}
	threshold := (window * 4) / 5
	if elapsed < threshold {
		return
	}

	log.Infow("liveness: sending heartbeat preemptively",
		"address", n.address,
		"elapsed", elapsed,
		"window", window,
		"threshold", threshold,
		"lastActiveBlock", node.LastActiveBlock,
		"head", head)
	if err := n.sendHeartbeat(ctx); err != nil {
		log.Warnw("liveness: heartbeat failed", "err", err)
	}
}

// sendHeartbeat dispatches a registry.heartbeat() transaction.
func (n *Node) sendHeartbeat(ctx context.Context) error {
	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := n.registry.Heartbeat(auth)
	if err != nil {
		return fmt.Errorf("heartbeat tx: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		return fmt.Errorf("wait heartbeat: %w", err)
	}
	log.Infow("liveness: heartbeat confirmed", "address", n.address)
	return nil
}

// sendReactivate dispatches a registry.reactivate() transaction.
func (n *Node) sendReactivate(ctx context.Context) error {
	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := n.registry.Reactivate(auth)
	if err != nil {
		return fmt.Errorf("reactivate tx: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		return fmt.Errorf("wait reactivate: %w", err)
	}
	log.Infow("liveness: reactivate confirmed", "address", n.address)
	return nil
}

// Run is the main participation loop; blocks until ctx is done.
func (n *Node) Run(ctx context.Context, cfg *Config) {
	// Fee-bump / resubmit stuck transactions in the background.
	n.txm.Start(ctx)
	defer n.txm.Stop()
	ticker := time.NewTicker(cfg.PollInterval)
	defer ticker.Stop()
	// Emit balance and gas-spent info every 10 minutes regardless of poll interval.
	fundsTicker := time.NewTicker(10 * time.Minute)
	defer fundsTicker.Stop()
	log.Infow(
		"node running",
		"address", n.address,
		"poll", cfg.PollInterval,
		"auto-create", cfg.AutoCreateEpochs,
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-fundsTicker.C:
			n.logFunds(ctx)
		case <-ticker.C:
			// Keep our on-chain liveness row healthy before scanning epochs.
			// This guarantees heartbeat()/reactivate() fire even when there
			// are no active epochs to participate in.
			n.maintainLiveness(ctx)
			if cfg.AutoCreateEpochs {
				n.maybeScheduleAutoCreate(ctx, cfg)
			}
			if err := n.tick(ctx); err != nil {
				log.Errorw(err, "participation tick")
			}
		}
	}
}

// maybeScheduleAutoCreate races other nodes to fire `createEpoch` once the
// contract's `nextEpochStartBlock()` cadence threshold has been reached.
// Each candidate sleeps a uniform-random delay in [0, AutoCreateJitter)
// before firing, so the population spreads out and most loser txs revert
// cheaply at the contract's `block.number < nextEpochStartBlock()` guard.
//
// Idempotent within a cadence window: we cache the nextEpochStartBlock()
// value the most-recent attempt was scheduled against, and skip
// re-scheduling for the same threshold.
func (n *Node) maybeScheduleAutoCreate(ctx context.Context, cfg *Config) {
	callOpts := &bind.CallOpts{Context: ctx}
	next, err := n.manager.NextEpochStartBlock(callOpts)
	if err != nil {
		log.Warnw("auto-create: read nextEpochStartBlock failed", "err", err)
		return
	}
	currentBlock, err := n.contracts.Pool().Current().BlockNumber(ctx)
	if err != nil {
		log.Warnw("auto-create: read block number failed", "err", err)
		return
	}
	early := false
	if currentBlock < next {
		// Not due by the cadence, but the newest epoch may be nearly claimed
		// out (or dead): the next epoch has to exist before the last key
		// goes, and the contract allows createEpoch early in exactly those
		// two cases (docs/pool-keys-v4.md §9).
		slot, allowed := n.earlyCreateAllowed(ctx)
		if !allowed || n.autoCreateEarlyPool == slot {
			return
		}
		n.autoCreateEarlyPool = slot
		early = true
	} else {
		if n.autoCreateNextStart == next {
			return // already scheduled / fired for this window
		}
		n.autoCreateNextStart = next
	}

	jitter := time.Duration(0)
	if cfg.AutoCreateJitter > 0 {
		jitter = time.Duration(mrand.Int63n(int64(cfg.AutoCreateJitter)))
	}
	log.Infow(
		"auto-create: scheduling createEpoch attempt",
		"nextStart", next,
		"currentBlock", currentBlock,
		"early", early,
		"jitter", jitter,
	)
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(jitter):
		}
		// Re-check: another node may have won the race during our sleep.
		check, err := n.manager.NextEpochStartBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Warnw("auto-create: re-read nextEpochStartBlock failed", "err", err)
			return
		}
		if check != next {
			log.Debugw("auto-create: another node won the race",
				"originalNext", next, "currentNext", check)
			return
		}
		if err := n.fireCreateEpoch(ctx, cfg); err != nil {
			if early {
				// The pool-drain trigger races the cadence: a revert here
				// just means "not yet", and the cadence path will retry.
				log.Debugw("auto-create: early attempt refused", "err", decodeContractError(err))
				return
			}
			log.Warnw("auto-create: createEpoch failed (likely lost race)", "err", err)
			return
		}
		log.Infow("auto-create: createEpoch landed", "nextStart", next)
	}()
}

// fireCreateEpoch sends the createEpoch transaction with the policy
// configured in cfg.EpochPolicy. The decryption policy is left empty
// (no owner-only restriction, no time locks) — operators wanting tighter
// gating should configure it via per-application AppPolicy at
// registerApplication time.
// adaptivePolicy sizes the committee from the registry so a fleet that grows
// or shrinks keeps filling committees without operators coordinating: n is
// three quarters of the active operators (rounded up, capped at MaxN and
// floored by the contract), t a majority of n, and m_min two thirds of n.
// With α ≥ 1.34 every active operator is admissible, so the first n to claim
// form the committee and a quarter of the fleet may be offline.
func (n *Node) adaptivePolicy(ctx context.Context, alphaBps uint16) (EpochPolicyConfig, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	active, err := n.registry.ActiveCount(callOpts)
	if err != nil {
		return EpochPolicyConfig{}, fmt.Errorf("active count: %w", err)
	}
	minN, err := n.manager.MINCOMMITTEESIZE(callOpts)
	if err != nil {
		return EpochPolicyConfig{}, fmt.Errorf("min committee size: %w", err)
	}
	minT, err := n.manager.MINTHRESHOLD(callOpts)
	if err != nil {
		return EpochPolicyConfig{}, fmt.Errorf("min threshold: %w", err)
	}
	maxAlpha, err := n.manager.MAXLOTTERYALPHABPS(callOpts)
	if err != nil {
		return EpochPolicyConfig{}, fmt.Errorf("max lottery alpha: %w", err)
	}
	size := (active*3 + 3) / 4 // ceil(0.75·active)
	size = min(size, uint64(ccommon.MaxN))
	size = max(size, uint64(minN), 1)
	t := max(size/2+1, uint64(minT))
	mMin := max(t, (size*2+2)/3) // ceil(2n/3)
	alpha := min(alphaBps, maxAlpha)
	return EpochPolicyConfig{
		Threshold:             uint16(t),
		CommitteeSize:         uint16(size),
		MinValidContributions: uint16(mMin),
		LotteryAlphaBps:       alpha,
	}, nil
}

// abortedEpochSlot marks an early-creation observation made against an
// Aborted newest epoch (which has no pool cursor to key it by).
const abortedEpochSlot = 0xff

// earlyCreateAllowed mirrors the contract's early-creation rule: before the
// cadence, createEpoch only succeeds when the newest epoch is Live with at
// most one unclaimed pool key (poolNext ≥ MaxK − 1) or Aborted. It returns
// the (epoch, poolNext) observation the decision was made against so the
// early trigger fires once per drain, not once per tick.
func (n *Node) earlyCreateAllowed(ctx context.Context) (poolSlot, bool) {
	callOpts := &bind.CallOpts{Context: ctx}
	nonce, err := n.manager.EpochNonce(callOpts)
	if err != nil || nonce == 0 {
		return poolSlot{}, false
	}
	prefix, err := n.manager.EPOCHPREFIX(callOpts)
	if err != nil {
		return poolSlot{}, false
	}
	epochID := web3.EpochID(prefix, nonce)
	epoch, err := n.manager.GetEpoch(callOpts, epochID)
	if err != nil {
		return poolSlot{}, false
	}
	switch epoch.Status {
	case epochAborted:
		return poolSlot{epoch: epochID, key: abortedEpochSlot}, true
	case epochLive:
		next, err := n.manager.GetPoolStatus(callOpts, epochID)
		if err != nil {
			return poolSlot{}, false
		}
		return poolSlot{epoch: epochID, key: next}, int(next) >= ccommon.MaxK-1
	default:
		return poolSlot{}, false
	}
}

func (n *Node) fireCreateEpoch(ctx context.Context, cfg *Config) error {
	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts: %w", err)
	}
	policy := cfg.EpochPolicy
	if policy.Adaptive() {
		derived, err := n.adaptivePolicy(ctx, policy.LotteryAlphaBps)
		if err != nil {
			return fmt.Errorf("derive epoch policy: %w", err)
		}
		policy = derived
		log.Infow("auto-create: derived epoch policy from the registry",
			"n", policy.CommitteeSize, "t", policy.Threshold, "mMin", policy.MinValidContributions)
	}
	tx, err := n.manager.CreateEpoch(
		auth,
		policy.Threshold,
		policy.CommitteeSize,
		policy.MinValidContributions,
		policy.LotteryAlphaBps,
	)
	if err != nil {
		return fmt.Errorf("create epoch: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 30*time.Second); err != nil {
		return fmt.Errorf("wait tx: %w", err)
	}
	return nil
}

// epochLookback is how many of the newest epochs the lifecycle scan always
// covers. Normally only the newest two can still be in CommitteeSelection /
// KeyAssembly, but an epoch that qualified and was never finalized (every
// member offline through the finalize window) stays in KeyAssembly across
// cadences and must remain discoverable after a restart; visiting a few
// more costs one getEpoch per tick until each one is seen terminal. Past
// the window the scan keeps walking back while the chain still reports
// unfinished epochs (see epochsToVisit), so no outage is long enough to
// hide one.
const epochLookback = 8

// epochReader is the slice of the chain the lifecycle scan needs to look
// past its fixed window; *web3.Contracts implements it.
type epochReader interface {
	GetEpoch(ctx context.Context, epochID [12]byte) (web3.EpochView, error)
}

// tick runs one poll cycle: key-generation lifecycle for the newest epochs,
// then decryption service for every pending ciphertext.
//
// Live is key-generation-terminal: every pool key and share root is stored
// by finalizeEpoch, so a Live epoch needs no further lifecycle work and is
// dropped from the scan at once. Decryption work is discovered from
// CiphertextSubmitted events instead of by walking epochs, because
// applications outlive epochs by design.
func (n *Node) tick(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}
	epochNonce, err := n.manager.EpochNonce(callOpts)
	if err != nil {
		return fmt.Errorf("epoch nonce: %w", err)
	}
	prefix, err := n.manager.EPOCHPREFIX(callOpts)
	if err != nil {
		return fmt.Errorf("epoch prefix: %w", err)
	}
	for _, epochID := range n.epochsToVisit(ctx, n.contracts, prefix, epochNonce) {
		if err := n.participate(ctx, epochID); err != nil {
			log.Warnw("participate failed", "epoch", roundHex(epochID), "err", decodeContractError(err))
		}
	}
	if err := n.scanCiphertexts(ctx); err != nil {
		log.Warnw("ciphertext scan failed", "err", err)
	}
	n.serviceCiphertexts(ctx)
	return nil
}

// epochsToVisit is the newest epochLookback epochs minus the terminal ones,
// oldest first, preceded by every older epoch the chain still reports as
// unfinished: past the window the scan steps back one nonce at a time while
// getEpoch says CommitteeSelection or KeyAssembly and stops at the first
// Live / Aborted (or otherwise closed) epoch, or at nonce 1, so an unfinished
// qualifying epoch stays discoverable however many cadences have passed
// (docs/pool-keys-v4.md §10). A read failure ends the walk for this tick.
func (n *Node) epochsToVisit(ctx context.Context, chain epochReader, prefix uint32, epochNonce uint64) [][12]byte {
	first := uint64(1)
	if epochNonce >= epochLookback {
		first = epochNonce - epochLookback + 1
	}
	var older [][12]byte // newest first
	for nonce := first - 1; nonce >= 1; nonce-- {
		id := web3.EpochID(prefix, nonce)
		epoch, err := chain.GetEpoch(ctx, id)
		if err != nil {
			log.Warnw("lifecycle scan: cannot read epoch past the lookback window",
				"epoch", roundHex(id), "err", err)
			break
		}
		if epoch.Status != epochCommitteeSelection && epoch.Status != epochKeyAssembly {
			break
		}
		if !n.terminal[id] {
			older = append(older, id)
		}
	}
	out := make([][12]byte, 0, len(older)+epochLookback)
	for i := len(older) - 1; i >= 0; i-- {
		out = append(out, older[i])
	}
	for nonce := first; nonce <= epochNonce; nonce++ {
		if id := web3.EpochID(prefix, nonce); !n.terminal[id] {
			out = append(out, id)
		}
	}
	return out
}

// finish marks an epoch's key-generation lifecycle as needing no more work.
func (n *Node) finish(epochID [12]byte) {
	n.terminal[epochID] = true
	delete(n.finalizeRetry, epochID)
}

func (n *Node) participate(ctx context.Context, epochID [12]byte) error {
	epoch, err := n.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return fmt.Errorf("get epoch: %w", err)
	}
	switch epoch.Status {
	case 0: // None — epoch slot exists on-chain but is uninitialised (should not happen)
		return nil

	case epochCommitteeSelection: // try to claim a slot in the lottery
		return n.doClaimSlot(ctx, epochID, epoch)

	case epochKeyAssembly: // selected participants submit ZK shares,
		//                  then race on a deterministic stagger to prove and call finalizeEpoch.
		// The committee is frozen once KeyAssembly starts, so it is safe to
		// cache it like the decryption path does for Live epochs.
		selected, err := n.selected(ctx, epochID)
		if err != nil {
			return err
		}
		idx := myIndex(selected, n.address)
		if idx == 0 {
			// Not selected: contributing and finalizing are the committee's
			// job, and nothing here ever needs this node again.
			n.finish(epochID)
			return nil
		}
		if err := n.doContribution(ctx, epochID, idx, epoch, selected); err != nil {
			return err
		}
		// After contributing, every selected participant rotates through a
		// deterministic finalize stagger so normally one node proves and
		// submits at a time (any race-loser sees AlreadyLive and stops).
		if err := n.tryAutoFinalize(ctx, epochID, idx, selected); err != nil {
			log.Warnw("auto-finalize attempt failed",
				"epoch", roundHex(epochID), "err", err)
		}
		return nil

	case epochLive:
		// Every pool key and share root landed with finalizeEpoch: key
		// generation is over. Ciphertexts are the decryption scanner's job.
		n.finish(epochID)
		return nil

	case epochAborted:
		log.Warnw("epoch aborted — no further participation possible",
			"epoch", roundHex(epochID))
		n.finish(epochID)
		return nil

	case epochCompleted:
		n.finish(epochID)
		return nil

	default:
		log.Warnw("unknown epoch status — skipping", "epoch", roundHex(epochID), "status", epoch.Status)
		return nil
	}
}

// ---- Lottery slot claim ----

// doClaimSlot races to claim a committee slot for the epoch. Eligibility is
// derived deterministically from the epoch seed; if the seed has not been
// resolved yet (block.number < epoch.SeedBlock), the call will revert with
// SeedNotReady and we'll retry on the next poll. If the node is not eligible we
// silently no-op for the rest of the epoch.
func (n *Node) doClaimSlot(ctx context.Context, epochID [12]byte, epoch web3.EpochView) error {
	if n.signaled[epochID] {
		return nil
	}

	// Read the current head once; used by all pre-flight checks below.
	head, headErr := n.contracts.Client().BlockNumber(ctx)

	// Pre-flight: if the seed block hasn't been reached yet, skip this tick silently.
	// The slot lottery cannot be resolved before the seed is committed; the contract
	// reverts with SeedNotReady when block.number <= seedBlock. We mirror that
	// condition exactly (<=) so we never simulate or broadcast at seedBlock itself.
	if headErr == nil && head > 0 && epoch.SeedBlock > 0 && head <= epoch.SeedBlock {
		log.Debugw("claim slot: seed block not yet reached — waiting",
			"epoch", roundHex(epochID),
			"head", head,
			"seedBlock", epoch.SeedBlock)
		return nil
	}

	// Pre-flight: if the committee is already full, we were not selected.
	if epoch.ClaimedCount >= epoch.Policy.CommitteeSize {
		log.Infow("claim slot: committee already full — not selected for this epoch",
			"epoch", roundHex(epochID),
			"claimed", epoch.ClaimedCount,
			"size", epoch.Policy.CommitteeSize)
		n.signaled[epochID] = true
		return nil
	}

	// Pre-flight: check registration deadline before sending any tx.
	if headErr == nil {
		if head >= epoch.Policy.CommitteeSelectionDeadlineBlock {
			log.Infow("registration deadline already passed — skipping slot claim",
				"epoch", roundHex(epochID),
				"head", head,
				"deadline", epoch.Policy.CommitteeSelectionDeadlineBlock)
			n.signaled[epochID] = true
			return nil
		}
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := n.manager.ClaimSlot(auth, epochID)
	if err != nil {
		// SeedNotReady: the seed block hasn't been mined yet. Retry next poll
		// without setting signaled so we keep trying until the seed arrives.
		// Use decodeContractError (not err.Error()) because custom errors are
		// returned as raw ABI bytes; err.Error() is just "execution reverted".
		if strings.Contains(decodeContractError(err), "SeedNotReady") {
			log.Debugw("claim slot: seed not ready yet, retrying next poll", "epoch", roundHex(epochID))
			return nil
		}
		// Definitively final reverts: the committee is decided without us.
		// Set signaled so we stop sending txs for this epoch.
		if isExpectedClaimRevert(err) {
			log.Debugw("claim slot: not selected for committee", "epoch", roundHex(epochID), "reason", decodeContractError(err))
			n.signaled[epochID] = true
			return nil
		}
		// Unexpected permanent revert — all pre-flights passed but the contract
		// still rejected us. Accept the result and stop retrying.
		if isPermanentRevert(err) {
			log.Warnw("claim slot: unexpected permanent revert — marking as not selected",
				"epoch", roundHex(epochID), "err", decodeContractError(err))
			n.signaled[epochID] = true
			return nil
		}
		return fmt.Errorf("claim slot: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		// A mined revert here is a lost race (SlotsFull) or a gas shortfall
		// (claimSlot grows costlier with every claimed slot, so simultaneous
		// claims can outgrow the estimate). Retrying next tick is one cheap
		// simulation that yields the definitive custom error either way.
		return fmt.Errorf("wait claim slot tx (retrying next tick): %w", err)
	}
	n.signaled[epochID] = true
	log.Infow("slot claimed", "epoch", roundHex(epochID))
	return nil
}

// ---- Contribution ----

func (n *Node) doContribution(
	ctx context.Context,
	epochID [12]byte,
	idx uint16,
	epoch web3.EpochView,
	selected []common.Address,
) error {
	if n.contributed[epochID] {
		return nil
	}
	n.critical.Add(1)
	defer n.critical.Add(-1)
	// Check on-chain (handles restarts).
	rec, err := n.manager.GetContribution(&bind.CallOpts{Context: ctx}, epochID, n.address)
	if err == nil && rec.Accepted {
		log.Infow("contribution already accepted on-chain", "epoch", roundHex(epochID))
		n.contributed[epochID] = true
		return nil
	}

	// Pre-flight: check contribution deadline before burning time on ZK proof.
	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		log.Warnw("doContribution: failed to read block number", "epoch", roundHex(epochID), "err", err)
		// Proceed optimistically; worst case the tx reverts and we catch it below.
	} else if head >= epoch.Policy.KeyAssemblyDeadlineBlock {
		log.Warnw("contribution deadline already passed — skipping epoch",
			"epoch", roundHex(epochID),
			"head", head,
			"deadline", epoch.Policy.KeyAssemblyDeadlineBlock)
		n.contributed[epochID] = true
		return nil
	}

	// Pre-flight: skip when the epoch already has enough contributions to
	// finalize. Unlike partial decryptions (where every participating node
	// is expected to be rewarded), late contributions land on a epoch
	// that's already finalize-eligible — they don't change the outcome,
	// they don't earn anything, they just burn ~seconds of prover CPU and
	// a chain transaction's gas.
	//
	// There's a benign race: another node could submit and push us past
	// the gate between this read and our hypothetical submit. In that
	// case our tx would land successfully but be unrewarded — the same
	// outcome we're trying to avoid, just costing more. This guard keeps
	// the common case cheap.
	if epoch.ContributionCount >= epoch.Policy.MinValidContributions {
		log.Infow(
			"contribution: epoch already has enough contributions to finalize — skipping",
			"epoch", roundHex(epochID),
			"contributions", epoch.ContributionCount,
			"required", epoch.Policy.MinValidContributions,
		)
		n.contributed[epochID] = true
		return nil
	}

	threshold := epoch.Policy.Threshold
	committeeSize := epoch.Policy.CommitteeSize

	roundHash := roundScalar(epochID)
	// One polynomial per pool key: f_{j,i}(x) = Σ a_{j,k} x^k with uniform
	// coefficients in the BabyJubJub scalar field; a_{j,0} is this node's
	// additive share of P_j. Every epoch deals all MaxK of them.
	coeffs := make([][]*big.Int, ccommon.MaxK)
	for j := range coeffs {
		keyCoeffs, err := randomScalars(int(threshold))
		if err != nil {
			return err
		}
		coeffs[j] = keyCoeffs
	}

	// Recipient keys come from the CommitteeSnapshot event frozen when the
	// committee filled, not the live registry: a post-fill updateKey rotation
	// must not change what contributors encrypt shares to (the on-chain prefix
	// hash pins the snapshot keys). Operator addresses still come from the
	// selected list, which the snapshot does not carry.
	// Deployments that predate the CommitteeSnapshot event never emit it;
	// fall back to the live registry there, which is exactly what those
	// contracts hash (no key rotation protection, same as before).
	snapshot, err := n.contracts.CommitteeSnapshot(ctx, epochID)
	if err != nil {
		log.Warnw("committee snapshot unavailable, using live registry keys", "epoch", roundHex(epochID), "err", err)
		snapshot = nil
	} else if len(snapshot) != int(committeeSize) {
		return fmt.Errorf("committee snapshot has %d members, want %d", len(snapshot), committeeSize)
	}
	recipientIdxs := make([]uint16, committeeSize)
	recipientKeys := make([]nodetypes.NodeKey, committeeSize)
	for i := uint16(0); i < committeeSize; i++ {
		recipientIdxs[i] = i + 1
		if snapshot != nil {
			recipientKeys[i] = snapshot[i]
			recipientKeys[i].Operator = selected[i]
			continue
		}
		nd, err := n.contracts.GetNode(ctx, selected[i])
		if err != nil {
			return fmt.Errorf("get node key idx=%d: %w", i+1, err)
		}
		recipientKeys[i] = nodetypes.NodeKey{Operator: selected[i], PubX: nd.PubX, PubY: nd.PubY}
	}

	// One fresh hashed-ElGamal nonce per recipient. The nonce is the only
	// thing hiding the share (mask = H(r·pub_j)); a predictable r would let
	// anyone unmask every share from calldata.
	nonces, err := randomScalars(int(committeeSize))
	if err != nil {
		return err
	}

	log.Infow(
		"contribution assignment",
		"epoch", roundHex(epochID),
		"index", idx,
		"threshold", threshold,
		"committeeSize", committeeSize,
		"deadline", epoch.Policy.KeyAssemblyDeadlineBlock,
		"head", head,
	)

	asgn := contribution.Assignment{
		RoundHash:        roundHash,
		Threshold:        threshold,
		CommitteeSize:    committeeSize,
		ContributorIndex: idx,
		Coefficients:     coeffs,
		RecipientIndexes: recipientIdxs,
		RecipientKeys:    recipientKeys,
		EncryptionNonces: nonces,
	}
	witness, pi, err := contribution.BuildWitness(asgn)
	if err != nil {
		return fmt.Errorf("build contribution witness: %w", err)
	}
	proof, err := n.runtimes.contribution.ProveAndVerify(witness)
	if err != nil {
		return fmt.Errorf("prove contribution: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return fmt.Errorf("marshal contribution proof: %w", err)
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return fmt.Errorf("encode contribution public witness: %w", err)
	}
	transcriptScalars, err := pi.TranscriptScalars()
	if err != nil {
		return fmt.Errorf("contribution transcript scalars: %w", err)
	}
	transcriptBytes, err := encodeWords(transcriptScalars...)
	if err != nil {
		return fmt.Errorf("encode contribution transcript: %w", err)
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for contribution: %w", err)
	}
	tx, err := n.manager.SubmitContribution(
		auth, epochID, idx,
		common.BigToHash(pi.CommitmentHash),
		common.BigToHash(pi.ShareHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		// AlreadyContributed: another copy of this node already submitted, or we
		// restarted after the on-chain pre-check passed but before recording locally.
		if strings.Contains(decodeContractError(err), "AlreadyContributed") {
			log.Infow("contribution already on-chain (benign race) — skipping",
				"epoch", roundHex(epochID))
			n.contributed[epochID] = true
			return nil
		}
		if isPermanentRevert(err) {
			log.Warnw("contribution tx permanently rejected — will not retry this epoch",
				"epoch", roundHex(epochID), "err", decodeContractError(err))
			n.contributed[epochID] = true
		}
		return fmt.Errorf("submit contribution: %w", err)
	}
	n.txm.RecordPending(tx)
	if err := n.txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		if strings.Contains(decodeContractError(err), "AlreadyContributed") {
			log.Infow("contribution already on-chain (benign race) — skipping",
				"epoch", roundHex(epochID))
			n.contributed[epochID] = true
			return nil
		}
		if isPermanentRevert(err) {
			log.Warnw("contribution tx reverted on-chain — will not retry this epoch",
				"epoch", roundHex(epochID), "err", decodeContractError(err))
			n.contributed[epochID] = true
		}
		return fmt.Errorf("wait contribution tx: %w", err)
	}
	n.contributed[epochID] = true
	n.ownContribs[epochID] = &savedContrib{
		coefficients:     coeffs,
		recipientIndexes: recipientIdxs,
		recipientKeys:    recipientKeys,
	}
	var gasUsed uint64
	if rec, err := n.contracts.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		gasUsed = rec.GasUsed
	}
	log.Infow("contribution submitted", "epoch", roundHex(epochID), "index", idx, "tx", tx.Hash().Hex(), "gas", gasUsed)
	return nil
}

// ---- Auto-finalize (deterministic stagger across selected participants) ----

// staggerBlocks is the per-slot delay between successive attempts in a
// rotation. With STAGGER_BLOCKS=3, slot 0 attempts at liveNotBeforeBlock,
// slot 1 at +3, slot 2 at +6, etc.
const staggerBlocks = 3

// tryAutoFinalize is called by every selected participant after their
// contribution lands. It computes the participant's slot in an epoch-specific
// rotation derived from the lottery seed, waits until that slot's window
// opens (anchored at liveNotBeforeBlock, the block from which the contract
// accepts the finalization), then reconstructs every accepted contribution
// from calldata, proves the batched finalization and submits finalizeEpoch.
// The first proof to land makes the epoch Live with every pool key and share
// root stored; everyone else sees AlreadyLive and stops.
//
// Pre-checks short-circuit cheaply: if the epoch is already past KeyAssembly
// or we know the race is over, skip. If the contribution count is below
// minValidContributions or block.number is below the wait-until block, defer
// to the next tick. A failed attempt (anything but a lost race) is retried
// with an exponential per-epoch backoff: the proof is expensive, so a
// persistent fault must not cost one per tick, while a transient RPC or
// transaction hiccup is retried soon.
func (n *Node) tryAutoFinalize(
	ctx context.Context,
	epochID [12]byte,
	myIdx uint16,
	selected []common.Address,
) error {
	if n.finalized[epochID] {
		return nil
	}
	if b := n.finalizeRetry[epochID]; b != nil && !b.due() {
		return nil
	}
	epoch, err := n.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return fmt.Errorf("get epoch: %w", err)
	}
	if epoch.Status != epochKeyAssembly { // not in KeyAssembly any more — Live, aborted, etc.
		n.finalized[epochID] = true
		return nil
	}
	if epoch.ContributionCount < epoch.Policy.MinValidContributions {
		return nil // not enough contributions yet; retry next tick
	}

	committeeSize := uint16(len(selected))
	if committeeSize == 0 {
		return nil
	}
	mySlot := staggerSlot(epoch.Seed, 0, myIdx, committeeSize)
	waitUntil := epoch.Policy.LiveNotBeforeBlock + mySlot*staggerBlocks

	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("read head: %w", err)
	}
	if head < waitUntil {
		return nil // not our turn yet; another tick will retry
	}

	n.critical.Add(1)
	defer n.critical.Add(-1)
	log.Infow(
		"auto-finalize: my turn",
		"epoch", roundHex(epochID),
		"myIdx", myIdx,
		"mySlot", mySlot,
		"head", head,
		"finalizeNotBefore", epoch.Policy.LiveNotBeforeBlock,
		"contributions", epoch.ContributionCount,
	)

	res, err := finalizer.ProveAndSubmitFinalize(
		ctx, n.contracts, n.manager, n.txm, n.runtimes.finalize, epochID, n.contribCache,
	)
	if err != nil {
		// Only a lost race is final: AlreadyLive from the contract (or from
		// the finalizer's own re-check), or the epoch having left
		// KeyAssembly. Anything else (a dealer's calldata not yet readable,
		// a mined revert, an RPC hiccup) is retried with backoff.
		reason := decodeContractError(err)
		status := epochKeyAssembly // unknown → assume still open so we retry
		if cur, rerr := n.contracts.GetEpoch(ctx, epochID); rerr == nil {
			status = cur.Status
		}
		if errors.Is(err, finalizer.ErrAlreadyLive) || finalizeRaceLost(reason, status) {
			log.Infow("auto-finalize: another node beat us", "epoch", roundHex(epochID), "reason", reason)
			n.finalized[epochID] = true
			return nil
		}
		b := n.finalizeRetry[epochID]
		if b == nil {
			b = &serviceBackoff{}
			n.finalizeRetry[epochID] = b
		}
		b.fail()
		return fmt.Errorf("auto-finalize failed (retrying in %d ticks): %s", b.wait, reason)
	}
	n.finalized[epochID] = true
	delete(n.finalizeRetry, epochID)
	log.Infow("auto-finalize: epoch finalized by us",
		"epoch", roundHex(epochID), "tx", res.TxHash.Hex(), "gas", res.GasUsed,
		"contributors", res.ParticipantIndexes, "keys", len(res.PoolKeys))
	return nil
}

// buildPrivateShare computes d_{j,i} = Σ_c f_{c,j}(i) for one pool key over
// every accepted contribution by:
//   - Own contribution (in-memory cache): evaluate own polynomial directly
//   - Own contribution (after restart, cache lost): fall back to calldata recovery
//   - Other contributions: scan on-chain txs for calldata and decrypt
//
// The result is cached only when every accepted contribution was recovered;
// a partial sum is a wrong share, so any failure is returned as a retryable
// error instead.
//
// The calldata scan starts at the epoch's seed block, which keeps it tight
// while still capturing contributions that arrive during the registration
// phase (nodes can contribute immediately after claiming a slot, which is
// only possible from seedBlock onward).
func (n *Node) buildPrivateShare(
	ctx context.Context,
	client calldataReader,
	epochID [12]byte,
	keyIndex uint8,
	myIdx uint16,
	selected []common.Address,
	epoch epochView,
	callOpts *bind.CallOpts,
) (*big.Int, error) {
	slot := poolSlot{epoch: epochID, key: keyIndex}
	if s, ok := n.privateShares[slot]; ok {
		return s, nil
	}
	roundHash := roundScalar(epochID)
	fromBlock := epoch.SeedBlock
	layout, err := contribution.NewLayout(int(epoch.Policy.Threshold), int(epoch.Policy.CommitteeSize))
	if err != nil {
		return nil, fmt.Errorf("epoch policy: %w", err)
	}

	shares := make([]*big.Int, 0, len(selected))
	expected := 0
	for i, addr := range selected {
		contribIdx := uint16(i + 1)
		rec, err := n.manager.GetContribution(callOpts, epochID, addr)
		if err != nil {
			return nil, fmt.Errorf("get contribution of %s: %w", addr.Hex(), err)
		}
		if !rec.Accepted {
			continue
		}
		expected++
		if addr == n.address {
			if share, ok := n.ownShare(epochID, keyIndex, myIdx); ok {
				shares = append(shares, share)
				continue
			}
		}
		share, err := n.recoverShareFrom(ctx, client, epochID, addr, contribIdx, keyIndex,
			roundHash, rec.CommitmentsHash, myIdx, fromBlock, layout)
		if err != nil {
			return nil, fmt.Errorf("recover share from %s (idx %d): %w", addr.Hex(), contribIdx, err)
		}
		shares = append(shares, share)
	}
	total, err := sumRecoveredShares(shares, expected)
	if err != nil {
		return nil, err
	}
	log.Infow("private share built",
		"epoch", roundHex(epochID), "key", keyIndex, "contributions", expected, "myIdx", myIdx)
	n.privateShares[slot] = total
	return total, nil
}

// ownShare evaluates this node's own polynomial for pool key `keyIndex` at
// myIdx from the in-memory contribution cache. ok=false means the caller must
// recover the share from calldata like any other contribution.
func (n *Node) ownShare(epochID [12]byte, keyIndex uint8, myIdx uint16) (*big.Int, bool) {
	sc := n.ownContribs[epochID]
	if sc == nil || int(keyIndex) >= len(sc.coefficients) {
		log.Infow("own contribution cache missing (node restarted?), recovering from calldata",
			"epoch", roundHex(epochID))
		return nil, false
	}
	share, err := ccommon.EvaluatePolynomialNative(sc.coefficients[keyIndex], big.NewInt(int64(myIdx)))
	if err != nil {
		log.Warnw("own polynomial evaluation failed, falling back to calldata",
			"epoch", roundHex(epochID), "err", err)
		return nil, false
	}
	return share, true
}

// sumRecoveredShares folds the recovered f_{c,j}(i) values into d_{j,i} modulo the
// group order. It refuses anything short of the full set of `expected`
// accepted contributions: the sum is only a valid share when complete.
func sumRecoveredShares(shares []*big.Int, expected int) (*big.Int, error) {
	if expected == 0 {
		return nil, fmt.Errorf("private share: no accepted contributions")
	}
	if len(shares) != expected {
		return nil, fmt.Errorf("private share: recovered %d/%d contributions — calldata not yet available, retry later",
			len(shares), expected)
	}
	modulus := group.ScalarField()
	total := new(big.Int)
	for _, share := range shares {
		total.Add(total, share)
		total.Mod(total, modulus)
	}
	if total.Sign() == 0 {
		return nil, fmt.Errorf("private share is zero after %d contributions — possible Shamir evaluation issue", expected)
	}
	return total, nil
}

// calldataReader is the slice of the chain client that contribution calldata
// recovery needs (what finalizer.ContributionCalldata takes); *web3.PooledBackend
// implements it.
type calldataReader interface {
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*ethtypes.Transaction, bool, error)
}

// recoverShareFrom fetches the submitContribution tx calldata for `contributor`
// (located through the ContributionSubmitted event log, or the on-disk cache
// when it already holds it) and decrypts the share slot destined for myIdx
// under pool key `keyIndex`. The compact transcript is decoded with the
// epoch's (t, n), which come from the epoch policy, never from calldata, and
// the calldata is trusted — and cached — only once it reproduces the dealer's
// stored commitmentsHash and the share it yields opens the dealer's
// commitments (see recoverShare); a cache entry that fails is refetched, a
// fetched body that fails is an error.
func (n *Node) recoverShareFrom(
	ctx context.Context,
	client calldataReader,
	epochID [12]byte,
	contributor common.Address,
	contribIdx uint16,
	keyIndex uint8,
	roundHash *big.Int,
	storedHash [32]byte,
	myIdx uint16,
	fromBlock uint64,
	layout contribution.Layout,
) (*big.Int, error) {
	if myIdx == 0 || int(myIdx) > layout.CommitteeSize {
		return nil, fmt.Errorf("member index %d outside the committee of %d", myIdx, layout.CommitteeSize)
	}
	recovered := func(data []byte) (*big.Int, error) {
		tr, err := decodeContribution(data, layout)
		if err != nil {
			return nil, err
		}
		return recoverShare(tr, layout, contribIdx, myIdx, keyIndex, roundHash, storedHash, n.bjjSecret)
	}
	data, cached := n.contribCache.Get(epochID, contributor)
	share, err := recovered(data)
	if err != nil {
		if cached {
			log.Warnw("cached contribution calldata unusable, refetching",
				"epoch", roundHex(epochID), "contributor", contributor, "err", err)
		}
		data, err = finalizer.ContributionCalldata(ctx, client, n.manager, epochID, contributor, fromBlock)
		if err != nil {
			return nil, err
		}
		if share, err = recovered(data); err != nil {
			return nil, err
		}
		n.contribCache.Put(epochID, contributor, data)
	}
	return share, nil
}

// recoverShare decrypts member myIdx's share of pool key keyIndex from one
// dealer's decoded transcript and accepts it only if the calldata is the
// dealer's accepted contribution: the decoded commitments must reproduce the
// commitmentsHash the contract stored for the dealer, and the share must open
// the dealer's key-keyIndex commitments (share·G = Σ_m myIdx^m · A[j][m]). A
// wrong RPC body or a corrupted cache entry would otherwise become part of
// d_{j,i} and make every partial decryption under that key revert.
func recoverShare(
	tr *contribution.Transcript,
	layout contribution.Layout,
	contribIdx, myIdx uint16,
	keyIndex uint8,
	roundHash *big.Int,
	storedHash [32]byte,
	secret *big.Int,
) (*big.Int, error) {
	if int(keyIndex) >= len(tr.Commitments) || int(keyIndex) >= len(tr.MaskedShares) {
		return nil, fmt.Errorf("pool key %d outside the transcript's %d keys", keyIndex, len(tr.Commitments))
	}
	got, err := finalizer.ContributionHash(roundHash, contribIdx, layout.Threshold, tr.Commitments)
	if err != nil {
		return nil, err
	}
	if got.Cmp(new(big.Int).SetBytes(storedHash[:])) != 0 {
		return nil, fmt.Errorf("commitments hash %s does not match the stored %x", got, storedHash)
	}
	// Decode enforces that recipient slot i carries index i+1, so this
	// node's ephemeral and masked share sit at slot myIdx−1.
	slot := int(myIdx) - 1
	ct := shareenc.Ciphertext{
		Ephemeral:   tr.Ephemerals[slot],
		MaskedShare: tr.MaskedShares[keyIndex][slot],
	}
	share, err := shareenc.DecryptShareRoundHash(roundHash, contribIdx, myIdx, keyIndex, ct, secret)
	if err != nil {
		return nil, err
	}
	if err := feldman.VerifyShare(tr.Commitments[keyIndex], myIdx, share); err != nil {
		return nil, fmt.Errorf("key %d share dealt by member %d: %w", keyIndex, contribIdx, err)
	}
	return share, nil
}

// decodeContribution decodes raw submitContribution calldata under the
// epoch's compact layout; nil calldata (a cache miss) is an error like any
// malformed payload, so callers fall back to the chain.
func decodeContribution(data []byte, layout contribution.Layout) (*contribution.Transcript, error) {
	if data == nil {
		return nil, fmt.Errorf("no contribution calldata")
	}
	tr, err := finalizer.DecodeContribution(data, layout)
	if err != nil {
		return nil, fmt.Errorf("decode contribution transcript: %w", err)
	}
	return tr, nil
}

// ---- small helpers ----

func myIndex(selected []common.Address, addr common.Address) uint16 {
	for i, a := range selected {
		if a == addr {
			return uint16(i + 1)
		}
	}
	return 0
}

func roundScalar(id [12]byte) *big.Int {
	return new(big.Int).SetBytes(id[:])
}

func roundHex(id [12]byte) string { return fmt.Sprintf("%x", id) }

// staggerSlot returns myIdx's position in the epoch's rotation: a
// permutation of [0, committeeSize) whose start is derived from the lottery
// seed (plus a caller-chosen salt, e.g. the ciphertext index) so that a
// different member goes first for every epoch and every slot.
func staggerSlot(seed common.Hash, salt uint64, myIdx, committeeSize uint16) uint64 {
	n := uint64(committeeSize)
	if n == 0 {
		return 0
	}
	startSlot := new(big.Int).SetBytes(seed[:])
	startSlot.Add(startSlot, new(big.Int).SetUint64(salt))
	start := startSlot.Mod(startSlot, new(big.Int).SetUint64(n)).Uint64()
	return (uint64(myIdx-1) + n - start) % n
}

// finalizeRaceLost reports whether a failed finalize attempt is final: the
// contract said AlreadyLive, or the epoch has left KeyAssembly. Any other
// failure must be retried.
func finalizeRaceLost(reason string, status uint8) bool {
	return strings.Contains(reason, "AlreadyLive") || status != epochKeyAssembly
}

// rpcHost returns only the host of an RPC URL for logging; provider API
// keys live in the path or userinfo and must not reach the logs.
func rpcHost(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "<unparseable rpc url>"
	}
	return u.Host
}

// isExpectedClaimRevert returns true if a claimSlot revert is "benign" — i.e.
// the node should silently accept it (not eligible, slot already gone, seed not
// yet available). The node will retry on the next poll for the SeedNotReady case
// since `signaled` only flips on definitively-final reverts.
// isPermanentRevert returns true when the error indicates the EVM rejected the
// transaction in a way that retrying will never succeed. Transient errors (RPC
// timeouts, network issues) do NOT match, so the node retries those naturally.
//
// We match the exact phrase "execution reverted" (the standard Ethereum error
// returned by eth_call / eth_estimateGas simulation) and the tx manager's
// "reverted (status 0)" for a mined-but-reverted receipt. We intentionally do
// NOT match plain "reverted" to avoid false-positives from RPC provider error
// messages that happen to contain that word.
func isPermanentRevert(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "execution reverted") || strings.Contains(s, "reverted (status")
}

func isExpectedClaimRevert(err error) bool {
	if err == nil {
		return false
	}
	// Decode the custom-error name first; err.Error() for custom Solidity errors
	// is just "execution reverted" without the name.
	s := decodeContractError(err)
	// Definitive: don't retry. NotEligible & SlotsFull & InvalidPhase mean the
	// committee is decided without us. AlreadyClaimed means we already won.
	// NotRegistered means our node is inactive in the registry.
	// SeedExpired means the seed beacon data is gone — epoch is unrecoverable.
	for _, tok := range []string{
		"NotEligible", "SlotsFull", "AlreadyClaimed",
		"InvalidPhase", "SeedExpired", "NotRegistered",
	} {
		if strings.Contains(s, tok) {
			return true
		}
	}
	return false
}

// decodeContractError attempts to extract a named custom-error identifier from
// an "execution reverted" error returned by go-ethereum. It uses
// ethclient.RevertErrorData to retrieve the raw revert bytes and then looks up
// the 4-byte selector against all errors defined in the DKGManager ABI.
// Returns the error name if found, or the original error string unchanged.
func decodeContractError(err error) string {
	if err == nil {
		return ""
	}
	data, ok := ethclient.RevertErrorData(err)
	if !ok || len(data) < 4 {
		return err.Error()
	}
	var sel [4]byte
	copy(sel[:], data[:4])
	for _, md := range []*bind.MetaData{gtypes.DKGManagerMetaData, gtypes.DKGAppManagerMetaData, gtypes.DKGRegistryMetaData} {
		parsed, parseErr := md.GetAbi()
		if parseErr != nil {
			continue
		}
		if abiErr, lookupErr := parsed.ErrorByID(sel); lookupErr == nil {
			return fmt.Sprintf("execution reverted: %s", abiErr.Name)
		}
	}
	return err.Error()
}
