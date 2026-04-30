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

// DKGTypesAppPolicy is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesAppPolicy struct {
	AuthorizedSubmitter common.Address
	MaxCiphertexts      uint16
	NotBeforeBlock      uint64
	NotAfterBlock       uint64
}

// DKGTypesApplication is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesApplication struct {
	Creator        common.Address
	Mode           uint8
	DerivationS    *big.Int
	OrganizerPK    DKGTypesPoint
	Policy         DKGTypesAppPolicy
	CreatedAtBlock uint64
	Exists         bool
}

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
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	DeltaHash        [32]byte
	Delta            DKGTypesPoint
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delta\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerApplicationCoDec\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitment0X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment0Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitOrganizerShare\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dleqProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dleqInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochEvicted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochFinalized\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerShareSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x610160806040523461021b5760c0816161b88038038091610020828561021f565b83398101031261021b5780519063ffffffff82169182810361021b5761004860208301610256565b61005460408401610256565b9061006160608501610256565b9261007a60a061007360808801610256565b9601610256565b9563ffffffff46160361020c576001600160a01b038216156101fd576001600160a01b0383161580156101ec575b80156101db575b80156101ca575b6101bb5763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b6024820152601881526100f260388261021f565b5190201660c05260e052610100526101205261014052604051615f4d908161026b823960805181611310015260a05181818161026c01528181611ebb01528181612ac601528181614a680152615a4d015260c0518181816108790152612b7f015260e0518181816102c001528181610d85015261482201526101005181818161135901528181611b0d01528181613e77015261449d015261012051818181610dd90152818161143701526134050152610140518181816110320152818161218901526139330152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100b6565b506001600160a01b038516156100af565b506001600160a01b038416156100a8565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761024257604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361021b5756fe60806040526004361015610011575f80fd5b5f3560e01c806306433b1b14610244578063074a75e11461023f5780630b1451f01461023a57806317476f001461023557806318287e5f1461023057806323488be51461022b5780632c7c7642146102265780632fed2529146102215780633353ec6e1461021c578063373877a6146102175780634554c0be1461021257806349c61a121461020d578063510ba2df1461020857806363f314cd14610203578063669a76a9146101fe5780636759e0e1146101f957806370f2469b146101f457806372517b4b146101ef57806377235ee1146101ea57806385250700146101e557806385e1f4d0146101e05780638dc1f53a146101db57806391c1f1ee146101d657806393c3d3a8146101d1578063a4adcd7f146101cc578063a9c4b25f146101c7578063be59b8ea146101c2578063be8a39ad146101bd578063bf192209146101b8578063ca3c0458146101b3578063d3720aac146101ae578063d6c29c9e146101a9578063d9933767146101a4578063fe1604b51461019f5763fe2348971461019a575f80fd5b6121b8565b612174565b611dad565b611d10565b611c06565b611b7e565b611af8565b611a7c565b61199c565b61148b565b611466565b611422565b61139c565b611334565b6112f4565b61110e565b611075565b61100d565b610ee4565b610e1c565b610db4565b610d70565b610d1e565b610c8a565b610c0a565b610b63565b610b28565b610a74565b6108e1565b61085d565b610771565b61056a565b61034f565b61029b565b610257565b5f91031261025357565b5f80fd5b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f91610304575b50604051908152602090f35b610326915060203d60201161032c575b61031e8183612356565b810190612377565b5f6102f8565b503d610314565b612386565b600435906001600160a01b03198216820361025357565b3461025357602036600319011261025357610368610338565b6001600160a01b031981165f90815260226020526040902080546001600160a01b03161561054a5760058101805492600160ff85166103a6816118b3565b0361053b57600883015461ffff1691600184019384549361ffff6103d96103d28761ffff9060101c1690565b61ffff1690565b91161461053b576001600160401b03605085901c16954387101561053b5761046361045361043961041761046f946001600160401b039060481c1690565b9961043361042c60408b901c61ffff166103d2565b809c6123a5565b906123a5565b61044d6001600160401b0343169a8b6123ca565b996123ca565b9560901c6001600160401b031690565b6001600160401b031690565b6001600160401b038516101561052c577f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c946104d887610503945f6006899601559067ffffffffffffffff60481b82549160481b169067ffffffffffffffff60481b1916179055565b805467ffffffffffffffff60501b191660509290921b67ffffffffffffffff60501b16919091179055565b604080516001600160401b0395861681529290941660208301526001600160a01b0319169290a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b608090604319011261025357604490565b346102535761016036600319011261025357610584610338565b60243561059036610559565b60c4359060e4356101043590610124359261014435936105c4886001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b03161561054a576005015460039060ff166105e7816118b3565b0361053b57861561076257610621876106148a6001600160601b0360a01b165f52602d60205260405f2090565b905f5260205260405f2090565b936006850195610636875460ff9060401c1690565b610753576106639261065f92868a8c8e6106508484614cbe565b61065a8686614cbe565b614d43565b1590565b6107445782546001600160a01b031916331783557f5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a857237929361073f9361070e926106e991600490805460ff60a01b1916600160a01b1781555f60018201556106e36106ca6123ea565b8b8152602001889052600282018b905560038201889055565b0161243f565b8054600160401b68ffffffffffffffffff199091166001600160401b03431617179055565b60408051600181525f602082015290810194909452606084015233946001600160a01b031916929081906080820190565b0390a4005b6327f7eb4d60e11b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b6378e9323b60e11b5f5260045ffd5b346102535760203660031901126102535761078a610338565b6001600160a01b031981165f90815260226020526040902080546001600160a01b0316801561054a576001600160a01b0316330361084f57600501805460ff166107d3816118b3565b6003811490811561083a575b8115610826575b5061053b57805460ff191660041790556001600160a01b0319167f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b60049150610833816118b3565b145f6107e6565b9050610845816118b3565b60058114906107df565b6282b42960e81b5f5260045ffd5b34610253575f36600319011261025357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b61ffff81160361025357565b604435906108b68261089d565b565b602435906108b68261089d565b6001600160401b0381160361025357565b35906108b6826108c5565b34610253576101c0366003190112610253576004356108ff8161089d565b60243561090b8161089d565b604435916109188361089d565b6064356109248161089d565b6084356109308161089d565b60a43561093c816108c5565b60c43591610949836108c5565b60e43593610956856108c5565b60c0366101031901126102535761098f976109749761010497612a12565b6040516001600160a01b031990911681529081906020820190565b0390f35b634e487b7160e01b5f52602160045260245ffd5b600211156109b157565b610993565b81516001600160a01b0316815260208201516101608201939290919060028310156109b15760c0610140916108b694602085015260408101516040850152610a106060820151606086019060208091805184520151910152565b6001600160401b036060608083015160018060a01b0381511660a088015261ffff602082015116858801528260408201511660e0880152015116610100850152610a6b60a08201516101208601906001600160401b03169052565b01511515910152565b346102535760403660031901126102535761098f610b1c610b17610a96610338565b602435905f60c0604051610aa981612292565b828152826020820152826040820152610ac061304c565b6060820152604051610ad1816122b2565b83815283602082015283604082015283606082015260808201528260a082015201526001600160601b0360a01b165f52602d60205260405f20905f5260205260405f2090565b6130c8565b604051918291826109b6565b34610253576020366003190112610253576040610b4b610b46610338565b613153565b610b618251809260208091805184520151910152565bf35b34610253576040366003190112610253576020610bb5610b81610338565b60243590610b8e8261089d565b6001600160601b0360a01b165f52602c835260405f209061ffff165f5260205260405f2090565b54604051908152f35b6001600160401b0360a0809280511515855261ffff6020820151166020860152826040820151166040860152826060820151166060860152826080820151166080860152015116910152565b3461025357602036600319011261025357610c23610338565b610c2b6131a0565b506001600160601b0360a01b165f52602260205260c0610c50600360405f20016131d0565b610b616040518092610bbe565b9181601f84011215610253578235916001600160401b038311610253576020838186019501011161025357565b346102535760e036600319011261025357610ca3610338565b602435604435916064356084356001600160401b03811161025357610ccc903690600401610c5d565b60a4929192356001600160401b03811161025357610cee903690600401610c5d565b93909260c435976001600160401b03891161025357610d14610d1c993690600401610c5d565b98909761332c565b005b34610253576040366003190112610253576020610bb5610d3c610338565b60243590610d498261089d565b6001600160601b0360a01b165f526029835260405f209061ffff165f5260205260405f2090565b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b346102535760403660031901126102535760206001610e70610e3c610338565b60243590610e498261089d565b6001600160601b0360a01b165f526028845260405f209061ffff165f5260205260405f2090565b0154604051908152f35b6001600160a01b0381160361025357565b91909160c060a060e0830194600180831b03815116845261ffff602082015116602085015261ffff604082015116604085015260608101516060850152610a6b6080820151608086019060208091805184520151910152565b346102535760603660031901126102535761098f610f90610f03610338565b610f7b60243591610f1383610e7a565b60443590610f208261089d565b5f60a0604051610f2f816122cd565b828152826020820152826040820152826060820152610f4c61304c565b608082015201526001600160601b0360a01b165f52602660205260405f209061ffff165f5260205260405f2090565b9060018060a01b03165f5260205260405f2090565b611001610ff8600460405193610fa5856122cd565b80546001600160a01b038116865261ffff60a082901c81166020880152610fd59160b01c1661ffff166040870152565b60018101546060860152610feb60028201613064565b6080860152015460ff1690565b151560a0830152565b60405191829182610e8b565b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b34610253576101003660031901126102535761108f610338565b60243561109a6108a9565b9160843560643560a4356001600160401b038111610253576110c0903690600401610c5d565b9060c4356001600160401b038111610253576110e0903690600401610c5d565b94909360e435986001600160401b038a1161025357611106610d1c9a3690600401610c5d565b9990986137ef565b346102535760c036600319011261025357611127610338565b60243561113336610559565b6001600160a01b031983165f90815260226020526040902080546001600160a01b03161561054a576005015460039060ff1661116e816118b3565b0361053b5781156107625761119b82610614856001600160601b0360a01b165f52602d60205260405f2090565b90600682016111af815460ff9060401c1690565b610753576001600160a01b031985165f908152602b6020526040902091600183015490811561054a576004611259886106e99461123061129d98546112228c604051948593602085019788929091606c94926001600160601b0360a01b168452600c840152602c830152604c8201520190565b03601f198101835282612356565b5190207f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1900690565b86546001600160a81b0319163360ff60a01b191617875560018701819055956106e36112836123ea565b5f8082526001602090920182905260028401556003830155565b7f5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a857237926040518061073f33966001600160601b0360a01b169482606060019193929360808101945f825260208201525f60408201520152565b34610253575f36600319011261025357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b346102535760e0366003190112610253576113b5610338565b602435604435916113c58361089d565b6064356113d18161089d565b6084359060a4356001600160401b038111610253576113f4903690600401610c5d565b92909160c435966001600160401b0388116102535761141a610d1c983690600401610c5d565b979096613d4a565b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610253575f3660031901126102535760206001600160401b035f5416604051908152f35b346102535760c0366003190112610253576114a4610338565b6024356114b08161089d565b6044356064359160843560a435936114dc866001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b031695909390861561054a576003611501600587015460ff1690565b61150a816118b3565b0361053b5761ffff82169687158015611823575b6118015761152f61065f8589615720565b8015611810575b61180157611546600387016131d0565b906115518251151590565b90816117ed575b506117de5760408101516001600160401b031680151590816117c2575b506117a05761159161046360608301516001600160401b031690565b80151590816117af575b506117a0576115b761046360808301516001600160401b031690565b801515908161178d575b5061176b576115dd61046360a08301516001600160401b031690565b801515908161177a575b5061176b57602001516115fd9061ffff166103d2565b801515908161174a575b5061173b5761163f8261162e8a6001600160601b0360a01b165f52602c60205260405f2090565b9061ffff165f5260205260405f2090565b5461172c576116fa60087fa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf968a6116c661073f9661162e8c8b6116a88c6112228c6040519485936020850197889094939260609260808301968352602083015260408201520152565b519020936001600160601b0360a01b165f52602c60205260405f2090565b550180546116df9060301c61ffff1660010161ffff1690565b61ffff60301b82549160301b169061ffff60301b1916179055565b60405193849333996001600160601b0360a01b1697859094939260609260808301968352602083015260408201520152565b6316feb18560e11b5f5260045ffd5b63464e67af60e01b5f5260045ffd5b905061ffff611762600888015461ffff9060301c1690565b1610155f611607565b630410ff2960e31b5f5260045ffd5b90506001600160401b034216115f6115e7565b90506001600160401b034316115f6115c1565b633deac39560e01b5f5260045ffd5b90506001600160401b034216105f61159b565b6001600160401b031690506001600160401b034316105f611575565b6330cd747160e01b5f5260045ffd5b6001600160a01b031690503314155f611558565b634c4d29cd60e11b5f5260045ffd5b5061181e61065f8387615720565b611536565b50610100881161151e565b9060e0806108b69361ffff815116845261ffff602082015116602085015261ffff604082015116604085015261186f6060820151606086019061ffff169052565b60808181015161ffff169085015260a0818101516001600160401b03169085015260c0818101516001600160401b03169085015201516001600160401b0316910152565b600611156109b157565b9060068210156109b15752565b81516001600160a01b03168152610300810192916108b691906102e090610160906118fd6020820151602086019061182e565b6119106040820151610120860190610bbe565b61192360608201516101e08601906118bd565b60808101516001600160401b031661020085015260a08101516001600160401b031661022085015260c081015161024085015260e081015161026085015261010081015161ffff1661028085015261012081015161ffff166102a085015261014081015161ffff166102c0850152015161ffff16910152565b346102535760203660031901126102535761098f611a70611a6b6119be610338565b5f6101606040516119ce816122e8565b8281526040516119dd81612304565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e08201526020820152611a176131a0565b60408201528260608201528260808201528260a08201528260c08201528260e082015282610100820152826101208201528261014082015201526001600160601b0360a01b165f52602260205260405f2090565b61425b565b604051918291826118ca565b346102535760e036600319011261025357611a95610338565b60243560443591611aa58361089d565b60843560643560a4356001600160401b03811161025357611aca903690600401610c5d565b92909160c435966001600160401b03881161025357611af0610d1c983690600401610c5d565b97909661437c565b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b60206040818301928281528451809452019201905f5b818110611b5f5750505090565b82516001600160a01b0316845260209384019390920191600101611b52565b34610253576020366003190112610253576001600160a01b0319611ba0610338565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b818110611be75761098f85611bdb81870382612356565b60405191829182611b3c565b82546001600160a01b0316845260209093019260019283019201611bc4565b346102535760403660031901126102535761098f611c69611c25610338565b60243590611c3282610e7a565b611c3a6131a0565b506001600160a01b0319165f9081526025602090815260408083206001600160a01b0390941683529290522090565b611cbe610ff8600460405193611c7e856122cd565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b6040519182918291909160a08060c0830194600180831b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b346102535761012036600319011261025357611d2a610338565b611d326108b8565b6044359160843560643560a43560c4356001600160401b03811161025357611d5e903690600401610c5d565b9160e4356001600160401b03811161025357611d7e903690600401610c5d565b95909461010435996001600160401b038b1161025357611da5610d1c9b3690600401610c5d565b9a90996146ff565b3461025357602036600319011261025357611dc6610338565b6001600160a01b031981165f90815260226020526040902080549091906001600160a01b03161561054a5760058201805460018401805494909391611e1e61065f6001600160401b03605089901c1660ff841661594e565b61053b57600882019561ffff611e476103d2611e3c8a5461ffff1690565b9360101c61ffff1690565b9116101561216557611e8c611e85611e73856001600160601b0360a01b165f52602360205260405f2090565b335f9081526020919091526040902090565b5460ff1690565b612156576006820180549182156120ce575b50506040516313a4120960e31b815233600482015260a0816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333576001916060915f9161209f575b500151611f0581614c5f565b611f0e81614c5f565b036120905760408051602081019283523360601b6bffffffffffffffffffffffff19169181019190915260079190611f498160548101611222565b519020910154111561208157611fd7611f64855461ffff1690565b94611f8c33611f87856001600160601b0360a01b165f52602460205260405f2090565b614c69565b611fbe611fb133610f7b866001600160601b0360a01b165f52602360205260405f2090565b805460ff19166001179055565b611fc786614caa565b61ffff1661ffff19825416179055565b60405161ffff851681526001600160a01b0319821694612031916120259190339088907f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b90602090a3614caa565b935460101c61ffff1690565b9261ffff80851691161461204157005b61205b9261204e91615a19565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b6120c1915060a03d60a0116120c7575b6120b98183612356565b810190614bec565b5f611ef9565b503d6120af565b9091506120e69060481c6001600160401b0316610463565b804311156121475740908115612138578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f80611e9e565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102535760403660031901126102535761098f6122256121d7610338565b602435906121e48261089d565b5f604080516121f281612320565b82815282602082015201526001600160601b0360a01b165f52602860205260405f209061ffff165f5260205260405f2090565b60016040519161223483612320565b60ff815461ffff8116855260101c16151560208401520154604082015260405191829182919091604080606083019461ffff81511684526020810151151560208501520151910152565b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b038211176122ad57604052565b61227e565b608081019081106001600160401b038211176122ad57604052565b60c081019081106001600160401b038211176122ad57604052565b61018081019081106001600160401b038211176122ad57604052565b61010081019081106001600160401b038211176122ad57604052565b606081019081106001600160401b038211176122ad57604052565b604081019081106001600160401b038211176122ad57604052565b90601f801991011681019081106001600160401b038211176122ad57604052565b90816020910312610253575190565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b906001600160401b03809116911603906001600160401b0382116123c557565b612391565b906001600160401b03809116911601906001600160401b0382116123c557565b604051906108b6604083612356565b604051906108b661010083612356565b604051906108b661018083612356565b604051906108b6606083612356565b356124328161089d565b90565b35612432816108c5565b600160606108b693612471813561245581610e7a565b85546001600160a01b0319166001600160a01b03909116178555565b61249e60208201356124828161089d565b855461ffff60a01b191660a09190911b61ffff60a01b16178555565b60408101356124ac816108c5565b845467ffffffffffffffff60b01b191660b09190911b67ffffffffffffffff60b01b161784550135916124de836108c5565b01906001600160401b03166001600160401b0319825416179055565b9060068110156109b15760ff80198354169116179055565b908160209103126102535751612432816108c5565b908160051b91808304602014901517156123c557565b908160011b91808304600214901517156123c557565b818102929181159184041417156123c557565b8115612570570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160401b03166001600160401b0381146123c55760010190565b634e487b7160e01b5f52603260045260245ffd5b9060408210156125ce57600c600183811c810193160290565b6125a1565b91908260c0910312610253576040516125eb816122cd565b809280359081151582036102535760a061264f918193855260208101356126118161089d565b6020860152612622604082016108d6565b6040860152612633606082016108d6565b6060860152612644608082016108d6565b6080860152016108d6565b910152565b60068210156109b15752565b600161279b60e06108b69461268661ffff825116869061ffff1661ffff19825416179055565b6020810151855463ffff0000191660109190911b63ffff0000161785556040810151855465ffff00000000191660209190911b65ffff00000000161785556060810151855467ffff000000000000191660309190911b61ffff60301b161785556080810151855469ffff0000000000000000191660409190911b69ffff00000000000000001617855561274e61272660a08301516001600160401b031690565b865467ffffffffffffffff60501b191660509190911b67ffffffffffffffff60501b16178655565b61278d61276560c08301516001600160401b031690565b865467ffffffffffffffff60901b191660909190911b67ffffffffffffffff60901b16178655565b01516001600160401b031690565b9101906001600160401b03166001600160401b0319825416179055565b600161279b60a06108b6946127dc81511515869060ff801983541691151516179055565b6020810151855460408301516affffffffffffffffffff001990911660089290921b62ffff00169190911760189190911b6affffffffffffffff000000161785556060810151612857906001600160401b0316865467ffffffffffffffff60581b191660589190911b67ffffffffffffffff60581b16178655565b61278d61286e60808301516001600160401b031690565b865467ffffffffffffffff60981b191660989190911b67ffffffffffffffff60981b16178655565b815181546001600160a01b0319166001600160a01b039091161781556108b691906116df90610160906008906128d3602086015160018301612660565b6128e46040860151600383016127b8565b6129706005820161290260608801516128fc816118b3565b826124fa565b61293d61291960808901516001600160401b031690565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b60a0870151815470ffffffffffffffff000000000000000000191660489190911b67ffffffffffffffff60481b16179055565b60c0850151600682015560e0850151600782015501926129ab61299961010083015161ffff1690565b855461ffff191661ffff909116178555565b6129d86129be61012083015161ffff1690565b855463ffff0000191660109190911b63ffff000016178555565b612a096129eb61014083015161ffff1690565b855465ffff00000000191660209190911b65ffff0000000016178555565b015161ffff1690565b91929597949661ffff8316801590811561303f575b8115613031575b8115613022575b8115613015575b8115613002575b8115612ff4575b508015612fe5575b8015612fd9575b8015612fca575b8015612f9e575b8015612f82575b8015612f66575b61052c57604086016001600160401b03612a8e82612435565b1615159081612f4d575b81612f27575b508015612ebc575b8015612ea2575b612e9357604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610333576001600160401b03915f91612e64575b5016801561052c57612b2061ffff841661ffff8a16612553565b906127108210612e2d5750505f19975b5f546001600160401b0316612b4490612584565b612b62906001600160401b03166001600160401b03195f5416175f55565b5f546021546001600160a01b03196bffffffff00000000000000007f000000000000000000000000000000000000000000000000000000000000000060401b166001600160401b039384161760a01b169b8c939092909116906001600160401b038216603f16916001600160401b031660401115612da09b612c91612d9b9a612c86612cd69a612c23612cc69a612c03612cb69a612c9c99612def576125b5565b9091906001600160601b0383549160031b9260a01c831b921b1916179055565b612c66612c4a612c3b6021546001600160401b031690565b6001016001600160401b031690565b6001600160401b03166001600160401b03196021541617602155565b612c7b612c716123f9565b61ffff909e168e52565b61ffff1660208d0152565b61ffff1660408b0152565b61ffff166060890152565b61ffff891660808801526001600160401b031660a0870152565b6001600160401b031660c0850152565b6001600160401b031660e0830152565b612d50612cea5f546001600160401b031690565b94612d4061ffff6001600160401b034316961696612d24612d0b89896123ca565b93612d14612409565b33815296602088015236906125d3565b6040860152600160608601526001600160401b03166080850152565b6001600160401b031660a0830152565b5f60c08201528560e08201525f6101008201525f6101208201525f6101408201525f610160820152612d96876001600160601b0360a01b165f52602260205260405f2090565b612896565b6123ca565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917f2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d069190a390565b612e08612dfb826125b5565b90549060031b1c60a01b90565b6001600160a01b03198116612e1e575b506125b5565b612e2790614ef1565b5f612e18565b612e59612e5e927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43612553565b612566565b97612b30565b612e86915060203d602011612e8c575b612e7e8183612356565b810190612512565b5f612b06565b503d612e74565b63148b7e9360e31b5f5260045ffd5b5061010061ffff612eb560208901612428565b1611612aad565b5060608601612ecd61046382612435565b15159081612f0e575b81612ee2575b50612aa6565b90506001600160401b03612f04610463612efe60a08b01612435565b93612435565b911611155f612edc565b9050612f1f61046360a08901612435565b151590612ed6565b90506001600160401b03612f43610463612efe60808b01612435565b911611155f612a9e565b9050612f5e61046360808901612435565b151590612a98565b506001600160401b0389166001600160401b0382161115612a75565b506001600160401b0388166001600160401b038a161115612a6e565b50612fb961046361ffff87166001600160401b0343166123ca565b6001600160401b0389161115612a67565b5061010061ffff861611612a60565b5061ffff851615612a59565b5061271061ffff881610612a52565b905061ffff8516105f612a4a565b905061ffff831661ffff86161190612a43565b61ffff8616159150612a3c565b602061ffff8516119150612a35565b61ffff841681119150612a2e565b61ffff8416159150612a27565b604051906130598261233b565b5f6020838281520152565b906040516130718161233b565b602060018294805484520154910152565b9060405161308f816122b2565b60606001600160401b0360018395828154838060a01b038116875261ffff8160a01c16602088015260b01c166040860152015416910152565b90604051916130d683612292565b80546001600160a01b0381168452839060a01c60ff1660028110156109b15761314c60066108b69460c09360208601526001810154604086015261311c60028201613064565b606086015261312d60048201613082565b608086015201546001600160401b03811660a085015260401c60ff1690565b1515910152565b61315b61304c565b506001600160601b0360a01b165f52602b60205260405f206001810154156131865761243290613064565b506040516131938161233b565b5f81526001602082015290565b604051906131ad826122cd565b5f60a0838281528260208201528260408201528260608201528260808201520152565b906108b66040516131e0816122cd565b60a061325c6001839661324e61323e825460ff81161515885261ffff8160081c1660208901526001600160401b03808260181c161660408901526001600160401b03808260581c161660608901526001600160401b039060981c1690565b6001600160401b03166080870152565b01546001600160401b031690565b6001600160401b0316910152565b908060209392818452848401375f828201840152601f01601f1916010190565b92906132a390612432959360408652604086019161326a565b92602081850391015261326a565b90610120828203126102535780601f8301121561025357610120604051926132d98285612356565b8391810192831161025357905b8282106132f35750505090565b81358152602091820191016132e6565b90600182018092116123c557565b90608082018092116123c557565b919082018092116123c557565b9592949399989791969097613355876001600160601b0360a01b165f52602260205260405f2090565b80549092906001600160a01b03161561054a576005830191613378835460ff1690565b613381816118b3565b6003811461371557806133956002926118b3565b0361053b576133b161046360028601546001600160401b031690565b431061053b5760088401946133cc865461ffff9060101c1690565b966001860154936133e56103d28661ffff9060201c1690565b61ffff8a1610613706578d1580156136fe575b80156136f6575b6136e7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f92859261345b60405198899586948594635c73957b60e11b86526004860161328a565b03915afa92831561033357613478936136cd575b508101906132b1565b9485516134946134888b60a01c90565b6001600160601b031690565b14918215926136b1575b821561368e575b50811561367b575b50801561366d575b801561365f575b8015613651575b6136425760408051602081018b81529181018a9052606081018890526134fa91906134f18160808101611222565b519020886152bc565b8060e086015103613642576135106108a0612527565b860361364257613551958c61354694613539936135316103d2996101000190565b51918c6153e3565b805460ff19166003179055565b5460101c61ffff1690565b61356561355f610860612527565b8861331f565b965f5b8281106135ca575050507f4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a472939495506135c5906040519384936001600160601b0360a01b1696846040919493926060820195825260208201520152565b0390a2565b8060019160061b8a016040516135fc816112226020820194602081013590358660209093929193604081019481520152565b51902061363b613620886001600160601b0360a01b165f52602960205260405f2090565b61ffff8460051b8701351661ffff165f5260205260405f2090565b5501613568565b63d1fed5fd60e01b5f5260045ffd5b508560c085015114156134c3565b508760a085015114156134bc565b5088608085015114156134b5565b6060860151915061ffff1614155f6134ad565b9091506136a86103d2604088015b519260101c61ffff1690565b1415905f6134a5565b915060208601516136c561ffff84166103d2565b14159161349e565b806136db5f6136e193612356565b80610249565b5f61346f565b63c5f680ed60e01b5f5260045ffd5b508a156133ff565b508c156133f8565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b906101a0828203126102535780601f83011215610253576101a06040519261374c8285612356565b8391810192831161025357905b8282106137665750505090565b8135815260209182019101613759565b906080116102535790608090565b909291928311610253579190565b90939293848311610253578411610253578101920390565b9291926001600160401b0382116122ad57604051916137d3601f8201601f191660200184612356565b829481845281830111610253578281602093845f960137010152565b959993969298949099613816876001600160601b0360a01b165f52602260205260405f2090565b80549095906001600160a01b03161561054a576003613839600588015460ff1690565b613842816118b3565b0361053b5761ffff84169b8c158015613cb1575b8015613ca9575b613c9a576138838561162e8b6001600160601b0360a01b165f52602c60205260405f2090565b54918215613c8b5761ffff966138bc6138b48861162e8e6001600160601b0360a01b165f52602760205260405f2090565b5461ffff1690565b6138ce6103d260018c015461ffff1690565b9889911610613c6d5761390a9c6138fd8861162e8e6001600160601b0360a01b165f52602860205260405f2090565b9d8e5460ff9060101c1690565b613c7c578f8c935f925f948761391e6123ea565b5f8152600160208201529781613bac575b50507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b15610253578f909284935f9361398a60405196879586948594635c73957b60e11b86526004860161328a565b03915afa8015610333578f929e6134889f6139bb936139b193613b98575b50810190613724565b9d8e519260a01c90565b1494851595613b89575b508415613b7a575b508315613b66575b508215613b57575b508115613b47575b8115613b33575b508015613b25575b8015613b16575b8015613b07575b613642576101008701948551106136425760408051602081018d81529181018b9052613a3f9190613a368160608101611222565b51902089615323565b91826101608901510361364257613a566064612527565b810361364257613a69613a709185613776565b36916137aa565b602081519101200361364257613a9382606494613aa097613a989751918b615546565b6156cc565b916101800190565b51036136425781600184613ae07ff00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336966201000062ff000019825416179055565b01556040805194855260208501929092526001600160a01b0319169290819081015b0390a3565b50886101408801511415613a02565b508a61012088015114156139fb565b508460e088015114156139f4565b9050602060c089015191015114155f6139ec565b60a08901518151141591506139e5565b60808a0151141591505f6139dd565b60608b015160ff909116141592505f6139d5565b60408c0151141593508f6139cd565b60208d0151141594505f6139c5565b806136db5f613ba693612356565b5f6139a8565b6001600160a01b0319165f908152602d6020526040902091969550613bd091610614565b6006810154613be39060401c60ff161590565b610762578f948c89613bfa845460ff9060a01c1690565b97613c04896109a7565b60ff8916613c1e575050505060010154945b8f889061392f565b61162e9293999a50613c499450610614906001600160601b0360a01b165f52602e60205260405f2090565b613c5a61065f600383015460ff1690565b613c6d57613c6790613064565b95613c16565b63032cddf960e11b5f5260045ffd5b63955c0c4960e01b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508b1561385d565b506101008d11613856565b61ffff5f199116019061ffff82116123c557565b80548210156125ce575f5260205f2001905f90565b90610200828203126102535780601f830112156102535761020060405192613d0d8285612356565b8391810192831161025357905b828210613d275750505090565b8135815260209182019101613d1a565b61ffff1661ffff81146123c55760010190565b979091969592939495613d71896001600160601b0360a01b165f52602260205260405f2090565b80549097906001600160a01b03161561054a576003613d9460058a015460ff1690565b613d9d816118b3565b0361053b57613dc961065f611e85611e738d6001600160601b0360a01b165f52602360205260405f2090565b6141af5761ffff89169485158015614192575b8015614186575b8015614177575b801561416f575b61416157613e25613e168c6001600160601b0360a01b165f52602460205260405f2090565b613e1f8c613cbc565b90613cd0565b90543360039290921b1c6001600160a01b03160361364257613e716004613e698d610f7b8b61162e33936001600160601b0360a01b165f52602660205260405f2090565b015460ff1690565b614152577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f928592613ecd60405198899586948594635c73957b60e11b86526004860161328a565b03915afa92831561033357613eea9361413e575b50810190613ce5565b91613f0d8761162e8a6001600160601b0360a01b165f52602960205260405f2090565b54908351613f1e6134888b60a01c90565b149283159361412f575b50821561411c575b821561410c575b82156140fd575b5081156140f4575b81156140bb575b50613642576101208101805190916101400190611222613f8483516040519283916020830195869091604092825260208201520190565b5190208403613642577f39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f2294614058926003600893613fdf8b610f7b8961162e33936001600160601b0360a01b165f52602660205260405f2090565b805460a08c901b61ffff60a01b1663ffffffff60a01b199091161760b089901b61ffff60b01b1617815560048101805460ff19166001179055925160028401555191015501805461403b9060201c61ffff16613d37565b613d37565b65ffff0000000082549160201b169065ffff000000001916179055565b61408c61407d8261162e886001600160601b0360a01b165f52602760205260405f2090565b611fc7614036825461ffff1690565b6040805161ffff958616815294909116602085015283015233926001600160a01b031916918060608101613b02565b60e0830151610100840151604080516020810193845290810191909152919250906140e98160608101611222565b51902014155f613f4d565b80159150613f46565b6080840151141591505f613f3e565b6060840151600114159250613f37565b604084015161ffff861614159250613f30565b6020850151141592505f613f28565b806136db5f61414c93612356565b5f613ee1565b633466526160e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b508715613df1565b5061010061ffff881611613dea565b5061ffff871615613de3565b5060018901546141a89060101c61ffff166103d2565b8611613ddc565b63965c290d60e01b5f5260045ffd5b906108b66040516141ce81612304565b60e061325c6001839661324e61424b825461ffff8116885261ffff808260101c1616602089015261420d61ffff8260201c1660408a019061ffff169052565b61ffff603082901c16606089015261ffff604082901c1660808901526001600160401b03605082901c1660a089015260901c6001600160401b031690565b6001600160401b031660c0870152565b906108b6614338600861426c612409565b85546001600160a01b0316815294614286600182016141be565b6020870152614297600382016131d0565b60408701526142ef6142df60058301546142bd6142b48260ff1690565b60608b01612654565b6001600160401b03600882901c1660808a015260481c6001600160401b031690565b6001600160401b031660a0880152565b600681015460c0870152600781015460e0870152015461ffff811661010086015261ffff601082901c1661012086015261ffff602082901c1661014086015260301c61ffff1690565b61ffff16610160840152565b600360406108b693602081518051865501516001850155602081015160028501550151151591019060ff801983541691151516179055565b929597909694976143a1846001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b03161561054a576005015460039060ff166143c4816118b3565b0361053b5761ffff831698891580156146a2575b611801576143fe89610614876001600160601b0360a01b165f52602d60205260405f2090565b9161441461065f600685015460ff9060401c1690565b61076257600160ff61442b855460ff9060a01c1690565b614434816109a7565b16036107625761445c8561162e886001600160601b0360a01b165f52602c60205260405f2090565b5415613c8b5761448d6003613e698761162e8e6106148c6001600160601b0360a01b165f52602e60205260405f2090565b6141525761449b878a614cbe565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f928a926144f360405196879586948594635c73957b60e11b86526004860161328a565b03915afa80156103335761468e575b508461451083880188613ce5565b805161451f6134888860a01c90565b149283159361467f575b8315614670575b8315614660575b8315614652575b831561463e575b8315614628575b508215614618575b8215614608575b5050613642576145e66145eb926145a67f8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948976145956123ea565b948a865288602087015236916137aa565b602081519101206145b5612419565b93845260208401526001604084015261162e89610614876001600160601b0360a01b165f52602e60205260405f2090565b614344565b6040805194855260208501929092526001600160a01b03191692a4565b610180015114159050845f61455b565b6101608101518914159250614554565b610140820151600390910154141592505f61454c565b610120820151600282015414159350614545565b60808201511515935061453e565b6060820151600214159350614537565b60408201518c14159350614530565b60208201518b14159350614529565b806136db5f61469c93612356565b5f614502565b506101008a116143d8565b90610140828203126102535780601f8301121561025357610140604051926146d58285612356565b8391810192831161025357905b8282106146ef5750505090565b81358152602091820191016146e2565b9a99949693959198929790996147298c6001600160601b0360a01b165f52602260205260405f2090565b80549096906001600160a01b03161561054a57600587015460ff169261476b61065f60018a015495614765876001600160401b039060901c1690565b906157c3565b61053b5761479761065f8f611e73611e85916001600160601b0360a01b165f52602360205260405f2090565b6141af5761ffff8d169586158015614bd3575b614bc4576147db8f8f906147d5613e1f916001600160601b0360a01b165f52602460205260405f2090565b91613cbc565b90543360039290921b1c6001600160a01b0316036136425761481c8f613e69600491610f7b33916001600160601b0360a01b165f52602560205260405f2090565b614bb5577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f92859261487860405198899586948594635c73957b60e11b86526004860161328a565b03915afa9283156103335761489593614ba1575b508101906146ad565b928b6148a661348886519260a01c90565b1491821592614b85575b8215614b6a575b508115614b5b575b508015614b4d575b8015614b3f575b6136425760408051602081018a81529181018990526148fe91906148f58160608101611222565b5190208b615383565b8060c084015103613642578561010084015114801590614b30575b6136425761010061492981612527565b850361364257614959613a6961080096614947613a69898389613784565b60208151910120976114009187613792565b6020815191012061497e8d6001600160601b0360a01b165f52602a60205260405f2090565b54036136425761499892614991926156cc565b9160e00190565b510361364257614a26916149fa60046008936149cd8c610f7b33916001600160601b0360a01b165f52602560205260405f2090565b805461ffff60a01b191660a08d901b61ffff60a01b1617815590600382015501805460ff19166001179055565b018054614a0d9060101c61ffff16613d37565b63ffff000082549160101b169063ffff00001916179055565b614a63614a47876001600160601b0360a01b165f52602b60205260405f2090565b928354926001850193845480155f14614b2a57506001906157ea565b9255557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561025357604051633c1bcdef60e21b8152336004820152925f908490602490829084905af1928315610333577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb593614b16575b506040805161ffff9095168552602085019190915283015233926001600160a01b031916918060608101613b02565b806136db5f614b2493612356565b5f614ae7565b906157ea565b50866101208401511415614919565b508660a083015114156148ce565b5087608083015114156148c7565b9050606083015114155f6148bf565b909150614b7c6103d26040860161369c565b1415905f6148b7565b91506020840151614b9961ffff84166103d2565b1415916148b0565b806136db5f614baf93612356565b5f61488c565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b50614be5601086901c61ffff166103d2565b87116147aa565b908160a0910312610253576040519060a082018281106001600160401b038211176122ad576040528051614c1f81610e7a565b8252602081015160208301526040810151604083015260608101519060038210156102535760809160608401520151614c57816108c5565b608082015290565b600311156109b157565b8054600160401b8110156122ad57614c8691600182018155613cd0565b81546001600160a01b0393841660039290921b91821b9390911b1916919091179055565b61ffff60019116019061ffff82116123c557565b905f516020615ef85f395f51905f52821080614d2d575b15614d1e57811580614d14575b614d0557614cef91615bb4565b15614cf657565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b5060018114614ce2565b63d7c7beeb60e01b5f5260045ffd5b505f516020615ef85f395f51905f528110614cd5565b614dc190969593919492967f118620ccc82d2ca83b6188d7846e48b2fef422584cfaa5ee9287682269c30b2790614d78615964565b91825260a01c6020820152876040820152826060820152602073__$e3e8552c3b274c2b7820ea318ff5a26c3e$__91604051809481926321c2736360e21b835260048301615c73565b0381845af490811561033357614e2a966020935f93614e9e575b50614dfe90614de8615964565b9384525f516020615ef85f395f51905f52900690565b602083015284604083015285606083015260405180809881946321c2736360e21b835260048301615c73565b03915af4938415610333575f94614e71575b50614e5690614e4e614e5e9596615ca0565b979096615d6e565b9290916157ea565b91149182614e6b57505090565b14919050565b614e5e945090614e4e614e95614e569360203d60201161032c5761031e8183612356565b95505090614e3c565b614dfe919350614eba90853d871161032c5761031e8183612356565b9290614ddb565b8054905f815581614ed0575050565b5f5260205f20908101905b818110614ee6575050565b5f8155600101614edb565b6001600160a01b031981165f908152602260205260409020546001600160a01b0316156152b9576001600160a01b031981165f9081526024602052604081208054915b82811061514f5750505060015b61010061ffff821611156150315750614f76614f71826001600160601b0360a01b165f52602460205260405f2090565b614ec1565b6001600160a01b031981165f908152602a60209081526040808320839055602b9091529020614fab905b60015f918281550155565b615001614fcc826001600160601b0360a01b165f52602260205260405f2090565b60085f918281558260018201558260028201558260038201558260048201558260058201558260068201558260078201550155565b6001600160a01b0319167f457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c9965f80a2565b8061505d6103d26138b46150c29461162e876001600160601b0360a01b165f52602760205260405f2090565b61511a575b6150916150878261162e866001600160601b0360a01b165f52602860205260405f2090565b5460101c60ff1690565b6150f0575b6150b88161162e856001600160601b0360a01b165f52602c60205260405f2090565b546150c757613d37565b614f41565b5f6150ea8261162e866001600160601b0360a01b165f52602c60205260405f2090565b55613d37565b615115614fa08261162e866001600160601b0360a01b165f52602860205260405f2090565b615096565b61514a61513f8261162e866001600160601b0360a01b165f52602760205260405f2090565b805461ffff19169055565b615062565b61517161515c8284613cd0565b905460039190911b1c6001600160a01b031690565b6151a061519682610f7b886001600160601b0360a01b165f52602360205260405f2090565b805460ff19169055565b6001600160a01b031985165f9081526029602052604081206151d8906151c86103d286613303565b61ffff165f5260205260405f2090565b5561521b6151fe82610f7b886001600160601b0360a01b165f52602560205260405f2090565b60045f918281558260018201558260028201558260038201550155565b60015b61010061ffff82161115615236575050600101614f34565b806152656004613e6985610f7b61526e9661162e8d6001600160601b0360a01b165f52602660205260405f2090565b61527357613d37565b61521e565b61403661529c84610f7b8461162e8c6001600160601b0360a01b165f52602660205260405f2090565b60045f918281558260018201558260016002830182815501550155565b50565b5f516020615ef85f395f51905f52916040519060208201926001600160601b0360a01b1683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c815261531c606c82612356565b5190200690565b5f516020615ef85f395f51905f52916040519060208201926001600160601b0360a01b1683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c815261531c606c82612356565b5f516020615ef85f395f51905f52916040519060208201926001600160601b0360a01b1683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c815261531c606c82612356565b9091939294615403906104006153fc620100008261331f565b9186613792565b9161542f6103d260016154226103d2600889015461ffff9060101c1690565b96015460101c61ffff1690565b92610800915f5b868110615459575050505050505090615452916108a0916156cc565b0361364257565b61ffff8160051b890135168015801561553d575b613642576154be6154a161515c615498866001600160601b0360a01b165f52602460205260405f2090565b613e1f85613cbc565b6001600160a01b031985165f908152602560205260409020610f7b565b906154d061065f600484015460ff1690565b908115615520575b50613642576003615509613a696154ef8886612553565b615501896154fc88613303565b612553565b90888b613792565b602081519101209101540361364257600101615436565b90506155356103d2835461ffff9060a01c1690565b14155f6154d8565b5086811161546d565b90939291936155766103d2600161556961555f88613311565b976104809061331f565b97015460101c61ffff1690565b905f5b8481106155cc57505050505b6020811061559257505050565b8060061b83018160051b83013561364257803515908115916155bc575b5061364257600101615585565b600191506020013514155f6155af565b8060061b870161ffff8260051b8801351690811580156156c3575b6136425761563d61561b61515c615612896001600160601b0360a01b165f52602460205260405f2090565b613e1f86613cbc565b610f7b8661162e8a6001600160601b0360a01b165f52602660205260405f2090565b9161564f61065f600485015460ff1690565b9081156156a6575b5061364257815460b01c61ffff1661ffff80861691160361364257600282015481351491821592615691575b505061364257600101615579565b60209192506003015491013514155f80615683565b90506156bb6103d2845461ffff9060a01c1690565b14155f615657565b508482116155e7565b9291905f516020615ef85f395f51905f525f940691829060051b8201915b8281106156f75750505050565b909192945f516020615ef85f395f51905f5283816020938186358b0990089709939291016156ea565b5f516020615ef85f395f51905f5281108015906157ac575b6157a6575f516020615ef85f395f51905f528082819309928009818080808487097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09600108937f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000009081490565b50505f90565b505f516020615ef85f395f51905f52821015615738565b60068110156109b15760021490816157d9575090565b6001600160401b0391501643111590565b9392919091841580615944575b61593c57811580615932575b61592d575f516020615ef85f395f51905f52828609945f516020615ef85f395f51905f528285095f516020615ef85f395f51905f528188095f516020615ef85f395f51905f52907f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09965f516020615ef85f395f51905f52907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000096158a791615e4a565b935f516020615ef85f395f51905f52876001086158c390615e91565b935f516020615ef85f395f51905f529109915f516020615ef85f395f51905f529109905f516020615ef85f395f51905f529108905f516020615ef85f395f51905f5291099261591190615e02565b61591a90615e91565b5f516020615ef85f395f51905f52910990565b505090565b5060018114615803565b935090509190565b50600183146157f7565b60068110156109b15760011490816157d9575090565b60405190615973608083612356565b6080368337565b6040519061040061598b8184612356565b368337565b6040519061080061598b8184612356565b9060208110156125ce5760051b0190565b9060408110156125ce5760051b0190565b91905f835b60208210615a035750505061040082015f905b604082106159ed57505050610c000190565b60208060019285518152019301910190916159db565b60208060019285518152019301910190916159c8565b919091615a2461597a565b615a2c615990565b93615a4b836001600160601b0360a01b165f52602460205260405f2090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110615af65750505061ffff165b60208110615ad15750611222615ab0615ace9394956040519283916020830195866159c3565b519020916001600160601b0360a01b165f52602a60205260405f2090565b55565b806001615aef615ae9615ae4839561253d565b613303565b886159b2565b5201615a8a565b80615b03615b4692613303565b615b0d82886159a1565b5260a0615b1d61515c8387613cd0565b6040516313a4120960e31b81526001600160a01b03909116600482015292839081906024820190565b0381865afa918215610333576001926040915f91615b96575b506020810151615b77615b718561253d565b8d6159b2565b520151615b8f615b89615ae48461253d565b8b6159b2565b5201615a77565b615bae915060a03d81116120c7576120b98183612356565b5f615b5f565b5f516020615ef85f395f51905f528110801590615c5c575b6157a6575f516020615ef85f395f51905f52818192099180095f516020615ef85f395f51905f5282065f516020615ef85f395f51905f5203905f516020615ef85f395f51905f5282116123c5575f516020615ef85f395f51905f528080838180965f0896097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b505f516020615ef85f395f51905f52821015615bcc565b919060808301925f905b60048210615c8a57505050565b6020806001928551815201930191019091615c7d565b7f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f19006908115615d66575f6001927f1561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9f90807f25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b915b615d1f575050509190565b600180821614615d50575b60011c80615d39575b80615d14565b9181818493615d47936157ea565b91909250615d33565b9290828196615d5e936157ea565b949092615d2a565b5f9150600190565b7f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1909392919306908115615df7575f918060019592915b615db0575050509190565b600180821614615de1575b60011c80615dca575b80615da5565b9181818493615dd8936157ea565b91909250615dc4565b9290828196615def936157ea565b949092615dbb565b505090505f90600190565b5f516020615ef85f395f51905f5290065f516020615ef85f395f51905f52035f516020615ef85f395f51905f5281116123c5575f516020615ef85f395f51905f529060010890565b905f516020615ef85f395f51905f5290065f516020615ef85f395f51905f52035f516020615ef85f395f51905f5281116123c5575f516020615ef85f395f51905f52910890565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f516020615ef85f395f51905f5260a082015260208160c08160055afa1561025357519056fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212205f710ccd99c66ddb532bb3fc69844a65bbd8d3e2535eb04cdd3aeaea5a22fa4b64736f6c634300081c0033",
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

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
func (_DKGManager *DKGManagerCaller) GetApplication(opts *bind.CallOpts, epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getApplication", epochId, aid)

	if err != nil {
		return *new(DKGTypesApplication), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesApplication)).(*DKGTypesApplication)

	return out0, err

}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
func (_DKGManager *DKGManagerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGManager.Contract.GetApplication(&_DKGManager.CallOpts, epochId, aid)
}

