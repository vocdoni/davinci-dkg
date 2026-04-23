package web3

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/davinci-dkg/types"
)

const (
	dkgManagerABIJSON = `[
		{"inputs":[],"name":"REGISTRY","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"CONTRIBUTION_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"FINALIZE_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"PARTIAL_DECRYPT_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"DECRYPT_COMBINE_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"REVEAL_SUBMIT_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"REVEAL_SHARE_VERIFIER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"}],"name":"getRound","outputs":[{"name":"organizer","type":"address"},{"components":[{"name":"threshold","type":"uint16"},{"name":"committeeSize","type":"uint16"},{"name":"minValidContributions","type":"uint16"},{"name":"lotteryAlphaBps","type":"uint16"},{"name":"seedDelay","type":"uint16"},{"name":"registrationDeadlineBlock","type":"uint64"},{"name":"contributionDeadlineBlock","type":"uint64"},{"name":"finalizeNotBeforeBlock","type":"uint64"},{"name":"disclosureAllowed","type":"bool"}],"name":"policy","type":"tuple"},{"components":[{"name":"ownerOnly","type":"bool"},{"name":"maxDecryptions","type":"uint16"},{"name":"notBeforeBlock","type":"uint64"},{"name":"notBeforeTimestamp","type":"uint64"},{"name":"notAfterBlock","type":"uint64"},{"name":"notAfterTimestamp","type":"uint64"}],"name":"decryptionPolicy","type":"tuple"},{"name":"status","type":"uint8"},{"name":"nonce","type":"uint64"},{"name":"seedBlock","type":"uint64"},{"name":"seed","type":"bytes32"},{"name":"lotteryThreshold","type":"uint256"},{"name":"claimedCount","type":"uint16"},{"name":"contributionCount","type":"uint16"},{"name":"partialDecryptionCount","type":"uint16"},{"name":"revealedShareCount","type":"uint16"},{"name":"ciphertextCount","type":"uint16"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getContributionVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getFinalizeVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getPartialDecryptVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getDecryptCombineVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getRevealSubmitVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getRevealShareVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"},{"internalType":"uint16","name":"participantIndex","type":"uint16"}],"name":"getShareCommitment","outputs":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"}],"name":"selectedParticipants","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"},{"internalType":"uint16","name":"ciphertextIndex","type":"uint16"}],"name":"getCombinedDecryption","outputs":[{"name":"ciphertextIndex","type":"uint16"},{"name":"completed","type":"bool"},{"name":"plaintext","type":"uint256"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"},{"internalType":"address","name":"participant","type":"address"}],"name":"getRevealedShare","outputs":[{"name":"participant","type":"address"},{"name":"participantIndex","type":"uint16"},{"name":"shareValue","type":"uint256"},{"name":"shareHash","type":"bytes32"},{"name":"accepted","type":"bool"}],"stateMutability":"view","type":"function"}
	]`
	dkgRegistryABIJSON = `[
		{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"getNode","outputs":[{"name":"operator","type":"address"},{"name":"pubX","type":"uint256"},{"name":"pubY","type":"uint256"},{"name":"status","type":"uint8"}],"stateMutability":"view","type":"function"}
	]`
)

var (
	managerABI  = mustParseABI(dkgManagerABIJSON)
	registryABI = mustParseABI(dkgRegistryABIJSON)
)

type Contracts struct {
	ChainID   uint64
	Addresses types.ContractAddresses

	pool        *RPCPool
	managerABI  *abi.ABI
	registryABI *abi.ABI
}

type RegistryNode struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Status   uint8
}

type RoundPolicy struct {
	Threshold                 uint16
	CommitteeSize             uint16
	MinValidContributions     uint16
	LotteryAlphaBps           uint16
	SeedDelay                 uint16
	RegistrationDeadlineBlock uint64
	ContributionDeadlineBlock uint64
	FinalizeNotBeforeBlock    uint64
	DisclosureAllowed         bool
}

type RoundView struct {
	Organizer              common.Address
	Policy                 RoundPolicy
	Status                 uint8
	Nonce                  uint64
	SeedBlock              uint64
	Seed                   common.Hash
	LotteryThreshold       *big.Int
	ClaimedCount           uint16
	ContributionCount      uint16
	PartialDecryptionCount uint16
	RevealedShareCount     uint16
}

type CombinedDecryptionView struct {
	CiphertextIndex uint16
	Completed       bool
	Plaintext       *big.Int
}

type RevealedShareView struct {
	Participant      common.Address
	ParticipantIndex uint16
	ShareValue       *big.Int
	ShareHash        common.Hash
	Accepted         bool
}

