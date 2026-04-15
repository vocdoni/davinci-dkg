package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/davinci-dkg/circuits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
	"github.com/vocdoni/davinci-node/crypto/ecc"
)

// bjjKeyDomain must match tests/helpers/nodekeys.go so registry keys are consistent.
const bjjKeyDomain = "davinci-dkg/bjj-key/v1"

// CiphertextFile is written to sharedDir/ciphertext-{roundHex}.json by the
// dkg-runner or by the webapp playground via the /api/ciphertext/ endpoint.
// C1 is the ephemeral ciphertext half (k*G); C2 is the encrypted half (m*G + k*PubKey).
// C2 is needed by the first node that combines the partial decryptions.
type CiphertextFile struct {
	CiphertextIndex uint16 `json:"ciphertext_index"`
	C1X             string `json:"c1x"`
	C1Y             string `json:"c1y"`
	C2X             string `json:"c2x,omitempty"`
	C2Y             string `json:"c2y,omitempty"`
}

// savedContrib caches data from the node's own submitted contribution
// so it can compute the own-polynomial component of d_i offline.
type savedContrib struct {
	coefficients     []*big.Int
	recipientIndexes []uint16
	recipientKeys    []nodetypes.NodeKey
}

// Node participates in every DKG round it can find on chain.
type Node struct {
	address   common.Address
	privKey   string
	bjjSecret *big.Int

	contracts *web3.Contracts
	manager   *gtypes.DKGManager
	registry  *gtypes.DKGRegistry
	txm       *txmanager.Manager

	sharedDir string

	// per-round local state
	signaled      map[[12]byte]bool
	contributed   map[[12]byte]bool
	decrypted     map[[12]byte]map[uint16]bool
	combined      map[[12]byte]bool
	terminal      map[[12]byte]bool // rounds where no further work is possible
	privateShares map[[12]byte]*big.Int
	ownContribs   map[[12]byte]*savedContrib
}

// newNode constructs a Node from the daemon config.
func newNode(cfg *Config) (*Node, error) {
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
	manager, err := gtypes.NewDKGManager(c.Addresses.Manager, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("manager binding: %w", err)
	}
	registry, err := gtypes.NewDKGRegistry(c.Addresses.Registry, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("registry binding: %w", err)
	}

	bjjSecret, err := deriveBJJSecret(cfg.PrivKey)
	if err != nil {
		return nil, fmt.Errorf("derive bjj key: %w", err)
	}

	// Override artifact path from env if set.
	if d := os.Getenv("DAVINCI_DKG_ARTIFACTS_DIR"); d != "" {
		circuits.BaseDir = d
	}

	return &Node{
		address:       txm.Address(),
		privKey:       cfg.PrivKey,
		bjjSecret:     bjjSecret,
		contracts:     c,
		manager:       manager,
		registry:      registry,
		txm:           txm,
		sharedDir:     cfg.SharedDir,
		signaled:      make(map[[12]byte]bool),
		contributed:   make(map[[12]byte]bool),
		decrypted:     make(map[[12]byte]map[uint16]bool),
		combined:      make(map[[12]byte]bool),
		terminal:      make(map[[12]byte]bool),
		privateShares: make(map[[12]byte]*big.Int),
		ownContribs:   make(map[[12]byte]*savedContrib),
	}, nil
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
		"datadir", cfg.Datadir,
		"sharedDir", cfg.SharedDir)
	log.Infow("config: chain connection",
		"network", cfg.Web3.Network,
		"chainId", n.contracts.ChainID,
		"rpc", cfg.Web3.RPC[0],
		"gasMultiplier", cfg.Web3.GasMultiplier)
	log.Infow("config: contracts",
		"registry", n.contracts.Addresses.Registry,
		"manager", cfg.ManagerAddr)
	log.Infow("config: participation",
		"pollInterval", cfg.PollInterval)
	log.Infow("config: explorer webapp",
		"enabled", cfg.Webapp.Enabled,
		"listen", cfg.Webapp.Listen,
		"publicRpc", cfg.Webapp.PublicRPC)

	// ── on-chain state ───────────────────────────────────────────────────
	callOpts := &bind.CallOpts{Context: ctx}

	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		log.Warnw("startup: failed to read chain head", "err", err)
	}

	prefix, err := n.manager.ROUNDPREFIX(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read ROUND_PREFIX", "err", err)
	}
	roundNonce, err := n.manager.RoundNonce(callOpts)
	if err != nil {
		log.Warnw("startup: failed to read roundNonce", "err", err)
	}
	log.Infow("chain: snapshot",
		"head", head,
		"roundPrefix", fmt.Sprintf("0x%08x", prefix),
		"roundNonce", roundNonce)

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

