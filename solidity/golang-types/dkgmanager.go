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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEpoch\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"epochNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Application\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.AppMode\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"organizerPK\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"createdAtBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEpoch\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Epoch\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.EpochPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.EpochPhase\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerApplication\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerApplicationCoDec\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.AppPolicy\",\"components\":[{\"name\":\"authorizedSubmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCiphertexts\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pkOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pkOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrAy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schnorrZ\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitOrganizerShare\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dleqProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"dleqInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApplicationRegistered\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mode\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"derivationS\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKx\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"organizerPKy\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochAborted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochCreated\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochEvicted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EpochFinalized\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrganizerShareSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"deltaOrgX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaOrgY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deltaY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"epochId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApplicationAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApplication\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSchnorrProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsIdentity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCanonical\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOnCurve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x610160806040523461021b5760c0816168078038038091610020828561021f565b83398101031261021b5780519063ffffffff82169182810361021b5761004860208301610256565b61005460408401610256565b9061006160608501610256565b9261007a60a061007360808801610256565b9601610256565b9563ffffffff46160361020c576001600160a01b038216156101fd576001600160a01b0383161580156101ec575b80156101db575b80156101ca575b6101bb5763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b6024820152601881526100f260388261021f565b5190201660c05260e05261010052610120526101405260405161659c908161026b82396080518161177f015260a05181818161026c0152818161213d01528181612dd3015281816148ea0152615ba7015260c0518181816108ad0152612e8c015260e0518181816102c0015281816113e401526146f30152610100518181816117c801528181611e1201528181613b3601526141b80152610120518181816114380152818161182001526135d70152610140518181816114a0015281816124200152613eb40152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038616156100b6565b506001600160a01b038516156100af565b506001600160a01b038416156100a8565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b0382119082101761024257604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361021b5756fe60806040526004361015610011575f80fd5b5f3560e01c806306433b1b14610244578063074a75e11461023f5780630b1451f01461023a57806317476f001461023557806318287e5f1461023057806323488be51461022b5780632648f567146102265780632c7c7642146102215780632de546d51461021c5780632fed2529146102175780633353ec6e146102125780634554c0be1461020d57806349c61a12146102085780634ba849e714610203578063510ba2df146101fe5780635a8f2bb3146101f957806363f314cd146101f4578063669a76a9146101ef57806372517b4b146101ea57806377235ee1146101e557806385250700146101e057806385e1f4d0146101db5780638dc1f53a146101d657806393c3d3a8146101d15780639bbada67146101cc578063a305e0f3146101c7578063a4adcd7f146101c2578063b7bca615146101bd578063be59b8ea146101b8578063bea5210d146101b3578063bf192209146101ae578063ca3c0458146101a9578063d3720aac146101a4578063d99337671461019f5763fe1604b51461019a575f80fd5b61240b565b61202f565b611f1c565b611e83565b611dfd565b611ca8565b611bc8565b6119bf565b61199a565b611908565b61184f565b61180b565b6117a3565b611763565b61157c565b6114e3565b61147b565b611413565b6113cf565b61137f565b61132d565b6112a1565b61120d565b61118d565b611106565b610ff4565b610ec6565b610de0565b6108f9565b610891565b6107a5565b61056a565b61034f565b61029b565b610257565b5f91031261025357565b5f80fd5b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f91610304575b50604051908152602090f35b610326915060203d60201161032c575b61031e8183612527565b810190612548565b5f6102f8565b503d610314565b612557565b600435906001600160a01b03198216820361025357565b3461025357602036600319011261025357610368610338565b6001600160a01b031981165f90815260226020526040902080546001600160a01b03161561054a5760058101805492600160ff85166103a681611adf565b0361053b57600883015461ffff1691600184019384549361ffff6103d96103d28761ffff9060101c1690565b61ffff1690565b91161461053b576001600160401b03605085901c16954387101561053b5761046361045361043961041761046f946001600160401b039060481c1690565b9961043361042c60408b901c61ffff166103d2565b809c612576565b90612576565b61044d6001600160401b0343169a8b61259b565b9961259b565b9560901c6001600160401b031690565b6001600160401b031690565b6001600160401b038516101561052c577f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c946104d887610503945f6006899601559067ffffffffffffffff60481b82549160481b169067ffffffffffffffff60481b1916179055565b805467ffffffffffffffff60501b191660509290921b67ffffffffffffffff60501b16919091179055565b604080516001600160401b0395861681529290941660208301526001600160a01b0319169290a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63d5b25b6360e01b5f5260045ffd5b608090604319011261025357604490565b346102535761016036600319011261025357610584610338565b60243561059036610559565b6001600160a01b031983165f908152602260205260409020805460c435926101043592610124359260e435929161014435916001600160a01b03161561054a576005015460039060ff166105e381611adf565b0361053b5786156107965761061d876106108a6001600160601b0360a01b165f52602e60205260405f2090565b905f5260205260405f2090565b936006850195610632875460ff9060401c1690565b610787578782826106709561066661066c968f968f878d99918a926106578484614cae565b6106618686614cae565b615e00565b90615f33565b1590565b6107785782546001600160a01b031916331783557f5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a85723792936107739361071b926106f691600490805460ff60a01b1916600160a01b1781555f60018201556106f06106d76125bb565b8b8152602001889052600282018b905560038201889055565b0161260d565b8054600160401b68ffffffffffffffffff199091166001600160401b03431617179055565b6107428561073d886001600160601b0360a01b165f52602f60205260405f2090565b6126f6565b60408051600181525f602082015290810194909452606084015233946001600160a01b031916929081906080820190565b0390a4005b6327f7eb4d60e11b5f5260045ffd5b630b792c8f60e01b5f5260045ffd5b6378e9323b60e11b5f5260045ffd5b34610253576020366003190112610253576107be610338565b6001600160a01b031981165f90815260226020526040902080546001600160a01b0316801561054a576001600160a01b0316330361088357600501805460ff1661080781611adf565b6003811490811561086e575b811561085a575b5061053b57805460ff191660041790556001600160a01b0319167f379d6214174fba4ddb78deda3bc869bf16579e3ecef2dc0e55d6f688f66e44be5f80a2005b6004915061086781611adf565b145f61081a565b905061087981611adf565b6005811490610813565b6282b42960e81b5f5260045ffd5b34610253575f36600319011261025357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b61ffff81160361025357565b604435906108ea826108d1565b565b606435906108ea826108d1565b346102535760e036600319011261025357610912610338565b602435604435610921816108d1565b6064359060843560a4359460c4359261094e826001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b031697909490881561054a576003610973600588015460ff1690565b61097c81611adf565b0361053b5761ffff83169889158015610db9575b610daa5761099e8689614d33565b6109a88383614d33565b88610c93576109b960038801612788565b906109c48251151590565b9081610c7f575b50610c705760408101516001600160401b03168015159081610c54575b50610c3257610a0461046360608301516001600160401b031690565b8015159081610c41575b50610c3257610a2a61046360808301516001600160401b031690565b8015159081610c1f575b50610bfd57610a5061046360a08301516001600160401b031690565b8015159081610c0c575b50610bfd5760200151610a709061ffff166103d2565b8015159081610bdc575b50610bcd575b610ab783610aa68a610610886001600160601b0360a01b165f52602d60205260405f2090565b9061ffff165f5260205260405f2090565b54610bbe57610b8460087f1c8665e7b6ffd238f0d8ba92b2923fbcdc5eccb9dc9c138d5614eb279484ddfc978a610b5061077397610aa68d88610b2e8e610b208d6040519485936020850197889094939260609260808301968352602083015260408201520152565b03601f198101835282612527565b519020936106108c6001600160601b0360a01b165f52602d60205260405f2090565b55018054610b699060301c61ffff1660010161ffff1690565b61ffff60301b82549160301b169061ffff60301b1916179055565b604080513381526020810198909852870194909452606086019390935260808501929092526001600160a01b0319169290819060a0820190565b6316feb18560e11b5f5260045ffd5b63464e67af60e01b5f5260045ffd5b905061ffff610bf4600889015461ffff9060301c1690565b1610155f610a7a565b630410ff2960e31b5f5260045ffd5b90506001600160401b034216115f610a5a565b90506001600160401b034316115f610a34565b633deac39560e01b5f5260045ffd5b90506001600160401b034216105f610a0e565b6001600160401b031690506001600160401b034316105f6109e8565b6330cd747160e01b5f5260045ffd5b6001600160a01b031690503314155f6109cb565b50610cb688610610866001600160601b0360a01b165f52602e60205260405f2090565b6006810154610cc99060401c60ff161590565b610796576004610cd99101612742565b80516001600160a01b03168015159081610d9f575b50610c705760408101516001600160401b03168015159081610d83575b50610c3257610d2761046360608301516001600160401b031690565b8015159081610d70575b50610bfd5760200151610d479061ffff166103d2565b8015159081610d66575b5015610a805763464e67af60e01b5f5260045ffd5b905089115f610d51565b90506001600160401b034316115f610d31565b6001600160401b031690506001600160401b034316105f610d0b565b90503314155f610cee565b634c4d29cd60e11b5f5260045ffd5b506101008a11610990565b6001600160401b0381160361025357565b35906108ea82610dc4565b34610253576101c036600319011261025357600435610dfe816108d1565b602435610e0a816108d1565b60443591610e17836108d1565b606435610e23816108d1565b608435610e2f816108d1565b60a435610e3b81610dc4565b60c43591610e4883610dc4565b60e43593610e5585610dc4565b60c03661010319011261025357610e8e97610e739761010497612d1f565b6040516001600160a01b031990911681529081906020820190565b0390f35b6060906003190112610253576004356001600160a01b031981168103610253579060243590604435610ec3816108d1565b90565b34610253576020610f0a610ed936610e92565b916001600160601b0360a01b165f52602d845260405f20905f52835260405f209061ffff165f5260205260405f2090565b54604051908152f35b634e487b7160e01b5f52602160045260245ffd5b60021115610f3157565b610f13565b81516001600160a01b031681526020820151610160820193929091906002831015610f315760c0610140916108ea94602085015260408101516040850152610f906060820151606086019060208091805184520151910152565b6001600160401b036060608083015160018060a01b0381511660a088015261ffff602082015116858801528260408201511660e0880152015116610100850152610feb60a08201516101208601906001600160401b03169052565b01511515910152565b346102535760403660031901126102535761100d610338565b602435905f60c060405161102081612463565b828152826020820152826040820152611037613359565b6060820152611044613371565b60808201528260a082015201526001600160601b0360a01b165f52602e60205260405f20905f52602052610e8e60405f206110fa6110f160066040519361108a85612463565b80546001600160a01b03811686526110ab9060a01c60ff1660208701613395565b600181015460408601526110c1600282016133a1565b60608601526110d260048201612742565b608086015201546001600160401b03811660a085015260401c60ff1690565b151560c0830152565b60405191829182610f36565b34610253576020366003190112610253576040611129611124610338565b6133bf565b61113f8251809260208091805184520151910152565bf35b6001600160401b0360a0809280511515855261ffff6020820151166020860152826040820151166040860152826060820151166060860152826080820151166080860152015116910152565b34610253576020366003190112610253576111a6610338565b6111ae61340c565b506001600160601b0360a01b165f52602260205260c06111d3600360405f2001612788565b61113f6040518092611141565b9181601f84011215610253578235916001600160401b038311610253576020838186019501011161025357565b346102535760e036600319011261025357611226610338565b602435604435916064356084356001600160401b0381116102535761124f9036906004016111e0565b60a4929192356001600160401b038111610253576112719036906004016111e0565b93909260c435976001600160401b0389116102535761129761129f9936906004016111e0565b9890976134fe565b005b3461025357610160366003190112610253576112bb610338565b602435604435916112cb836108d1565b6084356064356101043560e43560c43560a435610124356001600160401b038111610253576112fe9036906004016111e0565b969095610144359a6001600160401b038c116102535761132561129f9c36906004016111e0565b9b909a6139c6565b34610253576040366003190112610253576020610f0a61134b610338565b60243590611358826108d1565b6001600160601b0360a01b165f52602a835260405f209061ffff165f5260205260405f2090565b3461025357602060016113c561139436610e92565b916001600160601b0360a01b165f526029855260405f20905f52845260405f209061ffff165f5260205260405f2090565b0154604051908152f35b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b3461025357610100366003190112610253576114fd610338565b6024356115086108dd565b9160843560643560a4356001600160401b0381116102535761152e9036906004016111e0565b9060c4356001600160401b0381116102535761154e9036906004016111e0565b94909360e435986001600160401b038a116102535761157461129f9a36906004016111e0565b999098613d86565b346102535760c036600319011261025357611595610338565b6024356115a136610559565b6001600160a01b031983165f90815260226020526040902080546001600160a01b03161561054a576005015460039060ff166115dc81611adf565b0361053b5781156107965761160982610610856001600160601b0360a01b165f52602e60205260405f2090565b906006820161161d815460ff9060401c1690565b610787576001600160a01b031985165f908152602c6020526040902091600183015490811561054a5760046116a6886106f6946116906116ea9854610b208c604051948593602085019788929091606c94926001600160601b0360a01b168452600c840152602c830152604c8201520190565b5190205f5160206165475f395f51905f52900690565b86546001600160a81b0319163360ff60a01b191617875560018701819055956106f06116d06125bb565b5f8082526001602090920182905260028401556003830155565b61170c8261073d856001600160601b0360a01b165f52602f60205260405f2090565b7f5c1bc55eb261d6ac466922a422fe62e9de8433120dc04979463fd16a857237926040518061077333966001600160601b0360a01b169482606060019193929360808101945f825260208201525f60408201520152565b34610253575f36600319011261025357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610253575f3660031901126102535760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333575f916103045750604051908152602090f35b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461025357610e8e6118af61186336610e92565b915f6040805161187281612483565b82815282602082015201526001600160601b0360a01b165f52602960205260405f20905f5260205260405f209061ffff165f5260205260405f2090565b6001604051916118be83612483565b60ff815461ffff8116855260101c16151560208401520154604082015260405191829182919091604080606083019461ffff81511684526020810151151560208501520151910152565b346102535761016036600319011261025357611922610338565b60243560443591611932836108d1565b61193a6108ec565b608435906101043560e43560c43560a435610124356001600160401b0381116102535761196b9036906004016111e0565b969095610144359a6001600160401b038c116102535761199261129f9c36906004016111e0565b9b909a614012565b34610253575f3660031901126102535760206001600160401b035f5416604051908152f35b346102535760e0366003190112610253576119d8610338565b6024356119e4816108d1565b604435916064356084356001600160401b03811161025357611a0a9036906004016111e0565b60a4929192356001600160401b03811161025357611a2c9036906004016111e0565b93909260c435976001600160401b03891161025357611a5261129f9936906004016111e0565b9890976145db565b9060e0806108ea9361ffff815116845261ffff602082015116602085015261ffff6040820151166040850152611a9b6060820151606086019061ffff169052565b60808181015161ffff169085015260a0818101516001600160401b03169085015260c0818101516001600160401b03169085015201516001600160401b0316910152565b60061115610f3157565b906006821015610f315752565b81516001600160a01b03168152610300810192916108ea91906102e09061016090611b2960208201516020860190611a5a565b611b3c6040820151610120860190611141565b611b4f60608201516101e0860190611ae9565b60808101516001600160401b031661020085015260a08101516001600160401b031661022085015260c081015161024085015260e081015161026085015261010081015161ffff1661028085015261012081015161ffff166102a085015261014081015161ffff166102c0850152015161ffff16910152565b3461025357602036600319011261025357610e8e611c9c611c97611bea610338565b5f610160604051611bfa8161249e565b828152604051611c09816124ba565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e08201526020820152611c4361340c565b60408201528260608201528260808201528260a08201528260c08201528260e082015282610100820152826101208201528261014082015201526001600160601b0360a01b165f52602260205260405f2090565b614af3565b60405191829182611af6565b3461025357608036600319011261025357610e8e611cc4610338565b611d99610aa660243592611d8b611d69604435611ce0816108d1565b610aa6600180611d3584611d2b8c6064359b8c91611cfd836108d1565b611d05613371565b506001600160601b0360a01b8d165f52602760205260405f20905f5260205260405f2090565b549061ffff161c90565b1614958615611df357610aa683995b8815611dd45761061083986001600160601b0360a01b165f52602660205260405f2090565b5491611d80611d766125ca565b61ffff9097168752565b61ffff166020860152565b604084015215156060830152565b60405191829182919091606080608083019461ffff815116845261ffff60208201511660208501526040810151604085015201511515910152565b6106105f986001600160601b0360a01b165f52602660205260405f2090565b610aa65f99611d44565b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b60206040818301928281528451809452019201905f5b818110611e645750505090565b82516001600160a01b0316845260209384019390920191600101611e57565b34610253576020366003190112610253576001600160a01b0319611ea5610338565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b818110611eec57610e8e85611ee081870382612527565b60405191829182611e41565b82546001600160a01b0316845260209093019260019283019201611ec9565b6001600160a01b0381160361025357565b3461025357604036600319011261025357610e8e611f7f611f3b610338565b60243590611f4882611f0b565b611f5061340c565b506001600160a01b0319165f9081526025602090815260408083206001600160a01b0390941683529290522090565b611fdd611fd4600460405193611f94856124d6565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b151560a0830152565b6040519182918291909160a08060c0830194600180831b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b3461025357602036600319011261025357612048610338565b6001600160a01b031981165f90815260226020526040902080549091906001600160a01b03161561054a57600582018054600184018054949093916120a061066c6001600160401b03605089901c1660ff8416615a97565b61053b57600882019561ffff6120c96103d26120be8a5461ffff1690565b9360101c61ffff1690565b911610156123fc5761210e6121076120f5856001600160601b0360a01b165f52602360205260405f2090565b335f9081526020919091526040902090565b5460ff1690565b6123ed57600682018054918215612365575b50506040516313a4120960e31b815233600482015260a0816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610333576001916060915f91612336575b50015161218781614c4f565b61219081614c4f565b036123275760408051602081019283523360601b6bffffffffffffffffffffffff191691810191909152600791906121cb8160548101610b20565b51902091015411156123185761226e6121e6855461ffff1690565b9461220e33612209856001600160601b0360a01b165f52602460205260405f2090565b614c59565b61225561224833612233866001600160601b0360a01b165f52602360205260405f2090565b9060018060a01b03165f5260205260405f2090565b805460ff19166001179055565b61225e86614c9a565b61ffff1661ffff19825416179055565b60405161ffff851681526001600160a01b03198216946122c8916122bc9190339088907f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b90602090a3614c9a565b935460101c61ffff1690565b9261ffff8085169116146122d857005b6122f2926122e591615b73565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b612358915060a03d60a01161235e575b6123508183612527565b810190614bdc565b5f61217b565b503d612346565b90915061237d9060481c6001600160401b0316610463565b804311156123de57409081156123cf578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f80612120565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b34610253575f366003190112610253576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b634e487b7160e01b5f52604160045260245ffd5b60e081019081106001600160401b0382111761247e57604052565b61244f565b606081019081106001600160401b0382111761247e57604052565b61018081019081106001600160401b0382111761247e57604052565b61010081019081106001600160401b0382111761247e57604052565b60c081019081106001600160401b0382111761247e57604052565b608081019081106001600160401b0382111761247e57604052565b604081019081106001600160401b0382111761247e57604052565b90601f801991011681019081106001600160401b0382111761247e57604052565b90816020910312610253575190565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601160045260245ffd5b906001600160401b03809116911603906001600160401b03821161259657565b612562565b906001600160401b03809116911601906001600160401b03821161259657565b604051906108ea604083612527565b604051906108ea608083612527565b604051906108ea61010083612527565b604051906108ea61018083612527565b35610ec3816108d1565b35610ec381610dc4565b600160606108ea9361263f813561262381611f0b565b85546001600160a01b0319166001600160a01b03909116178555565b61266c6020820135612650816108d1565b855461ffff60a01b191660a09190911b61ffff60a01b16178555565b604081013561267a81610dc4565b845467ffffffffffffffff60b01b191660b09190911b67ffffffffffffffff60b01b161784550135916126ac83610dc4565b01906001600160401b03166001600160401b0319825416179055565b634e487b7160e01b5f52603260045260245ffd5b80548210156126f1575f5260205f2001905f90565b6126c8565b8054600160401b81101561247e57612713916001820181556126dc565b819291549060031b91821b915f19901b1916179055565b906006811015610f315760ff80198354169116179055565b9060405161274f816124f1565b60606001600160401b0360018395828154838060a01b038116875261ffff8160a01c16602088015260b01c166040860152015416910152565b906108ea604051612798816124d6565b60a0612814600183966128066127f6825460ff81161515885261ffff8160081c1660208901526001600160401b03808260181c161660408901526001600160401b03808260581c161660608901526001600160401b039060981c1690565b6001600160401b03166080870152565b01546001600160401b031690565b6001600160401b0316910152565b908160209103126102535751610ec381610dc4565b908160051b918083046020149015171561259657565b908160011b918083046002149015171561259657565b908160021b918083046004149015171561259657565b8181029291811591840414171561259657565b8115612896570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160401b03166001600160401b0381146125965760010190565b9060408210156126f157600c600183811c810193160290565b91908260c0910312610253576040516128f8816124d6565b809280359081151582036102535760a061295c9181938552602081013561291e816108d1565b602086015261292f60408201610dd5565b604086015261294060608201610dd5565b606086015261295160808201610dd5565b608086015201610dd5565b910152565b6006821015610f315752565b6001612aa860e06108ea9461299361ffff825116869061ffff1661ffff19825416179055565b6020810151855463ffff0000191660109190911b63ffff0000161785556040810151855465ffff00000000191660209190911b65ffff00000000161785556060810151855467ffff000000000000191660309190911b61ffff60301b161785556080810151855469ffff0000000000000000191660409190911b69ffff000000000000000016178555612a5b612a3360a08301516001600160401b031690565b865467ffffffffffffffff60501b191660509190911b67ffffffffffffffff60501b16178655565b612a9a612a7260c08301516001600160401b031690565b865467ffffffffffffffff60901b191660909190911b67ffffffffffffffff60901b16178655565b01516001600160401b031690565b9101906001600160401b03166001600160401b0319825416179055565b6001612aa860a06108ea94612ae981511515869060ff801983541691151516179055565b6020810151855460408301516affffffffffffffffffff001990911660089290921b62ffff00169190911760189190911b6affffffffffffffff000000161785556060810151612b64906001600160401b0316865467ffffffffffffffff60581b191660589190911b67ffffffffffffffff60581b16178655565b612a9a612b7b60808301516001600160401b031690565b865467ffffffffffffffff60981b191660989190911b67ffffffffffffffff60981b16178655565b815181546001600160a01b0319166001600160a01b039091161781556108ea9190610b699061016090600890612be060208601516001830161296d565b612bf1604086015160038301612ac5565b612c7d60058201612c0f6060880151612c0981611adf565b8261272a565b612c4a612c2660808901516001600160401b031690565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b60a0870151815470ffffffffffffffff000000000000000000191660489190911b67ffffffffffffffff60481b16179055565b60c0850151600682015560e085015160078201550192612cb8612ca661010083015161ffff1690565b855461ffff191661ffff909116178555565b612ce5612ccb61012083015161ffff1690565b855463ffff0000191660109190911b63ffff000016178555565b612d16612cf861014083015161ffff1690565b855465ffff00000000191660209190911b65ffff0000000016178555565b015161ffff1690565b91929597949661ffff8316801590811561334c575b811561333e575b811561332f575b8115613322575b811561330f575b8115613301575b5080156132f2575b80156132e6575b80156132d7575b80156132ab575b801561328f575b8015613273575b61052c57604086016001600160401b03612d9b82612603565b161515908161325a575b81613234575b5080156131c9575b80156131af575b6131a057604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015610333576001600160401b03915f91613171575b5016801561052c57612e2d61ffff841661ffff8a16612879565b90612710821061313a5750505f19975b5f546001600160401b0316612e51906128aa565b612e6f906001600160401b03166001600160401b03195f5416175f55565b5f546021546001600160a01b03196bffffffff00000000000000007f000000000000000000000000000000000000000000000000000000000000000060401b166001600160401b039384161760a01b169b8c939092909116906001600160401b038216603f16916001600160401b0316604011156130ad9b612f9e6130a89a612f93612fe39a612f30612fd39a612f10612fc39a612fa9996130fc576128c7565b9091906001600160601b0383549160031b9260a01c831b921b1916179055565b612f73612f57612f486021546001600160401b031690565b6001016001600160401b031690565b6001600160401b03166001600160401b03196021541617602155565b612f88612f7e6125d9565b61ffff909e168e52565b61ffff1660208d0152565b61ffff1660408b0152565b61ffff166060890152565b61ffff891660808801526001600160401b031660a0870152565b6001600160401b031660c0850152565b6001600160401b031660e0830152565b61305d612ff75f546001600160401b031690565b9461304d61ffff6001600160401b034316961696613031613018898961259b565b936130216125e9565b33815296602088015236906128e0565b6040860152600160608601526001600160401b03166080850152565b6001600160401b031660a0830152565b5f60c08201528560e08201525f6101008201525f6101208201525f6101408201525f6101608201526130a3876001600160601b0360a01b165f52602260205260405f2090565b612ba3565b61259b565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917f2de0c6a525550f171bc2279b676488dc1184538e19b168b6219b0b7e45978d069190a390565b613115613108826128c7565b90549060031b1c60a01b90565b6001600160a01b0319811661312b575b506128c7565b61313490614dfe565b5f613125565b61316661316b927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43612879565b61288c565b97612e3d565b613193915060203d602011613199575b61318b8183612527565b810190612822565b5f612e13565b503d613181565b63148b7e9360e31b5f5260045ffd5b5061010061ffff6131c2602089016125f9565b1611612dba565b50606086016131da61046382612603565b1515908161321b575b816131ef575b50612db3565b90506001600160401b0361321161046361320b60a08b01612603565b93612603565b911611155f6131e9565b905061322c61046360a08901612603565b1515906131e3565b90506001600160401b0361325061046361320b60808b01612603565b911611155f612dab565b905061326b61046360808901612603565b151590612da5565b506001600160401b0389166001600160401b0382161115612d82565b506001600160401b0388166001600160401b038a161115612d7b565b506132c661046361ffff87166001600160401b03431661259b565b6001600160401b0389161115612d74565b5061010061ffff861611612d6d565b5061ffff851615612d66565b5061271061ffff881610612d5f565b905061ffff8516105f612d57565b905061ffff831661ffff86161190612d50565b61ffff8616159150612d49565b602061ffff8516119150612d42565b61ffff841681119150612d3b565b61ffff8416159150612d34565b604051906133668261250c565b5f6020838281520152565b6040519061337e826124f1565b5f6060838281528260208201528260408201520152565b6002821015610f315752565b906040516133ae8161250c565b602060018294805484520154910152565b6133c7613359565b506001600160601b0360a01b165f52602c60205260405f206001810154156133f257610ec3906133a1565b506040516133ff8161250c565b5f81526001602082015290565b60405190613419826124d6565b5f60a0838281528260208201528260408201528260608201528260808201520152565b908060209392818452848401375f828201840152601f01601f1916010190565b929061347590610ec3959360408652604086019161343c565b92602081850391015261343c565b90610120828203126102535780601f8301121561025357610120604051926134ab8285612527565b8391810192831161025357905b8282106134c55750505090565b81358152602091820191016134b8565b906001820180921161259657565b906080820180921161259657565b9190820180921161259657565b9592949399989791969097613527876001600160601b0360a01b165f52602260205260405f2090565b80549092906001600160a01b03161561054a57600583019161354a835460ff1690565b61355381611adf565b600381146139275780613567600292611adf565b0361053b5761358361046360028601546001600160401b031690565b431061053b57600884019461359e865461ffff9060101c1690565b966001860154936135b76103d28661ffff9060201c1690565b61ffff8a1610613918578d158015613910575b8015613908575b6138f9577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f92859261362d60405198899586948594635c73957b60e11b86526004860161345c565b03915afa9283156103335761364a936138df575b50810190613483565b94855161366661365a8b60a01c90565b6001600160601b031690565b14918215926138c3575b82156138a0575b50811561388d575b50801561387f575b8015613871575b8015613863575b6138545760408051602081018b81529181018a9052606081018890526136cc91906136c38160808101610b20565b51902088615341565b8060e086015103613854576136e26108a0612837565b860361385457613763958c6137589461370b936137036103d2996101000190565b51918c615468565b61371f613719610820612837565b8c6134f1565b600160206137418a6001600160601b0360a01b165f52602c60205260405f2090565b83358155920135910155805460ff19166003179055565b5460101c61ffff1690565b613777613771610860612837565b886134f1565b965f5b8281106137dc575050507f4626ec91a37d133f9027eadd556f820c54a05b0da238327825d5e5983696a472939495506137d7906040519384936001600160601b0360a01b1696846040919493926060820195825260208201520152565b0390a2565b8060019160061b8a0160405161380e81610b206020820194602081013590358660209093929193604081019481520152565b51902061384d613832886001600160601b0360a01b165f52602a60205260405f2090565b61ffff8460051b8701351661ffff165f5260205260405f2090565b550161377a565b63d1fed5fd60e01b5f5260045ffd5b508560c08501511415613695565b508760a0850151141561368e565b508860808501511415613687565b6060860151915061ffff1614155f61367f565b9091506138ba6103d2604088015b519260101c61ffff1690565b1415905f613677565b915060208601516138d761ffff84166103d2565b141591613670565b806138ed5f6138f393612527565b80610249565b5f613641565b63c5f680ed60e01b5f5260045ffd5b508a156135d1565b508c156135ca565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b90610200828203126102535780601f83011215610253576102006040519261395e8285612527565b8391810192831161025357905b8282106139785750505090565b813581526020918201910161396b565b9060108110156126f15760051b0190565b600260206108ea9381815180518655015160018501550151151591019060ff801983541691151516179055565b96999b9492959b9a909a9891986139f1886001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b03161561054a576005015460039060ff16613a1481611adf565b0361053b5761ffff8a169c8d158e8115613d7a575b50610daa57613a508d6106108b6001600160601b0360a01b165f52602e60205260405f2090565b94613a6661066c600688015460ff9060401c1690565b61079657600160ff613a7d885460ff9060a01c1690565b613a8681610f27565b160361079657613ab38c8f610aa6906106108e6001600160601b0360a01b165f52602d60205260405f2090565b54918215613d6b5760408051602081018c81529181018a90526060810193909352608083019190915290613aea8160a08101610b20565b5190200361385457613b266002613b1e8c8f610aa6906106108e6001600160601b0360a01b165f52603060205260405f2090565b015460ff1690565b613d5c57613b34898c614cae565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692833b156102535786935f928492613b8c6040519a8b9586948594635c73957b60e11b86526004860161345c565b03915afa918215610333578895613bab93613d48575b50810190613936565b918251613bbb61365a8860a01c90565b1494851595613d39575b8515613d2a575b8515613d1a575b8515613d0c575b8515613cfd575b508415613cee575b508315613cdb575b8315613cc5575b508215613cb5575b8215613ca5575b505061385457613c807f8b6045276e66f28a1293f2044b947b82818f03c318251187680b22778c8f094893613c7b613c3d6125bb565b91878352856020840152613c4f6125bb565b92835260016020840152610aa689610610876001600160601b0360a01b165f52603060205260405f2090565b613999565b6040805194855260208501929092526001600160a01b0319169290819081015b0390a4565b610140015114159050825f613c07565b6101208101518714159250613c00565b610100820151600390910154141592505f613bf8565b60e0820151600282015414159350613bf1565b60c0830151141593505f613be9565b60a0840151141594505f613be1565b608084015115159550613bda565b6060840151600214159550613bd3565b60408401518c14159550613bcc565b60208401518b14159550613bc5565b806138ed5f613d5693612527565b5f613ba2565b633466526160e01b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b6101009150118e613a29565b94999297919998909896939596613db1866001600160601b0360a01b165f52602260205260405f2090565b8054909b906001600160a01b03161561054a576003613dd460058e015460ff1690565b613ddd81611adf565b0361053b5761ffff81169b8c158015613fe0575b8015613fd8575b613fc957613e2282610aa68e6106108c6001600160601b0360a01b165f52602d60205260405f2090565b54968715613d6b57613e5c613e54848f610aa6906106108e6001600160601b0360a01b165f52602860205260405f2090565b5461ffff1690565b61ffff613e716103d2600186015461ffff1690565b911610613fba57613e9e83610aa68f6106108d6001600160601b0360a01b165f52602960205260405f2090565b95613eae875460ff9060101c1690565b613fab577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535788925f928892613f0a60405196879586948594635c73957b60e11b86526004860161345c565b03915afa978815610333577f4c7dcb60e78f05c6d275f7243d256bbbd80718bd70167d6625266614573e1ae19b858f998f98868f8f9b819f60019f613f629e613f5a9b8e93613f97575b50615633565b918c896158c7565b805462ff000019166201000017815501556040805194855260208501929092526001600160a01b031916929081908101613ca0565b806138ed5f613fa593612527565b5f613f54565b63955c0c4960e01b5f5260045ffd5b63032cddf960e11b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508a15613df8565b506101008d11613df1565b61ffff5f199116019061ffff821161259657565b61ffff1661ffff81146125965760010190565b9b91999497989a909a969593929661403e8d6001600160601b0360a01b165f52602260205260405f2090565b80549097906001600160a01b03161561054a57600361406160058a015460ff1690565b61406a81611adf565b0361053b5761409761066c8f6120f5612107916001600160601b0360a01b165f52602360205260405f2090565b6145015761ffff8c1695861580156144e4575b80156144d8575b80156144c9575b80156144c1575b6144b3576140f68f6140f08f916140ea906001600160601b0360a01b165f52602460205260405f2090565b91613feb565b906126dc565b90543360039290921b1c6001600160a01b031603613854576141368f8f610aa6906106108e936001600160601b0360a01b165f52602d60205260405f2090565b54918215613d6b5760408051602081018981529181018e9052606081019390935260808301919091529061416d8160a08101610b20565b5190200361385457610610958d6141b08a8f8f90614193610aa69261ffff600191161b90565b9b8c956001600160601b0360a01b165f52602760205260405f2090565b5416613d5c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f92859261420e60405198899586948594635c73957b60e11b86526004860161345c565b03915afa9283156103335761422a93613d485750810190613936565b9561424d89610aa68d6001600160601b0360a01b165f52602a60205260405f2090565b54918b61425f61365a8a519260a01c90565b14938415946144a4575b8415614491575b8415614481575b8415614472575b508315614463575b508215614454575b50811561444b575b8115614412575b50613854576101208401906142b58251956101400190565b94610b206142da87516040519283916020830195869091604092825260208201520190565b5190208603613854576008613ca09361439f9261436387610aa68d8f808f9e7f22adff6e28e87e60c01f5d89cee122b88fbe9a7eb000159cd38220075a22a3029f61434890610aa688610aa688610610610610996001600160601b0360a01b165f52602660205260405f2090565b556001600160601b0360a01b165f52602760205260405f2090565b9081541790550161438261437d825461ffff9060201c1690565b613fff565b65ffff0000000082549160201b169065ffff000000001916179055565b6143d76143c884610aa68a6106108d6001600160601b0360a01b165f52602860205260405f2090565b61225e61437d825461ffff1690565b51925160405193849333996001600160601b0360a01b1697859094939261ffff60609381608085019816845216602083015260408201520152565b60e0860151610100870151604080516020810193845290810191909152919250906144408160608101610b20565b51902014155f61429d565b80159150614296565b60c0870151141591505f61428e565b60a0880151141592505f614286565b6080890151141593505f61427e565b6060890151600114159450614277565b604089015161ffff891614159450614270565b60208901518c14159450614269565b62d949df60e51b5f5260045ffd5b508b156140bf565b5061010061ffff8b16116140b8565b5061ffff8a16156140b1565b5060018901546144fa9060101c61ffff166103d2565b87116140aa565b63965c290d60e01b5f5260045ffd5b90610100828203126102535780601f8301121561025357610100604051926145388285612527565b8391810192831161025357905b8282106145525750505090565b8135815260209182019101614545565b909291928311610253579190565b906080116102535790608090565b90939293848311610253578411610253578101920390565b9291926001600160401b03821161247e57604051916145bf601f8201601f191660200184612527565b829481845281830111610253578281602093845f960137010152565b989794939196929590976146038a6001600160601b0360a01b165f52602260205260405f2090565b80549096906001600160a01b03161561054a57600587015460ff169261464561066c60018a01549561463f876001600160401b039060901c1690565b90615a1c565b61053b5761467061066c6121076120f58f6001600160601b0360a01b165f52602360205260405f2090565b6145015761ffff8b169586158015614a3d575b614a2e576146ad8c6140f08f6140ea906001600160601b0360a01b165f52602460205260405f2090565b90543360039290921b1c6001600160a01b031603613854576146ed6004613b1e8f61223333916001600160601b0360a01b165f52602560205260405f2090565b614a1f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102535784925f92859261474960405198899586948594635c73957b60e11b86526004860161345c565b03915afa9283156103335761476693614a0b575b50810190614510565b92835161477661365a8c60a01c90565b14918215926149ef575b82156149d4575b5081156149c5575b5080156149b7575b80156149a9575b6138545760408051602081018881529181018790526147ce91906147c58160608101610b20565b519020896153a8565b8060c084015103613854576101006147e581612837565b85036138545761481c6148036108009661480a614803898389614562565b3691614596565b6020815191012097611400918761457e565b602081519101206148418b6001600160601b0360a01b165f52602b60205260405f2090565b54036138545761485b9261485492615a43565b9160e00190565b5103613854576148e8916148bc600460089361488f336122338c6001600160601b0360a01b165f52602560205260405f2090565b805461ffff60a01b191660a08b901b61ffff60a01b1617815590600382015501805460ff19166001179055565b0180546148cf9060101c61ffff16613fff565b63ffff000082549160101b169063ffff00001916179055565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561025357604051633c1bcdef60e21b8152336004820152925f908490602490829084905af1928315610333577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb593614995575b506040805161ffff959095168552602085019190915283015233926001600160a01b03191691606090a3565b806138ed5f6149a393612527565b5f614969565b508460a0830151141561479e565b508560808301511415614797565b9050606083015114155f61478f565b9091506149e66103d2604086016138ae565b1415905f614787565b91506020840151614a0361ffff84166103d2565b141591614780565b806138ed5f614a1993612527565b5f61475d565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b50614a4f601086901c61ffff166103d2565b8711614683565b906108ea604051614a66816124ba565b60e061281460018396612806614ae3825461ffff8116885261ffff808260101c16166020890152614aa561ffff8260201c1660408a019061ffff169052565b61ffff603082901c16606089015261ffff604082901c1660808901526001600160401b03605082901c1660a089015260901c6001600160401b031690565b6001600160401b031660c0870152565b906108ea614bd06008614b046125e9565b85546001600160a01b0316815294614b1e60018201614a56565b6020870152614b2f60038201612788565b6040870152614b87614b776005830154614b55614b4c8260ff1690565b60608b01612961565b6001600160401b03600882901c1660808a015260481c6001600160401b031690565b6001600160401b031660a0880152565b600681015460c0870152600781015460e0870152015461ffff811661010086015261ffff601082901c1661012086015261ffff602082901c1661014086015260301c61ffff1690565b61ffff16610160840152565b908160a0910312610253576040519060a082018281106001600160401b0382111761247e576040528051614c0f81611f0b565b8252602081015160208301526040810151604083015260608101519060038210156102535760809160608401520151614c4781610dc4565b608082015290565b60031115610f3157565b8054600160401b81101561247e57614c76916001820181556126dc565b81546001600160a01b0393841660039290921b91821b9390911b1916919091179055565b61ffff60019116019061ffff821161259657565b905f5160206165275f395f51905f52821080614d1d575b15614d0e57811580614d04575b614cf557614cdf91615d0e565b15614ce657565b6361586bdd60e01b5f5260045ffd5b632b39517d60e21b5f5260045ffd5b5060018114614cd2565b63d7c7beeb60e01b5f5260045ffd5b505f5160206165275f395f51905f528110614cc5565b905f5160206165275f395f51905f52821080614d84575b15610daa57811580614d7a575b610daa57614d658183615d0e565b15610daa57614d739161620d565b15610daa57565b5060018114614d57565b505f5160206165275f395f51905f528110614d4a565b5f1981019190821161259657565b5f5160206165275f395f51905f5203905f5160206165275f395f51905f52821161259657565b8054905f815581614ddd575050565b5f5260205f20908101905b818110614df3575050565b5f8155600101614de8565b6001600160a01b031981165f908152602260205260409020546001600160a01b03161561533e576001600160a01b031981165f9081526024602090815260408083208054602f909352922090919091614e5783546134d5565b915f5b82811061526b575050505f5b818110614f7257505050614e96614e91826001600160601b0360a01b165f52602f60205260405f2090565b614dce565b614eb7614e91826001600160601b0360a01b165f52602460205260405f2090565b6001600160a01b031981165f908152602b60209081526040808320839055602c9091529020614eec905b60015f918281550155565b614f42614f0d826001600160601b0360a01b165f52602260205260405f2090565b60085f918281558260018201558260028201558260038201558260048201558260058201558260068201558260078201550155565b6001600160a01b0319167f457d47cb94f548852cc20fd99e9450eecfcf65ea0e2547389681f2e4bab9c9965f80a2565b80158015615246575f905b60015b61010061ffff82161115614ff55750906001929115614fa1575b5001614e66565b614fc6614fef91610610886001600160601b0360a01b165f52602e60205260405f2090565b60065f918281558260018201558260016002830182815501558260048201558260058201550155565b5f614f9a565b828761501d83610aa684610610856001600160601b0360a01b165f52602760205260405f2090565b54615159575b6103d2613e5484610aa66150c09695610610615053966001600160601b0360a01b165f52602860205260405f2090565b615120575b61508b61508182610aa6876106108d6001600160601b0360a01b165f52602960205260405f2090565b5460101c60ff1690565b6150f2575b6150b681610aa6866106108c6001600160601b0360a01b165f52602d60205260405f2090565b546150c557613fff565b614f80565b5f6150ec82610aa6876106108d6001600160601b0360a01b165f52602d60205260405f2090565b55613fff565b61511b614ee182610aa6876106108d6001600160601b0360a01b165f52602960205260405f2090565b615090565b61515461514982610aa6876106108d6001600160601b0360a01b165f52602860205260405f2090565b805461ffff19169055565b615058565b61517f83610aa684610610856001600160601b0360a01b165f52602760205260405f2090565b5460015b602061ffff821611156151df5750506103d2613e5484610aa685610610865f6151cf86610aa66150c09d9c6106106150539d6001600160601b0360a01b165f52602760205260405f2090565b5596505095505094505050615023565b9092615209928286926001806151f9848a9061ffff161c90565b1614615212575b50505050613fff565b84918991615183565b5f93610aa6610aa69261061061523c966001600160601b0360a01b165f52602660205260405f2090565b55838a8288615200565b61526561525b61525584614d9a565b866126dc565b90549060031b1c90565b90614f7d565b8061533861531b615296615281600195876126dc565b905460039190911b1c6001600160a01b031690565b6152c56152bb826122338d6001600160601b0360a01b165f52602360205260405f2090565b805460ff19169055565b6001600160a01b03198a165f908152602a602052604081206152fd906152ed6103d2886134d5565b61ffff165f5260205260405f2090565b556001600160a01b03198a165f908152602560205260409020612233565b60045f918281558260018201558260028201558260038201550155565b01614e5a565b50565b5f5160206165275f395f51905f52916040519060208201926001600160601b0360a01b1683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c81526153a1606c82612527565b5190200690565b5f5160206165275f395f51905f52916040519060208201926001600160601b0360a01b1683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c81526153a1606c82612527565b5f5160206165275f395f51905f52916040519060208201926001600160601b0360a01b1683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c81526153a1606c82612527565b90919392946154889061040061548162010000826134f1565b918661457e565b916154b46103d260016154a76103d2600889015461ffff9060101c1690565b96015460101c61ffff1690565b925f5f92610800935b8781106154e1575050505050505050906154da916108a091615a43565b0361385457565b61ffff8160051b8a01351692831580156155d8575b61385457600161ffff85161b9081811661385457179261555961553c615281615533866001600160601b0360a01b165f52602460205260405f2090565b6140f085613feb565b6001600160a01b031985165f908152602560205260409020612233565b9061556b61066c600484015460ff1690565b9081156155bb575b506138545760036155a461480361558a8986612879565b61559c8a615597886134d5565b612879565b90898c61457e565b6020815191012091015403613854576001016154bd565b90506155d06103d2835461ffff9060a01c1690565b14155f615573565b508784116154f6565b906101a0828203126102535780601f83011215610253576101a0604051926156098285612527565b8391810192831161025357905b8282106156235750505090565b8135815260209182019101615616565b939499959897969091985f5f916156486125bb565b5f815260016020820152938561580b575b90615666918101906155e1565b9a8b518760a01c14948515956157fc575b5084156157e7575b5083156157d3575b5082156157c4575b5081156157b4575b81156157a0575b508015615781575b8015615772575b8015615763575b613854576101008701976156d26103d260018b5193015461ffff1690565b116138545760208851116138545760408051602081019283529081019390935261570c926157038160608101610b20565b51902090615408565b928361016086015103613854576157236064612837565b8103613854576148036157369184614570565b60208151910120036138545760646157519161575993615a43565b916101800190565b5103613854575190565b508261014088015114156156b4565b508061012088015114156156ad565b5060e08701516157996103d260018b015461ffff1690565b14156156a6565b9050602060c089015191015114155f61569e565b60a0890151815114159150615697565b60808a0151141591505f61568f565b60608b015160ff909116141592505f615687565b60408c015161ffff909116141593505f61567f565b60208d0151141594505f615677565b92915061583085610610896001600160601b0360a01b165f52602e60205260405f2090565b60068101546158439060401c60ff161590565b61079657805460a01c60ff169261585984610f27565b60ff84166158755750906001615666920154935b909150615659565b9394505061589f8c610aa6876106108b6001600160601b0360a01b165f52603060205260405f2090565b906158b161066c600284015460ff1690565b613fba576158c1615666926133a1565b9461586d565b90949392946158f76103d260016158ea6158e0896134e3565b98610480906134f1565b98015460101c61ffff1690565b905f935f5b868110615951575050505050505b6020811061591757505050565b8060061b83018160051b8301356138545780351590811591615941575b506138545760010161590a565b600191506020013514155f615934565b8060061b890161ffff8260051b8a0135169081158015615a13575b61385457600161ffff83161b97888116613854578817976159a985610aa6886106108c6001600160601b0360a01b165f52602760205260405f2090565b54161561385457615a05906040516159db81610b20602082019460208101359035869091604092825260208201520190565b51902091610aa685610aa6886106108c6001600160601b0360a01b165f52602660205260405f2090565b5403613854576001016158fc565b5085821161596c565b6006811015610f31576002149081615a32575090565b6001600160401b0391501643111590565b9291905f5160206165275f395f51905f525f940691829060051b8201915b828110615a6e5750505050565b909192945f5160206165275f395f51905f5283816020938186358b099008970993929101615a61565b6006811015610f31576001149081615a32575090565b60405190610400615abe8184612527565b368337565b60405190610800615abe8184612527565b60405190615ae3608083612527565b6080368337565b60405190610200615abe8184612527565b9060208110156126f15760051b0190565b9060408110156126f15760051b0190565b91905f835b60208210615b5d5750505061040082015f905b60408210615b4757505050610c000190565b6020806001928551815201930191019091615b35565b6020806001928551815201930191019091615b22565b919091615b7e615aad565b615b86615ac3565b93615ba5836001600160601b0360a01b165f52602460205260405f2090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110615c505750505061ffff165b60208110615c2b5750610b20615c0a615c28939495604051928391602083019586615b1d565b519020916001600160601b0360a01b165f52602b60205260405f2090565b55565b806001615c49615c43615c3e839561284d565b6134d5565b88615b0c565b5201615be4565b80615c5d615ca0926134d5565b615c678288615afb565b5260a0615c7761528183876126dc565b6040516313a4120960e31b81526001600160a01b03909116600482015292839081906024820190565b0381865afa918215610333576001926040915f91615cf0575b506020810151615cd1615ccb8561284d565b8d615b0c565b520151615ce9615ce3615c3e8461284d565b8b615b0c565b5201615bd1565b615d08915060a03d811161235e576123508183612527565b5f615cb9565b5f5160206165275f395f51905f528110801590615dbc575b615db6575f5160206165275f395f51905f52818192099180095f5160206165275f395f51905f5282065f5160206165275f395f51905f5203905f5160206165275f395f51905f528211612596575f5160206165275f395f51905f528080838180965f0896097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e0960010892081490565b50505f90565b505f5160206165275f395f51905f52821015615d26565b919060808301925f905b60048210615dea57505050565b6020806001928551815201930191019091615ddd565b9192615e7b94939591957f118620ccc82d2ca83b6188d7846e48b2fef422584cfaa5ee9287682269c30b2793615e34615ad4565b94855260a01c602085015260408401526060830152602073__$1de98e108939035880e23eaafb1b0ea3e6$__92604051809681926321c2736360e21b835260048301615dd3565b0381855af492831561033357615ee2956020955f95615f10575b50615eb890615ea2615ad4565b9586525f5160206165275f395f51905f52900690565b60208501526040840152606083015260405180809581946321c2736360e21b835260048301615dd3565b03915af4908115610333575f91615ef7575090565b610ec3915060203d60201161032c5761031e8183612527565b615eb8919550615f2c90873d891161032c5761031e8183612527565b9490615e95565b615f53615f67919692939495965f5160206165475f395f51905f52900690565b925f5160206165475f395f51905f52900690565b94806161ff57505f5b61609d615f7b615aea565b92615f84615aea565b60018152927f1561ff836ce19d358a4eb7a4c199e94c377c749ae6f2a277f1f9195afe553f9f60808601527f25797203f7a0b24925572e1cd16bf9edfce0051fb9e133774b3c257a872d7d8b60808501527f0daaa7e6b25c28e6dc8dd1d48e9cc61cd07015c1d7c1b8d4590eb6f51d5346dc6101008601527f01666cafbf0a30da8b9ebeaf848a1da067a892296f1043188e1705402b6d68536101008501527f136d609c4c856f5d277fab08c730cbdd1a776ce4728c6a2eb20ff22bccf268946101808601527f21d66f0e2295ae954494f25889f9319cc1b4df71eff3f46ba9e4631b43fd7c9561018085015280602086015281602085015261608982828181616330565b604086810191825287018290525190616330565b6060830152606083015260015b6004811061616757505f92600196607e805b6160d75750505050501491826160d157505090565b14919050565b5f19016160e38161284d565b600383600c86841c60021b16921c1617861580159061615c575b61613b575b8061610f575b50806160bc565b9586829a61612c616124616133959a8a613988565b519288613988565b5192616330565b989095616108565b958981616154939b61614c93616330565b908181616330565b989095616102565b5060018a14156160fd565b61617961617382612863565b84613988565b5161618661617383612863565b519060015b6004811061619e575050506001016160aa565b806161f86161c56161b16001948a613988565b516161bc848a613988565b51908787616330565b6161da846161d58a959495612863565b6134f1565b906161f16161eb866161d58c612863565b8b613988565b5289613988565b520161618b565b61620890614da8565b615f70565b90811580616297575b616290576001915f915f5160206165475f395f51905f52805b61624957505050159081616241575090565b600191501490565b60018082161461627a575b60011c80616263575b8061622f565b918181849361627193616330565b9190925061625d565b929082819561628893616330565b939092616254565b5050600190565b5060018114616216565b5f5160206165275f395f51905f5290065f5160206165275f395f51905f52035f5160206165275f395f51905f528111612596575f5160206165275f395f51905f529060010890565b905f5160206165275f395f51905f5290065f5160206165275f395f51905f52035f5160206165275f395f51905f528111612596575f5160206165275f395f51905f52910890565b91939290938215806164b6575b6164af578015806164a5575b61649f575f5160206165275f395f51905f52818409915f5160206165275f395f51905f528187095f5160206165275f395f51905f528185095f5160206165275f395f51905f52907f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09935f5160206165275f395f51905f52907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000096163ed916162e9565b935f5160206165275f395f51905f52846001089361640a906162a1565b965f5160206165275f395f51905f52888609616425906164c0565b97885f5160206165275f395f51905f529109935f5160206165275f395f51905f529109915f5160206165275f395f51905f529109905f5160206165275f395f51905f529108905f5160206165275f395f51905f529109935f5160206165275f395f51905f5291095f5160206165275f395f51905f52910990565b50509190565b5060018214616349565b9350919050565b506001851461633d565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f5160206165275f395f51905f5260a082015260208160c08160055afa1561025357519056fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1a26469706673582212205789505788fc251d2d2ab8d9d9120897ef3918c7558b769f0bf58e73fd883fdb64736f6c634300081c0033",
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

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerTransactor) SubmitOrganizerShare(opts *bind.TransactOpts, epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitOrganizerShare", epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitOrganizerShare(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
}

// SubmitOrganizerShare is a paid mutator transaction binding the contract method 0x4ba849e7.
//
// Solidity: function submitOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y, uint256 deltaOrgX, uint256 deltaOrgY, bytes dleqProof, bytes dleqInput) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitOrganizerShare(epochId [12]byte, aid [32]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int, deltaOrgX *big.Int, deltaOrgY *big.Int, dleqProof []byte, dleqInput []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitOrganizerShare(&_DKGManager.TransactOpts, epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)
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
