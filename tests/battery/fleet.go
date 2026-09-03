// Package battery is the load / concurrency / adversarial test battery for
// davinci-dkg. Every test talks to an already-running fleet (Anvil + real
// davinci-dkg-node daemons) through the harness' external mode; nothing here
// starts Docker. All tests are gated on DAVINCI_DKG_BATTERY=1 and skip
// otherwise, so `make test` is unaffected.
package battery

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/vocdoni/davinci-dkg/config"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

const (
	envEnabled       = "DAVINCI_DKG_BATTERY"
	defaultRPCURL    = "http://127.0.0.1:8545"
	defaultAddresses = "/tmp/testnet-addresses.env"

	statusCommitteeSelection uint8 = 1
	statusKeyAssembly        uint8 = 2
	statusLive               uint8 = 3
	statusAborted            uint8 = 4

	// logRangeBlocks bounds every eth_getLogs call.
	logRangeBlocks = 10_000
	pollInterval   = 500 * time.Millisecond
)

// Fleet is the battery's handle on the running testnet: typed contract
// handles, a raw JSON-RPC client for Anvil cheat codes, and the deploy-time
// immutables every scenario needs to size its waits.
type Fleet struct {
	Services *helpers.TestServices
	rpc      *rpc.Client

	Prefix           uint32
	EpochDuration    uint64
	MinThreshold     uint16
	MinCommitteeSize uint16
	MaxAlphaBps      uint16
	MaxN             uint16

	txTimeout time.Duration
}

// actor is a fresh, funded EOA with its own nonce-managing signer.
type actor struct {
	*helpers.TestActor
	Label string
}

func enabled() bool { return os.Getenv(envEnabled) == "1" }

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}

func envUint64(key string, def uint64) uint64 {
	return uint64(envInt(key, int(def)))
}

func envDuration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}

// connect waits for the RPC and the deployer's addresses file, then builds
// the typed handles. The harness' TxManager is bound to a fresh random key
// (never an Anvil default account: those belong to the deployer and the
// nodes, whose signers allocate nonces locally).
func connect(ctx context.Context) (*Fleet, error) {
	rpcURL := envOr(helpers.TestRPCURLEnvVar, defaultRPCURL)
	addrPath := envOr(helpers.TestAddressesEnvVar, defaultAddresses)

	rpcClient, err := waitRPC(ctx, rpcURL)
	if err != nil {
		return nil, err
	}
	addresses, err := waitAddresses(ctx, addrPath)
	if err != nil {
		return nil, err
	}
	contracts, err := web3.New([]string{rpcURL}, addresses)
	if err != nil {
		return nil, fmt.Errorf("web3: %w", err)
	}
	treasuryKey, err := freshKey()
	if err != nil {
		return nil, err
	}
	txm, err := txmanager.New(contracts.Pool().Current, contracts.ChainID, treasuryKey)
	if err != nil {
		return nil, fmt.Errorf("txmanager: %w", err)
	}
	registry, err := golangtypes.NewDKGRegistry(contracts.Addresses.Registry, contracts.Client())
	if err != nil {
		return nil, err
	}
	manager, err := golangtypes.NewDKGManager(contracts.Addresses.Manager, contracts.Client())
	if err != nil {
		return nil, err
	}
	appManager, err := golangtypes.NewDKGAppManager(contracts.Addresses.AppManager, contracts.Client())
	if err != nil {
		return nil, err
	}
	f := &Fleet{
		Services: &helpers.TestServices{
			RPCURL:     rpcURL,
			Addresses:  contracts.Addresses,
			Contracts:  contracts,
			Registry:   registry,
			Manager:    manager,
			AppManager: appManager,
			TxManager:  txm,
		},
		rpc:       rpcClient,
		txTimeout: envDuration("BATTERY_TX_TIMEOUT", 3*time.Minute),
	}
	if err := f.fund(ctx, txm.Address()); err != nil {
		return nil, err
	}
	return f, f.readImmutables(ctx)
}

func (f *Fleet) close() {
	_ = f.Services.Contracts.Close()
	f.rpc.Close()
}