// RoundStatus enum mirror (matches IDKGRegistry.NodeStatus in Solidity).
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

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for registration: %w", err)
	}
	var tx *ethtypes.Transaction
	switch existing.Status {
	case nodeStatusNone:
		log.Infow("registering bjj key on-chain (first time)",
			"address", n.address)
		tx, err = n.registry.RegisterKey(auth, wantX, wantY)
	case nodeStatusInactive:
		log.Warnw("node is INACTIVE on-chain, reactivating via updateKey",
			"address", n.address,
			"lastActiveBlock", existing.LastActiveBlock)
		tx, err = n.registry.UpdateKey(auth, wantX, wantY)
	default: // ACTIVE but stale key
		log.Infow("rotating bjj key on-chain",
			"address", n.address,
			"oldPubX", existing.PubX, "newPubX", wantX)
		tx, err = n.registry.UpdateKey(auth, wantX, wantY)
	}
	if err != nil {
		return fmt.Errorf("register/update key tx: %w", err)
	}
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
//     — e.g. because the reaper ran before our first lucky round —
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
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		return fmt.Errorf("wait reactivate: %w", err)
	}
	log.Infow("liveness: reactivate confirmed", "address", n.address)
	return nil
}

// Run is the main participation loop; blocks until ctx is done.
func (n *Node) Run(ctx context.Context, pollInterval time.Duration) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	// Emit balance and gas-spent info every 10 minutes regardless of poll interval.
	fundsTicker := time.NewTicker(10 * time.Minute)
	defer fundsTicker.Stop()
	log.Infow("node running", "address", n.address, "poll", pollInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-fundsTicker.C:
			n.logFunds(ctx)
		case <-ticker.C:
			// Keep our on-chain liveness row healthy before scanning rounds.
			// This guarantees heartbeat()/reactivate() fire even when there
			// are no active rounds to participate in.
			n.maintainLiveness(ctx)
			if err := n.tick(ctx); err != nil {
				log.Errorw(err, "participation tick")
			}
		}
	}
}

func (n *Node) tick(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}
	roundNonce, err := n.manager.RoundNonce(callOpts)
	if err != nil {
		return fmt.Errorf("round nonce: %w", err)
	}
	prefix, err := n.manager.ROUNDPREFIX(callOpts)
	if err != nil {
		return fmt.Errorf("round prefix: %w", err)
	}
	for i := uint64(1); i <= roundNonce; i++ {
		roundID := makeRoundID(prefix, i)
		if n.terminal[roundID] {
			continue // round is in a terminal state; no further work possible
		}
		if err := n.participate(ctx, roundID, callOpts); err != nil {
			log.Warnw("participate failed", "round", roundHex(roundID), "err", err)
		}
	}
	return nil
}

func (n *Node) participate(ctx context.Context, roundID [12]byte, callOpts *bind.CallOpts) error {
	round, err := n.contracts.GetRound(ctx, roundID)
	if err != nil {
		return fmt.Errorf("get round: %w", err)
	}
	switch round.Status {
	case 0: // None — round slot exists on-chain but is uninitialised (should not happen)
		return nil

	case 1: // Registration — try to claim a slot in the lottery
		return n.doClaimSlot(ctx, roundID, round)

	case 2: // Contribution — selected participants submit ZK shares
		selected, err := n.contracts.SelectedParticipants(ctx, roundID)
		if err != nil {
			return fmt.Errorf("selected participants: %w", err)
		}
		idx := myIndex(selected, n.address)
		if idx == 0 {
			return nil // not selected for this round
		}
		return n.doContribution(ctx, roundID, idx, round, selected)

	case 3: // Finalized — partial decryptions, then combine
		selected, err := n.contracts.SelectedParticipants(ctx, roundID)
		if err != nil {
			return fmt.Errorf("selected participants: %w", err)
		}
		idx := myIndex(selected, n.address)
		if idx == 0 {
			// Not a selected participant; try to combine if threshold partials exist.
			if err := n.doCombineDecryption(ctx, roundID, round, selected, callOpts); err != nil {
				return err
			}
		} else {
			if err := n.doDecryption(ctx, roundID, idx, round, selected, callOpts); err != nil {
				return err
			}
			if err := n.doCombineDecryption(ctx, roundID, round, selected, callOpts); err != nil {
				return err
			}
		}
		// Once combined, this round requires no further work — mark terminal so
		// future ticks skip the GetRound RPC call entirely.
		if n.combined[roundID] {
			n.terminal[roundID] = true
			log.Infow("round fully processed, marked terminal",
				"round", roundHex(roundID))
		}
		return nil

	case 4: // Aborted — organizer cancelled the round
		log.Warnw("round aborted — no further participation possible",
			"round", roundHex(roundID))
		n.terminal[roundID] = true
		return nil

	case 5: // Completed — secret reconstructed (disclosureAllowed path)
		log.Infow("round completed (secret disclosed)", "round", roundHex(roundID))
		n.terminal[roundID] = true
		return nil

	default:
		log.Warnw("unknown round status — skipping", "round", roundHex(roundID), "status", round.Status)
		return nil
	}
}

