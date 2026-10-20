package txmanager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/vocdoni/davinci-dkg/config"
	"github.com/vocdoni/davinci-dkg/log"
)

const (
	defaultMaxPendingTime     = 5 * time.Minute
	defaultMaxRetries         = 10
	defaultFeeIncreasePercent = 50
	defaultMonitorInterval    = 15 * time.Second
	// defaultGasMultiplier is the headroom applied to bind's gas estimate.
	// The estimate is exact for the state it was simulated against; a
	// transaction that lands behind others touching the same storage in
	// one block (three claimSlot calls, each costlier than the last) runs
	// out of gas without it.
	defaultGasMultiplier = 1.2
	// rpcCallTimeout bounds every RPC issued by the background monitor.
	rpcCallTimeout = 15 * time.Second
)

// Config holds tunable parameters for the Manager.
type Config struct {
	// MaxPendingTime is how long a transaction may be pending before it is
	// considered stuck and retried with a higher fee.
	MaxPendingTime time.Duration
	// MaxRetries is the maximum number of retry attempts for a stuck tx.
	MaxRetries int
	// FeeIncreasePercent is how much the gas price is bumped on each retry (%).
	FeeIncreasePercent int
	// MonitorInterval controls how often the background goroutine checks for
	// stuck transactions.
	MonitorInterval time.Duration
	// GasMultiplier inflates estimated gas limits (≥ 1). Explicit
	// TransactOpts.GasLimit values are never touched.
	GasMultiplier float64
	// Chain is the per-chain fee profile (fallback tip, fee cap).
	Chain config.Chain
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() Config {
	return Config{
		MaxPendingTime:     defaultMaxPendingTime,
		MaxRetries:         defaultMaxRetries,
		FeeIncreasePercent: defaultFeeIncreasePercent,
		MonitorInterval:    defaultMonitorInterval,
		GasMultiplier:      defaultGasMultiplier,
	}
}

// pendingTx tracks a nonce handed out by the manager until the chain has
// consumed it, and then until the transaction is deep (see observe).
// `signed` is nil between allocation and signing.
type pendingTx struct {
	hash        common.Hash
	nonce       uint64
	signed      *gethtypes.Transaction
	submittedAt time.Time
	retries     int
	capLogged   bool   // the "fee cap reached" warning went out for this entry
	mined       bool   // the chain's mined nonce count covers it
	minedAt     uint64 // head of the observation that first saw it mined
}

// confirmDepth is how many blocks a mined transaction stays tracked: a reorg
// that drops it within that depth is answered by rebroadcasting it, and
// only past it does its nonce raise the floor no stale view can reopen.
const confirmDepth = 64

// Manager handles nonce management, EIP-1559 gas estimation, and basic
// stuck-transaction recovery. It is safe for concurrent use.
//
// Nonces are allocated from a local counter, under the lock, at signing
// time: bind calls the Signer only after gas estimation succeeded, so a
// call that reverts during estimation never burns a nonce. The counter is
// initialised from (and never falls behind) the pending nonce bind fetched
// from the chain, and is reconciled against the confirmed nonce whenever a
// receipt is observed.
type Manager struct {
	clientFn func() *ethclient.Client
	key      *ecdsa.PrivateKey
	chainID  *big.Int
	from     common.Address
	config   Config
	// noter, when set, is told the JSON-RPC method of every request the
	// manager issues on the raw client, so the caller's per-method request
	// counters (web3.RPCPool.Note) stay complete.
	noter func(method string)
	// receipts, when set, serves the receipt polls of WaitTxByHash instead
	// of the raw client.
	receipts ReceiptReader

	mu        sync.Mutex
	nextNonce uint64
	// floor is one above the highest nonce whose transaction is deep: no
	// nonce below it is handed out again, whatever any endpoint reports.
	// minedCount is the mined count of the newest accepted observation,
	// which covers nonces this process never tracked (an account restarted
	// at nonce 5): allocation never goes below it, and only an accepted
	// reorg observation lowers it, never below floor. head is the newest
	// head an observation came from; older ones are stale and ignored.
	floor, minedCount, head uint64
	pending                 map[uint64]*pendingTx
	totalGasSpent           *big.Int // accumulated gas cost in wei across confirmed txs

	monitorCancel context.CancelFunc
}

// New creates a new Manager. clientFn is called each time an ethclient is needed,
// allowing the caller to supply a pool's Current() method for failover support.
func New(clientFn func() *ethclient.Client, chainID uint64, privateKey string) (*Manager, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}
	publicKey, ok := key.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid public key type")
	}
	cfg := DefaultConfig()
	cfg.Chain = config.ChainByID(chainID)
	return &Manager{
		clientFn:      clientFn,
		key:           key,
		chainID:       new(big.Int).SetUint64(chainID),
		from:          crypto.PubkeyToAddress(*publicKey),
		config:        cfg,
		pending:       make(map[uint64]*pendingTx),
		totalGasSpent: new(big.Int),
	}, nil
}

