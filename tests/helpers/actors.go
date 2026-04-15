package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

type TestActor struct {
	Contracts *web3.Contracts
	Manager   *golangtypes.DKGManager
	Registry  *golangtypes.DKGRegistry
	TxManager *txmanager.Manager
}

func (a *TestActor) Address() common.Address {
	return a.TxManager.Address()
}

func (a *TestActor) CallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func (s *TestServices) ActorFromPrivateKey(privateKey string) (*TestActor, error) {
	txm, err := txmanager.New(s.Contracts.Pool().Current, s.Contracts.ChainID, privateKey)
	if err != nil {
		return nil, fmt.Errorf("new tx manager: %w", err)
	}
	return &TestActor{
		Contracts: s.Contracts,
		Manager:   s.Manager,
		Registry:  s.Registry,
		TxManager: txm,
	}, nil
}

func (s *TestServices) Actor(index int) (*TestActor, error) {
	if index < 0 || index >= len(DefaultAnvilPrivateKeys) {
		return nil, fmt.Errorf("actor index %d out of range", index)
	}
	return s.ActorFromPrivateKey(DefaultAnvilPrivateKeys[index])
}

func ClaimSlotAs(ctx context.Context, actor *TestActor, roundID [12]byte) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.ClaimSlot(auth, roundID)
	if err != nil {
		return fmt.Errorf("claim slot: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func SubmitContributionAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	contributorIndex uint16,
	commitmentsHash [32]byte,
	encryptedSharesHash [32]byte,
	transcript []byte,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.SubmitContribution(
		auth,
		roundID,
		contributorIndex,
		commitmentsHash,
		encryptedSharesHash,
		transcript,
		proof,
		input,
	)
	if err != nil {
		return fmt.Errorf("submit contribution: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func SubmitPartialDecryptionAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	participantIndex uint16,
	ciphertextIndex uint16,
	deltaHash [32]byte,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.SubmitPartialDecryption(
		auth,
		roundID,
		participantIndex,
		ciphertextIndex,
		deltaHash,
		proof,
		input,
	)
	if err != nil {
		return fmt.Errorf("submit partial decryption: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func SubmitRevealedShareAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	participantIndex uint16,
	shareValue *big.Int,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.SubmitRevealedShare(auth, roundID, participantIndex, shareValue, proof, input)
	if err != nil {
		return fmt.Errorf("submit revealed share: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func FinalizeRoundAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	aggregateCommitmentsHash [32]byte,
	collectivePublicKeyHash [32]byte,
	shareCommitmentHash [32]byte,
	transcript []byte,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.FinalizeRound(
		auth,
		roundID,
		aggregateCommitmentsHash,
		collectivePublicKeyHash,
		shareCommitmentHash,
		transcript,
		proof,
		input,
	)
	if err != nil {
		return fmt.Errorf("finalize round: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func CombineDecryptionAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	ciphertextIndex uint16,
	combineHash [32]byte,
	plaintextHash [32]byte,
	transcript []byte,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.CombineDecryption(auth, roundID, ciphertextIndex, combineHash, plaintextHash, transcript, proof, input)
	if err != nil {
		return fmt.Errorf("combine decryption: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func ReconstructSecretAs(
	ctx context.Context,
	actor *TestActor,
	roundID [12]byte,
	disclosureHash [32]byte,
	reconstructedSecretHash [32]byte,
	transcript []byte,
	proof []byte,
	input []byte,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.ReconstructSecret(auth, roundID, disclosureHash, reconstructedSecretHash, transcript, proof, input)
	if err != nil {
		return fmt.Errorf("reconstruct secret: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}
