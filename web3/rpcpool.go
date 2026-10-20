package web3

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
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

// RPCPool holds the node's JSON-RPC endpoints in priority order, the order
// they were configured in, and fails over between them. Every call goes to
// the earliest usable endpoint, so the primary takes the traffic back as
// soon as its cooldown expires. An endpoint is disabled for
// rpcCooldownDuration after rpcMaxFailures consecutive transport failures,
// or at once when it rate-limits us or refuses a class of request it does
// not serve (IsCapabilityError).
type RPCPool struct {
	mu      sync.Mutex
	entries []*rpcEntry
	current int // the endpoint picked last; NoteError and CurrentURL refer to it
	stats   rpcStats
	redact  redactor // every configured endpoint's secrets, for error text
}

// rpcStats counts the JSON-RPC requests sent through the pool, per method,
// so a long-running node can tell what one poll cycle costs its provider
// (the node logs the per-tick histogram; see Snapshot). Requests are counted
// where they are issued, not on the wire: a batch counts once as "batch"
// plus once per method it carries.
type rpcStats struct {
	mu     sync.Mutex
	counts map[string]uint64
}

func (s *rpcStats) note(method string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.counts == nil {
		s.counts = make(map[string]uint64)
	}
	s.counts[method]++
}

func (s *rpcStats) snapshot() map[string]uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]uint64, len(s.counts))
	for method, count := range s.counts {
		out[method] = count
	}
	return out
}

// Note records one JSON-RPC request of `method` (e.g. "eth_call") issued
// through the pool. PooledBackend and the Contracts readers call it for
// every request they send; callers that reach an endpoint some other way
// can keep the counters honest by calling it themselves.
func (p *RPCPool) Note(method string) { p.stats.note(method) }

// Snapshot returns a copy of the per-method request counts since the pool
// was created. Two snapshots some ticks apart give the cost of a tick.
func (p *RPCPool) Snapshot() map[string]uint64 { return p.stats.snapshot() }

type rpcEntry struct {
	url        string // the full URL: may carry an API key, never log it
	label      string // what is logged: "#<index> scheme://host"
	client     *ethclient.Client
	failures   int
	disabledAt time.Time // zero = enabled
}

// NewRPCPool dials all provided URLs. Returns error only if none succeed.
func NewRPCPool(urls []string) (*RPCPool, error) {
	if len(urls) == 0 {
		return nil, fmt.Errorf("rpc pool: no URLs provided")
	}
	pool := &RPCPool{redact: newRedactor(urls...)}
	for i, url := range urls {
		label := fmt.Sprintf("#%d %s", i, RedactURL(url))
		client, err := dialEndpoint(url)
		if err != nil {
			log.Warnw("rpc pool: failed to dial endpoint, skipping", "endpoint", label, "err", err)
			continue
		}
		pool.entries = append(pool.entries, &rpcEntry{url: url, label: label, client: client})
	}
	if len(pool.entries) == 0 {
		return nil, fmt.Errorf("rpc pool: all endpoints failed to dial")
	}
	return pool, nil
}

// Current returns the earliest usable endpoint in priority order,
// re-enabling entries whose cooldown has expired. If every entry is
// disabled, all of them are reset and the first one is returned.
func (p *RPCPool) Current() *ethclient.Client {
	_, client, _ := p.next(nil)
	return client
}

// next returns the earliest usable endpoint not yet marked in tried (indexed
// like the entries; nil when nothing was tried) and makes it current. ok is
// false once every usable endpoint has been tried. When no endpoint is
// usable at all and none was tried, every entry is reset and the first one
// is returned.
func (p *RPCPool) next(tried []bool) (int, *ethclient.Client, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	idx, ok := p.pickLocked(tried)
	if !ok {
		return 0, nil, false
	}
	return idx, p.entries[idx].client, true
}

// pickLocked is next without the client, for callers holding p.mu.
func (p *RPCPool) pickLocked(tried []bool) (int, bool) {
	for idx, e := range p.entries {
		if (tried != nil && tried[idx]) || !e.usable() {
			continue
		}
		p.current = idx
		return idx, true
	}
	if slices.Contains(tried, true) {
		return 0, false
	}
	log.Warnw("rpc pool: all endpoints failed, resetting")
	for _, e := range p.entries {
		e.disabledAt, e.failures = time.Time{}, 0
	}
	p.current = 0
	return 0, true
}

// usable reports whether the entry may serve a call, re-enabling it once its
// cooldown has expired.
func (e *rpcEntry) usable() bool {
	if e.disabledAt.IsZero() {
		return true
	}
	if time.Since(e.disabledAt) < rpcCooldownDuration {
		return false
	}
	e.disabledAt, e.failures = time.Time{}, 0
	log.Infow("rpc pool: re-enabled endpoint", "endpoint", e.label)
	return true
}

// CurrentURL returns the full URL of the active endpoint. It may carry an
// API key: log RedactURL(CurrentURL()) instead.
func (p *RPCPool) CurrentURL() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.entries[p.current].url
}

