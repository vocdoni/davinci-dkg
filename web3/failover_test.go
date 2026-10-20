package web3

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// ethCall is the JSON-RPC method of a contract read.
const ethCall = "eth_call"

// archiveRefusal is publicnode's answer, with HTTP 403, to an eth_getLogs
// range older than its window and to every eth_getTransactionReceipt.
const archiveRefusal = `{"jsonrpc":"2.0","id":1,"error":{"code":-32602,` +
	`"message":"Archive requests require a personal token. Get one at: https://example.com"}}`

// fakeEndpoint is a JSON-RPC endpoint at a fixed head serving one log and
// one receipt. archiveWindow > 0 makes it behave like publicnode: HTTP 403
// for eth_getLogs ranges starting more than archiveWindow blocks behind the
// head and for every receipt. status > 0 answers every request with that
// HTTP status instead (403 for a token-gated endpoint, 429 for a rate limit).
type fakeEndpoint struct {
	head          uint64
	archiveWindow uint64
	status        int
	log           types.Log
	receipt       *types.Receipt

	mu    sync.Mutex
	calls map[string]int
}

func newFakeEndpoint(t *testing.T, f *fakeEndpoint) string {
	t.Helper()
	f.calls = map[string]int{}
	srv := httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(srv.Close)
	return srv.URL
}

func (f *fakeEndpoint) count(method string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls[method]
}

func (f *fakeEndpoint) serve(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	batch := len(body) > 0 && body[0] == '['
	var reqs []testRPCRequest
	if batch {
		_ = json.Unmarshal(body, &reqs)
	} else {
		var one testRPCRequest
		_ = json.Unmarshal(body, &one)
		reqs = []testRPCRequest{one}
	}
	f.mu.Lock()
	for _, req := range reqs {
		f.calls[req.Method]++
	}
	f.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if f.status != 0 {
		w.WriteHeader(f.status)
		refusal := archiveRefusal
		if f.status == http.StatusTooManyRequests {
			refusal = `{"jsonrpc":"2.0","id":1,"error":{"code":-32005,"message":"rate limit exceeded"}}`
		}
		_, _ = w.Write([]byte(refusal))
		return
	}
	resps := make([]testRPCResponse, len(reqs))
	for i, req := range reqs {
		resp, refused := f.answer(req)
		if refused {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(archiveRefusal))
			return
		}
		resps[i] = resp
	}
	if batch {
		_ = json.NewEncoder(w).Encode(resps)
		return
	}
	_ = json.NewEncoder(w).Encode(resps[0])
}

func (f *fakeEndpoint) answer(req testRPCRequest) (testRPCResponse, bool) {
	resp := testRPCResponse{JSONRPC: "2.0", ID: req.ID}
	switch req.Method {
	case "eth_blockNumber":
		resp.Result = hexutil.EncodeUint64(f.head)
	case ethCall:
		resp.Result = "0x01"
	case "eth_getLogs":
		var q struct {
			FromBlock hexutil.Uint64 `json:"fromBlock"`
		}
		_ = json.Unmarshal(req.Params[0], &q)
		if f.archiveWindow > 0 && uint64(q.FromBlock)+f.archiveWindow < f.head {
			return resp, true
		}
		resp.Result = []types.Log{f.log}
	case "eth_getTransactionReceipt":
		if f.archiveWindow > 0 {
			return resp, true
		}
		if f.receipt == nil {
			resp.Result = json.RawMessage("null") // unknown transaction
			break
		}
		resp.Result = f.receipt
	default:
		resp.Error = &testRPCError{Code: -32601, Message: "not handled"}
	}
	return resp, false
}

// fakeLogAndReceipt is a log and the receipt of the transaction that
// emitted it, as an archive-capable endpoint serves them.
func fakeLogAndReceipt() (types.Log, *types.Receipt) {
	txHash := common.HexToHash("0xabcd")
	log := types.Log{
		Address:     common.HexToAddress("0x2000000000000000000000000000000000000002"),
		Topics:      []common.Hash{{1}},
		Data:        []byte{},
		BlockNumber: 100,
		TxHash:      txHash,
		BlockHash:   common.Hash{2},
	}
	receipt := &types.Receipt{
		Status:      types.ReceiptStatusSuccessful,
		Logs:        []*types.Log{},
		TxHash:      txHash,
		BlockHash:   common.Hash{2},
		BlockNumber: big.NewInt(100),
		GasUsed:     21_000,
	}
	return log, receipt
}

