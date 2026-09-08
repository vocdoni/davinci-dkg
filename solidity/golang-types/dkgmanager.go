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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epochDurationBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_committeeSelectionBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_keyAssemblyBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finalizeGapBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minCommitteeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_maxLotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LOTTERY_ALPHA_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_COMMITTEE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ciphertextCount\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimPoolKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochDurationBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"transcriptDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppPoolIndex\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSelectionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyAssemblyDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liveNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolShareRoot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolStatus\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"nextIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"shareProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeFilled\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeSnapshot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"pubKeys\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochLive\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolKeyClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"keyIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DirectCallRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInSnapshot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolExhausted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TranscriptWordNotInField\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6102608060405234610438576101a081614d0b8038038091610021828561043c565b8339810103126104385780519063ffffffff8216908183036104385761004960208201610473565b9261005660408301610473565b61006260608401610473565b9061006f60808501610473565b61007b60a08601610473565b60c08601519160e087015194610100880151946101208901519a6101408a016100a390610487565b986100b16101608c01610487565b9a610180016100bf90610487565b9b4663ffffffff1603610429576001600160a01b0382161561041a576001600160a01b038316158015610409575b80156103f8575b80156103e7575b6103d85763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261013760388261043c565b5190201660c05260e052610100526101205261014052806103d257506064915b806103cc57506019905b806103c657506019955b806103c057506005915b610188836101838985610496565b610496565b9260018311908115916103b7575b81156103ae575b5080156103a4575b8015610394575b61035657610160526001600160401b038181166101e05261ffff9690916101d39190610496565b16610200526001600160401b03166102205280841661038e57506001905b8161018052838116155f1461038857506001915b826101a052838116155f1461038057508280925b836101c052169283911611918215610375575b508115610365575b5061035657336102405260405161485390816104b8823960805181611940015260a0518181816103a6015281816114cc015281816125690152614096015260c05181818161261d01528181612b810152613eed015260e051818181611356015281816135a40152614040015261010051818181610bbe0152818161189e015261395201526101205181818161184801528181612d3401526130510152610140518181816102010152818161202a015261246d0152610160518181816141df01526142e00152610180518181816110350152612ad501526101a0518181816102530152612aa401526101c0518181816102bb0152612a7301526101e0518161265e01526102005181612695015261022051816126c901526102405181612c850152f35b63d06b96b160e01b5f5260045ffd5b612710915061ffff16105f610234565b60201091505f61022c565b839092610219565b91610205565b906101f1565b506001600160401b0381116101ac565b50808310156101a5565b9050155f61019d565b88159150610196565b91610175565b9561016b565b90610161565b91610157565b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100fb565b506001600160a01b038516156100f4565b506001600160a01b038416156100ed565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761045f57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361043857565b519061ffff8216820361043857565b919082018092116104a357565b634e487b7160e01b5f52601160045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816304da5740146140c55750806306433b1b14614081578063074a75e11461401b57806318287e5f14613f1157806323488be514613ed1578063268ae2a114613ea55780632de546d514613e5f578063368e2a2714613e2d578063421adfbb14613d2557806356cbb5f314613cae5780635a8f2bb314613c655780635b0c0347146135d357806363f314cd1461358e5780636657755014612d75578063669a76a914612d0d5780636d16897d14612c5c5780636f067f6314612c1357806371712c291461023357806371a5978c146124ae57806372517b4b1461244657806377235ee114611c325780637ade132414611b8f5780637b31b5661461196457806385e1f4d0146119235780638dc1f53a1461187757806393c3d3a8146118325780639bbada6714611796578063a4adcd7f1461176f578063b7bca61514611059578063bd11c4c01461101a578063be59b8ea14610d13578063bea5210d14610bed578063bf19220914610ba8578063ca3c045814610ae0578063d3720aac146109e8578063d3979253146109aa578063d9933767146102df578063d9e9ca2e146102a0578063ebe86c1314610277578063f03a489814610238578063fa8f5e96146102335763fe1604b5146101ec575f80fd5b346102305780600319360112610230576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b80fd5b6141c8565b5034610230578060031936011261023057602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102305780600319360112610230576001546040516001600160a01b039091168152602090f35b5034610230578060031936011261023057602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610230576020366003190112610230576001600160a01b03196103026140ea565b1680825260026020526040822080546001600160a01b03161561099b57600281019081549060ff8216600182019182549160068110156109875760011480610970575b1561096157600581019061ffff808354169360101c1683101561095257868852600360209081526040808a20335f908152925290205460ff16610943576003810180549081156108d0575b506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316969060c0816024818b5afa9081156108c5578b916108a6575b5060608101516003811015610892575f19016108835760a001516001600160401b0390811660489290921c16111561087457600490604051602081019182523360601b60408201526034815261043e60548261428a565b519020910154111561086557858752600460205260408720805490600160401b821015610851579161047f8261050494600161ffff99989795018155614333565b81546001600160a01b03600392831b90811b19909116339182901b17909255898b5260209081526040808c205f9384529091529020805460ff19166001179055856104c9836144a8565b168619825416179055604051818152877f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a36144a8565b905490838260101c16938491161461051a578580f35b8260030261ffff811690810361083d576105376201fffe916146eb565b916002600160f11b0390600f1c16166201fffe61fffe82169116810361083d5790610564869493926146eb565b91848852600460205260408820885b85811061072757505050865b838110610684575060405160208101918260208251919201908a5b81811061066b57505050816105b79103601f19810183528261428a565b519020838752600a60205260408720556040519160408301908352604060208401528151809152602060608401920190875b8181106106525750505090807f3687c3912698ec8584ed18d7c68916ff5fb9418efc8c6d52e0da152938e50f6a920390a2805460ff191660021790557f23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb0678280a25f808080808580f35b82518452879550602093840193909201916001016105e9565b825184528a98506020938401939092019160010161059a565b80919293945060011b81810460021482151715610713576106a5818561471d565b516106b96106b38388614409565b8561471d565b52600181018082116106ff576106d26106da918661471d565b519186614409565b600181018091116106ff57906106f3600193928561471d565b5201908593929161057f565b634e487b7160e01b89526011600452602489fd5b634e487b7160e01b88526011600452602488fd5b909192939495506107388183614333565b90546040516313a4120960e31b815260039290921b1c6001600160a01b0316600482015260c081602481875afa908115610832578a91610804575b50600182018083116107dc57610789838761471d565b526020810151908260011b91838304600214841517156107f057906040916107b1848a61471d565b52015190600181018091116107dc57906107ce600193928861471d565b520190879594939291610573565b634e487b7160e01b8b52601160045260248bfd5b634e487b7160e01b8c52601160045260248cfd5b610825915060c03d811161082b575b61081d818361428a565b8101906144d1565b5f610773565b503d610813565b6040513d8c823e3d90fd5b634e487b7160e01b87526011600452602487fd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b633802147960e11b8952600489fd5b63aba4733960e01b8b5260048bfd5b634e487b7160e01b8c52602160045260248cfd5b6108bf915060c03d60c01161082b5761081d818361428a565b5f6103e7565b6040513d8d823e3d90fd5b9050608886901c6001600160401b031643811015610934574090811561092557819055877fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602083604051908152a25f610390565b6302504bb360e61b8a5260048afd5b63172181cb60e21b8a5260048afd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b50604082901c6001600160401b0316431115610345565b634e487b7160e01b88526021600452602488fd5b63d5b25b6360e01b8352600483fd5b50346102305760203660031901126102305760209060ff906040906001600160a01b03196109d66140ea565b168152600d8452205416604051908152f35b503461023057604036600319011261023057610a026140ea565b60243591906001600160a01b0383168303610adc57906040918160808451610a298161426f565b8281528260208201528286820152826060820152015260018060a01b03191681526005602052209060018060a01b03165f5260205260a060405f20604051610a708161426f565b815491600180851b0383169283835261ffff6020840191861c16815261ffff60018301549160408501928352608060ff60036002870154966060890197885201541695019415158552604051958652511660208501525160408401525160608301525115156080820152f35b5080fd5b5034610230576020366003190112610230576001600160a01b0319610b036140ea565b168152600460205260408120604051908160208254918281520190819285526020852090855b818110610b895750505082610b3f91038361428a565b604051928392602084019060208552518091526040840192915b818110610b67575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610b59565b82546001600160a01b0316845260209093019260019283019201610b29565b50346102305780600319360112610230576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461023057608036600319011261023057608090610c0a6140ea565b9060243591610c17614101565b610c1f614112565b91836060604051610c2f81614254565b828152826020820152826040820152015260018060a01b031916918284526007602052604084208585526020526040842061ffff82165f5260205261ffff6040600180825f20548487161c161495865f14610d0c5784975b8715610d055784965b825260066020528282209082526020522091165f5260205261ffff60405f2091165f5260205261ffff60405f20549160608260405196610ccf88614254565b16958681528360208201931683526040810194855201938452604051948552511660208401525160408301525115156060820152f35b8196610c90565b8097610c87565b503461023057602036600319011261023057610d2d6140ea565b81610160604051610d3d8161421d565b828152604051610d4c81614202565b8381528360208201528360408201528360608201528360808201528360a08201528360c082015260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015282610140820152015260018060a01b031916815260026020526040812060405190610dd48261421d565b80546001600160a01b03168252604051610ded81614202565b600182015461ffff8116825261ffff8160101c16602083015261ffff8160201c16604083015261ffff8160301c16606083015260018060401b038160401c16608083015260018060401b038160801c1660a083015260c01c60c08201526020830190815260028201549160ff83169060408501916006811015611006578252606085019360018060401b038160081c168552608086019060018060401b038160481c16825260a087019060018060401b039060881c16815260038301549160c08801928352600560048501549460e08a0195865201549661010089019961ffff89168b526101208a019661ffff8a60101c1688526101408b019861ffff8b60201c168a5261ffff6101608d019b60301c168b526040519b60018060a01b039051168c525161ffff81511660208d015261ffff60208201511660408d015261ffff60408201511660608d015261ffff60608201511660808d015260018060401b0360808201511660a08d015260018060401b0360a08201511660c08d015260c060018060401b039101511660e08c015251906006821015610ff257506101008a0152516001600160401b039081166101208a01529051811661014089015290511661016087015251610180860152516101a0850152845161ffff9081166101c0860152905181166101e08501529051811661020084015281511661022083015261024082f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b87526021600452602487fd5b5034610230578060031936011261023057602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102305760e0366003190112610230576110736140ea565b61107b614123565b906084356001600160401b03811161176b5761109b90369060040161419b565b909260a4356001600160401b0381116115ec576110bc90369060040161419b565b90939060c4356001600160401b038111611767576110de90369060040161419b565b96909132331480159061175d575b61174e576001600160a01b03198416808a52600260205260408a2080549198909790916001600160a01b03161561173f5760ff600289015416600189015490600681101561172b5760021480611714575b1561170557898c52600360209081526040808e20335f908152925290205460ff16156116f65761ffff88169a8b1580156116e6575b6116d75761119460408e8d815260046020522061118e8b61431f565b90614333565b90543360039290921b1c6001600160a01b031603611638578a8d52600560209081526040808f20335f908152925290206003015460ff166116c85780870197610100888a03126116c45788601f890112156116c457604051986111f96101008b61428a565b896101008a019182116116ac5789905b82821061169c5750505088519060a01c1480159061168a575b8015611675575b8015611667575b8015611657575b8015611647575b6116385761ffff82811694601084901c90911692836201fffe6112649260011b16614409565b6001600160fc1b038116810361162457836005029060058204850361160f57906112909160041b614409565b978860051b8981046020148a15171561160f578503611600575f5160206147fe5f395f51905f528f8e906112c536898d614416565b602081519101206040516112f5816112e76020820194606435604435876144bc565b03601f19810183528261428a565b519020905060405190602082019283527f4b37311b22cd0f09ae11d49f42ab65dce8fccf2600e6e2e7d41f51dc3d44b752602c830152604c820152604c815261133f606c8261428a565b51902006968760c08c0151036115f0578f939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156115ec5785936113a860405196879586948594635c73957b60e11b865260048601614398565b03915afa801561157e576115d7575b50506113c2836143de565b928360051b93808504602014901517156115c3576113df906143de565b816003029160038304036115c357906113f791614409565b908160051b91808304602014901517156107f0578183116115bf5781116115bb578161142892850191033691614416565b60208151910120878a52600a60205260408a2054036115ac579161144e9160e0936145c4565b9101510361159d57906003600592848752836020526040872060018060a01b0333165f5260205260405f209081549061ffff60a01b9060a01b169061ffff60a01b1916178155604435600182015501600160ff198254161790550161ffff815460101c1661ffff8114611589579060016114c992019061446f565b827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b15610adc57818091602460405180948193633c1bcdef60e21b83523360048401525af1801561157e57611565575b5050604051918252604435602083015260643560408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b8161156f9161428a565b61157a57825f611524565b8280fd5b6040513d84823e3d90fd5b634e487b7160e01b85526011600452602485fd5b63d1fed5fd60e01b8552600485fd5b63d1fed5fd60e01b8952600489fd5b8a80fd5b8b80fd5b634e487b7160e01b8d52601160045260248dfd5b816115e19161428a565b6115bf578b5f6113b7565b8580fd5b5063d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8f5260048ffd5b50634e487b7160e01b8f52601160045260248ffd5b634e487b7160e01b8f52601160045260248ffd5b63d1fed5fd60e01b8d5260048dfd5b5060643560a0890151141561123e565b5060443560808901511415611237565b508b60608901511415611230565b50604088015161ffff8360101c161415611229565b50602088015161ffff83161415611222565b8135815260209182019101611209565b8f80fd5b634e487b7160e01b5f52604160045260245ffd5b8d80fd5b6305d252c360e01b8d5260048dfd5b63652122d960e01b8d5260048dfd5b5061ffff8260101c168c11611172565b63965c290d60e01b8c5260048cfd5b63268dbf6760e21b8c5260048cfd5b50608081901c6001600160401b031643111561113d565b634e487b7160e01b8d52602160045260248dfd5b63d5b25b6360e01b8b5260048bfd5b6346f4d6c560e11b8952600489fd5b50333b15156110ec565b8780fd5b8380fd5b5034610230578060031936011261023057546040516001600160401b039091168152602090f35b50346102305761ffff60406117aa36614134565b9491828480516117b981614239565b828152826020820152015260018060a01b031916825260096020528282209082526020522091165f52602052606060405f206040516117f781614239565b81546040600161ffff83169485855260ff602086019460101c1615158452015492019182526040519283525115156020830152516040820152f35b50346102305780600319360112610230576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461023057806003193601126102305760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561191757906118e0575b602090604051908152f35b506020813d60201161190f575b816118fa6020938361428a565b8101031261190b57602090516118d5565b5f80fd5b3d91506118ed565b604051903d90823e3d90fd5b5034610230578060031936011261023057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102305760c03660031901126102305761197e6140ea565b6001600160a01b03191680825260026020526040822080549091602435916044359060a4359060843590606435906001600160a01b031615611b805760ff6002880154166006811015611b6c57600303611b5d576119dc818561466d565b6119e6838361466d565b84885260086020526040882086895260205261ffff611a0a8160408b2054166144a8565b16966101008811611b4e5760015460ff8160a01c1615611b3f576001600160a01b031689813b156102305760849160405192838092633a54cd5d60e11b82528b60048301528c60248301528d60448301523360648301525afa801561083257611b26575b509260a092869592611b0360057f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc978d8d9c9b60209f5260088f52604081208c82528f52604081208d61ffff198254161790556040611acf8888888c61458e565b918c81526020600f90528d828220908252602052208d5f528f5260405f20550161ffff600181835460301c16011690614489565b604051933385528b850152604084015260608301526080820152a4604051908152f35b611b318a809261428a565b611b3b575f611a6e565b8880fd5b63023b34fb60e11b8a5260048afd5b63464e67af60e01b8952600489fd5b63268dbf6760e21b8852600488fd5b634e487b7160e01b89526021600452602489fd5b63d5b25b6360e01b8852600488fd5b503461023057611b9e3661416a565b6001600160a01b03199091168083526002602081905260408420015491929160ff166006811015611c1e57600303611c0f57601060ff84161015611c00578160409160ff9352600c6020522091165f52602052602060405f2054604051908152f35b63d1fed5fd60e01b8252600482fd5b63268dbf6760e21b8252600482fd5b634e487b7160e01b83526021600452602483fd5b50346102305761010036600319011261023057611c4d6140ea565b611c55614101565b9060a4356001600160401b03811161176b57611c7590369060040161419b565b9060c4356001600160401b0381116115ec57611c9590369060040161419b565b90919060e4356001600160401b03811161176757611cb790369060040161419b565b6001600160a01b0319871689526002602052604089208054929590926001600160a01b0316156124375760ff6002840154166006811015612423576003036124145761ffff8916158015612405575b80156123fb575b6123ec576001600160a01b031988168a52600f60209081526040808c206024358d528252808c2061ffff8c165f90815292529020549586156123dd57600154976001600160a01b0389163b156115bf578b611d83818c8c6040518080958194633db5946d60e21b8352602435906004840161435c565b03916001600160a01b03165afa801561157e576123c8575b50506001600160a01b03198a168c52600760209081526040808e206024358f528252808e2061ffff8e165f90815292529020548c9590805b6123945750600101549461ffff8616116123855760018060a01b03198a168c52600960205260408c206024358d5260205260408c2061ffff8c165f5260205260405f209889549860ff8a60101c16612376578386019291908e610120868603126102305784601f87011215610230575060405193611e536101208661428a565b84906101208701116116ac5785905b610120870182106123665750508d8d85519060a01c1490811591612355575b8115612342575b508015612330575b8015612320575b8015612310575b6116005761ffff881660808501511061160057602060808501511161160057610cc08303611600578e836080116102305760405190611ede60a08361428a565b608082523660808d01116102305760808c602084013760a082015260208151910120036123015760ff8160a01c16156122f2578d60408d611f3793825180809681946381fe92fb60e01b8352602435906004840161435c565b03916001600160a01b03165afa9182156122e55781926122ab575b5060808a0135149081159161229c575b5061163857611f815f5160206147fe5f395f51905f529136908a614416565b60208151910120604051611fa3816112e76020820194608435606435876144bc565b5190206040516001600160a01b03198d16602082019081527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c80830193909352918152611ff8606c8261428a565b519020068060e08301510361163857606688612013926145c4565b6101008201510361226f5760800151948b939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156115ec57859361207c60405196879586948594635c73957b60e11b865260048601614398565b03915afa801561157e57612287575b505060c083018311610713576104c0830183116107135760101c61ffff1687805b838110612181575050505b60208110612122575050620100009062ff00001916178155600160843591015561ffff604051926064358452608435602085015216917f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae160406024359360018060a01b03191692a480f35b6104c08160061b83010160c08260051b840101356121725780351590811591612162575b50612153576001016120b7565b63d1fed5fd60e01b8752600487fd5b600191506020013514155f612146565b63d1fed5fd60e01b8852600488fd5b60c08160051b860101356104c08260061b870101908015801561227e575b61226f5761ffff16906001821b938481166116385784179360018060a01b03198b168d52600760205260408d208d602435905260205260408d2061ffff8d165f5260205260405f2054161561226f5760405160208082019280358452013560408201526040815261221160608261428a565b5190209060018060a01b03198a168c52600660205260408c206024358d5260205260408c2061ffff8c165f5260205261ffff60405f2091165f5260205260405f205403612260576001016120ac565b63d1fed5fd60e01b8a5260048afd5b63d1fed5fd60e01b8c5260048cfd5b5084811161219f565b816122919161428a565b61176757875f61208b565b905060a089013514155f611f62565b9150506040813d6040116122dd575b816122c76040938361428a565b810103126116c45760208151910151905f611f52565b3d91506122ba565b50604051903d90823e3d90fd5b63023b34fb60e11b8e5260048efd5b63d1fed5fd60e01b8e5260048efd5b5060843560c08501511415611e9e565b5060643560a08501511415611e97565b5061ffff881660608501511415611e90565b61ffff915016604085015114158e611e88565b602086015160243514159150611e81565b8135815260209182019101611e62565b63955c0c4960e01b8e5260048efd5b63032cddf960e11b8c5260048cfd5b5f198101908082116116245716955f1981146123b4576001019580611dd3565b634e487b7160e01b8e52601160045260248efd5b816123d29161428a565b6115bf578b5f611d9b565b6346f551f560e01b8b5260048bfd5b636d28699160e01b8a5260048afd5b5060643515611d0d565b5061010061ffff8a1611611d06565b63268dbf6760e21b8a5260048afd5b634e487b7160e01b8b52602160045260248bfd5b63d5b25b6360e01b8a5260048afd5b503461023057806003193601126102305760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561191757906118e057602090604051908152f35b50346102305760803660031901126102305760043561ffff8116809103610adc576124d7614123565b6124df614101565b916124e8614112565b906001600160401b036124f96142c6565b164310612b66575b80158015612b5a575b8015612b4d575b8015612b3f575b8015612b35575b8015612b29575b8015612b18575b8015612b0b575b8015612afc575b8015612acf575b8015612a9e575b8015612a6d575b612a5e57604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115612a53578691612a19575b506001600160401b0316918215612a0a5761ffff8091169316916125c583856143f6565b9080612710029061271082040361083d578082106129dd5750505f19915b8554906001600160401b03808316908114610713576001600160401b036001909101166001600160401b03199290921682178755612641827f0000000000000000000000000000000000000000000000000000000000000000614645565b95436001600160401b03908116600101969087116106ff5761268c7f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166142ad565b9061ffff6126c37f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166142ad565b936126f77f0000000000000000000000000000000000000000000000000000000000000000436001600160401b03166142ad565b95604051986127058a614202565b895260208901521660408088019190915260608701919091526001600160401b03918216608087015291811660a08601529190911660c0840152519161274a8361421d565b338352602083015260016040830152606082015260018060401b034316608082015260018060401b03831660a08201528460c08201528160e08201528461010082015284610120820152846101408201528461016082015260018060a01b03198416855260026020526040852060018060a01b0382511660018060a01b031982541617815560018101602083015161ffff808251161661ffff198354161782556127fc61ffff6020830151168361446f565b61280e61ffff604083015116836143bf565b61282061ffff60608301511683614489565b608081810151835460a084015160c0948501516001600160401b03909216604093841b600160401b600160801b031617931b600160801b600160c01b0316929092179190921b6001600160c01b031916179091558201516006811015611006576002820180546060850151608086015160a08701516001600160c81b031990931660ff9095169490941760089190911b610100600160481b03161760489390931b600160481b600160881b03169290921760889290921b600160881b600160c81b031691909117905560c0820151600382015560e08201516004820155610100820151600591909101805461ffff191661ffff92831617815561012083015160209793612953939161016091906129399084168561446f565b61294a8361014083015116856143bf565b01511690614489565b8054600160401b600160801b03191643604081811b600160401b600160801b03169290921790925580516001600160401b03928316815291909316858201529182015233906001600160a01b03198316907f1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d90606090a36040516001600160a01b03199091168152f35b80156129f6576129f091905f19046143f6565b916125e3565b634e487b7160e01b87526012600452602487fd5b63d06b96b160e01b8652600486fd5b90506020813d602011612a4b575b81612a346020938361428a565b810103126115ec57612a459061445b565b5f6125a1565b3d9150612a27565b6040513d88823e3d90fd5b63d06b96b160e01b8552600485fd5b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff831611612550565b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff841610612549565b5061ffff7f0000000000000000000000000000000000000000000000000000000000000000168110612542565b5061271061ffff83161061253b565b508061ffff851610612534565b5061ffff831661ffff85161161252d565b5061ffff841615612526565b506020811161251f565b50602061ffff841611612518565b5061ffff83168111612511565b5061ffff83161561250a565b84546001600160a01b031990612ba5906001600160401b03167f0000000000000000000000000000000000000000000000000000000000000000614645565b16808652600260205260ff60026040882001541690600682101561100657600382149081612bf8575b50159081612bec575b50156125015763268dbf6760e21b8552600485fd5b6004915014155f612bd7565b90508652600d602052600f60ff60408820541610155f612bce565b50346102305760403660031901126102305760209061ffff906040906001600160a01b0319612c406140ea565b1681526008845281812060243582528452205416604051908152f35b5034610230576020366003190112610230576004356001600160a01b03811690819003610adc577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303612cff576001549060ff8260a01c16612cf0578015612ce1576001600160a81b031990911617600160a01b1760015580f35b63e6c4247b60e01b8352600483fd5b6373253a9760e01b8352600483fd5b6282b42960e81b8252600482fd5b503461023057806003193601126102305760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561191757906118e057602090604051908152f35b50346102305760a036600319011261023057612d8f6140ea565b6024356044356001600160401b03811161176b57612db190369060040161419b565b6064356001600160401b0381116115ec57612dd090369060040161419b565b9092906084356001600160401b03811161176757612df290369060040161419b565b9094323314801590613584575b61174e576001600160a01b03198816808a52600260205260408a2080549098919791906001600160a01b03161561173f57600289019660ff885416600681101561172b576003811461357557600119016117055760018a0154998a60c01c431061356657600561ffff91015460101c169961ffff8160201c168b106135575761ffff8160101c169b60018c10801561354e575b6123015761046093929190618bff1983016116005760e08814801590613542575b6116005785880194939291908f60e0888803126102305786601f89011215610230575060405195612ee560e08861428a565b8660e0890191821161353c5788905b82821061352c575050505f5160206147fe5f395f51905f52865110801590613511575b80156134f6575b80156134db575b80156134c0575b80156134a5575b801561348a575b6115f05785519060a01c1490811591613476575b508015613468575b801561345a575b801561344c575b6123015790612f835f5160206147fe5f395f51905f529236908b614416565b602081519101206040519060208201928352604082015260408152612fa960608261428a565b51902060405160208101918c83527fe28959afa6ea38549c61aff75344fc2c9f148f1259fcef44fdd297a1d9a39d0f602c830152604c820152604c8152612ff1606c8261428a565b519020068060a0840151036116385760c0918861300d926145c4565b91015103612260576104008501851161325c5789805b602082106133575750506108408501938486116107dc57908a93929184956042965b601081106132d95750507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156115ec5785936130a360405196879586948594635c73957b60e11b865260048601614398565b03915afa801561157e576132c0575b505b60ff8116601081101561328457828102818104841482151715613270576040018060401161325c578060051b908082046020149015171561325c576130f99085614409565b868a52600b60205260408a20825f5260205260405f208135815560016020830135910155610400906040519161312f818461428a565b3683376040810181116107dc578a5b8a8110613230575050885b602081106131fb57508960205b6005821061318f5750509060ff92915190878b52600c60205260408b20905f5260205260405f20551660ff8114610713576001016130b4565b60011c908b5b8281106131a6575060010190613156565b8060011b818104600214821517156123b457806131c48f9287614634565b5191506001810180911161162457600192916131e36131ea9288614634565b5190614731565b6131f48287614634565b5201613195565b807f11a12e535a08d28aa7434e11614f2eb9b34da3fcba5746a376ed981855fb01f061322960019385614634565b5201613149565b8061324b60019260061b8401604060608201359101356145ac565b6132558286614634565b520161313e565b634e487b7160e01b8a52601160045260248afd5b634e487b7160e01b5f52601160045260245ffd5b88867f0a9d060d45776170692faf35cb6f7bdda0152d8d49f36631cfa3547235467f6360208a89600360ff19825416179055604051908152a280f35b816132ca9161428a565b6132d557865f6130b2565b8680fd5b90919293949550868102818104881482151715613270578060051b90808204602014901517156123b45761330d9083614409565b8c5b60208110613328575050600101908c9594939291613045565b8060061b820180351590811591613347575b506116005760010161330f565b600191506020013514155f61333a565b8160051b87016104008135910135918b841015613428578115801561341f575b612301576001821b90818116611600578b8f52600460205260408f209117939291905f1982019082821161160f578f916133b390604092614333565b60018060a01b0391549060031b1c16918d81526005602052209060018060a01b03165f5260205260405f209060ff60038301541615908115613409575b5061230157600101540361226f576001905b0190613023565b905061ffff80835460a01c16911614155f6133f0565b508c8211613377565b9291901590811591613442575b5061226f57600190613402565b905015155f613435565b508160808501511415612f64565b508b60608501511415612f5d565b508c60408501511415612f56565b905061ffff6020860151911614155f612f4e565b505f5160206147fe5f395f51905f5260c08701511015612f3a565b505f5160206147fe5f395f51905f5260a08701511015612f33565b505f5160206147fe5f395f51905f5260808701511015612f2c565b505f5160206147fe5f395f51905f5260608701511015612f25565b505f5160206147fe5f395f51905f5260408701511015612f1e565b505f5160206147fe5f395f51905f5260208701511015612f17565b8135815260209182019101612ef4565b50508f80fd5b50610100891415612eb3565b508c8c11612e92565b63368f2d7d60e21b8d5260048dfd5b63268dbf6760e21b8d5260048dfd5b6337bca76b60e21b8d5260048dfd5b50333b1515612dff565b50346102305780600319360112610230576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461190b5761018036600319011261190b576135ee6140ea565b6135f6614101565b906135ff614112565b610124356001600160401b03811161190b5761361f90369060040161419b565b90610144356001600160401b03811161190b5761364090369060040161419b565b90916001600160401b03610164351161190b573660236101643501121561190b576001600160401b0360046101643501351161190b57366024610164356004013560051b6101643501011161190b576001600160a01b031986165f90815260026020526040902080549093906001600160a01b031615613c565760ff6002850154166006811015613c4257600303613c33576001600160a01b031987165f90815260036020908152604080832033845290915290205460ff1615613c245761ffff8816158015613c0c575b8015613c00575b8015613bf1575b8015613be6575b613bd8576001600160a01b031987165f9081526004602052604090206137499061118e8a61431f565b90543360039290921b1c6001600160a01b031603613bba5760018060a01b031987165f52600f60205260405f206024355f5260205260405f2061ffff87165f5260205260405f20548015613bc9576137ab60e43560c43560a43560843561458e565b03613bba576001546001600160a01b0316803b1561190b57604051633db5946d60e21b8152905f90829081806137e76024358e6004840161435c565b03915afa8015613baf57613b9a575b5060018060a01b0319871689526007602052604089206024358a526020526040892061ffff87165f5260205260405f20600161ffff8a161b905416613b8b57828101946101e082870312613b875785601f83011215613b87576040519561385f6101e08861428a565b86906101e08401116115bb5782905b6101e084018210613b7757505085518860a01c14801590613b67575b8015613b55575b8015613b43575b8015613b33575b8015613b23575b6122605761010086015161012087015160405190602082019283526040820152604081526138d560608261428a565b5190206101043503612260576001600160a01b031988168a52600c60205260408a2060ff6139056024358b61455a565b165f5260205260405f205461392360c088015160e0890151906145ac565b600561016435600401350361226f5761ffff8b165f19018c5b60058110613ac3575050036122605789939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156115ec5785936139a460405196879586948594635c73957b60e11b865260048601614398565b03915afa801561157e57613aae575b50506005613a4f9160018060a01b03198616885260066020526040882060243589526020526040882061ffff86165f5260205260405f2061ffff88165f5260205260405f2061010435905560018060a01b03198616885260076020526040882060243589526020526040882061ffff86165f5260205260405f20600161ffff89161b81541790550161ffff600181835460201c160116906143bf565b61ffff6101206101008301519201519281604051961686521660208501526040840152606083015233917f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30260806024359360018060a01b03191692a480f35b81613ab89161428a565b6115ec57855f6139b3565b90916101643560040135821015613b0f5760019061016435600584901b016024013590848316613b015790613af791614731565b92811c910161393c565b613b0a91614731565b613af7565b634e487b7160e01b8e52603260045260248efd5b5060a43560a087015114156138a6565b506084356080870151141561389f565b5061ffff891660608701511415613898565b50604086015161ffff88161415613891565b506024356020870151141561388a565b813581526020918201910161386e565b8980fd5b633466526160e01b8952600489fd5b613ba79199505f9061428a565b5f975f6137f6565b6040513d5f823e3d90fd5b63d1fed5fd60e01b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b506101043515613720565b5061010061ffff871611613719565b5061ffff861615613712565b5061ffff600185015460101c1661ffff89161161370b565b63965c290d60e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b3461190b57613c7336614134565b9160018060a01b0319165f52600960205260405f20905f5260205261ffff60405f2091165f526020526020600160405f200154604051908152f35b3461190b57613cbc3661416a565b6001600160a01b03199091165f818152600260208190526040909120015460ff166006811015613c4257600303613c3357601060ff83161015613bba575f52600b60205260ff60405f2091165f526020526040805f206001815491015482519182526020820152f35b3461190b57604036600319011261190b57613d3e6140ea565b60015460243591906001600160a01b03163303613e1f5760018060a01b03191690815f52600260205260ff600260405f200154166006811015613c4257600303613c3357815f52600d60205260ff60405f205416906010821015613e105760209260ff6001840116815f52600d855260405f2060ff821660ff19825416179055815f52600e855260405f20835f52855260ff60405f20911660ff198254161790557f587aa85ec6cb98aa5d1c21fbe47dbf442a2432b78629190c195062eb34a0303c84604051858152a3604051908152f35b6311a7ebfd60e31b5f5260045ffd5b6282b42960e81b5f5260045ffd5b3461190b57604036600319011261190b576020613e54613e4b6140ea565b6024359061455a565b60ff60405191168152f35b3461190b57613e6d36614134565b9160018060a01b0319165f52600f60205260405f20905f5260205261ffff60405f2091165f52602052602060405f2054604051908152f35b3461190b575f36600319011261190b576020613ebf6142c6565b6040516001600160401b039091168152f35b3461190b575f36600319011261190b57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461190b57602036600319011261190b576001600160a01b0319613f336140ea565b165f81815260026020526040902080546001600160a01b031615613c5657600281019060ff825416906006821015613c4257600182149182614001575b8215613faf575b505015613c3357805460ff191660041790557f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b600214915081613fe7575b81613fc8575b508380613f77565b905061ffff600181600584015460101c1692015460201c161183613fc0565b600181015460801c6001600160401b031643119150613fba565b600182015460401c6001600160401b031643119250613f70565b3461190b575f36600319011261190b5760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015613baf575f906118e057602090604051908152f35b3461190b575f36600319011261190b576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461190b575f36600319011261190b575f5460401c6001600160401b03168152602090f35b600435906001600160a01b03198216820361190b57565b6044359061ffff8216820361190b57565b6064359061ffff8216820361190b57565b6024359061ffff8216820361190b57565b606090600319011261190b576004356001600160a01b03198116810361190b57906024359060443561ffff8116810361190b5790565b604090600319011261190b576004356001600160a01b03198116810361190b579060243560ff8116810361190b5790565b9181601f8401121561190b578235916001600160401b03831161190b576020838186019501011161190b57565b3461190b575f36600319011261190b5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60e081019081106001600160401b038211176116b057604052565b61018081019081106001600160401b038211176116b057604052565b606081019081106001600160401b038211176116b057604052565b608081019081106001600160401b038211176116b057604052565b60a081019081106001600160401b038211176116b057604052565b601f909101601f19168101906001600160401b038211908210176116b057604052565b6001600160401b03918216908216019190821161327057565b5f5460401c6001600160401b031680156143115761430e907f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316906142ad565b90565b50436001600160401b031690565b61ffff5f199116019061ffff821161327057565b8054821015614348575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b03199091168152602081019190915260400190565b908060209392818452848401375f828201840152601f01601f1916010190565b92906143b19061430e9593604086526040860191614378565b926020818503910152614378565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b600581901b91906001600160fb1b0381160361327057565b8181029291811591840414171561327057565b9190820180921161327057565b9192916001600160401b0382116116b0576040519161443f601f8201601f19166020018461428a565b82948184528183011161190b578281602093845f960137010152565b51906001600160401b038216820361190b57565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b61ffff60019116019061ffff821161327057565b91606093918352602083015260408201520190565b908160c091031261190b576040519060c082016001600160401b038111838210176116b05760405280516001600160a01b038116810361190b57825260208101516020830152604081015160408301526060810151600381101561190b576145529160a09160608501526145476080820161445b565b60808501520161445b565b60a082015290565b6001600160a01b0319165f908152600e60209081526040808320938352929052205460ff168015613bba575f190160ff1690565b91608093916040519384526020840152604083015260608201522090565b604191604051915f8353600183015260218201522090565b92915f5160206147fe5f395f51905f525f940691600192809260051b8201915b8281106146055750505050156145f657565b63331835e560e21b5f5260045ffd5b909192955f5160206147fe5f395f51905f52838160209381863599818b1016998c0990089809939291016145e4565b9060208110156143485760051b0190565b60401b63ffffffff60401b166001600160401b039091161760a01b6001600160a01b03191690565b905f5160206147fe5f395f51905f528210806146be575b156146a5578115806146b4575b6146a55761469e9161474a565b156146a557565b634c4d29cd60e11b5f5260045ffd5b5060018114614691565b505f5160206147fe5f395f51905f528110614684565b6001600160401b0381116116b05760051b60200190565b906146f5826146d4565b614702604051918261428a565b8281528092614713601f19916146d4565b0190602036910137565b80518210156143485760209160051b010190565b6041916040519160018353600183015260218201522090565b5f5160206147fe5f395f51905f5281108015906147e6575b6147e0575f5160206147fe5f395f51905f528181920991800990805f5160206147fe5f395f51905f5203915f5160206147fe5f395f51905f528311613270575f5160206147fe5f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b50505f90565b505f5160206147fe5f395f51905f5282101561476256fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122061d78a074c204f382d029be6bb0cc8a379cd21afd7625798449cc79a527cc89e64736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _finalizeVerifier common.Address, _decryptCombineVerifier common.Address, _epochDurationBlocks *big.Int, _committeeSelectionBlocks *big.Int, _keyAssemblyBlocks *big.Int, _finalizeGapBlocks *big.Int, _minThreshold uint16, _minCommitteeSize uint16, _maxLotteryAlphaBps uint16) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _finalizeVerifier, _decryptCombineVerifier, _epochDurationBlocks, _committeeSelectionBlocks, _keyAssemblyBlocks, _finalizeGapBlocks, _minThreshold, _minCommitteeSize, _maxLotteryAlphaBps)
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

