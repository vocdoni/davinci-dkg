package web3

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// PooledBackend implements bind.ContractBackend (= ContractCaller + ContractTransactor +
// ContractFilterer) over the pool's endpoints. The ethclient-style reads callers need
// beyond that interface (BlockNumber, TransactionByHash, TransactionReceipt) are here
// too, as is BatchCallContext, which lets BatchCall send several reads in one round
// trip. Every read goes to the pool's earliest usable endpoint and, when that endpoint
// is at fault (unreachable, rate-limiting us or refusing the request, e.g. an old log
// range or a receipt on a free tier), to the next one within the same call (see
// failover); contract reverts and other answers are returned as they are. Every
// request is counted per JSON-RPC method (RPCPool.Note).
type PooledBackend struct {
	pool *RPCPool
}

// NewPooledBackend creates a new PooledBackend wrapping the given RPCPool.
func NewPooledBackend(pool *RPCPool) *PooledBackend {
	return &PooledBackend{pool: pool}
}

// failover sends one request of method through call to the earliest usable
// endpoint and, while the endpoint is at fault (endpointFault), to the next
// one not tried yet, so a refusing or unreachable endpoint costs a round trip
// rather than the caller's whole attempt. Each attempt is counted and its
// outcome noted against the endpoint that served it. Under WithFairAttempts
// each attempt also runs within its share of the deadline, and an endpoint
// that runs out of it is disabled for the cooldown. The last outcome is
// returned once an endpoint answered, every usable one was tried or ctx is
// done, its error text cleared of the endpoints' secrets (redactor).
func failover[T any](
	ctx context.Context, pool *RPCPool, method string, call func(context.Context, *ethclient.Client) (T, error),
) (T, error) {
	out, err := tryEndpoints(ctx, pool, method, call)
	return out, pool.redact.error(err)
}

// tryEndpoints is failover before the error text is redacted.
func tryEndpoints[T any](
	ctx context.Context, pool *RPCPool, method string, call func(context.Context, *ethclient.Client) (T, error),
) (T, error) {
	tried := make([]bool, len(pool.entries))
	var (
		out T
		err error
	)
	for {
		// A done caller is the caller's failure: no endpoint is picked
		// (which could reset the pool), tried or blamed.
		if cerr := ctx.Err(); cerr != nil {
			if err == nil {
				err = cerr
			}
			return out, err
		}
		idx, client, ok := pool.next(tried)
		if !ok {
			return out, err
		}
		attempt, cancel := attemptContext(ctx, tried)
		tried[idx] = true
		pool.Note(method)
		out, err = call(attempt, client)
		// Only the attempt's own deadline is the endpoint's fault; the
		// caller's deadline or cancellation, which ends the attempt with
		// it, is not.
		slow := errors.Is(context.Cause(attempt), errAttemptTimeout)
		cancel()
		switch {
		case slow:
			pool.disable(idx, "too slow for a bounded read")
			continue
		case ctx.Err() != nil:
			return out, err
		}
		pool.noteAt(idx, err)
		if !endpointFault(err) {
			return out, err
		}
	}
}

// fairAttemptsKey marks a context made by WithFairAttempts.
type fairAttemptsKey struct{}

// WithFairAttempts makes the pooled reads under ctx, which carries a
// deadline, leave time for every endpoint: while other endpoints remain
// untried, an attempt gets an equal share of what remains among them and no
// more, and an endpoint that runs out of its share is disabled for the
// cooldown, so the read goes on to the next one and later reads (the next
// retention pass) start there. The last endpoint gets whatever remains; its
// running out is the caller's deadline, no endpoint fault. It is for
// bounded maintenance (a retention pass) that a single endpoint slow on one
// kind of read must not spend whole.
func WithFairAttempts(ctx context.Context) context.Context {
	return context.WithValue(ctx, fairAttemptsKey{}, true)
}

// errAttemptTimeout is the cause of an attempt whose own deadline, its
// share under WithFairAttempts, ran out.
var errAttemptTimeout = errors.New("rpc pool: attempt ran out of its share of the deadline")

// attemptContext is the context of the next attempt of a failover whose
// endpoints so far are marked in tried. Under WithFairAttempts, while
// others remain untried, it has a deadline of its own, its share of what
// remains, whose expiry has errAttemptTimeout as its cause.
func attemptContext(ctx context.Context, tried []bool) (context.Context, context.CancelFunc) {
	deadline, bounded := ctx.Deadline()
	untried := 0
	for _, t := range tried {
		if !t {
			untried++
		}
	}
	if fair, _ := ctx.Value(fairAttemptsKey{}).(bool); !fair || !bounded || untried <= 1 {
		return ctx, func() {}
	}
	return context.WithTimeoutCause(ctx, time.Until(deadline)/time.Duration(untried), errAttemptTimeout)
}

// CodeAt implements ContractCaller.
func (p *PooledBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return failover(ctx, p.pool, "eth_getCode", func(ctx context.Context, c *ethclient.Client) ([]byte, error) {
		return c.CodeAt(ctx, contract, blockNumber)
	})
}

// CallContract implements ContractCaller.
func (p *PooledBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return failover(ctx, p.pool, "eth_call", func(ctx context.Context, c *ethclient.Client) ([]byte, error) {
		return c.CallContract(ctx, call, blockNumber)
	})
}