// Address returns the sender address controlled by this manager.
func (m *Manager) Address() common.Address {
	return m.from
}

// SetRequestNoter registers a callback told the JSON-RPC method of every
// request the manager sends (receipt polls, nonce and fee reads, balance
// reads, broadcasts), for request accounting. nil disables it.
func (m *Manager) SetRequestNoter(noter func(method string)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.noter = noter
}

// ReceiptReader fetches transaction receipts. *web3.PooledBackend is one: it
// moves to the next endpoint within the call when one refuses receipts (a
// free tier answering HTTP 403) or fails, where the raw client would poll the
// refusing endpoint until the wait times out.
type ReceiptReader interface {
	TransactionReceipt(ctx context.Context, hash common.Hash) (*gethtypes.Receipt, error)
}

// SetReceiptReader makes WaitTxByHash poll receipts through r, which counts
// its own requests, instead of the raw client.
func (m *Manager) SetReceiptReader(r ReceiptReader) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.receipts = r
}

// receipt reads one receipt through the receipt reader, or the raw client
// when none is set.
func (m *Manager) receipt(ctx context.Context, hash common.Hash) (*gethtypes.Receipt, error) {
	m.mu.Lock()
	r := m.receipts
	m.mu.Unlock()
	if r != nil {
		return r.TransactionReceipt(ctx, hash)
	}
	m.note("eth_getTransactionReceipt")
	return m.clientFn().TransactionReceipt(ctx, hash)
}

// note reports one request to the registered noter, if any.
func (m *Manager) note(method string) {
	m.mu.Lock()
	noter := m.noter
	m.mu.Unlock()
	if noter != nil {
		noter(method)
	}
}

// SetGasMultiplier sets the headroom applied to estimated gas limits.
// Values below 1 are ignored: headroom never shrinks an estimate.
func (m *Manager) SetGasMultiplier(x float64) {
	if x < 1 {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.config.GasMultiplier = x
}

// gasMultiplier reads the configured headroom under the lock.
func (m *Manager) gasMultiplier() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.config.GasMultiplier
}

// Start launches a background goroutine that monitors pending transactions
// and retries any that appear stuck (pending longer than MaxPendingTime).
func (m *Manager) Start(ctx context.Context) {
	monCtx, cancel := context.WithCancel(ctx)
	m.monitorCancel = cancel
	go func() {
		ticker := time.NewTicker(m.config.MonitorInterval)
		defer ticker.Stop()
		for {
			select {
			case <-monCtx.Done():
				return
			case <-ticker.C:
				_ = m.retryStuck(monCtx)
			}
		}
	}()
}

// Stop halts the background monitoring goroutine.
func (m *Manager) Stop() {
	if m.monitorCancel != nil {
		m.monitorCancel()
	}
}

