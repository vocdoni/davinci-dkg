package txmanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	qt "github.com/frankban/quicktest"
)

const (
	testChainID = uint64(31337)
	testKey     = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

// fakeRPC is a JSON-RPC server with per-method handlers; unhandled methods
// return an error so every code path that tolerates RPC failures is
// exercised by default.
type fakeRPC struct {
	mu       sync.Mutex
	handlers map[string]func(params []json.RawMessage) (any, error)
	calls    map[string]int
	srv      *httptest.Server
}

func newFakeRPC(t *testing.T) *fakeRPC {
	t.Helper()
	f := &fakeRPC{handlers: map[string]func([]json.RawMessage) (any, error){}, calls: map[string]int{}}
	f.srv = httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeRPC) handle(method string, h func(params []json.RawMessage) (any, error)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.handlers[method] = h
}

func (f *fakeRPC) serve(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     json.RawMessage   `json:"id"`
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f.mu.Lock()
	f.calls[req.Method]++
	h := f.handlers[req.Method]
	f.mu.Unlock()

	resp := map[string]any{"jsonrpc": "2.0", "id": req.ID}
	if h == nil {
		resp["error"] = map[string]any{"code": -32601, "message": "fakeRPC: " + req.Method + " not handled"}
	} else if result, err := h(req.Params); err != nil {
		resp["error"] = map[string]any{"code": -32000, "message": err.Error()}
	} else {
		resp["result"] = result
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func newTestManager(t *testing.T, f *fakeRPC) *Manager {
	t.Helper()
	client, err := ethclient.Dial(f.srv.URL)
	qt.Assert(t, err, qt.IsNil)
	t.Cleanup(client.Close)
	m, err := New(func() *ethclient.Client { return client }, testChainID, testKey)
	qt.Assert(t, err, qt.IsNil)
	return m
}

func unsignedTx(nonce uint64) *gethtypes.Transaction {
	to := common.HexToAddress("0xdead")
	return gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID:   new(big.Int).SetUint64(testChainID),
		Nonce:     nonce,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(2),
		Gas:       21_000,
		To:        &to,
		Data:      []byte{1, 2, 3},
	})
}

func signedTx(t *testing.T, m *Manager, nonce uint64) *gethtypes.Transaction {
	t.Helper()
	signed, err := gethtypes.SignTx(unsignedTx(nonce), gethtypes.LatestSignerForChainID(m.chainID), m.key)
	qt.Assert(t, err, qt.IsNil)
	return signed
}

func receiptJSON(hash common.Hash, block int64) any {
	return map[string]any{
		"type":              "0x2",
		"status":            "0x1",
		"cumulativeGasUsed": "0x5208",
		"gasUsed":           "0x5208",
		"effectiveGasPrice": "0x2",
		"logsBloom":         hexutil.Bytes(make([]byte, 256)),
		"logs":              []any{},
		"transactionHash":   hash,
		"blockHash":         common.HexToHash("0x01"),
		"blockNumber":       hexutil.Uint64(block),
		"transactionIndex":  "0x0",
	}
}

// bind hands the Signer a transaction carrying the RPC's pending nonce; two
// goroutines (auto-create and the tick) can be handed the same one. The
// manager must allocate nonces from its own counter under the lock so both
// transactions are distinct and consecutive.
func TestSignerAllocatesDistinctNoncesConcurrently(t *testing.T) {
	c := qt.New(t)
	m := newTestManager(t, newFakeRPC(t))
	const chainPending, workers = uint64(10), 25

	nonces := make([]uint64, workers)
	var wg sync.WaitGroup
	for i := range workers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			auth, err := m.NewTransactOpts(context.Background())
			if err != nil {
				t.Error(err)
				return
			}
			signed, err := auth.Signer(auth.From, unsignedTx(chainPending))
			if err != nil {
				t.Error(err)
				return
			}
			from, err := gethtypes.Sender(gethtypes.LatestSignerForChainID(m.chainID), signed)
			if err != nil || from != m.Address() {
				t.Errorf("bad signature: %v %s", err, from)
			}
			nonces[i] = signed.Nonce()
		}(i)
	}
	wg.Wait()
	sort.Slice(nonces, func(i, j int) bool { return nonces[i] < nonces[j] })
	for i, n := range nonces {
		c.Assert(n, qt.Equals, chainPending+uint64(i))
	}
	c.Assert(m.pending, qt.HasLen, workers)
}

