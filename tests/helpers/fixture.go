package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/config"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// NewTestServicesFromExternal creates a TestServices connected to an already-running
// testnet. Unlike NewTestServices with env vars, this also bootstraps all default
// Anvil node keys so that claimSlot works out of the box.
//
// Used by the TypeScript SDK integration test fixture (cmd/sdk-test-fixture).
func NewTestServicesFromExternal(
	ctx context.Context,
	rpcURL string,
	addressesContent []byte,
) (*TestServices, func(), error) {
	addresses, err := config.ParseContractAddressesEnv(addressesContent)
	if err != nil {
		return nil, nil, fmt.Errorf("parse contract addresses: %w", err)
	}

	contracts, err := web3.New([]string{rpcURL}, addresses)
	if err != nil {
		return nil, nil, fmt.Errorf("connect to chain: %w", err)
	}

	txm, err := txmanager.New(contracts.Pool().Current, contracts.ChainID, LocalAccountPrivKey)
	if err != nil {
		_ = contracts.Close()
		return nil, nil, fmt.Errorf("create tx manager: %w", err)
	}

	registry, err := golangtypes.NewDKGRegistry(addresses.Registry, contracts.Client())
	if err != nil {
		_ = contracts.Close()
		return nil, nil, err
	}

	manager, err := golangtypes.NewDKGManager(addresses.Manager, contracts.Client())
	if err != nil {
		_ = contracts.Close()
		return nil, nil, err
	}

	services := &TestServices{
		RPCURL:    rpcURL,
		Addresses: addresses,
		Contracts: contracts,
		Registry:  registry,
		Manager:   manager,
		TxManager: txm,
	}

	if err := bootstrapLocalNodeKeys(ctx, services); err != nil {
		_ = contracts.Close()
		return nil, nil, fmt.Errorf("bootstrap node keys: %w", err)
	}

	cleanup := func() { _ = contracts.Close() }
	return services, cleanup, nil
}

// CreateSDKTestFixture creates a finalized single-participant round on the given
// testnet and returns the round ID. Useful as a fixture for TypeScript SDK tests
// that need a finalized round without generating ZK proofs themselves.
func CreateSDKTestFixture(
	ctx context.Context,
	services *TestServices,
) (*FinalizedRoundResult, error) {
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("get block number: %w", err)
	}

	policy := types.RoundPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		DisclosureAllowed:         false,
	}
	coefficients := []*big.Int{big.NewInt(11)}

	return CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
}