// The incident: the pool sits on publicnode, which refuses eth_getLogs older
// than its window and every receipt with HTTP 403. The same call must be
// answered by the next endpoint, so a share is built (and a transaction
// confirmed) in the tick that needed it rather than failing the slot. A
// recent range stays on the endpoint that serves it.
func TestPooledBackendServesRefusedCallsFromTheNextEndpoint(t *testing.T) {
	ctx := context.Background()
	log, receipt := fakeLogAndReceipt()
	const head, window = 50_000, 10_000
	setup := func(t *testing.T) (*PooledBackend, *fakeEndpoint, *fakeEndpoint) {
		t.Helper()
		refusing := &fakeEndpoint{head: head, archiveWindow: window, log: log}
		archive := &fakeEndpoint{head: head, log: log, receipt: receipt}
		pool, err := NewRPCPool([]string{newFakeEndpoint(t, refusing), newFakeEndpoint(t, archive)})
		qt.Assert(t, err, qt.IsNil)
		t.Cleanup(pool.Close)
		return NewPooledBackend(pool), refusing, archive
	}

	t.Run("recent logs", func(t *testing.T) {
		c := qt.New(t)
		backend, refusing, archive := setup(t)
		logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(head - 100), ToBlock: big.NewInt(head)})
		c.Assert(err, qt.IsNil)
		c.Assert(logs, qt.HasLen, 1)
		c.Assert(refusing.count("eth_getLogs"), qt.Equals, 1)
		c.Assert(archive.count("eth_getLogs"), qt.Equals, 0, qt.Commentf("a served call must not be repeated"))
	})
	t.Run("old logs", func(t *testing.T) {
		c := qt.New(t)
		backend, refusing, archive := setup(t)
		from := big.NewInt(head - 3*window)
		logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: from, ToBlock: new(big.Int).Add(from, big.NewInt(9_999))})
		c.Assert(err, qt.IsNil, qt.Commentf("an old range refused by the first endpoint must be served by the next"))
		c.Assert(logs, qt.HasLen, 1)
		c.Assert(logs[0].TxHash, qt.Equals, log.TxHash)
		c.Assert(refusing.count("eth_getLogs"), qt.Equals, 1)
		c.Assert(archive.count("eth_getLogs"), qt.Equals, 1)
	})
	t.Run("receipt", func(t *testing.T) {
		c := qt.New(t)
		backend, refusing, archive := setup(t)
		got, err := backend.TransactionReceipt(ctx, receipt.TxHash)
		c.Assert(err, qt.IsNil, qt.Commentf("a refused receipt must be served by the next endpoint"))
		c.Assert(got.Status, qt.Equals, types.ReceiptStatusSuccessful)
		c.Assert(refusing.count("eth_getTransactionReceipt"), qt.Equals, 1)
		c.Assert(archive.count("eth_getTransactionReceipt"), qt.Equals, 1)
	})
	t.Run("batch", func(t *testing.T) {
		c := qt.New(t)
		refusing := &fakeEndpoint{head: head, status: http.StatusForbidden}
		archive := &fakeEndpoint{head: head}
		pool, err := NewRPCPool([]string{newFakeEndpoint(t, refusing), newFakeEndpoint(t, archive)})
		c.Assert(err, qt.IsNil)
		defer pool.Close()
		var out hexutil.Bytes
		batch := []rpc.BatchElem{{Method: ethCall, Args: []any{map[string]any{}, "latest"}, Result: &out}}
		c.Assert(NewPooledBackend(pool).BatchCallContext(ctx, batch), qt.IsNil)
		c.Assert(batch[0].Error, qt.IsNil)
		c.Assert([]byte(out), qt.DeepEquals, []byte{1})
		c.Assert(archive.count(ethCall), qt.Equals, 1)
	})
}

// A rate-limited or unreachable endpoint is skipped inside the call too, and
// every request sent is counted. With no other endpoint the error surfaces.
func TestPooledBackendFailsOverRateLimitsAndTransportErrors(t *testing.T) {
	ctx := context.Background()
	log, _ := fakeLogAndReceipt()
	unreachable := httptest.NewServer(http.NotFoundHandler())
	unreachable.Close()
	for _, tc := range []struct {
		name string
		url  func(t *testing.T) string
	}{
		{"rate limit", func(t *testing.T) string {
			return newFakeEndpoint(t, &fakeEndpoint{status: http.StatusTooManyRequests})
		}},
		{"token-gated", func(t *testing.T) string {
			return newFakeEndpoint(t, &fakeEndpoint{status: http.StatusForbidden})
		}},
		{"unreachable", func(*testing.T) string { return unreachable.URL }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := qt.New(t)
			bad := tc.url(t)
			good := &fakeEndpoint{head: 100, log: log}
			pool, err := NewRPCPool([]string{bad, newFakeEndpoint(t, good)})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			logs, err := NewPooledBackend(pool).FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
			c.Assert(err, qt.IsNil)
			c.Assert(logs, qt.HasLen, 1)
			c.Assert(pool.Snapshot()["eth_getLogs"], qt.Equals, uint64(2), qt.Commentf("both attempts are requests"))

			alone, err := NewRPCPool([]string{bad})
			c.Assert(err, qt.IsNil)
			defer alone.Close()
			_, err = NewPooledBackend(alone).FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
			c.Assert(err, qt.IsNotNil)
			c.Assert(endpointFault(err), qt.IsTrue, qt.Commentf("got %v", err))
		})
	}
}

