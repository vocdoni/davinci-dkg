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
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"}],"name":"getRound","outputs":[{"name":"organizer","type":"address"},{"name":"threshold","type":"uint16"},{"name":"committeeSize","type":"uint16"},{"name":"minValidContributions","type":"uint16"},{"name":"lotteryAlphaBps","type":"uint16"},{"name":"seedDelay","type":"uint16"},{"name":"registrationDeadlineBlock","type":"uint64"},{"name":"contributionDeadlineBlock","type":"uint64"},{"name":"disclosureAllowed","type":"bool"},{"name":"status","type":"uint8"},{"name":"nonce","type":"uint64"},{"name":"seedBlock","type":"uint64"},{"name":"seed","type":"bytes32"},{"name":"lotteryThreshold","type":"uint256"},{"name":"claimedCount","type":"uint16"},{"name":"contributionCount","type":"uint16"},{"name":"partialDecryptionCount","type":"uint16"},{"name":"revealedShareCount","type":"uint16"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getContributionVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getFinalizeVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getPartialDecryptVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getDecryptCombineVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getRevealSubmitVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[],"name":"getRevealShareVerifierVKeyHash","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"},{"internalType":"uint16","name":"participantIndex","type":"uint16"}],"name":"getShareCommitment","outputs":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"}],"name":"selectedParticipants","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
		{"inputs":[{"internalType":"bytes12","name":"roundId","type":"bytes12"},{"internalType":"uint16","name":"ciphertextIndex","type":"uint16"}],"name":"getCombinedDecryption","outputs":[{"name":"ciphertextIndex","type":"uint16"},{"name":"combineHash","type":"bytes32"},{"name":"plaintextHash","type":"bytes32"},{"name":"completed","type":"bool"}],"stateMutability":"view","type":"function"},
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

	client      *ethclient.Client
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
	CombineHash     common.Hash
	PlaintextHash   common.Hash
	Completed       bool
}

type RevealedShareView struct {
	Participant      common.Address
	ParticipantIndex uint16
	ShareValue       *big.Int
	ShareHash        common.Hash
	Accepted         bool
}

func New(rpcURL string, addresses types.ContractAddresses) (*Contracts, error) {
	if err := addresses.Validate(); err != nil {
		return nil, err
	}
	if rpcURL == "" {
		return nil, fmt.Errorf("rpc url is required")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	return &Contracts{
		ChainID:     chainID.Uint64(),
		Addresses:   addresses,
		client:      client,
		managerABI:  managerABI,
		registryABI: registryABI,
	}, nil
}

func (c *Contracts) Close() error {
	c.client.Close()
	return nil
}

func (c *Contracts) Client() *ethclient.Client {
	return c.client
}

func (c *Contracts) callHash(ctx context.Context, contract common.Address, contractABI *abi.ABI, method string) (common.Hash, error) {
	input, err := contractABI.Pack(method)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pack %s: %w", method, err)
	}

	output, err := c.client.CallContract(ctx, ethereum.CallMsg{
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