// NewTransactOpts builds an EIP-1559 TransactOpts whose Signer allocates the
// nonce from the manager's counter (see Manager), applies the gas headroom
// to estimated limits and records the signed transaction for monitoring.
// Fee caps are estimated from current network conditions.
func (m *Manager) NewTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(m.key, m.chainID)
	if err != nil {
		return nil, fmt.Errorf("create transact opts: %w", err)
	}
	auth.Context = ctx
	// Fee estimation is best effort: on failure the bound method falls back
	// to the node's suggested gas price, unless the chain has a fee cap,
	// which bind's own estimate would ignore.
	if tipCap, feeCap, err := m.suggestFees(ctx); err == nil {
		auth.GasTipCap = tipCap
		auth.GasFeeCap = feeCap
	} else if tipCap, feeCap, ok := m.fallbackFees(); ok {
		auth.GasTipCap = tipCap
		auth.GasFeeCap = feeCap
	}

	sign := auth.Signer
	auth.Signer = func(addr common.Address, tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
		// bind fetched the chain's pending nonce for tx; treat it as a
		// lower bound and allocate the real one locally.
		nonce := m.allocateNonce(tx.Nonce())
		gas := tx.Gas()
		if auth.GasLimit == 0 { // estimated, not chosen by the caller
			gas = uint64(float64(gas) * m.gasMultiplier())
		}
		signed, err := sign(addr, rebuildTx(tx, nonce, gas))
		if err != nil {
			m.mu.Lock()
			delete(m.pending, nonce)
			m.mu.Unlock()
			return nil, err
		}
		m.track(signed)
		return signed, nil
	}
	return auth, nil
}

// allocateNonce hands out the next nonce under the lock. chainPending is
// the chain's view of the sender's next nonce: the counter never falls
// behind it, and any nonce between it and the counter that is no longer in
// flight (a transaction the monitor gave up on) is re-used first so the
// account cannot stall on a gap. Mined nonces are never re-used, whatever a
// stale chainPending (an endpoint behind the chain) says: none goes below
// the accepted mined count and the deep floor, and those still tracked are
// skipped like any in flight.
func (m *Manager) allocateNonce(chainPending uint64) uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	floor := max(chainPending, m.minedCount, m.floor)
	if floor > m.nextNonce {
		m.nextNonce = floor
	}
	nonce := m.nextNonce
	for k := floor; k < m.nextNonce; k++ {
		if _, inflight := m.pending[k]; !inflight {
			nonce = k
			break
		}
	}
	if nonce == m.nextNonce {
		m.nextNonce++
	}
	m.pending[nonce] = &pendingTx{nonce: nonce, submittedAt: time.Now()}
	return nonce
}

// track records a signed transaction against its nonce.
func (m *Manager) track(tx *gethtypes.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	ptx, ok := m.pending[tx.Nonce()]
	if !ok {
		ptx = &pendingTx{nonce: tx.Nonce(), submittedAt: time.Now()}
		m.pending[tx.Nonce()] = ptx
	}
	ptx.hash = tx.Hash()
	ptx.signed = tx
}

// rebuildTx returns tx with its nonce and gas limit replaced (or tx itself
// when both already match).
func rebuildTx(tx *gethtypes.Transaction, nonce, gas uint64) *gethtypes.Transaction {
	if tx.Nonce() == nonce && tx.Gas() == gas {
		return tx
	}
	if tx.Type() == gethtypes.LegacyTxType {
		return gethtypes.NewTx(&gethtypes.LegacyTx{
			Nonce: nonce, GasPrice: tx.GasPrice(), Gas: gas, To: tx.To(), Value: tx.Value(), Data: tx.Data(),
		})
	}
	return gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID: tx.ChainId(), Nonce: nonce, GasTipCap: tx.GasTipCap(), GasFeeCap: tx.GasFeeCap(),
		Gas: gas, To: tx.To(), Value: tx.Value(), Data: tx.Data(), AccessList: tx.AccessList(),
	})
}

// RecordPending stores a submitted transaction for monitoring. Transactions
// signed through NewTransactOpts are already tracked; calling this for
// them is harmless. It remains for transactions signed elsewhere.
func (m *Manager) RecordPending(tx *gethtypes.Transaction) {
	m.track(tx)
}

