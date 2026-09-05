package web3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/vocdoni/davinci-dkg/log"
)

// rpcHTTPTimeout bounds every HTTP round trip of a pooled endpoint so a hung
// endpoint fails over instead of blocking a tick forever.
const rpcHTTPTimeout = 30 * time.Second

const (
	rpcCooldownDuration = 2 * time.Minute
	rpcMaxFailures      = 3
)

// rpcRateLimitCode is the JSON-RPC error code most public providers answer
// with once a client exceeds its request quota (the -32000..-32099 server
// error band); -32005 is the conventional "rate limit" spelling.
const rpcRateLimitCode = -32005

// RPCPool holds multiple ethclient instances and provides round-robin failover.
type RPCPool struct {
	mu      sync.Mutex
	entries []*rpcEntry
	current int
}

type rpcEntry struct {
	url        string
	client     *ethclient.Client
	failures   int
	disabledAt time.Time // zero = enabled
}

// NewRPCPool dials all provided URLs. Returns error only if none succeed.
func NewRPCPool(urls []string) (*RPCPool, error) {
	if len(urls) == 0 {
		return nil, fmt.Errorf("rpc pool: no URLs provided")
	}
	pool := &RPCPool{}
	for _, url := range urls {
		// Dial with a bounded HTTP client: a hung endpoint then fails the
		// round trip after rpcHTTPTimeout instead of blocking forever. The
		// option only applies to http(s) URLs; websocket and ipc endpoints
		// keep their own keepalive/deadline handling.
		rpcClient, err := rpc.DialOptions(context.Background(), url, rpc.WithHTTPClient(&http.Client{Timeout: rpcHTTPTimeout}))
		if err != nil {
			log.Warnw("rpc pool: failed to dial endpoint, skipping", "url", url, "err", err)
			continue
		}
		client := ethclient.NewClient(rpcClient)
		pool.entries = append(pool.entries, &rpcEntry{url: url, client: client})
	}
	if len(pool.entries) == 0 {
		return nil, fmt.Errorf("rpc pool: all endpoints failed to dial")
	}
	return pool, nil
}

// Current returns the active client, re-enabling cooled-down entries first.
// If all entries are disabled, they are all reset.
func (p *RPCPool) Current() *ethclient.Client {
	p.mu.Lock()
	defer p.mu.Unlock()

	n := len(p.entries)
	// Try to find an enabled entry starting from current.
	for i := 0; i < n; i++ {
		idx := (p.current + i) % n
		e := p.entries[idx]
		if e.disabledAt.IsZero() {
			// Enabled entry found.
			p.current = idx
			return e.client
		}
		// Check if cooldown has expired.
		if time.Since(e.disabledAt) >= rpcCooldownDuration {
			e.disabledAt = time.Time{}
			e.failures = 0
			log.Infow("rpc pool: re-enabled endpoint", "url", e.url)
			p.current = idx
			return e.client
		}
	}

	// All entries are disabled — reset all of them.
	log.Warnw("rpc pool: all endpoints failed, resetting")
	for _, e := range p.entries {
		e.disabledAt = time.Time{}
		e.failures = 0
	}
	return p.entries[p.current].client
}

// CurrentURL returns the URL of the active client (for logging).
func (p *RPCPool) CurrentURL() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.entries[p.current].url
}

// MarkFailed marks the current endpoint as having failed. After rpcMaxFailures
// it is disabled for rpcCooldownDuration and the pool rotates to the next one.
func (p *RPCPool) MarkFailed() {
	p.mu.Lock()
	defer p.mu.Unlock()

	e := p.entries[p.current]
	e.failures++
	if e.failures >= rpcMaxFailures {
		e.disabledAt = time.Now()
		log.Warnw("rpc pool: disabling endpoint", "url", e.url, "failures", e.failures)
		p.current = (p.current + 1) % len(p.entries)
	}
}

// MarkRateLimited disables the current endpoint for the cooldown at once and
// rotates: a 429 means the provider is already refusing our requests, so the
// rpcMaxFailures threshold would only buy two more rate-limited round trips.
func (p *RPCPool) MarkRateLimited() {
	p.mu.Lock()
	defer p.mu.Unlock()

	e := p.entries[p.current]
	e.failures++
	e.disabledAt = time.Now()
	log.Warnw("rpc pool: rate-limited, disabling endpoint", "url", e.url, "failures", e.failures)
	p.current = (p.current + 1) % len(p.entries)
}

// NoteError classifies an error returned by a call made to the current
// endpoint and rotates the pool when the endpoint is at fault: a rate-limit
// response disables it at once (MarkRateLimited), a connectivity failure
// counts toward the usual threshold (MarkFailed). Contract reverts and every
// other error are none of the endpoint's business. nil is a no-op, so callers
// can feed it every result unconditionally.
// Rotate advances to the next endpoint without disabling the current one;
// used by startup code that wants to try every endpoint once.
func (p *RPCPool) Rotate() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.current = (p.current + 1) % len(p.entries)
}

func (p *RPCPool) NoteError(err error) {
	if err == nil || !IsRPCTransportError(err) {
		return
	}
	if IsRateLimitError(err) {
		p.MarkRateLimited()
		return
	}
	p.MarkFailed()
}

// Close closes all underlying clients.
func (p *RPCPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, e := range p.entries {
		e.client.Close()
	}
}

// IsRPCTransportError returns true if the error is a connectivity/transport
// problem or a provider rate-limit response (not a contract revert). Such
// errors should trigger pool rotation.
func IsRPCTransportError(err error) bool {
	if err == nil {
		return false
	}
	if IsRateLimitError(err) {
		return true
	}
	s := err.Error()
	for _, substr := range []string{
		"connection refused",
		"EOF",
		"dial tcp",
		"i/o timeout",
		"no such host",
		"context deadline exceeded",
		"read tcp",
		"write tcp",
		"unexpected EOF",
		"use of closed network connection",
	} {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// IsRateLimitError reports whether err is an RPC provider throttling us: a
// JSON-RPC error code -32005 on any rpc.Error in the chain, or an HTTP 429 /
// "too many requests" / "rate limit" message (case-insensitive). A rate-limited
// endpoint is not broken, but retrying it immediately only buys more of the
// same response, so callers rotate to another endpoint right away.
func IsRateLimitError(err error) bool {
	if err == nil {
		return false
	}
	var rpcErr rpc.Error
	if errors.As(err, &rpcErr) && rpcErr.ErrorCode() == rpcRateLimitCode {
		return true
	}
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "429") ||
		strings.Contains(s, "too many requests") ||
		strings.Contains(s, "rate limit")
}