// BatchCallContext sends every element of the batch to one endpoint in one
// JSON-RPC batch request (one HTTP round trip). The batch moves to the next
// endpoint when the request fails or one of its elements is refused by the
// endpoint (rate limit, capability); the last endpoint's per-element errors,
// a revert for instance, are left in the elements and only a failed request
// is returned.
func (p *PooledBackend) BatchCallContext(ctx context.Context, batch []rpc.BatchElem) error {
	answered := false
	_, err := failover(ctx, p.pool, "batch", func(ctx context.Context, c *ethclient.Client) (struct{}, error) {
		for i := range batch {
			p.pool.Note(batch[i].Method)
			batch[i].Error = nil
		}
		if err := c.Client().BatchCallContext(ctx, batch); err != nil {
			answered = false
			return struct{}{}, err
		}
		answered = true
		for _, elem := range batch {
			if endpointFault(elem.Error) {
				return struct{}{}, elem.Error
			}
		}
		return struct{}{}, nil
	})
	for i := range batch {
		batch[i].Error = p.pool.redact.error(batch[i].Error)
	}
	if answered {
		return nil
	}
	return err
}

// EstimateGas implements GasEstimator (part of ContractTransactor). It is pinned
// to the latest block, the state PendingCodeAt checks, so a broken "pending"
// state cannot skew the estimate.
func (p *PooledBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return failover(ctx, p.pool, "eth_estimateGas", func(ctx context.Context, c *ethclient.Client) (uint64, error) {
		return c.EstimateGasAtBlock(ctx, call, nil)
	})
}

// SuggestGasPrice implements GasPricer (part of ContractTransactor).
func (p *PooledBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return failover(ctx, p.pool, "eth_gasPrice", func(ctx context.Context, c *ethclient.Client) (*big.Int, error) {
		return c.SuggestGasPrice(ctx)
	})
}

// SuggestGasTipCap implements GasPricer1559 (part of ContractTransactor).
func (p *PooledBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return failover(ctx, p.pool, "eth_maxPriorityFeePerGas", func(ctx context.Context, c *ethclient.Client) (*big.Int, error) {
		return c.SuggestGasTipCap(ctx)
	})
}

// SendTransaction implements TransactionSender (part of ContractTransactor).
// A send is not repeated on another endpoint: after a transport error the
// first one may well have taken it, and the txmanager rebroadcasts what
// stays unmined.
func (p *PooledBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	idx, client, _ := p.pool.next(nil)
	p.pool.Note("eth_sendRawTransaction")
	err := client.SendTransaction(ctx, tx)
	p.pool.noteAt(idx, err)
	return p.pool.redact.error(err)
}

// BlockNumber returns the head height.
func (p *PooledBackend) BlockNumber(ctx context.Context) (uint64, error) {
	return failover(ctx, p.pool, "eth_blockNumber", func(ctx context.Context, c *ethclient.Client) (uint64, error) {
		return c.BlockNumber(ctx)
	})
}

// HeaderByNumber implements ContractTransactor.
func (p *PooledBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return failover(ctx, p.pool, "eth_getBlockByNumber", func(ctx context.Context, c *ethclient.Client) (*types.Header, error) {
		return c.HeaderByNumber(ctx, number)
	})
}

// TransactionByHash fetches a transaction.
func (p *PooledBackend) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	type found struct {
		tx      *types.Transaction
		pending bool
	}
	got, err := failover(ctx, p.pool, "eth_getTransactionByHash", func(ctx context.Context, c *ethclient.Client) (found, error) {
		tx, pending, err := c.TransactionByHash(ctx, hash)
		return found{tx, pending}, err
	})
	return got.tx, got.pending, err
}

// TransactionReceipt fetches a receipt; it is what transaction waits poll.
func (p *PooledBackend) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return failover(ctx, p.pool, "eth_getTransactionReceipt", func(ctx context.Context, c *ethclient.Client) (*types.Receipt, error) {
		return c.TransactionReceipt(ctx, hash)
	})
}

// PendingCodeAt implements ContractTransactor. It reads the code at the latest
// block: a pending transaction cannot remove a contract, bind only uses this to
// refuse a call to an address with no code, and some public endpoints
// (mainnet.optimism.io) answer "pending" with empty code for a live contract.
// A genuinely codeless address still returns empty code.
func (p *PooledBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return p.CodeAt(ctx, account, nil)
}

// PendingNonceAt implements ContractTransactor. It returns the larger of the
// latest and pending transaction counts, both read from one endpoint: some
// endpoints answer "pending" with 0 (or a stale value), which would make every
// send fail with "nonce too low". The latest count can lag the pool by our own
// in-flight transactions, so the txmanager treats the result as a lower bound
// and keeps its own counter above every nonce it has handed out
// (Manager.allocateNonce).
func (p *PooledBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return failover(ctx, p.pool, "eth_getTransactionCount", func(ctx context.Context, c *ethclient.Client) (uint64, error) {
		latest, err := c.NonceAt(ctx, account, nil)
		if err != nil {
			return 0, err
		}
		p.pool.Note("eth_getTransactionCount")
		pending, err := c.PendingNonceAt(ctx, account)
		if err != nil {
			return 0, err
		}
		return max(latest, pending), nil
	})
}

// FilterLogs implements ContractFilterer.
func (p *PooledBackend) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return failover(ctx, p.pool, "eth_getLogs", func(ctx context.Context, c *ethclient.Client) ([]types.Log, error) {
		return c.FilterLogs(ctx, q)
	})
}

// SubscribeFilterLogs implements ContractFilterer.
func (p *PooledBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	idx, client, _ := p.pool.next(nil)
	p.pool.Note("eth_subscribe")
	sub, err := client.SubscribeFilterLogs(ctx, q, ch)
	p.pool.noteAt(idx, err)
	return sub, p.pool.redact.error(err)
}

// BalanceAt reads an account's balance at blockNumber (nil = latest).
func (p *PooledBackend) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return failover(ctx, p.pool, "eth_getBalance", func(ctx context.Context, c *ethclient.Client) (*big.Int, error) {
		return c.BalanceAt(ctx, account, blockNumber)
	})
}