// ---- Lottery slot claim ----

// doClaimSlot races to claim a committee slot for the round. Eligibility is
// derived deterministically from the round seed; if the seed has not been
// resolved yet (block.number < round.SeedBlock), the call will revert with
// SeedNotReady and we'll retry on the next poll. If the node is not eligible we
// silently no-op for the rest of the round.
func (n *Node) doClaimSlot(ctx context.Context, roundID [12]byte, round web3.RoundView) error {
	if n.signaled[roundID] {
		return nil
	}

	// Read the current head once; used by all pre-flight checks below.
	head, headErr := n.contracts.Client().BlockNumber(ctx)

	// Pre-flight: if the seed block hasn't been reached yet, skip this tick silently.
	// The slot lottery cannot be resolved before the seed is committed, and the
	// contract will revert with SeedNotReady — but many RPC providers return a
	// bare "execution reverted" without the named error, generating false warnings.
	if headErr == nil && head > 0 && round.SeedBlock > 0 && head < round.SeedBlock {
		log.Debugw("claim slot: seed block not yet reached — waiting",
			"round", roundHex(roundID),
			"head", head,
			"seedBlock", round.SeedBlock)
		return nil
	}

	// Pre-flight: if the committee is already full, we were not selected.
	if round.ClaimedCount >= round.Policy.CommitteeSize {
		log.Infow("claim slot: committee already full — not selected for this round",
			"round", roundHex(roundID),
			"claimed", round.ClaimedCount,
			"size", round.Policy.CommitteeSize)
		n.signaled[roundID] = true
		return nil
	}

	// Pre-flight: check registration deadline before sending any tx.
	if headErr == nil {
		if head >= round.Policy.RegistrationDeadlineBlock {
			log.Warnw("registration deadline already passed — skipping slot claim",
				"round", roundHex(roundID),
				"head", head,
				"deadline", round.Policy.RegistrationDeadlineBlock)
			n.signaled[roundID] = true
			return nil
		}
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := n.manager.ClaimSlot(auth, roundID)
	if err != nil {
		// SeedNotReady: the seed block hasn't been mined yet. Retry next poll
		// without setting signaled so we keep trying until the seed arrives.
		if strings.Contains(err.Error(), "SeedNotReady") {
			log.Debugw("claim slot: seed not ready yet, retrying next poll", "round", roundHex(roundID))
			return nil
		}
		// Definitively final reverts: the committee is decided without us.
		// Set signaled so we stop sending txs for this round.
		if isExpectedClaimRevert(err) {
			log.Debugw("claim slot: not selected for committee", "round", roundHex(roundID), "reason", decodeContractError(err))
			n.signaled[roundID] = true
			return nil
		}
		// Unexpected permanent revert — all pre-flights passed but the contract
		// still rejected us. Accept the result and stop retrying.
		if isPermanentRevert(err) {
			log.Warnw("claim slot: unexpected permanent revert — marking as not selected",
				"round", roundHex(roundID), "err", decodeContractError(err))
			n.signaled[roundID] = true
			return nil
		}
		return fmt.Errorf("claim slot: %w", err)
	}
	if err := n.txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		// On-chain revert — committee race or not eligible. Accept and stop retrying.
		if isPermanentRevert(err) {
			log.Infow("claim slot: tx reverted on-chain (race condition or not eligible) — will not retry",
				"round", roundHex(roundID), "err", err)
			n.signaled[roundID] = true
			return nil
		}
		return fmt.Errorf("wait claim slot tx: %w", err)
	}
	n.signaled[roundID] = true
	log.Infow("slot claimed", "round", roundHex(roundID))
	return nil
}

// ---- Contribution ----