// The local counter follows the chain when an external sender moves it
// ahead, and re-uses any nonce below the counter that is no longer in
// flight so a dropped transaction cannot stall every later one.
func TestAllocateNonceFollowsTheChainAndFillsGaps(t *testing.T) {
	c := qt.New(t)
	m := &Manager{pending: map[uint64]*pendingTx{}}

	c.Assert(m.allocateNonce(5), qt.Equals, uint64(5))
	c.Assert(m.allocateNonce(5), qt.Equals, uint64(6)) // 5 still in flight
	c.Assert(m.allocateNonce(9), qt.Equals, uint64(9)) // chain moved ahead of us

	delete(m.pending, 6) // gave up on 6; the chain never saw 7 and 8 either
	c.Assert(m.allocateNonce(6), qt.Equals, uint64(6))
	c.Assert(m.allocateNonce(6), qt.Equals, uint64(7))
	c.Assert(m.allocateNonce(6), qt.Equals, uint64(8))
	c.Assert(m.allocateNonce(6), qt.Equals, uint64(10)) // 9 is in flight
}

// A flaky RPC must not turn into a spurious "transaction failed": receipt
// errors other than not-found are retried until the timeout.
func TestWaitTxByHashKeepsPollingThroughTransientErrors(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newTestManager(t, f)
	hash := common.HexToHash("0xabc")

	var receiptCalls int
	f.handle("eth_getTransactionReceipt", func([]json.RawMessage) (any, error) {
		receiptCalls++
		if receiptCalls < 3 {
			return nil, errors.New("upstream timeout")
		}
		return receiptJSON(hash, 7), nil
	})
	f.handle("eth_getTransactionCount", func([]json.RawMessage) (any, error) { return "0x1", nil })

	c.Assert(m.WaitTxByHash(hash, 10*time.Second), qt.IsNil)
	c.Assert(receiptCalls, qt.Equals, 3)
	c.Assert(m.TotalGasSpent().Int64(), qt.Equals, int64(21_000*2))

	f.handle("eth_getTransactionReceipt", func([]json.RawMessage) (any, error) {
		return nil, errors.New("still broken")
	})
	err := m.WaitTxByHash(hash, 1500*time.Millisecond)
	c.Assert(err, qt.ErrorMatches, "timeout waiting for transaction.*still broken.*")
}

// retryStuck must not hold the manager lock while it talks to the RPC:
// otherwise a slow endpoint blocks every NewTransactOpts / RecordPending
// caller for the duration of the call.
func TestRetryStuckDoesNotHoldTheLockDuringRPC(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newTestManager(t, f)
	m.RecordPending(signedTx(t, m, 3))

	entered := make(chan struct{})
	release := make(chan struct{})
	var once sync.Once
	f.handle("eth_getTransactionCount", func([]json.RawMessage) (any, error) {
		once.Do(func() { close(entered) })
		select {
		case <-release:
		case <-time.After(5 * time.Second):
		}
		return "0x3", nil
	})

	done := make(chan error, 1)
	go func() { done <- m.retryStuck(context.Background()) }()
	<-entered

	locked := make(chan int, 1)
	go func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		locked <- len(m.pending)
	}()
	select {
	case n := <-locked:
		c.Assert(n, qt.Equals, 1)
	case <-time.After(2 * time.Second):
		c.Fatal("manager lock is held across the RPC call")
	}
	close(release)
	c.Assert(<-done, qt.IsNil)
}

// A transaction the RPC's mempool does not know about (the send was lost or
// the pool failed over) is re-broadcast as-is on the next monitor pass.
func TestRetryStuckRebroadcastsTxMissingFromMempool(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newTestManager(t, f)
	tx := signedTx(t, m, 5)
	m.RecordPending(tx)
	m.pending[5].submittedAt = time.Now().Add(-2 * m.config.MonitorInterval)

	f.handle("eth_getTransactionCount", func([]json.RawMessage) (any, error) { return "0x5", nil })
	var sent []*gethtypes.Transaction
	f.handle("eth_sendRawTransaction", func(params []json.RawMessage) (any, error) {
		var raw hexutil.Bytes
		if err := json.Unmarshal(params[0], &raw); err != nil {
			return nil, err
		}
		var decoded gethtypes.Transaction
		if err := decoded.UnmarshalBinary(raw); err != nil {
			return nil, err
		}
		sent = append(sent, &decoded)
		return decoded.Hash(), nil
	})

	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	c.Assert(sent, qt.HasLen, 1)
	c.Assert(sent[0].Hash(), qt.Equals, tx.Hash())
	c.Assert(m.pending[5].retries, qt.Equals, 1)
	c.Assert(m.pending[5].hash, qt.Equals, tx.Hash())
}

