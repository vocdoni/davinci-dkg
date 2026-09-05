package main

import (
	"testing"
	"time"
)

func TestParseDeadline(t *testing.T) {
	now := time.Date(2026, 9, 4, 8, 0, 0, 0, time.UTC)
	got, err := parseDeadline("48h", now)
	if err != nil || got != uint64(now.Add(48*time.Hour).Unix()) {
		t.Fatalf("duration: got %d, %v", got, err)
	}
	got, err = parseDeadline("2026-09-10T12:00:00Z", now)
	if err != nil || got != uint64(time.Date(2026, 9, 10, 12, 0, 0, 0, time.UTC).Unix()) {
		t.Fatalf("timestamp: got %d, %v", got, err)
	}
	for _, bad := range []string{"", "yesterday", "-1h", "2026-09-04T08:00:00Z"} {
		if _, err := parseDeadline(bad, now); err == nil {
			t.Fatalf("%q must be rejected", bad)
		}
	}
}