// GetApplication is a free data retrieval call binding the contract method 0x2fed2529.
//
// Solidity: function getApplication(bytes12 epochId, bytes32 aid) view returns((address,uint8,uint256,(uint256,uint256),(address,uint16,uint64,uint64),uint64,bool))
func (_DKGManager *DKGManagerCallerSession) GetApplication(epochId [12]byte, aid [32]byte) (DKGTypesApplication, error) {
	return _DKGManager.Contract.GetApplication(&_DKGManager.CallOpts, epochId, aid)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 epochId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetCiphertextHash(opts *bind.CallOpts, epochId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCiphertextHash", epochId, ciphertextIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 epochId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetCiphertextHash(epochId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, epochId, ciphertextIndex)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 epochId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetCiphertextHash(epochId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, epochId, ciphertextIndex)
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

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCaller) GetCombinedDecryption(opts *bind.CallOpts, epochId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCombinedDecryption", epochId, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesCombinedDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesCombinedDecryptionRecord)).(*DKGTypesCombinedDecryptionRecord)

	return out0, err

}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerSession) GetCombinedDecryption(epochId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, epochId, ciphertextIndex)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 epochId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCallerSession) GetCombinedDecryption(epochId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, epochId, ciphertextIndex)
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

// GetPartialDecryption is a free data retrieval call binding the contract method 0x70f2469b.
//
// Solidity: function getPartialDecryption(bytes12 epochId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerCaller) GetPartialDecryption(opts *bind.CallOpts, epochId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPartialDecryption", epochId, participant, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesPartialDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesPartialDecryptionRecord)).(*DKGTypesPartialDecryptionRecord)

	return out0, err

}

