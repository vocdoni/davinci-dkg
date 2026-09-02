package node

// Discrete-log recovery for the final step of threshold ElGamal decryption.
//
// After a committee combines partial decryptions, we are left with a point
// `m·G` and need to recover the scalar `m`. There is no efficient general
// algorithm for the discrete log on BabyJubJub, so we lean on the fact that
// `m` is bounded by the protocol contract: it can never exceed
// `MaxDLogPlaintext`, currently 2^50.
//
// We use baby-step / giant-step (BSGS):
//
//   1. Precompute a table of m baby steps  i·G  for i ∈ [0, m).
//   2. Compute the giant step  M = m·G  and its inverse  -M.
//   3. To recover the unknown scalar a (with 0 ≤ a < m²), iterate
//      T_j = target − j·M for j = 0, 1, … and look each T_j up in the
//      table. A hit at table index i means a = j·m + i.
//
// Total work is at most 2m point additions and uses O(m) memory for the
// table. We pick m = ⌈√MaxDLogPlaintext⌉ = 2^25, which makes both costs
// equal: ~33.5 M point additions for the build and, worst case, for a
// lookup; both are split across every CPU.
//
// The table stores one 8-byte key per baby step instead of the 32-byte
// point encoding: the high (64 − ⌈log2 m⌉) bits are the low bits of the
// point's y-coordinate and the low bits are the index i. Keys are kept in a
// sorted vector, so a lookup is a binary search on the y prefix followed by
// a scalar-multiplication check of each candidate (a false candidate is a
// 2^-14 event per giant step at m = 2^25). On a twisted Edwards curve P and
// −P share y, so a hit at T_j = −i·G is useful too (a = j·m − i), which
// halves the expected number of giant steps. Memory is m × 8 B = 256 MB at
// m = 2^25.
//
// The table is built lazily on the first call (so a node that never combines
// pays nothing) and cached for the lifetime of the process.

import (
	"fmt"
	"math/big"
	"math/bits"
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-node/crypto/ecc"
)

// MaxDLogPlaintext is the largest plaintext (exclusive) that combineRound
// can recover from a finalized ciphertext. The contract does not enforce
// this — callers must make sure their plaintexts stay strictly below.
//
// 2^50 ≈ 1.13 × 10^15 — enough for any realistic vote tally (the entire
// human population voting weighted up to a million each still uses < 2^45).
// The cap is set by what BSGS can comfortably afford in memory; raising
// it requires switching to a different algorithm (Pollard's kangaroo).
const MaxDLogPlaintext = uint64(1) << 50

// bsgsM = ⌈√MaxDLogPlaintext⌉. With MaxDLogPlaintext = 2^50 the optimal
// m is exactly 2^25.
const bsgsM = uint64(1) << 25

var (
	bsgsOnce  sync.Once
	bsgsCache *bsgsTable
)

// bsgsTable is the precomputed baby-step table for a given m.
type bsgsTable struct {
	m       uint64
	idxMask uint64   // low bits of a key hold the baby-step index
	keys    []uint64 // sorted: (yLow64 &^ idxMask) | i
	negM    ecc.Point
}

// newBSGSTable builds the baby-step table for m in parallel across CPUs.
func newBSGSTable(m uint64) *bsgsTable {
	if m == 0 {
		panic("bsgs: m must be positive")
	}
	start := time.Now()
	t := &bsgsTable{m: m, keys: make([]uint64, m)}
	if m > 1 {
		t.idxMask = uint64(1)<<bits.Len64(m-1) - 1
	}

	workers := uint64(runtime.NumCPU())
	workers = max(1, min(workers, m))
	chunk := (m + workers - 1) / workers
	var wg sync.WaitGroup
	for lo := uint64(0); lo < m; lo += chunk {
		hi := min(lo+chunk, m)
		wg.Add(1)
		go func(lo, hi uint64) {
			defer wg.Done()
			gen := group.Generator()
			cur := group.NewPoint()
			cur.ScalarBaseMult(new(big.Int).SetUint64(lo))
			for i := lo; i < hi; i++ {
				t.keys[i] = t.key(cur, i)
				cur.Add(cur, gen)
			}
		}(lo, hi)
	}
	wg.Wait()
	slices.Sort(t.keys)

	// Giant step: M = m·G, then negate so we can advance by repeated
	// addition rather than subtraction inside the hot loop.
	t.negM = group.NewPoint()
	t.negM.ScalarBaseMult(new(big.Int).SetUint64(m))
	t.negM.Neg(t.negM)

	if m >= 1<<20 {
		log.Infow("dlog: BSGS table ready", "m", m, "elapsed", time.Since(start).String(),
			"bytes", len(t.keys)*8, "workers", workers)
	}
	return t
}