// Pending reports whether the transaction signed at nonce is still in
// flight: tracked (the original or a fee-bumped replacement), not seen
// mined and not given up on after MaxRetries.
func (m *Manager) Pending(nonce uint64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	ptx, ok := m.pending[nonce]
	return ok && !ptx.mined
}

// WaitTxByHash blocks until the transaction is confirmed or the timeout expires.
// It returns an error if the transaction reverts or the context/timeout fires.
// Transient receipt errors are retried until the timeout.
func (m *Manager) WaitTxByHash(hash common.Hash, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var lastErr error
	for {
		select {
		case <-ctx.Done():
			if lastErr != nil {
				return fmt.Errorf("timeout waiting for transaction %s (last error: %w)", hash.Hex(), lastErr)
			}
			return fmt.Errorf("timeout waiting for transaction %s", hash.Hex())
		case <-ticker.C:
			receipt, err := m.receipt(ctx, hash)
			switch {
			case err == nil:
				// Always record gas cost — reverted txs still consume gas.
				m.recordGasSpent(receipt)
				m.pruneConfirmed(receipt.BlockNumber)
				if receipt.Status != gethtypes.ReceiptStatusSuccessful {
					return fmt.Errorf("transaction %s reverted (status %d)", hash.Hex(), receipt.Status)
				}
				return nil
			case errors.Is(err, ethereum.NotFound):
				// Still pending — keep polling.
			default:
				lastErr = err // transient RPC failure — keep polling
			}
		}
	}
}

// Balance returns the current ETH balance of the managed account.
func (m *Manager) Balance(ctx context.Context) (*big.Int, error) {
	m.note("eth_getBalance")
	return m.clientFn().BalanceAt(ctx, m.from, nil)
}

// TotalGasSpent returns the accumulated gas cost in wei for all confirmed
// transactions tracked by this manager since it was created.
func (m *Manager) TotalGasSpent() *big.Int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return new(big.Int).Set(m.totalGasSpent)
}

// recordGasSpent adds the gas cost of a confirmed receipt to the running total.
func (m *Manager) recordGasSpent(receipt *gethtypes.Receipt) {
	if receipt == nil || receipt.EffectiveGasPrice == nil {
		return
	}
	cost := new(big.Int).Mul(
		new(big.Int).SetUint64(receipt.GasUsed),
		receipt.EffectiveGasPrice,
	)
	m.mu.Lock()
	m.totalGasSpent.Add(m.totalGasSpent, cost)
	m.mu.Unlock()
}

// fallbackTip is the priority fee used when the node cannot suggest one.
func (m *Manager) fallbackTip() *big.Int {
	if t := m.config.Chain.FallbackTipCap; t != nil {
		return new(big.Int).Set(t)
	}
	return big.NewInt(1_000_000_000) // 1 gwei
}

// fallbackFees returns the fees to use when estimation failed outright;
// ok is false on chains without a fee cap, where bind estimates them.
func (m *Manager) fallbackFees() (tip, feeCap *big.Int, ok bool) {
	limit := m.config.Chain.MaxFeePerGas
	if limit == nil {
		return nil, nil, false
	}
	return minBig(m.fallbackTip(), limit), new(big.Int).Set(limit), true
}

// minReplacement is the lowest fee geth's txpool accepts for a same-nonce
// replacement: 10% above the original (its default PriceBump), and always
// above it.
func minReplacement(v *big.Int) *big.Int {
	r := new(big.Int).Mul(v, big.NewInt(110))
	r.Add(r, big.NewInt(99)).Div(r, big.NewInt(100)) // ceil
	return maxBig(r, new(big.Int).Add(v, big.NewInt(1)))
}

func maxBig(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return b
	}
	return a
}

func minBig(a, b *big.Int) *big.Int {
	if b.Cmp(a) < 0 {
		return new(big.Int).Set(b)
	}
	return new(big.Int).Set(a)
}

