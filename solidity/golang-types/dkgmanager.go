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
	Threshold                 uint16
	CommitteeSize             uint16
	MinValidContributions     uint16
	LotteryAlphaBps           uint16
	RegistrationDeadlineBlock uint64
	ContributionDeadlineBlock uint64
	FinalizeNotBeforeBlock    uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epochDurationBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochDurationBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochEvicted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochFinalized\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x61020080604052346103345760e081614c7480380380916100208285610338565b8339810103126103345780519063ffffffff821691828103610334576100486020830161036f565b6100546040840161036f565b6100606060850161036f565b9161006d6080860161036f565b9360c061007c60a0880161036f565b9601519663ffffffff461603610325576001600160a01b03821615610316576001600160a01b038316158015610305575b80156102f4575b80156102e3575b6102d45763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b6024820152601881526100f7603882610338565b5190201660c05260e052610100526101205261014052806102cf575060645b6101f481028115908281046101f4148217156102bb5761271090046001600160401b0316906105dc8381029184830490911417156102bb5761271090046001600160401b031691600182118015906102b3575b80156102ab575b801561029b575b61028c576101a5926101a0916101605282610180526101968184610383565b6101a05282610383565b610383565b6101c052336101e0526040516148d7908161039d823960805181611bd8015260a0518181816102e70152818161121a0152818161337e0152613b97015260c0518181816131af0152613424015260e0518181816110d3015281816125bf0152613b410152610100518181816109bf015281816117220152611b3a015261012051818181611ae40152818161256801526128ae0152610140518181816101bf01528181611f84015261244f015261016051818181613cff0152613f84015261018051816134d901526101a0518161350701526101c0518161353201526101e051816124b90152f35b63d06b96b160e01b5f5260045ffd5b506001600160401b038111610177565b508115610170565b508215610169565b634e487b7160e01b5f52601160045260245ffd5b610116565b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100bb565b506001600160a01b038516156100b4565b506001600160a01b038416156100ad565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761035b57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361033457565b6001600160401b0391821690821601919082116102bb5756fe60e0806040526004361015610012575f80fd5b5f905f3560e01c90816304da574014613bc65750806306433b1b14613b82578063074a75e114613b1c578063143b4616146132a357806318287e5f146131d357806323488be5146131935780632648f56714612e00578063268ae2a114612dd35780632de546d514612d8d5780633353ec6e14612d575780634554c0be14612d0357806349c61a1214612687578063510ba2df146126375780635a8f2bb3146125ee57806363f314cd146125a9578063669a76a9146125415780636d16897d1461249057806371712c29146101f157806372517b4b1461242857806377235ee114611bfc57806385e1f4d014611bbb5780638dc1f53a14611b1357806393c3d3a814611ace5780639bbada6714611a32578063a305e0f314611465578063a4adcd7f1461143e578063b7bca61514610e4e578063be59b8ea14610b14578063bea5210d146109ee578063bf192209146109a9578063ca3c0458146108e1578063d3720aac146107f0578063d99337671461021f578063ebe86c13146101f6578063fa8f5e96146101f15763fe1604b5146101aa575f80fd5b346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b80fd5b613ce8565b50346101ee57806003193601126101ee576001546040516001600160a01b039091168152602090f35b50346101ee5760203660031901126101ee576001600160a01b0319610242613c1e565b1680825260236020526040822080546001600160a01b0316156107e157600481019081549060ff8216600182019182549160068110156107cd57600114806107b6575b156107a757600781019061ffff808354169360101c1683101561079857868852602460209081526040808a20335f908152925290205460ff1661078957600581018054958615610716575b50506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316959060a0816024818a5afa801561067e576060918b916106f7575b50015160038110156106e3575f19016106d457600690604051602081019182523360601b604082015260348152610362605482613dc5565b51902091015411156106c557858752602560205260408720805490600160401b8210156106b1579261ffff959493926103a5838895600161042b960181556140b6565b81546001600160a01b0360039290921b91821b1916339182901b17909155898b52602460209081526040808d205f9384529091529020805460ff19166001179055836103f0836141b2565b168419825416179055604051818152887f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a36141b2565b915460101c169216821461043d578480f35b6040519161040061044e8185613dc5565b368437604051916108006104628185613dc5565b368437858752602560205260408720875b83811061057c575050505b6020811061053657506040519060208201928387905b6020821061052057505050610420820186905b6040821061050a57505050610c0081526104c3610c2082613dc5565b519020828452602c6020526040842055805460ff191660021790557fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13588280a25f8080808480f35b60208060019285518152019301910190916104a7565b6020806001928551815201930191019091610494565b8060011b8181046002148215171561056857600181018091116105685790600161056181938561478b565b520161047e565b634e487b7160e01b87526011600452602487fd5b6001810180821161069d576020821015610689578160051b8701526105a181836140b6565b90546040516313a4120960e31b815260039290921b1c6001600160a01b0316600482015260a081602481875afa90811561067e578a91610650575b506020810151908260011b918383046002148415171561063c5790604091610604848a61478b565b52015190600181018091116106285790610621600193928861478b565b5201610473565b634e487b7160e01b8b52601160045260248bfd5b634e487b7160e01b8c52601160045260248cfd5b610671915060a03d8111610677575b6106698183613dc5565b81019061413b565b5f6105dc565b503d61065f565b6040513d8c823e3d90fd5b634e487b7160e01b8a52603260045260248afd5b634e487b7160e01b8a52601160045260248afd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b63aba4733960e01b8952600489fd5b634e487b7160e01b8a52602160045260248afd5b610710915060a03d60a011610677576106698183613dc5565b5f61032a565b90955060881c6001600160401b03164381101561077a574094851561076b57859055867fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602087604051908152a25f806102d0565b6302504bb360e61b8952600489fd5b63172181cb60e21b8952600489fd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b50604082901c6001600160401b0316431115610285565b634e487b7160e01b88526021600452602488fd5b63d5b25b6360e01b8352600483fd5b50346101ee5760403660031901126101ee5761080a613c1e565b60243591906001600160a01b03831683036108dd579060409161082b61402b565b5060018060a01b03191681526026602052209060018060a01b03165f5260205260c060405f2060405161085d81613d59565b81549160018060a01b0383169283835261ffff602084019160a01c16815260018201546040840190815261ffff6002840154926060860193845260a060ff600460038801549760808a01988952015416960195151586526040519687525116602086015251604085015251606084015251608083015251151560a0820152f35b5080fd5b50346101ee5760203660031901126101ee576001600160a01b0319610904613c1e565b168152602560205260408120604051908160208254918281520190819285526020852090855b81811061098a5750505082610940910383613dc5565b604051928392602084019060208552518091526040840192915b818110610968575050500390f35b82516001600160a01b031684528594506020938401939092019160010161095a565b82546001600160a01b031684526020909301926001928301920161092a565b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee5760803660031901126101ee57608090610a0b613c1e565b9060243591610a18613bfc565b610a20613c0d565b91836060604051610a3081613daa565b828152826020820152826040820152015260018060a01b031916918284526028602052604084208585526020526040842061ffff82165f5260205261ffff6040600180825f20548487161c161495865f14610b0d5784975b8715610b065784965b825260276020528282209082526020522091165f5260205261ffff60405f2091165f5260205261ffff60405f20549160608260405196610ad088613daa565b16958681528360208201931683526040810194855201938452604051948552511660208401525160408301525115156060820152f35b8196610a91565b8097610a88565b50346101ee5760203660031901126101ee57610b2e613c1e565b81610180604051610b3e81613d3d565b828152604051610b4d81613d22565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201526020820152610b8161402b565b60408201528260608201528260808201528260a08201528260c08201528260e082015282610100820152826101208201528261014082015282610160820152015260018060a01b03191681526023602052604081209060405190610be482613d3d565b82546001600160a01b0316825260405191610bfe83613d22565b600184015461ffff8116845261ffff8160101c16602085015261ffff8160201c16604085015261ffff8160301c16606085015260018060401b038160401c16608085015260018060401b038160801c1660a085015260c01c60c084015260208101928352610c6e60028501613f08565b60408201908152600485015460608301959060ff81166006811015610e3a578752608084019060018060401b038160081c16825260a0850160018060401b038260481c16815260c086019160018060401b039060881c16825260058401549260e087019384526007600686015495610100890196875201549761012088019a61ffff8a168c52610db06101408a019861ffff8c60101c168a526101608b019a61ffff8d60201c168c5261ffff61018082019d60301c168d5260e06040519e8f9260018060a01b0390511683525161ffff815116602084015261ffff602082015116604084015261ffff604082015116606084015261ffff606082015116608084015260018060401b0360808201511660a084015260018060401b0360a08201511660c084015260c060018060401b0391015116910152516101008d0190613c6b565b51906006821015610e2657506101c08a0152516001600160401b039081166101e08a0152905181166102008901529051166102208701525161024086015251610260850152845161ffff908116610280860152905181166102a0850152905181166102c08401528151166102e083015261030082f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b86526021600452602486fd5b50346101ee5760e03660031901126101ee57610e68613c1e565b610e70613beb565b6064356044356084356001600160401b03811161131057610e95903690600401613cbb565b9360a4356001600160401b03811161143a57610eb5903690600401613cbb565b90959060c4356001600160401b03811161143657610ed7903690600401613cbb565b6001600160a01b03198a16808c52602360205260408c208054919a90979390916001600160a01b0316156114275760ff600489015416600189015490600681101561141357600214806113fc575b156113ed5760408e8d815260246020522060018060a01b0333165f5260205260ff60405f205416156113de5761ffff88169c8d1580156113ce575b6113bf57610f838f808f60409252602560205220610f7d8b6140a2565b906140b6565b90543360039290921b1c6001600160a01b0316036113145760408f8e815260266020522060018060a01b0333165f5260205260ff600460405f200154166113b057848301969594939291908f610100848a03126101ee5788601f850112156101ee575060405197610ff66101008a613dc5565b8861010085019182116113965784905b8282106113865750505087519060a01c1490811591611373575b811561135c575b50801561134e575b8015611340575b8015611332575b611323575f5160206148625f395f51905f528c8f8c8e6040519060208201928352604082015260408152611072606082613dc5565b519020905060405190602082019283527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c81526110bc606c82613dc5565b51902006948560c088015103611314578e939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561131057859361112560405196879586948594635c73957b60e11b86526004860161407b565b03915afa80156112c6576112f7575b50506101009161200085036112e0576108008581116112f3576111583682856140f6565b6020815191012095611400908183116112ef5781116112ef5781611181928501910336916140f6565b602081519101208a8d52602c60205260408d2054036112e057916111a79160e093614737565b910151036112d157858852602660209081526040808a20335f90815292529020805461ffff60a01b191660a09390931b61ffff60a01b169290921782556003820155600401805460ff19166001179055600701805461121791906112119060101c61ffff166140cb565b90613eb0565b847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156108dd57818091602460405180948193633c1bcdef60e21b83523360048401525af180156112c6576112ad575b5050604051938452602084015260408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b816112b791613dc5565b6112c257845f611272565b8480fd5b6040513d84823e3d90fd5b63d1fed5fd60e01b8852600488fd5b63d1fed5fd60e01b8c5260048cfd5b5f80fd5b8c80fd5b8161130191613dc5565b61130c578a5f611134565b8a80fd5b8580fd5b63d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8e5260048efd5b508a60a0870151141561103d565b508960808701511415611036565b508c6060870151141561102f565b905061ffff60408801519160101c1614155f611027565b602088015161ffff821614159150611020565b8135815260209182019101611006565b50508f80fd5b634e487b7160e01b5f52604160045260245ffd5b6305d252c360e01b8f5260048ffd5b63652122d960e01b8f5260048ffd5b508d61ffff8360101c1610610f60565b63965c290d60e01b8e5260048efd5b63268dbf6760e21b8e5260048efd5b50608081901c6001600160401b0316431115610f25565b634e487b7160e01b8f52602160045260248ffd5b63d5b25b6360e01b8d5260048dfd5b8980fd5b8780fd5b50346101ee57806003193601126101ee57546040516001600160401b039091168152602090f35b50346101ee576101603660031901126101ee57611480613c1e565b611488613bfc565b90611491613c0d565b610124356001600160401b0381116112c2576114b1903690600401613cbb565b90610144356001600160401b0381116118c6576114d2903690600401613cbb565b6001600160a01b03198616885260236020526040882080549294909390926001600160a01b031615611a235760ff60048501541660068110156106e357600303611a14576001600160a01b031987168952602460209081526040808b20335f908152925290205460ff1615611a055761ffff88161580156119ed575b80156119e1575b80156119d2575b80156119c7575b6119b9576001600160a01b03198716895260256020526040892061158a90610f7d8a6140a2565b90543360039290921b1c6001600160a01b03160361199b576001600160a01b031987168952602e60209081526040808b206024358c528252808b2061ffff89165f908152925290205480156119aa576115ed60e43560c43560a435608435614719565b0361199b5760018060a01b0319871689526028602052604089206024358a526020526040892061ffff87165f5260205260405f20600161ffff8a161b90541661198c5781850194610200818703126114365785601f82011215611436576040519561165a61020088613dc5565b869061020083011161130c5781905b6102008301821061197c57505060018060a01b031988168a52602b60205260408a2061ffff8a165f5260205260405f205486518960a01c149081159161196b575b8115611958575b8115611948575b8115611935575b8115611924575b8115611913575b811561190a575b81156118e8575b506118d95761012086019361014085519701968751604051906020820192835260408201526040815261170f606082613dc5565b51902061010435036118ca578a939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561131057859361177460405196879586948594635c73957b60e11b86526004860161407b565b03915afa80156112c6576118b1575b505061181b600761ffff9360018060a01b031988168a52602760205260408a206024358b5260205260408a208588165f5260205260405f20858a165f5260205260405f2061010435905560018060a01b031988168a52602860205260408a206024358b5260205260408a208588165f5260205260405f206001868b161b81541790550161181584825460201c166140cb565b90613eca565b60018060a01b0319851687526029602052604087206024358852602052604087208285165f5260205260405f2082611855818354166140cb565b1683198254161790555191519281604051961686521660208501526040840152606083015233917f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30260806024359360018060a01b03191692a480f35b816118bb91613dc5565b6118c657865f611783565b8680fd5b63d1fed5fd60e01b8b5260048bfd5b63d1fed5fd60e01b8a5260048afd5b9050604060e0880151610100890151825191825260208201522014155f6116db565b801591506116d4565b60c088015160a435141591506116cd565b60a0880151608435141591506116c6565b608088015161ffff8c16141591506116bf565b60608801516001141591506116b8565b604088015161ffff8a16141591506116b1565b6020880151602435141591506116aa565b8135815260209182019101611669565b633466526160e01b8952600489fd5b63d1fed5fd60e01b8952600489fd5b6346f551f560e01b8a5260048afd5b62d949df60e51b8952600489fd5b506101043515611563565b5061010061ffff87161161155c565b5061ffff861615611555565b5061ffff600185015460101c1661ffff89161161154e565b63965c290d60e01b8952600489fd5b63268dbf6760e21b8952600489fd5b63d5b25b6360e01b8952600489fd5b50346101ee5761ffff6040611a4636613c35565b949182848051611a5581613d8f565b828152826020820152015260018060a01b0319168252602a6020528282209082526020522091165f52602052606060405f20604051611a9381613d8f565b81546040600161ffff83169485855260ff602086019460101c1615158452015492019182526040519283525115156020830152516040820152f35b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611baf5790611b7c575b602090604051908152f35b506020813d602011611ba7575b81611b9660209383613dc5565b810103126112ef5760209051611b71565b3d9150611b89565b604051903d90823e3d90fd5b50346101ee57806003193601126101ee57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101ee576101003660031901126101ee57611c17613c1e565b60a052611c22613bfc565b60c05260a4356001600160401b0381116108dd57611c44903690600401613cbb565b90919060c4356001600160401b03811161242457611c66903690600401613cbb565b9060e4356001600160401b0381116112c257611c86903690600401613cbb565b60a0516001600160a01b03191686526023602052604086208054909392906001600160a01b0316156124155760ff60048501541660068110156107cd576003036107a75761ffff60c05116158015612404575b80156123fa575b6123eb5760a0516001600160a01b0319168752602e602090815260408089206024358a52825280892060c05161ffff165f90815292529020549283156123dc5760018060a01b031960a05116885260296020526040882060243589526020526040882061ffff60c051165f52602052600161ffff60405f2054169501549461ffff8616116123cd5760a0516001600160a01b0319168852602a60209081526040808a206024358b528252808a2060c05161ffff165f9081529252902060808190525496601088901c60ff166123be5788908991604051611dbf81613d74565b8b8152600160208201526024356122cd575b868601936101a0878603126112f35784601f880112156112f35760405194611dfb6101a087613dc5565b85906101a08901116122c95787905b6101a0890182106122b9575050845160a05160a01c14928315936122a8575b8315612292575b831561227e575b50821561226f575b50811561225f575b811561224b575b508015612239575b8015612228575b8015612217575b6118d95761ffff8716610100830151106118d9576020610100830151116118d9575f5160206148625f395f51905f5260405160208101906064358252608435604082015260408152611eb7606082613dc5565b51902060405160a0516001600160a01b031916602082019081527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c80830193909352918152611f0e606c82613dc5565b519020069081610160840151036118ca57610c8081036118ca576080116114365760405195611f3e60a088613dc5565b608087526020870160808d01973689116112f35760808e83378c60a0820152519020036118d95760648b611f7192614737565b6101808201510361199b576101000151957f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15611436579089939291611fd960405196879586948594635c73957b60e11b86526004860161407b565b03915afa801561220c579086916121f7575b505085116121e357610480850185116121e357839060101c61ffff16815b8381106120e6575050505b602081106120875750620100009062ff000019161760805155608435600160805101556040516064358152608435602082015261ffff60c0511690602435907f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1604060018060a01b031960a0511692a480f35b6104808160061b85010160808260051b860101356120d757803515908115916120c7575b506120b857600101612014565b63d1fed5fd60e01b8352600483fd5b600191506020013514155f6120ab565b63d1fed5fd60e01b8452600484fd5b6104808160061b88010161ffff60808360051b8a0101351690811580156121da575b6112d1576001821b9485811661199b5785179460018060a01b031960a0511689526028602052604089206024358a526020526040892061ffff60c051165f5260205260405f205416156112d157604051602080820192803584520135604082015260408152612178606082613dc5565b5190209060018060a01b031960a05116885260276020526040882060243589526020526040882061ffff60c051165f5260205261ffff60405f2091165f5260205260405f2054036121cb57600101612009565b63d1fed5fd60e01b8652600486fd5b50838211612108565b634e487b7160e01b84526011600452602484fd5b8161220191613dc5565b6112c257845f611feb565b6040513d88823e3d90fd5b506084356101408301511415611e64565b506064356101208301511415611e5d565b5061ffff871660e08301511415611e56565b9050602060c084015191015114155f611e4e565b60a0840151815114159150611e47565b6080850151141591505f611e3f565b606086015160ff909116141592505f611e37565b925061ffff60c051166040860151141592611e30565b602086015160243514159350611e29565b8135815260209182019101611e0a565b8d80fd5b5050905060015460ff8160a01c16156123af5760405163be5b346360e01b815260a0516001600160a01b0319166004820152602480359082015260c05161ffff1660448201529190608090839060649082906001600160a01b03165afa91821561067e578a908b918c918d9561235a575b5091936040519161234e83613d74565b82526020820152611dd1565b94505050506080823d6080116123a7575b8161237860809383613dc5565b810103126114365781519160ff8316830361130c5760208101516040820151606090920151939091905f61233e565b3d915061236b565b63023b34fb60e11b8a5260048afd5b63955c0c4960e01b8952600489fd5b63032cddf960e11b8852600488fd5b6346f551f560e01b8852600488fd5b636d28699160e01b8752600487fd5b5060643515611ce0565b5061010061ffff60c0511611611cd9565b63d5b25b6360e01b8752600487fd5b8280fd5b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611baf5790611b7c57602090604051908152f35b50346101ee5760203660031901126101ee576004356001600160a01b038116908190036108dd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303612533576001549060ff8260a01c16612524578015612515576001600160a81b031990911617600160a01b1760015580f35b63e6c4247b60e01b8352600483fd5b6373253a9760e01b8352600483fd5b6282b42960e81b8252600482fd5b50346101ee57806003193601126101ee5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611baf5790611b7c57602090604051908152f35b50346101ee57806003193601126101ee576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346101ee5761ffff604061260236613c35565b949160018060a01b0319168252602a6020528282209082526020522091165f526020526020600160405f200154604051908152f35b50346101ee5760403660031901126101ee5761ffff6040612656613c1e565b9261265f613beb565b9360018060a01b0319168152602b6020522091165f52602052602060405f2054604051908152f35b50346101ee5760e03660031901126101ee576126a1613c1e565b6084356001600160401b038111612424576126c0903690600401613cbb565b60a4356001600160401b0381116112c2576126df903690600401613cbb565b90919060c4356001600160401b0381116118c657612701903690600401613cbb565b6001600160a01b03198716885260236020526040882080549095919391906001600160a01b031615611a235760ff60048701541660068110156106e35760038114612cf45760011901611a14576001860154938460c01c4310612ce55761ffff600788015460101c169161ffff8660201c168310612cd657602435158015612ccc575b8015612cc2575b612cb3578185019461012081870312612caf5785601f82011215612caf57604051956101206127ba8189613dc5565b879083019182116122c95782905b828210612c9f5750505085518b60a01c14801590612c8d575b8015612c78575b8015612c6a575b8015612c5a575b8015612c4a575b8015612c3a575b6112e0575f5160206148625f395f51905f528b8d604051602081019060243582526044356040820152606435606082015260608152612844608082613dc5565b519020905060405190602082019260018060a01b03191683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c8152612897606c82613dc5565b51902006948560e088015103612c2b578c939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561131057859361290060405196879586948594635c73957b60e11b86526004860161407b565b03915afa80156112c657612c12575b50506108a0926201140086036118d9576101000151620104009590946104009187116112ef579392909182899384019088038c968d5b848110612a8357505050505061295b9350614737565b03612a74578201808311612a60576001600160a01b031984168552602d6020526040852090358155620104208301356001919091015560048101805460ff191660031790556007015460101c61ffff169062010c00810181116121e357835b828110612a0e575050507f4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a47260606040519260243584526044356020850152606435604085015260018060a01b03191692a280f35b80604062010c0060019360061b8501016020825191803583520135602082015220828060a01b031986168752602b6020526040872061ffff808460051b87013516165f5260205260405f2055016129ba565b634e487b7160e01b85526011600452602485fd5b63d1fed5fd60e01b8552600485fd5b919395509193959661ffff8260051b8d013516158015612bf8575b61132357600161ffff8d8460051b0135161b811661132357908c60408f612af88f9695600161ffff8760051b8a0135161b179660018060a01b0319851683526025602052610f7d61ffff858520928860051b0135166140a2565b90546001600160a01b03199094168252602660209081529290912060039190911b9290921c6001600160a01b03165f90815291905260409020600481015460ff16158015612bdc575b6113145761080082028281046108001483151715612bb25760018301808411612bc6576108008102908082046108001490151715612bb257612b8a612b9191600393898b6140de565b36916140f6565b602081519101209101540361132357600101908b9593919796949297612945565b634e487b7160e01b5f52601160045260245ffd5b5050634e487b7160e01b8f52601160045260248ffd5b5061ffff8d8360051b01351661ffff825460a01c161415612b41565b5061ffff8360101c1661ffff8d8460051b01351611612a9e565b81612c1c91613dc5565b612c2757885f61290f565b8880fd5b63d1fed5fd60e01b8d5260048dfd5b5060643560c08701511415612804565b5060443560a087015114156127fd565b50602435608087015114156127f6565b5083606087015114156127ef565b50604086015161ffff8860101c1614156127e8565b50602086015161ffff881614156127e1565b81358152602091820191016127c8565b8b80fd5b63c5f680ed60e01b8b5260048bfd5b506064351561278b565b5060443515612784565b63368f2d7d60e21b8b5260048bfd5b63268dbf6760e21b8a5260048afd5b63475a253560e01b8a5260048afd5b50346101ee5760203660031901126101ee57612d486002604060c093612d27613c1e565b612d2f61402b565b506001600160a01b031916815260236020522001613f08565b612d556040518092613c6b565bf35b50346101ee5760203660031901126101ee576040612d7b612d76613c1e565b613fc3565b60208251918051835201516020820152f35b50346101ee5761ffff6040612da136613c35565b949160018060a01b0319168252602e6020528282209082526020522091165f52602052602060405f2054604051908152f35b50346101ee57806003193601126101ee576020612dee613f6a565b6040516001600160401b039091168152f35b50346112ef5760e03660031901126112ef57612e1a613c1e565b602435612e25613bfc565b6001600160a01b03199092165f81815260236020526040902080546001600160a01b0316936064359160c4359160a435916084359188156131845760ff6004830154166006811015613170576003036131615761ffff169788158015613156575b61314757612e9483876146c1565b612e9e85856146c1565b876130b857612eaf60028301613f08565b908151151590816130ad575b5061309e5760408101516001600160401b0316801515908161308b575b506130695760608101516001600160401b03168015159081613078575b506130695760808101516001600160401b03168015159081613056575b506130345760a08101516001600160401b03168015159081613043575b50613034576020015161ffff16801515908161301e575b5061300f575b858952602e60205260408920878a5260205260408920885f5260205260405f2054613000579181612fe1600787958c6040612fad8a8860a09c9a7f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc9e614719565b918c8152602e6020528d828220908252602052208d5f5260205260405f20550161ffff600181835460301c16011690613ee9565b604051933385526020850152604084015260608301526080820152a480f35b6316feb18560e11b8952600489fd5b63464e67af60e01b8952600489fd5b905061ffff600783015460301c1610155f612f46565b630410ff2960e31b8a5260048afd5b426001600160401b03161190505f612f2f565b436001600160401b03161190505f612f12565b633deac39560e01b8a5260048afd5b426001600160401b03161090505f612ef5565b436001600160401b03161090505f612ed8565b6330cd747160e01b8a5260048afd5b90503314155f612ebb565b5060015460ff8160a01c1615613138576001600160a01b0316803b156112ef575f60849160405192838092633a54cd5d60e11b82528b60048301528c60248301528d60448301523360648301525afa801561312d57613118575b50612f4c565b6131259199505f90613dc5565b5f975f613112565b6040513d5f823e3d90fd5b63023b34fb60e11b5f5260045ffd5b634c4d29cd60e11b5f5260045ffd5b506101008911612e86565b63268dbf6760e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b346112ef575f3660031901126112ef57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346112ef5760203660031901126112ef576001600160a01b03196131f5613c1e565b165f81815260236020526040902080546001600160a01b031680156131845733036132955760040160ff815416600681101590816131705760038114918215613288575b8215613277575b505061316157805460ff191660041790557f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b909150613170576004148380613240565b506005811491505f613239565b6282b42960e81b5f5260045ffd5b346112ef576101403660031901126112ef5760043561ffff81168091036112ef576132cc613beb565b906132d5613bfc565b6132dd613c0d565b9160c03660831901126112ef576001600160401b036132fa613f6a565b1643106131615780158015613b10575b8015613b03575b8015613af5575b8015613ae9575b8015613ad8575b8015613acb575b8015613abc575b6139a1576001600160401b03613348613de8565b16151580613aa3575b80613a77575b8015613a12575b80156139f9575b6139ea57604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561312d575f916139b0575b506001600160401b03169283156139a15761ffff8091169416926133da8486613e56565b61271081106139745750505f19905b5f546001600160401b03808216908114612bb2576001600160401b03199091166001919091016001600160401b039081169182175f556022547f0000000000000000000000000000000000000000000000000000000000000000604090811b63ffffffff60401b1690931760a081901b6001600160a01b0319169793909261347d92603f8116929116101561393a57613e69565b81546001600160601b0360039290921b91821b191692901b919091179055602280546001600160401b031981166001600160401b03918216600190810183169190911790925543811691820197919591908811612bb2576134fe7f000000000000000000000000000000000000000000000000000000000000000087613e97565b9061ffff61352c7f000000000000000000000000000000000000000000000000000000000000000089613e97565b936135577f00000000000000000000000000000000000000000000000000000000000000008a613e97565b956040519761356589613d22565b885260208801521660408087019190915260608601919091526001600160401b03918216608086015291811660a085015291821660c08401525f549051911695906135af81613d3d565b33815260208101928352604051906135c682613d59565b60843580151581036112ef57825260a43561ffff811681036112ef57602083015260c4356001600160401b03811681036112ef57604083015260e4356001600160401b03811681036112ef576060830152610104356001600160401b03811681036112ef576080830152610124356001600160401b03811681036112ef5760a08301526040810191825286606082016001815260808301998a5260a083019088825260c084019560018060401b03169687875260e08501915f83526101008601938a85526101208701975f89526101408801905f82526101608901945f86526101808a01985f8a525f52602360205260405f209960018060a01b0390511660018060a01b03198b5416178a5560018a01905161ffff808251161661ffff198354161782556136fc61ffff60208301511683613eb0565b61370e61ffff60408301511683613eca565b61372061ffff60608301511683613ee9565b608081810151835460a08085015160c0958601516001600160401b03938416604095861b600160401b600160801b03161791861b600160801b600160c01b031691909117951b6001600160c01b031916949094179094559351805160028e0180546020840151978401516060850151958501516001600160d81b031990921660ff941515949094169390931760089890981b62ffff00169790971760189290921b6301000000600160581b03169190911760589390931b600160581b600160981b03169290921760989590951b600160981b600160d81b0316949094179055919091015160038a0180546001600160401b0319169190921617905591519860048801929060068b10156131705760209f600761ffff9a6138c8998c996138d09f988b98899788976138c09c60ff801986541691161784555191835490600160481b600160881b03905160481b1691600160881b600160c81b03905160881b1692610100600160481b039060081b1690610100600160c81b0319161717179055516005830155516006820155019a51161682198a5416178955511687613eb0565b511684613eca565b511690613ee9565b5f8054600160401b600160801b03191643604090811b600160401b600160801b031691909117909155805193845285840191909152820152339082907f1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d90606090a3604051908152f35b61394381613e69565b905460039190911b1c60a01b6001600160a01b03198116613965575b50613e69565b61396e906141c6565b8961395f565b805f5160206148825f395f51905f5202905f5160206148825f395f51905f52820403612bb25704906133e9565b63d06b96b160e01b5f5260045ffd5b90506020813d6020116139e2575b816139cb60209383613dc5565b810103126112ef576139dc90613e42565b856133b6565b3d91506139be565b63148b7e9360e31b5f5260045ffd5b5060a43561ffff81168091036112ef5761010010613365565b506001600160401b03613a23613e15565b16151580613a5e575b801561335e5750613a3b613e2b565b6001600160401b03613a4b613e15565b166001600160401b03909116111561335e565b506001600160401b03613a6f613e2b565b161515613a2c565b50613a80613dfe565b6001600160401b03613a90613de8565b166001600160401b039091161115613357565b506001600160401b03613ab4613dfe565b161515613351565b5061271061ffff841610613334565b508061ffff83161061332d565b5061ffff841661ffff831611613326565b5061ffff82161561331f565b50602061ffff851611613318565b5061ffff84168111613311565b5061ffff84161561330a565b346112ef575f3660031901126112ef5760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561312d575f90611b7c57602090604051908152f35b346112ef575f3660031901126112ef576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346112ef575f3660031901126112ef575f5460401c6001600160401b03168152602090f35b6024359061ffff821682036112ef57565b6044359061ffff821682036112ef57565b6064359061ffff821682036112ef57565b600435906001600160a01b0319821682036112ef57565b60609060031901126112ef576004356001600160a01b0319811681036112ef57906024359060443561ffff811681036112ef5790565b80511515825260208082015161ffff16908301526040808201516001600160401b039081169184019190915260608083015182169084015260808083015182169084015260a09182015116910152565b9181601f840112156112ef578235916001600160401b0383116112ef57602083818601950101116112ef57565b346112ef575f3660031901126112ef5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60e081019081106001600160401b0382111761139c57604052565b6101a081019081106001600160401b0382111761139c57604052565b60c081019081106001600160401b0382111761139c57604052565b604081019081106001600160401b0382111761139c57604052565b606081019081106001600160401b0382111761139c57604052565b608081019081106001600160401b0382111761139c57604052565b601f909101601f19168101906001600160401b0382119082101761139c57604052565b60c4356001600160401b03811681036112ef5790565b610104356001600160401b03811681036112ef5790565b60e4356001600160401b03811681036112ef5790565b610124356001600160401b03811681036112ef5790565b51906001600160401b03821682036112ef57565b81810292918115918404141715612bb257565b906040821015613e8357600c600183811c60020193160290565b634e487b7160e01b5f52603260045260245ffd5b6001600160401b039182169082160191908211612bb257565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b90604051613f1581613d59565b825460ff81161515825261ffff600882901c1660208301526001600160401b03601882901c81166040840152605882901c8116606084015260989190911c8116608083015260019093015490921660a0830152565b5f5460401c6001600160401b03168015613fb557613fb2907f00000000000000000000000000000000000000000000000000000000000000006001600160401b031690613e97565b90565b50436001600160401b031690565b604051613fcf81613d74565b5f81525f60208201525060018060a01b0319165f52602d60205260405f2060018101548015614010576040519161400583613d74565b548252602082015290565b505060405161401e81613d74565b5f81526001602082015290565b6040519061403882613d59565b5f60a0838281528260208201528260408201528260608201528260808201520152565b908060209392818452848401375f828201840152601f01601f1916010190565b929061409490613fb2959360408652604086019161405b565b92602081850391015261405b565b61ffff5f199116019061ffff8211612bb257565b8054821015613e83575f5260205f2001905f90565b61ffff1661ffff8114612bb25760010190565b909392938483116112ef5784116112ef578101920390565b9192916001600160401b03821161139c576040519161411f601f8201601f191660200184613dc5565b8294818452818301116112ef578281602093845f960137010152565b908160a09103126112ef576040519060a082016001600160401b0381118382101761139c5760405280516001600160a01b03811681036112ef5782526020810151602083015260408101516040830152606081015160038110156112ef5760608301526141aa90608001613e42565b608082015290565b61ffff60019116019061ffff8211612bb257565b6001600160a01b0319165f818152602360205260409020546001600160a01b0316156146be57805f52602560205260405f20805460609060015460ff8160a01c166145e3575b5081519260018401809411612bb2575f5b82811061453d575050505f5b8281106142eb57505050805f52602560205260405f208054905f8155816142cd575b5050805f52602c6020525f6040812055805f52602d6020525f6001604082208281550155805f5260236020525f60076040822082815582600182015582600282015582600382015582600482015582600582015582600682015501557f457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c9965f80a2565b5f5260205f20908101905b8181101561424b575f81556001016142d8565b8061451b575f5b60015b61010061ffff8216111561430d575050600101614229565b855f52602860205260405f20825f5260205260405f2061ffff82165f5260205260405f205461445a575b6143c790865f52602960205260405f20835f5260205260405f2061ffff82165f5260205261ffff60405f205416614428575b865f52602a60205260405f20835f5260205260405f2061ffff82165f5260205260ff60405f205460101c166143f7575b865f52602e60205260405f20835f5260205260405f2061ffff82165f5260205260405f20546143cc576140cb565b6142f5565b865f52602e60205260405f20835f5260205260405f2061ffff82165f526020525f60408120556140cb565b865f52602a60205260405f20835f5260205260405f2061ffff82165f526020525f6001604082208281550155614399565b865f52602960205260405f20835f5260205260405f2061ffff82165f5260205260405f2061ffff198154169055614369565b92855f979592969752602860205260405f20865f5260205260405f2061ffff85165f5260205260405f20549460015b61ffff8116602081116144e95790816001808a6144ae951c16146144b3575b506140cb565b614489565b8a5f52602760205260405f208a5f5260205260405f2061ffff89165f5260205260405f20905f526020525f60408120555f6144a8565b50505f8881526028602090815260408083208a8452825280832061ffff89168452909152812055969591945092614337565b5f198101818111612bb2578251811015613e835760051b8201602001516142f2565b61454781836140b6565b90545f88815260246020908152604080832060039590951b9390931c6001600160a01b0316808352938152828220805460ff19169055898252602b90522060018301929190838311612bb25761ffff8060019516165f526020525f6040812055875f52602660205260405f2090838060a01b03165f526020525f600460408220828155828582015582600282015582600382015501550161421d565b6040516376bc6b8f60e11b8152600481018690529192505f90829060249082906001600160a01b03165afa90811561312d575f91614624575b50905f61420c565b90503d805f833e6146358183613dc5565b8101906020818303126112ef578051906001600160401b0382116112ef57019080601f830112156112ef578151916001600160401b03831161139c578260051b90604051936146876020840186613dc5565b84526020808501928201019283116112ef57602001905b8282106146ae575050505f61461c565b815181526020918201910161469e565b50565b905f5160206148625f395f51905f52821080614703575b15613147578115806146f9575b613147576146f29161479c565b1561314757565b50600181146146e5565b505f5160206148625f395f51905f5281106146d8565b91608093916040519384526020840152604083015260608201522090565b9291905f5160206148625f395f51905f525f940691829060051b8201915b8281106147625750505050565b909192945f5160206148625f395f51905f5283816020938186358b099008970993929101614755565b906040811015613e835760051b0190565b5f5160206148625f395f51905f52811080159061484a575b614844575f5160206148625f395f51905f5281819209918009905f5160206148625f395f51905f5281065f5160206148625f395f51905f5203915f5160206148625f395f51905f528311612bb2575f5160206148625f395f51905f52808080848196097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09600108935f08081490565b50505f90565b505f5160206148625f395f51905f528210156147b456fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000100068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43a264697066735822122011f22e1748a49d73ed482022960d673c26c636498caa22beaaa66488dfe6ca5064736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _finalizeVerifier common.Address, _decryptCombineVerifier common.Address, _epochDurationBlocks *big.Int) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _finalizeVerifier, _decryptCombineVerifier, _epochDurationBlocks)
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

