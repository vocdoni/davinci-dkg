package helpers

import (
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

func RoundIDFromString(value string) [12]byte {
	var epochID [12]byte
	copy(epochID[:], []byte(value))
	return epochID
}

func RoundIDToString(epochID [12]byte) string {
	return strings.TrimRight(string(epochID[:]), "\x00")
}

func ComputeRoundID(prefix uint32, nonce uint64) [12]byte {
	var epochID [12]byte
	binary.BigEndian.PutUint32(epochID[:4], prefix)
	binary.BigEndian.PutUint64(epochID[4:], nonce)
	return epochID
}

// CreateEpoch creates an epoch once the cadence allows it, mining up to
// nextEpochStartBlock first.
func CreateEpoch(ctx context.Context, services *TestServices, policy types.EpochPolicy) ([12]byte, error) {
	var zero [12]byte

	// Honor the on-chain cadence guard: createEpoch reverts unless
	// block.number >= nextEpochStartBlock(). Tests run sequentially against a
	// single anvil instance, so without this advance the second epoch in a
	// suite would silently revert and the helper would hang waiting for the
	// epoch to reach Contribution.
	nextStart, err := services.Manager.NextEpochStartBlock(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get next epoch start block: %w", err)
	}
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return zero, fmt.Errorf("get block number: %w", err)
	}
	if head < nextStart {
		if err := MineBlocks(ctx, services, nextStart-head); err != nil {
			return zero, fmt.Errorf("mine to next epoch start: %w", err)
		}
	}
	return CreateEpochNow(ctx, services, policy)
}

// CreateEpochNow calls createEpoch at the current block, without honoring
// the cadence: the contract decides. Before nextEpochStartBlock it only
// succeeds when the newest epoch is Live with at most one unclaimed key or
// Aborted, which is exactly what the early-creation tests probe.
func CreateEpochNow(ctx context.Context, services *TestServices, policy types.EpochPolicy) ([12]byte, error) {
	var zero [12]byte

	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get epoch prefix: %w", err)
	}
	currentNonce, err := services.Manager.EpochNonce(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get epoch nonce: %w", err)
	}

	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return zero, err
	}
	tx, err := services.Manager.CreateEpoch(
		auth,
		policy.Threshold,
		policy.CommitteeSize,
		policy.MinValidContributions,
		policy.LotteryAlphaBps,
	)
	if err != nil {
		return zero, fmt.Errorf("create epoch: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return zero, err
	}

	return ComputeRoundID(prefix, currentNonce+1), nil
}

// ClaimSlot has the caller race for a committee slot in the epoch. The caller
// must be a registered node and pass the lottery for that epoch.
func ClaimSlot(ctx context.Context, services *TestServices, epochID [12]byte) error {
	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := services.Manager.ClaimSlot(auth, epochID)
	if err != nil {
		return fmt.Errorf("claim slot: %w", err)
	}
	return services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func WaitEpochPhase(ctx context.Context, services *TestServices, epochID [12]byte, status uint8) (web3.EpochView, error) {
	var epoch web3.EpochView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		epoch, fetchErr = services.Contracts.GetEpoch(ctx, epochID)
		return fetchErr == nil && epoch.Status == status
	})
	if err != nil {
		return web3.EpochView{}, err
	}
	return epoch, nil
}
