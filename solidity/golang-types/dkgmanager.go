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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_epochDurationBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_committeeSelectionBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_keyAssemblyBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finalizeGapBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minCommitteeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_maxLotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LOTTERY_ALPHA_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_COMMITTEE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"appManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ciphertextCount\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochDurationBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSelectionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyAssemblyDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liveNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEpochStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAppManager\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pokAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pokAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pokZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pokAx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pokAy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pokZ\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeFilled\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochLive\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInSnapshot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x610260806040523461043f576101a08161463380380380916100218285610443565b83398101031261043f5780519063ffffffff82169081830361043f576100496020820161047a565b926100566040830161047a565b6100626060840161047a565b9061006f6080850161047a565b61007b60a0860161047a565b60c08601519160e087015194610100880151946101208901519a6101408a016100a39061048e565b986100b16101608c0161048e565b9a610180016100bf9061048e565b9b4663ffffffff1603610430576001600160a01b03821615610421576001600160a01b038316158015610410575b80156103ff575b80156103ee575b6103df5763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b602482015260188152610137603882610443565b5190201660c05260e052610100526101205261014052806103d957506064915b806103d357506019905b806103cd57506019955b806103c757506005915b61018883610183898561049d565b61049d565b9260018311908115916103be575b81156103b5575b5080156103ab575b801561039b575b61035d57610160526001600160401b038181166101e05261ffff9690916101d3919061049d565b16610200526001600160401b03166102205280841661039557506001905b8161018052838116155f1461038f57506001915b826101a052838116155f1461038757508280925b836101c05216928391161191821561037c575b50811561036c575b5061035d57336102405260405161417490816104bf823960805181611f37015260a05181818161038501528181611342015281816128c20152613aa4015260c051818181612adc01528181612cf501528181612d5b01526138fb015260e051818181611201015281816130a90152613a4e015261010051818181610ab40152818161185a0152611e99015261012051818181611bf20152818161305201526133ca0152610140518181816101e0015281816122f201526127ce015261016051818181613bbc0152613cd8015261018051818181610f2b0152612e9b01526101a0518181816102320152612e6a01526101c05181818161029a0152612e3901526101e05181612984015261020051816129bb015261022051816129ef01526102405181612fa30152f35b63d06b96b160e01b5f5260045ffd5b612710915061ffff16105f610234565b60201091505f61022c565b839092610219565b91610205565b906101f1565b506001600160401b0381116101ac565b50808310156101a5565b9050155f61019d565b88159150610196565b91610175565b9561016b565b90610161565b91610157565b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100fb565b506001600160a01b038516156100f4565b506001600160a01b038416156100ed565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761046657604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361043f57565b519061ffff8216820361043f57565b919082018092116104aa57565b634e487b7160e01b5f52601160045260245ffdfe60e0806040526004361015610012575f80fd5b5f905f3560e01c90816304da574014613ad35750806306433b1b14613a8f578063074a75e114613a2957806318287e5f1461391f57806323488be5146138df578063268ae2a1146138b35780632de546d51461386d5780633353ec6e1461383857806349c61a1214613171578063510ba2df146131215780635a8f2bb3146130d857806363f314cd14613093578063669a76a91461302b5780636d16897d14612f7a5780636f067f6314612f3157806371712c291461021257806371a5978c1461280f57806372517b4b146127a757806377235ee114611f5b57806385e1f4d014611f1a5780638dc1f53a14611e725780638e515b5014611c2157806393c3d3a814611bdc5780639bbada6714611b40578063a305e0f31461159d578063a4adcd7f14611576578063b7bca61514610f4f578063bd11c4c014610f10578063be59b8ea14610c09578063bea5210d14610ae3578063bf19220914610a9e578063ca3c0458146109d6578063d3720aac146108c4578063d9933767146102be578063d9e9ca2e1461027f578063ebe86c1314610256578063f03a489814610217578063fa8f5e96146102125763fe1604b5146101cb575f80fd5b3461020f578060031936011261020f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b80fd5b613ba5565b503461020f578060031936011261020f57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461020f578060031936011261020f576001546040516001600160a01b039091168152602090f35b503461020f578060031936011261020f57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461020f57602036600319011261020f576001600160a01b03196102e1613af8565b1680825260026020526040822080546001600160a01b0316156108b557600281019081549060ff8216600182019182549160068110156108a1576001148061088a575b1561087b57600581019061ffff808354169360101c1683101561086c57868852600360209081526040808a20335f908152925290205460ff1661085d576003810180549081156107ea575b506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316969060c0816024818b5afa9081156107df578b916107c0575b50606081015160038110156107ac575f190161079d5760a001516001600160401b0390811660489290921c16111561078e57600490604051602081019182523360601b60408201526034815261041d605482613c82565b519020910154111561077f57858752600460205260408720805490600160401b82101561076b579261ffff9594939261046083889560016104e596018155613eb2565b81546001600160a01b03600392831b90811b19909116339182901b179092558a8c5260209081526040808d205f9384529091529020805460ff19166001179055836104aa83613e8a565b168419825416179055604051818152887f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a3613e8a565b915460101c16921682146104f7578480f35b604051916104006105088185613c82565b3684376040519161080061051c8185613c82565b368437858752600460205260408720875b838110610636575050505b602081106105f057506040519060208201928387905b602082106105da57505050610420820186905b604082106105c457505050610c00815261057d610c2082613c82565b519020828452600b6020526040842055805460ff191660021790557f23a9ea75665bd065d8fc1c53ceb8c23343c59630fcf7ad5083dc4b1057bbb0678280a25f8080808480f35b6020806001928551815201930191019091610561565b602080600192855181520193019101909161054e565b8060011b8181046002148215171561062257600181018091116106225790600161061b81938561405a565b5201610538565b634e487b7160e01b87526011600452602487fd5b60018101808211610757576020821015610743578160051b87015261065b8183613eb2565b90546040516313a4120960e31b815260039290921b1c6001600160a01b0316600482015260c081602481875afa908115610738578a9161070a575b506020810151908260011b91838304600214841517156106f657906040916106be848a61405a565b52015190600181018091116106e257906106db600193928861405a565b520161052d565b634e487b7160e01b8b52601160045260248bfd5b634e487b7160e01b8c52601160045260248cfd5b61072b915060c03d8111610731575b6107238183613c82565b810190613f08565b5f610696565b503d610719565b6040513d8c823e3d90fd5b634e487b7160e01b8a52603260045260248afd5b634e487b7160e01b8a52601160045260248afd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b633802147960e11b8952600489fd5b63aba4733960e01b8b5260048bfd5b634e487b7160e01b8c52602160045260248cfd5b6107d9915060c03d60c011610731576107238183613c82565b5f6103c6565b6040513d8d823e3d90fd5b9050608886901c6001600160401b03164381101561084e574090811561083f57819055877fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602083604051908152a25f61036f565b6302504bb360e61b8a5260048afd5b63172181cb60e21b8a5260048afd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b50604082901c6001600160401b0316431115610324565b634e487b7160e01b88526021600452602488fd5b63d5b25b6360e01b8352600483fd5b503461020f57604036600319011261020f576108de613af8565b60243591906001600160a01b03831683036109d257906040918160a0845161090581613c67565b8281528260208201528286820152826060820152826080820152015260018060a01b03191681526005602052209060018060a01b03165f5260205260c060405f2060405161095281613c67565b81549160018060a01b0383169283835261ffff602084019160a01c16815260018201546040840190815261ffff6002840154926060860193845260a060ff600460038801549760808a01988952015416960195151586526040519687525116602086015251604085015251606084015251608083015251151560a0820152f35b5080fd5b503461020f57602036600319011261020f576001600160a01b03196109f9613af8565b168152600460205260408120604051908160208254918281520190819285526020852090855b818110610a7f5750505082610a35910383613c82565b604051928392602084019060208552518091526040840192915b818110610a5d575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610a4f565b82546001600160a01b0316845260209093019260019283019201610a1f565b503461020f578060031936011261020f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461020f57608036600319011261020f57608090610b00613af8565b9060243591610b0d613b20565b610b15613b31565b91836060604051610b2581613c4c565b828152826020820152826040820152015260018060a01b031916918284526007602052604084208585526020526040842061ffff82165f5260205261ffff6040600180825f20548487161c161495865f14610c025784975b8715610bfb5784965b825260066020528282209082526020522091165f5260205261ffff60405f2091165f5260205261ffff60405f20549160608260405196610bc588613c4c565b16958681528360208201931683526040810194855201938452604051948552511660208401525160408301525115156060820152f35b8196610b86565b8097610b7d565b503461020f57602036600319011261020f57610c23613af8565b81610160604051610c3381613bfa565b828152604051610c4281613bdf565b8381528360208201528360408201528360608201528360808201528360a08201528360c082015260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015282610140820152015260018060a01b031916815260026020526040812060405190610cca82613bfa565b80546001600160a01b03168252604051610ce381613bdf565b600182015461ffff8116825261ffff8160101c16602083015261ffff8160201c16604083015261ffff8160301c16606083015260018060401b038160401c16608083015260018060401b038160801c1660a083015260c01c60c08201526020830190815260028201549160ff83169060408501916006811015610efc578252606085019360018060401b038160081c168552608086019060018060401b038160481c16825260a087019060018060401b039060881c16815260038301549160c08801928352600560048501549460e08a0195865201549661010089019961ffff89168b526101208a019661ffff8a60101c1688526101408b019861ffff8b60201c168a5261ffff6101608d019b60301c168b526040519b60018060a01b039051168c525161ffff81511660208d015261ffff60208201511660408d015261ffff60408201511660608d015261ffff60608201511660808d015260018060401b0360808201511660a08d015260018060401b0360a08201511660c08d015260c060018060401b039101511660e08c015251906006821015610ee857506101008a0152516001600160401b039081166101208a01529051811661014089015290511661016087015251610180860152516101a0850152845161ffff9081166101c0860152905181166101e08501529051811661020084015281511661022083015261024082f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b87526021600452602487fd5b503461020f578060031936011261020f57602060405161ffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461020f5760e036600319011261020f57610f69613af8565b610f71613b0f565b6064356044356084356001600160401b03811161143857610f96903690600401613b78565b909360a4356001600160401b03811161157257610fb7903690600401613b78565b90929060c4356001600160401b03811161156e57610fd9903690600401613b78565b6001600160a01b03198a16808c52600260205260408c208054919a9097939490939092916001600160a01b03161561155f5760ff600289015416600189015490600681101561154b5760021480611534575b156115255760408e8d815260036020522060018060a01b0333165f5260205260ff60405f205416156115165761ffff88169c8d158015611506575b6114f7576110898f808f604092526004602052206110838b613e9e565b90613eb2565b90543360039290921b1c6001600160a01b03160361144c5760408f8e815260056020522060018060a01b0333165f5260205260ff600460405f200154166114e857868401969594939291908f610100888a031261020f5788601f8901121561020f5750604051976110fc6101008a613c82565b8861010089019182116114ce5788905b8282106114be5750505087519060a01c14908115916114ab575b8115611494575b508015611486575b8015611478575b801561146a575b61145b5761010094612000880361144c575f51602061411f5f395f51905f528f8e908d8f6111a06111768e36908d613d92565b6020815191012091611192604051938492602084019687613edb565b03601f198101835282613c82565b519020905060405190602082019283527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c81526111ea606c82613c82565b51902006938460c08901510361143c578f939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b1561143857859361125360405196879586948594635c73957b60e11b865260048601613df7565b03915afa80156113ee5761141f575b505061080085811161141b57611279368285613d92565b60208151910120956114009081831161141757811161141757816112a292850191033691613d92565b602081519101208a8d52600b60205260408d20540361140857916112c89160e093614006565b910151036113f957858852600560208181526040808b20335f90815292529020805461ffff60a01b191660a09490941b61ffff60a01b169390931783556003830191909155600491909101805460ff1916600117905501805460101c61ffff9081169081146106225790600161133f920190613e32565b847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156109d257818091602460405180948193633c1bcdef60e21b83523360048401525af180156113ee576113d5575b5050604051938452602084015260408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b816113df91613c82565b6113ea57845f61139a565b8480fd5b6040513d84823e3d90fd5b63d1fed5fd60e01b8852600488fd5b63d1fed5fd60e01b8c5260048cfd5b5f80fd5b8c80fd5b8161142991613c82565b611434578b5f611262565b8b80fd5b8580fd5b5063d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8f5260048ffd5b63d1fed5fd60e01b8e5260048efd5b508a60a08701511415611143565b50896080870151141561113c565b508c60608701511415611135565b905061ffff60408801519160101c1614155f61112d565b602088015161ffff821614159150611126565b813581526020918201910161110c565b50508f80fd5b634e487b7160e01b5f52604160045260245ffd5b6305d252c360e01b8f5260048ffd5b63652122d960e01b8f5260048ffd5b508d61ffff8360101c1610611066565b63965c290d60e01b8e5260048efd5b63268dbf6760e21b8e5260048efd5b50608081901c6001600160401b031643111561102b565b634e487b7160e01b8f52602160045260248ffd5b63d5b25b6360e01b8d5260048dfd5b8980fd5b8780fd5b503461020f578060031936011261020f57546040516001600160401b039091168152602090f35b503461020f5761016036600319011261020f576115b8613af8565b6115c0613b20565b906115c9613b31565b610124356001600160401b0381116113ea576115e9903690600401613b78565b90610144356001600160401b0381116119bc5761160a903690600401613b78565b6001600160a01b03198616885260026020526040882080549294909390926001600160a01b031615611b315760ff6002850154166006811015611b1d57600303611b0e576001600160a01b031987168952600360209081526040808b20335f908152925290205460ff1615611aff5761ffff8816158015611ae7575b8015611adb575b8015611acc575b8015611ac1575b611ab3576001600160a01b0319871689526004602052604089206116c2906110838a613e9e565b90543360039290921b1c6001600160a01b031603611a95576001600160a01b031987168952600d60209081526040808b206024358c528252808b2061ffff89165f90815292529020548015611aa45761172560e43560c43560a435608435613fe8565b03611a955760018060a01b0319871689526007602052604089206024358a526020526040892061ffff87165f5260205260405f20600161ffff8a161b905416611a8657818501946102008187031261156e5785601f8201121561156e576040519561179261020088613c82565b8690610200830111611a825781905b61020083018210611a7257505060018060a01b031988168a52600a60205260408a2061ffff8a165f5260205260405f205486518960a01c1490811591611a61575b8115611a4e575b8115611a3e575b8115611a2b575b8115611a1a575b8115611a09575b8115611a00575b81156119de575b506119cf57610120860193610140855197019687516040519060208201928352604082015260408152611847606082613c82565b51902061010435036119c0578a939291907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156114385785936118ac60405196879586948594635c73957b60e11b865260048601613df7565b03915afa80156113ee576119a7575b5050611950600561ffff9360018060a01b031988168a52600660205260408a206024358b5260205260408a208588165f5260205260405f20858a165f5260205260405f2061010435905560018060a01b031988168a52600760205260408a206024358b5260205260408a208588165f5260205260405f206001868b161b81541790550183600181835460201c16011690613e4c565b51915160408051968316875291909316602086015284015260608301523391602435916001600160a01b031916907f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a30290608090a480f35b816119b191613c82565b6119bc57865f6118bb565b8680fd5b63d1fed5fd60e01b8b5260048bfd5b63d1fed5fd60e01b8a5260048afd5b9050604060e0880151610100890151825191825260208201522014155f611813565b8015915061180c565b60c088015160a43514159150611805565b60a0880151608435141591506117fe565b608088015161ffff8c16141591506117f7565b60608801516001141591506117f0565b604088015161ffff8a16141591506117e9565b6020880151602435141591506117e2565b81358152602091820191016117a1565b8a80fd5b633466526160e01b8952600489fd5b63d1fed5fd60e01b8952600489fd5b6346f551f560e01b8a5260048afd5b62d949df60e51b8952600489fd5b50610104351561169b565b5061010061ffff871611611694565b5061ffff86161561168d565b5061ffff600185015460101c1661ffff891611611686565b63965c290d60e01b8952600489fd5b63268dbf6760e21b8952600489fd5b634e487b7160e01b8a52602160045260248afd5b63d5b25b6360e01b8952600489fd5b503461020f5761ffff6040611b5436613b42565b949182848051611b6381613c31565b828152826020820152015260018060a01b031916825260096020528282209082526020522091165f52602052606060405f20604051611ba181613c31565b81546040600161ffff83169485855260ff602086019460101c1615158452015492019182526040519283525115156020830152516040820152f35b503461020f578060031936011261020f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461020f5761012036600319011261020f57611c3c613af8565b6001600160a01b03191680825260026020526040822080549091602435916044359060a4359060843590606435906001600160a01b031615611e635760ff6002880154166006811015611e4f57600303611e4057611c9a8185613f81565b611ca48383613f81565b84885260086020526040882086895260205261ffff611cc88160408b205416613e8a565b16966101008811611e315786611daa575b9261010092869592611d6e60057ff3ca88585465ab591836f08efb3f8993e60f41ef7f7d7b82bf35c776776432b2978d8d9c9b60209f5260088f52604081208c82528f52604081208d61ffff198254161790556040611d3a8888888c613fe8565b918c81526020600d90528d828220908252602052208d5f528f5260405f20550161ffff600181835460301c16011690613e6b565b604051933385528b85015260408401526060830152608082015260c43560a082015260e43560c08201526101043560e0820152a4604051908152f35b60015460ff8160a01c1615611e22576001600160a01b031689813b1561020f5760849160405192838092633a54cd5d60e11b82528b60048301528c60248301528d60448301523360648301525afa801561073857611e09575b50611cd9565b611e148a8092613c82565b611e1e575f611e03565b8880fd5b63023b34fb60e11b8a5260048afd5b63464e67af60e01b8952600489fd5b63268dbf6760e21b8852600488fd5b634e487b7160e01b89526021600452602489fd5b63d5b25b6360e01b8852600488fd5b503461020f578060031936011261020f5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611f0e5790611edb575b602090604051908152f35b506020813d602011611f06575b81611ef560209383613c82565b810103126114175760209051611ed0565b3d9150611ee8565b604051903d90823e3d90fd5b503461020f578060031936011261020f57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461020f5761010036600319011261020f57611f76613af8565b60a052611f81613b20565b60c05260a4356001600160401b0381116109d257611fa3903690600401613b78565b90919060c4356001600160401b0381116127a357611fc5903690600401613b78565b9060e4356001600160401b0381116113ea57611fe5903690600401613b78565b60a0516001600160a01b03191686526002602052604086208054909392906001600160a01b0316156127945760ff60028501541660068110156108a15760030361087b5761ffff60c05116158015612783575b8015612779575b61276a5760a0516001600160a01b0319168752600d602090815260408089206024358a52825280892060c05161ffff165f908152925290205492831561275b5760a0516001600160a01b0319168852600760209081526040808a206024358b528252808a2060c05161ffff165f9081529252902054889590805b61273b5750600101549461ffff86161161272c5760a0516001600160a01b0319168852600960209081526040808a206024358b528252808a2060c05161ffff165f9081529252902060808190525496601088901c60ff1661271d578890899160405161212481613c16565b8b81526001602082015260243561263b575b868601936101a08786031261141b5784601f8801121561141b57604051946121606101a087613c82565b85906101a08901116126375787905b6101a089018210612627575050845160a05160a01c1492831593612616575b8315612600575b83156125ec575b5082156125dd575b5081156125cd575b81156125b9575b5080156125a7575b8015612596575b8015612585575b6119cf5761ffff8716610100830151106119cf576020610100830151116119cf57610c8081036119cf575f51602061411f5f395f51905f5261220c36838e613d92565b6020815191012060405161222e81611192602082019460843560643587613edb565b51902060405160a0516001600160a01b031916602082019081527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c80830193909352918152612285606c82613c82565b519020069081610160840151036119c05760801161156e57604051956122ac60a088613c82565b608087526020870160808d019736891161141b5760808e83378c60a0820152519020036119cf5760648b6122df92614006565b61018082015103611a95576101000151957f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561156e57908993929161234760405196879586948594635c73957b60e11b865260048601613df7565b03915afa801561257a57908691612565575b50508511612551576104808501851161255157839060101c61ffff16815b838110612454575050505b602081106123f55750620100009062ff000019161760805155608435600160805101556040516064358152608435602082015261ffff60c0511690602435907f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae1604060018060a01b031960a0511692a480f35b6104808160061b85010160808260051b860101356124455780351590811591612435575b5061242657600101612382565b63d1fed5fd60e01b8352600483fd5b600191506020013514155f612419565b63d1fed5fd60e01b8452600484fd5b6104808160061b88010161ffff60808360051b8a010135169081158015612548575b6113f9576001821b94858116611a955785179460018060a01b031960a0511689526007602052604089206024358a526020526040892061ffff60c051165f5260205260405f205416156113f9576040516020808201928035845201356040820152604081526124e6606082613c82565b5190209060018060a01b031960a05116885260066020526040882060243589526020526040882061ffff60c051165f5260205261ffff60405f2091165f5260205260405f20540361253957600101612377565b63d1fed5fd60e01b8652600486fd5b50838211612476565b634e487b7160e01b84526011600452602484fd5b8161256f91613c82565b6113ea57845f612359565b6040513d88823e3d90fd5b5060843561014083015114156121c9565b5060643561012083015114156121c2565b5061ffff871660e083015114156121bb565b9050602060c084015191015114155f6121b3565b60a08401518151141591506121ac565b6080850151141591505f6121a4565b606086015160ff909116141592505f61219c565b925061ffff60c051166040860151141592612195565b60208601516024351415935061218e565b813581526020918201910161216f565b8d80fd5b5050905060015460ff8160a01c1615611e225760405163be5b346360e01b815260a0516001600160a01b0319166004820152602480359082015260c05161ffff1660448201529190608090839060649082906001600160a01b03165afa918215610738578a908b918c918d956126c8575b509193604051916126bc83613c16565b82526020820152612136565b94505050506080823d608011612715575b816126e660809383613c82565b8101031261156e5781519160ff83168303611a825760208101516040820151606090920151939091905f6126ac565b3d91506126d9565b63955c0c4960e01b8952600489fd5b63032cddf960e11b8852600488fd5b5f198101908082116106e25716955f1981146107575760010195806120b9565b6346f551f560e01b8852600488fd5b636d28699160e01b8752600487fd5b506064351561203f565b5061010061ffff60c0511611612038565b63d5b25b6360e01b8752600487fd5b8280fd5b503461020f578060031936011261020f5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611f0e5790611edb57602090604051908152f35b503461020f57608036600319011261020f5760043561ffff81168091036109d257612838613b0f565b612840613b20565b91612849613b31565b906001600160401b0361285a613cbe565b164310612f225780158015612f16575b8015612f09575b8015612efb575b8015612eef575b8015612ede575b8015612ed1575b8015612ec2575b8015612e95575b8015612e64575b8015612e33575b612e2457604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561257a578691612dea575b506001600160401b0316918215612ddb5761ffff80911693169161291e8385613d7f565b9080612710029061271082040361062257808210612dae5750505f19915b8554946001600160401b03808716908114612d9a576001600160401b03600191820181166001600160401b0319989098168817895543811690910195908611612d9a576129b27f0000000000000000000000000000000000000000000000000000000000000000436001600160401b0316613ca5565b9061ffff6129e97f0000000000000000000000000000000000000000000000000000000000000000436001600160401b0316613ca5565b93612a1d7f0000000000000000000000000000000000000000000000000000000000000000436001600160401b0316613ca5565b9560405197612a2b89613bdf565b885260208801521660408087019190915260608601919091526001600160401b03918216608086015291811660a08501529190911660c08301525190612a7082613bfa565b33825260208201526001604082015283606082015260018060401b034316608082015260018060401b03831660a08201528460c08201528160e08201528461010082015284610120820152846101408201528461016082015260018060a01b03198463ffffffff60401b7f000000000000000000000000000000000000000000000000000000000000000060401b161760a01b16855260026020526040852060018060a01b0382511660018060a01b031982541617815560018101602083015161ffff808251161661ffff19835416178255612b5461ffff60208301511683613e32565b612b6661ffff60408301511683613e4c565b612b7861ffff60608301511683613e6b565b608081810151835460a084015160c0948501516001600160401b03909216604093841b600160401b600160801b031617931b600160801b600160c01b0316929092179190921b6001600160c01b031916179091558201516006811015610efc576002820180546060850151608086015160a08701516001600160c81b031990931660ff9095169490941760089190911b610100600160481b03161760489390931b600160481b600160881b03169290921760889290921b600160881b600160c81b031691909117905560c0820151600382015560e08201516004820155610100820151600591909101805461ffff191661ffff92831617815561012083015160209793612cab93916101609190612c9190841685613e32565b612ca2836101408301511685613e4c565b01511690613e6b565b8054600160401b600160801b03191643604081811b600160401b600160801b03169290921790925580516001600160401b03928316815291909316858201528083019190915233917f0000000000000000000000000000000000000000000000000000000000000000901b63ffffffff60401b16831760a01b6001600160a01b031916907f1bd7dbfb91d6bbeee799f81d11452e0d0d87712734cbf66805ed6041d7d17a4d90606090a3604080517f000000000000000000000000000000000000000000000000000000000000000090911b63ffffffff60401b1690911760a01b6001600160a01b0319168152f35b634e487b7160e01b88526011600452602488fd5b8015612dc757612dc191905f1904613d7f565b9161293c565b634e487b7160e01b87526012600452602487fd5b63d06b96b160e01b8652600486fd5b90506020813d602011612e1c575b81612e0560209383613c82565b8101031261143857612e1690613e1e565b5f6128fa565b3d9150612df8565b63d06b96b160e01b8552600485fd5b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff8316116128a9565b5061ffff7f00000000000000000000000000000000000000000000000000000000000000001661ffff8416106128a2565b5061ffff7f000000000000000000000000000000000000000000000000000000000000000016811061289b565b5061271061ffff831610612894565b508061ffff85161061288d565b5061ffff831661ffff851611612886565b5061ffff84161561287f565b50602061ffff841611612878565b5061ffff83168111612871565b5061ffff83161561286a565b63268dbf6760e21b8552600485fd5b503461020f57604036600319011261020f5760209061ffff906040906001600160a01b0319612f5e613af8565b1681526008845281812060243582528452205416604051908152f35b503461020f57602036600319011261020f576004356001600160a01b038116908190036109d2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361301d576001549060ff8260a01c1661300e578015612fff576001600160a81b031990911617600160a01b1760015580f35b63e6c4247b60e01b8352600483fd5b6373253a9760e01b8352600483fd5b6282b42960e81b8252600482fd5b503461020f578060031936011261020f5760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115611f0e5790611edb57602090604051908152f35b503461020f578060031936011261020f576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b503461020f5761ffff60406130ec36613b42565b949160018060a01b031916825260096020528282209082526020522091165f526020526020600160405f200154604051908152f35b503461020f57604036600319011261020f5761ffff6040613140613af8565b92613149613b0f565b9360018060a01b0319168152600a6020522091165f52602052602060405f2054604051908152f35b346114175760e03660031901126114175761318a613af8565b6084356001600160401b038111611417576131a9903690600401613b78565b919060a4356001600160401b038111611417576131ca903690600401613b78565b909360c4356001600160401b038111611417576131eb903690600401613b78565b6001600160a01b031986165f90815260026020526040902080549297909491929091906001600160a01b0316156138295760ff6002860154166006811015613815576003811461380657600119016137f7576001850154928360c01c43106137f75761ffff600587015460101c169061ffff8560201c1682106137e8576024351580156137de575b80156137d4575b6137c557808a01906101408b8303126114175781601f8c01121561141757604051916101406132a98185613c82565b83908d01918211611417578c905b8282106137b55750505081518a60a01c148015906137a3575b801561378e575b8015613780575b8015613770575b8015613760575b8015613750575b613741576108a0945f9b620114008903613741575f51602061411f5f395f51905f528c8c61332860e0880151918d3691613d92565b60208151910120604051906020820192602435845260443560408401526064356060840152608083015260a082015260a0815261336660c082613c82565b51902060405190602082019260018060a01b03191683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c81526133b7606c82613c82565b51902006958661010086015103613741577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15611417575f9361341c60405196879586948594635c73957b60e11b865260048601613df7565b03915afa80156137365761371f575b50610120015162010400959094610400918711611417579392909182899384019088038c968d5b8481106135905750505050506134689350614006565b0361358157820180831161356d576001600160a01b031984168552600c6020526040852090358155620104208301356001919091015560028101805460ff191660031790556005015460101c61ffff169062010c008101811161255157835b82811061351b575050507f5d8b14aea3a3af8564f9576bdf230e2b3aad200f22c6268df330139c5634da5d60606040519260243584526044356020850152606435604085015260018060a01b03191692a280f35b80604062010c0060019360061b8501016020825191803583520135602082015220828060a01b031986168752600a6020526040872061ffff808460051b87013516165f5260205260405f2055016134c7565b634e487b7160e01b85526011600452602485fd5b63d1fed5fd60e01b8552600485fd5b919395509193959661ffff8260051b8d013516158015613705575b61145b57600161ffff8d8460051b0135161b811661145b57908c60408f6136058f9695600161ffff8760051b8a0135161b179660018060a01b031985168352600460205261108361ffff858520928860051b013516613e9e565b90546001600160a01b03199094168252600560209081529290912060039190911b9290921c6001600160a01b03165f90815291905260409020600481015460ff161580156136e9575b61144c57610800820282810461080014831517156136bf57600183018084116136d35761080081029080820461080014901517156136bf5761369761369e91600393898b613ef0565b3691613d92565b602081519101209101540361145b57600101908b9593919796949297613452565b634e487b7160e01b5f52601160045260245ffd5b5050634e487b7160e01b8f52601160045260248ffd5b5061ffff8d8360051b01351661ffff825460a01c16141561364e565b5061ffff8360101c1661ffff8d8460051b013516116135ab565b61372c919a505f90613c82565b5f9861012061342b565b6040513d5f823e3d90fd5b63d1fed5fd60e01b5f5260045ffd5b5060643560c083015114156132f3565b5060443560a083015114156132ec565b50602435608083015114156132e5565b5082606083015114156132de565b50604082015161ffff8760101c1614156132d7565b50602082015161ffff871614156132d0565b81358152602091820191016132b7565b63c5f680ed60e01b5f5260045ffd5b506064351561327a565b5060443515613273565b63368f2d7d60e21b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b6337bca76b60e21b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b63d5b25b6360e01b5f5260045ffd5b3461141757602036600319011261141757604061385b613856613af8565b613d17565b60208251918051835201516020820152f35b346114175761387b36613b42565b9160018060a01b0319165f52600d60205260405f20905f5260205261ffff60405f2091165f52602052602060405f2054604051908152f35b34611417575f3660031901126114175760206138cd613cbe565b6040516001600160401b039091168152f35b34611417575f36600319011261141757602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34611417576020366003190112611417576001600160a01b0319613941613af8565b165f81815260026020526040902080546001600160a01b03161561382957600281019060ff82541690600682101561381557600182149182613a0f575b82156139bd575b5050156137f757805460ff191660041790557f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b6002149150816139f5575b816139d6575b508380613985565b905061ffff600181600584015460101c1692015460201c1611836139ce565b600181015460801c6001600160401b0316431191506139c8565b600182015460401c6001600160401b03164311925061397e565b34611417575f3660031901126114175760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015613736575f90611edb57602090604051908152f35b34611417575f366003190112611417576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34611417575f366003190112611417575f5460401c6001600160401b03168152602090f35b600435906001600160a01b03198216820361141757565b6024359061ffff8216820361141757565b6044359061ffff8216820361141757565b6064359061ffff8216820361141757565b6060906003190112611417576004356001600160a01b03198116810361141757906024359060443561ffff811681036114175790565b9181601f84011215611417578235916001600160401b038311611417576020838186019501011161141757565b34611417575f3660031901126114175760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b60e081019081106001600160401b038211176114d457604052565b61018081019081106001600160401b038211176114d457604052565b604081019081106001600160401b038211176114d457604052565b606081019081106001600160401b038211176114d457604052565b608081019081106001600160401b038211176114d457604052565b60c081019081106001600160401b038211176114d457604052565b601f909101601f19168101906001600160401b038211908210176114d457604052565b6001600160401b0391821690821601919082116136bf57565b5f5460401c6001600160401b03168015613d0957613d06907f00000000000000000000000000000000000000000000000000000000000000006001600160401b031690613ca5565b90565b50436001600160401b031690565b604051613d2381613c16565b5f81525f60208201525060018060a01b0319165f52600c60205260405f2060018101548015613d645760405191613d5983613c16565b548252602082015290565b5050604051613d7281613c16565b5f81526001602082015290565b818102929181159184041417156136bf57565b9192916001600160401b0382116114d45760405191613dbb601f8201601f191660200184613c82565b829481845281830111611417578281602093845f960137010152565b908060209392818452848401375f828201840152601f01601f1916010190565b9290613e1090613d069593604086526040860191613dd7565b926020818503910152613dd7565b51906001600160401b038216820361141757565b9063ffff000082549160101b169063ffff00001916179055565b805461ffff60201b191660209290921b61ffff60201b16919091179055565b805461ffff60301b191660309290921b61ffff60301b16919091179055565b61ffff60019116019061ffff82116136bf57565b61ffff5f199116019061ffff82116136bf57565b8054821015613ec7575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b91606093918352602083015260408201520190565b90939293848311611417578411611417578101920390565b908160c09103126114175760405190613f2082613c67565b80516001600160a01b038116810361141757825260208101516020830152604081015160408301526060810151600381101561141757613f799160a0916060850152613f6e60808201613e1e565b608085015201613e1e565b60a082015290565b905f51602061411f5f395f51905f52821080613fd2575b15613fb957811580613fc8575b613fb957613fb29161406b565b15613fb957565b634c4d29cd60e11b5f5260045ffd5b5060018114613fa5565b505f51602061411f5f395f51905f528110613f98565b91608093916040519384526020840152604083015260608201522090565b9291905f51602061411f5f395f51905f525f940691829060051b8201915b8281106140315750505050565b909192945f51602061411f5f395f51905f5283816020938186358b099008970993929101614024565b906040811015613ec75760051b0190565b5f51602061411f5f395f51905f528110801590614107575b614101575f51602061411f5f395f51905f528181920991800990805f51602061411f5f395f51905f5203915f51602061411f5f395f51905f5283116136bf575f51602061411f5f395f51905f528080838195097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b50505f90565b505f51602061411f5f395f51905f5282101561408356fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220c0c346fa520d2c61e5b67fcc584afc7fbb01d7835b1322ba8e9a119cc473936e64736f6c634300081c0033",
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

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x8e515b50.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerTransactor) SubmitCiphertext(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, pokAx *big.Int, pokAy *big.Int, pokZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitCiphertext", epochId, aid, c1x, c1y, c2x, c2y, pokAx, pokAy, pokZ)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x8e515b50.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, pokAx *big.Int, pokAy *big.Int, pokZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, c1x, c1y, c2x, c2y, pokAx, pokAy, pokZ)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0x8e515b50.
//
// Solidity: function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ) returns(uint16 ciphertextIndex)
func (_DKGManager *DKGManagerTransactorSession) SubmitCiphertext(epochId [12]byte, aid [32]byte, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, pokAx *big.Int, pokAy *big.Int, pokZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, aid, c1x, c1y, c2x, c2y, pokAx, pokAy, pokZ)
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
	PokAx           *big.Int
	PokAy           *big.Int
	PokZ            *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCiphertextSubmitted is a free log retrieval operation binding the contract event 0xf3ca88585465ab591836f08efb3f8993e60f41ef7f7d7b82bf35c776776432b2.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ)
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

// WatchCiphertextSubmitted is a free log subscription operation binding the contract event 0xf3ca88585465ab591836f08efb3f8993e60f41ef7f7d7b82bf35c776776432b2.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ)
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

// ParseCiphertextSubmitted is a log parse operation binding the contract event 0xf3ca88585465ab591836f08efb3f8993e60f41ef7f7d7b82bf35c776776432b2.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, address submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 pokAx, uint256 pokAy, uint256 pokZ)
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