// New creates a Contracts handle connected to the given RPC endpoints.
//
// The only required address is Manager. All other addresses (Registry plus the
// six verifiers) may be left as zero: New will query the manager's public
// immutable fields on-chain and fill them in automatically. This means that
// when using a well-known network preset (see config/networks.go) only the
// DKGManager address needs to be stored — everything else is derived at runtime.
//
// Explicitly-supplied non-zero addresses always take precedence over the
// on-chain values, so individual verifier overrides still work as before.
//
// Multiple RPC URLs may be provided; they are used in a round-robin pool with
// automatic failover (see RPCPool).
func New(rpcURLs []string, addresses types.ContractAddresses) (*Contracts, error) {
	if addresses.Manager == (common.Address{}) {
		return nil, fmt.Errorf("manager address is required")
	}
	if len(rpcURLs) == 0 {
		return nil, fmt.Errorf("at least one rpc url is required")
	}

	pool, err := NewRPCPool(rpcURLs)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	client := pool.Current()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	// Derive the registry address from the manager contract when not supplied.
	if addresses.Registry == (common.Address{}) {
		addr, err := fetchAddressFromManager(client, addresses.Manager, "REGISTRY")
		if err != nil {
			pool.Close()
			return nil, fmt.Errorf("derive registry from manager: %w", err)
		}
		addresses.Registry = addr
	}

	// Derive verifier addresses from the manager's public immutable fields when
	// not supplied. This allows callers to provide only the Manager address (e.g.
	// via a network preset) and have the full address book filled in on-chain.
	verifierFields := []struct {
		method string
		dest   *common.Address
	}{
		{"CONTRIBUTION_VERIFIER", &addresses.ContributionVerifier},
		{"FINALIZE_VERIFIER", &addresses.FinalizeVerifier},
		{"PARTIAL_DECRYPT_VERIFIER", &addresses.PartialDecryptVerifier},
		{"DECRYPT_COMBINE_VERIFIER", &addresses.DecryptCombineVerifier},
		{"REVEAL_SUBMIT_VERIFIER", &addresses.RevealSubmitVerifier},
		{"REVEAL_SHARE_VERIFIER", &addresses.RevealShareVerifier},
	}
	for _, vf := range verifierFields {
		if *vf.dest != (common.Address{}) {
			continue // explicit override — keep it
		}
		addr, err := fetchAddressFromManager(client, addresses.Manager, vf.method)
		if err != nil {
			pool.Close()
			return nil, fmt.Errorf("derive %s from manager: %w", vf.method, err)
		}
		*vf.dest = addr
	}

	if err := addresses.Validate(); err != nil {
		pool.Close()
		return nil, err
	}

	return &Contracts{
		ChainID:     chainID.Uint64(),
		Addresses:   addresses,
		pool:        pool,
		managerABI:  managerABI,
		registryABI: registryABI,
	}, nil
}

// fetchAddressFromManager calls a named view function on the DKGManager contract
// and returns the address it returns. Used to derive Registry and the six verifier
// addresses from the manager's public immutable fields.
func fetchAddressFromManager(client *ethclient.Client, manager common.Address, method string) (common.Address, error) {
	input, err := managerABI.Pack(method)
	if err != nil {
		return common.Address{}, fmt.Errorf("pack %s: %w", method, err)
	}
	output, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &manager,
		Data: input,
	}, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("call %s: %w", method, err)
	}
	values, err := managerABI.Unpack(method, output)
	if err != nil {
		return common.Address{}, fmt.Errorf("unpack %s: %w", method, err)
	}
	if len(values) != 1 {
		return common.Address{}, fmt.Errorf("unexpected output count for %s", method)
	}
	addr, ok := values[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("unexpected output type for %s", method)
	}
	return addr, nil
}

func (c *Contracts) Close() error {
	c.pool.Close()
	return nil
}

// Client returns the current active RPC client from the pool.
func (c *Contracts) Client() *ethclient.Client {
	return c.pool.Current()
}

// Pool returns the underlying RPCPool.
func (c *Contracts) Pool() *RPCPool {
	return c.pool
}

// PooledBackend returns a bind.ContractBackend backed by the pool.
func (c *Contracts) PooledBackend() *PooledBackend {
	return NewPooledBackend(c.pool)
}

func (c *Contracts) callHash(ctx context.Context, contract common.Address, contractABI *abi.ABI, method string) (common.Hash, error) {
	input, err := contractABI.Pack(method)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pack %s: %w", method, err)
	}

	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: input,
	}, nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("call %s: %w", method, err)
	}

	values, err := contractABI.Unpack(method, output)
	if err != nil {
		return common.Hash{}, fmt.Errorf("unpack %s: %w", method, err)
	}
	if len(values) != 1 {
		return common.Hash{}, fmt.Errorf("unexpected output count for %s", method)
	}

	switch value := values[0].(type) {
	case [32]byte:
		return common.BytesToHash(value[:]), nil
	case common.Hash:
		return value, nil
	default:
		return common.Hash{}, fmt.Errorf("unexpected output type for %s", method)
	}
}

func parseABI(raw string) (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	return &parsed, nil
}

func mustParseABI(raw string) *abi.ABI {
	parsed, err := parseABI(raw)
	if err != nil {
		panic(err)
	}
	return parsed
}
