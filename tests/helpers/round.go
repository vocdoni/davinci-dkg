package helpers

import (
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// ZeroDecryptionPolicy is an all-zero decryption policy: no owner restriction,
// no time locks, no submission cap. Used by tests that don't care about
// submission gating; callers constructing CreateEpoch calls directly should
// pass this to keep behaviour equivalent to the pre-DecryptionPolicy era.
func ZeroDecryptionPolicy() golangtypes.DKGTypesDecryptionPolicy {
	return golangtypes.DKGTypesDecryptionPolicy{}
}

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

func CreateEpoch(ctx context.Context, services *TestServices, policy types.EpochPolicy) ([12]byte, error) {
	var zero [12]byte

	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get epoch prefix: %w", err)
	}
	currentNonce, err := services.Manager.EpochNonce(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get epoch nonce: %w", err)
	}

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
		golangtypes.DKGTypesDecryptionPolicy{
			OwnerOnly:          policy.DecryptionPolicy.OwnerOnly,
			MaxDecryptions:     policy.DecryptionPolicy.MaxDecryptions,
			NotBeforeBlock:     policy.DecryptionPolicy.NotBeforeBlock,
			NotBeforeTimestamp: policy.DecryptionPolicy.NotBeforeTimestamp,
			NotAfterBlock:      policy.DecryptionPolicy.NotAfterBlock,
			NotAfterTimestamp:  policy.DecryptionPolicy.NotAfterTimestamp,
		},
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