// MarkFailed counts a transport failure against the current endpoint. After
// rpcMaxFailures in a row it is disabled for rpcCooldownDuration and the
// pool moves to the next usable one.
func (p *RPCPool) MarkFailed() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.failedLocked(p.current, "unreachable")
}

// MarkRateLimited disables the current endpoint for the cooldown at once and
// moves on: a 429 means the provider is already refusing our requests, so the
// rpcMaxFailures threshold would only buy two more rate-limited round trips.
func (p *RPCPool) MarkRateLimited() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.disableLocked(p.current, "rate-limited")
}

// NoteError classifies the outcome of a call the current endpoint served
// (see noteAt). nil counts as an answer, so callers can feed it every result
// unconditionally.
func (p *RPCPool) NoteError(err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.noteLocked(p.current, err)
}

// disable takes the endpoint at idx out of service for the cooldown.
func (p *RPCPool) disable(idx int, reason string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.disableLocked(idx, reason)
}

// noteAt is NoteError for the endpoint at idx, whichever is current now.
func (p *RPCPool) noteAt(idx int, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.noteLocked(idx, err)
}

// noteLocked records one call's outcome against entry idx. Any answer,
// including an error that is the call's own (a revert, not found), proves
// the endpoint works and clears its failure count, so only consecutive
// failures disable it. A rate limit or a capability refusal disables it at
// once; a transport failure or an HTTP server error counts toward
// rpcMaxFailures.
func (p *RPCPool) noteLocked(idx int, err error) {
	switch {
	case IsRateLimitError(err):
		p.disableLocked(idx, "rate-limited")
	case IsCapabilityError(err):
		p.disableLocked(idx, "refused the request")
	case isRedirect(err):
		p.disableLocked(idx, "redirected (not followed: list the endpoint's canonical URL)")
	case IsRPCTransportError(err):
		p.failedLocked(idx, "unreachable")
	case isServerError(err):
		p.failedLocked(idx, "server error")
	default:
		p.entries[idx].failures = 0
	}
}

func (p *RPCPool) failedLocked(idx int, reason string) {
	e := p.entries[idx]
	e.failures++
	if e.failures >= rpcMaxFailures {
		p.disableLocked(idx, reason)
	}
}

// disableLocked takes entry idx out of service for the cooldown and makes
// the next usable endpoint current.
func (p *RPCPool) disableLocked(idx int, reason string) {
	e := p.entries[idx]
	if e.disabledAt.IsZero() {
		log.Warnw("rpc pool: disabling endpoint", "endpoint", e.label, "reason", reason, "failures", e.failures)
	}
	e.disabledAt = time.Now()
	p.pickLocked(nil)
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

// IsCapabilityError reports whether err is an endpoint refusing a class of
// request it does not serve rather than failing it: HTTP 401 / 403 (a free
// tier answering "Archive requests require a personal token" for old log
// ranges and for every receipt), or a pruned node missing the state, header
// or history asked for. Another endpoint may well answer, so the pool
// rotates at once, as for a rate limit. A revert is never one.
func IsCapabilityError(err error) bool {
	if err == nil {
		return false
	}
	var httpErr rpc.HTTPError
	if errors.As(err, &httpErr) &&
		(httpErr.StatusCode == http.StatusForbidden || httpErr.StatusCode == http.StatusUnauthorized) {
		return true
	}
	s := strings.ToLower(err.Error())
	if strings.Contains(s, "execution reverted") {
		return false
	}
	for _, substr := range []string{
		"403 forbidden",
		"401 unauthorized",
		"archive",
		"missing trie node",
		"pruned",
		"header not found",
	} {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// endpointFault reports whether err is the endpoint's fault rather than the
// call's own outcome: a transport failure, a rate limit, a capability
// refusal, an HTTP 5xx or a redirect. Such a call is worth
// repeating on another endpoint.
func endpointFault(err error) bool {
	return IsRPCTransportError(err) || IsCapabilityError(err) || isServerError(err) || isRedirect(err)
}

// isRedirect reports whether the endpoint answered with an HTTP redirect.
// The node never follows one (see dialEndpoint): the credentials in the
// endpoint's URL must not reach another, so the endpoint is unusable until
// its URL is fixed.
func isRedirect(err error) bool {
	var httpErr rpc.HTTPError
	return errors.As(err, &httpErr) && httpErr.StatusCode >= 300 && httpErr.StatusCode < 400
}

// isServerError reports whether err is an HTTP 5xx from the endpoint (or a
// gateway in front of it, 52x included): the endpoint failed the request,
// which another endpoint can serve. The status comes from go-ethereum's
// structured rpc.HTTPError, never from message text.
func isServerError(err error) bool {
	var httpErr rpc.HTTPError
	return errors.As(err, &httpErr) && httpErr.StatusCode >= 500 && httpErr.StatusCode <= 599
}
