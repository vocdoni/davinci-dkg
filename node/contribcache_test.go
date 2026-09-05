package node

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

// The cache stores plain files under
// <datadir>/contributions/<epochId hex>/<dealer address>.bin and round-trips
// them; dealers and epochs never share entries.
func TestContributionCacheRoundTripsUnderTheDatadir(t *testing.T) {
	c := qt.New(t)
	dir := t.TempDir()
	cache := &contributionCache{dir: contributionCacheDir(dir)}
	epoch := [12]byte{1, 2, 3}
	dealer := common.HexToAddress("0x1234567890AbCdEf1234567890aBcDeF12345678")

	_, ok := cache.Get(epoch, dealer)
	c.Assert(ok, qt.IsFalse, qt.Commentf("an unseen contribution must be a miss"))

	data := []byte("raw submitContribution calldata")
	cache.Put(epoch, dealer, data)

	got, ok := cache.Get(epoch, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(string(got), qt.Equals, string(data))

	// The file must live at the documented path so operators can inspect it.
	want := filepath.Join(dir, "contributions",
		"010203000000000000000000", strings.ToLower(dealer.Hex())+".bin")
	_, err := os.Stat(want)
	c.Assert(err, qt.IsNil, qt.Commentf("cached calldata not at %s", want))

	// Other dealers and epochs are separate entries.
	_, ok = cache.Get(epoch, common.HexToAddress("0xdead"))
	c.Assert(ok, qt.IsFalse)
	_, ok = cache.Get([12]byte{9}, dealer)
	c.Assert(ok, qt.IsFalse)

	// A rewritten entry keeps a single canonical file (atomic replace).
	cache.Put(epoch, dealer, []byte("new calldata"))
	got, ok = cache.Get(epoch, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(string(got), qt.Equals, "new calldata")
}

// A nil cache (node built without one, e.g. in tests) and a cache without a
// datadir are disabled — Get misses, Put is a no-op — and read problems are
// misses that make the caller fall back to the RPC, never errors or panics.
func TestContributionCacheDisabledAndUnreadableEntriesAreMisses(t *testing.T) {
	c := qt.New(t)

	var nilCache *contributionCache
	_, ok := nilCache.Get([12]byte{1}, common.HexToAddress("0x1"))
	c.Assert(ok, qt.IsFalse)
	nilCache.Put([12]byte{1}, common.HexToAddress("0x1"), []byte("x"))

	empty := &contributionCache{}
	empty.Put([12]byte{1}, common.HexToAddress("0x1"), []byte("x"))
	_, ok = empty.Get([12]byte{1}, common.HexToAddress("0x1"))
	c.Assert(ok, qt.IsFalse)

	// A directory that does not exist yet (or cannot be read) is a miss.
	missing := &contributionCache{dir: filepath.Join(t.TempDir(), "contributions")}
	_, ok = missing.Get([12]byte{1}, common.HexToAddress("0x1"))
	c.Assert(ok, qt.IsFalse)

	// An empty datadir disables the cache entirely.
	c.Assert(contributionCacheDir(""), qt.Equals, "")
	c.Assert(contributionCacheDir("/data"), qt.Equals, filepath.Join("/data", "contributions"))
}