// FINALIZEVERIFIER is a free data retrieval call binding the contract method 0x93c3d3a8.
//
// Solidity: function FINALIZE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) FINALIZEVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "FINALIZE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FINALIZEVERIFIER is a free data retrieval call binding the contract method 0x93c3d3a8.
//
// Solidity: function FINALIZE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) FINALIZEVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.FINALIZEVERIFIER(&_DKGManager.CallOpts)
}

// FINALIZEVERIFIER is a free data retrieval call binding the contract method 0x93c3d3a8.
//
// Solidity: function FINALIZE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) FINALIZEVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.FINALIZEVERIFIER(&_DKGManager.CallOpts)
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

// GetFinalizeVerifierVKeyHash is a free data retrieval call binding the contract method 0x669a76a9.
//
// Solidity: function getFinalizeVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetFinalizeVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getFinalizeVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizeVerifierVKeyHash is a free data retrieval call binding the contract method 0x669a76a9.
//
// Solidity: function getFinalizeVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetFinalizeVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetFinalizeVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetFinalizeVerifierVKeyHash is a free data retrieval call binding the contract method 0x669a76a9.
//
// Solidity: function getFinalizeVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetFinalizeVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetFinalizeVerifierVKeyHash(&_DKGManager.CallOpts)
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
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex)
func (_DKGManager *DKGManagerCaller) GetPoolStatus(opts *bind.CallOpts, epochId [12]byte) (uint8, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPoolStatus", epochId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPoolStatus is a free data retrieval call binding the contract method 0xd3979253.
//
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex)
func (_DKGManager *DKGManagerSession) GetPoolStatus(epochId [12]byte) (uint8, error) {
	return _DKGManager.Contract.GetPoolStatus(&_DKGManager.CallOpts, epochId)
}

// GetPoolStatus is a free data retrieval call binding the contract method 0xd3979253.
//
// Solidity: function getPoolStatus(bytes12 epochId) view returns(uint8 nextIndex)
func (_DKGManager *DKGManagerCallerSession) GetPoolStatus(epochId [12]byte) (uint8, error) {
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

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x66577550.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) FinalizeEpoch(opts *bind.TransactOpts, epochId [12]byte, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "finalizeEpoch", epochId, transcriptDigest, transcript, proof, input)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x66577550.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) FinalizeEpoch(epochId [12]byte, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId, transcriptDigest, transcript, proof, input)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x66577550.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 transcriptDigest, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) FinalizeEpoch(epochId [12]byte, transcriptDigest [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId, transcriptDigest, transcript, proof, input)
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
