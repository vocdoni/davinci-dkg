package helpers

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func IsIntegrationEnabled() bool {
	return os.Getenv("RUN_INTEGRATION_TESTS") == "true"
}

// IsBenchmarkEnabled reports whether gas-profile / multi-size benchmark tests
// should run. Benchmarks are always skipped unless RUN_BENCHMARKS=true, even
// when integration tests are enabled — they are the data source for
// BENCHMARKS.md and are far too slow to run on every PR (a full MaxN=32 sweep
// takes upwards of 15 minutes of proving alone).
func IsBenchmarkEnabled() bool {
	return os.Getenv("RUN_BENCHMARKS") == "true"
}

func MaxTestTimeout(t *testing.T) time.Duration {
	t.Helper()

	if deadline, ok := t.Deadline(); ok {
		remaining := time.Until(deadline)
		buffer := 15 * time.Second
		if remaining <= buffer {
			buffer = remaining / 2
		}
		return remaining - buffer
	}
	return 10 * time.Minute
}

func WaitUntilCondition(ctx context.Context, interval time.Duration, condition func() bool) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for condition")
		case <-ticker.C:
			if condition() {
				return nil
			}
		}
	}
}
