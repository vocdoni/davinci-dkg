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
	SeedDelay                 uint16
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochEvicted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochFinalized\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x61018080604052346102235760c0816153c580380380916100208285610227565b8339810103126102235780519063ffffffff821691828103610223576100486020830161025e565b6100546040840161025e565b906100616060850161025e565b9261007a60a06100736080880161025e565b960161025e565b9563ffffffff461603610214576001600160a01b03821615610205576001600160a01b0383161580156101f4575b80156101e3575b80156101d2575b6101c35763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b6024820152601881526100f2603882610227565b5190201660c05260e05261010052610120526101405233610160526040516151529081610273823960805181611047015260a05181818161026b01528181611987015281816125a801528181613a590152614b4c015260c0518181816106080152612651015260e0518181816102bf01528181610dcc01526139230152610100518181816110900152818161169c01526134be015261012051818181610e20015281816110e80152612da3015261014051818181610f4d01528181611c5f015261310501526101605181610e930152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100b6565b506001600160a01b038516156100af565b506001600160a01b038416156100a8565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761024a57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102235756fe60806040526004361015610011575f80fd5b5f3560e01c806306433b1b14610224578063074a75e11461021f5780630b1451f01461021a57806318287e5f1461021557806323488be5146102105780632648f5671461020b5780632c7c7642146102065780632de546d5146102015780633353ec6e146101fc5780634554c0be146101f757806349c61a12146101f2578063510ba2df146101ed5780635a8f2bb3146101e857806363f314cd146101e3578063669a76a9146101de5780636d16897d146101d957806372517b4b146101d457806377235ee1146101cf57806385e1f4d0146101ca5780638dc1f53a146101c557806393c3d3a8146101c05780639bbada67146101bb578063a305e0f3146101b6578063a4adcd7f146101b1578063b7bca615146101ac578063be59b8ea146101a7578063bea5210d146101a2578063bf1922091461019d578063ca3c045814610198578063d3720aac14610193578063d99337671461018e578063ebe86c13146101895763fe1604b514610184575f80fd5b611c4a565b611c21565b611896565b611795565b61170d565b611687565b611561565b611496565b61127c565b611255565b6111c1565b611117565b6110d3565b61106b565b61102b565b610f90565b610f28565b610e74565b610dfb565b610db7565b610d77565b610d31565b610c9b565b610c17565b610b85565b610b44565b610a6d565b610654565b6105ec565b610507565b61034e565b61029a565b610256565b5f91031261023357565b5f80fd5b6001600160a01b031690565b6001600160a01b03909116815260200190565b34610233575f366003190112610233576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610233575f3660031901126102335760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610332575f91610303575b50604051908152602090f35b610325915060203d60201161032b575b61031d8183611d4b565b810190611d6e565b5f6102f7565b503d610313565b611d7d565b600435906001600160a01b03198216820361023357565b3461023357602036600319011261023357610367610337565b61037081611d88565b80546001600160a01b039061038490610237565b16156104f85760058101805492600160ff85166103a0816113aa565b036104e957600883015461ffff1691600184019384549361ffff6103d26103cb8760101c61ffff1690565b61ffff1690565b9116146104e957605084901c6001600160401b0316956103f187610a31565b4311156104e9576104669061046190610451906104369060481c6001600160401b03169961043061042960408b901c61ffff166103cb565b809c611ed4565b90611ed4565b436001600160401b03169961044b908b611ef2565b99611ef2565b9560901c6001600160401b031690565b610a31565b6001600160401b03851610156104da577f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c946104ac876104b1945f600689960155611f0b565b611f34565b604080516001600160401b0395861681529290941660208301526001600160a01b0319169290a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b3461023357602036600319011261023357610520610337565b61052981611d88565b6105338154610237565b6001600160a01b038116156104f85761054b90610237565b33036105de57600501805460ff16610562816113aa565b600381149081156105c9575b81156105b5575b506104e957805460ff191660041790556001600160a01b0319167f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b600491506105c2816113aa565b145f610575565b90506105d4816113aa565b600581149061056e565b6282b42960e81b5f5260045ffd5b34610233575f36600319011261023357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b61ffff81160361023357565b604435906106458261062c565b565b606435906106458261062c565b346102335760e03660031901126102335761066d610337565b60243560443561067c8161062c565b6064359060843560a4359460c4359261069482611d88565b9361069f8554610237565b976001600160a01b038916156104f85760036106bf600588015460ff1690565b6106c8816113aa565b036104e95761ffff83169889158015610a26575b610a17576106ea8689613e3b565b6106f48383613e3b565b886109615761070560038801611fc2565b906107108251151590565b908161094e575b5061093f576107296040820151610a31565b6001600160401b03811615159081610924575b50610902576107516104616060830151610a31565b8015159081610911575b50610902576107706104616080830151610a31565b80151590816108ef575b506108cd5761078f61046160a0830151610a31565b80151590816108dc575b506108cd57602001516107af9061ffff166103cb565b80151590816108ac575b5061089d575b6107e2836107dd8a6107d088611da2565b905f5260205260405f2090565b61203c565b5461088e5761084f60087f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc978a610830610889976107dd610826898f8e8b91613ea2565b936107d08c611da2565b550180546108499060301c61ffff1660010161ffff1690565b9061204d565b604080513381526020810198909852870194909452606086019390935260808501929092526001600160a01b0319169290819060a0820190565b0390a4005b6316feb18560e11b5f5260045ffd5b63464e67af60e01b5f5260045ffd5b905061ffff6108c4600889015461ffff9060301c1690565b1610155f6107b9565b630410ff2960e31b5f5260045ffd5b426001600160401b03161190505f610799565b436001600160401b03161190505f61077a565b633deac39560e01b5f5260045ffd5b426001600160401b03161090505f61075b565b61092e9150610a31565b436001600160401b0316105f61073c565b6330cd747160e01b5f5260045ffd5b6109589150610237565b3314155f610717565b505f5460ff60e082901c1615610976565b1590565b610a0857610996906109919060401c6001600160a01b031681565b610237565b803b1561023357604051633a54cd5d60e11b81526001600160a01b031986166004820152602481018a905261ffff85166044820152336064820152905f90829060849082905afa8015610332576109ee575b506107bf565b806109fc5f610a0293611d4b565b80610229565b5f6109e8565b63023b34fb60e11b5f5260045ffd5b634c4d29cd60e11b5f5260045ffd5b506101008a116106dc565b6001600160401b031690565b6001600160401b0381160361023357565b359061064582610a3d565b6001600160a01b0319909116815260200190565b34610233576101c036600319011261023357600435610a8b8161062c565b602435610a978161062c565b60443591610aa48361062c565b606435610ab08161062c565b608435610abc8161062c565b60a435610ac881610a3d565b60c43591610ad583610a3d565b60e43593610ae285610a3d565b60c03661010319011261023357610b0c97610b0097610104976124f4565b60405191829182610a59565b0390f35b6060906003190112610233576004356001600160a01b031981168103610233579060243590604435610b418161062c565b90565b34610233576020610b7c610b5736610b10565b6001600160a01b03199092165f908152602d855260408082209282529185522061203c565b54604051908152f35b34610233576020366003190112610233576040610ba8610ba3610337565b612aef565b60208251918051835201516020820152f35b6001600160401b03169052565b80511515825260208082015161ffff16908301526040808201516001600160401b039081169184019190915260608083015182169084015260808083015182169084015260a09182015116910152565b3461023357602036600319011261023357610c30610337565b610c38612b57565b506001600160a01b0319165f90815260226020526040902060c090610c5f90600301611fc2565b610c6c6040518092610bc7565bf35b9181601f84011215610233578235916001600160401b038311610233576020838186019501011161023357565b346102335760e036600319011261023357610cb4610337565b602435604435916064356084356001600160401b03811161023357610cdd903690600401610c6e565b90919060a4356001600160401b03811161023357610cff903690600401610c6e565b60c4359790949193906001600160401b03891161023357610d27610d2f993690600401610c6e565b989097612c49565b005b34610233576040366003190112610233576020610b7c610d4f610337565b60243590610d5c8261062c565b6001600160a01b0319165f908152602a84526040902061203c565b346102335760206001610dad610d8c36610b10565b91848060a01b0319165f526029855260405f20905f52845260405f2061203c565b0154604051908152f35b34610233575f366003190112610233576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610233575f3660031901126102335760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610332575f916103035750604051908152602090f35b6001600160a01b0381160361023357565b3461023357602036600319011261023357600435610e9181610e63565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036105de575f549060ff8260e01c16610f19576001600160a01b03811615610f0a57600160401b600160e81b031990911660409190911b600160401b600160e01b031617600160e01b175f55005b63e6c4247b60e01b5f5260045ffd5b6373253a9760e01b5f5260045ffd5b34610233575f3660031901126102335760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610332575f916103035750604051908152602090f35b346102335761010036600319011261023357610faa610337565b602435610fb5610638565b9160843560643560a4356001600160401b03811161023357610fdb903690600401610c6e565b9060c4356001600160401b03811161023357610ffb903690600401610c6e565b60e4359890959194906001600160401b038a1161023357611023610d2f9a3690600401610c6e565b99909861300f565b34610233575f36600319011261023357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610233575f3660031901126102335760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610332575f916103035750604051908152602090f35b34610233575f366003190112610233576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461023357610b0c61116861112b36610b10565b915f6040805161113a81611ca2565b828152826020820152015260018060a01b0319165f52602960205260405f20905f5260205260405f2061203c565b60016040519161117783611ca2565b60ff815461ffff8116855260101c16151560208401520154604082015260405191829182919091604080606083019461ffff81511684526020810151151560208501520151910152565b3461023357610160366003190112610233576111db610337565b602435604435916111eb8361062c565b6111f3610647565b608435906101043560e43560c43560a435610124356001600160401b03811161023357611224903690600401610c6e565b610144359a90979196906001600160401b038c116102335761124d610d2f9c3690600401610c6e565b9b909a6132e6565b34610233575f366003190112610233575f546040516001600160401b039091168152602090f35b346102335760e036600319011261023357611295610337565b6024356112a18161062c565b604435916064356084356001600160401b038111610233576112c7903690600401610c6e565b90919060a4356001600160401b038111610233576112e9903690600401610c6e565b60c4359790949193906001600160401b03891161023357611311610d2f993690600401610c6e565b9890976137d7565b9060e0806106459361ffff815116845261ffff602082015116602085015261ffff604082015116604085015261135a6060820151606086019061ffff169052565b60808181015161ffff169085015261137a60a082015160a0860190610bba565b61138c60c082015160c0860190610bba565b0151910190610bba565b634e487b7160e01b5f52602160045260245ffd5b600611156113b457565b611396565b9060068210156113b45752565b81516001600160a01b031681526103008101929161064591906102e090610160906113f960208201516020860190611319565b61140c6040820151610120860190610bc7565b61141f60608201516101e08601906113b9565b6114326080820151610200860190610bba565b61144560a0820151610220860190610bba565b60c081015161024085015260e081015161026085015261010081015161ffff1661028085015261012081015161ffff166102a085015261014081015161ffff166102c0850152015161ffff16910152565b3461023357602036600319011261023357610b0c6115556115506114b8610337565b5f6101606040516114c881611cc2565b8281526040516114d781611cde565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e08201526020820152611511612b57565b60408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201520152611d88565b613c65565b604051918291826113c6565b3461023357608036600319011261023357610b0c61157d610337565b6116386107dd6024359261162a6116086044356115998161062c565b6107dd6001806115e9846115df8c6064359b8c916115b68361062c565b5f60606040516115c581611cfa565b82815282602082015282604082015201526107d08d611dbc565b549061ffff161c90565b161495861561167d576107dd83995b8815611673576107d08398611dd6565b549161161f611615611f84565b61ffff9097168752565b61ffff166020860152565b604084015215156060830152565b60405191829182919091606080608083019461ffff815116845261ffff60208201511660208501526040810151604085015201511515910152565b6107d05f98611dd6565b6107dd5f996115f8565b34610233575f366003190112610233576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b60206040818301928281528451809452019201905f5b8181106116ee5750505090565b82516001600160a01b03168452602093840193909201916001016116e1565b34610233576020366003190112610233576001600160a01b031961172f610337565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b81811061177657610b0c8561176a81870382611d4b565b604051918291826116cb565b82546001600160a01b0316845260209093019260019283019201611753565b3461023357604036600319011261023357610b0c6117e66117b4610337565b602435906117c182610e63565b6117c9612b57565b506001600160a01b0319165f908152602560205260409020613233565b61184461183b6004604051936117fb85611d15565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b151560a0830152565b6040519182918291909160a08060c0830194600180831b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b34610233576020366003190112610233576118af610337565b6118b881611d88565b9060018060a01b036118ca8354610237565b16156104f857600582018054600184018054949093916118fd6109726001600160401b03605089901c1660ff8416614a78565b6104e957600882019561ffff6119266103cb61191b8a5461ffff1690565b9360101c61ffff1690565b91161015611bff5761194b61194461193d85611df0565b3390613233565b5460ff1690565b611bf057600682018054918215611b68575b50506040516313a4120960e31b815260a0818061197d3360048301610243565b0381600180851b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610332576001916060915f91611b39575b5001516119c981613ddc565b6119d281613ddc565b03611b2a5760408051602081019283523360601b6001600160601b0319169181019190915260079190611a1281605481015b03601f198101835282611d4b565b5190209101541115611b1b57611a71611a2d855461ffff1690565b94611a4033611a3b85611e0a565b613de6565b611a62611a5533611a5086611df0565b613233565b805460ff19166001179055565b611a6b86613e27565b906121c9565b60405161ffff851681526001600160a01b0319821694611acb91611abf9190339088907f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b90602090a3613e27565b935460101c61ffff1690565b9261ffff808516911614611adb57005b611af592611ae891614b2d565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b611b5b915060a03d60a011611b61575b611b538183611d4b565b810190613d69565b5f6119bd565b503d611b49565b909150611b809060481c6001600160401b0316610a31565b80431115611be15740908115611bd2578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f8061195d565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b60039190911b1c6001600160a01b031690565b34610233575f366003190112610233575f546040805191901c6001600160a01b03168152602090f35b34610233575f366003190112610233576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b634e487b7160e01b5f52604160045260245ffd5b606081019081106001600160401b03821117611cbd57604052565b611c8e565b61018081019081106001600160401b03821117611cbd57604052565b61010081019081106001600160401b03821117611cbd57604052565b608081019081106001600160401b03821117611cbd57604052565b60c081019081106001600160401b03821117611cbd57604052565b604081019081106001600160401b03821117611cbd57604052565b601f909101601f19168101906001600160401b03821190821017611cbd57604052565b90816020910312610233575190565b6040513d5f823e3d90fd5b6001600160a01b0319165f90815260226020526040902090565b6001600160a01b0319165f908152602d6020526040902090565b6001600160a01b0319165f90815260276020526040902090565b6001600160a01b0319165f90815260266020526040902090565b6001600160a01b0319165f90815260236020526040902090565b6001600160a01b0319165f90815260246020526040902090565b6001600160a01b0319165f908152602c6020526040902090565b6001600160a01b0319165f908152602a6020526040902090565b6001600160a01b0319165f90815260286020526040902090565b6001600160a01b0319165f90815260296020526040902090565b6001600160a01b0319165f90815260256020526040902090565b6001600160a01b0319165f908152602b6020526040902090565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b039182169082160391908211611eed57565b611ec0565b6001600160401b039182169082160191908211611eed57565b8054600160481b600160881b03191660489290921b600160481b600160881b0316919091179055565b8054600160501b600160901b03191660509290921b600160501b600160901b0316919091179055565b9060068110156113b45760ff80198354169116179055565b6001600160401b039091169052565b60405190610645608083611d4b565b6040519061064561010083611d4b565b6040519061064561018083611d4b565b60405190610645604083611d4b565b90610645604051611fd281611d15565b835460ff81161515825261ffff600882901c1660208301526001600160401b03601882901c81166040840152605882901c811660608401529194859260a09261203592600192909161202e9160989190911c1660808701611f75565b0154610a31565b9101611f75565b9061ffff165f5260205260405f2090565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b35610b4181610a3d565b35610b418161062c565b908160209103126102335751610b4181610a3d565b908160051b9180830460201490151715611eed57565b908160011b9180830460021490151715611eed57565b81810292918115918404141715611eed57565b81156120de570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160401b03908116908114611eed5760010190565b634e487b7160e01b5f52603260045260245ffd5b90604082101561213757600c600183811c810193160290565b61210a565b91908260c09103126102335760405161215481611d15565b809280359081151582036102335760a06121b89181938552602081013561217a8161062c565b602086015261218b60408201610a4e565b604086015261219c60608201610a4e565b60608601526121ad60808201610a4e565b608086015201610a4e565b910152565b60068210156113b45752565b9061ffff1661ffff19825416179055565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b60016122d560e06106459461222d61ffff825116866121c9565b61223f61ffff602083015116866121da565b61225161ffff604083015116866121f4565b61226361ffff6060830151168661204d565b6080810151855461ffff60401b191660409190911b61ffff60401b1617855561229861229260a0830151610a31565b86611f34565b6122ce6122a860c0830151610a31565b8654600160901b600160d01b03191660909190911b600160901b600160d01b0316178655565b0151610a31565b910180546001600160401b0319166001600160401b03909216919091179055565b60016122d560a0610645948051151560ff801987541691151516178555602081015162ffff0086549160081b169062ffff00191617855561236361233d6040830151610a31565b86546301000000600160581b03191660189190911b6301000000600160581b0316178655565b6123996123736060830151610a31565b8654600160581b600160981b03191660589190911b600160581b600160981b0316178655565b6122ce6123a96080830151610a31565b8654600160981b600160d81b03191660989190911b600160981b600160d81b0316178655565b906108496101606008610645946124056123e98651610237565b82546001600160a01b0319166001600160a01b03909116178255565b612416602086015160018301612213565b6124276040860151600383016122f6565b61248a60058201612445606088015161243f816113aa565b82611f5d565b6124776124556080890151610a31565b8254610100600160481b03191660089190911b610100600160481b0316178255565b61248460a0880151610a31565b90611f0b565b60c0850151600682015560e0850151600782015501926124b96124b361010083015161ffff1690565b856121c9565b6124d26124cc61012083015161ffff1690565b856121da565b6124eb6124e561014083015161ffff1690565b856121f4565b015161ffff1690565b91929597949661ffff83168015908115612ae2575b8115612ad4575b8115612ac5575b8115612ab8575b8115612aa5575b8115612a97575b508015612a88575b8015612a7c575b8015612a6d575b8015612a41575b8015612a2b575b8015612a15575b6104da57604086016001600160401b036125708261206c565b16151590816129fc575b816129d5575b508015612969575b801561294f575b61294057604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610332575f91612911575b506001600160401b031680156104da5761260261ffff841661ffff8a166120c1565b9061271082106128da5750505f19975b5f5461261d90610a31565b612626906120f2565b5f80546001600160401b0319166001600160401b039092169190911790555f5461264f90610a31565b7f000000000000000000000000000000000000000000000000000000000000000060401b63ffffffff60401b166001600160401b039091161760a01b6001600160a01b031916998a916021546126a490610a31565b906126ae82610a31565b603f1691600160401b60019003166040111561284d9b61277a6128489a61276f6127aa9a61270d6127a19a6126ed6127989a6127859961289c5761211e565b90919082549060031b9160a01c821b9160018060601b03901b1916179055565b61274f61272d61271e602154610a31565b6001016001600160401b031690565b602180546001600160401b0319166001600160401b0392909216919091179055565b61276461275a611f93565b61ffff909e168e52565b61ffff1660208d0152565b61ffff1660408b0152565b61ffff166060890152565b61ffff8916608088015260a08701611f75565b60c08501611f75565b60e08301611f75565b6128126127b75f54610a31565b61ffff9490941694436001600160401b031694612809906127f46127db8989611ef2565b936127e4611fa3565b338152966020880152369061213c565b60408601526001606086015260808501611f75565b60a08301611f75565b5f60c08201528560e08201525f6101008201525f6101208201525f6101408201525f61016082015261284387611d88565b6123cf565b611ef2565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917f2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d069190a390565b6128b56128a88261211e565b90549060031b1c60a01b90565b6001600160a01b031981166128cb575b5061211e565b6128d490613f95565b5f6128c5565b61290661290b927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa436120c1565b6120d4565b97612612565b612933915060203d602011612939575b61292b8183611d4b565b810190612080565b5f6125e0565b503d612921565b63148b7e9360e31b5f5260045ffd5b5061010061ffff61296260208901612076565b161161258f565b506060860161297a6104618261206c565b151590816129bc575b8161298f575b50612588565b90506129a96104616129a360a08a0161206c565b9261206c565b6001600160401b0390911611155f612989565b90506129cd61046160a0890161206c565b151590612983565b90506129e96104616129a360808a0161206c565b6001600160401b0390911611155f612580565b9050612a0d6104616080890161206c565b15159061257a565b506001600160401b038981169082161115612557565b506001600160401b03888116908a161115612550565b50612a5c61046161ffff8716436001600160401b0316611ef2565b6001600160401b0389161115612549565b5061010061ffff861611612542565b5061ffff85161561253b565b5061271061ffff881610612534565b905061ffff8516105f61252c565b905061ffff831661ffff86161190612525565b61ffff861615915061251e565b602061ffff8516119150612517565b61ffff841681119150612510565b61ffff8416159150612509565b604051612afb81611d30565b5f81525f60208201525060018060a01b0319165f52602c60205260405f2060018101548015612b3c5760405191612b3183611d30565b548252602082015290565b5050604051612b4a81611d30565b5f81526001602082015290565b60405190612b6482611d15565b5f60a0838281528260208201528260408201528260608201528260808201520152565b90610120828203126102335780601f830112156102335761012060405192612baf8285611d4b565b8391810192831161023357905b828210612bc95750505090565b8135815260209182019101612bbc565b908060209392818452848401375f828201840152601f01601f1916010190565b9290612c1290610b419593604086526040860191612bd9565b926020818503910152612bd9565b9060018201809211611eed57565b9060808201809211611eed57565b91908201809211611eed57565b95919690979294939998612c5c87611d88565b9160018060a01b03612c6e8454610237565b16156104f8576005830191612c84835460ff1690565b612c8d816113aa565b600381146130005780612ca16002926113aa565b036104e957612cb66104616002860154610a31565b43106104e9576008840194612cd0865460101c61ffff1690565b966001860154612ce86103cb8261ffff9060201c1690565b61ffff8a1610612ff1578d158015612fe9575b8015612fe1575b612fd257612d1284830183612b87565b988c612d2f612d238c519260a01c90565b6001600160601b031690565b1491821592612fb6575b8215612f93575b508115612f80575b508015612f72575b8015612f64575b8015612f56575b612f335760408051602081018f81529181018e9052606081018c9052612d959190612d8c8160808101611a04565b5190208c61432d565b938460e08a015103612f33577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610233575f93612df560405196879586948594635c73957b60e11b865260048601612bf9565b03915afa801561033257612f42575b50612e106108a0612095565b8603612f3357612e7c958c612e7194612e3993612e316103cb996101000190565b51918c61444b565b612e4d612e47610820612095565b8c612c3c565b60016020612e5a8a611e24565b83358155920135910155805460ff19166003179055565b5460101c61ffff1690565b612e90612e8a610860612095565b88612c3c565b965f5b828110612eeb57505060408051958652602086019490945250918301919091529192506001600160a01b0319909116907f4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a47290606090a2565b80612f0f60019260061b8b0160208101359035604091825191825260208201522090565b612f2c612f1b88611e3e565b61ffff8460051b870135169061203c565b5501612e93565b63d1fed5fd60e01b5f5260045ffd5b806109fc5f612f5093611d4b565b5f612e04565b508960c08901511415612d5e565b508b60a08901511415612d57565b508c60808901511415612d50565b60608a0151915061ffff1614155f612d48565b909150612fad6103cb60408c015b519260101c61ffff1690565b1415905f612d40565b915060208a0151612fca61ffff84166103cb565b141591612d39565b63c5f680ed60e01b5f5260045ffd5b508a15612d02565b508c15612cfb565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b949992979098999591939561302386611d88565b9760018060a01b036130358a54610237565b16156104f857600361304b60058b015460ff1690565b613054816113aa565b036104e95761ffff86169b8c158015613228575b8015613220575b61321157613084876107dd8e6107d08c611da2565b54948515613202578a61ffff6130c18f60016130b76130af8f8f906103cb956107d06107dd92611e58565b5461ffff1690565b94015461ffff1690565b9116106131f3576130d9886107dd8f6107d08d611e72565b956130e9875460ff9060101c1690565b6131e45761310291858f8f8f8f8f928f8f958b95614618565b937f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610233575f93613157604051998a9586948594635c73957b60e11b865260048601612bf9565b03915afa918215610332577f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae198889660019661319a956131d0575b508c896148dd565b805462ff000019166201000017815501556040805194855260208501929092526001600160a01b0319169290819081015b0390a4565b806109fc5f6131de93611d4b565b5f613192565b63955c0c4960e01b5f5260045ffd5b63032cddf960e11b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508a1561306f565b506101008d11613068565b9060018060a01b03165f5260205260405f2090565b61ffff5f199116019061ffff8211611eed57565b8054821015612137575f5260205f2001905f90565b90610200828203126102335780601f8301121561023357610200604051926132998285611d4b565b8391810192831161023357905b8282106132b35750505090565b81358152602091820191016132a6565b9091604092825260208201520190565b61ffff1661ffff8114611eed5760010190565b9b939a909a999497959891999692966132fe8d611d88565b9360018060a01b036133108654610237565b16156104f8576003613326600587015460ff1690565b61332f816113aa565b036104e9576133476109728f61193d61194491611df0565b6136fd5761ffff8c169a8b1580156136e0575b80156136d4575b80156136c5575b80156136bd575b6136af576133918f61338b8f9161338590611e0a565b91613248565b9061325c565b905433916001600160a01b03916133a89190611c0e565b1603612f33576133c18f8f6107dd906107d08e93611da2565b54918215613202576133d4918c85613ea2565b03612f33578b998d89600161ffff8f161b9c6133f08e93611dbc565b5f91825260205260409020906134059161203c565b54166136a0578b998d928f9261341d868a0187613271565b9c61342785611e3e565b906134319161203c565b548d51909460a01c6001600160601b03161494851595613691575b50841561367e575b841561366e575b841561365f575b508315613650575b508215613641575b508115613638575b811561360e575b50612f33576101208801946134998651996101400190565b98611a046134b38b516040519283916020830195866132c3565b5190208703612f33577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610233575f9361351060405197889586948594635c73957b60e11b865260048601612bf9565b03915afa908115610332577f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30297613577876107dd8d8f6107d0908f9b6135716131cb9d61359c9c60089c6135fa575b506107dd886107dd886107d088611dd6565b55611dbc565b90815417905501613596613591825461ffff9060201c1690565b6132d3565b906121f4565b6135bf6135b0846107dd8a6107d08d611e58565b611a6b613591825461ffff1690565b5192516040805161ffff978816815296909316602087015291850192909252606084015233946001600160a01b031916929081906080820190565b806109fc5f61360893611d4b565b5f61355f565b905061363060e08a01516101008b015190604091825191825260208201522090565b14155f613481565b8015915061347a565b60c08b0151141591505f613472565b60a08c0151141592505f61346a565b60808d0151141593505f613462565b60608d015160011415945061345b565b60408d015161ffff8d1614159450613454565b60208e0151141594508f61344c565b633466526160e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b50881561336f565b5061010061ffff8b1611613368565b5061ffff8a1615613361565b506136f66103cb600188015460101c61ffff1690565b8c1161335a565b63965c290d60e01b5f5260045ffd5b90610100828203126102335780601f8301121561023357610100604051926137348285611d4b565b8391810192831161023357905b82821061374e5750505090565b8135815260209182019101613741565b909291928311610233579190565b906080116102335790608090565b90939293848311610233578411610233578101920390565b9192916001600160401b038211611cbd57604051916137bb601f8201601f191660200184611d4b565b829481845281830111610233578281602093845f960137010152565b989792959493919690976137ea8a611d88565b9560018060a01b036137fc8854610237565b16156104f857600587015460ff166001880154959061382d9061097290609089901c6001600160401b0316906149fd565b6104e95761384361097261194461193d8f611df0565b6136fd5761ffff8b169586158015613baf575b613ba05761386b8c61338b8f61338590611e0a565b905433916001600160a01b03916138829190611c0e565b1603612f33576138a3600461389b8f611a503391611e8c565b015460ff1690565b613b91576138b38483018361370c565b968d6138c4612d238a519260a01c90565b1491821592613b75575b8215613b5a575b508115613b4b575b508015613b3d575b8015613b2f575b612f33576139158a611a0461390c8c6040519283916020830195866132c3565b5190208d614391565b938460c088015103612f33577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610233575f9361397560405196879586948594635c73957b60e11b865260048601612bf9565b03915afa801561033257613b1b575b5061010061399181612095565b8503612f33576139c86139af610800966139b66139af89838961375e565b3691613792565b6020815191012097611400918761377a565b602081519101206139d88b611ea6565b5403612f33576139f2926139eb92614a24565b9160e00190565b5103612f3357613a5791613a3e6004600893613a1133611a508c611e8c565b805461ffff60a01b191660a08b901b61ffff60a01b1617815590600382015501805460ff19166001179055565b01613a51613591825460101c61ffff1690565b906121da565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15610233575f6040518094633c1bcdef60e21b8252818381613aa93360048301610243565b03925af1928315610332577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb593613b07575b506040805161ffff959095168552602085019190915283015233926001600160a01b03191691606090a3565b806109fc5f613b1593611d4b565b5f613adb565b806109fc5f613b2993611d4b565b5f613984565b508860a087015114156138ec565b5089608087015114156138e5565b9050606087015114155f6138dd565b909150613b6c6103cb60408a01612fa1565b1415905f6138d5565b91506020880151613b8961ffff84166103cb565b1415916138ce565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b50613bc16103cb8260101c61ffff1690565b8711613856565b90610645604051613bd881611cde565b60e06120356001839661202e613c5c825461ffff8116885261ffff808260101c16166020890152613c1761ffff8260201c1660408a019061ffff169052565b61ffff603082901c16606089015261ffff604082901c166080890152613c4d605082901c6001600160401b031660a08a01611f75565b60901c6001600160401b031690565b60c08701611f75565b90610645613d5d6008613c76611fa3565b94613c91613c848254610237565b6001600160a01b03168752565b613c9d60018201613bc8565b6020870152613cae60038201611fc2565b6040870152613d06613cfd6005830154613cd4613ccb8260ff1690565b60608b016121bd565b613cee600882901c6001600160401b031660808b01611f75565b60481c6001600160401b031690565b60a08801611f75565b600681015460c0870152600781015460e0870152015461ffff8116610100860152613d44613d388260101c61ffff1690565b61ffff16610120870152565b61ffff602082901c1661014086015260301c61ffff1690565b61ffff16610160840152565b908160a0910312610233576040519060a082016001600160401b03811183821017611cbd576040528051613d9c81610e63565b8252602081015160208301526040810151604083015260608101519060038210156102335760809160608401520151613dd481610a3d565b608082015290565b600311156113b457565b8054600160401b811015611cbd57613e039160018201815561325c565b81546001600160a01b0393841660039290921b91821b9390911b1916919091179055565b61ffff60019116019061ffff8211611eed57565b905f5160206150fd5f395f51905f52821080613e8c575b15610a1757811580613e82575b610a1757613e6d8183614c8f565b15610a1757613e7b91614d41565b15610a1757565b5060018114613e5f565b505f5160206150fd5f395f51905f528110613e52565b91608093916040519384526020840152604083015260608201522090565b602081830312610233578051906001600160401b03821161023357019080601f83011215610233578151916001600160401b038311611cbd578260051b9060405193613f0f6020840186611d4b565b845260208085019282010192831161023357602001905b828210613f335750505090565b8151815260209182019101613f26565b5f19810191908211611eed57565b80518210156121375760209160051b010190565b8054905f815581613f74575050565b5f5260205f20908101905b818110613f8a575050565b5f8155600101613f7f565b6001600160a01b03613faf613fa983611d88565b54610237565b161561432a57613fbe81611e0a565b80546060905f54613fd38160ff9060e01c1690565b6142b5575b50613fe38251612c20565b925f5b828110614232575050505f5b8281106140a05750505061400d61400882611e0a565b613f65565b5f61401782611ea6565b5561402f61402482611e24565b60015f918281550155565b61407061403b82611d88565b60085f918281558260018201558260028201558260038201558260048201558260058201558260068201558260078201550155565b6001600160a01b0319167f457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c9965f80a2565b8061421a575f5b60015b61010061ffff821611156140c2575050600101613ff2565b6140d3816107dd846107d08a611dbc565b54614191575b806140f46103cb6130af614137946107dd876107d08d611e58565b61416d575b61411761410d826107dd866107d08c611e72565b5460101c60ff1690565b614154575b61412d816107dd856107d08b611da2565b5461413c576132d3565b6140aa565b5f61414e826107dd866107d08c611da2565b556132d3565b614168614024826107dd866107d08c611e72565b61411c565b61418c614181826107dd866107d08c611e58565b805461ffff19169055565b6140f9565b6141a2816107dd846107d08a611dbc565b5460015b602061ffff821611156141d3575050805f6141cb614137936107dd866107d08c611dbc565b5590506140d9565b6141f6908389828760018061ffff841689901c16146141fb575b505050506132d3565b6141a6565b5f936107dd6107dd926107d061421096611dd6565b55838982876141ed565b61422c61422682613f43565b83613f51565b516140a7565b806142af61429261424f6142486001958761325c565b9054611c0e565b61426961425f82611a508d611df0565b805460ff19169055565b5f6142886142768c611e3e565b6142826103cb88612c20565b9061203c565b55611a508a611e8c565b60045f918281558260018201558260028201558260038201550155565b01613fe6565b9091506142cf906109919060401c6001600160a01b031681565b5f60405180926376bc6b8f60e11b825281806142ee8960048301610a59565b03915afa908115610332575f91614308575b50905f613fd8565b61432491503d805f833e61431c8183611d4b565b810190613ec0565b5f614300565b50565b5f5160206150fd5f395f51905f529160405190602082019260018060a01b03191683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c815261438a606c82611d4b565b5190200690565b5f5160206150fd5f395f51905f529160405190602082019260018060a01b03191683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c815261438a606c82611d4b565b5f5160206150fd5f395f51905f529160405190602082019260018060a01b03191683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c815261438a606c82611d4b565b909193929461446b906104006144646201000082612c3c565b918661377a565b916144966103cb60016144896103cb600889015460101c61ffff1690565b96015460101c61ffff1690565b925f5f92610800935b8781106144c3575050505050505050906144bc916108a091614a24565b03612f3357565b61ffff8160051b8a0135169283158015614591575b612f3357600161ffff85161b90818116612f3357179261451261450961424861450086611e0a565b61338b85613248565b611a5085611e8c565b90614524610972600484015460ff1690565b908115614574575b50612f3357600361455d6139af61454389866120c1565b6145558a61455088612c20565b6120c1565b90898c61377a565b6020815191012091015403612f335760010161449f565b90506145896103cb835461ffff9060a01c1690565b14155f61452c565b508784116144d8565b919082608091031261023357815160ff8116810361023357916020810151916060604083015192015190565b906101a0828203126102335780601f83011215610233576101a0604051926145ee8285611d4b565b8391810192831161023357905b8282106146085750505090565b81358152602091820191016145fb565b939499959897969091985f5f9161462d611fb3565b5f81526001602082015293856147e8575b9061464b918101906145c6565b9a8b518760a01c14948515956147d9575b5084156147c4575b5083156147b0575b5082156147a1575b508115614791575b811561477d575b50801561475e575b801561474f575b8015614740575b612f33576101008701976146b76103cb60018b5193015461ffff1690565b11612f33576020885111612f3357611a046146e06146e9946040519283916020830195866132c3565b519020906143ee565b928361016086015103612f33576147006064612095565b8103612f33576139af614713918461376c565b6020815191012003612f3357606461472e9161473693614a24565b916101800190565b5103612f33575190565b50826101408801511415614699565b50806101208801511415614692565b5060e08701516147766103cb60018b015461ffff1690565b141561468b565b9050602060c089015191015114155f614683565b60a089015181511415915061467c565b60808a0151141591505f614674565b60608b015160ff909116141592505f61466c565b60408c015161ffff909116141593505f614664565b60208d0151141594505f61465c565b93509150505f54906148016109728360ff9060e01c1690565b610a0857614864916080908c90614825906109919060401c6001600160a01b031681565b60405163be5b346360e01b81526001600160a01b03198a1660048201526024810188905261ffff90921660448301529093849190829081906064820190565b03915afa908115610332575f905f915f905f9461489e575b61464b93949550939461488d611fb3565b91825260208201529490915061463e565b50509261464b92506148c8915060803d6080116148d6575b6148c08183611d4b565b81019061459a565b92949293509091908461487c565b503d6148b6565b909493929461490d6103cb60016149006148f689612c2e565b9861048090612c3c565b98015460101c61ffff1690565b905f935f5b868110614967575050505050505b6020811061492d57505050565b8060061b83018160051b830135612f335780351590811591614957575b50612f3357600101614920565b600191506020013514155f61494a565b8060061b890161ffff8260051b8a01351690811580156149f4575b612f3357600161ffff83161b97888116612f33578817976149aa856107dd886107d08c611dbc565b541615612f33576149e6906040516149d181611a04602082019460208101359035866132c3565b519020916107dd856107dd886107d08c611dd6565b5403612f3357600101614912565b50858211614982565b60068110156113b4576002149081614a13575090565b6001600160401b0316431115919050565b9291905f5160206150fd5f395f51905f525f940691829060051b8201915b828110614a4f5750505050565b909192945f5160206150fd5f395f51905f5283816020938186358b099008970993929101614a42565b60068110156113b4576001149081614a13575090565b60405190610400614a9f8184611d4b565b368337565b60405190610800614a9f8184611d4b565b9060208110156121375760051b0190565b9060408110156121375760051b0190565b91905f835b60208210614b175750505061040082015f905b60408210614b0157505050610c000190565b6020806001928551815201930191019091614aef565b6020806001928551815201930191019091614adc565b919091614b38614a8e565b614b40614aa4565b93614b4a83611e0a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110614be05750505061ffff165b60208110614bbb5750611a04614baf614bb8939495604051928391602083019586614ad7565b51902091611ea6565b55565b806001614bd9614bd3614bce83956120ab565b612c20565b88614ac6565b5201614b89565b80614bed614c2192612c20565b614bf78288614ab5565b5260a0614c07614248838761325c565b604051809481926313a4120960e31b835260048301610243565b0381865afa918215610332576001926040915f91614c71575b506020810151614c52614c4c856120ab565b8d614ac6565b520151614c6a614c64614bce846120ab565b8b614ac6565b5201614b76565b614c89915060a03d8111611b6157611b538183611d4b565b5f614c3a565b5f5160206150fd5f395f51905f528110801590614d2a575b614d24575f5160206150fd5f395f51905f52818192099180095f5160206150fd5f395f51905f5282065f5160206150fd5f395f51905f5203905f5160206150fd5f395f51905f528211611eed575f5160206150fd5f395f51905f528080838180965f0896095f5160206150dd5f395f51905f520960010892081490565b50505f90565b505f5160206150fd5f395f51905f52821015614ca7565b90811580614e60575b614e5957614d5a81838181614ef9565b919092614d6982828587614ef9565b9390945f94600196607e805b614d935750505050505050159081614d8b575090565b600191501490565b5f19019660037f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1614dc38a6120ab565b1c168115801590614e4e575b614e30575b60018103614df257508686899a614dea93614ef9565b989097614d75565b60028103614e0857508484899a614dea93614ef9565b909790600314614e19575b80614d75565b978282614e27928b9a614ef9565b98909750614e13565b988082614e4092614e4894614ef9565b908181614ef9565b98614dd4565b5060018a1415614dcf565b5050600190565b5060018114614d4a565b5f5160206150fd5f395f51905f5290065f5160206150fd5f395f51905f52035f5160206150fd5f395f51905f528111611eed575f5160206150fd5f395f51905f529060010890565b905f5160206150fd5f395f51905f5290065f5160206150fd5f395f51905f52035f5160206150fd5f395f51905f528111611eed575f5160206150fd5f395f51905f52910890565b919392909382158061506c575b6150655780158061505b575b615055575f5160206150fd5f395f51905f52818409915f5160206150fd5f395f51905f528187095f5160206150fd5f395f51905f528185095f5160206150fd5f395f51905f52905f5160206150dd5f395f51905f5209935f5160206150fd5f395f51905f52907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000009614fa391614eb2565b935f5160206150fd5f395f51905f528460010893614fc090614e6a565b965f5160206150fd5f395f51905f52888609614fdb90615076565b97885f5160206150fd5f395f51905f529109935f5160206150fd5f395f51905f529109915f5160206150fd5f395f51905f529109905f5160206150fd5f395f51905f529108905f5160206150fd5f395f51905f529109935f5160206150fd5f395f51905f5291095f5160206150fd5f395f51905f52910990565b50509190565b5060018214614f12565b9350919050565b5060018514614f06565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f5160206150fd5f395f51905f5260a082015260208160c08160055afa1561023357519056fe1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220168bca377d91c725177fb2659003d86bb046b89aa72a0e8bc3c8c87d6c6ff13364736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _finalizeVerifier common.Address, _decryptCombineVerifier common.Address) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _finalizeVerifier, _decryptCombineVerifier)
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
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetEpoch(epochId [12]byte) (IDKGManagerEpoch, error) {
	return _DKGManager.Contract.GetEpoch(&_DKGManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbe59b8ea.
//
// Solidity: function getEpoch(bytes12 epochId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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

// CreateEpoch is a paid mutator transaction binding the contract method 0x2c7c7642.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateEpoch(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createEpoch", threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, decryptionPolicy)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x2c7c7642.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, decryptionPolicy)
}

// CreateEpoch is a paid mutator transaction binding the contract method 0x2c7c7642.
//
// Solidity: function createEpoch(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateEpoch(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateEpoch(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, decryptionPolicy)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactor) ExtendRegistration(opts *bind.TransactOpts, epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "extendRegistration", epochId)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 epochId) returns()
func (_DKGManager *DKGManagerSession) ExtendRegistration(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ExtendRegistration(&_DKGManager.TransactOpts, epochId)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 epochId) returns()
func (_DKGManager *DKGManagerTransactorSession) ExtendRegistration(epochId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ExtendRegistration(&_DKGManager.TransactOpts, epochId)
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
	SeedBlock        uint64
	LotteryThreshold *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEpochCreated is a free log retrieval operation binding the contract event 0x2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d06.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
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

// WatchEpochCreated is a free log subscription operation binding the contract event 0x2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d06.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
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

// ParseEpochCreated is a log parse operation binding the contract event 0x2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d06.
//
// Solidity: event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
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

// DKGManagerRegistrationExtendedIterator is returned from FilterRegistrationExtended and is used to iterate over the raw logs and unpacked data for RegistrationExtended events raised by the DKGManager contract.
type DKGManagerRegistrationExtendedIterator struct {
	Event *DKGManagerRegistrationExtended // Event containing the contract specifics and raw log

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
func (it *DKGManagerRegistrationExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRegistrationExtended)
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
		it.Event = new(DKGManagerRegistrationExtended)
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
func (it *DKGManagerRegistrationExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRegistrationExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRegistrationExtended represents a RegistrationExtended event raised by the DKGManager contract.
type DKGManagerRegistrationExtended struct {
	EpochId                 [12]byte
	NewSeedBlock            uint64
	NewRegistrationDeadline uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRegistrationExtended is a free log retrieval operation binding the contract event 0x9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c.
//
// Solidity: event RegistrationExtended(bytes12 indexed epochId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) FilterRegistrationExtended(opts *bind.FilterOpts, epochId [][12]byte) (*DKGManagerRegistrationExtendedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RegistrationExtended", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRegistrationExtendedIterator{contract: _DKGManager.contract, event: "RegistrationExtended", logs: logs, sub: sub}, nil
}

// WatchRegistrationExtended is a free log subscription operation binding the contract event 0x9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c.
//
// Solidity: event RegistrationExtended(bytes12 indexed epochId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) WatchRegistrationExtended(opts *bind.WatchOpts, sink chan<- *DKGManagerRegistrationExtended, epochId [][12]byte) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RegistrationExtended", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRegistrationExtended)
				if err := _DKGManager.contract.UnpackLog(event, "RegistrationExtended", log); err != nil {
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

// ParseRegistrationExtended is a log parse operation binding the contract event 0x9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c.
//
// Solidity: event RegistrationExtended(bytes12 indexed epochId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) ParseRegistrationExtended(log types.Log) (*DKGManagerRegistrationExtended, error) {
	event := new(DKGManagerRegistrationExtended)
	if err := _DKGManager.contract.UnpackLog(event, "RegistrationExtended", log); err != nil {
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