func waitRPC(ctx context.Context, rpcURL string) (*rpc.Client, error) {
	deadline := time.Now().Add(envDuration("BATTERY_CONNECT_TIMEOUT", 10*time.Minute))
	for {
		client, err := rpc.DialContext(ctx, rpcURL)
		if err == nil {
			var chainID hexutil.Big
			if err = client.CallContext(ctx, &chainID, "eth_chainId"); err == nil {
				return client, nil
			}
			client.Close()
		}
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("rpc %s not reachable: %w", rpcURL, err)
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func waitAddresses(ctx context.Context, path string) (types.ContractAddresses, error) {
	deadline := time.Now().Add(envDuration("BATTERY_CONNECT_TIMEOUT", 10*time.Minute))
	for {
		addresses, err := config.LoadContractAddressesFile(path)
		if err == nil {
			return addresses, nil
		}
		if time.Now().After(deadline) {
			return types.ContractAddresses{}, fmt.Errorf("addresses file %s: %w", path, err)
		}
		select {
		case <-ctx.Done():
			return types.ContractAddresses{}, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func (f *Fleet) readImmutables(ctx context.Context) error {
	opts := f.callOpts(ctx)
	prefix, err := f.Services.Manager.EPOCHPREFIX(opts)
	if err != nil {
		return fmt.Errorf("epoch prefix: %w", err)
	}
	duration, err := f.Services.Manager.EPOCHDURATIONBLOCKS(opts)
	if err != nil {
		return fmt.Errorf("epoch duration: %w", err)
	}
	minT, err := f.Services.Manager.MINTHRESHOLD(opts)
	if err != nil {
		return fmt.Errorf("min threshold: %w", err)
	}
	minN, err := f.Services.Manager.MINCOMMITTEESIZE(opts)
	if err != nil {
		return fmt.Errorf("min committee size: %w", err)
	}
	maxAlpha, err := f.Services.Manager.MAXLOTTERYALPHABPS(opts)
	if err != nil {
		return fmt.Errorf("max lottery alpha: %w", err)
	}
	f.Prefix = prefix
	f.EpochDuration = duration.Uint64()
	f.MinThreshold = minT
	f.MinCommitteeSize = minN
	f.MaxAlphaBps = maxAlpha
	f.MaxN = 32
	return nil
}

// ─── actors ──────────────────────────────────────────────────────────────

func freshKey() (string, error) {
	key, err := ethcrypto.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}
	return "0x" + hex.EncodeToString(ethcrypto.FromECDSA(key)), nil
}

// newActor creates a random key, funds it through anvil_setBalance and wraps
// it in a harness TestActor.
func (f *Fleet) newActor(ctx context.Context, label string) (*actor, error) {
	key, err := freshKey()
	if err != nil {
		return nil, err
	}
	base, err := f.Services.ActorFromPrivateKey(key)
	if err != nil {
		return nil, err
	}
	if err := f.fund(ctx, base.Address()); err != nil {
		return nil, err
	}
	return &actor{TestActor: base, Label: label}, nil
}

// newActors is newActor in bulk.
func (f *Fleet) newActors(ctx context.Context, prefix string, count int) ([]*actor, error) {
	out := make([]*actor, 0, count)
	for i := range count {
		a, err := f.newActor(ctx, fmt.Sprintf("%s-%d", prefix, i))
		if err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, nil
}

func (f *Fleet) fund(ctx context.Context, addr common.Address) error {
	balance := new(big.Int).Mul(big.NewInt(10_000), big.NewInt(1e18))
	if err := f.rpc.CallContext(ctx, nil, "anvil_setBalance", addr, hexutil.EncodeBig(balance)); err != nil {
		return fmt.Errorf("anvil_setBalance %s: %w", addr.Hex(), err)
	}
	return nil
}

// ─── transactions ────────────────────────────────────────────────────────

// txOutcome is what every mined transaction reports: hash, inclusion block,
// gas, the head observed when it was sent and the wall-clock wait.
type txOutcome struct {
	Hash     common.Hash
	Block    uint64
	Gas      uint64
	SentHead uint64
	Seconds  float64
}

func (o txOutcome) latencyBlocks() int64 { return int64(o.Block) - int64(o.SentHead) }

// result turns a mined tx into a report row.
func (o txOutcome) result(step, kind, notes string) Result {
	return Result{
		Step: step, Kind: kind, Tx: o.Hash.Hex(), Gas: o.Gas, Block: o.Block,
		LatencyBlocks: o.latencyBlocks(), LatencySeconds: o.Seconds, Pass: true, Notes: notes,
	}
}

// revertError is a transaction rejected by the EVM, either at gas
// estimation (the common case: bind simulates before signing, so no nonce
// is burnt) or as a mined-but-reverted receipt.
type revertError struct {
	Name string
	Err  error
}

func (e *revertError) Error() string { return "revert " + e.Name + ": " + e.Err.Error() }
func (e *revertError) Unwrap() error { return e.Err }

// send signs and submits one transaction through the actor's nonce manager,
// waits for the receipt and reports gas and latency. A revert at estimation
// or on chain is returned as *revertError.
func (f *Fleet) send(
	ctx context.Context,
	a *actor,
	build func(*bind.TransactOpts) (*ethtypes.Transaction, error),
) (txOutcome, error) {
	sentHead, err := f.head(ctx)
	if err != nil {
		return txOutcome{}, err
	}
	start := time.Now()
	auth, err := a.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return txOutcome{}, err
	}
	tx, err := build(auth)
	if err != nil {
		return txOutcome{SentHead: sentHead}, classify(err)
	}
	out := txOutcome{Hash: tx.Hash(), SentHead: sentHead}
	if err := a.TxManager.WaitTxByHash(tx.Hash(), f.txTimeout); err != nil {
		return out, classify(err)
	}
	receipt, err := f.Services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return out, fmt.Errorf("receipt %s: %w", tx.Hash().Hex(), err)
	}
	out.Block = receipt.BlockNumber.Uint64()
	out.Gas = receipt.GasUsed
	out.Seconds = time.Since(start).Seconds()
	return out, nil
}

// classify wraps EVM rejections in *revertError, leaving transport errors
// untouched.
func classify(err error) error {
	if err == nil {
		return nil
	}
	if name, ok := revertName(err); ok {
		return &revertError{Name: name, Err: err}
	}
	return err
}

// knownErrorSelectors covers the errors that are not in the three contract
// ABIs: the generated verifier wrappers.
var knownErrorSelectors = func() map[[4]byte]string {
	m := map[[4]byte]string{}
	for _, sig := range []string{
		"ProofInvalid()", "InvalidProofEncoding()", "InvalidInputEncoding()", "PublicInputNotInField()",
		"AppManagerNotSet()", "AppManagerAlreadySet()",
	} {
		var sel [4]byte
		copy(sel[:], ethcrypto.Keccak256([]byte(sig))[:4])
		m[sel] = strings.TrimSuffix(sig, "()")
	}
	return m
}()

// revertName extracts a custom-error name from an EVM error. ok=false means
// the error is not a revert at all.
func revertName(err error) (string, bool) {
	data, hasData := ethclient.RevertErrorData(err)
	if hasData && len(data) >= 4 {
		var sel [4]byte
		copy(sel[:], data[:4])
		for _, md := range []*bind.MetaData{
			golangtypes.DKGManagerMetaData, golangtypes.DKGAppManagerMetaData, golangtypes.DKGRegistryMetaData,
		} {
			parsed, parseErr := md.GetAbi()
			if parseErr != nil {
				continue
			}
			if abiErr, lookupErr := parsed.ErrorByID(sel); lookupErr == nil {
				return abiErr.Name, true
			}
		}
		if name, ok := knownErrorSelectors[sel]; ok {
			return name, true
		}
		return "custom:0x" + hex.EncodeToString(data[:4]), true
	}
	msg := err.Error()
	switch {
	case strings.Contains(msg, "reverted (status"):
		return "mined-revert", true
	case strings.Contains(msg, "execution reverted"):
		return "execution-reverted", true
	}
	return "", false
}

// expectRevert records whether err is a revert with one of the wanted names.
// A mismatch is reported through t.Errorf so the scenario keeps going and the
// report stays complete.
func expectRevert(t *testing.T, step string, err error, want ...string) bool {
	t.Helper()
	var rev *revertError
	got := "<no revert>"
	if errors.As(err, &rev) {
		got = rev.Name
	} else if err != nil {
		got = "error: " + err.Error()
	}
	pass := false
	for _, w := range want {
		if got == w {
			pass = true
		}
	}
	record(t, Result{
		Step: step, Kind: "revert", Pass: pass,
		Notes: fmt.Sprintf("want %s, got %s", strings.Join(want, "|"), got),
	})
	if !pass {
		t.Errorf("%s: want revert %v, got %s", step, want, got)
	}
	return pass
}

// expectRevertOrInconclusive is expectRevert for checks that race the real
// nodes: `want` is the revert that proves the property, `inconclusive` are
// the reverts that mean the fleet moved on first (e.g. the committee filled
// before our claim reached the check). Those are recorded as passing rows
// with an explicit note instead of failing the scenario.
func expectRevertOrInconclusive(t *testing.T, step string, err error, want string, inconclusive ...string) {
	t.Helper()
	var rev *revertError
	if errors.As(err, &rev) {
		for _, name := range inconclusive {
			if rev.Name == name {
				record(t, Result{
					Step: step, Kind: "revert", Pass: true,
					Notes: fmt.Sprintf("inconclusive: got %s (fleet moved on before the %s check), want %s", name, want, want),
				})
				return
			}
		}
	}
	expectRevert(t, step, err, want)
}

// expectOK records a successful tx and fails the test (non-fatally) when the
// tx errored instead.
func expectOK(t *testing.T, step, kind string, out txOutcome, err error, notes string) bool {
	t.Helper()
	if err != nil {
		record(t, Result{Step: step, Kind: kind, Tx: out.Hash.Hex(), Pass: false, Notes: err.Error()})
		t.Errorf("%s: %v", step, err)
		return false
	}
	record(t, out.result(step, kind, notes))
	return true
}

// ─── chain reads ─────────────────────────────────────────────────────────

func (f *Fleet) callOpts(ctx context.Context) *bind.CallOpts { return &bind.CallOpts{Context: ctx} }

func (f *Fleet) head(ctx context.Context) (uint64, error) {
	return f.Services.Contracts.Client().BlockNumber(ctx)
}

func (f *Fleet) blockTime(ctx context.Context, number uint64) (time.Time, error) {
	header, err := f.Services.Contracts.Client().HeaderByNumber(ctx, new(big.Int).SetUint64(number))
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(header.Time), 0), nil
}

// blocksToSeconds measures the wall-clock span between two blocks from their
// timestamps.
func (f *Fleet) blocksToSeconds(ctx context.Context, from, to uint64) float64 {
	a, errA := f.blockTime(ctx, from)
	b, errB := f.blockTime(ctx, to)
	if errA != nil || errB != nil {
		return 0
	}
	return b.Sub(a).Seconds()
}

func (f *Fleet) epochNonce(ctx context.Context) (uint64, error) {
	return f.Services.Manager.EpochNonce(f.callOpts(ctx))
}

func (f *Fleet) epochID(nonce uint64) [12]byte { return web3.EpochID(f.Prefix, nonce) }

func (f *Fleet) epoch(ctx context.Context, id [12]byte) (web3.EpochView, error) {
	return f.Services.Contracts.GetEpoch(ctx, id)
}

// serviceEnd is the cadence boundary of an epoch: the block at which the
// next epoch may be created. The epoch itself stays Live on chain forever,
// but the nodes only participate in the newest epochs, so this is the useful
// notion of "end".
func (f *Fleet) serviceEnd(e web3.EpochView) uint64 { return e.StartBlock + f.EpochDuration }

func (f *Fleet) committee(ctx context.Context, id [12]byte) ([]common.Address, error) {
	return f.Services.Contracts.SelectedParticipants(ctx, id)
}

func (f *Fleet) collectiveKey(ctx context.Context, id [12]byte) (types.CurvePoint, error) {
	pk, err := f.Services.Manager.GetCollectivePublicKey(f.callOpts(ctx), id)
	if err != nil {
		return types.CurvePoint{}, err
	}
	return types.CurvePoint{X: pk.X, Y: pk.Y}, nil
}

func (f *Fleet) nextEpochStart(ctx context.Context) (uint64, error) {
	return f.Services.Manager.NextEpochStartBlock(f.callOpts(ctx))
}

// waitBlock blocks until head >= target.
func (f *Fleet) waitBlock(ctx context.Context, target uint64) (uint64, error) {
	for {
		head, err := f.head(ctx)
		if err == nil && head >= target {
			return head, nil
		}
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("waiting for block %d: %w", target, ctx.Err())
		case <-time.After(pollInterval):
		}
	}
}

// waitStatus polls until the epoch reaches the given status (or is aborted).
func (f *Fleet) waitStatus(ctx context.Context, id [12]byte, status uint8) (web3.EpochView, error) {
	for {
		e, err := f.epoch(ctx, id)
		if err == nil && e.Status == status {
			return e, nil
		}
		if err == nil && e.Status == statusAborted {
			return e, fmt.Errorf("epoch %x aborted while waiting for status %d", id, status)
		}
		select {
		case <-ctx.Done():
			return web3.EpochView{}, fmt.Errorf("waiting for epoch %x status %d: %w", id, status, ctx.Err())
		case <-time.After(pollInterval):
		}
	}
}

// waitNonceAbove blocks until epochNonce > nonce and returns the new value.
func (f *Fleet) waitNonceAbove(ctx context.Context, nonce uint64) (uint64, error) {
	for {
		cur, err := f.epochNonce(ctx)
		if err == nil && cur > nonce {
			return cur, nil
		}
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("waiting for epoch nonce > %d: %w", nonce, ctx.Err())
		case <-time.After(pollInterval):
		}
	}
}

