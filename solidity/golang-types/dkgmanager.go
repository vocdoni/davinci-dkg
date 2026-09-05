// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package golangtypes

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DKGTypesCombinedDecryptionRecord is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesCombinedDecryptionRecord struct {
	CiphertextIndex uint16
	Completed       bool
	Plaintext       *big.Int
}

// DKGTypesContributionRecord is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesContributionRecord struct {
	Contributor         common.Address
	ContributorIndex    uint16
	CommitmentsHash     [32]byte
	EncryptedSharesHash [32]byte
	Accepted            bool
}

// DKGTypesEpochPolicy is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesEpochPolicy struct {
	Threshold                       uint16
	CommitteeSize                   uint16
	MinValidContributions           uint16
	LotteryAlphaBps                 uint16
	CommitteeSelectionDeadlineBlock uint64
	KeyAssemblyDeadlineBlock        uint64
	LiveNotBeforeBlock              uint64
}

// DKGTypesPartialDecryptionRecord is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesPartialDecryptionRecord struct {
	ParticipantIndex uint16
	CiphertextIndex  uint16
	DeltaHash        [32]byte
	Accepted         bool
}

// IDKGManagerEpoch is an auto generated low-level Go binding around an user-defined struct.
type IDKGManagerEpoch struct {
	Organizer              common.Address
	Policy                 DKGTypesEpochPolicy
	Status                 uint8
	Nonce                  uint64
	StartBlock             uint64
	SeedBlock              uint64
	Seed                   [32]byte
	LotteryThreshold       *big.Int
	ClaimedCount           uint16
	ContributionCount      uint16
	PartialDecryptionCount uint16
	CiphertextCount        uint16
}