// clampFees applies the chain's fee cap, keeping tip <= feeCap.
func (m *Manager) clampFees(tip, feeCap *big.Int) (*big.Int, *big.Int) {
	if limit := m.config.Chain.MaxFeePerGas; limit != nil {
		feeCap = minBig(feeCap, limit)
	}
	return minBig(tip, feeCap), feeCap
}

// suggestFees returns (gasTipCap, gasFeeCap) for an EIP-1559 transaction.
func (m *Manager) suggestFees(ctx context.Context) (*big.Int, *big.Int, error) {
	m.note("eth_maxPriorityFeePerGas")
	tipCap, err := m.clientFn().SuggestGasTipCap(ctx)
	if err != nil {
		tipCap = m.fallbackTip()
	}
	m.note("eth_getBlockByNumber")
	header, err := m.clientFn().HeaderByNumber(ctx, nil)
	if err != nil || header.BaseFee == nil {
		// Legacy chain or no base fee available — just use a high gas price.
		tip, feeCap := m.clampFees(tipCap, new(big.Int).Add(tipCap, big.NewInt(1_000_000_000)))
		return tip, feeCap, nil
	}
	// feeCap = 2 * baseFee + tipCap (gives headroom for a few blocks of fee increase)
	feeCap := new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), tipCap)
	tip, feeCap := m.clampFees(tipCap, feeCap)
	return tip, feeCap, nil
}

// retryStuck is one monitor pass. It snapshots the pending set under the
// lock, does all RPC work without it, and applies the outcome per entry
// only if that entry has not changed meanwhile:
//
//   - nonces below the confirmed nonce are mined, and stay tracked until
//     they are deep (observe), so a reorg that drops one is rebroadcast;
//   - a transaction at or above the chain's pending nonce is unknown to the
//     RPC's mempool (lost send or pool failover) and is re-broadcast as-is;
//   - a transaction pending longer than MaxPendingTime is replaced by a
//     fee-bumped copy, up to MaxRetries attempts (refused broadcasts
//     included), after which it is dropped so its nonce can be re-used.
func (m *Manager) retryStuck(ctx context.Context) error {
	m.mu.Lock()
	tracked := len(m.pending)
	m.mu.Unlock()
	if tracked == 0 {
		return nil
	}

	client := m.clientFn()
	// The mined count is read at the head just read, so the pair is one
	// observation. Without a head it is read at latest for this pass only
	// and observes nothing (see observe).
	m.note("eth_blockNumber")
	head, headErr := withTimeout(ctx, func(c context.Context) (uint64, error) { return client.BlockNumber(c) })
	var at *big.Int
	if headErr == nil {
		at = new(big.Int).SetUint64(head)
	}
	m.note("eth_getTransactionCount")
	confirmed, err := withTimeout(ctx, func(c context.Context) (uint64, error) { return client.NonceAt(c, m.from, at) })
	if err != nil {
		return fmt.Errorf("confirmed nonce: %w", err)
	}
	m.note("eth_getTransactionCount")
	chainPending, err := withTimeout(ctx, func(c context.Context) (uint64, error) { return client.PendingNonceAt(c, m.from) })
	if err != nil {
		return fmt.Errorf("pending nonce: %w", err)
	}
	if headErr == nil {
		m.applyObservation(ctx, confirmed, head)
	}

	m.mu.Lock()
	snapshot := make([]pendingTx, 0, len(m.pending))
	for _, ptx := range m.pending {
		snapshot = append(snapshot, *ptx)
	}
	m.mu.Unlock()
	now := time.Now()
	for _, ptx := range snapshot {
		if ptx.mined || ptx.nonce < confirmed || ptx.signed == nil {
			continue // mined, or still being signed
		}
		age := now.Sub(ptx.submittedAt)
		switch {
		case ptx.retries >= m.config.MaxRetries:
			m.dropExhausted(ptx)
		case age >= m.config.MaxPendingTime:
			// Before the unknown-to-the-mempool case: OP has no public mempool,
			// and a load-balanced RPC can answer the pending nonce with the
			// confirmed one, so an old tx stuck under the base fee would
			// otherwise be rebroadcast as is forever and never bumped.
			bumped, err := m.bumpedCopy(ptx.signed)
			if err != nil {
				continue
			}
			atCap := bumped.Hash() == ptx.signed.Hash()
			if atCap && !ptx.capLogged {
				log.Warnw("tx fee cap reached; rebroadcasting without a bump",
					"nonce", ptx.nonce, "tx", ptx.hash.Hex(), "maxFeePerGas", ptx.signed.GasFeeCap())
			}
			// At the cap this is a plain rebroadcast that still counts as an
			// attempt and restarts the pending clock, so the usual MaxRetries
			// bound and MaxPendingTime cadence hold.
			m.resend(ctx, ptx, bumped, true)
			if atCap {
				// A rebroadcast of a tx the pool already holds errors
				// ("already known") and resend then leaves the clock alone;
				// restart it here so the cadence stays MaxPendingTime.
				m.mu.Lock()
				if cur, ok := m.pending[ptx.nonce]; ok && cur.hash == ptx.hash {
					cur.capLogged = true
					cur.submittedAt = time.Now()
				}
				m.mu.Unlock()
			}
		case ptx.nonce >= chainPending && age >= m.config.MonitorInterval:
			m.resend(ctx, ptx, ptx.signed, false)
		}
	}
	return nil
}