// A transaction that has sat in the mempool for MaxPendingTime is replaced
// with a fee-bumped copy under the same nonce; confirmed nonces are pruned.
func TestRetryStuckBumpsFeesAndPrunesConfirmed(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newTestManager(t, f)
	m.RecordPending(signedTx(t, m, 4)) // confirmed below
	stuck := signedTx(t, m, 5)
	m.RecordPending(stuck)
	m.pending[5].submittedAt = time.Now().Add(-m.config.MaxPendingTime - time.Second)

	f.handle("eth_getTransactionCount", func(params []json.RawMessage) (any, error) {
		var tag string
		_ = json.Unmarshal(params[1], &tag)
		if tag == "pending" {
			return "0x6", nil // 5 is in the mempool
		}
		return "0x5", nil // 4 is mined
	})
	var sent []*gethtypes.Transaction
	f.handle("eth_sendRawTransaction", func(params []json.RawMessage) (any, error) {
		var raw hexutil.Bytes
		if err := json.Unmarshal(params[0], &raw); err != nil {
			return nil, err
		}
		var decoded gethtypes.Transaction
		if err := decoded.UnmarshalBinary(raw); err != nil {
			return nil, err
		}
		sent = append(sent, &decoded)
		return decoded.Hash(), nil
	})

	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	_, has4 := m.pending[4]
	c.Assert(has4, qt.IsFalse, qt.Commentf("nonce 4 is confirmed and must be pruned"))
	c.Assert(sent, qt.HasLen, 1)
	c.Assert(sent[0].Nonce(), qt.Equals, uint64(5))
	c.Assert(sent[0].GasFeeCap().Int64(), qt.Equals, int64(3)) // 2 × 1.5
	c.Assert(sent[0].Hash() != stuck.Hash(), qt.IsTrue)
	c.Assert(m.pending[5].hash, qt.Equals, sent[0].Hash())
	c.Assert(m.pending[5].retries, qt.Equals, 1)
	c.Assert(m.nextNonce, qt.Equals, uint64(5), qt.Commentf("counter follows the confirmed nonce"))

	// After MaxRetries the manager gives up so the nonce can be re-used.
	m.pending[5].retries = m.config.MaxRetries
	m.pending[5].submittedAt = time.Now().Add(-m.config.MaxPendingTime - time.Second)
	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	_, has5 := m.pending[5]
	c.Assert(has5, qt.IsFalse)
	c.Assert(fmt.Sprint(m.allocateNonce(5)), qt.Equals, "5")
}

// bind's gas estimate is exact for the state it was simulated against;
// transactions racing into the same block (three claimSlot calls, each
// costlier than the last) run out of gas without headroom. The Signer
// inflates estimated limits; an explicit GasLimit is left alone.
func TestSignerAppliesGasHeadroomToEstimatedGas(t *testing.T) {
	c := qt.New(t)
	m := newTestManager(t, newFakeRPC(t))

	auth, err := m.NewTransactOpts(context.Background())
	c.Assert(err, qt.IsNil)
	signed, err := auth.Signer(auth.From, unsignedTx(1))
	c.Assert(err, qt.IsNil)
	c.Assert(signed.Gas(), qt.Equals, uint64(21_000*120/100))

	m.SetGasMultiplier(1.5)
	auth, err = m.NewTransactOpts(context.Background())
	c.Assert(err, qt.IsNil)
	signed, err = auth.Signer(auth.From, unsignedTx(2))
	c.Assert(err, qt.IsNil)
	c.Assert(signed.Gas(), qt.Equals, uint64(21_000*150/100))

	auth, err = m.NewTransactOpts(context.Background())
	c.Assert(err, qt.IsNil)
	auth.GasLimit = 21_000 // caller-chosen: not an estimate, must not be touched
	signed, err = auth.Signer(auth.From, unsignedTx(3))
	c.Assert(err, qt.IsNil)
	c.Assert(signed.Gas(), qt.Equals, uint64(21_000))
}