// DKGManagerMetaData contains all meta data concerning the DKGManager contract.
var DKGManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolKeyVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epochDurationBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_committeeSelectionBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_keyAssemblyBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finalizeGapBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minCommitteeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_maxLotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LOTTERY_ALPHA_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_COMMITTEE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_KEY_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activatePoolKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"transcriptDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ciphertextCount\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimPoolKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochDurationBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppPoolIndex\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSelectionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyAssemblyDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liveNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolKeyVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolShareRoot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolStatus\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"nextIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"activated\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"shareProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeFilled\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeSnapshot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"pubKeys\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochLive\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolKeyActivated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolKeyClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DirectCallRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInSnapshot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExhausted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolKeyAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolKeyNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TranscriptWordNotInField\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6102608060405234610438576101a081614b848038038091610021828561043c565b8339810103126104385780519063ffffffff8216908183036104385761004960208201610473565b9261005660408301610473565b61006260608401610473565b9061006f60808501610473565b61007b60a08601610473565b60c08601519160e087015194610100880151946101208901519a6101408a016100a390610487565b986100b16101608c01610487565b9a610180016100bf90610487565b9b4663ffffffff1603610429576001600160a01b0382161561041a576001600160a01b038316158015610409575b80156103f8575b80156103e7575b6103d85763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261013760388261043c565b5190201660c05260e052610100526101205261014052806103d257506064915b806103cc57506019905b806103c657506019955b806103c057506005915b610188836101838985610496565b610496565b9260018311908115916103b7575b81156103ae575b5080156103a4575b8015610394575b61035657610160526001600160401b038181166101e05261ffff9690916101d39190610496565b16610200526001600160401b03166102205280841661038e57506001905b8161018052838116155f1461038857506001915b826101a052838116155f1461038057508280925b836101c052169283911611918215610375575b508115610365575b506103565733610240526040516146cc90816104b8823960805181611815015260a0518181816103b001528181611456015281816123ea0152613eff015260c05181818161249e01528181612a0c01526136e0015260e05181818161133801528181612baf0152613ea9015261010051818181610c16015281816117730152612f5c015261012051818181613387015281816136670152613a0601526101405181818161020c01528181611eb401526122f5015261016051818181614027015261423201526101805181818161108d015261296a01526101a05181818161025e015261293901526101c0518181816102c6015261290801526101e051816124df015261020051816125160152610220518161254a01526102405181612b110152f35b63d06b96b160e01b5f5260045ffd5b612710915061ffff16105f610234565b60201091505f61022c565b839092610219565b91610205565b906101f1565b506001600160401b0381116101ac565b50808310156101a5565b9050155f61019d565b88159150610196565b91610175565b9561016b565b90610161565b91610157565b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100fb565b506001600160a01b038516156100f4565b506001600160a01b038416156100ed565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761045f57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361043857565b519061ffff8216820361043857565b919082018092116104a357565b634e487b7160e01b5f52601160045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816304da574014613f2e5750806306433b1b14613eea578063074a75e114613e84578063134c98e01461381e57806318287e5f1461370457806323488be5146136c3578063268ae2a1146136965780632b743f04146136515780632de546d51461360b5780633681b55914613506578063368e2a27146134d3578063421adfbb146133c857806343e50afa1461336057806356cbb5f3146132c15780635a8f2bb3146132785780635b0c034714612bde57806363f314cd14612b995780636d16897d14612ae85780636f067f6314612a9f57806371712c291461023e57806371a5978c1461233657806372517b4b146122ce57806377235ee114611abe5780637ade132414611a6f5780637b31b5661461183957806385e1f4d0146117f85780638dc1f53a1461174c5780639bbada67146116b0578063a4adcd7f14611689578063b7bca615146110b1578063bd11c4c014611072578063be59b8ea14610d6b578063bea5210d14610c45578063bf19220914610c00578063ca3c045814610b38578063d3720aac14610a40578063d3979253146109f0578063d9933767146102ea578063d9e9ca2e146102ab578063ebe86c1314610282578063f03a489814610243578063fa8f5e961461023e5763fe1604b5146101f7575f80fd5b3461023b578060031936011261023b576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b80fd5b614010565b503461023b578060031936011261023b57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461023b578060031936011261023b576001546040516001600160a01b039091168152602090f35b503461023b578060031936011261023b57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461023b57602036600319011261023b576001600160a01b031961030d613f53565b1680825260026020526040822080546001600160a01b0316156109e1576002810190815460ff8116600183019283549160068110156109cd57600114806109b6575b156109a757600581019061ffff808354169360101c1683101561099857868852600360209081526040808a20335f908152925290205460ff1661098957600381018054908115610916575b506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316959060c0816024818a5afa90811561090b578b916108ec575b50606081015160038110156108d8575f19016108c95760a001516001600160401b0390811660489290921c1611156108ba57600490604051602081019182523360601b6040820152603481526104486054826140d2565b51902091015411156108ab57858752600460205260408720805490600160401b82101561089757916104878261050c94600161ffff9795018155614282565b81546001600160a01b03600392831b90811b19909116339182901b17909255898b5260209081526040808c205f9384529091529020805460ff19166001179055836104d183614333565b168419825416179055604051818152877f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a3614333565b925492818460101c169182911614610522578580f35b60405161040061053281836140d2565b3682376040519361080061054681876140d2565b368637600f1c936201fffe61fffe861695168503610883579084939291610586610570899761456e565b9561057e60405197886140d2565b80875261456e565b602086019490601f1901368637868a52600460205260408a20908a5b85811061072b57505050825b602081106106d75750604051906020820192838b905b602082106106bd5750505061042082018a905b604082106106a357505050610c0081526105f3610c20826140d2565b519020848852600a602052604088205560405192604084019184526040602085015251809152606083019190875b81811061068a5750505090807f3687c3912698ec8584ed18d7c68916ff5fb9418efc8c6d52e0da152938e50f6a920390a2805460ff191660021790557f23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb0678280a25f808080808580f35b8251845287955060209384019390920191600101610621565b825181528a985060209283019260019290920191016105d7565b825181528b995060209283019260019290920191016105c4565b809192939495965060011b81810460021482151715610717576001810180911161071757906001610709819385614585565b5201908795949392916105ae565b634e487b7160e01b8a52601160045260248afd5b9091929394959697506001810180821161086f5761074982876143e5565b526107548184614282565b90546040516313a4120960e31b815260039290921b1c6001600160a01b0316600482015260c081602481865afa908115610864578c91610836575b5060208101908151908360011b9184830460021485151715610822578e9392916040916107bc848b614585565b52019182519350600182019384831161080d57928f936107f36107fd948f929488956107eb60019b9a8f614585565b525192614596565b525192508b614596565b52019089979695949392916105a2565b50634e487b7160e01b8f52601160045260248ffd5b634e487b7160e01b8f52601160045260248ffd5b610857915060c03d811161085d575b61084f81836140d2565b81019061435c565b5f61078f565b503d610845565b6040513d8e823e3d90fd5b634e487b7160e01b8c52601160045260248cfd5b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b633802147960e11b8952600489fd5b63aba4733960e01b8b5260048bfd5b634e487b7160e01b8c52602160045260248cfd5b610905915060c03d60c01161085d5761084f81836140d2565b5f6103f1565b6040513d8d823e3d90fd5b9050608885901c6001600160401b03164381101561097a574090811561096b57819055877fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602083604051908152a25f61039a565b6302504bb360e61b8a5260048afd5b63172181cb60e21b8a5260048afd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b50604082901c6001600160401b031643111561034f565b634e487b7160e01b88526021600452602488fd5b63d5b25b6360e01b8352600483fd5b503461023b57602036600319011261023b5760409060ff826001600160a01b0319610a19613f53565b1692838152600e602052828282205416938152600d60205220541682519182526020820152f35b503461023b57604036600319011261023b57610a5a613f53565b60243591906001600160a01b0383168303610b3457906040918160808451610a81816140b7565b8281528260208201528286820152826060820152015260018060a01b03191681526005602052209060018060a01b03165f5260205260a060405f20604051610ac8816140b7565b815491600180851b0383169283835261ffff6020840191861c16815261ffff60018301549160408501928352608060ff60036002870154966060890197885201541695019415158552604051958652511660208501525160408401525160608301525115156080820152f35b5080fd5b503461023b57602036600319011261023b576001600160a01b0319610b5b613f53565b168152600460205260408120604051908160208254918281520190819285526020852090855b818110610be15750505082610b979103836140d2565b604051928392602084019060208552518091526040840192915b818110610bbf575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610bb1565b82546001600160a01b0316845260209093019260019283019201610b81565b503461023b578060031936011261023b576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461023b57608036600319011261023b57608090610c62613f53565b9060243591610c6f613fa7565b610c77613fb8565b91836060604051610c878161409c565b828152826020820152826040820152015260018060a01b031916918284526007602052604084208585526020526040842061ffff82165f5260205261ffff6040600180825f20548487161c161495865f14610d645784975b8715610d5d5784965b825260066020528282209082526020522091165f5260205261ffff60405f2091165f5260205261ffff60405f20549160608260405196610d278861409c565b16958681528360208201931683526040810194855201938452604051948552511660208401525160408301525115156060820152f35b8196610ce8565b8097610cdf565b503461023b57602036600319011261023b57610d85613f53565b81610160604051610d9581614065565b828152604051610da48161404a565b8381528360208201528360408201528360608201528360808201528360a08201528360c082015260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015282610140820152015260018060a01b031916815260026020526040812060405190610e2c82614065565b80546001600160a01b03168252604051610e458161404a565b600182015461ffff8116825261ffff8160101c16602083015261ffff8160201c16604083015261ffff8160301c16606083015260018060401b038160401c16608083015260018060401b038160801c1660a083015260c01c60c08201526020830190815260028201549160ff8316906040850191600681101561105e578252606085019360018060401b038160081c168552608086019060018060401b038160481c16825260a087019060018060401b039060881c16815260038301549160c08801928352600560048501549460e08a0195865201549661010089019961ffff89168b526101208a019661ffff8a60101c1688526101408b019861ffff8b60201c168a5261ffff6101608d019b60301c168b526040519b60018060a01b039051168c525161ffff81511660208d015261ffff60208201511660408d015261ffff60408201511660608d015261ffff60608201511660808d015260018060401b0360808201511660a08d015260018060401b0360a08201511660c08d015260c060018060401b039101511660e08c01525190600682101561104a57506101008a0152516001600160401b039081166101208a01529051811661014089015290511661016087015251610180860152516101a0850152845161ffff9081166101c0860152905181166101e08501529051811661020084015281511661022083015261024082f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b87526021600452602487fd5b503461023b578060031936011261023b57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461023b5760e036600319011261023b576110cb613f53565b6110d3613fc9565b906064356044356084356001600160401b03811161155c576110f9903690600401613f7a565b9460a4356001600160401b03811161168557611119903690600401613f7a565b95909260c4356001600160401b0381116116815761113b903690600401613f7a565b979098323314801590611677575b611668576001600160a01b03198416808c52600260205260408c208054919a909790916001600160a01b0316156116595760ff6002890154169b60018901549c6006811015611645576002148061162e575b1561161f5760408e8d815260036020522060018060a01b0333165f5260205260ff60405f205416156116105761ffff88169c8d158015611600575b6115f1576111f98f808f604092526004602052206111f38b61426e565b90614282565b90543360039290921b1c6001600160a01b0316036115705760408f8e815260056020522060018060a01b0333165f5260205260ff600360405f200154166115e2576112468483018361411c565b9788519060a01c14908115916115cf575b81156115b8575b5080156115aa575b801561159c575b801561158e575b61157f576103a09594939291906173ff198401611570575f5160206146775f395f51905f528f8e908d8f6112d76112ad8c8b3691614170565b60208151910120916112c9604051938492602084019687614347565b03601f1981018352826140d2565b519020905060405190602082019283527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c8152611321606c826140d2565b51902006948560c08a015103611560578f939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561155c57859361138a60405196879586948594635c73957b60e11b8652600486016141d5565b03915afa801561150257611543575b5050614c001161153f576113b436610c006140008501614170565b60208151910120898c52600a60205260408c20540361153057916113da9160e0936143f6565b9101510361152157906003600592868952836020526040892060018060a01b0333165f5260205260405f209081549061ffff60a01b9060a01b169061ffff60a01b191617815584600182015501600160ff198254161790550161ffff815460101c1661ffff811461150d579060016114539201906142fa565b847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b15610b3457818091602460405180948193633c1bcdef60e21b83523360048401525af18015611502576114e9575b5050604051938452602084015260408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b816114f3916140d2565b6114fe57845f6114ae565b8480fd5b6040513d84823e3d90fd5b634e487b7160e01b87526011600452602487fd5b63d1fed5fd60e01b8752600487fd5b63d1fed5fd60e01b8b5260048bfd5b8a80fd5b8161154d916140d2565b611558578b5f611399565b8b80fd5b8580fd5b5063d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8e5260048efd5b508a60a08801511415611274565b50896080880151141561126d565b508c60608801511415611266565b905061ffff60408901519160101c1614155f61125e565b602089015161ffff821614159150611257565b6305d252c360e01b8f5260048ffd5b63652122d960e01b8f5260048ffd5b508d61ffff8260101c16106111d6565b63965c290d60e01b8e5260048efd5b63268dbf6760e21b8e5260048efd5b5060808d901c6001600160401b031643111561119b565b634e487b7160e01b8f52602160045260248ffd5b63d5b25b6360e01b8d5260048dfd5b6346f4d6c560e11b8b5260048bfd5b50333b1515611149565b8980fd5b8780fd5b503461023b578060031936011261023b57546040516001600160401b039091168152602090f35b503461023b5761ffff60406116c436613fda565b9491828480516116d381614081565b828152826020820152015260018060a01b031916825260096020528282209082526020522091165f52602052606060405f2060405161171181614081565b81546040600161ffff83169485855260ff602086019460101c1615158452015492019182526040519283525115156020830152516040820152f35b503461023b578060031936011261023b5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156117ec57906117b5575b602090604051908152f35b506020813d6020116117e4575b816117cf602093836140d2565b810103126117e057602090516117aa565b5f80fd5b3d91506117c2565b604051903d90823e3d90fd5b503461023b578060031936011261023b57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461023b5760c036600319011261023b57611853613f53565b6001600160a01b03191680825260026020526040822080549091602435916044359060a4359060843590606435906001600160a01b031615611a605760ff6002880154166006811015611a4c57600303611a3d576118b18185614507565b6118bb8383614507565b84885260086020526040882086895260205261ffff6118df8160408b205416614333565b16966101008811611a2e5760015460ff8160a01c1615611a1f576001600160a01b031689813b1561023b5760849160405192838092633a54cd5d60e11b82528b60048301528c60248301528d60448301523360648301525afa8015611a14576119fb575b509260a0928695926119d860057f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc978d8d9c9b60209f5260088f52604081208c82528f52604081208d61ffff1982541617905560406119a48888888c6144a9565b918c81526020601090528d828220908252602052208d5f528f5260405f20550161ffff600181835460301c16011690614314565b604051933385528b850152604084015260608301526080820152a4604051908152f35b611a068a80926140d2565b611a10575f611943565b8880fd5b6040513d8c823e3d90fd5b63023b34fb60e11b8a5260048afd5b63464e67af60e01b8952600489fd5b63268dbf6760e21b8852600488fd5b634e487b7160e01b89526021600452602489fd5b63d5b25b6360e01b8852600488fd5b503461023b57604036600319011261023b5760ff6040611a8d613f53565b92611a96613f6a565b9360018060a01b0319168152600c6020522091165f52602052602060405f2054604051908152f35b503461023b5761010036600319011261023b57611ad9613f53565b611ae1613fa7565b9060a4356001600160401b0381116122ca57611b01903690600401613f7a565b9060c4356001600160401b03811161155c57611b21903690600401613f7a565b90919060e4356001600160401b03811161168557611b43903690600401613f7a565b6001600160a01b0319871689526002602052604089208054929590926001600160a01b0316156122bb5760ff60028401541660068110156122a7576003036122985761ffff8916158015612289575b801561227f575b612270576001600160a01b031988168a52601060209081526040808c206024358d528252808c2061ffff8c165f908152925290205495861561226157600154976001600160a01b0389163b15611558578b611c0f818c8c6040518080958194633db5946d60e21b835260243590600484016142ab565b03916001600160a01b03165afa80156115025761224c575b50506001600160a01b03198a168c52600760209081526040808e206024358f528252808e2061ffff8e165f90815292529020548c9590805b6122185750600101549461ffff8616116122095760018060a01b03198a168c52600960205260408c206024358d5260205260408c2061ffff8c165f5260205260405f209889549860ff8a60101c166121fa578386019291908e6101208686031261023b5784601f8701121561023b575061012060405194611ce082876140d2565b8590828801116121f55786905b82880182106121e5575050508d8d85519060a01c14908115916121d4575b81156121c1575b5080156121af575b801561219f575b801561218f575b6115705761ffff881660808501511061157057602060808501511161157057610cc08303611570578e8360801161023b5760405190611d6860a0836140d2565b608082523660808d011161023b5760808c602084013760a0820152602081519101200361157f5760ff8160a01c1615612180578d60408d611dc193825180809681946381fe92fb60e01b835260243590600484016142ab565b03916001600160a01b03165afa918215612173578192612135575b5060808a01351490811591612126575b506120f957611e0b5f5160206146775f395f51905f529136908a614170565b60208151910120604051611e2d816112c9602082019460843560643587614347565b5190206040516001600160a01b03198d16602082019081527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c80830193909352918152611e82606c826140d2565b519020068060e0830151036120f957606688611e9d926143f6565b610100820151036120ea5760800151948b939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561155c578593611f0660405196879586948594635c73957b60e11b8652600486016141d5565b03915afa801561150257612111575b505060c083018311610883576104c0830183116108835760101c61ffff1687805b838110611ffc575050505b60208110611fac575050620100009062ff00001916178155600160843591015561ffff604051926064358452608435602085015216917f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae160406024359360018060a01b03191692a480f35b6104c08160061b83010160c08260051b84010135611fed5780351590811591611fdd575b5061152157600101611f41565b600191506020013514155f611fd0565b63d1fed5fd60e01b8852600488fd5b60c08160051b860101356104c08260061b8701019080158015612108575b6120ea5761ffff16906001821b938481166120f95784179360018060a01b03198b168d52600760205260408d208d602435905260205260408d2061ffff8d165f5260205260405f205416156120ea5760405160208082019280358452013560408201526040815261208c6060826140d2565b5190209060018060a01b03198a168c52600660205260408c206024358d5260205260408c2061ffff8c165f5260205261ffff60405f2091165f5260205260405f2054036120db57600101611f36565b63d1fed5fd60e01b8a5260048afd5b63d1fed5fd60e01b8c5260048cfd5b63d1fed5fd60e01b8d5260048dfd5b5084811161201a565b8161211b916140d2565b61168557875f611f15565b905060a089013514155f611dec565b9150506040813d60401161216b575b81612151604093836140d2565b810103126121675760208151910151905f611ddc565b8d80fd5b3d9150612144565b50604051903d90823e3d90fd5b63023b34fb60e11b8e5260048efd5b5060843560c08501511415611d28565b5060643560a08501511415611d21565b5061ffff881660608501511415611d1a565b61ffff915016604085015114158e611d12565b602086015160243514159150611d0b565b8135815260209182019101611ced565b508f80fd5b63955c0c4960e01b8e5260048efd5b63032cddf960e11b8c5260048cfd5b5f198101908082116108225716955f198114612238576001019580611c5f565b634e487b7160e01b8e52601160045260248efd5b81612256916140d2565b611558578b5f611c27565b6346f551f560e01b8b5260048bfd5b636d28699160e01b8a5260048afd5b5060643515611b99565b5061010061ffff8a1611611b92565b63268dbf6760e21b8a5260048afd5b634e487b7160e01b8b52602160045260248bfd5b63d5b25b6360e01b8a5260048afd5b8380fd5b503461023b578060031936011261023b5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156117ec57906117b557602090604051908152f35b503461023b57608036600319011261023b5760043561ffff8116809103610b345761235f613fc9565b612367613fa7565b91612370613fb8565b906001600160401b03612381614218565b1643106129f1575b801580156129e5575b80156129d8575b80156129ca575b80156129be575b80156129ad575b80156129a0575b8015612991575b8015612964575b8015612933575b8015612902575b6128f357604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156128e85786916128ae575b506001600160401b031691821561289f5761ffff80911693169161244683856140f5565b9080612710029061271082040361150d578082106128725750505f19915b8554906001600160401b03808316908114610883576001600160401b036001909101166001600160401b031992909216821787556124c2827f00000000000000000000000000000000000000000000000000000000000000006144df565b95436001600160401b039081166001019690871161285e5761250d7f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166141ff565b9061ffff6125447f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166141ff565b936125787f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166141ff565b95604051986125868a61404a565b895260208901521660408088019190915260608701919091526001600160401b03918216608087015291811660a08601529190911660c084015251916125cb83614065565b338352602083015260016040830152606082015260018060401b034316608082015260018060401b03831660a08201528460c08201528160e08201528461010082015284610120820152846101408201528461016082015260018060a01b03198416855260026020526040852060018060a01b0382511660018060a01b031982541617815560018101602083015161ffff808251161661ffff1983541617825561267d61ffff602083015116836142fa565b61268f61ffff604083015116836142c7565b6126a161ffff60608301511683614314565b608081810151835460a084015160c0948501516001600160401b03909216604093841b600160401b600160801b031617931b600160801b600160c01b0316929092179190921b6001600160c01b03191617909155820151600681101561105e576002820180546060850151608086015160a08701516001600160c81b031990931660ff9095169490941760089190911b610100600160481b03161760489390931b600160481b600160881b03169290921760889290921b600160881b600160c81b031691909117905560c0820151600382015560e08201516004820155610100820151600591909101805461ffff191661ffff928316178155610120830151602097936127d4939161016091906127ba908416856142fa565b6127cb8361014083015116856142c7565b01511690614314565b8054600160401b600160801b03191643604081811b600160401b600160801b03169290921790925580516001600160401b03928316815291909316858201529182015233906001600160a01b03198316907f1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d90606090a36040516001600160a01b03199091168152f35b634e487b7160e01b89526011600452602489fd5b801561288b5761288591905f19046140f5565b91612464565b634e487b7160e01b87526012600452602487fd5b63d06b96b160e01b8652600486fd5b90506020813d6020116128e0575b816128c9602093836140d2565b8101031261155c576128da906142e6565b5f612422565b3d91506128bc565b6040513d88823e3d90fd5b63d06b96b160e01b8552600485fd5b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff8316116123d1565b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff8416106123ca565b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001681106123c3565b5061271061ffff8316106123bc565b508061ffff8516106123b5565b5061ffff831661ffff8516116123ae565b5061ffff8416156123a7565b50602061ffff8416116123a0565b5061ffff83168111612399565b5061ffff831615612392565b84546001600160a01b031990612a30906001600160401b03167f00000000000000000000000000000000000000000000000000000000000000006144df565b16808652600260205260ff60026040882001541690600682101561105e57600382149081612a84575b50159081612a78575b5015612389575b63268dbf6760e21b8552600485fd5b6004915014155f612a62565b90508652600e602052600760ff60408820541610155f612a59565b503461023b57604036600319011261023b5760209061ffff906040906001600160a01b0319612acc613f53565b1681526008845281812060243582528452205416604051908152f35b503461023b57602036600319011261023b576004356001600160a01b03811690819003610b34577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303612b8b576001549060ff8260a01c16612b7c578015612b6d576001600160a81b031990911617600160a01b1760015580f35b63e6c4247b60e01b8352600483fd5b6373253a9760e01b8352600483fd5b6282b42960e81b8252600482fd5b503461023b578060031936011261023b576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461023b5761018036600319011261023b57612bf9613f53565b612c01613fa7565b90612c0a613fb8565b610124356001600160401b0381116114fe57612c2a903690600401613f7a565b90610144356001600160401b03811161327457612c4b903690600401613f7a565b90916001600160401b0361016435116116855736602361016435011215611685576001600160401b0360046101643501351161168557366024610164356004013560051b61016435010111611685576001600160a01b03198616885260026020526040882080549093906001600160a01b0316156132655760ff600285015416600681101561325157600303613242576001600160a01b031987168952600360209081526040808b20335f908152925290205460ff16156132335761ffff881615801561321b575b801561320f575b8015613200575b80156131f5575b6131e7576001600160a01b031987168952600460205260408920612d4f906111f38a61426e565b90543360039290921b1c6001600160a01b0316036131c9576001600160a01b031987168952601060209081526040808b206024358c528252808b2061ffff89165f908152925290205480156131d857612db260e43560c43560a4356084356144a9565b036131c95760015489906001600160a01b0316803b15610b3457604051633db5946d60e21b815290829082908180612df06024358f600484016142ab565b03915afa8015611502576131b4575b505060018060a01b0319871689526007602052604089206024358a526020526040892061ffff87165f5260205260405f20600161ffff8a161b9054166131a557828101946101e0828703126116815785601f830112156116815760405195612e696101e0886140d2565b86906101e084011161153f5782905b6101e08401821061318157505085518860a01c14801590613171575b801561315f575b801561314d575b801561313d575b801561312d575b6120db576101008601516101208701516040519060208201928352604082015260408152612edf6060826140d2565b51902061010435036120db576001600160a01b031988168a52600c60205260408a2060ff612f0f6024358b614466565b165f5260205260405f2054612f2d60c088015160e0890151906144c7565b60056101643560040135036120ea5761ffff8b165f19018c5b600581106130cd575050036120db5789939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561155c578593612fae60405196879586948594635c73957b60e11b8652600486016141d5565b03915afa8015611502576130b8575b505060056130599160018060a01b03198616885260066020526040882060243589526020526040882061ffff86165f5260205260405f2061ffff88165f5260205260405f2061010435905560018060a01b03198616885260076020526040882060243589526020526040882061ffff86165f5260205260405f20600161ffff89161b81541790550161ffff600181835460201c160116906142c7565b61ffff6101206101008301519201519281604051961686521660208501526040840152606083015233917f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30260806024359360018060a01b03191692a480f35b816130c2916140d2565b61155c57855f612fbd565b909161016435600401358210156131195760019061016435600584901b01602401359084831661310b5790613101916145aa565b92811c9101612f46565b613114916145aa565b613101565b634e487b7160e01b8e52603260045260248efd5b5060a43560a08701511415612eb0565b5060843560808701511415612ea9565b5061ffff891660608701511415612ea2565b50604086015161ffff88161415612e9b565b5060243560208701511415612e94565b8135815260209182019101612e78565b634e487b7160e01b5f52604160045260245ffd5b633466526160e01b8952600489fd5b816131be916140d2565b611a1057885f612dff565b63d1fed5fd60e01b8952600489fd5b6346f551f560e01b8a5260048afd5b62d949df60e51b8952600489fd5b506101043515612d28565b5061010061ffff871611612d21565b5061ffff861615612d1a565b5061ffff600185015460101c1661ffff891611612d13565b63965c290d60e01b8952600489fd5b63268dbf6760e21b8952600489fd5b634e487b7160e01b8a52602160045260248afd5b63d5b25b6360e01b8952600489fd5b8680fd5b503461023b5761ffff604061328c36613fda565b949160018060a01b031916825260096020528282209082526020522091165f526020526020600160405f200154604051908152f35b503461023b57604036600319011261023b576132db613f53565b9060ff6132e6613f6a565b16916008831080159061333a575b61332b579060409160018060a01b0319168152600b60205220905f526020526040805f206001815491015482519182526020820152f35b633b948c2560e11b8252600482fd5b506001600160a01b031981168252600d60205260408220546001841b1660ff16156132f4565b503461023b578060031936011261023b5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa9081156117ec57906117b557602090604051908152f35b503461023b57604036600319011261023b576133e2613f53565b60015460243591906001600160a01b031633036134c5576001600160a01b031916808352600e602052604083205460ff16919060088310156134b657808452600d60205260ff60408520546001851b1616156134a75760209360ff604081600187011692848152600e88528181208385168419825416179055848152600f8852818120868252885220911660ff198254161790557f587aa85ec6cb98aa5d1c21fbe47dbf442a2432b78629190c195062eb34a0303c84604051858152a3604051908152f35b633b948c2560e11b8452600484fd5b6311a7ebfd60e31b8452600484fd5b6282b42960e81b8352600483fd5b503461023b57604036600319011261023b5760206134fb6134f2613f53565b60243590614466565b60ff60405191168152f35b503461023b57602036600319011261023b576001600160a01b0319613529613f53565b1680825260026020526040822080546001600160a01b0316156109e1576002810160ff81541660068110156135f757600381146135e857600119016135d9576001820154918260c01c4310612a695761ffff60058192015460101c169260201c1682106135ca57805460ff191660031790556040519081527f0a9d060d45776170692faf35cb6f7bdda0152d8d49f36631cfa3547235467f6390602090a280f35b63368f2d7d60e21b8452600484fd5b63268dbf6760e21b8452600484fd5b6337bca76b60e21b8552600485fd5b634e487b7160e01b85526021600452602485fd5b503461023b5761ffff604061361f36613fda565b949160018060a01b031916825260106020528282209082526020522091165f52602052602060405f2054604051908152f35b503461023b578060031936011261023b576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461023b578060031936011261023b5760206136b1614218565b6040516001600160401b039091168152f35b503461023b578060031936011261023b57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461023b57602036600319011261023b576001600160a01b0319613727613f53565b1680825260026020526040822080546001600160a01b0316156109e157600281019060ff8254169060068210156135f757600182149182613804575b82156137b2575b5050156137a357805460ff191660041790557f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be8280a280f35b63268dbf6760e21b8352600483fd5b6002149150816137ea575b816137cb575b505f8061376a565b905061ffff600181600584015460101c1692015460201c16115f6137c3565b600181015460801c6001600160401b0316431191506137bd565b600182015460401c6001600160401b031643119250613763565b346117e05760c03660031901126117e057613837613f53565b61383f613f6a565b604435916064356001600160401b0381116117e057613862903690600401613f7a565b6084356001600160401b0381116117e057613881903690600401613f7a565b909390919060a4356001600160401b0381116117e0576138a5903690600401613f7a565b969096323314801590613e7a575b613e6b576001600160a01b031983165f81815260026020526040902080549198909390916001600160a01b031615613e5c5760ff6002850154166006811015613e4857600303613e395760ff169860088a1015613db75760018a1b95895f52600d60205260ff8760405f20541616613e2a575f9b6118008203613db75761393c8584018461411c565b9687519060a01c14801590613e14575b8015613dfb575b8015613de2575b8015613dd4575b8015613dc6575b613db7576139865f5160206146775f395f51905f529236908c614170565b6020815191012060405190602082019283526040820152604081526139ac6060826140d2565b51902060405160208101918c83527fae031fc261aed61242596185b006e57bdcba774d6ea39d3348a9a570b38d9ff4602c830152604c820152604c81526139f4606c826140d2565b51902006968760c087015103613db7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156117e0575f93613a5860405196879586948594635c73957b60e11b8652600486016141d5565b03915afa8015613dac57613d97575b506104008501851161088357611000850185116108835761ffff600181600584015460101c1692015460101c16908890895b60208110613ca9575050506104009060405191613ab681846140d2565b368337885b60208110613c2057505084602089905b60058210613ba157505060c0613ae59160e09351966143f6565b91015103613b9257610800830180841161150d579160ff7f66cc041d38da90dceb4c4eb68068985126297a2212bd69d65b5b75e2415f729694926108206040953594013592878a52600b602052858a20895f52602052836001875f208781550155878a52600c602052858a20895f52602052855f2055868952600d6020521660ff848920541617858852600d60205260ff848920911660ff1982541617905582519182526020820152a380f35b63d1fed5fd60e01b8652600486fd5b60011c9150895b828110613bba57506001018691613acb565b8060011b8181046002148215171561086f57613bd681866143e5565b519060018101809111613c0c5760019291613bf4613bfb92886143e5565b51906145aa565b613c0582876143e5565b5201613ba8565b634e487b7160e01b8d52601160045260248dfd5b6110008160061b88010160208135910135908383105f14613c5a5760019291613c48916144c7565b613c5282866143e5565b525b01613abb565b1590811591613c9d575b506120db57807f11a12e535a08d28aa7434e11614f2eb9b34da3fcba5746a376ed981855fb01f0613c97600193866143e5565b52613c54565b6001915014158b613c64565b8060051b88019261040084359401359083831015613d755784158015613d6c575b6120f9576001851b9081811661157f5717938a8d52600460205260408d205f1982019082821161082257613d018f92604092614282565b60018060a01b0391549060031b1c16918d81526005602052209060018060a01b03165f5260205260405f209060ff60038301541615908115613d56575b506120f9576001015403611530576001905b01613a99565b905061ffff80835460a01c16911614158e613d3e565b50858511613cca565b931590811591613d8d575b5061153057600190613d50565b905015158c613d80565b613da49198505f906140d2565b5f9688613a67565b6040513d5f823e3d90fd5b63d1fed5fd60e01b5f5260045ffd5b508060a08801511415613968565b508b60808801511415613961565b50606087015161ffff600588015460101c16141561395a565b50604087015161ffff600188015460101c161415613953565b50602087015161ffff600188015416141561394c565b63615943bb60e11b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b6346f4d6c560e11b5f5260045ffd5b50333b15156138b3565b346117e0575f3660031901126117e05760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015613dac575f906117b557602090604051908152f35b346117e0575f3660031901126117e0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346117e0575f3660031901126117e0575f5460401c6001600160401b03168152602090f35b600435906001600160a01b0319821682036117e057565b6024359060ff821682036117e057565b9181601f840112156117e0578235916001600160401b0383116117e057602083818601950101116117e057565b6044359061ffff821682036117e057565b6064359061ffff821682036117e057565b6024359061ffff821682036117e057565b60609060031901126117e0576004356001600160a01b0319811681036117e057906024359060443561ffff811681036117e05790565b346117e0575f3660031901126117e05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60e081019081106001600160401b0382111761319157604052565b61018081019081106001600160401b0382111761319157604052565b606081019081106001600160401b0382111761319157604052565b608081019081106001600160401b0382111761319157604052565b60a081019081106001600160401b0382111761319157604052565b601f909101601f19168101906001600160401b0382119082101761319157604052565b8181029291811591840414171561410857565b634e487b7160e01b5f52601160045260245ffd5b90610100828203126117e05780601f830112156117e05760405191614143610100846140d2565b829061010081019283116117e057905b8282106141605750505090565b8135815260209182019101614153565b9192916001600160401b0382116131915760405191614199601f8201601f1916602001846140d2565b8294818452818301116117e0578281602093845f960137010152565b908060209392818452848401375f828201840152601f01601f1916010190565b92906141ee906141fc95936040865260408601916141b5565b9260208185039101526141b5565b90565b6001600160401b03918216908216019190821161410857565b5f5460401c6001600160401b03168015614260576141fc907f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316906141ff565b50436001600160401b031690565b61ffff5f199116019061ffff821161410857565b8054821015614297575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b03199091168152602081019190915260400190565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b51906001600160401b03821682036117e057565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b61ffff60019116019061ffff821161410857565b91606093918352602083015260408201520190565b908160c09103126117e0576040519060c082016001600160401b038111838210176131915760405280516001600160a01b03811681036117e05782526020810151602083015260408101516040830152606081015160038110156117e0576143dd9160a09160608501526143d2608082016142e6565b6080850152016142e6565b60a082015290565b9060208110156142975760051b0190565b92915f5160206146775f395f51905f525f940691600192809260051b8201915b82811061443757505050501561442857565b63331835e560e21b5f5260045ffd5b909192955f5160206146775f395f51905f52838160209381863599818b1016998c099008980993929101614416565b6001600160a01b0319165f908152600f60209081526040808320938352929052205460ff16801561449a575f190160ff1690565b633b948c2560e11b5f5260045ffd5b91608093916040519384526020840152604083015260608201522090565b604191604051915f8353600183015260218201522090565b60401b63ffffffff60401b166001600160401b039091161760a01b6001600160a01b03191690565b905f5160206146775f395f51905f52821080614558575b1561453f5781158061454e575b61453f57614538916145c3565b1561453f57565b634c4d29cd60e11b5f5260045ffd5b506001811461452b565b505f5160206146775f395f51905f52811061451e565b6001600160401b0381116131915760051b60200190565b9060408110156142975760051b0190565b80518210156142975760209160051b010190565b6041916040519160018353600183015260218201522090565b5f5160206146775f395f51905f52811080159061465f575b614659575f5160206146775f395f51905f528181920991800990805f5160206146775f395f51905f5203915f5160206146775f395f51905f528311614108575f5160206146775f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b50505f90565b505f5160206146775f395f51905f528210156145db56fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212208de39e7443a5eb1df14dcabc627df3ccc8c6a4f975231e697bc936f670b9add464736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _poolKeyVerifier common.Address, _decryptCombineVerifier common.Address, _epochDurationBlocks *big.Int, _committeeSelectionBlocks *big.Int, _keyAssemblyBlocks *big.Int, _finalizeGapBlocks *big.Int, _minThreshold uint16, _minCommitteeSize uint16, _maxLotteryAlphaBps uint16) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _poolKeyVerifier, _decryptCombineVerifier, _epochDurationBlocks, _committeeSelectionBlocks, _keyAssemblyBlocks, _finalizeGapBlocks, _minThreshold, _minCommitteeSize, _maxLotteryAlphaBps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKGManager{DKGManagerCaller: DKGManagerCaller{contract: contract}, DKGManagerTransactor: DKGManagerTransactor{contract: contract}, DKGManagerFilterer: DKGManagerFilterer{contract: contract}}, nil
}