// dropExhausted gives up on a transaction a monitor pass snapshotted at
// MaxRetries, if the tracker still holds it in that state: one a receipt
// observation marked mined meanwhile stays tracked until it is deep.
func (m *Manager) dropExhausted(ptx pendingTx) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cur, ok := m.pending[ptx.nonce]
	if ok && cur.hash == ptx.hash && !cur.mined && cur.retries >= m.config.MaxRetries {
		delete(m.pending, ptx.nonce)
	}
}

// resend broadcasts tx and, if the pending entry is still the one we
// snapshotted, records the new hash. A bump restarts the pending clock; a
// plain re-broadcast does not, so a stuck transaction still gets bumped on
// schedule.
func (m *Manager) resend(ctx context.Context, ptx pendingTx, tx *gethtypes.Transaction, bumped bool) {
	m.note("eth_sendRawTransaction")
	_, err := withTimeout(ctx, func(c context.Context) (struct{}, error) {
		return struct{}{}, m.clientFn().SendTransaction(c, tx)
	})
	m.mu.Lock()
	defer m.mu.Unlock()
	cur, ok := m.pending[ptx.nonce]
	if !ok || cur.hash != ptx.hash {
		return
	}
	// A bump counts as an attempt (MaxRetries bounds the fee escalation), and
	// so does a refused broadcast: a transaction the mempool rejects for good
	// (insufficient funds, underpriced) would otherwise be re-sent on every
	// pass forever, and the pending set — fed a fresh nonce by every caller
	// retry — would grow without bound. A plain rebroadcast the network takes
	// or already holds does not: on OP a load-balanced RPC can make a live tx
	// look unknown on every pass, and counting those would drop it, freeing
	// its nonce, long before it is old enough to bump.
	if bumped || (err != nil && !alreadyKnown(err)) {
		cur.retries++
	}
	if err != nil {
		return // "already known" and friends: the original may still confirm
	}
	cur.hash = tx.Hash()
	cur.signed = tx
	if bumped {
		cur.submittedAt = time.Now()
	}
}

// alreadyKnown reports whether a send failed only because the pool already
// holds the transaction (geth "already known", other clients' variants).
func alreadyKnown(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "already known") || strings.Contains(msg, "known transaction") ||
		strings.Contains(msg, "alreadyknown")
}

// withTimeout runs one RPC call under rpcCallTimeout.
func withTimeout[T any](ctx context.Context, call func(context.Context) (T, error)) (T, error) {
	cctx, cancel := context.WithTimeout(ctx, rpcCallTimeout)
	defer cancel()
	return call(cctx)
}

