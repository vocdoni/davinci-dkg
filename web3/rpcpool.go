package web3

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/davinci-dkg/log"
)

const (
	rpcCooldownDuration = 2 * time.Minute
	rpcMaxFailures      = 3
)

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
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Warnw("rpc pool: failed to dial endpoint, skipping", "url", url, "err", err)
			continue
		}
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

// Close closes all underlying clients.
func (p *RPCPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, e := range p.entries {
		e.client.Close()
	}
}

// IsRPCTransportError returns true if the error is a connectivity/transport
// problem (not a contract revert). Such errors should trigger pool rotation.
func IsRPCTransportError(err error) bool {
	if err == nil {
		return false
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
