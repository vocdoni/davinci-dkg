package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

const (
	roundStatusContribution uint8 = 2
	roundStatusFinalized    uint8 = 3
)

type FinalizedRoundResult struct {
	RoundID                  [12]byte
	Round                    web3.RoundView
	RoundHash                *big.Int
	Participant              common.Address
	AggregateCommitmentsHash common.Hash
	CollectivePublicKeyHash  common.Hash
	ShareCommitments         []types.CurvePoint
}

func CreateContributionRound(ctx context.Context, services *TestServices, policy types.RoundPolicy) ([12]byte, error) {
	var zero [12]byte

	if err := policy.Validate(); err != nil {
		return zero, err
	}

	roundID, err := CreateRound(ctx, services, policy)
	if err != nil {
		return zero, err
	}

	// Lottery flow: advance past seedBlock so blockhash is available, then claim.
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return zero, fmt.Errorf("get block number: %w", err)
	}
	seedBlock := head + uint64(policy.SeedDelay)
	if head <= seedBlock {
		if err := MineBlocks(ctx, services, seedBlock-head+1); err != nil {
			return zero, err
		}
	}
	if err := ClaimSlot(ctx, services, roundID); err != nil {
		return zero, err
	}
	if _, err := WaitRoundStatus(ctx, services, roundID, roundStatusContribution); err != nil {
		return zero, err
	}

	return roundID, nil
}

func CreateFinalizedSingleParticipantRound(
	ctx context.Context,
	services *TestServices,
	policy types.RoundPolicy,
	coefficients []*big.Int,
) (*FinalizedRoundResult, error) {
	roundID, err := CreateContributionRound(ctx, services, policy)
	if err != nil {
		return nil, err
	}

	submission, err := BuildContributionSubmission(ctx, services, roundID, 1, 1, 1, coefficients, []uint16{1})
	if err != nil {
		return nil, err
	}

	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := services.Manager.SubmitContribution(
		auth,
		roundID,
		1,
		submission.CommitmentsHash,
		submission.EncryptedSharesHash,
		submission.Transcript,
		submission.Proof,
		submission.Input,
	)
	if err != nil {
		return nil, fmt.Errorf("submit contribution: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return nil, err
	}

	finalizeOutput, err := BuildFinalizeRoundOutput(ctx, roundID, 1, 1, []uint16{1}, [][]*big.Int{coefficients})
	if err != nil {
		return nil, err
	}

	auth, err = services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err = services.Manager.FinalizeRound(
		auth,
		roundID,
		finalizeOutput.AggregateCommitmentsHash,
		finalizeOutput.CollectivePublicKeyHash,
		finalizeOutput.ShareCommitmentHash,
		finalizeOutput.Transcript,
		finalizeOutput.Proof,
		finalizeOutput.Input,
	)
	if err != nil {
		return nil, fmt.Errorf("finalize round: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return nil, err
	}

	round, err := WaitRoundStatus(ctx, services, roundID, roundStatusFinalized)
	if err != nil {
		return nil, err
	}

	return &FinalizedRoundResult{
		RoundID:                  roundID,
		Round:                    round,
		RoundHash:                finalizeOutput.RoundHash,
		Participant:              services.TxManager.Address(),
		AggregateCommitmentsHash: finalizeOutput.AggregateCommitmentsHash,
		CollectivePublicKeyHash:  finalizeOutput.CollectivePublicKeyHash,
		ShareCommitments:         finalizeOutput.ShareCommitments,
	}, nil
}
