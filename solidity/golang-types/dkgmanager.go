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
	Contributor            common.Address
	ContributorIndex       uint16
	CommitmentsHash        [32]byte
	EncryptedSharesHash    [32]byte
	CommitmentVectorDigest [32]byte
	Accepted               bool
}

// DKGTypesDecryptionPolicy is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesDecryptionPolicy struct {
	OwnerOnly          bool
	MaxDecryptions     uint16
	NotBeforeBlock     uint64
	NotBeforeTimestamp uint64
	NotAfterBlock      uint64
	NotAfterTimestamp  uint64
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

// DKGTypesPoint is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesPoint struct {
	X *big.Int
	Y *big.Int
}

// IDKGManagerEpoch is an auto generated low-level Go binding around an user-defined struct.
type IDKGManagerEpoch struct {
	Organizer              common.Address
	Policy                 DKGTypesEpochPolicy
	DecryptionPolicy       DKGTypesDecryptionPolicy
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epochDurationBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_committeeSelectionBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_keyAssemblyBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finalizeGapBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochDurationBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSelectionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyAssemblyDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liveNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeFilled\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochLive\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x610200806040523461034d57610140816146e480380380916100218285610351565b83398101031261034d5780519063ffffffff82169182810361034d5761004960208301610388565b61005560408401610388565b61006160608501610388565b61006d60808601610388565b9061007a60a08701610388565b9260c08701519460e0880151966101206101008a01519901519963ffffffff46160361033e576001600160a01b0382161561032f576001600160a01b03831615801561031e575b801561030d575b80156102fc575b6102ed5763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261010b603882610351565b5190201660c05260e052610100526101205261014052806102e757506064905b806102e157506019915b806102db57506019925b806102d557506005925b61015c84610157838661039c565b61039c565b9360018411908115916102cc575b81156102c3575b5080156102b9575b80156102a9575b61029a57610160919091526001600160401b0382811661018052916101a5919061039c565b166101a0526001600160401b03166101c052336101e05260405161432690816103be823960805181611bd6015260a0518181816102e7015281816112190152818161339d0152613b1b015260c0518181816131ad015261345e015260e0518181816110d2015281816125bd0152613ac50152610100518181816109be015281816117210152611b38015261012051818181611ae20152818161256601526128ac0152610140518181816101bf01528181611f82015261244d015261016051818181613c830152613eda0152610180518161349e01526101a051816134cc01526101c051816134f701526101e051816124b70152f35b63d06b96b160e01b5f5260045ffd5b506001600160401b038211610180565b5081841015610179565b9050155f610171565b8215915061016a565b92610149565b9261013f565b91610135565b9061012b565b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100cf565b506001600160a01b038516156100c8565b506001600160a01b038416156100c1565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761037457604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361034d57565b919082018092116103a957565b634e487b7160e01b5f52601160045260245ffdfe60e0806040526004361015610012575f80fd5b5f905f3560e01c90816304da574014613b4a5750806306433b1b14613b06578063074a75e114613aa0578063143b4616146132c257806318287e5f146131d157806323488be5146131915780632648f56714612dfe578063268ae2a114612dd15780632de546d514612d8b5780633353ec6e14612d555780634554c0be14612d0157806349c61a1214612685578063510ba2df146126355780635a8f2bb3146125ec57806363f314cd146125a7578063669a76a91461253f5780636d16897d1461248e57806371712c29146101f157806372517b4b1461242657806377235ee114611bfa57806385e1f4d014611bb95780638dc1f53a14611b1157806393c3d3a814611acc5780639bbada6714611a30578063a305e0f314611464578063a4adcd7f1461143d578063b7bca61514610e4d578063be59b8ea14610b13578063bea5210d146109ed578063bf192209146109a8578063ca3c0458146108e0578063d3720aac146107ef578063d99337671461021f578063ebe86c13146101f6578063fa8f5e96146101f15763fe1604b5146101aa575f80fd5b346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b80fd5b613c6c565b50346101ee57806003193601126101ee576001546040516001600160a01b039091168152602090f35b50346101ee5760203660031901126101ee576001600160a01b0319610242613ba2565b1680825260026020526040822080546001600160a01b0316156107e057600481019081549060ff8216600182019182549160068110156107cc57600114806107b5575b156107a657600781019061ffff808354169360101c1683101561079757868852600360209081526040808a20335f908152925290205460ff1661078857600581018054958615610715575b50506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316959060a0816024818a5afa801561067d576060918b916106f6575b50015160038110156106e2575f19016106d357600690604051602081019182523360601b604082015260348152610362605482613d49565b51902091015411156106c457858752600460205260408720805490600160401b8210156106b0579261ffff959493926103a5838895600161042a9601815561400c565b81546001600160a01b03600392831b90811b19909116339182901b179092558a8c5260209081526040808d205f9384529091529020805460ff19166001179055836103ef8361411c565b168419825416179055604051818152887f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a361411c565b915460101c169216821461043c578480f35b6040519161040061044d8185613d49565b368437604051916108006104618185613d49565b368437858752600460205260408720875b83811061057b575050505b6020811061053557506040519060208201928387905b6020821061051f57505050610420820186905b6040821061050957505050610c0081526104c2610c2082613d49565b519020828452600b6020526040842055805460ff191660021790557f23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb0678280a25f8080808480f35b60208060019285518152019301910190916104a6565b6020806001928551815201930191019091610493565b8060011b818104600214821517156105675760018101809111610567579060016105608193856141fa565b520161047d565b634e487b7160e01b87526011600452602487fd5b6001810180821161069c576020821015610688578160051b8701526105a0818361400c565b90546040516313a4120960e31b815260039290921b1c6001600160a01b0316600482015260a081602481875afa90811561067d578a9161064f575b506020810151908260011b918383046002148415171561063b5790604091610603848a6141fa565b5201519060018101809111610627579061062060019392886141fa565b5201610472565b634e487b7160e01b8b52601160045260248bfd5b634e487b7160e01b8c52601160045260248cfd5b610670915060a03d8111610676575b6106688183613d49565b8101906140a5565b5f6105db565b503d61065e565b6040513d8c823e3d90fd5b634e487b7160e01b8a52603260045260248afd5b634e487b7160e01b8a52601160045260248afd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b63aba4733960e01b8952600489fd5b634e487b7160e01b8a52602160045260248afd5b61070f915060a03d60a011610676576106688183613d49565b5f61032a565b90955060881c6001600160401b031643811015610779574094851561076a57859055867fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602087604051908152a25f806102d0565b6302504bb360e61b8952600489fd5b63172181cb60e21b8952600489fd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b50604082901c6001600160401b0316431115610285565b634e487b7160e01b88526021600452602488fd5b63d5b25b6360e01b8352600483fd5b50346101ee5760403660031901126101ee57610809613ba2565b60243591906001600160a01b03831683036108dc579060409161082a613f81565b5060018060a01b03191681526005602052209060018060a01b03165f5260205260c060405f2060405161085c81613cdd565b81549160018060a01b0383169283835261ffff602084019160a01c16815260018201546040840190815261ffff6002840154926060860193845260a060ff600460038801549760808a01988952015416960195151586526040519687525116602086015251604085015251606084015251608083015251151560a0820152f35b5080fd5b50346101ee5760203660031901126101ee576001600160a01b0319610903613ba2565b168152600460205260408120604051908160208254918281520190819285526020852090855b818110610989575050508261093f910383613d49565b604051928392602084019060208552518091526040840192915b818110610967575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610959565b82546001600160a01b0316845260209093019260019283019201610929565b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee5760803660031901126101ee57608090610a0a613ba2565b9060243591610a17613b80565b610a1f613b91565b91836060604051610a2f81613d2e565b828152826020820152826040820152015260018060a01b031916918284526007602052604084208585526020526040842061ffff82165f5260205261ffff6040600180825f20548487161c161495865f14610b0c5784975b8715610b055784965b825260066020528282209082526020522091165f5260205261ffff60405f2091165f5260205261ffff60405f20549160608260405196610acf88613d2e565b16958681528360208201931683526040810194855201938452604051948552511660208401525160408301525115156060820152f35b8196610a90565b8097610a87565b50346101ee5760203660031901126101ee57610b2d613ba2565b81610180604051610b3d81613cc1565b828152604051610b4c81613ca6565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201526020820152610b80613f81565b60408201528260608201528260808201528260a08201528260c08201528260e082015282610100820152826101208201528261014082015282610160820152015260018060a01b03191681526002602052604081209060405190610be382613cc1565b82546001600160a01b0316825260405191610bfd83613ca6565b600184015461ffff8116845261ffff8160101c16602085015261ffff8160201c16604085015261ffff8160301c16606085015260018060401b038160401c16608085015260018060401b038160801c1660a085015260c01c60c084015260208101928352610c6d60028501613e5e565b60408201908152600485015460608301959060ff81166006811015610e39578752608084019060018060401b038160081c16825260a0850160018060401b038260481c16815260c086019160018060401b039060881c16825260058401549260e087019384526007600686015495610100890196875201549761012088019a61ffff8a168c52610daf6101408a019861ffff8c60101c168a526101608b019a61ffff8d60201c168c5261ffff61018082019d60301c168d5260e06040519e8f9260018060a01b0390511683525161ffff815116602084015261ffff602082015116604084015261ffff604082015116606084015261ffff606082015116608084015260018060401b0360808201511660a084015260018060401b0360a08201511660c084015260c060018060401b0391015116910152516101008d0190613bef565b51906006821015610e2557506101c08a0152516001600160401b039081166101e08a0152905181166102008901529051166102208701525161024086015251610260850152845161ffff908116610280860152905181166102a0850152905181166102c08401528151166102e083015261030082f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b86526021600452602486fd5b50346101ee5760e03660031901126101ee57610e67613ba2565b610e6f613b6f565b6064356044356084356001600160401b03811161130f57610e94903690600401613c3f565b9360a4356001600160401b03811161143957610eb4903690600401613c3f565b90959060c4356001600160401b03811161143557610ed6903690600401613c3f565b6001600160a01b03198a16808c52600260205260408c208054919a90979390916001600160a01b0316156114265760ff600489015416600189015490600681101561141257600214806113fb575b156113ec5760408e8d815260036020522060018060a01b0333165f5260205260ff60405f205416156113dd5761ffff88169c8d1580156113cd575b6113be57610f828f808f60409252600460205220610f7c8b613ff8565b9061400c565b90543360039290921b1c6001600160a01b0316036113135760408f8e815260056020522060018060a01b0333165f5260205260ff600460405f200154166113af57848301969594939291908f610100848a03126101ee5788601f850112156101ee575060405197610ff56101008a613d49565b8861010085019182116113955784905b8282106113855750505087519060a01c1490811591611372575b811561135b575b50801561134d575b801561133f575b8015611331575b611322575f5160206142d15f395f51905f528c8f8c8e6040519060208201928352604082015260408152611071606082613d49565b519020905060405190602082019283527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c81526110bb606c82613d49565b51902006948560c088015103611313578e939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561130f57859361112460405196879586948594635c73957b60e11b865260048601613fd1565b03915afa80156112c5576112f6575b50506101009161200085036112df576108008581116112f257611157368285614060565b6020815191012095611400908183116112ee5781116112ee578161118092850191033691614060565b602081519101208a8d52600b60205260408d2054036112df57916111a69160e0936141a6565b910151036112d057858852600560209081526040808a20335f90815292529020805461ffff60a01b191660a09390931b61ffff60a01b169290921782556003820155600401805460ff19166001179055600701805461121691906112109060101c61ffff16614035565b90613e06565b847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156108dc57818091602460405180948193633c1bcdef60e21b83523360048401525af180156112c5576112ac575b5050604051938452602084015260408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b816112b691613d49565b6112c157845f611271565b8480fd5b6040513d84823e3d90fd5b63d1fed5fd60e01b8852600488fd5b63d1fed5fd60e01b8c5260048cfd5b5f80fd5b8c80fd5b8161130091613d49565b61130b578a5f611133565b8a80fd5b8580fd5b63d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8e5260048efd5b508a60a0870151141561103c565b508960808701511415611035565b508c6060870151141561102e565b905061ffff60408801519160101c1614155f611026565b602088015161ffff82161415915061101f565b8135815260209182019101611005565b50508f80fd5b634e487b7160e01b5f52604160045260245ffd5b6305d252c360e01b8f5260048ffd5b63652122d960e01b8f5260048ffd5b508d61ffff8360101c1610610f5f565b63965c290d60e01b8e5260048efd5b63268dbf6760e21b8e5260048efd5b50608081901c6001600160401b0316431115610f24565b634e487b7160e01b8f52602160045260248ffd5b63d5b25b6360e01b8d5260048dfd5b8980fd5b8780fd5b50346101ee57806003193601126101ee57546040516001600160401b039091168152602090f35b50346101ee576101603660031901126101ee5761147f613ba2565b611487613b80565b90611490613b91565b610124356001600160401b0381116112c1576114b0903690600401613c3f565b90610144356001600160401b0381116118c4576114d1903690600401613c3f565b6001600160a01b03198616885260026020526040882080549294909390926001600160a01b031615611a215760ff60048501541660068110156106e257600303611a12576001600160a01b031987168952600360209081526040808b20335f908152925290205460ff1615611a035761ffff88161580156119eb575b80156119df575b80156119d0575b80156119c5575b6119b7576001600160a01b03198716895260046020526040892061158990610f7c8a613ff8565b90543360039290921b1c6001600160a01b031603611999576001600160a01b031987168952600d60209081526040808b206024358c528252808b2061ffff89165f908152925290205480156119a8576115ec60e43560c43560a435608435614188565b036119995760018060a01b0319871689526007602052604089206024358a526020526040892061ffff87165f5260205260405f20600161ffff8a161b90541661198a5781850194610200818703126114355785601f82011215611435576040519561165961020088613d49565b869061020083011161130b5781905b6102008301821061197a57505060018060a01b031988168a52600a60205260408a2061ffff8a165f5260205260405f205486518960a01c1490811591611969575b8115611956575b8115611946575b8115611933575b8115611922575b8115611911575b8115611908575b81156118e6575b506118d75761012086019361014085519701968751604051906020820192835260408201526040815261170e606082613d49565b51902061010435036118c8578a939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561130f57859361177360405196879586948594635c73957b60e11b865260048601613fd1565b03915afa80156112c5576118af575b5050611819600761ffff9360018060a01b031988168a52600660205260408a206024358b5260205260408a208588165f5260205260405f20858a165f5260205260405f2061010435905560018060a01b031988168a528160205260408a206024358b5260205260408a208588165f5260205260405f206001868b161b81541790550161181384825460201c16614035565b90613e20565b60018060a01b0319851687526008602052604087206024358852602052604087208285165f5260205260405f208261185381835416614035565b1683198254161790555191519281604051961686521660208501526040840152606083015233917f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30260806024359360018060a01b03191692a480f35b816118b991613d49565b6118c457865f611782565b8680fd5b63d1fed5fd60e01b8b5260048bfd5b63d1fed5fd60e01b8a5260048afd5b9050604060e0880151610100890151825191825260208201522014155f6116da565b801591506116d3565b60c088015160a435141591506116cc565b60a0880151608435141591506116c5565b608088015161ffff8c16141591506116be565b60608801516001141591506116b7565b604088015161ffff8a16141591506116b0565b6020880151602435141591506116a9565b8135815260209182019101611668565b633466526160e01b8952600489fd5b63d1fed5fd60e01b8952600489fd5b6346f551f560e01b8a5260048afd5b62d949df60e51b8952600489fd5b506101043515611562565b5061010061ffff87161161155b565b5061ffff861615611554565b5061ffff600185015460101c1661ffff89161161154d565b63965c290d60e01b8952600489fd5b63268dbf6760e21b8952600489fd5b63d5b25b6360e01b8952600489fd5b50346101ee5761ffff6040611a4436613bb9565b949182848051611a5381613d13565b828152826020820152015260018060a01b031916825260096020528282209082526020522091165f52602052606060405f20604051611a9181613d13565b81546040600161ffff83169485855260ff602086019460101c1615158452015492019182526040519283525115156020830152516040820152f35b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611bad5790611b7a575b602090604051908152f35b506020813d602011611ba5575b81611b9460209383613d49565b810103126112ee5760209051611b6f565b3d9150611b87565b604051903d90823e3d90fd5b50346101ee57806003193601126101ee57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101ee576101003660031901126101ee57611c15613ba2565b60a052611c20613b80565b60c05260a4356001600160401b0381116108dc57611c42903690600401613c3f565b90919060c4356001600160401b03811161242257611c64903690600401613c3f565b9060e4356001600160401b0381116112c157611c84903690600401613c3f565b60a0516001600160a01b03191686526002602052604086208054909392906001600160a01b0316156124135760ff60048501541660068110156107cc576003036107a65761ffff60c05116158015612402575b80156123f8575b6123e95760a0516001600160a01b0319168752600d602090815260408089206024358a52825280892060c05161ffff165f90815292529020549283156123da5760018060a01b031960a05116885260086020526040882060243589526020526040882061ffff60c051165f52602052600161ffff60405f2054169501549461ffff8616116123cb5760a0516001600160a01b0319168852600960209081526040808a206024358b528252808a2060c05161ffff165f9081529252902060808190525496601088901c60ff166123bc5788908991604051611dbd81613cf8565b8b8152600160208201526024356122cb575b868601936101a0878603126112f25784601f880112156112f25760405194611df96101a087613d49565b85906101a08901116122c75787905b6101a0890182106122b7575050845160a05160a01c14928315936122a6575b8315612290575b831561227c575b50821561226d575b50811561225d575b8115612249575b508015612237575b8015612226575b8015612215575b6118d75761ffff8716610100830151106118d7576020610100830151116118d7575f5160206142d15f395f51905f5260405160208101906064358252608435604082015260408152611eb5606082613d49565b51902060405160a0516001600160a01b031916602082019081527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c80830193909352918152611f0c606c82613d49565b519020069081610160840151036118c857610c8081036118c8576080116114355760405195611f3c60a088613d49565b608087526020870160808d01973689116112f25760808e83378c60a0820152519020036118d75760648b611f6f926141a6565b61018082015103611999576101000151957f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15611435579089939291611fd760405196879586948594635c73957b60e11b865260048601613fd1565b03915afa801561220a579086916121f5575b505085116121e157610480850185116121e157839060101c61ffff16815b8381106120e4575050505b602081106120855750620100009062ff000019161760805155608435600160805101556040516064358152608435602082015261ffff60c0511690602435907f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1604060018060a01b031960a0511692a480f35b6104808160061b85010160808260051b860101356120d557803515908115916120c5575b506120b657600101612012565b63d1fed5fd60e01b8352600483fd5b600191506020013514155f6120a9565b63d1fed5fd60e01b8452600484fd5b6104808160061b88010161ffff60808360051b8a0101351690811580156121d8575b6112d0576001821b948581166119995785179460018060a01b031960a0511689526007602052604089206024358a526020526040892061ffff60c051165f5260205260405f205416156112d057604051602080820192803584520135604082015260408152612176606082613d49565b5190209060018060a01b031960a05116885260066020526040882060243589526020526040882061ffff60c051165f5260205261ffff60405f2091165f5260205260405f2054036121c957600101612007565b63d1fed5fd60e01b8652600486fd5b50838211612106565b634e487b7160e01b84526011600452602484fd5b816121ff91613d49565b6112c157845f611fe9565b6040513d88823e3d90fd5b506084356101408301511415611e62565b506064356101208301511415611e5b565b5061ffff871660e08301511415611e54565b9050602060c084015191015114155f611e4c565b60a0840151815114159150611e45565b6080850151141591505f611e3d565b606086015160ff909116141592505f611e35565b925061ffff60c051166040860151141592611e2e565b602086015160243514159350611e27565b8135815260209182019101611e08565b8d80fd5b5050905060015460ff8160a01c16156123ad5760405163be5b346360e01b815260a0516001600160a01b0319166004820152602480359082015260c05161ffff1660448201529190608090839060649082906001600160a01b03165afa91821561067d578a908b918c918d95612358575b5091936040519161234c83613cf8565b82526020820152611dcf565b94505050506080823d6080116123a5575b8161237660809383613d49565b810103126114355781519160ff8316830361130b5760208101516040820151606090920151939091905f61233c565b3d9150612369565b63023b34fb60e11b8a5260048afd5b63955c0c4960e01b8952600489fd5b63032cddf960e11b8852600488fd5b6346f551f560e01b8852600488fd5b636d28699160e01b8752600487fd5b5060643515611cde565b5061010061ffff60c0511611611cd7565b63d5b25b6360e01b8752600487fd5b8280fd5b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611bad5790611b7a57602090604051908152f35b50346101ee5760203660031901126101ee576004356001600160a01b038116908190036108dc577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303612531576001549060ff8260a01c16612522578015612513576001600160a81b031990911617600160a01b1760015580f35b63e6c4247b60e01b8352600483fd5b6373253a9760e01b8352600483fd5b6282b42960e81b8252600482fd5b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611bad5790611b7a57602090604051908152f35b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee5761ffff604061260036613bb9565b949160018060a01b031916825260096020528282209082526020522091165f526020526020600160405f200154604051908152f35b50346101ee5760403660031901126101ee5761ffff6040612654613ba2565b9261265d613b6f565b9360018060a01b0319168152600a6020522091165f52602052602060405f2054604051908152f35b50346101ee5760e03660031901126101ee5761269f613ba2565b6084356001600160401b038111612422576126be903690600401613c3f565b60a4356001600160401b0381116112c1576126dd903690600401613c3f565b90919060c4356001600160401b0381116118c4576126ff903690600401613c3f565b6001600160a01b03198716885260026020526040882080549095919391906001600160a01b031615611a215760ff60048701541660068110156106e25760038114612cf25760011901611a12576001860154938460c01c4310612ce35761ffff600788015460101c169161ffff8660201c168310612cd457602435158015612cca575b8015612cc0575b612cb1578185019461012081870312612cad5785601f82011215612cad57604051956101206127b88189613d49565b879083019182116122c75782905b828210612c9d5750505085518b60a01c14801590612c8b575b8015612c76575b8015612c68575b8015612c58575b8015612c48575b8015612c38575b6112df575f5160206142d15f395f51905f528b8d604051602081019060243582526044356040820152606435606082015260608152612842608082613d49565b519020905060405190602082019260018060a01b03191683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c8152612895606c82613d49565b51902006948560e088015103612c29578c939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561130f5785936128fe60405196879586948594635c73957b60e11b865260048601613fd1565b03915afa80156112c557612c10575b50506108a0926201140086036118d7576101000151620104009590946104009187116112ee579392909182899384019088038c968d5b848110612a8157505050505061295993506141a6565b03612a72578201808311612a5e576001600160a01b031984168552600c6020526040852090358155620104208301356001919091015560048101805460ff191660031790556007015460101c61ffff169062010c00810181116121e157835b828110612a0c575050507f5d8b14aea3a3af8564f9576bdf230e2b3aad200f22c6268df330139c5634da5d60606040519260243584526044356020850152606435604085015260018060a01b03191692a280f35b80604062010c0060019360061b8501016020825191803583520135602082015220828060a01b031986168752600a6020526040872061ffff808460051b87013516165f5260205260405f2055016129b8565b634e487b7160e01b85526011600452602485fd5b63d1fed5fd60e01b8552600485fd5b919395509193959661ffff8260051b8d013516158015612bf6575b61132257600161ffff8d8460051b0135161b811661132257908c60408f612af68f9695600161ffff8760051b8a0135161b179660018060a01b0319851683526004602052610f7c61ffff858520928860051b013516613ff8565b90546001600160a01b03199094168252600560209081529290912060039190911b9290921c6001600160a01b03165f90815291905260409020600481015460ff16158015612bda575b6113135761080082028281046108001483151715612bb05760018301808411612bc4576108008102908082046108001490151715612bb057612b88612b8f91600393898b614048565b3691614060565b602081519101209101540361132257600101908b9593919796949297612943565b634e487b7160e01b5f52601160045260245ffd5b5050634e487b7160e01b8f52601160045260248ffd5b5061ffff8d8360051b01351661ffff825460a01c161415612b3f565b5061ffff8360101c1661ffff8d8460051b01351611612a9c565b81612c1a91613d49565b612c2557885f61290d565b8880fd5b63d1fed5fd60e01b8d5260048dfd5b5060643560c08701511415612802565b5060443560a087015114156127fb565b50602435608087015114156127f4565b5083606087015114156127ed565b50604086015161ffff8860101c1614156127e6565b50602086015161ffff881614156127df565b81358152602091820191016127c6565b8b80fd5b63c5f680ed60e01b8b5260048bfd5b5060643515612789565b5060443515612782565b63368f2d7d60e21b8b5260048bfd5b63268dbf6760e21b8a5260048afd5b6337bca76b60e21b8a5260048afd5b50346101ee5760203660031901126101ee57612d466002604060c093612d25613ba2565b612d2d613f81565b506001600160a01b031916815260208390522001613e5e565b612d536040518092613bef565bf35b50346101ee5760203660031901126101ee576040612d79612d74613ba2565b613f19565b60208251918051835201516020820152f35b50346101ee5761ffff6040612d9f36613bb9565b949160018060a01b0319168252600d6020528282209082526020522091165f52602052602060405f2054604051908152f35b50346101ee57806003193601126101ee576020612dec613ec0565b6040516001600160401b039091168152f35b50346112ee5760e03660031901126112ee57612e18613ba2565b602435612e23613b80565b6001600160a01b03199092165f81815260026020526040902080546001600160a01b0316936064359160c4359160a435916084359188156131825760ff600483015416600681101561316e5760030361315f5761ffff169788158015613154575b61314557612e928387614130565b612e9c8585614130565b876130b657612ead60028301613e5e565b908151151590816130ab575b5061309c5760408101516001600160401b03168015159081613089575b506130675760608101516001600160401b03168015159081613076575b506130675760808101516001600160401b03168015159081613054575b506130325760a08101516001600160401b03168015159081613041575b50613032576020015161ffff16801515908161301c575b5061300d575b858952600d60205260408920878a5260205260408920885f5260205260405f2054612ffe579181612fdf600787958c6040612fab8a8860a09c9a7f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc9e614188565b918c8152600d6020528d828220908252602052208d5f5260205260405f20550161ffff600181835460301c16011690613e3f565b604051933385526020850152604084015260608301526080820152a480f35b6316feb18560e11b8952600489fd5b63464e67af60e01b8952600489fd5b905061ffff600783015460301c1610155f612f44565b630410ff2960e31b8a5260048afd5b426001600160401b03161190505f612f2d565b436001600160401b03161190505f612f10565b633deac39560e01b8a5260048afd5b426001600160401b03161090505f612ef3565b436001600160401b03161090505f612ed6565b6330cd747160e01b8a5260048afd5b90503314155f612eb9565b5060015460ff8160a01c1615613136576001600160a01b0316803b156112ee575f60849160405192838092633a54cd5d60e11b82528b60048301528c60248301528d60448301523360648301525afa801561312b57613116575b50612f4a565b6131239199505f90613d49565b5f975f613110565b6040513d5f823e3d90fd5b63023b34fb60e11b5f5260045ffd5b634c4d29cd60e11b5f5260045ffd5b506101008911612e84565b63268dbf6760e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b346112ee575f3660031901126112ee57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346112ee5760203660031901126112ee576001600160a01b03196131f3613ba2565b165f81815260026020526040902080546001600160a01b0316801561318257331415806132a8575b61329a5760040160ff8154166006811015908161316e576003811491821561328d575b821561327c575b505061315f57805460ff191660041790557f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b90915061316e576004148380613245565b506005811491505f61323e565b6282b42960e81b5f5260045ffd5b50600181015460801c6001600160401b031643111561321b565b346112ee576101403660031901126112ee5760043561ffff81168091036112ee576132eb613b6f565b906132f4613b80565b6132fc613b91565b9160c03660831901126112ee576001600160401b03613319613ec0565b16431061315f5780158015613a94575b8015613a87575b8015613a79575b8015613a6d575b8015613a5c575b8015613a4f575b8015613a40575b613925576001600160401b03613367613d6c565b16151580613a27575b806139fb575b8015613996575b801561397d575b61396e57604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561312b575f91613934575b506001600160401b03169283156139255761ffff8091169416926133f98486613dda565b816127100291612710830403612bb0578181106138f95750505f19905b5f54946001600160401b03808716908114612bb0576001600160401b03600191820181166001600160401b03199890981688175f556001600160a01b031963ffffffff60401b7f000000000000000000000000000000000000000000000000000000000000000060401b16891760a01b169643821696928701939092918411612bb0576134c37f000000000000000000000000000000000000000000000000000000000000000088613ded565b9061ffff6134f17f00000000000000000000000000000000000000000000000000000000000000008a613ded565b9361351c7f00000000000000000000000000000000000000000000000000000000000000008b613ded565b956040519861352a8a613ca6565b895260208901521660408088019190915260608701919091526001600160401b03918216608087015291811660a08601529190911660c08401525161356e81613cc1565b338152602081019283526040519061358582613cdd565b60843580151581036112ee57825260a43561ffff811681036112ee57602083015260c4356001600160401b03811681036112ee57604083015260e4356001600160401b03811681036112ee576060830152610104356001600160401b03811681036112ee576080830152610124356001600160401b03811681036112ee5760a08301526040810191825286606082016001815260808301998a5260a083019088825260c084019560018060401b03169687875260e08501915f83526101008601938a85526101208701975f89526101408801905f82526101608901945f86526101808a01985f8a525f52600260205260405f209960018060a01b0390511660018060a01b03198b5416178a5560018a01905161ffff808251161661ffff198354161782556136bb61ffff60208301511683613e06565b6136cd61ffff60408301511683613e20565b6136df61ffff60608301511683613e3f565b608081810151835460a08085015160c0958601516001600160401b03938416604095861b600160401b600160801b03161791861b600160801b600160c01b031691909117951b6001600160c01b031916949094179094559351805160028e0180546020840151978401516060850151958501516001600160d81b031990921660ff941515949094169390931760089890981b62ffff00169790971760189290921b6301000000600160581b03169190911760589390931b600160581b600160981b03169290921760989590951b600160981b600160d81b0316949094179055919091015160038a0180546001600160401b0319169190921617905591519860048801929060068b101561316e5760209f600761ffff9a613887998c9961388f9f988b988997889761387f9c60ff801986541691161784555191835490600160481b600160881b03905160481b1691600160881b600160c81b03905160881b1692610100600160481b039060081b1690610100600160c81b0319161717179055516005830155516006820155019a51161682198a5416178955511687613e06565b511684613e20565b511690613e3f565b5f8054600160401b600160801b03191643604090811b600160401b600160801b031691909117909155805193845285840191909152820152339082907f1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d90606090a3604051908152f35b81156139115761390b915f1904613dda565b90613416565b634e487b7160e01b5f52601260045260245ffd5b63d06b96b160e01b5f5260045ffd5b90506020813d602011613966575b8161394f60209383613d49565b810103126112ee5761396090613dc6565b856133d5565b3d9150613942565b63148b7e9360e31b5f5260045ffd5b5060a43561ffff81168091036112ee5761010010613384565b506001600160401b036139a7613d99565b161515806139e2575b801561337d57506139bf613daf565b6001600160401b036139cf613d99565b166001600160401b03909116111561337d565b506001600160401b036139f3613daf565b1615156139b0565b50613a04613d82565b6001600160401b03613a14613d6c565b166001600160401b039091161115613376565b506001600160401b03613a38613d82565b161515613370565b5061271061ffff841610613353565b508061ffff83161061334c565b5061ffff841661ffff831611613345565b5061ffff82161561333e565b50602061ffff851611613337565b5061ffff84168111613330565b5061ffff841615613329565b346112ee575f3660031901126112ee5760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561312b575f90611b7a57602090604051908152f35b346112ee575f3660031901126112ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346112ee575f3660031901126112ee575f5460401c6001600160401b03168152602090f35b6024359061ffff821682036112ee57565b6044359061ffff821682036112ee57565b6064359061ffff821682036112ee57565b600435906001600160a01b0319821682036112ee57565b60609060031901126112ee576004356001600160a01b0319811681036112ee57906024359060443561ffff811681036112ee5790565b80511515825260208082015161ffff16908301526040808201516001600160401b039081169184019190915260608083015182169084015260808083015182169084015260a09182015116910152565b9181601f840112156112ee578235916001600160401b0383116112ee57602083818601950101116112ee57565b346112ee575f3660031901126112ee5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60e081019081106001600160401b0382111761139b57604052565b6101a081019081106001600160401b0382111761139b57604052565b60c081019081106001600160401b0382111761139b57604052565b604081019081106001600160401b0382111761139b57604052565b606081019081106001600160401b0382111761139b57604052565b608081019081106001600160401b0382111761139b57604052565b601f909101601f19168101906001600160401b0382119082101761139b57604052565b60c4356001600160401b03811681036112ee5790565b610104356001600160401b03811681036112ee5790565b60e4356001600160401b03811681036112ee5790565b610124356001600160401b03811681036112ee5790565b51906001600160401b03821682036112ee57565b81810292918115918404141715612bb057565b6001600160401b039182169082160191908211612bb057565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b90604051613e6b81613cdd565b825460ff81161515825261ffff600882901c1660208301526001600160401b03601882901c81166040840152605882901c8116606084015260989190911c8116608083015260019093015490921660a0830152565b5f5460401c6001600160401b03168015613f0b57613f08907f00000000000000000000000000000000000000000000000000000000000000006001600160401b031690613ded565b90565b50436001600160401b031690565b604051613f2581613cf8565b5f81525f60208201525060018060a01b0319165f52600c60205260405f2060018101548015613f665760405191613f5b83613cf8565b548252602082015290565b5050604051613f7481613cf8565b5f81526001602082015290565b60405190613f8e82613cdd565b5f60a0838281528260208201528260408201528260608201528260808201520152565b908060209392818452848401375f828201840152601f01601f1916010190565b9290613fea90613f089593604086526040860191613fb1565b926020818503910152613fb1565b61ffff5f199116019061ffff8211612bb057565b8054821015614021575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b61ffff1661ffff8114612bb05760010190565b909392938483116112ee5784116112ee578101920390565b9192916001600160401b03821161139b5760405191614089601f8201601f191660200184613d49565b8294818452818301116112ee578281602093845f960137010152565b908160a09103126112ee576040519060a082016001600160401b0381118382101761139b5760405280516001600160a01b03811681036112ee5782526020810151602083015260408101516040830152606081015160038110156112ee57606083015261411490608001613dc6565b608082015290565b61ffff60019116019061ffff8211612bb057565b905f5160206142d15f395f51905f52821080614172575b1561314557811580614168575b613145576141619161420b565b1561314557565b5060018114614154565b505f5160206142d15f395f51905f528110614147565b91608093916040519384526020840152604083015260608201522090565b9291905f5160206142d15f395f51905f525f940691829060051b8201915b8281106141d15750505050565b909192945f5160206142d15f395f51905f5283816020938186358b0990089709939291016141c4565b9060408110156140215760051b0190565b5f5160206142d15f395f51905f5281108015906142b9575b6142b3575f5160206142d15f395f51905f5281819209918009905f5160206142d15f395f51905f5281065f5160206142d15f395f51905f5203915f5160206142d15f395f51905f528311612bb0575f5160206142d15f395f51905f52808080848196097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09600108935f08081490565b50505f90565b505f5160206142d15f395f51905f5282101561422356fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212200052b456cf283b2fc70e82b0118515fba3a3d7b47747779c7d669fe63739f1bf64736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _finalizeVerifier common.Address, _decryptCombineVerifier common.Address, _epochDurationBlocks *big.Int, _committeeSelectionBlocks *big.Int, _keyAssemblyBlocks *big.Int, _finalizeGapBlocks *big.Int) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _finalizeVerifier, _decryptCombineVerifier, _epochDurationBlocks, _committeeSelectionBlocks, _keyAssemblyBlocks, _finalizeGapBlocks)
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

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 epochId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerCaller) GetCollectivePublicKey(opts *bind.CallOpts, epochId [12]byte) (DKGTypesPoint, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCollectivePublicKey", epochId)

	if err != nil {
		return *new(DKGTypesPoint), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesPoint)).(*DKGTypesPoint)

	return out0, err

}

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 epochId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerSession) GetCollectivePublicKey(epochId [12]byte) (DKGTypesPoint, error) {
	return _DKGManager.Contract.GetCollectivePublicKey(&_DKGManager.CallOpts, epochId)
}

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 epochId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerCallerSession) GetCollectivePublicKey(epochId [12]byte) (DKGTypesPoint, error) {
	return _DKGManager.Contract.GetCollectivePublicKey(&_DKGManager.CallOpts, epochId)
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
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
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
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetContribution(epochId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	return _DKGManager.Contract.GetContribution(&_DKGManager.CallOpts, epochId, contributor)
}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 epochId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
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

// GetDecryptionPolicy is a free data retrieval call binding the contract method 0x4554c0be.
//
// Solidity: function getDecryptionPolicy(bytes12 epochId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerCaller) GetDecryptionPolicy(opts *bind.CallOpts, epochId [12]byte) (DKGTypesDecryptionPolicy, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getDecryptionPolicy", epochId)

	if err != nil {
		return *new(DKGTypesDecryptionPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesDecryptionPolicy)).(*DKGTypesDecryptionPolicy)

	return out0, err

}

