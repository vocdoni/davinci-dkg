package node

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"slices"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

func pointOf(k uint64) *big.Int {
	return new(big.Int).SetUint64(k)
}

func TestBSGSRoundTripsAcrossRange(t *testing.T) {
	const m = uint64(64) // covers values in [0, 4096)
	table := newBSGSTable(m)
	cases := []uint64{
		0, 1, 2, 63, 64, 65, // boundary around m
		100, 999, 4095, // last in-range
	}
	for _, want := range cases {
		t.Run(fmt.Sprintf("k=%d", want), func(t *testing.T) {
			target := group.NewPoint()
			target.ScalarBaseMult(pointOf(want))
			got, err := table.lookup(target)
			if err != nil {
				t.Fatalf("lookup failed: %v", err)
			}
			if got.Uint64() != want {
				t.Fatalf("got %d, want %d", got.Uint64(), want)
			}
		})
	}
}

func TestBSGSOutOfRangeErrors(t *testing.T) {
	const m = uint64(32)
	table := newBSGSTable(m)
	target := group.NewPoint()
	// m*m = 1024, so 1024 is the smallest out-of-range value.
	target.ScalarBaseMult(pointOf(m * m))
	if _, err := table.lookup(target); err == nil {
		t.Fatalf("expected out-of-range error, got nil")
	}
}

// The baby-step table is a sorted vector of 8-byte keys (truncated
// x-coordinate ‖ index) rather than a map of 32-byte encodings, so the
// production 2^25 table fits in ~256 MB. Every index must be present exactly
// once and recoverable from its key.
func TestBSGSTableIsSortedAndCompact(t *testing.T) {
	c := qt.New(t)
	const m = uint64(1000) // not a power of two: exercises the index-width rounding
	table := newBSGSTable(m)
	c.Assert(table.keys, qt.HasLen, int(m))
	c.Assert(slices.IsSorted(table.keys), qt.IsTrue)

	seen := make([]bool, m)
	for _, k := range table.keys {
		i := k & table.idxMask
		c.Assert(i < m, qt.IsTrue)
		c.Assert(seen[i], qt.IsFalse, qt.Commentf("index %d appears twice", i))
		seen[i] = true
	}
	// Random probes across the whole [0, m²) domain, including ones that hit
	// the −i·G branch, all resolve.
	rng := rand.New(rand.NewPCG(1, 2))
	for range 64 {
		want := rng.Uint64N(m * m)
		target := group.NewPoint()
		target.ScalarBaseMult(pointOf(want))
		got, err := table.lookup(target)
		c.Assert(err, qt.IsNil)
		c.Assert(got.Uint64(), qt.Equals, want)
	}
}

func BenchmarkBSGSTableBuild(b *testing.B) {
	const m = uint64(1) << 18
	for range b.N {
		newBSGSTable(m)
	}
}

func BenchmarkBSGSLookup(b *testing.B) {
	const m = uint64(1) << 18
	table := newBSGSTable(m)
	rng := rand.New(rand.NewPCG(3, 4))
	targets := make([]*big.Int, 32)
	for i := range targets {
		targets[i] = pointOf(rng.Uint64N(m * m))
	}
	b.ResetTimer()
	for i := range b.N {
		target := group.NewPoint()
		target.ScalarBaseMult(targets[i%len(targets)])
		if _, err := table.lookup(target); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDLogBSGSFull exercises the production table (m = 2^25, ~256 MB)
// with a worst-case-sized plaintext just below MaxDLogPlaintext.
func BenchmarkDLogBSGSFull(b *testing.B) {
	target := group.NewPoint()
	target.ScalarBaseMult(pointOf(MaxDLogPlaintext - 12345))
	b.ResetTimer()
	for range b.N {
		got, err := dlogBSGS(target)
		if err != nil {
			b.Fatal(err)
		}
		if got.Uint64() != MaxDLogPlaintext-12345 {
			b.Fatalf("got %d", got.Uint64())
		}
	}
}