// DKGManager is an auto generated Go binding around an Ethereum contract.
type DKGManager struct {
	DKGManagerCaller     // Read-only binding to the contract
	DKGManagerTransactor // Write-only binding to the contract
	DKGManagerFilterer   // Log filterer for contract events
}

// DKGManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGManagerSession struct {
	Contract     *DKGManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGManagerCallerSession struct {
	Contract *DKGManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DKGManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGManagerTransactorSession struct {
	Contract     *DKGManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DKGManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGManagerRaw struct {
	Contract *DKGManager // Generic contract binding to access the raw methods on
}

// DKGManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGManagerCallerRaw struct {
	Contract *DKGManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DKGManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGManagerTransactorRaw struct {
	Contract *DKGManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKGManager creates a new instance of DKGManager, bound to a specific deployed contract.
func NewDKGManager(address common.Address, backend bind.ContractBackend) (*DKGManager, error) {
	contract, err := bindDKGManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKGManager{DKGManagerCaller: DKGManagerCaller{contract: contract}, DKGManagerTransactor: DKGManagerTransactor{contract: contract}, DKGManagerFilterer: DKGManagerFilterer{contract: contract}}, nil
}

// NewDKGManagerCaller creates a new read-only instance of DKGManager, bound to a specific deployed contract.
func NewDKGManagerCaller(address common.Address, caller bind.ContractCaller) (*DKGManagerCaller, error) {
	contract, err := bindDKGManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCaller{contract: contract}, nil
}

// NewDKGManagerTransactor creates a new write-only instance of DKGManager, bound to a specific deployed contract.
func NewDKGManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGManagerTransactor, error) {
	contract, err := bindDKGManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGManagerTransactor{contract: contract}, nil
}

// NewDKGManagerFilterer creates a new log filterer instance of DKGManager, bound to a specific deployed contract.
func NewDKGManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGManagerFilterer, error) {
	contract, err := bindDKGManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGManagerFilterer{contract: contract}, nil
}

// bindDKGManager binds a generic wrapper to an already deployed contract.
func bindDKGManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGManager *DKGManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGManager.Contract.DKGManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGManager *DKGManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGManager.Contract.DKGManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGManager *DKGManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGManager.Contract.DKGManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGManager *DKGManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGManager *DKGManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGManager *DKGManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGManager.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint32)
func (_DKGManager *DKGManagerCaller) CHAINID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint32)
func (_DKGManager *DKGManagerSession) CHAINID() (uint32, error) {
	return _DKGManager.Contract.CHAINID(&_DKGManager.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint32)
func (_DKGManager *DKGManagerCallerSession) CHAINID() (uint32, error) {
	return _DKGManager.Contract.CHAINID(&_DKGManager.CallOpts)
}

// CONTRIBUTIONVERIFIER is a free data retrieval call binding the contract method 0x63f314cd.
//
// Solidity: function CONTRIBUTION_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) CONTRIBUTIONVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "CONTRIBUTION_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTRIBUTIONVERIFIER is a free data retrieval call binding the contract method 0x63f314cd.
//
// Solidity: function CONTRIBUTION_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) CONTRIBUTIONVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.CONTRIBUTIONVERIFIER(&_DKGManager.CallOpts)
}

