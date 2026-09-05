package web3

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
)

// GetContributionVerifierVKeyHash returns the configured contribution proving key hash.
func (c *Contracts) GetContributionVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getContributionVerifierVKeyHash")
}

// GetPartialDecryptVerifierVKeyHash returns the configured partial decrypt proving key hash.
func (c *Contracts) GetPartialDecryptVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getPartialDecryptVerifierVKeyHash")
}

// GetPoolKeyVerifierVKeyHash returns the configured pool-key proving key hash.
func (c *Contracts) GetPoolKeyVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getPoolKeyVerifierVKeyHash")
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

// CommitteeSnapshot returns the committee members' share-encryption keys
// exactly as they were frozen when the committee filled (the CommitteeSnapshot
// event), not the live registry keys: a later updateKey rotation must not
// change what contributors encrypt shares to. Entry i is committee position
// i+1; Operator is left zero and must be filled by the caller from the
// selected-participants list.
func (c *Contracts) CommitteeSnapshot(ctx context.Context, epochID [12]byte) ([]types.NodeKey, error) {
	epoch, err := c.GetEpoch(ctx, epochID)
	if err != nil {
		return nil, fmt.Errorf("get epoch: %w", err)
	}
	filterer, err := gtypes.NewDKGManagerFilterer(c.Addresses.Manager, c.PooledBackend())
	if err != nil {
		return nil, fmt.Errorf("bind DKGManager filterer: %w", err)
	}
	it, err := filterer.FilterCommitteeSnapshot(&bind.FilterOpts{Context: ctx, Start: epoch.StartBlock}, [][12]byte{epochID})
	if err != nil {
		return nil, fmt.Errorf("filter CommitteeSnapshot from block %d: %w", epoch.StartBlock, err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return nil, fmt.Errorf("iterate CommitteeSnapshot: %w", err)
		}
		return nil, fmt.Errorf("committee snapshot event not found for epoch %x", epochID)
	}
	ev := it.Event
	if len(ev.PubKeys) != int(ev.CommitteeSize)*2 {
		return nil, fmt.Errorf(
			"committee snapshot for epoch %x carries %d keys for committeeSize=%d",
			epochID, len(ev.PubKeys)/2, ev.CommitteeSize)
	}
	keys := make([]types.NodeKey, ev.CommitteeSize)
	for i := range keys {
		keys[i] = types.NodeKey{PubX: ev.PubKeys[i*2], PubY: ev.PubKeys[i*2+1]}
	}
	return keys, nil
}

// EpochID builds the 12-byte epoch identifier `prefix ‖ nonce` used by the
// DKGManager (see DKGIdLib.computeEpochId).
func EpochID(prefix uint32, nonce uint64) [12]byte {
	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], prefix)
	binary.BigEndian.PutUint64(id[4:], nonce)
	return id
}
