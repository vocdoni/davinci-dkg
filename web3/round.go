package web3

import (
	"context"
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

// GetRevealSubmitVerifierVKeyHash returns the configured reveal-submit proving key hash.
func (c *Contracts) GetRevealSubmitVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getRevealSubmitVerifierVKeyHash")
}

// GetRevealShareVerifierVKeyHash returns the configured reveal-share proving key hash.
func (c *Contracts) GetRevealShareVerifierVKeyHash(ctx context.Context) (common.Hash, error) {
	return c.callHash(ctx, c.Addresses.Manager, c.managerABI, "getRevealShareVerifierVKeyHash")
}

// GetRound returns the on-chain round view.
func (c *Contracts) GetRound(ctx context.Context, roundID [12]byte) (RoundView, error) {
	input, err := c.managerABI.Pack("getRound", roundID)
	if err != nil {
		return RoundView{}, fmt.Errorf("pack getRound: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Manager,
		Data: input,
	}, nil)
	if err != nil {
		return RoundView{}, fmt.Errorf("call getRound: %w", err)
	}
	values, err := c.managerABI.Unpack("getRound", output)
	if err != nil {
		return RoundView{}, fmt.Errorf("unpack getRound: %w", err)
	}
	// Flat layout (per hand-written ABI; tuples count as single values):
	//   0 organizer                7 lotteryThreshold
	//   1 policy (tuple)           8 claimedCount
	//   2 decryptionPolicy (tuple) 9 contributionCount
	//   3 status                  10 partialDecryptionCount
	//   4 nonce                   11 revealedShareCount
	//   5 seedBlock               12 ciphertextCount
	//   6 seed
	if len(values) != 13 {
		return RoundView{}, fmt.Errorf("unexpected output count for getRound: %d", len(values))
	}
	policy, ok := values[1].(struct {
		Threshold                 uint16 `json:"threshold"`
		CommitteeSize             uint16 `json:"committeeSize"`
		MinValidContributions     uint16 `json:"minValidContributions"`
		LotteryAlphaBps           uint16 `json:"lotteryAlphaBps"`
		SeedDelay                 uint16 `json:"seedDelay"`
		RegistrationDeadlineBlock uint64 `json:"registrationDeadlineBlock"`
		ContributionDeadlineBlock uint64 `json:"contributionDeadlineBlock"`
		FinalizeNotBeforeBlock    uint64 `json:"finalizeNotBeforeBlock"`
		DisclosureAllowed         bool   `json:"disclosureAllowed"`
	})
	if !ok {
		return RoundView{}, fmt.Errorf("unexpected policy tuple shape")
	}
	seedBytes := values[6].([32]byte)
	return RoundView{
		Organizer: values[0].(common.Address),
		Policy: RoundPolicy{
			Threshold:                 policy.Threshold,
			CommitteeSize:             policy.CommitteeSize,
			MinValidContributions:     policy.MinValidContributions,
			LotteryAlphaBps:           policy.LotteryAlphaBps,
			SeedDelay:                 policy.SeedDelay,
			RegistrationDeadlineBlock: policy.RegistrationDeadlineBlock,
			ContributionDeadlineBlock: policy.ContributionDeadlineBlock,
			FinalizeNotBeforeBlock:    policy.FinalizeNotBeforeBlock,
			DisclosureAllowed:         policy.DisclosureAllowed,
		},
		Status:                 values[3].(uint8),
		Nonce:                  values[4].(uint64),
		SeedBlock:              values[5].(uint64),
		Seed:                   common.BytesToHash(seedBytes[:]),
		LotteryThreshold:       values[7].(*big.Int),
		ClaimedCount:           values[8].(uint16),
		ContributionCount:      values[9].(uint16),
		PartialDecryptionCount: values[10].(uint16),
		RevealedShareCount:     values[11].(uint16),
	}, nil
}

// SelectedParticipants returns the ordered participant set for a round.
func (c *Contracts) SelectedParticipants(ctx context.Context, roundID [12]byte) ([]common.Address, error) {
	input, err := c.managerABI.Pack("selectedParticipants", roundID)
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