func (n *Node) doContribution(
	ctx context.Context,
	roundID [12]byte,
	idx uint16,
	round web3.RoundView,
	selected []common.Address,
) error {
	if n.contributed[roundID] {
		return nil
	}
	// Check on-chain (handles restarts).
	rec, err := n.manager.GetContribution(&bind.CallOpts{Context: ctx}, roundID, n.address)
	if err == nil && rec.Accepted {
		log.Infow("contribution already accepted on-chain", "round", roundHex(roundID))
		n.contributed[roundID] = true
		return nil
	}

	// Pre-flight: check contribution deadline before burning time on ZK proof.
	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		log.Warnw("doContribution: failed to read block number", "round", roundHex(roundID), "err", err)
		// Proceed optimistically; worst case the tx reverts and we catch it below.
	} else if head >= round.Policy.ContributionDeadlineBlock {
		log.Warnw("contribution deadline already passed — skipping round",
			"round", roundHex(roundID),
			"head", head,
			"deadline", round.Policy.ContributionDeadlineBlock)
		n.contributed[roundID] = true
		return nil
	}

	threshold := round.Policy.Threshold
	committeeSize := round.Policy.CommitteeSize

	roundHash := roundScalar(roundID)
	coeffs := make([]*big.Int, threshold)
	for i := range coeffs {
		// Use 128-bit random coefficients to avoid overflowing the BabyJubJub
		// subgroup order during circuit polynomial evaluation (circuit evaluates mod BN254).
		c, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
		if err != nil {
			return err
		}
		if c.Sign() == 0 {
			c.SetInt64(1)
		}
		coeffs[i] = c
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

	// Deterministic nonces
	nonces := make([]*big.Int, committeeSize)
	for i := range nonces {
		nonces[i] = big.NewInt(int64(1000 + recipientIdxs[i]))
	}

	log.Infow("contribution assignment",
		"round", roundHex(roundID),
		"index", idx,
		"threshold", threshold,
		"committeeSize", committeeSize,
		"deadline", round.Policy.ContributionDeadlineBlock,
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
	runtime, err := contribution.Artifacts.LoadOrSetupForCircuit(ctx, &contribution.ContributionCircuit{})
	if err != nil {
		return fmt.Errorf("load contribution circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
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
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return fmt.Errorf("encode contribution transcript: %w", err)
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for contribution: %w", err)
	}
	tx, err := n.manager.SubmitContribution(
		auth, roundID, idx,
		common.BigToHash(pi.CommitmentHash),
		common.BigToHash(pi.ShareHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		if isPermanentRevert(err) {
			log.Warnw("contribution tx permanently rejected — will not retry this round",
				"round", roundHex(roundID), "err", err)
			n.contributed[roundID] = true
		}
		return fmt.Errorf("submit contribution: %w", err)
	}
	if err := n.txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		if isPermanentRevert(err) {
			log.Warnw("contribution tx reverted on-chain — will not retry this round",
				"round", roundHex(roundID), "err", err)
			n.contributed[roundID] = true
		}
		return fmt.Errorf("wait contribution tx: %w", err)
	}
	n.contributed[roundID] = true
	n.ownContribs[roundID] = &savedContrib{
		coefficients:     coeffs,
		recipientIndexes: recipientIdxs,
		recipientKeys:    recipientKeys,
	}
	log.Infow("contribution submitted", "round", roundHex(roundID), "index", idx)
	return nil
}

// ---- Partial decryption ----

func (n *Node) doDecryption(
	ctx context.Context,
	roundID [12]byte,
	idx uint16,
	round web3.RoundView,
	selected []common.Address,
	callOpts *bind.CallOpts,
) error {
	const ctIdx = uint16(1)
	if n.decrypted[roundID] == nil {
		n.decrypted[roundID] = make(map[uint16]bool)
	}
	if n.decrypted[roundID][ctIdx] {
		return nil
	}
	rec, err := n.manager.GetPartialDecryption(callOpts, roundID, n.address, ctIdx)
	if err == nil && rec.Accepted {
		n.decrypted[roundID][ctIdx] = true
		return nil
	}

	ct, err := n.readCiphertext(roundID)
	if err != nil {
		return nil // ciphertext not written yet — wait
	}

	dShare, err := n.buildPrivateShare(ctx, roundID, idx, selected, round, callOpts)
	if err != nil {
		return fmt.Errorf("build private share: %w", err)
	}

	c1X, ok1 := new(big.Int).SetString(ct.C1X, 10)
	c1Y, ok2 := new(big.Int).SetString(ct.C1Y, 10)
	if !ok1 || !ok2 {
		return fmt.Errorf("parse C1 from ciphertext file")
	}

	nonce := n.decNonce(roundID, idx, ctIdx)
	asgn := partialdecrypt.Assignment{
		RoundHash:        roundScalar(roundID),
		ParticipantIndex: idx,
		Base:             nodetypes.CurvePoint{X: c1X, Y: c1Y},
		Secret:           dShare,
		Nonce:            nonce,
	}
	witness, pi, err := partialdecrypt.BuildWitness(asgn)
	if err != nil {
		return fmt.Errorf("build partial decrypt witness: %w", err)
	}
	runtime, err := partialdecrypt.Artifacts.LoadOrSetupForCircuit(ctx, &partialdecrypt.PartialDecryptCircuit{})
	if err != nil {
		return fmt.Errorf("load partial decrypt circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return fmt.Errorf("prove partial decrypt: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return fmt.Errorf("marshal partial decrypt proof: %w", err)
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return fmt.Errorf("encode partial decrypt public witness: %w", err)
	}
	dHash := ethcrypto.Keccak256Hash(
		common.LeftPadBytes(pi.Delta.X.Bytes(), 32),
		common.LeftPadBytes(pi.Delta.Y.Bytes(), 32),
	)

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for partial decryption: %w", err)
	}
	tx, err := n.manager.SubmitPartialDecryption(auth, roundID, idx, ctIdx, dHash, proofBytes, inputBytes)
	if err != nil {
		if isPermanentRevert(err) {
			log.Warnw("partial decryption tx permanently rejected — will not retry",
				"round", roundHex(roundID), "err", err)
			n.decrypted[roundID][ctIdx] = true
		}
		return fmt.Errorf("submit partial decryption: %w", err)
	}
	if err := n.txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		if isPermanentRevert(err) {
			log.Warnw("partial decryption tx reverted on-chain — will not retry",
				"round", roundHex(roundID), "err", err)
			n.decrypted[roundID][ctIdx] = true
		}
		return fmt.Errorf("wait partial decryption tx: %w", err)
	}
	n.decrypted[roundID][ctIdx] = true
	log.Infow("partial decryption submitted", "round", roundHex(roundID), "index", idx)
	return nil
}

// buildPrivateShare computes d_i = Σ_j f_j(i) by:
//   - Own contribution (in-memory cache): evaluate own polynomial directly
//   - Own contribution (after restart, cache lost): fall back to calldata recovery
//   - Other contributions: scan on-chain txs for calldata and decrypt
//
// fromBlock is used as the lower bound for the calldata scan; passing the round's
// registration deadline block keeps the scan tight and avoids the O(n) hit of
// scanning from genesis.
func (n *Node) buildPrivateShare(
	ctx context.Context,
	roundID [12]byte,
	myIdx uint16,
	selected []common.Address,
	round web3.RoundView,
	callOpts *bind.CallOpts,
) (*big.Int, error) {
	if s, ok := n.privateShares[roundID]; ok {
		return s, nil
	}
	modulus := group.ScalarField()
	roundHash := roundScalar(roundID)
	total := new(big.Int)

	// Use the registration deadline as the natural lower bound for calldata scans.
	// Contributions must appear between registration end and contribution deadline,
	// so this avoids scanning blocks before the round even started.
	fromBlock := round.Policy.RegistrationDeadlineBlock

	recovered := 0
	expected := 0
	for i, addr := range selected {
		contribIdx := uint16(i + 1)
		rec, err := n.manager.GetContribution(callOpts, roundID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		expected++

		if addr == n.address {
			// Own contribution: use in-memory polynomial if available (normal path).
			if sc := n.ownContribs[roundID]; sc != nil {
				x := big.NewInt(int64(myIdx))
				share, err := ccommon.EvaluatePolynomialNative(sc.coefficients, x)
				if err == nil {
					total.Add(total, share)
					total.Mod(total, modulus)
					recovered++
					continue
				}
				log.Warnw("own polynomial evaluation failed, falling back to calldata",
					"round", roundHex(roundID), "err", err)
			} else {
				log.Warnw("own contribution cache missing (node restarted?), recovering from calldata",
					"round", roundHex(roundID))
			}
			// Fall through to calldata recovery for own contribution too.
		}

		// Recover encrypted share from on-chain calldata.
		share, err := n.recoverShareFrom(ctx, roundID, addr, contribIdx, roundHash, myIdx, fromBlock)
		if err != nil {
			log.Warnw("share recovery failed — contribution will be excluded from private share",
				"round", roundHex(roundID), "contributor", addr.Hex(), "idx", contribIdx, "err", err)
			continue
		}
		total.Add(total, share)
		total.Mod(total, modulus)
		recovered++
	}

	log.Infow("private share built",
		"round", roundHex(roundID),
		"recovered", recovered,
		"expected", expected,
		"myIdx", myIdx)

	if recovered == 0 {
		return nil, fmt.Errorf("private share: recovered 0/%d contributions — calldata not yet available or no contributions accepted", expected)
	}
	if total.Sign() == 0 {
		return nil, fmt.Errorf("private share is zero after %d/%d contributions — possible Shamir evaluation issue", recovered, expected)
	}
	n.privateShares[roundID] = total
	return total, nil
}

// recoverShareFrom fetches the submitContribution tx calldata for `contributor`
// and decrypts the share slot destined for myIdx.
//
// fromBlock is the earliest block to scan. Pass the round's
// registrationDeadlineBlock so the scan starts at the earliest plausible point
// instead of walking back an arbitrary fixed number of blocks from the head.
func (n *Node) recoverShareFrom(
	ctx context.Context,
	roundID [12]byte,
	contributor common.Address,
	contribIdx uint16,
	roundHash *big.Int,
	myIdx uint16,
	fromBlock uint64,
) (*big.Int, error) {
	client := n.contracts.Client()
	chainID := new(big.Int).SetUint64(n.contracts.ChainID)

	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	start := fromBlock

	managerAddr := n.contracts.Addresses.Manager
	signer := ethtypes.NewCancunSigner(chainID)
	for blk := start; blk <= latest; blk++ {
		block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blk))
		if err != nil {
			continue
		}
		for _, tx := range block.Transactions() {
			if tx.To() == nil || *tx.To() != managerAddr {
				continue
			}
			from, err := ethtypes.Sender(signer, tx)
			if err != nil || from != contributor {
				continue
			}
			data := tx.Data()
			if len(data) < 4 {
				continue
			}
			// submitContribution selector: first 4 bytes
			// We try to decode and ignore on error
			eph, masked, recipIdxs, err := decodeContributionTranscript(data)
			if err != nil {
				continue
			}
			for slot, ridx := range recipIdxs {
				if ridx != myIdx {
					continue
				}
				if slot >= len(eph) || slot >= len(masked) {
					continue
				}
				ct := shareenc.Ciphertext{
					Ephemeral:   nodetypes.CurvePoint{X: eph[slot][0], Y: eph[slot][1]},
					MaskedShare: masked[slot],
				}
				share, err := shareenc.DecryptShareRoundHash(roundHash, contribIdx, myIdx, ct, n.bjjSecret)
				if err != nil {
					continue
				}
				return share, nil
			}
		}
	}
	return nil, fmt.Errorf("share not found in calldata for round %s contributor %s", roundHex(roundID), contributor.Hex())
}

