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
// throttled endpoint on every scan. A healthy second endpoint takes over,
// within the same call; alone, the throttled endpoint's error surfaces and
// classifies as a rate limit.
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
	c.Assert(err, qt.IsNil, qt.Commentf("the next endpoint must answer the throttled call"))
	c.Assert(pool.CurrentURL(), qt.Not(qt.Equals), before, qt.Commentf("a FilterLogs 429 must rotate the pool"))
	c.Assert(pool.entries[0].disabledAt.IsZero(), qt.IsFalse)

	alone, err := NewRPCPool([]string{rateLimited.URL})
	c.Assert(err, qt.IsNil)
	defer alone.Close()
	_, err = NewPooledBackend(alone).FilterLogs(context.Background(), ethereum.FilterQuery{FromBlock: big.NewInt(1)})
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(IsRateLimitError(err), qt.IsTrue, qt.Commentf("the 429 response must classify as a rate limit, got %v", err))
}

// errConnReset is the scattered failure the OP operators saw from their
// private RPC: a transport error, not a refusal.
var errConnReset = errors.New("read tcp 10.0.0.2:41234->10.0.0.1:443: read: connection reset by peer")

// Only consecutive transport failures disable an endpoint: a primary that
// drops a connection now and then, with calls it answered in between (a
// result, or an error that is the call's own such as a revert), stays the
// active endpoint however many such failures it collects over the node's
// life.
func TestScatteredTransportErrorsKeepThePrimary(t *testing.T) {
	for _, answer := range []struct {
		name string
		err  error
	}{
		{"success", nil},
		{"revert", errors.New("execution reverted: AlreadyLive()")},
	} {
		t.Run(answer.name, func(t *testing.T) {
			c := qt.New(t)
			p := newTestPool("primary", "fallback")
			for i := range 3 * rpcMaxFailures {
				_ = p.Current()
				p.NoteError(errConnReset)
				_ = p.Current()
				c.Assert(p.CurrentURL(), qt.Equals, "primary", qt.Commentf("after scattered failure %d", i+1))
				p.NoteError(answer.err)
			}
			c.Assert(p.entries[0].disabledAt.IsZero(), qt.IsTrue, qt.Commentf("the primary must stay enabled"))

			// rpcMaxFailures failures in a row still disable it.
			for range rpcMaxFailures {
				_ = p.Current()
				p.NoteError(errConnReset)
			}
			c.Assert(p.entries[0].disabledAt.IsZero(), qt.IsFalse)
			_ = p.Current()
			c.Assert(p.CurrentURL(), qt.Equals, "fallback")
		})
	}
}

// The pool prefers its endpoints in the configured order: once the
// primary's cooldown has expired it takes the traffic back, even while the
// fallback it moved to keeps answering.
func TestPoolReturnsToThePrimaryAfterTheCooldown(t *testing.T) {
	for _, disable := range []struct {
		name string
		mark func(p *RPCPool)
	}{
		{"transport failures", func(p *RPCPool) {
			for range rpcMaxFailures {
				p.NoteError(errConnReset)
			}
		}},
		{"rate limit", func(p *RPCPool) { p.NoteError(errors.New("429 Too Many Requests")) }},
	} {
		t.Run(disable.name, func(t *testing.T) {
			c := qt.New(t)
			p := newTestPool("primary", "fallback", "last")
			_ = p.Current()
			disable.mark(p)
			_ = p.Current()
			c.Assert(p.CurrentURL(), qt.Equals, "fallback")

			// The fallback answers; the primary is still cooling down.
			p.NoteError(nil)
			_ = p.Current()
			c.Assert(p.CurrentURL(), qt.Equals, "fallback")

			p.entries[0].disabledAt = time.Now().Add(-rpcCooldownDuration)
			_ = p.Current()
			c.Assert(p.CurrentURL(), qt.Equals, "primary", qt.Commentf("a cooled-down primary must take the traffic back"))
			c.Assert(p.entries[0].failures, qt.Equals, 0)
		})
	}
}

// An endpoint that refuses a class of request (a free tier without archive
// access, a pruned node) is disabled for the cooldown at once and the pool
// rotates, like a rate limit: retrying it only buys the same refusal. The
// incident's spelling is publicnode's HTTP 403 with JSON-RPC -32602 "Archive
// requests require a personal token". A malformed request of ours (-32602
// invalid argument) and a revert are not the endpoint's fault.
func TestCapabilityRefusalsRotateAtOnce(t *testing.T) {
	archive := `{"code":-32602,"message":"Archive requests require a personal token. Get one at: https://example.com"}`
	cases := []struct {
		name   string
		err    error
		refuse bool
	}{
		{"http 403 archive", rpc.HTTPError{StatusCode: http.StatusForbidden, Status: "403 Forbidden", Body: []byte(archive)}, true},
		{"wrapped http 403", fmt.Errorf("filter logs: %w", rpc.HTTPError{StatusCode: http.StatusForbidden, Status: "403 Forbidden"}), true},
		{"http 401", rpc.HTTPError{StatusCode: http.StatusUnauthorized, Status: "401 Unauthorized"}, true},
		{"json-rpc archive", fakeRPCError{code: -32602, msg: "Archive requests require a personal token."}, true},
		{"missing trie node", errors.New("missing trie node 9ae5b4e3e1c0e2c5 (path ) state 0x1 is not available"), true},
		{"pruned history", errors.New("pruned history unavailable"), true},
		{"header not found", fakeRPCError{code: -32000, msg: "header not found"}, true},
		{"invalid argument", fakeRPCError{code: -32602, msg: "invalid argument 0: hex string without 0x prefix"}, false},
		{"revert", errors.New("execution reverted: NotEligible()"), false},
		{"not found", ethereum.NotFound, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := qt.New(t)
			p := newTestPool("a", "b")
			_ = p.Current()
			p.NoteError(tc.err)
			c.Assert(p.entries[0].disabledAt.IsZero(), qt.Equals, !tc.refuse)
			_ = p.Current()
			want := "a"
			if tc.refuse {
				want = "b"
			}
			c.Assert(p.CurrentURL(), qt.Equals, want)
		})
	}
}
