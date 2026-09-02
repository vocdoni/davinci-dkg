package web3

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// GetContributionVerifierVKeyHash returns the configured contribution proving key hash.
func (c *Contracts) GetContributionVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getContributionVerifierVKeyHash")
}

// GetPartialDecryptVerifierVKeyHash returns the configured partial decrypt proving key hash.
func (c *Contracts) GetPartialDecryptVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getPartialDecryptVerifierVKeyHash")
}

// GetFinalizeVerifierVKeyHash returns the configured finalize proving key hash.
func (c *Contracts) GetFinalizeVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getFinalizeVerifierVKeyHash")
}

// GetDecryptCombineVerifierVKeyHash returns the configured decrypt-combine proving key hash.
func (c *Contracts) GetDecryptCombineVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getDecryptCombineVerifierVKeyHash")
}

// GetEpoch returns the on-chain epoch view.
func (c *Contracts) GetEpoch(ctx context.Context, epochID [12]byte) (EpochView, error) {
	input, err := c.managerABI.Pack("getEpoch", epochID)
	if err != nil {
		return EpochView{}, fmt.Errorf("pack getEpoch: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Manager,
		Data: input,
	}, nil)
	if err != nil {
		return EpochView{}, fmt.Errorf("call getEpoch: %w", err)
	}
	values, err := c.managerABI.Unpack("getEpoch", output)
	if err != nil {
		return EpochView{}, fmt.Errorf("unpack getEpoch: %w", err)
	}
	// Flat layout (per hand-written ABI; tuples count as single values):
	//   0  organizer         6  seed
	//   1  policy (tuple)    7  lotteryThreshold
	//   2  status            8  claimedCount
	//   3  nonce             9  contributionCount
	//   4  startBlock       10  partialDecryptionCount
	//   5  seedBlock        11  ciphertextCount
	if len(values) != 12 {
		return EpochView{}, fmt.Errorf("unexpected output count for getEpoch: %d", len(values))
	}
	policy, ok := values[1].(struct {
		Threshold                       uint16 `json:"threshold"`
		CommitteeSize                   uint16 `json:"committeeSize"`
		MinValidContributions           uint16 `json:"minValidContributions"`
		LotteryAlphaBps                 uint16 `json:"lotteryAlphaBps"`
		CommitteeSelectionDeadlineBlock uint64 `json:"committeeSelectionDeadlineBlock"`
		KeyAssemblyDeadlineBlock        uint64 `json:"keyAssemblyDeadlineBlock"`
		LiveNotBeforeBlock              uint64 `json:"liveNotBeforeBlock"`
	})
	if !ok {
		return EpochView{}, fmt.Errorf("unexpected policy tuple shape")
	}
	seedBytes := values[6].([32]byte)
	return EpochView{
		Organizer: values[0].(common.Address),
		Policy: EpochPolicy{
			Threshold:                       policy.Threshold,
			CommitteeSize:                   policy.CommitteeSize,
			MinValidContributions:           policy.MinValidContributions,
			LotteryAlphaBps:                 policy.LotteryAlphaBps,
			CommitteeSelectionDeadlineBlock: policy.CommitteeSelectionDeadlineBlock,
			KeyAssemblyDeadlineBlock:        policy.KeyAssemblyDeadlineBlock,
			LiveNotBeforeBlock:              policy.LiveNotBeforeBlock,
		},
		Status:                 values[2].(uint8),
		Nonce:                  values[3].(uint64),
		StartBlock:             values[4].(uint64),
		SeedBlock:              values[5].(uint64),
		Seed:                   common.BytesToHash(seedBytes[:]),
		LotteryThreshold:       values[7].(*big.Int),
		ClaimedCount:           values[8].(uint16),
		ContributionCount:      values[9].(uint16),
		PartialDecryptionCount: values[10].(uint16),
		CiphertextCount:        values[11].(uint16),
	}, nil
}

// SelectedParticipants returns the ordered participant set for a epoch.
func (c *Contracts) SelectedParticipants(ctx context.Context, epochID [12]byte) ([]common.Address, error) {
	input, err := c.managerABI.Pack("selectedParticipants", epochID)
	if err != nil {
		return nil, fmt.Errorf("pack selectedParticipants: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Manager,
		Data: input,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("call selectedParticipants: %w", err)
	}
	values, err := c.managerABI.Unpack("selectedParticipants", output)
	if err != nil {
		return nil, fmt.Errorf("unpack selectedParticipants: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("unexpected output count for selectedParticipants")
	}
	participants, ok := values[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected output type for selectedParticipants")
	}
	return participants, nil
}

// EpochID builds the 12-byte epoch identifier `prefix ‖ nonce` used by the
// DKGManager (see DKGIdLib.computeEpochId).
func EpochID(prefix uint32, nonce uint64) [12]byte {
	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], prefix)
	binary.BigEndian.PutUint64(id[4:], nonce)
	return id
}
