package web3

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PooledBackend implements bind.ContractBackend (= ContractCaller + ContractTransactor +
// ContractFilterer) by delegating every call to the pool's current active client.
// The ethclient-style reads callers need beyond that interface (BlockNumber,
// TransactionByHash, TransactionReceipt) are here too. Every endpoint error is
// routed through RPCPool.NoteError, so an endpoint that is rate-limiting us or
// unreachable rotates instead of being hammered until the caller's next retry;
// contract reverts and other call errors change nothing.
type PooledBackend struct {
	pool *RPCPool
}

// NewPooledBackend creates a new PooledBackend wrapping the given RPCPool.
func NewPooledBackend(pool *RPCPool) *PooledBackend {
	return &PooledBackend{pool: pool}
}

// CodeAt implements ContractCaller.
func (p *PooledBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	code, err := p.pool.Current().CodeAt(ctx, contract, blockNumber)
	p.pool.NoteError(err)
	return code, err
}

// CallContract implements ContractCaller.
func (p *PooledBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	output, err := p.pool.Current().CallContract(ctx, call, blockNumber)
	p.pool.NoteError(err)
	return output, err
}

// EstimateGas implements GasEstimator (part of ContractTransactor).
func (p *PooledBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	gas, err := p.pool.Current().EstimateGas(ctx, call)
	p.pool.NoteError(err)
	return gas, err
}

// SuggestGasPrice implements GasPricer (part of ContractTransactor).
func (p *PooledBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	price, err := p.pool.Current().SuggestGasPrice(ctx)
	p.pool.NoteError(err)
	return price, err
}

// SuggestGasTipCap implements GasPricer1559 (part of ContractTransactor).
func (p *PooledBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	tip, err := p.pool.Current().SuggestGasTipCap(ctx)
	p.pool.NoteError(err)
	return tip, err
}

// SendTransaction implements TransactionSender (part of ContractTransactor).
func (p *PooledBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := p.pool.Current().SendTransaction(ctx, tx)
	p.pool.NoteError(err)
	return err
}

// BlockNumber returns the head height of the current endpoint.
func (p *PooledBackend) BlockNumber(ctx context.Context) (uint64, error) {
	head, err := p.pool.Current().BlockNumber(ctx)
	p.pool.NoteError(err)
	return head, err
}

// HeaderByNumber implements ContractTransactor.
func (p *PooledBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := p.pool.Current().HeaderByNumber(ctx, number)
	p.pool.NoteError(err)
	return header, err
}

// TransactionByHash fetches a transaction from the current endpoint.
func (p *PooledBackend) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	tx, pending, err := p.pool.Current().TransactionByHash(ctx, hash)
	p.pool.NoteError(err)
	return tx, pending, err
}

// TransactionReceipt fetches a receipt from the current endpoint.
func (p *PooledBackend) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	receipt, err := p.pool.Current().TransactionReceipt(ctx, hash)
	p.pool.NoteError(err)
	return receipt, err
}

// PendingCodeAt implements ContractTransactor.
func (p *PooledBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	code, err := p.pool.Current().PendingCodeAt(ctx, account)
	p.pool.NoteError(err)
	return code, err
}

// PendingNonceAt implements ContractTransactor.
func (p *PooledBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	nonce, err := p.pool.Current().PendingNonceAt(ctx, account)
	p.pool.NoteError(err)
	return nonce, err
}

// FilterLogs implements ContractFilterer.
func (p *PooledBackend) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := p.pool.Current().FilterLogs(ctx, q)
	p.pool.NoteError(err)
	return logs, err
}

// SubscribeFilterLogs implements ContractFilterer.
func (p *PooledBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	sub, err := p.pool.Current().SubscribeFilterLogs(ctx, q, ch)
	p.pool.NoteError(err)
	return sub, err
}