// decodeContributionTranscript extracts (ephemerals, maskedShares, recipientIndexes)
// from the raw calldata of a submitContribution transaction.
//
// submitContribution ABI:
//
//	(bytes12,uint16,bytes32,bytes32,bytes transcript,bytes proof,bytes input)
//
// transcript layout = abi.encode(
//
//	uint256[2N] commitmentPoints,
//	uint256[N]  recipientIndexes,
//	uint256[2N] recipientPubKeys,
//	uint256[2N] ephemerals,
//	uint256[N]  maskedShares,
//
// )
// where N = circuits/common.MaxN.
func decodeContributionTranscript(data []byte) (ephemerals [][2]*big.Int, maskedShares []*big.Int, recipientIndexes []uint16, err error) {
	if len(data) < 4 {
		return nil, nil, nil, fmt.Errorf("calldata too short")
	}
	// Skip 4-byte selector
	payload := data[4:]

	// ABI-decode 7 parameters; each static head is 32 bytes.
	// roundId (bytes12)=32, contributorIndex (uint16)=32, commitmentsHash=32, encSharesHash=32
	// transcript=offset(32), proof=offset(32), input=offset(32)
	// Total head = 7*32 = 224 bytes
	if len(payload) < 224 {
		return nil, nil, nil, fmt.Errorf("payload too short for head")
	}
	transcriptOffset := int(new(big.Int).SetBytes(padTo32(payload[128:160])).Int64())
	if transcriptOffset+32 > len(payload) {
		return nil, nil, nil, fmt.Errorf("transcript offset out of range")
	}
	transcriptLen := int(new(big.Int).SetBytes(padTo32(payload[transcriptOffset : transcriptOffset+32])).Int64())
	transcriptStart := transcriptOffset + 32
	if transcriptStart+transcriptLen > len(payload) {
		return nil, nil, nil, fmt.Errorf("transcript bytes out of range")
	}
	transcript := payload[transcriptStart : transcriptStart+transcriptLen]

	const N = ccommon.MaxN
	// total = 8N words = 256N bytes
	totalBytes := 8 * N * 32
	if len(transcript) < totalBytes {
		return nil, nil, nil, fmt.Errorf("transcript too short: %d", len(transcript))
	}

	// commitmentPoints occupy the first 2N*32 bytes (skipped here).
	// Section offsets (in bytes):
	//   recipientIndexes: 2N*32
	//   recipientPubKeys: 3N*32
	//   ephemerals:       5N*32
	//   maskedShares:     7N*32
	ridxOffset := 2 * N * 32
	ridxs := make([]uint16, N)
	for i := range ridxs {
		word := new(big.Int).SetBytes(transcript[ridxOffset+i*32 : ridxOffset+i*32+32])
		ridxs[i] = uint16(word.Uint64())
	}
	ephOffset := 5 * N * 32
	ephs := make([][2]*big.Int, N)
	for i := range ephs {
		x := new(big.Int).SetBytes(transcript[ephOffset+i*64 : ephOffset+i*64+32])
		y := new(big.Int).SetBytes(transcript[ephOffset+i*64+32 : ephOffset+i*64+64])
		ephs[i] = [2]*big.Int{x, y}
	}
	maskedOffset := 7 * N * 32
	masked := make([]*big.Int, N)
	for i := range masked {
		masked[i] = new(big.Int).SetBytes(transcript[maskedOffset+i*32 : maskedOffset+i*32+32])
	}

	return ephs, masked, ridxs, nil
}