// GetPartialDecryption is a free data retrieval call binding the contract method 0x70f2469b.
//
// Solidity: function getPartialDecryption(bytes12 epochId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerSession) GetPartialDecryption(epochId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, epochId, participant, ciphertextIndex)
}

// GetPartialDecryption is a free data retrieval call binding the contract method 0x70f2469b.
//
// Solidity: function getPartialDecryption(bytes12 epochId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerCallerSession) GetPartialDecryption(epochId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, epochId, participant, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 epochId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCaller) GetPlaintext(opts *bind.CallOpts, epochId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPlaintext", epochId, ciphertextIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 epochId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerSession) GetPlaintext(epochId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, epochId, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 epochId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCallerSession) GetPlaintext(epochId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, epochId, ciphertextIndex)
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

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGManager *DKGManagerTransactor) RegisterApplication(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "registerApplication", epochId, aid, policy)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGManager *DKGManagerSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.RegisterApplication(&_DKGManager.TransactOpts, epochId, aid, policy)
}

// RegisterApplication is a paid mutator transaction binding the contract method 0x85250700.
//
// Solidity: function registerApplication(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy) returns()
func (_DKGManager *DKGManagerTransactorSession) RegisterApplication(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.RegisterApplication(&_DKGManager.TransactOpts, epochId, aid, policy)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGManager *DKGManagerTransactor) RegisterApplicationCoDec(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "registerApplicationCoDec", epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGManager *DKGManagerSession) RegisterApplicationCoDec(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.RegisterApplicationCoDec(&_DKGManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// RegisterApplicationCoDec is a paid mutator transaction binding the contract method 0x17476f00.
//
// Solidity: function registerApplicationCoDec(bytes12 epochId, bytes32 aid, (address,uint16,uint64,uint64) policy, uint256 pkOrgX, uint256 pkOrgY, uint256 schnorrAx, uint256 schnorrAy, uint256 schnorrZ) returns()
func (_DKGManager *DKGManagerTransactorSession) RegisterApplicationCoDec(epochId [12]byte, aid [32]byte, policy DKGTypesAppPolicy, pkOrgX *big.Int, pkOrgY *big.Int, schnorrAx *big.Int, schnorrAy *big.Int, schnorrZ *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.RegisterApplicationCoDec(&_DKGManager.TransactOpts, epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 epochId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactor) SubmitCiphertext(opts *bind.TransactOpts, epochId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitCiphertext", epochId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 epochId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerSession) SubmitCiphertext(epochId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 epochId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitCiphertext(epochId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, epochId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitContribution(opts *bind.TransactOpts, epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitContribution", epochId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitContribution(epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, epochId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 epochId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitContribution(epochId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, epochId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0xbe8a39ad.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerTransactor) SubmitOrganizerShare(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitOrganizerShare", epochId, aid, ciphertextIndex, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0xbe8a39ad.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitOrganizerShare(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0xbe8a39ad.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitOrganizerShare(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x91c1f1ee.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitPartialDecryption(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitPartialDecryption", epochId, aid, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x91c1f1ee.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x91c1f1ee.
//
// Solidity: function submitPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitPartialDecryption(epochId [12]byte, aid [32]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, epochId, aid, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// DKGManagerApplicationRegisteredIterator is returned from FilterApplicationRegistered and is used to iterate over the raw logs and unpacked data for ApplicationRegistered events raised by the DKGManager contract.
type DKGManagerApplicationRegisteredIterator struct {
	Event *DKGManagerApplicationRegistered // Event containing the contract specifics and raw log

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
func (it *DKGManagerApplicationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerApplicationRegistered)
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
		it.Event = new(DKGManagerApplicationRegistered)
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
func (it *DKGManagerApplicationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerApplicationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerApplicationRegistered represents a ApplicationRegistered event raised by the DKGManager contract.
type DKGManagerApplicationRegistered struct {
	EpochId      [12]byte
	Aid          [32]byte
	Creator      common.Address
	Mode         uint8
	DerivationS  *big.Int
	OrganizerPKx *big.Int
	OrganizerPKy *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApplicationRegistered is a free log retrieval operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGManager *DKGManagerFilterer) FilterApplicationRegistered(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, creator []common.Address) (*DKGManagerApplicationRegisteredIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "ApplicationRegistered", epochIdRule, aidRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerApplicationRegisteredIterator{contract: _DKGManager.contract, event: "ApplicationRegistered", logs: logs, sub: sub}, nil
}

// WatchApplicationRegistered is a free log subscription operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGManager *DKGManagerFilterer) WatchApplicationRegistered(opts *bind.WatchOpts, sink chan<- *DKGManagerApplicationRegistered, epochId [][12]byte, aid [][32]byte, creator []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var aidRule []interface{}
	for _, aidItem := range aid {
		aidRule = append(aidRule, aidItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "ApplicationRegistered", epochIdRule, aidRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerApplicationRegistered)
				if err := _DKGManager.contract.UnpackLog(event, "ApplicationRegistered", log); err != nil {
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

// ParseApplicationRegistered is a log parse operation binding the contract event 0x5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792.
//
// Solidity: event ApplicationRegistered(bytes12 indexed epochId, bytes32 indexed aid, address indexed creator, uint8 mode, uint256 derivationS, uint256 organizerPKx, uint256 organizerPKy)
func (_DKGManager *DKGManagerFilterer) ParseApplicationRegistered(log types.Log) (*DKGManagerApplicationRegistered, error) {
	event := new(DKGManagerApplicationRegistered)
	if err := _DKGManager.contract.UnpackLog(event, "ApplicationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	CiphertextIndex uint16
	Submitter       common.Address
	C1x             *big.Int
	C1y             *big.Int
	C2x             *big.Int
	C2y             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCiphertextSubmitted is a free log retrieval operation binding the contract event 0xa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) FilterCiphertextSubmitted(opts *bind.FilterOpts, epochId [][12]byte, ciphertextIndex []uint16, submitter []common.Address) (*DKGManagerCiphertextSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "CiphertextSubmitted", epochIdRule, ciphertextIndexRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCiphertextSubmittedIterator{contract: _DKGManager.contract, event: "CiphertextSubmitted", logs: logs, sub: sub}, nil
}

// WatchCiphertextSubmitted is a free log subscription operation binding the contract event 0xa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) WatchCiphertextSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerCiphertextSubmitted, epochId [][12]byte, ciphertextIndex []uint16, submitter []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "CiphertextSubmitted", epochIdRule, ciphertextIndexRule, submitterRule)
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

// ParseCiphertextSubmitted is a log parse operation binding the contract event 0xa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed epochId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
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
	CiphertextIndex uint16
	CombineHash     [32]byte
	Plaintext       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecryptionCombined is a free log retrieval operation binding the contract event 0xf00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) FilterDecryptionCombined(opts *bind.FilterOpts, epochId [][12]byte, ciphertextIndex []uint16) (*DKGManagerDecryptionCombinedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "DecryptionCombined", epochIdRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerDecryptionCombinedIterator{contract: _DKGManager.contract, event: "DecryptionCombined", logs: logs, sub: sub}, nil
}

// WatchDecryptionCombined is a free log subscription operation binding the contract event 0xf00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) WatchDecryptionCombined(opts *bind.WatchOpts, sink chan<- *DKGManagerDecryptionCombined, epochId [][12]byte, ciphertextIndex []uint16) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "DecryptionCombined", epochIdRule, ciphertextIndexRule)
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

// ParseDecryptionCombined is a log parse operation binding the contract event 0xf00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336.
//
// Solidity: event DecryptionCombined(bytes12 indexed epochId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
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

// DKGManagerOrganizerShareSubmittedIterator is returned from FilterOrganizerShareSubmitted and is used to iterate over the raw logs and unpacked data for OrganizerShareSubmitted events raised by the DKGManager contract.
type DKGManagerOrganizerShareSubmittedIterator struct {
	Event *DKGManagerOrganizerShareSubmitted // Event containing the contract specifics and raw log

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
func (it *DKGManagerOrganizerShareSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerOrganizerShareSubmitted)
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
		it.Event = new(DKGManagerOrganizerShareSubmitted)
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
func (it *DKGManagerOrganizerShareSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerOrganizerShareSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerOrganizerShareSubmitted represents a OrganizerShareSubmitted event raised by the DKGManager contract.
type DKGManagerOrganizerShareSubmitted struct {
	EpochId         [12]byte
	Aid             [32]byte
	CiphertextIndex uint16
	DeltaOrgX       *big.Int
	DeltaOrgY       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrganizerShareSubmitted is a free log retrieval operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGManager *DKGManagerFilterer) FilterOrganizerShareSubmitted(opts *bind.FilterOpts, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (*DKGManagerOrganizerShareSubmittedIterator, error) {

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

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "OrganizerShareSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerOrganizerShareSubmittedIterator{contract: _DKGManager.contract, event: "OrganizerShareSubmitted", logs: logs, sub: sub}, nil
}

// WatchOrganizerShareSubmitted is a free log subscription operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGManager *DKGManagerFilterer) WatchOrganizerShareSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerOrganizerShareSubmitted, epochId [][12]byte, aid [][32]byte, ciphertextIndex []uint16) (event.Subscription, error) {

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

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "OrganizerShareSubmitted", epochIdRule, aidRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerOrganizerShareSubmitted)
				if err := _DKGManager.contract.UnpackLog(event, "OrganizerShareSubmitted", log); err != nil {
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

// ParseOrganizerShareSubmitted is a log parse operation binding the contract event 0x8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f0948.
//
// Solidity: event OrganizerShareSubmitted(bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, uint256 deltaOrgX, uint256 deltaOrgY)
func (_DKGManager *DKGManagerFilterer) ParseOrganizerShareSubmitted(log types.Log) (*DKGManagerOrganizerShareSubmitted, error) {
	event := new(DKGManagerOrganizerShareSubmitted)
	if err := _DKGManager.contract.UnpackLog(event, "OrganizerShareSubmitted", log); err != nil {
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
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	DeltaHash        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartialDecryptionSubmitted is a free log retrieval operation binding the contract event 0x39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
func (_DKGManager *DKGManagerFilterer) FilterPartialDecryptionSubmitted(opts *bind.FilterOpts, epochId [][12]byte, participant []common.Address) (*DKGManagerPartialDecryptionSubmittedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "PartialDecryptionSubmitted", epochIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerPartialDecryptionSubmittedIterator{contract: _DKGManager.contract, event: "PartialDecryptionSubmitted", logs: logs, sub: sub}, nil
}

// WatchPartialDecryptionSubmitted is a free log subscription operation binding the contract event 0x39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
func (_DKGManager *DKGManagerFilterer) WatchPartialDecryptionSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerPartialDecryptionSubmitted, epochId [][12]byte, participant []common.Address) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "PartialDecryptionSubmitted", epochIdRule, participantRule)
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

// ParsePartialDecryptionSubmitted is a log parse operation binding the contract event 0x39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed epochId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
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