// An answer that is the call's own outcome (not found, a revert, a method
// the endpoint does not know) is returned as is: asking the next endpoint
// would only double the request count.
func TestPooledBackendDoesNotRepeatAnsweredCalls(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	first := &fakeEndpoint{head: 100}
	second := &fakeEndpoint{head: 100}
	pool, err := NewRPCPool([]string{newFakeEndpoint(t, first), newFakeEndpoint(t, second)})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := NewPooledBackend(pool)

	_, err = backend.TransactionReceipt(ctx, common.Hash{9})
	c.Assert(errors.Is(err, ethereum.NotFound), qt.IsTrue, qt.Commentf("got %v", err))
	_, err = backend.HeaderByNumber(ctx, nil) // -32601 from the fake
	c.Assert(err, qt.IsNotNil)
	c.Assert(first.count("eth_getTransactionReceipt"), qt.Equals, 1)
	c.Assert(second.count("eth_getTransactionReceipt")+second.count("eth_getBlockByNumber"), qt.Equals, 0)
	c.Assert(pool.CurrentURL(), qt.Equals, pool.entries[0].url)
}

// A transaction wait polls receipts through the pooled backend: with the
// pool's first endpoint refusing every receipt (publicnode) the wait still
// sees the transaction mined, where polling the raw current client times
// out ("wait contribution tx: timeout" on the OP operators).
func TestTxWaitSurvivesAnEndpointThatRefusesReceipts(t *testing.T) {
	c := qt.New(t)
	log, receipt := fakeLogAndReceipt()
	refusing := &fakeEndpoint{head: 100, archiveWindow: 10, log: log}
	archive := &fakeEndpoint{head: 100, log: log, receipt: receipt}
	pool, err := NewRPCPool([]string{newFakeEndpoint(t, refusing), newFakeEndpoint(t, archive)})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	const anvilKey = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	raw, err := txmanager.New(pool.Current, 31337, anvilKey)
	c.Assert(err, qt.IsNil)
	c.Assert(raw.WaitTxByHash(receipt.TxHash, 1500*time.Millisecond), qt.ErrorMatches, "timeout waiting for transaction .*403 Forbidden.*")

	txm, err := txmanager.New(pool.Current, 31337, anvilKey)
	c.Assert(err, qt.IsNil)
	txm.SetReceiptReader(NewPooledBackend(pool))
	c.Assert(txm.WaitTxByHash(receipt.TxHash, 5*time.Second), qt.IsNil)
	c.Assert(refusing.count("eth_getTransactionReceipt"), qt.Not(qt.Equals), 0)
	c.Assert(archive.count("eth_getTransactionReceipt"), qt.Equals, 1)
}

// gatewayDown answers every request with an HTTP server error and a gateway's
// plain body, as a primary behind an unavailable load balancer does.
func gatewayDown(t *testing.T, status int) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(status)
		_, _ = w.Write([]byte("<html><body>origin unavailable</body></html>"))
	}))
	t.Cleanup(srv.Close)
	return srv.URL
}

