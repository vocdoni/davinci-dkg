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
// submission gating; callers constructing CreateRound calls directly should
// pass this to keep behaviour equivalent to the pre-DecryptionPolicy era.
func ZeroDecryptionPolicy() golangtypes.DKGTypesDecryptionPolicy {
	return golangtypes.DKGTypesDecryptionPolicy{}
}

func RoundIDFromString(value string) [12]byte {
	var roundID [12]byte
	copy(roundID[:], []byte(value))
	return roundID
}

func RoundIDToString(roundID [12]byte) string {
	return strings.TrimRight(string(roundID[:]), "\x00")
}

func ComputeRoundID(prefix uint32, nonce uint64) [12]byte {
	var roundID [12]byte
	binary.BigEndian.PutUint32(roundID[:4], prefix)
	binary.BigEndian.PutUint64(roundID[4:], nonce)
	return roundID
}

func CreateRound(ctx context.Context, services *TestServices, policy types.RoundPolicy) ([12]byte, error) {
	var zero [12]byte

	prefix, err := services.Manager.ROUNDPREFIX(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get round prefix: %w", err)
	}
	currentNonce, err := services.Manager.RoundNonce(services.CallOpts(ctx))
	if err != nil {
		return zero, fmt.Errorf("get round nonce: %w", err)
	}

	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return zero, err
	}
	tx, err := services.Manager.CreateRound(
		auth,
		policy.Threshold,
		policy.CommitteeSize,
		policy.MinValidContributions,
		policy.LotteryAlphaBps,
		policy.SeedDelay,
		policy.RegistrationDeadlineBlock,
		policy.ContributionDeadlineBlock,
		policy.DisclosureAllowed,
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
		return zero, fmt.Errorf("create round: %w", err)
	}
	if err := services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return zero, err
	}

	return ComputeRoundID(prefix, currentNonce+1), nil
}

// ClaimSlot has the caller race for a committee slot in the round. The caller
// must be a registered node and pass the lottery for that round.
func ClaimSlot(ctx context.Context, services *TestServices, roundID [12]byte) error {
	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := services.Manager.ClaimSlot(auth, roundID)
	if err != nil {
		return fmt.Errorf("claim slot: %w", err)
	}
	return services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func WaitRoundStatus(ctx context.Context, services *TestServices, roundID [12]byte, status uint8) (web3.RoundView, error) {
	var round web3.RoundView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		round, fetchErr = services.Contracts.GetRound(ctx, roundID)
		return fetchErr == nil && round.Status == status
	})
	if err != nil {
		return web3.RoundView{}, err
	}
	return round, nil
}
