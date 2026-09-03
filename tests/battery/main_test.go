package battery

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/vocdoni/davinci-dkg/log"
)

// fleet is the shared connection every scenario uses. It stays nil when the
// battery is not enabled, in which case every test skips.
var fleet *Fleet

func TestMain(m *testing.M) {
	if enabled() {
		log.Init(envOr("BATTERY_LOG_LEVEL", "warn"), "stdout", nil)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		f, err := connect(ctx)
		cancel()
		if err != nil {
			fmt.Fprintf(os.Stderr, "battery: cannot connect to the fleet: %v\n", err)
			os.Exit(1)
		}
		fleet = f
	} else {
		fmt.Println("battery: DAVINCI_DKG_BATTERY is not 1 — every test skips")
	}
	m.Run()
	if fleet != nil {
		if err := report.writeMarkdown(); err != nil {
			fmt.Fprintf(os.Stderr, "battery: write markdown summary: %v\n", err)
		} else {
			fmt.Printf("battery: report %s and %s\n", report.path, report.markdownPath())
		}
		fleet.close()
	}
}

// requireFleet skips the test unless the battery is enabled and connected.
func requireFleet(t *testing.T) *Fleet {
	t.Helper()
	if fleet == nil {
		t.Skip("set DAVINCI_DKG_BATTERY=1 (and DAVINCI_DKG_TEST_RPC_URL / DAVINCI_DKG_TEST_ADDRESSES) to run the battery")
	}
	return fleet
}

// testContext derives a context from the test deadline.
func testContext(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	timeout := 30 * time.Minute
	if deadline, ok := t.Deadline(); ok {
		timeout = time.Until(deadline) - 30*time.Second
	}
	return context.WithTimeout(context.Background(), timeout)
}
