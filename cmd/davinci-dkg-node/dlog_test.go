package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-node/crypto/ecc"
)

// bsgsLookup is the same algorithm as dlogBSGS, but parameterised so we
// can exercise it with a tiny cap in unit tests instead of building the
// full 2^25-entry table for the real deployment. The production cap stays
// MaxDLogPlaintext / bsgsM and is verified by separate end-to-end tests.
func bsgsLookup(target ecc.Point, m uint64) (*big.Int, error) {
	gen := group.Generator()
	cur := group.NewPoint()
	cur.SetZero()
	table := make(map[string]uint32, m)
	for i := uint32(0); uint64(i) < m; i++ {
		table[string(cur.Marshal())] = i
		cur.Add(cur, gen)
	}

	negM := group.NewPoint()
	negM.ScalarBaseMult(new(big.Int).SetUint64(m))
	negM.Neg(negM)

	cur.Set(target)
	for j := uint64(0); j < m; j++ {
		if i, ok := table[string(cur.Marshal())]; ok {
			return new(big.Int).SetUint64(j*m + uint64(i)), nil
		}
		cur.Add(cur, negM)
	}
	return nil, fmt.Errorf("dlog: out of range (>= %d)", m*m)
}

func TestBSGSRoundTripsAcrossRange(t *testing.T) {
	const m = uint64(64) // covers values in [0, 4096)
	cases := []uint64{
		0, 1, 2, 63, 64, 65, // boundary around m
		100, 999, 4095, // last in-range
	}
	for _, want := range cases {
		t.Run(fmt.Sprintf("k=%d", want), func(t *testing.T) {
			target := group.NewPoint()
			target.ScalarBaseMult(new(big.Int).SetUint64(want))
			got, err := bsgsLookup(target, m)
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
	target := group.NewPoint()
	// m*m = 1024, so 1024 is the smallest out-of-range value.
	target.ScalarBaseMult(new(big.Int).SetUint64(m * m))
	if _, err := bsgsLookup(target, m); err == nil {
		t.Fatalf("expected out-of-range error, got nil")
	}
}
