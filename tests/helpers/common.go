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
