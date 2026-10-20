package node

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/log"
)

// A node re-reads every accepted dealer's submitContribution calldata whenever
// it rebuilds a private share or a finalization statement (the encrypted shares
// and commitment points only live in that transaction). Under RPC rate
// limiting each miss costs a full event-log rescan of the epoch's age, so
// validated calldata is remembered on disk, next to the taint file, as plain
// bytes under <datadir>/<chainid>-<manager>/contributions/<epochId hex>/
// <dealer address>.bin (see datadir.go). A contribution is immutable once
// accepted (the contract refuses a second one), so entries never go stale.
// A committee member fills it for every accepted dealer when the epoch goes
// Live (keepContributions), while the logs are recent: a share for a pool key
// claimed days later then needs no old-range eth_getLogs, which many
// endpoints refuse.

// contributionCache memoises raw submitContribution calldata under the node
// datadir. It implements finalizer.CalldataCache. A nil receiver or an empty
// dir disables it: Get is always a miss and Put does nothing.
type contributionCache struct {
	dir string // <datadir>/<chainid>-<manager>/contributions, "" disables
}

// contributionCacheDir returns the cache directory of a datadir ("" for no
// datadir, which disables the cache).
func contributionCacheDir(datadir string) string {
	if datadir == "" {
		return ""
	}
	return filepath.Join(datadir, "contributions")
}

func (c *contributionCache) enabled() bool { return c != nil && c.dir != "" }

// path of one dealer's cached calldata. The dealer address is lowercased so
// the file name has a single canonical spelling across filesystems.
func (c *contributionCache) path(epochID [12]byte, dealer common.Address) string {
	return filepath.Join(c.dir, hex.EncodeToString(epochID[:]), strings.ToLower(dealer.Hex())+".bin")
}

// Get returns the cached calldata of dealer's contribution to the epoch. Any
// read problem (disabled cache, missing or unreadable file) is a miss: the
// caller falls back to the RPC.
func (c *contributionCache) Get(epochID [12]byte, dealer common.Address) ([]byte, bool) {
	if !c.enabled() {
		return nil, false
	}
	data, err := os.ReadFile(c.path(epochID, dealer))
	if err != nil {
		return nil, false
	}
	return data, true
}

// Put stores calldata written after it has been validated (selector, ABI
// shape, transcript size), best effort: a failed write is logged and only
// costs a later refetch. Retention, which must know, uses Store.
func (c *contributionCache) Put(epochID [12]byte, dealer common.Address, data []byte) {
	if err := c.Store(epochID, dealer, data); err != nil {
		log.Warnw("contribution cache: cannot persist calldata", "err", err)
	}
}

// Store is Put returning the write error. The write is atomic (temp file +
// rename) so a crash never leaves a truncated entry behind. A disabled cache
// stores nothing and reports no error.
func (c *contributionCache) Store(epochID [12]byte, dealer common.Address, data []byte) error {
	if !c.enabled() {
		return nil
	}
	p := c.path(epochID, dealer)
	tmp := p + ".tmp"
	err := os.MkdirAll(filepath.Dir(p), 0o700)
	if err == nil {
		err = os.WriteFile(tmp, data, 0o600)
	}
	if err == nil {
		err = os.Rename(tmp, p)
	}
	if err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("persist %s: %w", p, err)
	}
	return nil
}
