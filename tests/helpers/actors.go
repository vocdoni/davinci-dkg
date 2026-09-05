package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
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

func ClaimSlotAs(ctx context.Context, actor *TestActor, epochID [12]byte) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.ClaimSlot(auth, epochID)
	if err != nil {
		return fmt.Errorf("claim slot: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func SubmitContributionAs(
	ctx context.Context,
	actor *TestActor,
	epochID [12]byte,
	contributorIndex uint16,
	sub *ContributionSubmission,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.SubmitContribution(
		auth,
		epochID,
		contributorIndex,
		sub.CommitmentsHash,
		sub.EncryptedSharesHash,
		sub.Transcript,
		sub.Proof,
		sub.Input,
	)
	if err != nil {
		return fmt.Errorf("submit contribution: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// SubmitPartialDecryptionAs posts one member's partial with the Merkle path
// proving its share commitment against the pool key's root.
func SubmitPartialDecryptionAs(
	ctx context.Context,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
	participantIndex uint16,
	ciphertextIndex uint16,
	partial *PartialDecryptionSubmission,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.SubmitPartialDecryption(
		auth,
		epochID,
		aid,
		participantIndex,
		ciphertextIndex,
		partial.C1.X, partial.C1.Y, partial.C2.X, partial.C2.Y,
		partial.DeltaHash,
		partial.Proof,
		partial.Input,
		partial.ShareProof,
	)
	if err != nil {
		return fmt.Errorf("submit partial decryption: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// FinalizeEpochAs closes the key-assembly phase with the batched finalization
// proof: the contract checks the phase, the liveNotBeforeBlock gate, the
// accepted contribution count, every dealer row against storage and the
// proof, then stores all pool keys and share roots. Permissionless.
func FinalizeEpochAs(ctx context.Context, actor *TestActor, epochID [12]byte, sub *FinalizeSubmission) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.FinalizeEpoch(auth, epochID, sub.TranscriptDigest, sub.Transcript, sub.Proof, sub.Input)
	if err != nil {
		return fmt.Errorf("finalize epoch: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// AbortEpochAs records a provably dead epoch as Aborted. Permissionless; the
// contract refuses it for any epoch that can still progress.
func AbortEpochAs(ctx context.Context, actor *TestActor, epochID [12]byte) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.AbortEpoch(auth, epochID)
	if err != nil {
		return fmt.Errorf("abort epoch: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

func CombineDecryptionAs(
	ctx context.Context,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	out *DecryptCombineOutput,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := actor.Manager.CombineDecryption(
		auth, epochID, aid, ciphertextIndex, out.CombineHash, out.Plaintext, out.Transcript, out.Proof, out.Input,
	)
	if err != nil {
		return fmt.Errorf("combine decryption: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// RegisterApplication claims the epoch's next unclaimed pool key for `aid`.
// An organizer-locked registration publishes PK_org = skOrg·G with its
// Schnorr proof of possession; an automatic one passes zero key and proof
// words (the contract stores the identity and ignores them), so `skOrg` may
// be nil or zero there.
func RegisterApplication(
	ctx context.Context,
	actor *TestActor,
	appManager *golangtypes.DKGAppManager,
	epochID [12]byte,
	aid [32]byte,
	skOrg *big.Int,
	policy golangtypes.DKGTypesAppPolicy,
) error {
	pkX, pkY := new(big.Int), new(big.Int)
	ax, ay, z := new(big.Int), new(big.Int), new(big.Int)
	if policy.Mode != uint8(types.AppModeAutomatic) {
		pkOrgX, pkOrgY, proof, err := schnorr.ProveOrganizerRegister(skOrg, epochID, aid)
		if err != nil {
			return fmt.Errorf("organizer schnorr proof: %w", err)
		}
		pkX, pkY = pkOrgX, pkOrgY
		ax, ay, z = proof.Ax, proof.Ay, proof.Z
	}
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := appManager.RegisterApplication(auth, epochID, aid, policy, pkX, pkY, ax, ay, z)
	if err != nil {
		return fmt.Errorf("register application: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// RevealOrganizerSecretAs publishes sk_org for an organizer-locked
// application. Permissionless and one-shot: from then on the committee
// combines by itself.
func RevealOrganizerSecretAs(
	ctx context.Context,
	actor *TestActor,
	appManager *golangtypes.DKGAppManager,
	epochID [12]byte,
	aid [32]byte,
	skOrg *big.Int,
) error {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := appManager.RevealOrganizerSecret(auth, epochID, aid, skOrg)
	if err != nil {
		return fmt.Errorf("reveal organizer secret: %w", err)
	}
	return actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// SubmitCiphertextAs submits a ciphertext for a registered application and
// returns the index the contract assigned to it. There is no proof of
// knowledge of the randomness — see DKGManager.submitCiphertext.
func SubmitCiphertextAs(
	ctx context.Context,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
	c1, c2 types.CurvePoint,
) (uint16, error) {
	index, _, err := SubmitCiphertextMeasured(ctx, actor, epochID, aid, c1, c2)
	return index, err
}

// SubmitCiphertextMeasured is SubmitCiphertextAs plus the gas used.
func SubmitCiphertextMeasured(
	ctx context.Context,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
	c1, c2 types.CurvePoint,
) (uint16, uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, 0, err
	}
	tx, err := actor.Manager.SubmitCiphertext(auth, epochID, aid, c1.X, c1.Y, c2.X, c2.Y)
	if err != nil {
		return 0, 0, fmt.Errorf("submit ciphertext: %w", err)
	}
	if err := actor.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout); err != nil {
		return 0, 0, err
	}
	receipt, err := actor.Contracts.Client().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, 0, fmt.Errorf("ciphertext receipt: %w", err)
	}
	for _, lg := range receipt.Logs {
		if ev, err := actor.Manager.ParseCiphertextSubmitted(*lg); err == nil {
			return ev.CiphertextIndex, receipt.GasUsed, nil
		}
	}
	return 0, receipt.GasUsed, fmt.Errorf("CiphertextSubmitted event not found in tx %s", tx.Hash().Hex())
}

// EnsureNodeKeyRegistered registers or updates the BJJ key for actor if it is
// not already registered with the correct key. The key is derived deterministically
// from actor.PrivKey using the same domain as the DKG node binary.
func EnsureNodeKeyRegistered(ctx context.Context, services *TestServices, actor *TestActor) error {
	expectedX, expectedY, secret, err := deterministicNodeKeyMaterial(actor.PrivKey)
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

	_, _, proof, err := schnorr.ProveOperatorRegister(secret, actor.Address())
	if err != nil {
		return fmt.Errorf("schnorr proof for %s: %w", actor.Address().Hex(), err)
	}

	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("tx opts for %s: %w", actor.Address().Hex(), err)
	}

	var txHash common.Hash
	if node.Status == 0 {
		tx, err := actor.Registry.RegisterKey(auth, expectedX, expectedY, proof.Ax, proof.Ay, proof.Z)
		if err != nil {
			return fmt.Errorf("register key for %s: %w", actor.Address().Hex(), err)
		}
		txHash = tx.Hash()
	} else {
		tx, err := actor.Registry.UpdateKey(auth, expectedX, expectedY, proof.Ax, proof.Ay, proof.Z)
		if err != nil {
			return fmt.Errorf("update key for %s: %w", actor.Address().Hex(), err)
		}
		txHash = tx.Hash()
	}
	return actor.TxManager.WaitTxByHash(txHash, DefaultTxTimeout)
}

// ── measured variants (gas profiles) ─────────────────────────────────────────

// ClaimSlotMeasured claims a slot for actor and returns the gas used.
func ClaimSlotMeasured(ctx context.Context, services *TestServices, actor *TestActor, epochID [12]byte) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.ClaimSlot(auth, epochID)
	if err != nil {
		return 0, fmt.Errorf("claim slot: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// SubmitContributionMeasured submits a contribution for actor and returns the gas used.
func SubmitContributionMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	epochID [12]byte,
	contributorIndex uint16,
	sub *ContributionSubmission,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.SubmitContribution(
		auth,
		epochID,
		contributorIndex,
		sub.CommitmentsHash,
		sub.EncryptedSharesHash,
		sub.Transcript,
		sub.Proof,
		sub.Input,
	)
	if err != nil {
		return 0, fmt.Errorf("submit contribution: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// FinalizeEpochMeasured finalizes and returns the gas used.
func FinalizeEpochMeasured(
	ctx context.Context, services *TestServices, actor *TestActor, epochID [12]byte, sub *FinalizeSubmission,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.FinalizeEpoch(auth, epochID, sub.TranscriptDigest, sub.Transcript, sub.Proof, sub.Input)
	if err != nil {
		return 0, fmt.Errorf("finalize epoch: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// RegisterApplicationMeasured registers an application and returns the gas used.
func RegisterApplicationMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	appManager *golangtypes.DKGAppManager,
	epochID [12]byte,
	aid [32]byte,
	skOrg *big.Int,
	policy golangtypes.DKGTypesAppPolicy,
) (uint64, error) {
	if err := RegisterApplication(ctx, actor, appManager, epochID, aid, skOrg, policy); err != nil {
		return 0, err
	}
	return lastGasFor(ctx, services, actor, appManager, epochID, aid)
}

// lastGasFor reads the gas of the ApplicationRegistered transaction. Kept
// separate so RegisterApplication itself stays free of receipt plumbing.
func lastGasFor(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	appManager *golangtypes.DKGAppManager,
	epochID [12]byte,
	aid [32]byte,
) (uint64, error) {
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	start := uint64(0)
	if head > 200 {
		start = head - 200
	}
	it, err := appManager.FilterApplicationRegistered(
		&bind.FilterOpts{Context: ctx, Start: start, End: &head},
		[][12]byte{epochID}, [][32]byte{aid}, nil,
	)
	if err != nil {
		return 0, fmt.Errorf("filter ApplicationRegistered: %w", err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		return 0, fmt.Errorf("no ApplicationRegistered event for aid %x", aid)
	}
	return gasOf(ctx, services, actor, it.Event.Raw.TxHash)
}

// RevealOrganizerSecretMeasured reveals sk_org and returns the gas used.
func RevealOrganizerSecretMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	appManager *golangtypes.DKGAppManager,
	epochID [12]byte,
	aid [32]byte,
	skOrg *big.Int,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := appManager.RevealOrganizerSecret(auth, epochID, aid, skOrg)
	if err != nil {
		return 0, fmt.Errorf("reveal organizer secret: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// SubmitPartialDecryptionMeasured submits a partial decryption for actor and returns the gas used.
func SubmitPartialDecryptionMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
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
		epochID,
		aid,
		participantIndex,
		ciphertextIndex,
		partial.C1.X, partial.C1.Y, partial.C2.X, partial.C2.Y,
		partial.DeltaHash,
		partial.Proof,
		partial.Input,
		partial.ShareProof,
	)
	if err != nil {
		return 0, fmt.Errorf("submit partial decryption: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// CombineDecryptionMeasured combines and returns the gas used.
func CombineDecryptionMeasured(
	ctx context.Context,
	services *TestServices,
	actor *TestActor,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	out *DecryptCombineOutput,
) (uint64, error) {
	auth, err := actor.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := actor.Manager.CombineDecryption(
		auth, epochID, aid, ciphertextIndex, out.CombineHash, out.Plaintext, out.Transcript, out.Proof, out.Input,
	)
	if err != nil {
		return 0, fmt.Errorf("combine decryption: %w", err)
	}
	return gasOf(ctx, services, actor, tx.Hash())
}

// gasOf waits for the transaction and returns its gas usage.
func gasOf(ctx context.Context, services *TestServices, actor *TestActor, hash common.Hash) (uint64, error) {
	if err := actor.TxManager.WaitTxByHash(hash, DefaultTxTimeout); err != nil {
		return 0, err
	}
	receipt, err := services.Contracts.Client().TransactionReceipt(ctx, hash)
	if err != nil {
		return 0, err
	}
	return receipt.GasUsed, nil
}