func padTo32(b []byte) []byte {
	if len(b) >= 32 {
		return b[len(b)-32:]
	}
	padded := make([]byte, 32)
	copy(padded[32-len(b):], b)
	return padded
}

func (n *Node) decNonce(roundID [12]byte, idx, ctIdx uint16) *big.Int {
	h, err := dkghash.HashFieldElements(
		roundScalar(roundID),
		new(big.Int).SetUint64(uint64(idx)),
		new(big.Int).SetUint64(uint64(ctIdx)),
	)
	if err != nil {
		h = big.NewInt(999)
	}
	h.Xor(h, n.bjjSecret)
	h.Mod(h, group.ScalarField())
	if h.Sign() == 0 {
		h.SetInt64(1)
	}
	return h
}

func (n *Node) readCiphertext(roundID [12]byte) (*CiphertextFile, error) {
	path := filepath.Join(n.sharedDir, fmt.Sprintf("ciphertext-%x.json", roundID))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ct CiphertextFile
	if err := json.Unmarshal(data, &ct); err != nil {
		return nil, err
	}
	return &ct, nil
}

// doCombineDecryption attempts to combine threshold-many partial decryptions
// into a single on-chain combined decryption proof. It is called on every
// Finalized-round poll cycle by any selected participant; the first node to
// succeed wins the race.
//
// Prerequisites:
//   - A ciphertext file (including C2X, C2Y) must be present in sharedDir.
//   - partialDecryptionCount >= threshold accepted partial decryptions on-chain.
func (n *Node) doCombineDecryption(
	ctx context.Context,
	roundID [12]byte,
	round web3.RoundView,
	selected []common.Address,
	callOpts *bind.CallOpts,
) error {
	if n.combined[roundID] {
		return nil
	}
	// Check whether already combined on-chain.
	rec, err := n.contracts.GetCombinedDecryption(ctx, roundID, 1)
	if err == nil && rec.Completed {
		n.combined[roundID] = true
		return nil
	}

	threshold := round.Policy.Threshold
	if round.PartialDecryptionCount < threshold {
		return nil // not enough partial decryptions yet
	}

	ct, err := n.readCiphertext(roundID)
	if err != nil {
		log.Debugw("combine: waiting for ciphertext file",
			"round", roundHex(roundID), "sharedDir", n.sharedDir)
		return nil
	}
	if ct.C2X == "" || ct.C2Y == "" {
		log.Warnw("combine: ciphertext file present but missing C2 coordinates — cannot combine; "+
			"resubmit ciphertext with c2x/c2y fields",
			"round", roundHex(roundID))
		return nil
	}

	// Gather accepted partial decryptions.
	var partialIndexes []uint16
	var partialDeltas []nodetypes.CurvePoint
	for _, addr := range selected {
		pd, err := n.manager.GetPartialDecryption(callOpts, roundID, addr, 1)
		if err != nil || !pd.Accepted {
			continue
		}
		partialIndexes = append(partialIndexes, pd.ParticipantIndex)
		partialDeltas = append(partialDeltas, nodetypes.CurvePoint{X: pd.Delta.X, Y: pd.Delta.Y})
		if len(partialIndexes) >= int(threshold) {
			break
		}
	}
	if len(partialIndexes) < int(threshold) {
		return nil // not enough accepted partial decryptions yet
	}

	// Parse ciphertext points.
	c1X, ok1 := new(big.Int).SetString(ct.C1X, 10)
	c1Y, ok2 := new(big.Int).SetString(ct.C1Y, 10)
	c2X, ok3 := new(big.Int).SetString(ct.C2X, 10)
	c2Y, ok4 := new(big.Int).SetString(ct.C2Y, 10)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return fmt.Errorf("combine: parse ciphertext coordinates")
	}

	// Interpolate partial decryptions to recover k*PubKey.
	indexes := ccommon.Uint16sToBigInts(partialIndexes)
	combinedEncoded, err := ccommon.InterpolatePointsAtZeroNative(indexes, partialDeltas)
	if err != nil {
		return fmt.Errorf("combine: interpolate partial decryptions: %w", err)
	}
	combined, err := group.Decode(combinedEncoded)
	if err != nil {
		return fmt.Errorf("combine: decode combined point: %w", err)
	}

	// mG = C2 - combined.
	c2, err := group.Decode(nodetypes.CurvePoint{X: c2X, Y: c2Y})
	if err != nil {
		return fmt.Errorf("combine: decode C2: %w", err)
	}
	negCombined := group.NewPoint()
	negCombined.Neg(combined)
	mG := group.NewPoint()
	mG.Add(c2, negCombined)

	// Brute-force discrete log to recover the plaintext scalar.
	plaintext, err := bruteForceLog(mG)
	if err != nil {
		// DLOG failed — plaintext is ≥ 2²⁰. Retrying will always produce the
		// same failure, so mark the round terminal to avoid burning CPU every tick.
		n.combined[roundID] = true
		n.terminal[roundID] = true
		return fmt.Errorf("combine: dlog failed (plaintext must be < 2²⁰ ≈ 1M): %w", err)
	}

	// Build the ZK witness.
	assignment := decryptcombine.Assignment{
		RoundHash:          roundScalar(roundID),
		Threshold:          threshold,
		CiphertextC1:       nodetypes.CurvePoint{X: c1X, Y: c1Y},
		CiphertextC2:       nodetypes.CurvePoint{X: c2X, Y: c2Y},
		ParticipantIndexes: partialIndexes,
		PartialDecryptions: partialDeltas,
		Plaintext:          plaintext,
	}
	witness, pi, err := decryptcombine.BuildWitness(assignment)
	if err != nil {
		return fmt.Errorf("combine: build witness: %w", err)
	}
	runtime, err := decryptcombine.Artifacts.LoadOrSetupForCircuit(ctx, &decryptcombine.DecryptCombineCircuit{})
	if err != nil {
		return fmt.Errorf("combine: load circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return fmt.Errorf("combine: prove: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return fmt.Errorf("combine: marshal proof: %w", err)
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return fmt.Errorf("combine: encode input: %w", err)
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return fmt.Errorf("combine: encode transcript: %w", err)
	}

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("combine: tx opts: %w", err)
	}
	tx, err := n.manager.CombineDecryption(
		auth, roundID, 1,
		common.BigToHash(pi.CombineHash),
		common.BigToHash(pi.PlaintextHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		if strings.Contains(err.Error(), "AlreadyCombined") || isPermanentRevert(err) {
			log.Warnw("combine decryption tx permanently rejected — will not retry",
				"round", roundHex(roundID), "err", err)
			n.combined[roundID] = true
			if strings.Contains(err.Error(), "AlreadyCombined") {
				return nil
			}
		}
		return fmt.Errorf("combine: submit tx: %w", err)
	}
	if err := n.txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		if isPermanentRevert(err) {
			log.Warnw("combine decryption tx reverted on-chain — will not retry",
				"round", roundHex(roundID), "err", err)
			n.combined[roundID] = true
		}
		return fmt.Errorf("combine: wait tx: %w", err)
	}
	n.combined[roundID] = true
	log.Infow("decryption combined", "round", roundHex(roundID), "plaintext", plaintext.String())
	return nil
}

