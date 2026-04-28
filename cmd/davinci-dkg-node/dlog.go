package main

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
// equal — ~33.5 M point additions and a ~2 GB hash table on the heap.
//
// The table is built lazily on the first call (so a node that never combines
// pays nothing) and cached for the lifetime of the process. After that, every
// subsequent decryption costs at most m giant-step iterations (~30–60 s wall
// clock on a modern CPU; faster in practice because most plaintexts land
// well before the worst case).
//
// Memory note for operators: the precomputed table needs roughly
// (32 + 4) bytes per entry × 33.5 M entries ≈ 1.2 GB raw, which works out
// to ~2 GB on the Go heap once map overhead is counted. A node operator
// who never expects to combine won't pay this — initBSGSTable is only run
// the first time `dlogBSGS` is called. To raise the cap beyond 2^50 we'd
// need a different algorithm (e.g. Pollard's kangaroo) — see README.

import (
	"fmt"
	"math/big"
	"sync"
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
	bsgsTable map[string]uint32 // 32-byte point Marshal() → baby-step index
	bsgsNegM  ecc.Point         // -m·G, the giant-step direction
)

// initBSGSTable runs the one-time precomputation. Marked package-level so
// tests and benchmarks can drive it explicitly via dlogBSGS, but it's only
// ever called via the sync.Once below.
func initBSGSTable() {
	start := time.Now()
	log.Infow("dlog: building BSGS table", "m", bsgsM, "max", MaxDLogPlaintext)

	gen := group.Generator()
	cur := group.NewPoint()
	cur.SetZero()
	bsgsTable = make(map[string]uint32, bsgsM)
	for i := uint32(0); uint64(i) < bsgsM; i++ {
		// Marshal() output is the canonical encoding for the curve, so
		// equal points always produce equal byte strings. string([]byte)
		// in Go performs a copy — we need that, otherwise the next loop
		// iteration would mutate the key in place.
		bsgsTable[string(cur.Marshal())] = i
		cur.Add(cur, gen)
	}

	// Giant step: M = m·G, then negate so we can advance by repeated
	// addition rather than subtraction inside the hot loop.
	bsgsNegM = group.NewPoint()
	bsgsNegM.ScalarBaseMult(new(big.Int).SetUint64(bsgsM))
	bsgsNegM.Neg(bsgsNegM)

	log.Infow("dlog: BSGS table ready", "elapsed", time.Since(start).String(), "entries", len(bsgsTable))
}

// dlogBSGS recovers a scalar `a` such that a·G = target with 0 ≤ a <
// MaxDLogPlaintext. Returns an error when the search exhausts the
// configured range — that's a hard signal that the round produced a
// plaintext outside the documented domain and the result is unrecoverable
// without a different algorithm.
//
// First call lazily builds the precomputed table (~30–60 s on a modern
// CPU, ~2 GB heap). Subsequent calls reuse it and run at most m
// giant-step iterations (~30–60 s worst case; usually much faster since
// most plaintexts land well before the upper bound).
func dlogBSGS(target ecc.Point) (*big.Int, error) {
	bsgsOnce.Do(initBSGSTable)

	// Walk T_j = target − j·M for j = 0, 1, …, looking each up in the
	// baby-step table. Found at index i means a = j·m + i.
	cur := group.NewPoint()
	cur.Set(target)
	for j := uint64(0); j < bsgsM; j++ {
		if i, ok := bsgsTable[string(cur.Marshal())]; ok {
			return new(big.Int).SetUint64(j*bsgsM + uint64(i)), nil
		}
		cur.Add(cur, bsgsNegM)
	}
	return nil, fmt.Errorf("dlog: plaintext out of range (>= 2^50 ≈ 10^15)")
}