// DKGManagerEpochEvictedIterator is returned from FilterEpochEvicted and is used to iterate over the raw logs and unpacked data for EpochEvicted events raised by the DKGManager contract.
type DKGManagerEpochEvictedIterator struct {
	Event *DKGManagerEpochEvicted // Event containing the contract specifics and raw log

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
func (it *DKGManagerEpochEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerEpochEvicted)
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
		it.Event = new(DKGManagerEpochEvicted)
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
func (it *DKGManagerEpochEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerEpochEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerEpochEvicted represents a EpochEvicted event raised by the DKGManager contract.
type DKGManagerEpochEvicted struct {
	EpochId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEpochEvicted is a free log retrieval operation binding the contract event 0x457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c996.
//
// Solidity: event EpochEvicted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) FilterEpochEvicted(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerEpochEvictedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "EpochEvicted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerEpochEvictedIterator{contract: _DKGManager.contract, event: "EpochEvicted", logs: logs, sub: sub}, nil
}

// WatchEpochEvicted is a free log subscription operation binding the contract event 0x457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c996.
//
// Solidity: event EpochEvicted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) WatchEpochEvicted(opts *bind.WatchOpts, sink chan<- *DKGManagerEpochEvicted, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "EpochEvicted", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerEpochEvicted)
				if err := _DKGManager.contract.UnpackLog(event, "EpochEvicted", log); err != nil {
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

// ParseEpochEvicted is a log parse operation binding the contract event 0x457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c996.
//
// Solidity: event EpochEvicted(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) ParseEpochEvicted(log types.Log) (*DKGManagerEpochEvicted, error) {
	event := new(DKGManagerEpochEvicted)
	if err := _DKGManager.contract.UnpackLog(event, "EpochEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerEpochFinalizedIterator is returned from FilterEpochFinalized and is used to iterate over the raw logs and unpacked data for EpochFinalized events raised by the DKGManager contract.
type DKGManagerEpochFinalizedIterator struct {
	Event *DKGManagerEpochFinalized // Event containing the contract specifics and raw log

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
func (it *DKGManagerEpochFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerEpochFinalized)
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
		it.Event = new(DKGManagerEpochFinalized)
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
func (it *DKGManagerEpochFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerEpochFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerEpochFinalized represents a EpochFinalized event raised by the DKGManager contract.
type DKGManagerEpochFinalized struct {
	EpochId                  [12]byte
	AggregateCommitmentsHash [32]byte
	CollectivePublicKeyHash  [32]byte
	ShareCommitmentHash      [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterEpochFinalized is a free log retrieval operation binding the contract event 0x4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a472.
//
// Solidity: event EpochFinalized(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) FilterEpochFinalized(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerEpochFinalizedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "EpochFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerEpochFinalizedIterator{contract: _DKGManager.contract, event: "EpochFinalized", logs: logs, sub: sub}, nil
}

// WatchEpochFinalized is a free log subscription operation binding the contract event 0x4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a472.
//
// Solidity: event EpochFinalized(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) WatchEpochFinalized(opts *bind.WatchOpts, sink chan<- *DKGManagerEpochFinalized, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "EpochFinalized", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerEpochFinalized)
				if err := _DKGManager.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// ParseEpochFinalized is a log parse operation binding the contract event 0x4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a472.
//
// Solidity: event EpochFinalized(bytes12 indexed epochId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) ParseEpochFinalized(log types.Log) (*DKGManagerEpochFinalized, error) {
	event := new(DKGManagerEpochFinalized)
	if err := _DKGManager.contract.UnpackLog(event, "EpochFinalized", log); err != nil {
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

// DKGManagerRegistrationClosedIterator is returned from FilterRegistrationClosed and is used to iterate over the raw logs and unpacked data for RegistrationClosed events raised by the DKGManager contract.
type DKGManagerRegistrationClosedIterator struct {
	Event *DKGManagerRegistrationClosed // Event containing the contract specifics and raw log

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
func (it *DKGManagerRegistrationClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRegistrationClosed)
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
		it.Event = new(DKGManagerRegistrationClosed)
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
func (it *DKGManagerRegistrationClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRegistrationClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRegistrationClosed represents a RegistrationClosed event raised by the DKGManager contract.
type DKGManagerRegistrationClosed struct {
	EpochId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistrationClosed is a free log retrieval operation binding the contract event 0xca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f1358.
//
// Solidity: event RegistrationClosed(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) FilterRegistrationClosed(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerRegistrationClosedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RegistrationClosed", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRegistrationClosedIterator{contract: _DKGManager.contract, event: "RegistrationClosed", logs: logs, sub: sub}, nil
}

// WatchRegistrationClosed is a free log subscription operation binding the contract event 0xca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f1358.
//
// Solidity: event RegistrationClosed(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) WatchRegistrationClosed(opts *bind.WatchOpts, sink chan<- *DKGManagerRegistrationClosed, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RegistrationClosed", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRegistrationClosed)
				if err := _DKGManager.contract.UnpackLog(event, "RegistrationClosed", log); err != nil {
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

// ParseRegistrationClosed is a log parse operation binding the contract event 0xca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f1358.
//
// Solidity: event RegistrationClosed(bytes12 indexed epochId)
func (_DKGManager *DKGManagerFilterer) ParseRegistrationClosed(log types.Log) (*DKGManagerRegistrationClosed, error) {
	event := new(DKGManagerRegistrationClosed)
	if err := _DKGManager.contract.UnpackLog(event, "RegistrationClosed", log); err != nil {
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
