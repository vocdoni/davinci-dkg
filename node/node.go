package node

import (
	"context"
	"fmt"
	"math/big"
	"math/bits"
	mrand "math/rand"
	"net/url"
	"os"
	"slices"
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
	finalized     map[[12]byte]bool     // tracks epochs we've already attempted to auto-finalize
	terminal      map[[12]byte]bool     // epochs whose key-generation lifecycle needs no more work
	privateShares map[poolSlot]*big.Int // d_{j,i}, one per pool key we needed
	ownContribs   map[[12]byte]*savedContrib
	selectedCache map[[12]byte][]common.Address
	// liveEpochs are the Live epochs (by nonce) this process has seen whose
	// pool still has unclaimed, unactivated keys. Registrations may target
	// any Live epoch, so its reserve has to be kept up even once it is no
	// longer among the newest two; the set is bounded by
	// maxTrackedLiveEpochs (oldest dropped) and by epochs turning terminal.
	liveEpochs map[[12]byte]uint64

	// activateAnchor is the head at which this node first saw a pool key as
	// due for activation; the stagger rotation counts from there. Epochs go
	// Live long before most keys are needed, so there is no chain-wide
	// anchor to share (unlike the finalize window).
	activateAnchor map[poolSlot]uint64
	activateAhead  uint8

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
	// autoCreateEarlyPool is the (epoch, poolNext) the most recent
	// pool-drain-triggered attempt was fired against, so one attempt is made
	// per registration that eats into the reserve.
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
		address:        txm.Address(),
		bjjSecret:      bjjSecret,
		contracts:      c,
		manager:        manager,
		appManager:     appManager,
		registry:       registry,
		txm:            txm,
		runtimes:       runtimes,
		signaled:       make(map[[12]byte]bool),
		contributed:    make(map[[12]byte]bool),
		finalized:      make(map[[12]byte]bool),
		terminal:       make(map[[12]byte]bool),
		privateShares:  make(map[poolSlot]*big.Int),
		ownContribs:    make(map[[12]byte]*savedContrib),
		selectedCache:  make(map[[12]byte][]common.Address),
		liveEpochs:     make(map[[12]byte]uint64),
		activateAnchor: make(map[poolSlot]uint64),
		activateAhead:  max(cfg.ActivateAhead, 1),
		lookback:       cfg.DecryptLookbackBlocks,
		pending:        make(map[ctKey]*ciphertext),
		parked:         make(map[ctKey]*parkedSlot),
		partialDone:    make(map[ctKey]bool),
		served:         make(map[ctKey]uint64),
		shareProofs:    make(map[poolSlot][][32]byte),
		taints:         make(map[taintKey]bool),
		taintFile:      taintPath(cfg.Datadir),
		backoff:        make(map[ctKey]*serviceBackoff),
		inflight:       make(map[ctKey]inflightTx),
		combineJobs:    make(map[ctKey]*combineResult),
		combineSem:     make(chan struct{}, 1),
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
		"pollInterval", cfg.PollInterval,
		"activateAhead", cfg.ActivateAhead)

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
		// Not due by the cadence, but the newest epoch's pool may be nearly
		// claimed out: applications can only register against an activated,
		// unclaimed key, so the next epoch has to exist before the last one
		// goes. The contract allows createEpoch early in exactly that case.
		slot, low := n.poolRunningLow(ctx)
		if !low || n.autoCreateEarlyPool == slot {
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

// poolRunningLow reports whether the newest epoch has fewer than
// ActivateAhead unclaimed pool keys left, and the (epoch, poolNext) pair that
// observation was made against so the early trigger fires once per drain.
// A non-Live epoch reads back as an untouched pool, which is never low.
func (n *Node) poolRunningLow(ctx context.Context) (poolSlot, bool) {
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
	status, err := n.manager.GetPoolStatus(callOpts, epochID)
	if err != nil {
		return poolSlot{}, false
	}
	unclaimed := ccommon.MaxK - int(status.NextIndex)
	return poolSlot{epoch: epochID, key: status.NextIndex}, unclaimed < int(n.activateAhead)
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

// maxTrackedLiveEpochs bounds liveEpochs: each tracked epoch costs a few RPC
// reads per tick, and an epoch whose pool nobody claims from any more is
// better served by the fresh ones than by a node polling it forever.
const maxTrackedLiveEpochs = 6

// tick runs one poll cycle: key-generation lifecycle for the newest epochs
// (plus any older Live epoch whose pool still needs activations), then
// decryption service for every pending ciphertext.
//
// Only the two most recent epochs can still be in CommitteeSelection /
// KeyAssembly (a new epoch cannot start until EPOCH_DURATION_BLOCKS after
// the previous one, and Preparation is shorter than that), so the lifecycle
// scan is O(1) RPC calls no matter how old the deployment is. Decryption
// work is discovered from CiphertextSubmitted events instead of by walking
// epochs, because applications outlive epochs by design.
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
	for _, epochID := range n.epochsToVisit(prefix, epochNonce) {
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

// epochsToVisit is the newest two epochs plus every tracked Live epoch, minus
// the terminal ones, oldest first.
func (n *Node) epochsToVisit(prefix uint32, epochNonce uint64) [][12]byte {
	byNonce := make(map[uint64][12]byte, len(n.liveEpochs)+2)
	for id, nonce := range n.liveEpochs {
		byNonce[nonce] = id
	}
	first := uint64(1)
	if epochNonce > 1 {
		first = epochNonce - 1
	}
	for i := first; i <= epochNonce; i++ {
		byNonce[i] = web3.EpochID(prefix, i)
	}
	nonces := make([]uint64, 0, len(byNonce))
	for nonce := range byNonce {
		nonces = append(nonces, nonce)
	}
	slices.Sort(nonces)
	out := make([][12]byte, 0, len(nonces))
	for _, nonce := range nonces {
		if id := byNonce[nonce]; !n.terminal[id] {
			out = append(out, id)
		}
	}
	return out
}

// trackLive remembers a Live epoch that still needs activations so later
// ticks keep visiting it, dropping the oldest tracked epoch past the bound.
func (n *Node) trackLive(epochID [12]byte, nonce uint64) {
	if _, ok := n.liveEpochs[epochID]; ok {
		return
	}
	for len(n.liveEpochs) >= maxTrackedLiveEpochs {
		var oldest [12]byte
		oldestNonce := ^uint64(0)
		for id, nonce := range n.liveEpochs {
			if nonce < oldestNonce {
				oldest, oldestNonce = id, nonce
			}
		}
		delete(n.liveEpochs, oldest)
	}
	n.liveEpochs[epochID] = nonce
}

// finish marks an epoch's key-generation lifecycle as needing no more work.
func (n *Node) finish(epochID [12]byte) {
	n.terminal[epochID] = true
	delete(n.liveEpochs, epochID)
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
		//                  then race on a deterministic stagger to call finalizeEpoch.
		selected, err := n.contracts.SelectedParticipants(ctx, epochID)
		if err != nil {
			return fmt.Errorf("selected participants: %w", err)
		}
		idx := myIndex(selected, n.address)
		if idx == 0 {
			return nil // not selected for this epoch
		}
		if err := n.doContribution(ctx, epochID, idx, epoch, selected); err != nil {
			return err
		}
		// After contributing, every selected participant rotates through a
		// deterministic finalize stagger so exactly one node submits at a time
		// (any race-loser sees AlreadyFinalized and stops).
		if err := n.tryAutoFinalize(ctx, epochID, idx, selected); err != nil {
			log.Warnw("auto-finalize attempt failed",
				"epoch", roundHex(epochID), "err", err)
		}
		return nil

	case epochLive: // the pool is being activated; ciphertexts are served by the decryption scanner
		selected, err := n.selected(ctx, epochID)
		if err != nil {
			return err
		}
		idx := myIndex(selected, n.address)
		if idx == 0 {
			n.finish(epochID)
			return nil // not selected for this epoch
		}
		n.trackLive(epochID, epoch.Nonce)
		return n.tryActivatePoolKeys(ctx, epochID, idx, epoch, selected)

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

	recipientIdxs := make([]uint16, committeeSize)
	recipientKeys := make([]nodetypes.NodeKey, committeeSize)
	for i := uint16(0); i < committeeSize; i++ {
		recipientIdxs[i] = i + 1
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

// staggerBlocks is the per-slot delay between successive finalize attempts in
// the rotation. With STAGGER_BLOCKS=3, slot 0 attempts at finalizeNotBefore,
// slot 1 at +3, slot 2 at +6, etc.
const staggerBlocks = 3

// tryAutoFinalize is called by every selected participant after their
// contribution lands. It computes the participant's slot in a epoch-specific
// rotation derived from the lottery seed, waits until that slot's window
// opens, then races with the other slots to submit finalizeEpoch. The first
// to land wins; everyone else sees AlreadyFinalized and stops.
//
// Pre-checks short-circuit cheaply: if the epoch is already past Contribution
// or we've already attempted, skip. If the contribution count is below
// minValidContributions or block.number is below the wait-until block, defer
// to the next tick.
func (n *Node) tryAutoFinalize(
	ctx context.Context,
	epochID [12]byte,
	myIdx uint16,
	selected []common.Address,
) error {
	if n.finalized[epochID] {
		return nil
	}
	epoch, err := n.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return fmt.Errorf("get epoch: %w", err)
	}
	if epoch.Status != 2 { // not in Contribution any more — finalized, aborted, etc.
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

	// Re-read status one last time to avoid burning a proof for a epoch
	// another node has just finalized.
	epoch, err = n.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return fmt.Errorf("re-read epoch: %w", err)
	}
	if epoch.Status != 2 {
		n.finalized[epochID] = true
		return nil
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
	)

	gasUsed, err := finalizer.FinalizeEpoch(ctx, n.contracts, n.manager, n.txm, epochID)
	if err != nil {
		// Only a lost race is final: AlreadyLive from the contract, or the
		// epoch having left KeyAssembly. Anything else (a bad proof, a
		// mined revert, an RPC hiccup) is retried on the next tick.
		reason := decodeContractError(err)
		status := epochKeyAssembly // unknown → assume still open so we retry
		if cur, rerr := n.contracts.GetEpoch(ctx, epochID); rerr == nil {
			status = cur.Status
		}
		if finalizeRaceLost(reason, status) {
			log.Infow("auto-finalize: another node beat us", "epoch", roundHex(epochID), "reason", reason)
			n.finalized[epochID] = true
			return nil
		}
		return fmt.Errorf("auto-finalize submit failed (will retry next tick): %s", reason)
	}
	n.finalized[epochID] = true
	log.Infow("auto-finalize: epoch finalized by us",
		"epoch", roundHex(epochID), "gas", gasUsed)
	return nil
}

// tryActivatePoolKeys keeps the epoch's pool stocked while it is Live: key 0
// as soon as the epoch goes Live, then enough keys ahead of the claimed ones
// that a registration never waits for a proof. Committee members take turns
// in a seed-derived rotation like auto-finalize; activation is permissionless
// so a race loser only pays a revert.
//
// The rotation is anchored on the head at which this node first saw the key
// as due rather than on a chain-wide block: epochs stay Live indefinitely and
// keys past the first are needed whenever registrations drain the pool, so
// there is no shared anchor to count from.
func (n *Node) tryActivatePoolKeys(
	ctx context.Context,
	epochID [12]byte,
	myIdx uint16,
	epoch epochView,
	selected []common.Address,
) error {
	status, err := n.manager.GetPoolStatus(&bind.CallOpts{Context: ctx}, epochID)
	if err != nil {
		return fmt.Errorf("pool status: %w", err)
	}
	next, ok := nextKeyToActivate(status.NextIndex, status.Activated, n.activateAhead)
	if !ok {
		if int(status.NextIndex) >= ccommon.MaxK || bits.OnesCount8(status.Activated) >= ccommon.MaxK {
			n.finish(epochID) // the whole pool is claimed or up; ciphertexts are the scanner's job
		}
		return nil
	}
	keyIndex := next
	slot := poolSlot{epoch: epochID, key: keyIndex}

	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("read head: %w", err)
	}
	anchor, seen := n.activateAnchor[slot]
	if !seen {
		anchor = head
		n.activateAnchor[slot] = anchor
	}
	mySlot := staggerSlot(epoch.Seed, uint64(keyIndex), myIdx, uint16(len(selected)))
	if head < anchor+mySlot*staggerBlocks {
		return nil // not our turn yet
	}

	n.critical.Add(1)
	defer n.critical.Add(-1)
	log.Infow("pool key activation: my turn",
		"epoch", roundHex(epochID), "key", keyIndex, "myIdx", myIdx, "mySlot", mySlot, "head", head)

	res, err := finalizer.ProveAndSubmitActivation(
		ctx, n.contracts, n.manager, n.txm, n.runtimes.poolKey, epochID,
		epoch.Policy.Threshold, epoch.Policy.CommitteeSize, selected, keyIndex,
	)
	if err != nil {
		reason := decodeContractError(err)
		if strings.Contains(reason, "PoolKeyAlreadyActive") {
			log.Infow("pool key activation: another node beat us", "epoch", roundHex(epochID), "key", keyIndex)
			delete(n.activateAnchor, slot)
			return nil
		}
		return fmt.Errorf("activate pool key %d (will retry next tick): %s", keyIndex, reason)
	}
	delete(n.activateAnchor, slot)
	log.Infow("pool key activated",
		"epoch", roundHex(epochID), "key", keyIndex, "x", res.PoolKey.X, "gas", res.GasUsed)
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
		share, err := n.recoverShareFrom(ctx, epochID, addr, contribIdx, keyIndex, roundHash, myIdx, fromBlock)
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

// recoverShareFrom fetches the submitContribution tx calldata for `contributor`
// (located through the ContributionSubmitted event log) and decrypts the
// share slot destined for myIdx under pool key `keyIndex`.
func (n *Node) recoverShareFrom(
	ctx context.Context,
	epochID [12]byte,
	contributor common.Address,
	contribIdx uint16,
	keyIndex uint8,
	roundHash *big.Int,
	myIdx uint16,
	fromBlock uint64,
) (*big.Int, error) {
	data, err := finalizer.ContributionCalldata(ctx, n.contracts.Client(), n.manager, epochID, contributor, fromBlock)
	if err != nil {
		return nil, err
	}
	eph, masked, recipIdxs, err := decodeContributionTranscript(data)
	if err != nil {
		return nil, fmt.Errorf("decode contribution transcript: %w", err)
	}
	if int(keyIndex) >= len(masked) {
		return nil, fmt.Errorf("pool key %d is outside the transcript", keyIndex)
	}
	for slot, ridx := range recipIdxs {
		if ridx != myIdx || slot >= len(eph) || slot >= len(masked[keyIndex]) {
			continue
		}
		ct := shareenc.Ciphertext{
			Ephemeral:   nodetypes.CurvePoint{X: eph[slot][0], Y: eph[slot][1]},
			MaskedShare: masked[keyIndex][slot],
		}
		return shareenc.DecryptShareRoundHash(roundHash, contribIdx, myIdx, keyIndex, ct, n.bjjSecret)
	}
	return nil, fmt.Errorf("no share slot for index %d in contribution of %s", myIdx, contributor.Hex())
}

// decodeContributionTranscript extracts (ephemerals, maskedShares,
// recipientIndexes) from raw submitContribution calldata. maskedShares is
// indexed by pool key then recipient slot. Transcript layout (N =
// circuits/common.MaxN, K = circuits/common.MaxK, words of 32 bytes):
//
//	[0, 2KN)        commitments, key-major     [2KN, 2KN+N)     recipientIndexes
//	[2KN+N, 2KN+3N) recipientPubKeys           [2KN+3N, 2KN+5N) ephemerals
//	[2KN+5N, 3KN+5N) maskedShares, key-major
func decodeContributionTranscript(
	data []byte,
) (ephemerals [][2]*big.Int, maskedShares [][]*big.Int, recipientIndexes []uint16, err error) {
	transcript, err := finalizer.ContributionTranscript(data)
	if err != nil {
		return nil, nil, nil, err
	}
	const (
		nn      = ccommon.MaxN
		kk      = ccommon.MaxK
		idxOff  = 2 * kk * nn
		ephOff  = idxOff + 3*nn
		maskOff = idxOff + 5*nn
	)
	word := func(i int) *big.Int { return new(big.Int).SetBytes(transcript[i*32 : (i+1)*32]) }
	ridxs := make([]uint16, nn)
	ephs := make([][2]*big.Int, nn)
	for i := range nn {
		ridxs[i] = uint16(word(idxOff + i).Uint64())
		ephs[i] = [2]*big.Int{word(ephOff + 2*i), word(ephOff + 2*i + 1)}
	}
	masked := make([][]*big.Int, kk)
	for j := range kk {
		masked[j] = make([]*big.Int, nn)
		for i := range nn {
			masked[j][i] = word(maskOff + j*nn + i)
		}
	}
	return ephs, masked, ridxs, nil
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

// nextKeyToActivate picks the pool key to activate, if any. The reserve is
// the run of *contiguous* activated keys from nextIndex: registration claims
// keys in order and reverts PoolKeyNotActive on the first gap, so an
// activated key past a gap is no reserve at all (which a popcount would
// count). Every key in [nextIndex, min(MaxK, nextIndex+ahead)) must be set;
// the first unset one at or after nextIndex is the key to activate.
func nextKeyToActivate(nextIndex, activated, ahead uint8) (uint8, bool) {
	want := min(ccommon.MaxK, int(nextIndex)+int(ahead))
	for j := int(nextIndex); j < want; j++ {
		if activated&(1<<j) == 0 {
			return uint8(j), true
		}
	}
	return 0, false
}

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