// CONTRIBUTIONVERIFIER is a free data retrieval call binding the contract method 0x63f314cd.
//
// Solidity: function CONTRIBUTION_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) CONTRIBUTIONVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.CONTRIBUTIONVERIFIER(&_DKGManager.CallOpts)
}

// DECRYPTCOMBINEVERIFIER is a free data retrieval call binding the contract method 0xfe1604b5.
//
// Solidity: function DECRYPT_COMBINE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) DECRYPTCOMBINEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "DECRYPT_COMBINE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DECRYPTCOMBINEVERIFIER is a free data retrieval call binding the contract method 0xfe1604b5.
//
// Solidity: function DECRYPT_COMBINE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) DECRYPTCOMBINEVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.DECRYPTCOMBINEVERIFIER(&_DKGManager.CallOpts)
}

// DECRYPTCOMBINEVERIFIER is a free data retrieval call binding the contract method 0xfe1604b5.
//
// Solidity: function DECRYPT_COMBINE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) DECRYPTCOMBINEVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.DECRYPTCOMBINEVERIFIER(&_DKGManager.CallOpts)
}

// EPOCHDURATIONBLOCKS is a free data retrieval call binding the contract method 0xfa8f5e96.
//
// Solidity: function EPOCH_DURATION_BLOCKS() view returns(uint256)
func (_DKGManager *DKGManagerCaller) EPOCHDURATIONBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "EPOCH_DURATION_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHDURATIONBLOCKS is a free data retrieval call binding the contract method 0xfa8f5e96.
//
// Solidity: function EPOCH_DURATION_BLOCKS() view returns(uint256)
func (_DKGManager *DKGManagerSession) EPOCHDURATIONBLOCKS() (*big.Int, error) {
	return _DKGManager.Contract.EPOCHDURATIONBLOCKS(&_DKGManager.CallOpts)
}

// EPOCHDURATIONBLOCKS is a free data retrieval call binding the contract method 0xfa8f5e96.
//
// Solidity: function EPOCH_DURATION_BLOCKS() view returns(uint256)
func (_DKGManager *DKGManagerCallerSession) EPOCHDURATIONBLOCKS() (*big.Int, error) {
	return _DKGManager.Contract.EPOCHDURATIONBLOCKS(&_DKGManager.CallOpts)
}

// EPOCHPREFIX is a free data retrieval call binding the contract method 0x23488be5.
//
// Solidity: function EPOCH_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerCaller) EPOCHPREFIX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "EPOCH_PREFIX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// EPOCHPREFIX is a free data retrieval call binding the contract method 0x23488be5.
//
// Solidity: function EPOCH_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerSession) EPOCHPREFIX() (uint32, error) {
	return _DKGManager.Contract.EPOCHPREFIX(&_DKGManager.CallOpts)
}

// EPOCHPREFIX is a free data retrieval call binding the contract method 0x23488be5.
//
// Solidity: function EPOCH_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerCallerSession) EPOCHPREFIX() (uint32, error) {
	return _DKGManager.Contract.EPOCHPREFIX(&_DKGManager.CallOpts)
}

// MAXLOTTERYALPHABPS is a free data retrieval call binding the contract method 0xd9e9ca2e.
//
// Solidity: function MAX_LOTTERY_ALPHA_BPS() view returns(uint16)
func (_DKGManager *DKGManagerCaller) MAXLOTTERYALPHABPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "MAX_LOTTERY_ALPHA_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXLOTTERYALPHABPS is a free data retrieval call binding the contract method 0xd9e9ca2e.
//
// Solidity: function MAX_LOTTERY_ALPHA_BPS() view returns(uint16)
func (_DKGManager *DKGManagerSession) MAXLOTTERYALPHABPS() (uint16, error) {
	return _DKGManager.Contract.MAXLOTTERYALPHABPS(&_DKGManager.CallOpts)
}

// MAXLOTTERYALPHABPS is a free data retrieval call binding the contract method 0xd9e9ca2e.
//
// Solidity: function MAX_LOTTERY_ALPHA_BPS() view returns(uint16)
func (_DKGManager *DKGManagerCallerSession) MAXLOTTERYALPHABPS() (uint16, error) {
	return _DKGManager.Contract.MAXLOTTERYALPHABPS(&_DKGManager.CallOpts)
}

// MINCOMMITTEESIZE is a free data retrieval call binding the contract method 0xf03a4898.
//
// Solidity: function MIN_COMMITTEE_SIZE() view returns(uint16)
func (_DKGManager *DKGManagerCaller) MINCOMMITTEESIZE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "MIN_COMMITTEE_SIZE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MINCOMMITTEESIZE is a free data retrieval call binding the contract method 0xf03a4898.
//
// Solidity: function MIN_COMMITTEE_SIZE() view returns(uint16)
func (_DKGManager *DKGManagerSession) MINCOMMITTEESIZE() (uint16, error) {
	return _DKGManager.Contract.MINCOMMITTEESIZE(&_DKGManager.CallOpts)
}

// MINCOMMITTEESIZE is a free data retrieval call binding the contract method 0xf03a4898.
//
// Solidity: function MIN_COMMITTEE_SIZE() view returns(uint16)
func (_DKGManager *DKGManagerCallerSession) MINCOMMITTEESIZE() (uint16, error) {
	return _DKGManager.Contract.MINCOMMITTEESIZE(&_DKGManager.CallOpts)
}

// MINTHRESHOLD is a free data retrieval call binding the contract method 0xbd11c4c0.
//
// Solidity: function MIN_THRESHOLD() view returns(uint16)
func (_DKGManager *DKGManagerCaller) MINTHRESHOLD(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "MIN_THRESHOLD")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MINTHRESHOLD is a free data retrieval call binding the contract method 0xbd11c4c0.
//
// Solidity: function MIN_THRESHOLD() view returns(uint16)
func (_DKGManager *DKGManagerSession) MINTHRESHOLD() (uint16, error) {
	return _DKGManager.Contract.MINTHRESHOLD(&_DKGManager.CallOpts)
}

// MINTHRESHOLD is a free data retrieval call binding the contract method 0xbd11c4c0.
//
// Solidity: function MIN_THRESHOLD() view returns(uint16)
func (_DKGManager *DKGManagerCallerSession) MINTHRESHOLD() (uint16, error) {
	return _DKGManager.Contract.MINTHRESHOLD(&_DKGManager.CallOpts)
}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) PARTIALDECRYPTVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "PARTIAL_DECRYPT_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) PARTIALDECRYPTVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.PARTIALDECRYPTVERIFIER(&_DKGManager.CallOpts)
}

