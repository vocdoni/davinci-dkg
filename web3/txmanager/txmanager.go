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
	defaultTxTimeout          = 90 * time.Second
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
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() Config {
	return Config{
		MaxPendingTime:     defaultMaxPendingTime,
		MaxRetries:         defaultMaxRetries,
		FeeIncreasePercent: defaultFeeIncreasePercent,
		MonitorInterval:    defaultMonitorInterval,
	}
}

// pendingTx tracks a submitted transaction that has not yet been confirmed.
type pendingTx struct {
	hash        common.Hash
	nonce       uint64
	to          common.Address
	data        []byte
	gasFeeCap   *big.Int
	gasTipCap   *big.Int
	gasLimit    uint64
	submittedAt time.Time
	retries     int
}

// Manager handles nonce management, EIP-1559 gas estimation, and basic
// stuck-transaction recovery. It is safe for concurrent use.
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
				m.mu.Lock()
				_ = m.retryStuck(monCtx)
				m.mu.Unlock()
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

// NewTransactOpts builds an EIP-1559 TransactOpts.
// Fee caps are estimated from current network conditions.
func (m *Manager) NewTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	auth, err := bind.NewKeyedTransactorWithChainID(m.key, m.chainID)
	if err != nil {
		return nil, fmt.Errorf("create transact opts: %w", err)
	}
	auth.Context = ctx

	tipCap, feeCap, err := m.suggestFees(ctx)
	if err != nil {
		// Non-fatal: let the bound method use default gas price.
		return auth, nil
	}
	auth.GasTipCap = tipCap
	auth.GasFeeCap = feeCap

	return auth, nil
}

// RecordPending stores a submitted transaction for monitoring. Call this
// after successfully broadcasting a transaction obtained through NewTransactOpts.
func (m *Manager) RecordPending(tx *gethtypes.Transaction) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.pending[tx.Nonce()] = &pendingTx{
		hash:        tx.Hash(),
		nonce:       tx.Nonce(),
		to:          *tx.To(),
		data:        tx.Data(),
		gasFeeCap:   tx.GasFeeCap(),
		gasTipCap:   tx.GasTipCap(),
		gasLimit:    tx.Gas(),
		submittedAt: time.Now(),
	}
}

// WaitTxByHash blocks until the transaction is confirmed or the timeout expires.
// It returns an error if the transaction reverts or the context/timeout fires.
func (m *Manager) WaitTxByHash(hash common.Hash, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for transaction %s", hash.Hex())
		case <-ticker.C:
			receipt, err := m.clientFn().TransactionReceipt(ctx, hash)
			switch {
			case err == nil:
				// Always record gas cost — reverted txs still consume gas.
				m.recordGasSpent(receipt)
				if receipt.Status != gethtypes.ReceiptStatusSuccessful {
					return fmt.Errorf("transaction %s reverted (status %d)", hash.Hex(), receipt.Status)
				}
				m.pruneConfirmed(receipt.BlockNumber)
				return nil
			case errors.Is(err, ethereum.NotFound):
				// Still pending — keep polling.
			default:
				return fmt.Errorf("get receipt for %s: %w", hash.Hex(), err)
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

// retryStuck resubmits transactions that have been pending longer than
// MaxPendingTime with a bumped fee cap. Must be called with m.mu held.
func (m *Manager) retryStuck(ctx context.Context) error {
	now := time.Now()
	for nonce, ptx := range m.pending {
		if now.Sub(ptx.submittedAt) < m.config.MaxPendingTime {
			continue
		}
		if ptx.retries >= m.config.MaxRetries {
			// Give up and remove from tracking.
			delete(m.pending, nonce)
			continue
		}

		// Check if already confirmed.
		receipt, err := m.clientFn().TransactionReceipt(ctx, ptx.hash)
		if err == nil && receipt != nil {
			// Confirmed — remove.
			delete(m.pending, nonce)
			continue
		}

		// Bump fee and resubmit.
		bumped, bumpedTip, err := m.bumpedFees(ptx)
		if err != nil {
			continue
		}
		tx := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
			ChainID:   m.chainID,
			Nonce:     ptx.nonce,
			GasTipCap: bumpedTip,
			GasFeeCap: bumped,
			Gas:       ptx.gasLimit,
			To:        &ptx.to,
			Data:      ptx.data,
		})
		signer := gethtypes.NewCancunSigner(m.chainID)
		signed, err := gethtypes.SignTx(tx, signer, m.key)
		if err != nil {
			continue
		}
		if err := m.clientFn().SendTransaction(ctx, signed); err != nil {
			// Ignore "already known" errors; the original tx may still confirm.
			continue
		}
		ptx.hash = signed.Hash()
		ptx.gasFeeCap = bumped
		ptx.gasTipCap = bumpedTip
		ptx.submittedAt = now
		ptx.retries++
	}
	return nil
}

// bumpedFees returns fee caps increased by FeeIncreasePercent.
func (m *Manager) bumpedFees(ptx *pendingTx) (feeCap, tipCap *big.Int, err error) {
	pct := big.NewInt(int64(100 + m.config.FeeIncreasePercent))
	feeCap = new(big.Int).Mul(ptx.gasFeeCap, pct)
	feeCap.Div(feeCap, big.NewInt(100))
	tipCap = new(big.Int).Mul(ptx.gasTipCap, pct)
	tipCap.Div(tipCap, big.NewInt(100))
	return feeCap, tipCap, nil
}

// pruneConfirmed removes all pending entries with nonce ≤ confirmedBlockNonce.
// We use the block number as a rough proxy (all txs mined by this block are done).
func (m *Manager) pruneConfirmed(blockNumber *big.Int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if blockNumber == nil {
		return
	}
	// Refresh confirmed nonce to detect externally mined transactions.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	confirmed, err := m.clientFn().NonceAt(ctx, m.from, blockNumber)
	if err != nil {
		return
	}
	for nonce := range m.pending {
		if nonce < confirmed {
			delete(m.pending, nonce)
		}
	}
	// Advance our local nonce if the confirmed nonce has overtaken it.
	if confirmed > m.nextNonce {
		m.nextNonce = confirmed
	}
}

// ResetNonce re-reads the confirmed nonce from the chain. Call this after a
// node restart or when the nonce counter becomes out-of-sync.
func (m *Manager) ResetNonce(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	nonce, err := m.clientFn().NonceAt(ctx, m.from, nil)
	if err != nil {
		return fmt.Errorf("reset nonce: %w", err)
	}
	m.nextNonce = nonce
	return nil
}
