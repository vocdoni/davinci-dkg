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
	PrivKey   string // hex private key; used for deterministic BJJ key derivation
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
		PrivKey:   privateKey,
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
	commitment0X *big.Int,
	commitment0Y *big.Int,
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
		commitment0X,
		commitment0Y,
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

// EnsureNodeKeyRegistered registers or updates the BJJ key for actor if it is
// not already registered with the correct key. The key is derived deterministically
// from actor.PrivKey using the same domain as the DKG node binary.
func EnsureNodeKeyRegistered(ctx context.Context, services *TestServices, actor *TestActor) error {
	expectedX, expectedY, _, err := deterministicNodeKeyMaterial(actor.PrivKey)
	if err != nil {
		return fmt.Errorf("derive deterministic node key for %s: %w", actor.Address().Hex(), err)
	}

	node, err := services.Contracts.GetNode(ctx, actor.Address())
	if err != nil {
		return fmt.Errorf("get node for %s: %w", actor.Address().Hex(), err)
	}
	if node.Status != 0 &&
		node.Operator == actor.Address() &&
		node.PubX != nil && node.PubY != nil &&
		node.PubX.Cmp(expectedX) == 0 &&
		node.PubY.Cmp(expectedY) == 0 {
		return nil
	}

	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for %s: %w", actor.Address().Hex(), err)
	}

	var txHash common.Hash
	if node.Status == 0 {
		tx, err := actor.Registry.RegisterKey(auth, expectedX, expectedY)
		if err != nil {
			return fmt.Errorf("register key for %s: %w", actor.Address().Hex(), err)
		}
		txHash = tx.Hash()
	} else {
		tx, err := actor.Registry.UpdateKey(auth, expectedX, expectedY)
		if err != nil {
			return fmt.Errorf("update key for %s: %w", actor.Address().Hex(), err)
		}
		txHash = tx.Hash()
	}
	return actor.TxManager.WaitTxByHash(txHash, DefaultTxTimeout)
}

// ClaimSlotMeasured claims a slot for actor and returns the gas used.
func ClaimSlotMeasured(ctx context.Context, services *TestServices, actor *TestActor, roundID [12]byte) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.ClaimSlot(auth, roundID)
	if err != nil {
		return 0, fmt.Errorf("claim slot: %w", err)
	}
	if err := actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return 0, err
	}
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, err
	}
	return receipt.GasUsed, nil
}

// SubmitContributionMeasured submits a contribution for actor and returns the gas used.
func SubmitContributionMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	roundID [12]byte,
	contributorIndex uint16,
	sub *ContributionSubmission,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.SubmitContribution(
		auth,
		roundID,
		contributorIndex,
		sub.CommitmentsHash,
		sub.EncryptedSharesHash,
		sub.Commitment0X,
		sub.Commitment0Y,
		sub.Transcript,
		sub.Proof,
		sub.Input,
	)
	if err != nil {
		return 0, fmt.Errorf("submit contribution: %w", err)
	}
	if err := actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return 0, err
	}
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, err
	}
	return receipt.GasUsed, nil
}

// SubmitPartialDecryptionMeasured submits a partial decryption for actor and returns the gas used.
func SubmitPartialDecryptionMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	roundID [12]byte,
	participantIndex uint16,
	ciphertextIndex uint16,
	partial *PartialDecryptionSubmission,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.SubmitPartialDecryption(
		auth,
		roundID,
		participantIndex,
		ciphertextIndex,
		partial.DeltaHash,
		partial.Proof,
		partial.Input,
	)
	if err != nil {
		return 0, fmt.Errorf("submit partial decryption: %w", err)
	}
	if err := actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return 0, err
	}
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, err
	}
	return receipt.GasUsed, nil
}
