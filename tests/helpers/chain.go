package helpers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func MineBlocks(ctx context.Context, services *TestServices, count uint64) error {
	if count == 0 {
		return nil
	}

	client, err := rpc.DialContext(ctx, services.RPCURL)
	if err != nil {
		return fmt.Errorf("dial rpc client: %w", err)
	}
	defer client.Close()

	if err := client.CallContext(ctx, nil, "anvil_mine", count); err != nil {
		return fmt.Errorf("mine %d blocks: %w", count, err)
	}
	return nil
}

// WaitForFinalizeGate mines blocks (on Anvil) until block.number >=
// the epoch's liveNotBeforeBlock, opening the on-chain finalize gate.
// Used by integration tests/helpers that drive finalize directly.
func WaitForFinalizeGate(ctx context.Context, services *TestServices, epochID [12]byte) error {
	epoch, err := services.Contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return fmt.Errorf("get epoch: %w", err)
	}
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("read head: %w", err)
	}
	if head >= epoch.Policy.LiveNotBeforeBlock {
		return nil
	}
	return MineBlocks(ctx, services, epoch.Policy.LiveNotBeforeBlock-head)
}

// ChainTimestamp is the timestamp of the latest block — the value the
// application decryption window (`decryptNotBefore` / `decryptNotAfter`) is
// compared against.
func ChainTimestamp(ctx context.Context, services *TestServices) (uint64, error) {
	header, err := services.Contracts.Client().HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("read head header: %w", err)
	}
	return header.Time, nil
}

// MineUntilTimestamp mines blocks until the chain's clock has passed `ts`.
// Anvil timestamps follow the wall clock, so callers should keep the window
// they wait for short.
func MineUntilTimestamp(ctx context.Context, services *TestServices, ts uint64) error {
	return WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		if err := MineBlocks(ctx, services, 1); err != nil {
			return false
		}
		now, err := ChainTimestamp(ctx, services)
		return err == nil && now > ts
	})
}
