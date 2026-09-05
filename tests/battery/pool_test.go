package battery

import (
	"testing"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// TestPoolCounts pins the pool arithmetic every registration and swarm wave
// is sized with; it needs no fleet. Since finalizeEpoch stores every key of
// the pool at once, every unclaimed key is claimable.
func TestPoolCounts(t *testing.T) {
	cases := []struct {
		name      string
		next      uint8
		unclaimed int
	}{
		{"fresh", 0, ccommon.MaxK},
		{"three claimed", 3, ccommon.MaxK - 3},
		{"one left", ccommon.MaxK - 1, 1},
		{"exhausted", ccommon.MaxK, 0},
	}
	for _, c := range cases {
		unclaimed, ready := poolCounts(c.next)
		if unclaimed != c.unclaimed || ready != c.unclaimed {
			t.Errorf("%s: poolCounts(%d) = (%d, %d), want (%d, %d)",
				c.name, c.next, unclaimed, ready, c.unclaimed, c.unclaimed)
		}
	}
}