// PARTIALDECRYPTVERIFIER is a free data retrieval call binding the contract method 0xbf192209.
//
// Solidity: function PARTIAL_DECRYPT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) PARTIALDECRYPTVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.PARTIALDECRYPTVERIFIER(&_DKGManager.CallOpts)
}

// POOLKEYVERIFIER is a free data retrieval call binding the contract method 0x2b743f04.
//
// Solidity: function POOL_KEY_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) POOLKEYVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "POOL_KEY_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOLKEYVERIFIER is a free data retrieval call binding the contract method 0x2b743f04.
//
// Solidity: function POOL_KEY_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) POOLKEYVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.POOLKEYVERIFIER(&_DKGManager.CallOpts)
}

// POOLKEYVERIFIER is a free data retrieval call binding the contract method 0x2b743f04.
//
// Solidity: function POOL_KEY_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) POOLKEYVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.POOLKEYVERIFIER(&_DKGManager.CallOpts)
}

// REGISTRY is a free data retrieval call binding the contract method 0x06433b1b.
//
// Solidity: function REGISTRY() view returns(address)
func (_DKGManager *DKGManagerCaller) REGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REGISTRY is a free data retrieval call binding the contract method 0x06433b1b.
//
// Solidity: function REGISTRY() view returns(address)
func (_DKGManager *DKGManagerSession) REGISTRY() (common.Address, error) {
	return _DKGManager.Contract.REGISTRY(&_DKGManager.CallOpts)
}

// REGISTRY is a free data retrieval call binding the contract method 0x06433b1b.
//
// Solidity: function REGISTRY() view returns(address)
func (_DKGManager *DKGManagerCallerSession) REGISTRY() (common.Address, error) {
	return _DKGManager.Contract.REGISTRY(&_DKGManager.CallOpts)
}

// AppManager is a free data retrieval call binding the contract method 0xebe86c13.
//
// Solidity: function appManager() view returns(address)
func (_DKGManager *DKGManagerCaller) AppManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "appManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppManager is a free data retrieval call binding the contract method 0xebe86c13.
//
// Solidity: function appManager() view returns(address)
func (_DKGManager *DKGManagerSession) AppManager() (common.Address, error) {
	return _DKGManager.Contract.AppManager(&_DKGManager.CallOpts)
}

// AppManager is a free data retrieval call binding the contract method 0xebe86c13.
//
// Solidity: function appManager() view returns(address)
func (_DKGManager *DKGManagerCallerSession) AppManager() (common.Address, error) {
	return _DKGManager.Contract.AppManager(&_DKGManager.CallOpts)
}

// CiphertextCount is a free data retrieval call binding the contract method 0x6f067f63.
//
// Solidity: function ciphertextCount(bytes12 epochId, bytes32 aid) view returns(uint16)
func (_DKGManager *DKGManagerCaller) CiphertextCount(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (uint16, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "ciphertextCount", epochId, aid)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CiphertextCount is a free data retrieval call binding the contract method 0x6f067f63.
//
// Solidity: function ciphertextCount(bytes12 epochId, bytes32 aid) view returns(uint16)
func (_DKGManager *DKGManagerSession) CiphertextCount(epochId [12]byte, aid [32]byte) (uint16, error) {
	return _DKGManager.Contract.CiphertextCount(&_DKGManager.CallOpts, epochId, aid)
}

// CiphertextCount is a free data retrieval call binding the contract method 0x6f067f63.
//
// Solidity: function ciphertextCount(bytes12 epochId, bytes32 aid) view returns(uint16)
func (_DKGManager *DKGManagerCallerSession) CiphertextCount(epochId [12]byte, aid [32]byte) (uint16, error) {
	return _DKGManager.Contract.CiphertextCount(&_DKGManager.CallOpts, epochId, aid)
}

// EpochDurationBlocks is a free data retrieval call binding the contract method 0x71712c29.
//
// Solidity: function epochDurationBlocks() view returns(uint256)
func (_DKGManager *DKGManagerCaller) EpochDurationBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "epochDurationBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDurationBlocks is a free data retrieval call binding the contract method 0x71712c29.
//
// Solidity: function epochDurationBlocks() view returns(uint256)
func (_DKGManager *DKGManagerSession) EpochDurationBlocks() (*big.Int, error) {
	return _DKGManager.Contract.EpochDurationBlocks(&_DKGManager.CallOpts)
}

// EpochDurationBlocks is a free data retrieval call binding the contract method 0x71712c29.
//
// Solidity: function epochDurationBlocks() view returns(uint256)
func (_DKGManager *DKGManagerCallerSession) EpochDurationBlocks() (*big.Int, error) {
	return _DKGManager.Contract.EpochDurationBlocks(&_DKGManager.CallOpts)
}

// EpochNonce is a free data retrieval call binding the contract method 0xa4adcd7f.
//
// Solidity: function epochNonce() view returns(uint64)
func (_DKGManager *DKGManagerCaller) EpochNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "epochNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EpochNonce is a free data retrieval call binding the contract method 0xa4adcd7f.
//
// Solidity: function epochNonce() view returns(uint64)
func (_DKGManager *DKGManagerSession) EpochNonce() (uint64, error) {
	return _DKGManager.Contract.EpochNonce(&_DKGManager.CallOpts)
}

// EpochNonce is a free data retrieval call binding the contract method 0xa4adcd7f.
//
// Solidity: function epochNonce() view returns(uint64)
func (_DKGManager *DKGManagerCallerSession) EpochNonce() (uint64, error) {
	return _DKGManager.Contract.EpochNonce(&_DKGManager.CallOpts)
}

// GetAppPoolIndex is a free data retrieval call binding the contract method 0x368e2a27.
//
// Solidity: function getAppPoolIndex(bytes12 epochId, bytes32 aid) view returns(uint8)
func (_DKGManager *DKGManagerCaller) GetAppPoolIndex(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (uint8, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getAppPoolIndex", epochId, aid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAppPoolIndex is a free data retrieval call binding the contract method 0x368e2a27.
//
// Solidity: function getAppPoolIndex(bytes12 epochId, bytes32 aid) view returns(uint8)
func (_DKGManager *DKGManagerSession) GetAppPoolIndex(epochId [12]byte, aid [32]byte) (uint8, error) {
	return _DKGManager.Contract.GetAppPoolIndex(&_DKGManager.CallOpts, epochId, aid)
}

// GetAppPoolIndex is a free data retrieval call binding the contract method 0x368e2a27.
//
// Solidity: function getAppPoolIndex(bytes12 epochId, bytes32 aid) view returns(uint8)
func (_DKGManager *DKGManagerCallerSession) GetAppPoolIndex(epochId [12]byte, aid [32]byte) (uint8, error) {
	return _DKGManager.Contract.GetAppPoolIndex(&_DKGManager.CallOpts, epochId, aid)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x2de546d5.
//
// Solidity: function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetCiphertextHash(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCiphertextHash", epochId, aid, ciphertextIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x2de546d5.
//
// Solidity: function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetCiphertextHash(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x2de546d5.
//
// Solidity: function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetCiphertextHash(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0x9bbada67.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCaller) GetCombinedDecryption(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCombinedDecryption", epochId, aid, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesCombinedDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesCombinedDecryptionRecord)).(*DKGTypesCombinedDecryptionRecord)

	return out0, err

}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0x9bbada67.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerSession) GetCombinedDecryption(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0x9bbada67.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCallerSession) GetCombinedDecryption(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerCaller) GetContribution(opts *bind.CallOpts, epochId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getContribution", epochId, contributor)

	if err != nil {
		return *new(DKGTypesContributionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesContributionRecord)).(*DKGTypesContributionRecord)

	return out0, err

}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetContribution(epochId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	return _DKGManager.Contract.GetContribution(&_DKGManager.CallOpts, epochId, contributor)
}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerCallerSession) GetContribution(epochId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	return _DKGManager.Contract.GetContribution(&_DKGManager.CallOpts, epochId, contributor)
}

// GetContributionVerifierVKeyHash is a free data retrieval call binding the contract method 0x074a75e1.
//
// Solidity: function getContributionVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetContributionVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getContributionVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetContributionVerifierVKeyHash is a free data retrieval call binding the contract method 0x074a75e1.
//
// Solidity: function getContributionVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetContributionVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetContributionVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetContributionVerifierVKeyHash is a free data retrieval call binding the contract method 0x074a75e1.
//
// Solidity: function getContributionVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetContributionVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetContributionVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetDecryptCombineVerifierVKeyHash is a free data retrieval call binding the contract method 0x72517b4b.
//
// Solidity: function getDecryptCombineVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetDecryptCombineVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getDecryptCombineVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDecryptCombineVerifierVKeyHash is a free data retrieval call binding the contract method 0x72517b4b.
//
// Solidity: function getDecryptCombineVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetDecryptCombineVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetDecryptCombineVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetDecryptCombineVerifierVKeyHash is a free data retrieval call binding the contract method 0x72517b4b.
//
// Solidity: function getDecryptCombineVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetDecryptCombineVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetDecryptCombineVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerCaller) GetEpoch(opts *bind.CallOpts, epochId [12]byte) (IDKGManagerEpoch, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getEpoch", epochId)

	if err != nil {
		return *new(IDKGManagerEpoch), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGManagerEpoch)).(*IDKGManagerEpoch)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetEpoch(epochId [12]byte) (IDKGManagerEpoch, error) {
	return _DKGManager.Contract.GetEpoch(&_DKGManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerCallerSession) GetEpoch(epochId [12]byte) (IDKGManagerEpoch, error) {
	return _DKGManager.Contract.GetEpoch(&_DKGManager.CallOpts, epochId)
}

// GetPartialDecryptVerifierVKeyHash is a free data retrieval call binding the contract method 0x8dc1f53a.
//
// Solidity: function getPartialDecryptVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetPartialDecryptVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPartialDecryptVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPartialDecryptVerifierVKeyHash is a free data retrieval call binding the contract method 0x8dc1f53a.
//
// Solidity: function getPartialDecryptVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetPartialDecryptVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetPartialDecryptVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetPartialDecryptVerifierVKeyHash is a free data retrieval call binding the contract method 0x8dc1f53a.
//
// Solidity: function getPartialDecryptVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetPartialDecryptVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetPartialDecryptVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetPartialDecryption is a free data retrieval call binding the contract method 0xbea5210d.
//
// Solidity: function getPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex) view returns((uint16,uint16,bytes32,bool))
func (_DKGManager *DKGManagerCaller) GetPartialDecryption(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPartialDecryption", epochId, aid, participantIndex, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesPartialDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesPartialDecryptionRecord)).(*DKGTypesPartialDecryptionRecord)

	return out0, err

}

// GetPartialDecryption is a free data retrieval call binding the contract method 0xbea5210d.
//
// Solidity: function getPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex) view returns((uint16,uint16,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, epochId, aid, participantIndex, ciphertextIndex)
}

// GetPartialDecryption is a free data retrieval call binding the contract method 0xbea5210d.
//
// Solidity: function getPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex) view returns((uint16,uint16,bytes32,bool))
func (_DKGManager *DKGManagerCallerSession) GetPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, epochId, aid, participantIndex, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x5a8f2bb3.
//
// Solidity: function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCaller) GetPlaintext(opts *bind.CallOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (*big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPlaintext", epochId, aid, ciphertextIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlaintext is a free data retrieval call binding the contract method 0x5a8f2bb3.
//
// Solidity: function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerSession) GetPlaintext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x5a8f2bb3.
//
// Solidity: function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCallerSession) GetPlaintext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, epochId, aid, ciphertextIndex)
}

