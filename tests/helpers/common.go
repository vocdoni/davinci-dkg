package helpers

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
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

// RevertName returns the custom-error name a transaction (or its gas
// estimation) reverted with, looked up across the manager, app manager and
// registry ABIs — an app-manager error such as OrganizerSecretNotRevealed
// bubbles up through the manager's external call and is not in the manager's
// own ABI. An unknown selector comes back as "custom:0x…"; ok is false when
// the error carries no revert data at all.
func RevertName(err error) (name string, ok bool) {
	data, hasData := ethclient.RevertErrorData(err)
	if !hasData || len(data) < 4 {
		return "", false
	}
	var selector [4]byte
	copy(selector[:], data[:4])
	for _, md := range []*bind.MetaData{
		golangtypes.DKGManagerMetaData, golangtypes.DKGAppManagerMetaData, golangtypes.DKGRegistryMetaData,
	} {
		parsed, parseErr := md.GetAbi()
		if parseErr != nil {
			continue
		}
		if abiErr, lookupErr := parsed.ErrorByID(selector); lookupErr == nil {
			return abiErr.Name, true
		}
	}
	return "custom:0x" + hex.EncodeToString(selector[:]), true
}

// RevertsWith reports whether err is a revert with the named custom error,
// and otherwise describes what it was instead (for assertion messages).
func RevertsWith(err error, want string) (bool, string) {
	if err == nil {
		return false, "no error"
	}
	name, ok := RevertName(err)
	if !ok {
		return false, "not a revert: " + err.Error()
	}
	return name == want, name
}