// waitLiveEpoch returns the newest Live epoch that still has at least
// minBlocksLeft blocks before its cadence boundary, waiting for the next one
// otherwise. Progress is logged so a long wait is readable.
func (f *Fleet) waitLiveEpoch(ctx context.Context, t *testing.T, minBlocksLeft uint64) ([12]byte, web3.EpochView, error) {
	t.Helper()
	lastLog := time.Time{}
	for {
		id, e, ok, status := f.liveCandidate(ctx, minBlocksLeft)
		if ok {
			return id, e, nil
		}
		if time.Since(lastLog) > 20*time.Second {
			t.Logf("waitLiveEpoch: %s", status)
			lastLog = time.Now()
		}
		select {
		case <-ctx.Done():
			return [12]byte{}, web3.EpochView{}, fmt.Errorf("waiting for a live epoch: %w", ctx.Err())
		case <-time.After(2 * time.Second):
		}
	}
}

// liveCandidate inspects the newest epoch and reports whether it is usable.
func (f *Fleet) liveCandidate(ctx context.Context, minBlocksLeft uint64) ([12]byte, web3.EpochView, bool, string) {
	nonce, err := f.epochNonce(ctx)
	if err != nil || nonce == 0 {
		return [12]byte{}, web3.EpochView{}, false, fmt.Sprintf("no epoch yet (err=%v)", err)
	}
	head, err := f.head(ctx)
	if err != nil {
		return [12]byte{}, web3.EpochView{}, false, err.Error()
	}
	id := f.epochID(nonce)
	e, err := f.epoch(ctx, id)
	if err != nil {
		return id, e, false, err.Error()
	}
	left := int64(f.serviceEnd(e)) - int64(head)
	status := fmt.Sprintf("epoch nonce=%d status=%d head=%d serviceEnd=%d left=%d liveNotBefore=%d",
		nonce, e.Status, head, f.serviceEnd(e), left, e.Policy.LiveNotBeforeBlock)
	if e.Status == statusLive && left >= int64(minBlocksLeft) {
		return id, e, true, status
	}
	return id, e, false, status
}

