package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
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

// DefaultLotteryAlphaBps is the over-subscription factor applied to integration
// test round policies when the caller leaves LotteryAlphaBps at zero. Matches
// the runtime default used by cmd/dkg-runner.
const DefaultLotteryAlphaBps uint16 = 15000

// DefaultSeedDelay is the seed-block offset used by integration test policies
// when the caller does not specify one. Matches cmd/dkg-runner.
const DefaultSeedDelay uint16 = 1

func CreateContributionRound(ctx context.Context, services *TestServices, policy types.RoundPolicy) ([12]byte, error) {
	var zero [12]byte

	if policy.LotteryAlphaBps == 0 {
		policy.LotteryAlphaBps = DefaultLotteryAlphaBps
	}
	if policy.SeedDelay == 0 {
		policy.SeedDelay = DefaultSeedDelay
	}
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
		submission.Commitment0X,
		submission.Commitment0Y,
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

	// Wait until block.number >= finalizeNotBeforeBlock so the on-chain gate is open.
	if err := WaitForFinalizeGate(ctx, services, roundID); err != nil {
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

// CombineSingleParticipantDecryption drives partial decryption + combine for a
// ciphertext that is already on-chain at (roundID, ciphertextIndex), assuming
// the round was created by CreateFinalizedSingleParticipantRound (committee=1,
// threshold=1, single participant index 1 owned by services.TxManager).
//
// `share` is the polynomial share value held by participant 1 — for a single
// coefficient list `coefficients`, this is f(1) = sum(coefficients).
//
// Used by the SDK end-to-end ciphertext test (sdk/tests/ciphertext-e2e):
// the SDK submits an encrypted ciphertext, then this helper drives the
// committee-side decryption and the combineDecryption call so the SDK can
// read the recovered plaintext via getPlaintext and assert correctness.
func CombineSingleParticipantDecryption(
	ctx context.Context,
	services *TestServices,
	roundID [12]byte,
	ciphertextIndex uint16,
	share *big.Int,
) error {
	ciphertextHash, err := services.Manager.GetCiphertextHash(services.CallOpts(ctx), roundID, ciphertextIndex)
	if err != nil {
		return fmt.Errorf("get ciphertext hash: %w", err)
	}
	var zero common.Hash
	if ciphertextHash == zero {
		return fmt.Errorf("ciphertext at (%x, %d) not yet submitted", roundID, ciphertextIndex)
	}

	// Recover the actual ciphertext coordinates from the CiphertextSubmitted
	// event log (the contract only stores the keccak hash). We scan from the
	// round's seedBlock to limit the filter range.
	round, err := services.Contracts.GetRound(ctx, roundID)
	if err != nil {
		return fmt.Errorf("get round: %w", err)
	}
	startBlock := uint64(0)
	if round.SeedBlock > 0 {
		startBlock = uint64(round.SeedBlock) - 1
	}
	latest, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("read head: %w", err)
	}
	filterOpts := &bind.FilterOpts{Context: ctx, Start: startBlock, End: &latest}
	it, err := services.Manager.FilterCiphertextSubmitted(filterOpts, [][12]byte{roundID}, []uint16{ciphertextIndex}, nil)
	if err != nil {
		return fmt.Errorf("filter CiphertextSubmitted: %w", err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return fmt.Errorf("iterate CiphertextSubmitted: %w", err)
		}
		return fmt.Errorf("no CiphertextSubmitted event for (%x, %d)", roundID, ciphertextIndex)
	}
	c1 := types.CurvePoint{X: new(big.Int).Set(it.Event.C1x), Y: new(big.Int).Set(it.Event.C1y)}
	c2 := types.CurvePoint{X: new(big.Int).Set(it.Event.C2x), Y: new(big.Int).Set(it.Event.C2y)}

	// Build the partial decryption (delta = share * c1) using the gnark proof.
	// For a single-participant committee the partial decryption IS the combined
	// decryption (Lagrange interpolation at zero with a single share is the
	// share itself), but we still go through both txs so this exercises the
	// full on-chain decryption path the SDK consumers depend on.
	const partialNonce = 1
	partial, err := BuildPartialDecryptionSubmissionFromBase(ctx, roundID, 1, c1, share, big.NewInt(partialNonce))
	if err != nil {
		return fmt.Errorf("build partial decryption: %w", err)
	}

	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := services.Manager.SubmitPartialDecryption(auth, roundID, 1, ciphertextIndex, partial.DeltaHash, partial.Proof, partial.Input)
	if err != nil {
		return fmt.Errorf("submit partial decryption: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return err
	}

	// Recover the plaintext by brute-force discrete log over a small window.
	// Fixture rounds always submit ciphertexts of small integers so this
	// terminates immediately; the helper bounds the search to 2^20 to match
	// the SDK's `decrypt` and avoid hangs on bad inputs.
	plaintext, err := bruteForceELGamalPlaintext(c2, partial.Delta)
	if err != nil {
		return fmt.Errorf("recover plaintext: %w", err)
	}

	combineOutput, err := BuildDecryptCombineOutputFromCiphertext(
		ctx, roundID, 1, c1, c2, []uint16{1}, []types.CurvePoint{partial.Delta}, plaintext,
	)
	if err != nil {
		return fmt.Errorf("build combine output: %w", err)
	}

	auth, err = services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err = services.Manager.CombineDecryption(
		auth, roundID, ciphertextIndex,
		combineOutput.CombineHash, combineOutput.Plaintext,
		combineOutput.Transcript, combineOutput.Proof, combineOutput.Input,
	)
	if err != nil {
		return fmt.Errorf("combine decryption: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return err
	}
	return nil
}

// bruteForceELGamalPlaintext recovers m from c2 - delta = m·G by trying every
// m in [0, 2^20). Used by CombineSingleParticipantDecryption to discover the
// plaintext that the SDK encrypted (the SDK chose a random k, so this is the
// only way the fixture can learn what was sent without round-tripping through
// the original encryption).
//
// Note on the loop: gnark's PointAffine zero value is the affine origin (0, 0),
// which is NOT a point on twisted Edwards (the identity is (0, 1)). We can't
// start from `group.NewPoint()` and add G repeatedly, because adding (0, 0) + G
// produces an invalid result. Instead we use the identity encoded as (0, 1)
// for the m=0 check and start the iteration from G itself for m=1+.
func bruteForceELGamalPlaintext(c2 types.CurvePoint, delta types.CurvePoint) (*big.Int, error) {
	c2Native, err := group.Decode(c2)
	if err != nil {
		return nil, fmt.Errorf("decode c2: %w", err)
	}
	deltaNative, err := group.Decode(delta)
	if err != nil {
		return nil, fmt.Errorf("decode delta: %w", err)
	}
	negDelta := group.NewPoint()
	negDelta.Neg(deltaNative)

	target := group.NewPoint()
	target.Add(c2Native, negDelta)
	targetEnc := group.Encode(target)

	// m = 0 → target is the curve identity (0, 1)
	if targetEnc.X.Sign() == 0 && targetEnc.Y.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0), nil
	}

	// m >= 1: candidate starts at G, iterate candidate += G.
	g := group.Generator()
	candidate := group.NewPoint()
	candidate.Set(g)
	for i := int64(1); i < 1<<20; i++ {
		candEnc := group.Encode(candidate)
		if candEnc.X.Cmp(targetEnc.X) == 0 && candEnc.Y.Cmp(targetEnc.Y) == 0 {
			return big.NewInt(i), nil
		}
		next := group.NewPoint()
		next.Add(candidate, g)
		candidate = next
	}
	return nil, fmt.Errorf("plaintext out of brute-force range (> 2^20)")
}

// silence "imported and not used" — `bind` is referenced inside the
// CombineSingleParticipantDecryption body via &bind.FilterOpts{}.
var _ = bind.CallOpts{}