// key packs the low 64 bits of p's y-coordinate (above the index bits)
// together with the baby-step index.
func (t *bsgsTable) key(p ecc.Point, i uint64) uint64 {
	return t.prefix(p) | i
}

// prefix returns the y-coordinate part of a key with the index bits cleared.
func (t *bsgsTable) prefix(p ecc.Point) uint64 {
	_, y := p.Point()
	words := y.Bits()
	if len(words) == 0 {
		return 0
	}
	return uint64(words[0]) &^ t.idxMask
}

// lookup recovers a such that a·G = target with 0 ≤ a < m². The giant-step
// walk is split across CPUs; the first verified hit wins.
func (t *bsgsTable) lookup(target ecc.Point) (*big.Int, error) {
	workers := uint64(runtime.NumCPU())
	workers = max(1, min(workers, t.m))
	chunk := (t.m + workers - 1) / workers

	var (
		found  atomic.Bool
		mu     sync.Mutex
		result uint64
		wg     sync.WaitGroup
	)
	for j0 := uint64(0); j0 < t.m; j0 += chunk {
		j1 := min(j0+chunk, t.m)
		wg.Add(1)
		go func(j0, j1 uint64) {
			defer wg.Done()
			// cur = target − j0·M
			cur := group.NewPoint()
			cur.ScalarMult(t.negM, new(big.Int).SetUint64(j0))
			cur.Add(cur, target)
			for j := j0; j < j1 && !found.Load(); j++ {
				if a, ok := t.candidates(cur, target, j); ok {
					mu.Lock()
					if !found.Swap(true) {
						result = a
					}
					mu.Unlock()
					return
				}
				cur.Add(cur, t.negM)
			}
		}(j0, j1)
	}
	wg.Wait()
	if !found.Load() {
		return nil, fmt.Errorf("dlog: plaintext out of range (>= %d)", t.m*t.m)
	}
	return new(big.Int).SetUint64(result), nil
}

// candidates checks every baby step whose y prefix matches cur = target −
// j·M. A match at index i means cur = ±i·G, i.e. a = j·m ± i; each
// candidate is verified by a scalar multiplication before being accepted.
func (t *bsgsTable) candidates(cur, target ecc.Point, j uint64) (uint64, bool) {
	prefix := t.prefix(cur)
	pos, _ := slices.BinarySearch(t.keys, prefix)
	check := group.NewPoint()
	for ; pos < len(t.keys) && t.keys[pos]&^t.idxMask == prefix; pos++ {
		i := t.keys[pos] & t.idxMask
		a := j*t.m + i
		check.ScalarBaseMult(new(big.Int).SetUint64(a))
		if check.Equal(target) {
			return a, true
		}
		if i == 0 || j*t.m < i {
			continue
		}
		a = j*t.m - i
		check.ScalarBaseMult(new(big.Int).SetUint64(a))
		if check.Equal(target) {
			return a, true
		}
	}
	return 0, false
}

// dlogBSGS recovers a scalar `a` such that a·G = target with 0 ≤ a <
// MaxDLogPlaintext. Returns an error when the search exhausts the
// configured range — that's a hard signal that the epoch produced a
// plaintext outside the documented domain and the result is unrecoverable
// without a different algorithm.
//
// First call lazily builds the precomputed table (2^25 baby steps spread
// over every CPU, ~256 MB). Subsequent calls reuse it.
func dlogBSGS(target ecc.Point) (*big.Int, error) {
	bsgsOnce.Do(func() {
		log.Infow("dlog: building BSGS table", "m", bsgsM, "max", MaxDLogPlaintext)
		bsgsCache = newBSGSTable(bsgsM)
	})
	return bsgsCache.lookup(target)
}