// bumpedCopy re-signs tx with fee caps increased by FeeIncreasePercent,
// clamped to the chain's fee cap. When the clamp leaves the tip or the fee
// cap less than 10% above the original, geth would reject a replacement as
// underpriced, so tx itself is returned to be rebroadcast as is.
func (m *Manager) bumpedCopy(tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
	pct := big.NewInt(int64(100 + m.config.FeeIncreasePercent))
	bump := func(v *big.Int) *big.Int {
		out := new(big.Int).Mul(v, pct)
		return out.Div(out, big.NewInt(100))
	}
	tip, feeCap := m.clampFees(bump(tx.GasTipCap()), bump(tx.GasFeeCap()))
	if m.config.Chain.MaxFeePerGas != nil &&
		(tip.Cmp(minReplacement(tx.GasTipCap())) < 0 || feeCap.Cmp(minReplacement(tx.GasFeeCap())) < 0) {
		return tx, nil
	}
	raw := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID:    m.chainID,
		Nonce:      tx.Nonce(),
		GasTipCap:  tip,
		GasFeeCap:  feeCap,
		Gas:        tx.Gas(),
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	})
	return gethtypes.SignTx(raw, gethtypes.LatestSignerForChainID(m.chainID), m.key)
}

// observe applies one chain observation: the sender's mined nonce count
// read at block head. Only an anchored count is an observation: a zero head
// (unknown) changes nothing, and one older than a head already seen comes
// from an endpoint behind the chain and is ignored, as is a count below the
// deep floor. An accepted observation sets the mined count allocation never
// goes below, lower or higher. Tracked transactions the count covers are
// marked mined; one marked mined that it no longer covers was dropped by a
// reorg: it is in flight again, the counter comes back down to just above
// what is still tracked, and it is returned for the caller to rebroadcast.
// A transaction still mined confirmDepth blocks after it was first seen
// mined is deep: it is dropped from the tracker and raises the floor.
func (m *Manager) observe(mined, head uint64) []*gethtypes.Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()
	if head == 0 || head < m.head || mined < m.floor {
		return nil
	}
	m.head, m.minedCount = head, mined
	var reorged []*gethtypes.Transaction
	for nonce, ptx := range m.pending {
		switch {
		case nonce < mined && !ptx.mined:
			ptx.mined, ptx.minedAt = true, head
		case nonce >= mined && ptx.mined:
			ptx.mined, ptx.minedAt = false, 0
			ptx.submittedAt = time.Now()
			if ptx.signed != nil {
				reorged = append(reorged, ptx.signed)
			}
		}
	}
	for nonce, ptx := range m.pending {
		if ptx.mined && head >= ptx.minedAt+confirmDepth {
			delete(m.pending, nonce)
			m.floor = max(m.floor, nonce+1)
		}
	}
	if len(reorged) > 0 {
		next := max(mined, m.floor)
		for nonce := range m.pending {
			next = max(next, nonce+1)
		}
		m.nextNonce = next
	}
	if mined > m.nextNonce {
		m.nextNonce = mined
	}
	return reorged
}

// applyObservation observes the mined nonce count at head and rebroadcasts
// what a reorg dropped.
func (m *Manager) applyObservation(ctx context.Context, mined, head uint64) {
	for _, tx := range m.observe(mined, head) {
		log.Warnw("tx manager: a reorg dropped a mined transaction, rebroadcasting",
			"nonce", tx.Nonce(), "tx", tx.Hash().Hex())
		m.resend(ctx, pendingTx{nonce: tx.Nonce(), hash: tx.Hash()}, tx, false)
	}
}

// pruneConfirmed observes the sender's nonce at blockNumber, the block a
// receipt was seen in.
func (m *Manager) pruneConfirmed(blockNumber *big.Int) {
	if blockNumber == nil || !blockNumber.IsUint64() {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), rpcCallTimeout)
	defer cancel()
	m.note("eth_getTransactionCount")
	mined, err := m.clientFn().NonceAt(ctx, m.from, blockNumber)
	if err != nil {
		return
	}
	m.applyObservation(ctx, mined, blockNumber.Uint64())
}