// Every HTTP 5xx is the endpoint's fault, gateway-specific ones (521, 524)
// included, whatever the body says: the call moves to the next endpoint
// (startup, logs, receipts and transaction waits all work through a healthy
// fallback) and they count toward the consecutive-failure threshold like a
// transport error. A 4xx such as 404 does not.
func TestServerErrorsFailOver(t *testing.T) {
	ctx := context.Background()
	log0, receipt := fakeLogAndReceipt()
	for _, status := range []int{
		http.StatusInternalServerError, http.StatusNotImplemented, http.StatusBadGateway,
		http.StatusServiceUnavailable, http.StatusGatewayTimeout, 521, 524, 599,
	} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			c := qt.New(t)
			down := gatewayDown(t, status)
			healthy := &fakeEndpoint{head: 100, log: log0, receipt: receipt}
			pool, err := NewRPCPool([]string{down, newFakeEndpoint(t, healthy)})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			backend := NewPooledBackend(pool)

			logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
			c.Assert(err, qt.IsNil)
			c.Assert(logs, qt.HasLen, 1)
			got, err := backend.TransactionReceipt(ctx, receipt.TxHash)
			c.Assert(err, qt.IsNil)
			c.Assert(got.TxHash, qt.Equals, receipt.TxHash)
			c.Assert(pool.entries[0].failures, qt.Equals, 2)
			c.Assert(pool.entries[0].disabledAt.IsZero(), qt.IsTrue, qt.Commentf("two server errors in a row are below the threshold"))

			txm, err := txmanager.New(pool.Current, 31337, "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
			c.Assert(err, qt.IsNil)
			txm.SetReceiptReader(backend)
			c.Assert(txm.WaitTxByHash(receipt.TxHash, 5*time.Second), qt.IsNil)
			c.Assert(pool.entries[0].disabledAt.IsZero(), qt.IsFalse, qt.Commentf("the third in a row disables the endpoint"))

			startup := testRPCServer()
			defer startup.Close()
			contracts, err := NewOnChain([]string{down, startup.URL}, testAddresses(), 31337)
			c.Assert(err, qt.IsNil, qt.Commentf("startup must fail over a %d", status))
			c.Assert(contracts.Close(), qt.IsNil)
		})
	}
	notFound := rpc.HTTPError{StatusCode: http.StatusNotFound, Status: "404 Not Found"}
	qt.Check(t, endpointFault(notFound), qt.IsFalse)
}

// A bounded read that starts with little time left (here 1.5 s, under the
// old 2 s floor) still leaves the secondary its share: the primary gets at
// most its fair share of what remains, and running out of it puts the
// primary in cooldown instead of expiring the caller's deadline with it.
func TestFairAttemptsLeaveTheSecondaryItsShare(t *testing.T) {
	c := qt.New(t)
	stalled := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body) // read the request, so the server notices the client leave
		<-r.Context().Done()
	}))
	defer stalled.Close()
	log0, _ := fakeLogAndReceipt()
	pool, err := NewRPCPool([]string{stalled.URL, newFakeEndpoint(t, &fakeEndpoint{head: 100, log: log0})})
	c.Assert(err, qt.IsNil)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	logs, err := NewPooledBackend(pool).FilterLogs(WithFairAttempts(ctx), ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
	c.Assert(err, qt.IsNil, qt.Commentf("the secondary must answer within the caller's deadline"))
	c.Assert(logs, qt.HasLen, 1)
	c.Assert(pool.entries[0].disabledAt.IsZero(), qt.IsFalse, qt.Commentf("the stalled primary must enter cooldown"))
}

// A read whose caller is already done (an expired pass deadline, a batch
// falling back to single calls with the same context) is the caller's
// failure: it sends nothing and leaves every endpoint's health as it was,
// cooling-down endpoints included, fair attempts or not.
func TestADoneCallerLeavesEndpointHealthAlone(t *testing.T) {
	log0, _ := fakeLogAndReceipt()
	query := ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)}
	expired, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	<-expired.Done()
	for name, ctx := range map[string]context.Context{"plain": expired, "fair": WithFairAttempts(expired)} {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			first, second := &fakeEndpoint{head: 100, log: log0}, &fakeEndpoint{head: 100, log: log0}
			pool, err := NewRPCPool([]string{newFakeEndpoint(t, first), newFakeEndpoint(t, second)})
			c.Assert(err, qt.IsNil)
			defer pool.Close()
			backend := NewPooledBackend(pool)

			_, err = backend.FilterLogs(ctx, query)
			c.Assert(errors.Is(err, context.DeadlineExceeded), qt.IsTrue, qt.Commentf("got %v", err))
			for i, e := range pool.entries {
				c.Assert(e.disabledAt.IsZero(), qt.IsTrue, qt.Commentf("endpoint %d put in cooldown", i))
				c.Assert(e.failures, qt.Equals, 0, qt.Commentf("endpoint %d blamed", i))
			}
			c.Assert(first.count("eth_getLogs")+second.count("eth_getLogs"), qt.Equals, 0)

			// Every endpoint cooling down: a done caller does not reset them either.
			for _, e := range pool.entries {
				e.disabledAt = time.Now()
			}
			_, err = backend.FilterLogs(ctx, query)
			c.Assert(err, qt.IsNotNil)
			for i, e := range pool.entries {
				c.Assert(e.disabledAt.IsZero(), qt.IsFalse, qt.Commentf("endpoint %d reset", i))
			}
		})
	}
}
