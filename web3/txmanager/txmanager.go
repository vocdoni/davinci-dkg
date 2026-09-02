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
// consumed it. `signed` is nil between allocation and signing.
type pendingTx struct {
	hash        common.Hash
	nonce       uint64
	signed      *gethtypes.Transaction
	submittedAt time.Time
	retries     int
}

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

	mu            sync.Mutex
	nextNonce     uint64
	pending       map[uint64]*pendingTx
	totalGasSpent *big.Int // accumulated gas cost in wei across confirmed txs

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
	return &Manager{
		clientFn:      clientFn,
		key:           key,
		chainID:       new(big.Int).SetUint64(chainID),
		from:          crypto.PubkeyToAddress(*publicKey),
		config:        DefaultConfig(),
		pending:       make(map[uint64]*pendingTx),
		totalGasSpent: new(big.Int),
	}, nil
}

// Address returns the sender address controlled by this manager.
func (m *Manager) Address() common.Address {
	return m.from
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
	// to the node's suggested gas price.
	if tipCap, feeCap, err := m.suggestFees(ctx); err == nil {
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
// account cannot stall on a gap.
func (m *Manager) allocateNonce(chainPending uint64) uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	if chainPending > m.nextNonce {
		m.nextNonce = chainPending
	}
	nonce := m.nextNonce
	for k := chainPending; k < m.nextNonce; k++ {
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
			receipt, err := m.clientFn().TransactionReceipt(ctx, hash)
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

// suggestFees returns (gasTipCap, gasFeeCap) for an EIP-1559 transaction.
func (m *Manager) suggestFees(ctx context.Context) (*big.Int, *big.Int, error) {
	tipCap, err := m.clientFn().SuggestGasTipCap(ctx)
	if err != nil {
		tipCap = big.NewInt(1_000_000_000) // 1 gwei fallback
	}
	header, err := m.clientFn().HeaderByNumber(ctx, nil)
	if err != nil || header.BaseFee == nil {
		// Legacy chain or no base fee available — just use a high gas price.
		return tipCap, new(big.Int).Add(tipCap, big.NewInt(1_000_000_000)), nil
	}
	// feeCap = 2 * baseFee + tipCap (gives headroom for a few blocks of fee increase)
	feeCap := new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), tipCap)
	return tipCap, feeCap, nil
}

// retryStuck is one monitor pass. It snapshots the pending set under the
// lock, does all RPC work without it, and applies the outcome per entry
// only if that entry has not changed meanwhile:
//
//   - nonces below the confirmed nonce are done (mined or replaced);
//   - a transaction at or above the chain's pending nonce is unknown to the
//     RPC's mempool (lost send or pool failover) and is re-broadcast as-is;
//   - a transaction pending longer than MaxPendingTime is replaced by a
//     fee-bumped copy, up to MaxRetries times, after which it is dropped so
//     its nonce can be re-used.
func (m *Manager) retryStuck(ctx context.Context) error {
	m.mu.Lock()
	snapshot := make([]pendingTx, 0, len(m.pending))
	for _, ptx := range m.pending {
		snapshot = append(snapshot, *ptx)
	}
	m.mu.Unlock()
	if len(snapshot) == 0 {
		return nil
	}

	client := m.clientFn()
	confirmed, err := withTimeout(ctx, func(c context.Context) (uint64, error) { return client.NonceAt(c, m.from, nil) })
	if err != nil {
		return fmt.Errorf("confirmed nonce: %w", err)
	}
	chainPending, err := withTimeout(ctx, func(c context.Context) (uint64, error) { return client.PendingNonceAt(c, m.from) })
	if err != nil {
		return fmt.Errorf("pending nonce: %w", err)
	}
	m.reconcile(confirmed)

	now := time.Now()
	for _, ptx := range snapshot {
		if ptx.nonce < confirmed || ptx.signed == nil {
			continue // done, or still being signed
		}
		age := now.Sub(ptx.submittedAt)
		switch {
		case ptx.retries >= m.config.MaxRetries:
			m.mu.Lock()
			if cur, ok := m.pending[ptx.nonce]; ok && cur.hash == ptx.hash {
				delete(m.pending, ptx.nonce)
			}
			m.mu.Unlock()
		case ptx.nonce >= chainPending && age >= m.config.MonitorInterval:
			m.resend(ctx, ptx, ptx.signed, false)
		case age >= m.config.MaxPendingTime:
			bumped, err := m.bumpedCopy(ptx.signed)
			if err != nil {
				continue
			}
			m.resend(ctx, ptx, bumped, true)
		}
	}
	return nil
}

// resend broadcasts tx and, if the pending entry is still the one we
// snapshotted, records the new hash. A bump restarts the pending clock; a
// plain re-broadcast does not, so a stuck transaction still gets bumped on
// schedule.
func (m *Manager) resend(ctx context.Context, ptx pendingTx, tx *gethtypes.Transaction, bumped bool) {
	if _, err := withTimeout(ctx, func(c context.Context) (struct{}, error) {
		return struct{}{}, m.clientFn().SendTransaction(c, tx)
	}); err != nil {
		return // "already known" and friends: the original may still confirm
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	cur, ok := m.pending[ptx.nonce]
	if !ok || cur.hash != ptx.hash {
		return
	}
	cur.hash = tx.Hash()
	cur.signed = tx
	cur.retries++
	if bumped {
		cur.submittedAt = time.Now()
	}
}

// withTimeout runs one RPC call under rpcCallTimeout.
func withTimeout[T any](ctx context.Context, call func(context.Context) (T, error)) (T, error) {
	cctx, cancel := context.WithTimeout(ctx, rpcCallTimeout)
	defer cancel()
	return call(cctx)
}

// bumpedCopy re-signs tx with fee caps increased by FeeIncreasePercent.
func (m *Manager) bumpedCopy(tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
	pct := big.NewInt(int64(100 + m.config.FeeIncreasePercent))
	bump := func(v *big.Int) *big.Int {
		out := new(big.Int).Mul(v, pct)
		return out.Div(out, big.NewInt(100))
	}
	raw := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID:    m.chainID,
		Nonce:      tx.Nonce(),
		GasTipCap:  bump(tx.GasTipCap()),
		GasFeeCap:  bump(tx.GasFeeCap()),
		Gas:        tx.Gas(),
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	})
	return gethtypes.SignTx(raw, gethtypes.LatestSignerForChainID(m.chainID), m.key)
}

// reconcile drops every pending entry below the confirmed nonce and keeps
// the local counter from falling behind it.
func (m *Manager) reconcile(confirmed uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for nonce := range m.pending {
		if nonce < confirmed {
			delete(m.pending, nonce)
		}
	}
	if confirmed > m.nextNonce {
		m.nextNonce = confirmed
	}
}

// pruneConfirmed reconciles the pending set against the sender's nonce at
// blockNumber (all our transactions mined by then are done).
func (m *Manager) pruneConfirmed(blockNumber *big.Int) {
	if blockNumber == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), rpcCallTimeout)
	defer cancel()
	confirmed, err := m.clientFn().NonceAt(ctx, m.from, blockNumber)
	if err != nil {
		return
	}
	m.reconcile(confirmed)
}

// ResetNonce re-reads the confirmed nonce from the chain and reconciles the
// pending set against it. Call this when the nonce counter is suspected to
// be out of sync.
func (m *Manager) ResetNonce(ctx context.Context) error {
	nonce, err := m.clientFn().NonceAt(ctx, m.from, nil)
	if err != nil {
		return fmt.Errorf("reset nonce: %w", err)
	}
	m.reconcile(nonce)
	return nil
}