// GetDecryptionPolicy is a free data retrieval call binding the contract method 0x4554c0be.
//
// Solidity: function getDecryptionPolicy(bytes12 epochId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerSession) GetDecryptionPolicy(epochId [12]byte) (DKGTypesDecryptionPolicy, error) {
	return _DKGManager.Contract.GetDecryptionPolicy(&_DKGManager.CallOpts, epochId)
}

// GetDecryptionPolicy is a free data retrieval call binding the contract method 0x4554c0be.
//
// Solidity: function getDecryptionPolicy(bytes12 epochId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerCallerSession) GetDecryptionPolicy(epochId [12]byte) (DKGTypesDecryptionPolicy, error) {
	return _DKGManager.Contract.GetDecryptionPolicy(&_DKGManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetEpoch(epochId [12]byte) (IDKGManagerEpoch, error) {
	return _DKGManager.Contract.GetEpoch(&_DKGManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetShareCommitmentHash(opts *bind.CallOpts, epochId [12]byte, participantIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getShareCommitmentHash", epochId, participantIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetShareCommitmentHash(epochId [12]byte, participantIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetShareCommitmentHash(&_DKGManager.CallOpts, epochId, participantIndex)
}

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetShareCommitmentHash(epochId [12]byte, participantIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetShareCommitmentHash(&_DKGManager.CallOpts, epochId, participantIndex)
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

// CreateEpoch is a paid mutator transaction binding the contract method 0x143b4616.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateEpoch(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createEpoch", threshold, committeeSize, minValidContributions, lotteryAlphaBps, decryptionPolicy)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x143b4616.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, decryptionPolicy)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x143b4616.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, decryptionPolicy)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x49c61a12.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) FinalizeEpoch(opts *bind.TransactOpts, epochId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "finalizeEpoch", epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x49c61a12.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) FinalizeEpoch(epochId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x49c61a12.
//
// Solidity: function finalizeEpoch(bytes12 epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) FinalizeEpoch(epochId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeEpoch(&_DKGManager.TransactOpts, epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
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

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x2648f567.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactor) SubmitCiphertext(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitCiphertext", epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x2648f567.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x2648f567.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y)
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

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0xa305e0f3.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitPartialDecryption(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitPartialDecryption", epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0xa305e0f3.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0xa305e0f3.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input)
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
	EpochId                  [12]byte
	AggregateCommitmentsHash [32]byte
	CollectivePublicKeyHash  [32]byte
	ShareCommitmentHash      [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterEpochLive is a free log retrieval operation binding the contract event 0x5d8b14aea3a3af8564f9576bdf230e2b3aad200f22c6268df330139c5634da5d.
//
// Solidity: event EpochLive(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
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

// WatchEpochLive is a free log subscription operation binding the contract event 0x5d8b14aea3a3af8564f9576bdf230e2b3aad200f22c6268df330139c5634da5d.
//
// Solidity: event EpochLive(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
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

// ParseEpochLive is a log parse operation binding the contract event 0x5d8b14aea3a3af8564f9576bdf230e2b3aad200f22c6268df330139c5634da5d.
//
// Solidity: event EpochLive(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
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
