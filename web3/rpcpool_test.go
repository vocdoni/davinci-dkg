package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	qt "github.com/frankban/quicktest"
)

// fakeRPCError spells a provider JSON-RPC error the way go-ethereum surfaces
// it: a type implementing rpc.Error with ErrorCode().
type fakeRPCError struct {
	code int
	msg  string
}

func (e fakeRPCError) Error() string  { return e.msg }
func (e fakeRPCError) ErrorCode() int { return e.code }

var _ rpc.Error = fakeRPCError{}

// Rate-limit responses (HTTP 429, JSON-RPC -32005 or a message naming it)
// must count as endpoint failures so the pool rotates instead of hammering
// the same throttled endpoint; contract reverts and unrelated JSON-RPC codes
// must not.
func TestRateLimitErrorsAreTransportErrors(t *testing.T) {
	cases := []struct {
		name      string
		err       error
		rateLimit bool
		transport bool
	}{
		{"nil", nil, false, false},
		{"json-rpc -32005", fakeRPCError{code: -32005, msg: "rate limit exceeded"}, true, true},
		{"wrapped json-rpc -32005", fmt.Errorf("filter logs: %w", fakeRPCError{code: -32005, msg: "too many requests"}), true, true},
		{"json-rpc -32005 code with neutral message", fakeRPCError{code: -32005, msg: "exceeded quota"}, true, true},
		{"json-rpc other code", fakeRPCError{code: -32603, msg: "internal error"}, false, false},
		{"http 429", errors.New("429 Too Many Requests"), true, true},
		{"http 429 lowercase", errors.New("http 429: quota exceeded"), true, true},
		{"rate limit mixed case", errors.New("Rate Limit Exceeded"), true, true},
		{"connectivity", errors.New("dial tcp 10.0.0.1:443: connection refused"), false, true},
		{"contract revert", errors.New("execution reverted: PoolExhausted()"), false, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := qt.New(t)
			c.Assert(IsRateLimitError(tc.err), qt.Equals, tc.rateLimit)
			c.Assert(IsRPCTransportError(tc.err), qt.Equals, tc.transport)
		})
	}
}

// newTestPool builds a pool of entries without dialing anything; the failure
// bookkeeping under test never touches the clients.
func newTestPool(urls ...string) *RPCPool {
	entries := make([]*rpcEntry, len(urls))
	for i, u := range urls {
		entries[i] = &rpcEntry{url: u}
	}
	return &RPCPool{entries: entries}
}

// A rate-limited endpoint rotates after a single failure, while connectivity
// failures keep the rpcMaxFailures threshold; the cooldown eventually brings
// the rate-limited endpoint back.
func TestRateLimitRotatesImmediatelyConnectivityKeepsThreshold(t *testing.T) {
	c := qt.New(t)

	p := newTestPool("a", "b")
	c.Assert(p.CurrentURL(), qt.Equals, "a")
	p.MarkRateLimited()
	c.Assert(p.CurrentURL(), qt.Equals, "b", qt.Commentf("a rate-limited endpoint must rotate after one failure"))
	c.Assert(p.entries[0].disabledAt.IsZero(), qt.IsFalse, qt.Commentf("the rate-limited endpoint must be disabled"))

	// The pool keeps serving "b" while it is healthy; once the cooldown on
	// "a" has expired and "b" is rate-limited too, Current() re-enables "a".
	_ = p.Current()
	c.Assert(p.CurrentURL(), qt.Equals, "b")
	p.entries[0].disabledAt = time.Now().Add(-rpcCooldownDuration)
	p.MarkRateLimited() // "b" rotates to the cooled-down "a"
	c.Assert(p.CurrentURL(), qt.Equals, "a")

	q := newTestPool("a", "b")
	for i := 1; i < rpcMaxFailures; i++ {
		q.MarkFailed()
		c.Assert(q.CurrentURL(), qt.Equals, "a", qt.Commentf("connectivity failure %d rotated before the threshold", i))
		c.Assert(q.entries[0].disabledAt.IsZero(), qt.IsTrue, qt.Commentf("connectivity failure %d disabled the endpoint", i))
	}
	q.MarkFailed()
	c.Assert(q.CurrentURL(), qt.Equals, "b", qt.Commentf("the third connectivity failure must rotate"))
	c.Assert(q.entries[0].disabledAt.IsZero(), qt.IsFalse)
}

// The incident path: a FilterLogs (eth_getLogs) 429 from the current endpoint
// must disable it at once and rotate, instead of the pool hammering the same
// throttled endpoint on every scan. A healthy second endpoint takes over.
func TestFilterLogsRateLimitRotatesPool(t *testing.T) {
	c := qt.New(t)
	rateLimited := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"error":{"code":-32005,"message":"rate limit exceeded"}}`))
	}))
	defer rateLimited.Close()
	healthy := testRPCServer() // answers eth_getLogs with an empty result
	defer healthy.Close()

	pool, err := NewRPCPool([]string{rateLimited.URL, healthy.URL})
	c.Assert(err, qt.IsNil)
	defer pool.Close()

	backend := NewPooledBackend(pool)
	before := pool.CurrentURL()
	_, err = backend.FilterLogs(context.Background(), ethereum.FilterQuery{FromBlock: big.NewInt(1)})
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(IsRateLimitError(err), qt.IsTrue, qt.Commentf("the 429 response must classify as a rate limit, got %v", err))
	c.Assert(pool.CurrentURL(), qt.Not(qt.Equals), before, qt.Commentf("a FilterLogs 429 must rotate the pool"))
}
