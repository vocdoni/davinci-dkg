package battery

import "testing"

// TestPoolCounts pins the pool arithmetic every registration and swarm wave
// is sized with; it needs no fleet.
func TestPoolCounts(t *testing.T) {
	cases := []struct {
		name            string
		next, activated uint8
		unclaimed       int
		ready           int
	}{
		{"fresh, nothing activated", 0, 0b00000000, 8, 0},
		{"fresh, two ahead", 0, 0b00000011, 8, 2},
		{"three claimed, two ahead", 3, 0b00011111, 5, 2},
		{"gap after the cursor", 3, 0b00010111, 5, 0},
		{"gap past the ready run", 2, 0b10001111, 6, 2},
		{"whole pool activated", 5, 0b11111111, 3, 3},
		{"exhausted", 8, 0b11111111, 0, 0},
	}
	for _, c := range cases {
		unclaimed, ready := poolCounts(c.next, c.activated)
		if unclaimed != c.unclaimed || ready != c.ready {
			t.Errorf("%s: poolCounts(%d, %08b) = (%d, %d), want (%d, %d)",
				c.name, c.next, c.activated, unclaimed, ready, c.unclaimed, c.ready)
		}
	}
}