// GetPoolKey is a free data retrieval call binding the contract method 0x56cbb5f3.
//
// Solidity: function getPoolKey(bytes12 epochId, uint8 keyIndex) view returns(uint256, uint256)
func (_DKGManager *DKGManagerCaller) GetPoolKey(opts *bind.CallOpts, epochId [12]byte, keyIndex uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPoolKey", epochId, keyIndex)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPoolKey is a free data retrieval call binding the contract method 0x56cbb5f3.
//
// Solidity: function getPoolKey(bytes12 epochId, uint8 keyIndex) view returns(uint256, uint256)
func (_DKGManager *DKGManagerSession) GetPoolKey(epochId [12]byte, keyIndex uint8) (*big.Int, *big.Int, error) {
	return _DKGManager.Contract.GetPoolKey(&_DKGManager.CallOpts, epochId, keyIndex)
}

// GetPoolKey is a free data retrieval call binding the contract method 0x56cbb5f3.
//
// Solidity: function getPoolKey(bytes12 epochId, uint8 keyIndex) view returns(uint256, uint256)
func (_DKGManager *DKGManagerCallerSession) GetPoolKey(epochId [12]byte, keyIndex uint8) (*big.Int, *big.Int, error) {
	return _DKGManager.Contract.GetPoolKey(&_DKGManager.CallOpts, epochId, keyIndex)
}

// GetPoolKeyVerifierVKeyHash is a free data retrieval call binding the contract method 0x43e50afa.
//
// Solidity: function getPoolKeyVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetPoolKeyVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPoolKeyVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolKeyVerifierVKeyHash is a free data retrieval call binding the contract method 0x43e50afa.
//
// Solidity: function getPoolKeyVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetPoolKeyVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetPoolKeyVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetPoolKeyVerifierVKeyHash is a free data retrieval call binding the contract method 0x43e50afa.
//
// Solidity: function getPoolKeyVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetPoolKeyVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetPoolKeyVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetPoolShareRoot is a free data retrieval call binding the contract method 0x7ade1324.
//
// Solidity: function getPoolShareRoot(bytes12 epochId, uint8 keyIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetPoolShareRoot(opts *bind.CallOpts, epochId [12]byte, keyIndex uint8) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPoolShareRoot", epochId, keyIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolShareRoot is a free data retrieval call binding the contract method 0x7ade1324.
//
// Solidity: function getPoolShareRoot(bytes12 epochId, uint8 keyIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetPoolShareRoot(epochId [12]byte, keyIndex uint8) ([32]byte, error) {
	return _DKGManager.Contract.GetPoolShareRoot(&_DKGManager.CallOpts, epochId, keyIndex)
}

// GetPoolShareRoot is a free data retrieval call binding the contract method 0x7ade1324.
//
// Solidity: function getPoolShareRoot(bytes12 epochId, uint8 keyIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetPoolShareRoot(epochId [12]byte, keyIndex uint8) ([32]byte, error) {
	return _DKGManager.Contract.GetPoolShareRoot(&_DKGManager.CallOpts, epochId, keyIndex)
}

// GetPoolStatus is a free data retrieval call binding the contract method 0xd3979253.
//
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex, uint8 activated)
func (_DKGManager *DKGManagerCaller) GetPoolStatus(opts *bind.CallOpts, epochId [12]byte) (struct {
	NextIndex uint8
	Activated uint8
}, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPoolStatus", epochId)

	outstruct := new(struct {
		NextIndex uint8
		Activated uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NextIndex = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Activated = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetPoolStatus is a free data retrieval call binding the contract method 0xd3979253.
//
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex, uint8 activated)
func (_DKGManager *DKGManagerSession) GetPoolStatus(epochId [12]byte) (struct {
	NextIndex uint8
	Activated uint8
}, error) {
	return _DKGManager.Contract.GetPoolStatus(&_DKGManager.CallOpts, epochId)
}

// GetPoolStatus is a free data retrieval call binding the contract method 0xd3979253.
//
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex, uint8 activated)
func (_DKGManager *DKGManagerCallerSession) GetPoolStatus(epochId [12]byte) (struct {
	NextIndex uint8
	Activated uint8
}, error) {
	return _DKGManager.Contract.GetPoolStatus(&_DKGManager.CallOpts, epochId)
}

// LastEpochStartBlock is a free data retrieval call binding the contract method 0x04da5740.
//
// Solidity: function lastEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerCaller) LastEpochStartBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "lastEpochStartBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastEpochStartBlock is a free data retrieval call binding the contract method 0x04da5740.
//
// Solidity: function lastEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerSession) LastEpochStartBlock() (uint64, error) {
	return _DKGManager.Contract.LastEpochStartBlock(&_DKGManager.CallOpts)
}

// LastEpochStartBlock is a free data retrieval call binding the contract method 0x04da5740.
//
// Solidity: function lastEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerCallerSession) LastEpochStartBlock() (uint64, error) {
	return _DKGManager.Contract.LastEpochStartBlock(&_DKGManager.CallOpts)
}

// NextEpochStartBlock is a free data retrieval call binding the contract method 0x268ae2a1.
//
// Solidity: function nextEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerCaller) NextEpochStartBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "nextEpochStartBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextEpochStartBlock is a free data retrieval call binding the contract method 0x268ae2a1.
//
// Solidity: function nextEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerSession) NextEpochStartBlock() (uint64, error) {
	return _DKGManager.Contract.NextEpochStartBlock(&_DKGManager.CallOpts)
}

// NextEpochStartBlock is a free data retrieval call binding the contract method 0x268ae2a1.
//
// Solidity: function nextEpochStartBlock() view returns(uint64)
func (_DKGManager *DKGManagerCallerSession) NextEpochStartBlock() (uint64, error) {
	return _DKGManager.Contract.NextEpochStartBlock(&_DKGManager.CallOpts)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 epochId) view returns(address[])
func (_DKGManager *DKGManagerCaller) SelectedParticipants(opts *bind.CallOpts, epochId [12]byte) ([]common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "selectedParticipants", epochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 epochId) view returns(address[])
func (_DKGManager *DKGManagerSession) SelectedParticipants(epochId [12]byte) ([]common.Address, error) {
	return _DKGManager.Contract.SelectedParticipants(&_DKGManager.CallOpts, epochId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 epochId) view returns(address[])
func (_DKGManager *DKGManagerCallerSession) SelectedParticipants(epochId [12]byte) ([]common.Address, error) {
	return _DKGManager.Contract.SelectedParticipants(&_DKGManager.CallOpts, epochId)
}

// AbortEpoch is a paid mutator transaction binding the contract method 0x18287e5f.
//
// Solidity: function abortEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactor) AbortEpoch(opts *bind.TransactOpts, epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "abortEpoch", epochId)
}

// AbortEpoch is a paid mutator transaction binding the contract method 0x18287e5f.
//
// Solidity: function abortEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerSession) AbortEpoch(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.AbortEpoch(&_DKGManager.TransactOpts, epochId)
}

// AbortEpoch is a paid mutator transaction binding the contract method 0x18287e5f.
//
// Solidity: function abortEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactorSession) AbortEpoch(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.AbortEpoch(&_DKGManager.TransactOpts, epochId)
}

// ActivatePoolKey is a paid mutator transaction binding the contract method 0x134c98e0.
//
// Solidity: function activatePoolKey(bytes12 epochId, uint8 keyIndex, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) ActivatePoolKey(opts *bind.TransactOpts, epochId [12]byte, keyIndex uint8, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "activatePoolKey", epochId, keyIndex, transcriptDigest, transcript, proof, input)
}

// ActivatePoolKey is a paid mutator transaction binding the contract method 0x134c98e0.
//
// Solidity: function activatePoolKey(bytes12 epochId, uint8 keyIndex, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) ActivatePoolKey(epochId [12]byte, keyIndex uint8, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ActivatePoolKey(&_DKGManager.TransactOpts, epochId, keyIndex, transcriptDigest, transcript, proof, input)
}

// ActivatePoolKey is a paid mutator transaction binding the contract method 0x134c98e0.
//
// Solidity: function activatePoolKey(bytes12 epochId, uint8 keyIndex, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) ActivatePoolKey(epochId [12]byte, keyIndex uint8, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ActivatePoolKey(&_DKGManager.TransactOpts, epochId, keyIndex, transcriptDigest, transcript, proof, input)
}

// ClaimPoolKey is a paid mutator transaction binding the contract method 0x421adfbb.
//
// Solidity: function claimPoolKey(bytes12 epochId, bytes32 aid) returns(uint8 keyIndex)
func (_DKGManager *DKGManagerTransactor) ClaimPoolKey(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "claimPoolKey", epochId, aid)
}

// ClaimPoolKey is a paid mutator transaction binding the contract method 0x421adfbb.
//
// Solidity: function claimPoolKey(bytes12 epochId, bytes32 aid) returns(uint8 keyIndex)
func (_DKGManager *DKGManagerSession) ClaimPoolKey(epochId [12]byte, aid [32]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimPoolKey(&_DKGManager.TransactOpts, epochId, aid)
}

// ClaimPoolKey is a paid mutator transaction binding the contract method 0x421adfbb.
//
// Solidity: function claimPoolKey(bytes12 epochId, bytes32 aid) returns(uint8 keyIndex)
func (_DKGManager *DKGManagerTransactorSession) ClaimPoolKey(epochId [12]byte, aid [32]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimPoolKey(&_DKGManager.TransactOpts, epochId, aid)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactor) ClaimSlot(opts *bind.TransactOpts, epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "claimSlot", epochId)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 epochId) returns()
func (_DKGManager *DKGManagerSession) ClaimSlot(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimSlot(&_DKGManager.TransactOpts, epochId)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactorSession) ClaimSlot(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimSlot(&_DKGManager.TransactOpts, epochId)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0x77235ee1.
//
// Solidity: function combineDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) CombineDecryption(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "combineDecryption", epochId, aid, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0x77235ee1.
//
// Solidity: function combineDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) CombineDecryption(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0x77235ee1.
//
// Solidity: function combineDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) CombineDecryption(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x71a5978c.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateEpoch(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createEpoch", threshold, committeeSize, minValidContributions, lotteryAlphaBps)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x71a5978c.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x71a5978c.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x3681b559.
//
// Solidity: function finalizeEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactor) FinalizeEpoch(opts *bind.TransactOpts, epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "finalizeEpoch", epochId)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x3681b559.
//
// Solidity: function finalizeEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerSession) FinalizeEpoch(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x3681b559.
//
// Solidity: function finalizeEpoch(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactorSession) FinalizeEpoch(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId)
}

// SetAppManager is a paid mutator transaction binding the contract method 0x6d16897d.
//
// Solidity: function setAppManager(address a) returns()
func (_DKGManager *DKGManagerTransactor) SetAppManager(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "setAppManager", a)
}

// SetAppManager is a paid mutator transaction binding the contract method 0x6d16897d.
//
// Solidity: function setAppManager(address a) returns()
func (_DKGManager *DKGManagerSession) SetAppManager(a common.Address) (*types.Transaction, error) {
	return _DKGManager.Contract.SetAppManager(&_DKGManager.TransactOpts, a)
}

// SetAppManager is a paid mutator transaction binding the contract method 0x6d16897d.
//
// Solidity: function setAppManager(address a) returns()
func (_DKGManager *DKGManagerTransactorSession) SetAppManager(a common.Address) (*types.Transaction, error) {
	return _DKGManager.Contract.SetAppManager(&_DKGManager.TransactOpts, a)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x7b31b566.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerTransactor) SubmitCiphertext(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitCiphertext", epochId, aid, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x7b31b566.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x7b31b566.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerTransactorSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, c1x, c1y, c2x, c2y)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitContribution(opts *bind.TransactOpts, epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitContribution", epochId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitContribution(epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, epochId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitContribution(epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, epochId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x5b0c0347.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input, bytes32[] shareProof) returns()
func (_DKGManager *DKGManagerTransactor) SubmitPartialDecryption(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte, shareProof [][32]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitPartialDecryption", epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input, shareProof)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x5b0c0347.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input, bytes32[] shareProof) returns()
func (_DKGManager *DKGManagerSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte, shareProof [][32]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input, shareProof)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x5b0c0347.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input, bytes32[] shareProof) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte, shareProof [][32]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input, shareProof)
}

