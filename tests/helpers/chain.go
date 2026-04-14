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
