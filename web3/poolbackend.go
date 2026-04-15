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
type PooledBackend struct {
	pool *RPCPool
}

// NewPooledBackend creates a new PooledBackend wrapping the given RPCPool.
func NewPooledBackend(pool *RPCPool) *PooledBackend {
	return &PooledBackend{pool: pool}
}

// CodeAt implements ContractCaller.
func (p *PooledBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return p.pool.Current().CodeAt(ctx, contract, blockNumber)
}

// CallContract implements ContractCaller.
func (p *PooledBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return p.pool.Current().CallContract(ctx, call, blockNumber)
}

// EstimateGas implements GasEstimator (part of ContractTransactor).
func (p *PooledBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return p.pool.Current().EstimateGas(ctx, call)
}

// SuggestGasPrice implements GasPricer (part of ContractTransactor).
func (p *PooledBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return p.pool.Current().SuggestGasPrice(ctx)
}

// SuggestGasTipCap implements GasPricer1559 (part of ContractTransactor).
func (p *PooledBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return p.pool.Current().SuggestGasTipCap(ctx)
}

// SendTransaction implements TransactionSender (part of ContractTransactor).
// On transport errors it also marks the current endpoint as failed.
func (p *PooledBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := p.pool.Current().SendTransaction(ctx, tx)
	if err != nil && IsRPCTransportError(err) {
		p.pool.MarkFailed()
	}
	return err
}

// HeaderByNumber implements ContractTransactor.
func (p *PooledBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return p.pool.Current().HeaderByNumber(ctx, number)
}

// PendingCodeAt implements ContractTransactor.
func (p *PooledBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return p.pool.Current().PendingCodeAt(ctx, account)
}

// PendingNonceAt implements ContractTransactor.
func (p *PooledBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return p.pool.Current().PendingNonceAt(ctx, account)
}

// FilterLogs implements ContractFilterer.
func (p *PooledBackend) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return p.pool.Current().FilterLogs(ctx, q)
}

// SubscribeFilterLogs implements ContractFilterer.
func (p *PooledBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return p.pool.Current().SubscribeFilterLogs(ctx, q, ch)
}