// ─── event scans ─────────────────────────────────────────────────────────

type partialEvent struct {
	Index       uint16
	Participant common.Address
	Delta       types.CurvePoint
	Block       uint64
	Tx          common.Hash
}

// partials lists the PartialDecryptionSubmitted events for one ciphertext
// slot, in log order.
func (f *Fleet) partials(ctx context.Context, id [12]byte, aid [32]byte, idx uint16, from uint64) ([]partialEvent, error) {
	head, err := f.head(ctx)
	if err != nil {
		return nil, err
	}
	var out []partialEvent
	for start := from; start <= head; start += logRangeBlocks {
		end := min(start+logRangeBlocks-1, head)
		it, err := f.Services.Manager.FilterPartialDecryptionSubmitted(
			&bind.FilterOpts{Context: ctx, Start: start, End: &end}, [][12]byte{id}, [][32]byte{aid}, nil)
		if err != nil {
			return nil, fmt.Errorf("filter PartialDecryptionSubmitted: %w", err)
		}
		for it.Next() {
			ev := it.Event
			if ev.CiphertextIndex != idx {
				continue
			}
			out = append(out, partialEvent{
				Index:       ev.ParticipantIndex,
				Participant: ev.Participant,
				Delta:       types.CurvePoint{X: new(big.Int).Set(ev.DeltaX), Y: new(big.Int).Set(ev.DeltaY)},
				Block:       ev.Raw.BlockNumber,
				Tx:          ev.Raw.TxHash,
			})
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

type combineEvent struct {
	Block     uint64
	Tx        common.Hash
	Gas       uint64
	Plaintext *big.Int
	Sender    common.Address
}

// combined returns the DecryptionCombined event of a slot if it exists.
func (f *Fleet) combined(ctx context.Context, id [12]byte, aid [32]byte, idx uint16, from uint64) (*combineEvent, error) {
	head, err := f.head(ctx)
	if err != nil {
		return nil, err
	}
	for start := from; start <= head; start += logRangeBlocks {
		end := min(start+logRangeBlocks-1, head)
		it, err := f.Services.Manager.FilterDecryptionCombined(
			&bind.FilterOpts{Context: ctx, Start: start, End: &end}, [][12]byte{id}, [][32]byte{aid}, []uint16{idx})
		if err != nil {
			return nil, fmt.Errorf("filter DecryptionCombined: %w", err)
		}
		var found *combineEvent
		if it.Next() {
			ev := it.Event
			found = &combineEvent{Block: ev.Raw.BlockNumber, Tx: ev.Raw.TxHash, Plaintext: new(big.Int).Set(ev.Plaintext)}
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return nil, err
		}
		if found != nil {
			if receipt, rerr := f.Services.Contracts.Client().TransactionReceipt(ctx, found.Tx); rerr == nil {
				found.Gas = receipt.GasUsed
			}
			if tx, _, terr := f.Services.Contracts.Client().TransactionByHash(ctx, found.Tx); terr == nil {
				if sender, serr := ethtypes.Sender(ethtypes.LatestSignerForChainID(tx.ChainId()), tx); serr == nil {
					found.Sender = sender
				}
			}
			return found, nil
		}
	}
	return nil, nil
}

// waitCombine polls until the slot is combined or the chain passes
// deadlineBlock. ok=false with a nil error means the deadline passed.
func (f *Fleet) waitCombine(
	ctx context.Context, id [12]byte, aid [32]byte, idx uint16, from, deadlineBlock uint64,
) (*combineEvent, bool, error) {
	for {
		rec, err := f.Services.Contracts.GetCombinedDecryption(ctx, id, aid, idx)
		if err == nil && rec.Completed {
			ev, err := f.combined(ctx, id, aid, idx, from)
			if err != nil {
				return nil, false, err
			}
			if ev == nil {
				ev = &combineEvent{Plaintext: rec.Plaintext}
			}
			return ev, true, nil
		}
		head, herr := f.head(ctx)
		if herr == nil && head > deadlineBlock {
			return nil, false, nil
		}
		select {
		case <-ctx.Done():
			return nil, false, ctx.Err()
		case <-time.After(time.Second):
		}
	}
}

// scanFrom is the block every per-epoch event scan starts at.
func scanFrom(e web3.EpochView) uint64 {
	if e.SeedBlock > 0 {
		return e.SeedBlock - 1
	}
	return 0
}