// bruteForceLog recovers scalar m from m*G via brute-force DLOG.
// Only feasible for small plaintexts (m < 2^20 ≈ 1 million).
func bruteForceLog(target ecc.Point) (*big.Int, error) {
	gen := group.Generator()
	candidate := group.NewPoint()
	candidate.SetZero()
	for i := int64(0); i < 1<<20; i++ {
		if candidate.Equal(target) {
			return big.NewInt(i), nil
		}
		candidate.Add(candidate, gen)
	}
	return nil, fmt.Errorf("message out of brute-force range (>= 2^20)")
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

func makeRoundID(prefix uint32, nonce uint64) [12]byte {
	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], prefix)
	binary.BigEndian.PutUint64(id[4:], nonce)
	return id
}

func roundHex(id [12]byte) string { return fmt.Sprintf("%x", id) }

// isExpectedClaimRevert returns true if a claimSlot revert is "benign" — i.e.
// the node should silently accept it (not eligible, slot already gone, seed not
// yet available). The node will retry on the next poll for the SeedNotReady case
// since `signaled` only flips on definitively-final reverts.
// isPermanentRevert returns true when the error indicates the EVM rejected the
// transaction in a way that retrying will never succeed. Transient errors (RPC
// timeouts, network issues) do NOT match, so the node retries those naturally.
//
// We match the exact phrase "execution reverted" (the standard Ethereum error
// returned by both eth_call simulation and mined-but-reverted receipts).
// We intentionally do NOT match plain "reverted" to avoid false-positives from
// RPC provider error messages that happen to contain that word.
func isPermanentRevert(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "execution reverted")
}

func isExpectedClaimRevert(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	// Definitive: don't retry. NotEligible & SlotsFull & InvalidPhase mean the
	// committee is decided without us. AlreadyClaimed means we already won.
	// SeedExpired means the seed beacon data is gone — round is unrecoverable.
	for _, tok := range []string{"NotEligible", "SlotsFull", "AlreadyClaimed", "InvalidPhase", "SeedExpired"} {
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
	parsed, parseErr := gtypes.DKGManagerMetaData.GetAbi()
	if parseErr != nil {
		return err.Error()
	}
	if abiErr, lookupErr := parsed.ErrorByID(sel); lookupErr == nil {
		return fmt.Sprintf("execution reverted: %s", abiErr.Name)
	}
	return err.Error()
}