// DKGManagerCiphertextSubmittedIterator is returned from FilterCiphertextSubmitted and is used to iterate over the raw logs and unpacked data for CiphertextSubmitted events raised by the DKGManager contract.
type DKGManagerCiphertextSubmittedIterator struct {
	Event *DKGManagerCiphertextSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerCiphertextSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerCiphertextSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerCiphertextSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerCiphertextSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerCiphertextSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerCiphertextSubmitted represents a CiphertextSubmitted event raised by the DKGManager contract.
type DKGManagerCiphertextSubmitted struct {
	EpochId         [12]byte
	Aid             [32]byte
	CiphertextIndex uint16
	Submitter       common.Address
	C1x             *big.Int
	C1y             *big.Int
	C2x             *big.Int
	C2y             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCiphertextSubmitted is a free log retrieval operation binding the contract event 0x1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) FilterCiphertextSubmitted(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (*DKGManagerCiphertextSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "CiphertextSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCiphertextSubmittedIterator{contract: _DKGManager.contract, event: "CiphertextSubmitted", logs: logs, sub: sub}, nil
}

// WatchCiphertextSubmitted is a free log subscription operation binding the contract event 0x1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) WatchCiphertextSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerCiphertextSubmitted, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "CiphertextSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerCiphertextSubmitted)
				if err := _DKGManager.contract.UnpackLog(event, "CiphertextSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCiphertextSubmitted is a log parse operation binding the contract event 0x1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) ParseCiphertextSubmitted(log types.Log) (*DKGManagerCiphertextSubmitted, error) {
	event := new(DKGManagerCiphertextSubmitted)
	if err := _DKGManager.contract.UnpackLog(event, "CiphertextSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerCommitteeFilledIterator is returned from FilterCommitteeFilled and is used to iterate over the raw logs and unpacked data for CommitteeFilled events raised by the DKGManager contract.
type DKGManagerCommitteeFilledIterator struct {
	Event *DKGManagerCommitteeFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerCommitteeFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerCommitteeFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerCommitteeFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerCommitteeFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerCommitteeFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerCommitteeFilled represents a CommitteeFilled event raised by the DKGManager contract.
type DKGManagerCommitteeFilled struct {
	EpochId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommitteeFilled is a free log retrieval operation binding the contract event 0x23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb067.
//
// Solidity: event CommitteeFilled(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) FilterCommitteeFilled(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerCommitteeFilledIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "CommitteeFilled", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCommitteeFilledIterator{contract: _DKGManager.contract, event: "CommitteeFilled", logs: logs, sub: sub}, nil
}

// WatchCommitteeFilled is a free log subscription operation binding the contract event 0x23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb067.
//
// Solidity: event CommitteeFilled(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) WatchCommitteeFilled(opts *bind.WatchOpts, sink chan<- *DKGManagerCommitteeFilled, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "CommitteeFilled", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerCommitteeFilled)
				if err := _DKGManager.contract.UnpackLog(event, "CommitteeFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitteeFilled is a log parse operation binding the contract event 0x23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb067.
//
// Solidity: event CommitteeFilled(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) ParseCommitteeFilled(log types.Log) (*DKGManagerCommitteeFilled, error) {
	event := new(DKGManagerCommitteeFilled)
	if err := _DKGManager.contract.UnpackLog(event, "CommitteeFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerCommitteeSnapshotIterator is returned from FilterCommitteeSnapshot and is used to iterate over the raw logs and unpacked data for CommitteeSnapshot events raised by the DKGManager contract.
type DKGManagerCommitteeSnapshotIterator struct {
	Event *DKGManagerCommitteeSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerCommitteeSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerCommitteeSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerCommitteeSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerCommitteeSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerCommitteeSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerCommitteeSnapshot represents a CommitteeSnapshot event raised by the DKGManager contract.
type DKGManagerCommitteeSnapshot struct {
	EpochId       [12]byte
	CommitteeSize uint16
	PubKeys       []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitteeSnapshot is a free log retrieval operation binding the contract event 0x3687c3912698ec8584ed18d7c68916ff5fb9418efc8c6d52e0da152938e50f6a.
//
// Solidity: event CommitteeSnapshot(bytes12 indexed epochId, uint16 committeeSize, uint256[] pubKeys)
func (_DKGManager *DKGManagerFilterer) FilterCommitteeSnapshot(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerCommitteeSnapshotIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "CommitteeSnapshot", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCommitteeSnapshotIterator{contract: _DKGManager.contract, event: "CommitteeSnapshot", logs: logs, sub: sub}, nil
}

// WatchCommitteeSnapshot is a free log subscription operation binding the contract event 0x3687c3912698ec8584ed18d7c68916ff5fb9418efc8c6d52e0da152938e50f6a.
//
// Solidity: event CommitteeSnapshot(bytes12 indexed epochId, uint16 committeeSize, uint256[] pubKeys)
func (_DKGManager *DKGManagerFilterer) WatchCommitteeSnapshot(opts *bind.WatchOpts, sink chan<- *DKGManagerCommitteeSnapshot, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "CommitteeSnapshot", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerCommitteeSnapshot)
				if err := _DKGManager.contract.UnpackLog(event, "CommitteeSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitteeSnapshot is a log parse operation binding the contract event 0x3687c3912698ec8584ed18d7c68916ff5fb9418efc8c6d52e0da152938e50f6a.
//
// Solidity: event CommitteeSnapshot(bytes12 indexed epochId, uint16 committeeSize, uint256[] pubKeys)
func (_DKGManager *DKGManagerFilterer) ParseCommitteeSnapshot(log types.Log) (*DKGManagerCommitteeSnapshot, error) {
	event := new(DKGManagerCommitteeSnapshot)
	if err := _DKGManager.contract.UnpackLog(event, "CommitteeSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerContributionSubmittedIterator is returned from FilterContributionSubmitted and is used to iterate over the raw logs and unpacked data for ContributionSubmitted events raised by the DKGManager contract.
type DKGManagerContributionSubmittedIterator struct {
	Event *DKGManagerContributionSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerContributionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerContributionSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerContributionSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerContributionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerContributionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerContributionSubmitted represents a ContributionSubmitted event raised by the DKGManager contract.
type DKGManagerContributionSubmitted struct {
	EpochId             [12]byte
	Contributor         common.Address
	ContributorIndex    uint16
	CommitmentsHash     [32]byte
	EncryptedSharesHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterContributionSubmitted is a free log retrieval operation binding the contract event 0x8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5.
//
// Solidity: event ContributionSubmitted(bytes12 indexed epochId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
func (_DKGManager *DKGManagerFilterer) FilterContributionSubmitted(opts *bind.FilterOpts, epochId [][12]byte, contributor []common.Address) (*DKGManagerContributionSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "ContributionSubmitted", epochIdRule, contributorRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerContributionSubmittedIterator{contract: _DKGManager.contract, event: "ContributionSubmitted", logs: logs, sub: sub}, nil
}

// WatchContributionSubmitted is a free log subscription operation binding the contract event 0x8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5.
//
// Solidity: event ContributionSubmitted(bytes12 indexed epochId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
func (_DKGManager *DKGManagerFilterer) WatchContributionSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerContributionSubmitted, epochId [][12]byte, contributor []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "ContributionSubmitted", epochIdRule, contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerContributionSubmitted)
				if err := _DKGManager.contract.UnpackLog(event, "ContributionSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContributionSubmitted is a log parse operation binding the contract event 0x8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5.
//
// Solidity: event ContributionSubmitted(bytes12 indexed epochId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
func (_DKGManager *DKGManagerFilterer) ParseContributionSubmitted(log types.Log) (*DKGManagerContributionSubmitted, error) {
	event := new(DKGManagerContributionSubmitted)
	if err := _DKGManager.contract.UnpackLog(event, "ContributionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerDecryptionCombinedIterator is returned from FilterDecryptionCombined and is used to iterate over the raw logs and unpacked data for DecryptionCombined events raised by the DKGManager contract.
type DKGManagerDecryptionCombinedIterator struct {
	Event *DKGManagerDecryptionCombined // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerDecryptionCombinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerDecryptionCombined)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerDecryptionCombined)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerDecryptionCombinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerDecryptionCombinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerDecryptionCombined represents a DecryptionCombined event raised by the DKGManager contract.
type DKGManagerDecryptionCombined struct {
	EpochId         [12]byte
	Aid             [32]byte
	CiphertextIndex uint16
	CombineHash     [32]byte
	Plaintext       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecryptionCombined is a free log retrieval operation binding the contract event 0x4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) FilterDecryptionCombined(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (*DKGManagerDecryptionCombinedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "DecryptionCombined", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerDecryptionCombinedIterator{contract: _DKGManager.contract, event: "DecryptionCombined", logs: logs, sub: sub}, nil
}

// WatchDecryptionCombined is a free log subscription operation binding the contract event 0x4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) WatchDecryptionCombined(opts *bind.WatchOpts, sink chan<- *DKGManagerDecryptionCombined, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "DecryptionCombined", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerDecryptionCombined)
				if err := _DKGManager.contract.UnpackLog(event, "DecryptionCombined", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecryptionCombined is a log parse operation binding the contract event 0x4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) ParseDecryptionCombined(log types.Log) (*DKGManagerDecryptionCombined, error) {
	event := new(DKGManagerDecryptionCombined)
	if err := _DKGManager.contract.UnpackLog(event, "DecryptionCombined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerEpochAbortedIterator is returned from FilterEpochAborted and is used to iterate over the raw logs and unpacked data for EpochAborted events raised by the DKGManager contract.
type DKGManagerEpochAbortedIterator struct {
	Event *DKGManagerEpochAborted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerEpochAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerEpochAborted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerEpochAborted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerEpochAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerEpochAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerEpochAborted represents a EpochAborted event raised by the DKGManager contract.
type DKGManagerEpochAborted struct {
	EpochId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEpochAborted is a free log retrieval operation binding the contract event 0x379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be.
//
// Solidity: event EpochAborted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) FilterEpochAborted(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerEpochAbortedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "EpochAborted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerEpochAbortedIterator{contract: _DKGManager.contract, event: "EpochAborted", logs: logs, sub: sub}, nil
}

// WatchEpochAborted is a free log subscription operation binding the contract event 0x379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be.
//
// Solidity: event EpochAborted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) WatchEpochAborted(opts *bind.WatchOpts, sink chan<- *DKGManagerEpochAborted, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "EpochAborted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerEpochAborted)
				if err := _DKGManager.contract.UnpackLog(event, "EpochAborted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochAborted is a log parse operation binding the contract event 0x379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be.
//
// Solidity: event EpochAborted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) ParseEpochAborted(log types.Log) (*DKGManagerEpochAborted, error) {
	event := new(DKGManagerEpochAborted)
	if err := _DKGManager.contract.UnpackLog(event, "EpochAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerEpochCreatedIterator is returned from FilterEpochCreated and is used to iterate over the raw logs and unpacked data for EpochCreated events raised by the DKGManager contract.
type DKGManagerEpochCreatedIterator struct {
	Event *DKGManagerEpochCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerEpochCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerEpochCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerEpochCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerEpochCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerEpochCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerEpochCreated represents a EpochCreated event raised by the DKGManager contract.
type DKGManagerEpochCreated struct {
	EpochId          [12]byte
	Organizer        common.Address
	StartBlock       uint64
	SeedBlock        uint64
	LotteryThreshold *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochCreated is a free log retrieval operation binding the contract event 0x1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 startBlock, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) FilterEpochCreated(opts *bind.FilterOpts, epochId [][12]byte, organizer []common.Address) (*DKGManagerEpochCreatedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var organizerRule []interface{}
	for _, organizerItem := range organizer {
		organizerRule = append(organizerRule, organizerItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "EpochCreated", epochIdRule, organizerRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerEpochCreatedIterator{contract: _DKGManager.contract, event: "EpochCreated", logs: logs, sub: sub}, nil
}

// WatchEpochCreated is a free log subscription operation binding the contract event 0x1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 startBlock, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) WatchEpochCreated(opts *bind.WatchOpts, sink chan<- *DKGManagerEpochCreated, epochId [][12]byte, organizer []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var organizerRule []interface{}
	for _, organizerItem := range organizer {
		organizerRule = append(organizerRule, organizerItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "EpochCreated", epochIdRule, organizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerEpochCreated)
				if err := _DKGManager.contract.UnpackLog(event, "EpochCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochCreated is a log parse operation binding the contract event 0x1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 startBlock, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) ParseEpochCreated(log types.Log) (*DKGManagerEpochCreated, error) {
	event := new(DKGManagerEpochCreated)
	if err := _DKGManager.contract.UnpackLog(event, "EpochCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerEpochLiveIterator is returned from FilterEpochLive and is used to iterate over the raw logs and unpacked data for EpochLive events raised by the DKGManager contract.
type DKGManagerEpochLiveIterator struct {
	Event *DKGManagerEpochLive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerEpochLiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerEpochLive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerEpochLive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerEpochLiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerEpochLiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerEpochLive represents a EpochLive event raised by the DKGManager contract.
type DKGManagerEpochLive struct {
	EpochId           [12]byte
	ContributionCount uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEpochLive is a free log retrieval operation binding the contract event 0x0a9d060d45776170692faf35cb6f7bdda0152d8d49f36631cfa3547235467f63.
//
// Solidity: event EpochLive(bytes12 indexed epochId, uint16 contributionCount)
func (_DKGManager *DKGManagerFilterer) FilterEpochLive(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerEpochLiveIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "EpochLive", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerEpochLiveIterator{contract: _DKGManager.contract, event: "EpochLive", logs: logs, sub: sub}, nil
}

// WatchEpochLive is a free log subscription operation binding the contract event 0x0a9d060d45776170692faf35cb6f7bdda0152d8d49f36631cfa3547235467f63.
//
// Solidity: event EpochLive(bytes12 indexed epochId, uint16 contributionCount)
func (_DKGManager *DKGManagerFilterer) WatchEpochLive(opts *bind.WatchOpts, sink chan<- *DKGManagerEpochLive, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "EpochLive", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerEpochLive)
				if err := _DKGManager.contract.UnpackLog(event, "EpochLive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochLive is a log parse operation binding the contract event 0x0a9d060d45776170692faf35cb6f7bdda0152d8d49f36631cfa3547235467f63.
//
// Solidity: event EpochLive(bytes12 indexed epochId, uint16 contributionCount)
func (_DKGManager *DKGManagerFilterer) ParseEpochLive(log types.Log) (*DKGManagerEpochLive, error) {
	event := new(DKGManagerEpochLive)
	if err := _DKGManager.contract.UnpackLog(event, "EpochLive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerPartialDecryptionSubmittedIterator is returned from FilterPartialDecryptionSubmitted and is used to iterate over the raw logs and unpacked data for PartialDecryptionSubmitted events raised by the DKGManager contract.
type DKGManagerPartialDecryptionSubmittedIterator struct {
	Event *DKGManagerPartialDecryptionSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerPartialDecryptionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerPartialDecryptionSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerPartialDecryptionSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerPartialDecryptionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerPartialDecryptionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerPartialDecryptionSubmitted represents a PartialDecryptionSubmitted event raised by the DKGManager contract.
type DKGManagerPartialDecryptionSubmitted struct {
	EpochId          [12]byte
	Aid              [32]byte
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	DeltaX           *big.Int
	DeltaY           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartialDecryptionSubmitted is a free log retrieval operation binding the contract event 0x22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a302.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, bytes32 indexed aid, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, uint256 deltaX, uint256 deltaY)
func (_DKGManager *DKGManagerFilterer) FilterPartialDecryptionSubmitted(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, participant []common.Address) (*DKGManagerPartialDecryptionSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "PartialDecryptionSubmitted", epochIdRule, aidRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerPartialDecryptionSubmittedIterator{contract: _DKGManager.contract, event: "PartialDecryptionSubmitted", logs: logs, sub: sub}, nil
}

// WatchPartialDecryptionSubmitted is a free log subscription operation binding the contract event 0x22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a302.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, bytes32 indexed aid, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, uint256 deltaX, uint256 deltaY)
func (_DKGManager *DKGManagerFilterer) WatchPartialDecryptionSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerPartialDecryptionSubmitted, epochId [][12]byte, aid [][32]byte, participant []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "PartialDecryptionSubmitted", epochIdRule, aidRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerPartialDecryptionSubmitted)
				if err := _DKGManager.contract.UnpackLog(event, "PartialDecryptionSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePartialDecryptionSubmitted is a log parse operation binding the contract event 0x22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a302.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, bytes32 indexed aid, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, uint256 deltaX, uint256 deltaY)
func (_DKGManager *DKGManagerFilterer) ParsePartialDecryptionSubmitted(log types.Log) (*DKGManagerPartialDecryptionSubmitted, error) {
	event := new(DKGManagerPartialDecryptionSubmitted)
	if err := _DKGManager.contract.UnpackLog(event, "PartialDecryptionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerPoolKeyActivatedIterator is returned from FilterPoolKeyActivated and is used to iterate over the raw logs and unpacked data for PoolKeyActivated events raised by the DKGManager contract.
type DKGManagerPoolKeyActivatedIterator struct {
	Event *DKGManagerPoolKeyActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerPoolKeyActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerPoolKeyActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerPoolKeyActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerPoolKeyActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerPoolKeyActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerPoolKeyActivated represents a PoolKeyActivated event raised by the DKGManager contract.
type DKGManagerPoolKeyActivated struct {
	EpochId  [12]byte
	KeyIndex uint8
	X        *big.Int
	Y        *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolKeyActivated is a free log retrieval operation binding the contract event 0x66cc041d38da90dceb4c4eb68068985126297a2212bd69d65b5b75e2415f7296.
//
// Solidity: event PoolKeyActivated(bytes12 indexed epochId, uint8 indexed keyIndex, uint256 x, uint256 y)
func (_DKGManager *DKGManagerFilterer) FilterPoolKeyActivated(opts *bind.FilterOpts, epochId [][12]byte, keyIndex []uint8) (*DKGManagerPoolKeyActivatedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "PoolKeyActivated", epochIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerPoolKeyActivatedIterator{contract: _DKGManager.contract, event: "PoolKeyActivated", logs: logs, sub: sub}, nil
}

// WatchPoolKeyActivated is a free log subscription operation binding the contract event 0x66cc041d38da90dceb4c4eb68068985126297a2212bd69d65b5b75e2415f7296.
//
// Solidity: event PoolKeyActivated(bytes12 indexed epochId, uint8 indexed keyIndex, uint256 x, uint256 y)
func (_DKGManager *DKGManagerFilterer) WatchPoolKeyActivated(opts *bind.WatchOpts, sink chan<- *DKGManagerPoolKeyActivated, epochId [][12]byte, keyIndex []uint8) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "PoolKeyActivated", epochIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerPoolKeyActivated)
				if err := _DKGManager.contract.UnpackLog(event, "PoolKeyActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolKeyActivated is a log parse operation binding the contract event 0x66cc041d38da90dceb4c4eb68068985126297a2212bd69d65b5b75e2415f7296.
//
// Solidity: event PoolKeyActivated(bytes12 indexed epochId, uint8 indexed keyIndex, uint256 x, uint256 y)
func (_DKGManager *DKGManagerFilterer) ParsePoolKeyActivated(log types.Log) (*DKGManagerPoolKeyActivated, error) {
	event := new(DKGManagerPoolKeyActivated)
	if err := _DKGManager.contract.UnpackLog(event, "PoolKeyActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerPoolKeyClaimedIterator is returned from FilterPoolKeyClaimed and is used to iterate over the raw logs and unpacked data for PoolKeyClaimed events raised by the DKGManager contract.
type DKGManagerPoolKeyClaimedIterator struct {
	Event *DKGManagerPoolKeyClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerPoolKeyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerPoolKeyClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerPoolKeyClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerPoolKeyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerPoolKeyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerPoolKeyClaimed represents a PoolKeyClaimed event raised by the DKGManager contract.
type DKGManagerPoolKeyClaimed struct {
	EpochId  [12]byte
	Aid      [32]byte
	KeyIndex uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolKeyClaimed is a free log retrieval operation binding the contract event 0x587aa85ec6cb98aa5d1c21fbe47dbf442a2432b78629190c195062eb34a0303c.
//
// Solidity: event PoolKeyClaimed(bytes12 indexed epochId, bytes32 indexed aid, uint8 keyIndex)
func (_DKGManager *DKGManagerFilterer) FilterPoolKeyClaimed(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte) (*DKGManagerPoolKeyClaimedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "PoolKeyClaimed", epochIdRule, aidRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerPoolKeyClaimedIterator{contract: _DKGManager.contract, event: "PoolKeyClaimed", logs: logs, sub: sub}, nil
}

// WatchPoolKeyClaimed is a free log subscription operation binding the contract event 0x587aa85ec6cb98aa5d1c21fbe47dbf442a2432b78629190c195062eb34a0303c.
//
// Solidity: event PoolKeyClaimed(bytes12 indexed epochId, bytes32 indexed aid, uint8 keyIndex)
func (_DKGManager *DKGManagerFilterer) WatchPoolKeyClaimed(opts *bind.WatchOpts, sink chan<- *DKGManagerPoolKeyClaimed, epochId [][12]byte, aid [][32]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "PoolKeyClaimed", epochIdRule, aidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerPoolKeyClaimed)
				if err := _DKGManager.contract.UnpackLog(event, "PoolKeyClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolKeyClaimed is a log parse operation binding the contract event 0x587aa85ec6cb98aa5d1c21fbe47dbf442a2432b78629190c195062eb34a0303c.
//
// Solidity: event PoolKeyClaimed(bytes12 indexed epochId, bytes32 indexed aid, uint8 keyIndex)
func (_DKGManager *DKGManagerFilterer) ParsePoolKeyClaimed(log types.Log) (*DKGManagerPoolKeyClaimed, error) {
	event := new(DKGManagerPoolKeyClaimed)
	if err := _DKGManager.contract.UnpackLog(event, "PoolKeyClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerSeedResolvedIterator is returned from FilterSeedResolved and is used to iterate over the raw logs and unpacked data for SeedResolved events raised by the DKGManager contract.
type DKGManagerSeedResolvedIterator struct {
	Event *DKGManagerSeedResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerSeedResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerSeedResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerSeedResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerSeedResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerSeedResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerSeedResolved represents a SeedResolved event raised by the DKGManager contract.
type DKGManagerSeedResolved struct {
	EpochId [12]byte
	Seed    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSeedResolved is a free log retrieval operation binding the contract event 0xc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652.
//
// Solidity: event SeedResolved(bytes12 indexed epochId, bytes32 seed)
func (_DKGManager *DKGManagerFilterer) FilterSeedResolved(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerSeedResolvedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "SeedResolved", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerSeedResolvedIterator{contract: _DKGManager.contract, event: "SeedResolved", logs: logs, sub: sub}, nil
}

// WatchSeedResolved is a free log subscription operation binding the contract event 0xc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652.
//
// Solidity: event SeedResolved(bytes12 indexed epochId, bytes32 seed)
func (_DKGManager *DKGManagerFilterer) WatchSeedResolved(opts *bind.WatchOpts, sink chan<- *DKGManagerSeedResolved, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "SeedResolved", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerSeedResolved)
				if err := _DKGManager.contract.UnpackLog(event, "SeedResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSeedResolved is a log parse operation binding the contract event 0xc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652.
//
// Solidity: event SeedResolved(bytes12 indexed epochId, bytes32 seed)
func (_DKGManager *DKGManagerFilterer) ParseSeedResolved(log types.Log) (*DKGManagerSeedResolved, error) {
	event := new(DKGManagerSeedResolved)
	if err := _DKGManager.contract.UnpackLog(event, "SeedResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerSlotClaimedIterator is returned from FilterSlotClaimed and is used to iterate over the raw logs and unpacked data for SlotClaimed events raised by the DKGManager contract.
type DKGManagerSlotClaimedIterator struct {
	Event *DKGManagerSlotClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DKGManagerSlotClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerSlotClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DKGManagerSlotClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DKGManagerSlotClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerSlotClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerSlotClaimed represents a SlotClaimed event raised by the DKGManager contract.
type DKGManagerSlotClaimed struct {
	EpochId [12]byte
	Claimer common.Address
	Slot    uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlotClaimed is a free log retrieval operation binding the contract event 0x80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b.
//
// Solidity: event SlotClaimed(bytes12 indexed epochId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) FilterSlotClaimed(opts *bind.FilterOpts, epochId [][12]byte, claimer []common.Address) (*DKGManagerSlotClaimedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "SlotClaimed", epochIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerSlotClaimedIterator{contract: _DKGManager.contract, event: "SlotClaimed", logs: logs, sub: sub}, nil
}

// WatchSlotClaimed is a free log subscription operation binding the contract event 0x80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b.
//
// Solidity: event SlotClaimed(bytes12 indexed epochId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) WatchSlotClaimed(opts *bind.WatchOpts, sink chan<- *DKGManagerSlotClaimed, epochId [][12]byte, claimer []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "SlotClaimed", epochIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerSlotClaimed)
				if err := _DKGManager.contract.UnpackLog(event, "SlotClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlotClaimed is a log parse operation binding the contract event 0x80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b.
//
// Solidity: event SlotClaimed(bytes12 indexed epochId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) ParseSlotClaimed(log types.Log) (*DKGManagerSlotClaimed, error) {
	event := new(DKGManagerSlotClaimed)
	if err := _DKGManager.contract.UnpackLog(event, "SlotClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
