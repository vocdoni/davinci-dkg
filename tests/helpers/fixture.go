package helpers

import (
	"context"
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
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

	// web3.New fills AppManager in from the manager's `appManager()` view when
	// the addresses file predates APP_MANAGER, so read it back from there.
	appManager, err := golangtypes.NewDKGAppManager(contracts.Addresses.AppManager, contracts.Client())
	if err != nil {
		_ = contracts.Close()
		return nil, nil, err
	}

	services := &TestServices{
		RPCURL:     rpcURL,
		Addresses:  contracts.Addresses,
		Contracts:  contracts,
		Registry:   registry,
		Manager:    manager,
		AppManager: appManager,
		TxManager:  txm,
	}

	if err := bootstrapLocalNodeKeys(ctx, services); err != nil {
		_ = contracts.Close()
		return nil, nil, fmt.Errorf("bootstrap node keys: %w", err)
	}

	cleanup := func() { _ = contracts.Close() }
	return services, cleanup, nil
}

// CreateSDKTestFixture creates a Live single-participant epoch on the given
// testnet with `keys` of its pool activated (at least key 0). Useful as a
// fixture for TypeScript SDK tests that need a Live epoch with a claimable
// pool key without generating ZK proofs themselves.
func CreateSDKTestFixture(
	ctx context.Context,
	services *TestServices,
	keys int,
) (*FinalizedRoundResult, error) {
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("get block number: %w", err)
	}

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}
	coefficients := []*big.Int{big.NewInt(FixtureShare)}

	result, err := CreateFinalizedSingleParticipantRound(ctx, services, policy, coefficients)
	if err != nil {
		return nil, err
	}
	for keyIndex := 1; keyIndex < keys; keyIndex++ {
		if keyIndex >= ccommon.MaxK {
			return nil, fmt.Errorf("cannot activate %d keys, the pool holds %d", keys, ccommon.MaxK)
		}
		if _, err := ActivateRoundPoolKey(ctx, services, result, uint8(keyIndex)); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// FixtureShare is the single contribution coefficient CreateSDKTestFixture
// deals, and therefore participant 1's share of pool key 0.
const FixtureShare int64 = 11
