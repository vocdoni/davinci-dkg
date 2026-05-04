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
