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

// DKGTypesRevealedShareRecord is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesRevealedShareRecord struct {
	Participant      common.Address
	ParticipantIndex uint16
	ShareValue       *big.Int
	ShareHash        [32]byte
	Accepted         bool
}

// DKGTypesRoundPolicy is an auto generated low-level Go binding around an user-defined struct.
type DKGTypesRoundPolicy struct {
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

// IDKGManagerRound is an auto generated low-level Go binding around an user-defined struct.
type IDKGManagerRound struct {
	Organizer              common.Address
	Policy                 DKGTypesRoundPolicy
	DecryptionPolicy       DKGTypesDecryptionPolicy
	Status                 uint8
	Nonce                  uint64
	SeedBlock              uint64
	Seed                   [32]byte
	LotteryThreshold       *big.Int
	ClaimedCount           uint16
	ContributionCount      uint16
	PartialDecryptionCount uint16
	RevealedShareCount     uint16
	CiphertextCount        uint16
}

// DKGManagerMetaData contains all meta data concerning the DKGManager contract.
var DKGManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealSubmitVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealShareVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SHARE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SUBMIT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROUND_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRound\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delta\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealShareVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealSubmitVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RevealedShareRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Round\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RoundPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizeNotBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.RoundStatus\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"revealedShareCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reconstructSecret\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roundNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitment0X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment0Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevealedShareSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundAborted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundCreated\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundEvicted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundFinalized\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecretReconstructed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisclosureDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRevealedShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReconstruction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRevealedShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShareCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LagrangeMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6101a0806040523461029657610100816161848038038091610021828561029a565b8339810103126102965780519063ffffffff82169182810361029657610049602083016102d1565b610055604084016102d1565b610061606085016102d1565b9061006e608086016102d1565b9261007b60a087016102d1565b9461009460e061008d60c08a016102d1565b98016102d1565b9763ffffffff461603610287576001600160a01b03821615610278576001600160a01b038316158015610267575b8015610256575b8015610245575b8015610234575b8015610223575b6102145763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261011a60388261029a565b5190201660c05260e0526101005261012052610140526101605261018052604051615e9e90816102e6823960805181611483015260a051818181610374015281816121ce015281816138540152818161468d0152615b11015260c051818181610bfd015261390a015260e05181818161040c01528181610c7a015261444701526101005181818161112a015281816114cc0152611a02015261012051818181610cce015281816115240152612b3e015261014051818181610f1f0152818161271e01526140790152610160518181816103b80152818161190f0152611c1a015261018051818181610c3601528181611a560152612f8f0152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038816156100de565b506001600160a01b038716156100d7565b506001600160a01b038616156100d0565b506001600160a01b038516156100c9565b506001600160a01b038416156100c2565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b038211908210176102bd57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102965756fe60806040526004361015610011575f80fd5b5f3560e01c8063058994a11461027457806306433b1b1461026f578063070c74921461026a578063074a75e1146102655780630b1451f0146102605780630e2c53f71461025b5780633353ec6e14610256578063349181a214610251578063373877a61461024c5780633caf448714610247578063415a1b86146102425780634554c0be1461023d578063510ba2df1461023857806353d721841461023357806356664d151461022e5780635ddd06261461022957806363f314cd14610224578063669a76a91461021f5780636759e0e11461021a57806370f2469b1461021557806372517b4b14610210578063802ae2311461020b57806385e1f4d0146102065780638dc1f53a1461020157806393c3d3a8146101fc578063a9c4b25f146101f7578063b18730c2146101f2578063b58aab90146101ed578063bf192209146101e8578063c2440e16146101e3578063c9396bf0146101de578063ca3c0458146101d9578063d3720aac146101d4578063d6c29c9e146101cf578063d9933767146101ca578063f4e34945146101c5578063fe1604b5146101c05763fe234897146101bb575f80fd5b61274d565b612709565b61261b565b6120d9565b61203c565b611f32565b611eaa565b611a99565b611a31565b6119ed565b611952565b6118ea565b611553565b61150f565b6114a7565b611467565b610f62565b610efa565b610dd1565b610d11565b610ca9565b610c65565b610c21565b610be1565b610ad2565b610a6f565b610a1c565b6109ab565b6108f9565b61085d565b610756565b61071b565b61068e565b610484565b6103e7565b6103a3565b61035f565b6102c1565b600435906001600160a01b03198216820361029057565b5f80fd5b9181601f84011215610290578235916001600160401b038311610290576020838186019501011161029057565b346102905760e0366003190112610290576102da610279565b602435604435916064356084356001600160401b03811161029057610303903690600401610294565b60a4929192356001600160401b03811161029057610325903690600401610294565b93909260c435976001600160401b0389116102905761034b610353993690600401610294565b989097612a65565b005b5f91031261029057565b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f91610450575b50604051908152602090f35b610472915060203d602011610478575b61046a81836128d0565b810190612e3c565b5f610444565b503d610460565b61293b565b346102905760203660031901126102905761049d610279565b6001600160a01b031981165f90815260226020526040902080546001600160a01b03161561067f5760058101805492600160ff85166104db8161251d565b0361067057600883015461ffff1691600184019384549361ffff61050e6105078761ffff9060101c1690565b61ffff1690565b911614610670576001600160401b03605085901c1695438710156106705761059861058861056e61054c6105a4946001600160401b039060481c1690565b9961056861056160408b901c61ffff16610507565b809c612e4b565b90612e4b565b6105826001600160401b0343169a8b612e6b565b99612e6b565b9560901c6001600160401b031690565b6001600160401b031690565b6001600160401b0385161015610661577f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c9461060d87610638945f6006899601559067ffffffffffffffff60481b82549160481b169067ffffffffffffffff60481b1916179055565b805467ffffffffffffffff60501b191660509290921b67ffffffffffffffff60501b16919091179055565b604080516001600160401b0395861681529290941660208301526001600160a01b0319169290a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b6328ad4a9560e21b5f5260045ffd5b346102905760c0366003190112610290576106a7610279565b602435604435916064356001600160401b038111610290576106cd903690600401610294565b906084356001600160401b038111610290576106ed903690600401610294565b92909160a435966001600160401b03881161029057610713610353983690600401610294565b979096612edc565b3461029057602036600319011261029057604061073e610739610279565b6131b3565b6107548251809260208091805184520151910152565bf35b346102905760203660031901126102905761076f610279565b6001600160a01b031981165f90815260226020526040902080546001600160a01b0316801561067f576001600160a01b0316330361083457600501805460ff166107b88161251d565b6003811490811561081f575b811561080b575b5061067057805460ff191660041790556001600160a01b0319167f97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e4105f80a2005b600491506108188161251d565b145f6107cb565b905061082a8161251d565b60058114906107c4565b6282b42960e81b5f5260045ffd5b61ffff81160361029057565b6024359061085b82610842565b565b346102905760403660031901126102905760206108af61087b610279565b6024359061088882610842565b6001600160601b0360a01b165f52602d835260405f209061ffff165f5260205260405f2090565b54604051908152f35b6001600160401b0381160361029057565b6101043590811515820361029057565b3590811515820361029057565b60c0906101231901126102905761012490565b34610290576101e0366003190112610290576109a761098c60043561091d81610842565b60243561092981610842565b60443561093581610842565b60643561094181610842565b60843561094d81610842565b60a4359061095a826108b8565b60c43592610967846108b8565b60e43594610974866108b8565b61097c6108c9565b96610986366108e6565b986137ab565b6040516001600160a01b031990911681529081906020820190565b0390f35b34610290575f3660031901126102905760206001600160401b035f5416604051908152f35b6001600160401b0360a0809280511515855261ffff6020820151166020860152826040820151166040860152826060820151166060860152826080820151166080860152015116910152565b3461029057602036600319011261029057610a35610279565b610a3d613d8b565b506001600160601b0360a01b165f52602260205260c0610a62600360405f2001613dbb565b61075460405180926109d0565b346102905760403660031901126102905760206108af610a8d610279565b60243590610a9a82610842565b6001600160601b0360a01b165f52602a835260405f209061ffff165f5260205260405f2090565b6001600160a01b0381160361029057565b34610290576040366003190112610290576109a7610b50610af1610279565b60243590610afe82610ac1565b5f6080604051610b0d81612827565b82815282602082015282604082015282606082015201526001600160601b0360a01b165f52602960205260405f209060018060a01b03165f5260205260405f2090565b60ff600360405192610b6184612827565b61ffff815460018060a01b038116865260a01c1660208501526001810154604085015260028101546060850152015416151560808201526040519182918291909160808060a083019460018060a01b03815116845261ffff6020820151166020850152604081015160408501526060810151606085015201511515910152565b34610290575f36600319011261029057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f916104505750604051908152602090f35b346102905760403660031901126102905760206001610d65610d31610279565b60243590610d3e82610842565b6001600160601b0360a01b165f526028845260405f209061ffff165f5260205260405f2090565b0154604051908152f35b91909160c060a060e0830194600180831b03815116845261ffff602082015116602085015261ffff604082015116604085015260608101516060850152610dc86080820151608086019060208091805184520151910152565b01511515910152565b34610290576060366003190112610290576109a7610e7d610df0610279565b610e6860243591610e0083610ac1565b60443590610e0d82610842565b5f60a0604051610e1c81612847565b828152826020820152826040820152826060820152610e3961317d565b608082015201526001600160601b0360a01b165f52602660205260405f209061ffff165f5260205260405f2090565b9060018060a01b03165f5260205260405f2090565b610eee610ee5600460405193610e9285612847565b80546001600160a01b038116865261ffff60a082901c81166020880152610ec29160b01c1661ffff166040870152565b60018101546060860152610ed860028201613195565b6080860152015460ff1690565b151560a0830152565b60405191829182610d6f565b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f916104505750604051908152602090f35b346102905760c036600319011261029057610f7b610279565b602435610f8781610842565b604435610f9381610842565b6064356084356001600160401b03811161029057610fb5903690600401610294565b60a4939193356001600160401b03811161029057610fd7903690600401610294565b90610ff6886001600160601b0360a01b165f52602260205260405f2090565b80549096906001600160a01b03161561067f576003611019600589015460ff1690565b6110228161251d565b036106705761106b61106761106061104e8c6001600160601b0360a01b165f52602360205260405f2090565b335f9081526020919091526040902090565b5460ff1690565b1590565b6114585761ffff8816938415801561143b575b801561142f575b8015611420575b8015611418575b61140a576110c76110b88b6001600160601b0360a01b165f52602460205260405f2090565b6110c18b613e55565b90613e69565b90543360039290921b1c6001600160a01b03160361138057611124600461111c8c610e688a61110b33936001600160601b0360a01b165f52602660205260405f2090565b9061ffff165f5260205260405f2090565b015460ff1690565b6113fb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f92859261118060405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f5761119d936113e1575b50810190613e7e565b906111c08661110b896001600160601b0360a01b165f52602a60205260405f2090565b5482516111dc6111d08a60a01c90565b6001600160601b031690565b14918215926113d2575b5081156113c9575b811561138f575b506113805760c081018051909160e0019061122a61123883516040519283916020830195869091604092825260208201520190565b03601f1981018352826128d0565b5190208403611380577f39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f229461130c9260036008936112938b610e688961110b33936001600160601b0360a01b165f52602660205260405f2090565b805460a08c901b61ffff60a01b1663ffffffff60a01b199091161760b089901b61ffff60b01b1617815560048101805460ff1916600117905592516002840155519101550180546112ef9060201c61ffff16613ed0565b613ed0565b65ffff0000000082549160201b169065ffff000000001916179055565b6113506113318261110b886001600160601b0360a01b165f52602760205260405f2090565b6113406112ea825461ffff1690565b61ffff1661ffff19825416179055565b6040805161ffff958616815294909116602085015283015233926001600160a01b0319169180606081015b0390a3005b63d1fed5fd60e01b5f5260045ffd5b9050608082015161122a6113be60a085015b516040805160208101958652908101919091529182906060820190565b51902014155f6111f5565b801591506111ee565b6020840151141591505f6111e6565b806113ef5f6113f5936128d0565b80610355565b5f611194565b633466526160e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b508615611093565b5061010061ffff87161161108c565b5061ffff861615611085565b5060018801546114519060101c61ffff16610507565b851161107e565b63965c290d60e01b5f5260045ffd5b34610290575f36600319011261029057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f916104505750604051908152602090f35b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102905760c03660031901126102905761156c610279565b60243561157881610842565b6044356064359160843560a435936115a4866001600160601b0360a01b165f52602260205260405f2090565b80546001600160a01b031695909390861561067f5760036115c9600587015460ff1690565b6115d28161251d565b036106705761ffff821696871580156118df575b6118bd576115f761106785896156b7565b80156118cc575b6118bd5761160e60038701613dbb565b906116198251151590565b90816118a9575b5061189a5760408101516001600160401b0316801515908161187e575b5061185c5761165961059860608301516001600160401b031690565b801515908161186b575b5061185c5761167f61059860808301516001600160401b031690565b8015159081611849575b50611827576116a561059860a08301516001600160401b031690565b8015159081611836575b5061182757602001516116c59061ffff16610507565b8015159081611806575b506117f7576116f68261110b8a6001600160601b0360a01b165f52602d60205260405f2090565b546117e8576117b160087fa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf968a61177d6117e39661110b8c8b61175f8c61122a8c6040519485936020850197889094939260609260808301968352602083015260408201520152565b519020936001600160601b0360a01b165f52602d60205260405f2090565b550180546117969060401c61ffff1660010161ffff1690565b61ffff60401b82549160401b169061ffff60401b1916179055565b60405193849333996001600160601b0360a01b1697859094939260609260808301968352602083015260408201520152565b0390a4005b6316feb18560e11b5f5260045ffd5b63464e67af60e01b5f5260045ffd5b905061ffff61181e600888015461ffff9060401c1690565b1610155f6116cf565b630410ff2960e31b5f5260045ffd5b90506001600160401b034216115f6116af565b90506001600160401b034316115f611689565b633deac39560e01b5f5260045ffd5b90506001600160401b034216105f611663565b6001600160401b031690506001600160401b034316105f61163d565b6330cd747160e01b5f5260045ffd5b6001600160a01b031690503314155f611620565b634c4d29cd60e11b5f5260045ffd5b506118da61106783876156b7565b6115fe565b5061010088116115e6565b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f916104505750604051908152602090f35b346102905760e03660031901126102905761196b610279565b60243561197781610842565b604435916064356084356001600160401b0381116102905761199d903690600401610294565b60a4929192356001600160401b038111610290576119bf903690600401610294565b93909260c435976001600160401b038911610290576119e5610353993690600401610294565b989097613f5c565b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f575f916104505750604051908152602090f35b346102905760a036600319011261029057611ab2610279565b602435611abe81610842565b6044356064356001600160401b03811161029057611ae0903690600401610294565b6084356001600160401b03811161029057611aff903690600401610294565b90611b1e876001600160601b0360a01b165f52602260205260405f2090565b80549094906001600160a01b03161561067f576002850154611b449060401c60ff161590565b611e59576003611b58600587015460ff1690565b611b618161251d565b0361067057611b8d61106761106061104e8b6001600160601b0360a01b165f52602360205260405f2090565b6114585761ffff87169384158015611e3c575b8015611e34575b611e2557611bd5611bcc8a6001600160601b0360a01b165f52602460205260405f2090565b6110c18a613e55565b90543360039290921b1c6001600160a01b03160361138057611c14600361111c33610e688d6001600160601b0360a01b165f52602960205260405f2090565b611e16577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f928592611c7060405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f57611c8d93611e02575b50810190614281565b611caf8561110b886001600160601b0360a01b165f52602a60205260405f2090565b54908051611cc06111d08960a01c90565b1492831593611df3575b508215611de4575b8215611ddb575b8215611db7575b50506113805781611d8e600861137b937f5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac956001611d3633610e688c6001600160601b0360a01b165f52602960205260405f2090565b805461ffff60a01b191660a08b901b61ffff60a01b1617815560038101805460ff191660011790550155018054611d739060301c61ffff16613ed0565b61ffff60301b82549160301b169061ffff60301b1916179055565b6040805161ffff9095168552602085019190915233946001600160a01b03191693918291820190565b90915061122a611dcf6113a160608401519360800190565b51902014155f80611ce0565b81159250611cd9565b60408101518514159250611cd2565b6020820151141592505f611cca565b806113ef5f611e10936128d0565b5f611c84565b63a89ac15160e01b5f5260045ffd5b639eae062d60e01b5f5260045ffd5b508615611ba7565b506001860154611e529060101c61ffff16610507565b8511611ba0565b630ba0cb2f60e21b5f5260045ffd5b60206040818301928281528451809452019201905f5b818110611e8b5750505090565b82516001600160a01b0316845260209384019390920191600101611e7e565b34610290576020366003190112610290576001600160a01b0319611ecc610279565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b818110611f13576109a785611f07818703826128d0565b60405191829182611e68565b82546001600160a01b0316845260209093019260019283019201611ef0565b34610290576040366003190112610290576109a7611f95611f51610279565b60243590611f5e82610ac1565b611f66613d8b565b506001600160a01b0319165f9081526025602090815260408083206001600160a01b0390941683529290522090565b611fea610ee5600460405193611faa85612847565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b6040519182918291909160a08060c0830194600180831b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b346102905761012036600319011261029057612056610279565b61205e61084e565b6044359160843560643560a43560c4356001600160401b0381116102905761208a903690600401610294565b9160e4356001600160401b038111610290576120aa903690600401610294565b95909461010435996001600160401b038b11610290576120d16103539b3690600401610294565b9a9099614324565b34610290576020366003190112610290576120f2610279565b6001600160a01b031981165f90815260226020526040902080549091906001600160a01b03161561067f576005820180546001840180549490939161214a6110676001600160401b03605089901c1660ff8416615a28565b61067057600882019561ffff6121736105076121688a5461ffff1690565b9360101c61ffff1690565b911610156124685761219f61106061104e856001600160601b0360a01b165f52602360205260405f2090565b612459576006820180549182156123d1575b50506040516313a4120960e31b815233600482015260a0816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047f576001916060915f916123a2575b50015161221881614874565b61222181614874565b036123935760408051602081019283523360601b6bffffffffffffffffffffffff1916918101919091526007919061225c816054810161122a565b5190209101541115612384576122da612277855461ffff1690565b9461229f3361229a856001600160601b0360a01b165f52602460205260405f2090565b61487e565b6122d16122c433610e68866001600160601b0360a01b165f52602360205260405f2090565b805460ff19166001179055565b611340866148c4565b60405161ffff851681526001600160a01b0319821694612334916123289190339088907f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b90602090a36148c4565b935460101c61ffff1690565b9261ffff80851691161461234457005b61235e9261235191615add565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b6123c4915060a03d60a0116123ca575b6123bc81836128d0565b810190614811565b5f61220c565b503d6123b2565b9091506123e99060481c6001600160401b0316610598565b8043111561244a574090811561243b578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f806121b1565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b906101008061085b9361ffff815116845261ffff60208201511660208501526124ab6040820151604086019061ffff169052565b60608181015161ffff169085015260808181015161ffff169085015260a0818101516001600160401b03169085015260c0818101516001600160401b03169085015260e0818101516001600160401b03169085015201511515910152565b634e487b7160e01b5f52602160045260245ffd5b6006111561252757565b612509565b9060068210156125275752565b81516001600160a01b031681526103408101929161085b9190610320906101809061256c60208201516020860190612477565b61257f60408201516101408601906109d0565b612592606082015161020086019061252c565b60808101516001600160401b031661022085015260a08101516001600160401b031661024085015260c081015161026085015260e081015161028085015261010081015161ffff166102a085015261012081015161ffff166102c085015261014081015161ffff166102e085015261016081015161ffff16610300850152015161ffff16910152565b34610290576020366003190112610290576109a76126fd6126f861263d610279565b5f61018060405161264d81612862565b82815260405161265c8161287e565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e082015283610100820152602082015261269d613d8b565b60408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015201526001600160601b0360a01b165f52602260205260405f2090565b614996565b60405191829182612539565b34610290575f366003190112610290576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610290576040366003190112610290576109a76127ba61276c610279565b6024359061277982610842565b5f604080516127878161289a565b82815282602082015201526001600160601b0360a01b165f52602860205260405f209061ffff165f5260205260405f2090565b6001604051916127c98361289a565b60ff815461ffff8116855260101c16151560208401520154604082015260405191829182919091604080606083019461ffff81511684526020810151151560208501520151910152565b634e487b7160e01b5f52604160045260245ffd5b60a081019081106001600160401b0382111761284257604052565b612813565b60c081019081106001600160401b0382111761284257604052565b6101a081019081106001600160401b0382111761284257604052565b61012081019081106001600160401b0382111761284257604052565b606081019081106001600160401b0382111761284257604052565b604081019081106001600160401b0382111761284257604052565b90601f801991011681019081106001600160401b0382111761284257604052565b908060209392818452848401375f828201840152601f01601f1916010190565b929061292a9061293895936040865260408601916128f1565b9260208185039101526128f1565b90565b6040513d5f823e3d90fd5b6040519061085b610120836128d0565b6040519061085b6101a0836128d0565b90610120828203126102905780601f83011215610290576101206040519261298e82856128d0565b8391810192831161029057905b8282106129a85750505090565b813581526020918201910161299b565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b908160051b91808304602014901517156129f657565b6129cc565b908160011b91808304600214901517156129f657565b818102929181159184041417156129f657565b90600182018092116129f657565b90608082018092116129f657565b919082018092116129f657565b9060068110156125275760ff80198354169116179055565b9592949399989791969097612a8e876001600160601b0360a01b165f52602260205260405f2090565b80549092906001600160a01b03161561067f576005830191612ab1835460ff1690565b612aba8161251d565b60038114612e2d5780612ace60029261251d565b0361067057612aea61059860028601546001600160401b031690565b4310610670576008840194612b05865461ffff9060101c1690565b96600186015493612b1e6105078661ffff9060201c1690565b61ffff8a1610612e1e578d158015612e16575b8015612e0e575b612dff577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f928592612b9460405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f57612bb193612deb575b50810190612966565b948551612bc16111d08b60a01c90565b1491821592612dcf575b8215612dac575b508115612d99575b508015612d8b575b8015612d7d575b8015612d6f575b6113805760408051602081018b81529181018a905260608101889052612c279190612c1e816080810161122a565b51902088614a8e565b8060e08601510361138057612c3d6108a06129e0565b860361138057612c7e958c612c7394612c6693612c5e610507996101000190565b51918c614c15565b805460ff19166003179055565b5460101c61ffff1690565b612c92612c8c6108606129e0565b88612a40565b965f5b828110612cf7575050507f5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a7245593949550612cf2906040519384936001600160601b0360a01b1696846040919493926060820195825260208201520152565b0390a2565b8060019160061b8a01604051612d298161122a6020820194602081013590358660209093929193604081019481520152565b519020612d68612d4d886001600160601b0360a01b165f52602a60205260405f2090565b61ffff8460051b8701351661ffff165f5260205260405f2090565b5501612c95565b508560c08501511415612bf0565b508760a08501511415612be9565b508860808501511415612be2565b6060860151915061ffff1614155f612bda565b909150612dc6610507604088015b519260101c61ffff1690565b1415905f612bd2565b91506020860151612de361ffff8416610507565b141591612bcb565b806113ef5f612df9936128d0565b5f612ba8565b63c5f680ed60e01b5f5260045ffd5b508a15612b38565b508c15612b31565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b90816020910312610290575190565b906001600160401b03809116911603906001600160401b0382116129f657565b906001600160401b03809116911601906001600160401b0382116129f657565b9060e0828203126102905780601f830112156102905760405191612eb060e0846128d0565b829060e0810192831161029057905b828210612ecc5750505090565b8135815260209182019101612ebf565b949293909796919596612f03866001600160601b0360a01b165f52602260205260405f2090565b80549093906001600160a01b03161561067f576002840154612f299060401c60ff161590565b611e595760058401946003612f3f875460ff1690565b612f488161251d565b03610670578a158015613175575b6131665761ffff996001612f73600888015461ffff9060301c1690565b96015495612f8461ffff8816610507565b9b8c911610613157577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f928592612fe560405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f5761300293613143575b50810190612e8b565b9182516130126111d08860a01c90565b14801590613135575b8015613127575b8015613119575b6113805760408301978851106113805760408051602081018b8152918101899052613065919061305c816060810161122a565b51902087614af5565b918260a0850151036113805761307b60406129e0565b03611380576040856130b2936130a68261309d6130ab9661ffff9060101c1690565b8d51908c614d8d565b614ea9565b9160c00190565b5103611380576130e86130f592857fbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b4829751614f67565b805460ff19166005179055565b6040805194855260208501929092526001600160a01b031916929081908101612cf2565b508660808401511415613029565b508860608401511415613022565b50876020840151141561301b565b806113ef5f613151936128d0565b5f612ff9565b63957674fd60e01b5f5260045ffd5b6314141ce560e21b5f5260045ffd5b508815612f56565b6040519061318a826128b5565b5f6020838281520152565b906040516131a2816128b5565b602060018294805484520154910152565b6131bb61317d565b506001600160601b0360a01b165f52602c60205260405f206001810154156131e65761293890613195565b506040516131f3816128b5565b5f81526001602082015290565b35612938816108b8565b3561293881610842565b908160209103126102905751612938816108b8565b634e487b7160e01b5f52601260045260245ffd5b8115613247570490565b613229565b6001600160401b03166001600160401b0381146129f65760010190565b90604082101561328257600c600183811c810193160290565b6129b8565b91908260c09103126102905760405161329f81612847565b60a08082946132ad816108d9565b845260208101356132bd81610842565b602085015260408101356132d0816108b8565b604085015260608101356132e3816108b8565b606085015260808101356132f6816108b8565b6080850152013591613307836108b8565b0152565b60068210156125275752565b906134ba610100600161085b9461333f61ffff865116829061ffff1661ffff19825416179055565b61336b613351602087015161ffff1690565b825463ffff0000191660109190911b63ffff000016178255565b61339b61337d604087015161ffff1690565b825465ffff00000000191660209190911b65ffff0000000016178255565b6133cc6133ad606087015161ffff1690565b825467ffff000000000000191660309190911b61ffff60301b16178255565b6133ff6133de608087015161ffff1690565b825469ffff0000000000000000191660409190911b61ffff60401b16178255565b61343e61341660a08701516001600160401b031690565b825467ffffffffffffffff60501b191660509190911b67ffffffffffffffff60501b16178255565b61347d61345560c08701516001600160401b031690565b825467ffffffffffffffff60901b191660909190911b67ffffffffffffffff60901b16178255565b01926134b361349660e08301516001600160401b031690565b855467ffffffffffffffff19166001600160401b03909116178555565b0151151590565b815468ff0000000000000000191690151560401b68ff000000000000000016179055565b60016135e060a061085b948051151560ff801987541691151516178555602081015162ffff0086549160081b169062ffff00191617855561355461352c60408301516001600160401b031690565b86546affffffffffffffff000000191660189190911b6affffffffffffffff00000016178655565b61359361356b60608301516001600160401b031690565b865467ffffffffffffffff60581b191660589190911b67ffffffffffffffff60581b16178655565b6135d26135aa60808301516001600160401b031690565b865467ffffffffffffffff60981b191660989190911b67ffffffffffffffff60981b16178655565b01516001600160401b031690565b9101906001600160401b03166001600160401b0319825416179055565b815181546001600160a01b0319166001600160a01b0390911617815561085b9190611796906101809060089061363a602086015160018301613317565b61364b6040860151600383016134de565b6136d76005820161366960608801516136638161251d565b82612a4d565b6136a461368060808901516001600160401b031690565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b60a0870151815470ffffffffffffffff000000000000000000191660489190911b67ffffffffffffffff60481b16179055565b60c0850151600682015560e08501516007820155019261371261370061010083015161ffff1690565b855461ffff191661ffff909116178555565b61373f61372561012083015161ffff1690565b855463ffff0000191660109190911b63ffff000016178555565b61377061375261014083015161ffff1690565b855465ffff00000000191660209190911b65ffff0000000016178555565b6137a261378361016083015161ffff1690565b855467ffff000000000000191660309190911b61ffff60301b16178555565b015161ffff1690565b94909895919793969261ffff86168015908115613d7e575b8115613d70575b508015613d64575b8015613d53575b8015613d44575b8015613d38575b8015613d29575b8015613cfd575b8015613ce1575b8015613cc5575b61066157604087016001600160401b0361381c82613200565b1615159081613cac575b81613c86575b508015613c1b575b8015613c01575b613bf257604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561047f576001600160401b03915f91613bc3575b50168015610661576138ae61ffff8c1661ffff8416612a11565b906127108210613b8c5750505f19985b6138f46138da6138d55f546001600160401b031690565b61324c565b6001600160401b03166001600160401b03195f5416175f55565b6139576139085f546001600160401b031690565b7f000000000000000000000000000000000000000000000000000000000000000060401b6bffffffff0000000000000000166001600160401b03919091161760a01b6001600160a01b03191690565b9a8b61398961396e6021546001600160401b031690565b60406001600160401b03603f831692161015613b4e57613269565b6139ab92906001600160601b0383549160031b9260a01c831b921b1916179055565b602180546001600160401b038082166001011667ffffffffffffffff199091161790556139d6612946565b61ffff909816885261ffff16602088015261ffff16604087015261ffff16606086015261ffff871660808601526001600160401b031660a08501526001600160401b031660c08401526001600160401b031660e083015215156101008201525f546001600160401b031691436001600160401b03169361ffff1692613a5b8486612e6b565b91613a64612956565b33815293602085015236613a7791613287565b6040840152600160608401526001600160401b031660808301526001600160401b031660a08201525f60c0820181905260e08201859052610100820181905261012082018190526101408201819052610160820181905261018082018190526001600160a01b03198616815260226020526040902090613af6916135fd565b613aff91612e6b565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917fcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d79190a390565b613b67613b5a82613269565b90549060031b1c60a01b90565b6001600160a01b03198116613b7d575b50613269565b613b86906152c5565b5f613b77565b613bb8613bbd927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43612a11565b61323d565b986138be565b613be5915060203d602011613beb575b613bdd81836128d0565b810190613214565b5f613894565b503d613bd3565b63148b7e9360e31b5f5260045ffd5b5061010061ffff613c1460208a0161320a565b161161383b565b5060608701613c2c61059882613200565b15159081613c6d575b81613c41575b50613834565b90506001600160401b03613c63610598613c5d60a08c01613200565b93613200565b911611155f613c3b565b9050613c7e61059860a08a01613200565b151590613c35565b90506001600160401b03613ca2610598613c5d60808c01613200565b911611155f61382c565b9050613cbd61059860808a01613200565b151590613826565b506001600160401b0383166001600160401b0385161115613803565b506001600160401b0382166001600160401b03841611156137fc565b50613d1861059861ffff8a166001600160401b034316612e6b565b6001600160401b03831611156137f5565b5061010061ffff8916116137ee565b5061ffff8816156137e7565b5061271061ffff8216106137e0565b5061ffff8a1661ffff8a16116137d9565b5061ffff8916156137d2565b905061ffff8b16105f6137ca565b61ffff8c161591506137c3565b60405190613d9882612847565b5f60a0838281528260208201528260408201528260608201528260808201520152565b9061085b604051613dcb81612847565b60a0613e4760018396613e39613e29825460ff81161515885261ffff8160081c1660208901526001600160401b03808260181c161660408901526001600160401b03808260581c161660608901526001600160401b039060981c1690565b6001600160401b03166080870152565b01546001600160401b031690565b6001600160401b0316910152565b61ffff5f199116019061ffff82116129f657565b8054821015613282575f5260205f2001905f90565b906101a0828203126102905780601f83011215610290576101a060405192613ea682856128d0565b8391810192831161029057905b828210613ec05750505090565b8135815260209182019101613eb3565b61ffff1661ffff81146129f65760010190565b906080116102905790608090565b909291928311610290579190565b90939293848311610290578411610290578101920390565b9291926001600160401b0382116128425760405191613f40601f8201601f1916602001846128d0565b829481845281830111610290578281602093845f960137010152565b949295919793613f80866001600160601b0360a01b165f52602260205260405f2090565b80549094906001600160a01b03161561067f576003613fa3600587015460ff1690565b613fac8161251d565b036106705761ffff83169a8b158015614276575b801561426e575b61425f57613fed8461110b8a6001600160601b0360a01b165f52602d60205260405f2090565b549687156142505761ffff9561402661401e8761110b8d6001600160601b0360a01b165f52602760205260405f2090565b5461ffff1690565b61403861050760018b015461ffff1690565b9788911610614241576140638661110b8c6001600160601b0360a01b165f52602860205260405f2090565b9b6140738d5460ff9060101c1690565b614232577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f9285926140cf60405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f576140eb936131435750810190612e8b565b9586516140fb6111d08a60a01c90565b14801590614224575b8015614216575b8015614208575b6113805760408701948551106113805760408051602081018d81529181018b905261414e9190614145816060810161122a565b51902089614b55565b918260a0890151036113805761416460646129e0565b81036113805761417761417e9185613ee3565b3691613f17565b6020815191012003611380576130a6826064946141a1976130ab9751918b61575a565b510361138057816001846141e17ff00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336966201000062ff000019825416179055565b01556040805194855260208501929092526001600160a01b0319169290819081015b0390a3565b508860808801511415614112565b508a6060880151141561410b565b508460208801511415614104565b63955c0c4960e01b5f5260045ffd5b63032cddf960e11b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508a15613fc7565b506101008c11613fc0565b9060a0828203126102905780601f8301121561029057604051916142a660a0846128d0565b829060a0810192831161029057905b8282106142c25750505090565b81358152602091820191016142b5565b90610140828203126102905780601f8301121561029057610140604051926142fa82856128d0565b8391810192831161029057905b8282106143145750505090565b8135815260209182019101614307565b9a999496939591989297909961434e8c6001600160601b0360a01b165f52602260205260405f2090565b80549096906001600160a01b03161561067f57600587015460ff169261439061106760018a01549561438a876001600160401b039060901c1690565b906158d7565b610670576143bc6110678f61104e611060916001600160601b0360a01b165f52602360205260405f2090565b6114585761ffff8d1695861580156147f8575b6147e9576144008f8f906143fa6110c1916001600160601b0360a01b165f52602460205260405f2090565b91613e55565b90543360039290921b1c6001600160a01b031603611380576144418f61111c600491610e6833916001600160601b0360a01b165f52602560205260405f2090565b6147da577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156102905784925f92859261449d60405198899586948594635c73957b60e11b865260048601612911565b03915afa92831561047f576144ba936147c6575b508101906142d2565b928b6144cb6111d086519260a01c90565b14918215926147aa575b821561478f575b508115614780575b508015614772575b8015614764575b6113805760408051602081018a8152918101899052614523919061451a816060810161122a565b5190208b614bb5565b8060c084015103611380578561010084015114801590614755575b6113805761010061454e816129e0565b85036113805761457e6141776108009661456c614177898389613ef1565b60208151910120976114009187613eff565b602081519101206145a38d6001600160601b0360a01b165f52602b60205260405f2090565b5403611380576145bd926145b692614ea9565b9160e00190565b51036113805761464b9161461f60046008936145f28c610e6833916001600160601b0360a01b165f52602560205260405f2090565b805461ffff60a01b191660a08d901b61ffff60a01b1617815590600382015501805460ff19166001179055565b0180546146329060101c61ffff16613ed0565b63ffff000082549160101b169063ffff00001916179055565b61468861466c876001600160601b0360a01b165f52602c60205260405f2090565b928354926001850193845480155f1461474f57506001906158fe565b9255557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b1561029057604051633c1bcdef60e21b8152336004820152925f908490602490829084905af192831561047f577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb59361473b575b506040805161ffff9095168552602085019190915283015233926001600160a01b031916918060608101614203565b806113ef5f614749936128d0565b5f61470c565b906158fe565b5086610120840151141561453e565b508660a083015114156144f3565b5087608083015114156144ec565b9050606083015114155f6144e4565b9091506147a161050760408601612dba565b1415905f6144dc565b915060208401516147be61ffff8416610507565b1415916144d5565b806113ef5f6147d4936128d0565b5f6144b1565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b5061480a601086901c61ffff16610507565b87116143cf565b908160a0910312610290576040519061482982612827565b805161483481610ac1565b825260208101516020830152604081015160408301526060810151906003821015610290576080916060840152015161486c816108b8565b608082015290565b6003111561252757565b805468010000000000000000811015612842576148a091600182018155613e69565b81546001600160a01b0393841660039290921b91821b9390911b1916919091179055565b61ffff60019116019061ffff82116129f657565b9061085b6040516148e88161287e565b61010061498f60018396614975614965825461ffff8116885261491961ffff8260101c1660208a019061ffff169052565b61ffff602082901c16604089015261ffff603082901c16606089015261ffff604082901c1660808901526001600160401b03605082901c1660a089015260901c6001600160401b031690565b6001600160401b031660c0870152565b01546001600160401b03811660e085015260401c60ff1690565b1515910152565b9061085b614a8260086149a7612956565b85546001600160a01b03168152946149c1600182016148d8565b60208701526149d260038201613dbb565b6040870152614a2a614a1a60058301546149f86149ef8260ff1690565b60608b0161330b565b6001600160401b03600882901c1660808a015260481c6001600160401b031690565b6001600160401b031660a0880152565b600681015460c0870152600781015460e0870152015461ffff811661010086015261ffff601082901c1661012086015261ffff602082901c1661014086015261ffff603082901c1661016086015260401c61ffff1690565b61ffff16610180840152565b5f516020615e295f395f51905f52916040519060208201926001600160601b0360a01b1683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c8152614aee606c826128d0565b5190200690565b5f516020615e295f395f51905f52916040519060208201926001600160601b0360a01b1683527fc5cb4182e179e0279f50e2d772929d40dc9d4db3b30ec2ebbefbe6b9bb543075602c830152604c820152604c8152614aee606c826128d0565b5f516020615e295f395f51905f52916040519060208201926001600160601b0360a01b1683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c8152614aee606c826128d0565b5f516020615e295f395f51905f52916040519060208201926001600160601b0360a01b1683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c8152614aee606c826128d0565b9091939294614c3590610400614c2e6201000082612a40565b9186613eff565b91614c616105076001614c54610507600889015461ffff9060101c1690565b96015460101c61ffff1690565b92610800915f5b868110614c8b575050505050505090614c84916108a091614ea9565b0361138057565b61ffff8160051b8901351680158015614d84575b61138057614d05614ce8614cd3614cca866001600160601b0360a01b165f52602460205260405f2090565b6110c185613e55565b905460039190911b1c6001600160a01b031690565b6001600160a01b031985165f908152602560205260409020610e68565b90614d17611067600484015460ff1690565b908115614d67575b50611380576003614d50614177614d368886612a11565b614d4889614d4388612a24565b612a11565b90888b613eff565b602081519101209101540361138057600101614c68565b9050614d7c610507835461ffff9060a01c1690565b14155f614d1f565b50868111614c9f565b919091614d9c61040085612a40565b925f5b838110614de3575050505b60208110614db757505050565b8060051b808401351590811591614dd6575b5061138057600101614daa565b905082013515155f614dc9565b8060051b61ffff81880135169081158015614e9c575b61138057614e4a614e2d614cd3614e24886001600160601b0360a01b165f52602460205260405f2090565b6110c186613e55565b6001600160a01b031987165f908152602960205260409020610e68565b91614e5c611067600385015460ff1690565b908115614e7f575b50611380576001908701359101540361138057600101614d9f565b9050614e94610507845461ffff9060a01c1690565b14155f614e64565b5061ffff84168211614df9565b9291905f516020615e295f395f51905f525f940691829060051b8201915b828110614ed45750505050565b909192945f516020615e295f395f51905f5283816020938186358b099008970993929101614ec7565b6001600160401b0381116128425760051b60200190565b90614f1e82614efd565b614f2b60405191826128d0565b8281528092614f3c601f1991614efd565b0190602036910137565b8051156132825760200190565b80518210156132825760209160051b010190565b91614f7183614f14565b92614f7b81614f14565b916104008101908181116129f6575f5b83811061526b5750505080158015615261575b8015615257575b61524857614fb581949294614f14565b90614fbf81614f14565b945f93845b838110615195575050614fd682614f14565b614fdf84614f46565b51614fe982614f46565b5260015b838110615141575061501061500a61500485615c9e565b83614f53565b51615dad565b9061501a84614f14565b9484915b600183116150d95750505061503284614f46565b525f955f945b8386106150725750505050505061505c905f516020615e495f395f51905f52900690565b0361506357565b6373bdb71560e11b5f5260045ffd5b6150828683999495969799614f53565b5161508d8988614f53565b5191613247575f516020615e495f395f51905f528091816001946150c86150b48e8b614f53565b515f516020615e495f395f51905f52900690565b9209095f9408970194939291615038565b909192939495966150e984615c9e565b6150fb6150f582615c9e565b84614f53565b51916132475761512b5f516020615e495f395f51905f529182886151359509615124828d614f53565b5285614f53565b515f960993615cac565b9190969594939661501e565b948561515a615154839596979498615c9e565b88614f53565b516151658285614f53565b5192613247576001925f516020615e495f395f51905f5291096151888289614f53565b5201949093929194614fed565b6151a76150b482849895969798614f53565b600180915f905b8882106151dd5750509082916151c66001948c614f53565b526151d18289614f53565b52019493929194614fc4565b9092959183851461523d576151f56150b48588614f53565b9283831461522e57613247575f516020615e495f395f51905f52808085600194099461522085615c78565b90085f9809935b01906151ae565b63027639eb60e31b5f5260045ffd5b919592600190615227565b630a4960f960e31b5f5260045ffd5b5081518111614fa5565b5083518111614f9e565b8060019160051b80840135615280838b614f53565b5284013561528e8288614f53565b5201614f8b565b8054905f8155816152a4575050565b5f5260205f20908101905b8181106152ba575050565b5f81556001016152af565b6001600160a01b031981165f908152602260205260409020546001600160a01b0316156156b4576001600160a01b031981165f9081526024602052604081208054915b8281106155235750505060015b61010061ffff82161115615405575061534a615345826001600160601b0360a01b165f52602460205260405f2090565b615295565b6001600160a01b031981165f908152602b60209081526040808320839055602c909152902061537f905b60015f918281550155565b6153d56153a0826001600160601b0360a01b165f52602260205260405f2090565b60085f918281558260018201558260028201558260038201558260048201558260058201558260068201558260078201550155565b6001600160a01b0319167f98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef265f80a2565b8061543161050761401e6154969461110b876001600160601b0360a01b165f52602760205260405f2090565b6154ee575b61546561545b8261110b866001600160601b0360a01b165f52602860205260405f2090565b5460101c60ff1690565b6154c4575b61548c8161110b856001600160601b0360a01b165f52602d60205260405f2090565b5461549b57613ed0565b615315565b5f6154be8261110b866001600160601b0360a01b165f52602d60205260405f2090565b55613ed0565b6154e96153748261110b866001600160601b0360a01b165f52602860205260405f2090565b61546a565b61551e6155138261110b866001600160601b0360a01b165f52602760205260405f2090565b805461ffff19169055565b615436565b615530614cd38284613e69565b61555f61555582610e68886001600160601b0360a01b165f52602360205260405f2090565b805460ff19169055565b6001600160a01b031985165f908152602a602052604081206155979061558761050786612a24565b61ffff165f5260205260405f2090565b556155da6155bd82610e68886001600160601b0360a01b165f52602560205260405f2090565b60045f918281558260018201558260028201558260038201550155565b6156166155ff82610e68886001600160601b0360a01b165f52602960205260405f2090565b60035f918281558260018201558260028201550155565b60015b61010061ffff82161115615631575050600101615308565b80615660600461111c85610e686156699661110b8d6001600160601b0360a01b165f52602660205260405f2090565b61566e57613ed0565b615619565b6112ea61569784610e688461110b8c6001600160601b0360a01b165f52602660205260405f2090565b60045f918281558260018201558260016002830182815501550155565b50565b5f516020615e295f395f51905f528110801590615743575b61573d575f516020615e295f395f51905f528082819309928009818080808487097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09600108937f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000009081490565b50505f90565b505f516020615e295f395f51905f528210156156cf565b909392919361578a610507600161577d61577388612a32565b9761048090612a40565b97015460101c61ffff1690565b905f5b8481106157e057505050505b602081106157a657505050565b8060061b83018160051b83013561138057803515908115916157d0575b5061138057600101615799565b600191506020013514155f6157c3565b8060061b870161ffff8260051b8801351690811580156158ce575b61138057615848615826614cd3614e24896001600160601b0360a01b165f52602460205260405f2090565b610e688661110b8a6001600160601b0360a01b165f52602660205260405f2090565b9161585a611067600485015460ff1690565b9081156158b1575b5061138057815460b01c61ffff1661ffff8086169116036113805760028201548135149182159261589c575b50506113805760010161578d565b60209192506003015491013514155f8061588e565b90506158c6610507845461ffff9060a01c1690565b14155f615862565b508482116157fb565b60068110156125275760021490816158ed575090565b6001600160401b0391501643111590565b9392919091841580615a1e575b615a1657811580615a0c575b615a07575f516020615e295f395f51905f52828609945f516020615e295f395f51905f528285095f516020615e295f395f51905f528188095f516020615e295f395f51905f5290620292f809965f516020615e295f395f51905f5290620292fc0961598191615d00565b935f516020615e295f395f51905f528760010861599d90615d47565b935f516020615e295f395f51905f529109915f516020615e295f395f51905f529109905f516020615e295f395f51905f529108905f516020615e295f395f51905f529109926159eb90615cb8565b6159f490615d47565b5f516020615e295f395f51905f52910990565b505090565b5060018114615917565b935090509190565b506001831461590b565b60068110156125275760011490816158ed575090565b60405190610400615a4f81846128d0565b368337565b60405190610800615a4f81846128d0565b9060208110156132825760051b0190565b9060408110156132825760051b0190565b91905f835b60208210615ac75750505061040082015f905b60408210615ab157505050610c000190565b6020806001928551815201930191019091615a9f565b6020806001928551815201930191019091615a8c565b919091615ae8615a3e565b615af0615a54565b93615b0f836001600160601b0360a01b165f52602460205260405f2090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110615bba5750505061ffff165b60208110615b95575061122a615b74615b92939495604051928391602083019586615a87565b519020916001600160601b0360a01b165f52602b60205260405f2090565b55565b806001615bb3615bad615ba883956129fb565b612a24565b88615a76565b5201615b4e565b80615bc7615c0a92612a24565b615bd18288615a65565b5260a0615be1614cd38387613e69565b6040516313a4120960e31b81526001600160a01b03909116600482015292839081906024820190565b0381865afa91821561047f576001926040915f91615c5a575b506020810151615c3b615c35856129fb565b8d615a76565b520151615c53615c4d615ba8846129fb565b8b615a76565b5201615b3b565b615c72915060a03d81116123ca576123bc81836128d0565b5f615c23565b5f516020615e495f395f51905f5203905f516020615e495f395f51905f5282116129f657565b5f198101919082116129f657565b80156129f6575f190190565b5f516020615e295f395f51905f5290065f516020615e295f395f51905f52035f516020615e295f395f51905f5281116129f6575f516020615e295f395f51905f529060010890565b905f516020615e295f395f51905f5290065f516020615e295f395f51905f52035f516020615e295f395f51905f5281116129f6575f516020615e295f395f51905f52910890565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f516020615e295f395f51905f5260a082015260208160c08160055afa15610290575190565b9060405191602083526020808401526020604084015260608301527f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126ef60808301525f516020615e495f395f51905f5260a083015260208260c08160055afa15615e1b5760c082519201604052565b639e44e6e05f526004601cfdfe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1a26469706673582212207251f14ce6a353c5163c3849cb046a8feb45754a1bf7eb07327f0c0453f1579164736f6c634300081c0033",
}

// DKGManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGManagerMetaData.ABI instead.
var DKGManagerABI = DKGManagerMetaData.ABI

// DKGManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGManagerMetaData.Bin instead.
var DKGManagerBin = DKGManagerMetaData.Bin

// DeployDKGManager deploys a new Ethereum contract, binding an instance of DKGManager to it.
func DeployDKGManager(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint32, _registry common.Address, _contributionVerifier common.Address, _partialDecryptVerifier common.Address, _finalizeVerifier common.Address, _decryptCombineVerifier common.Address, _revealSubmitVerifier common.Address, _revealShareVerifier common.Address) (common.Address, *types.Transaction, *DKGManager, error) {
	parsed, err := DKGManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGManagerBin), backend, _chainId, _registry, _contributionVerifier, _partialDecryptVerifier, _finalizeVerifier, _decryptCombineVerifier, _revealSubmitVerifier, _revealShareVerifier)
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

// REVEALSHAREVERIFIER is a free data retrieval call binding the contract method 0x5ddd0626.
//
// Solidity: function REVEAL_SHARE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) REVEALSHAREVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "REVEAL_SHARE_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REVEALSHAREVERIFIER is a free data retrieval call binding the contract method 0x5ddd0626.
//
// Solidity: function REVEAL_SHARE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) REVEALSHAREVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.REVEALSHAREVERIFIER(&_DKGManager.CallOpts)
}

// REVEALSHAREVERIFIER is a free data retrieval call binding the contract method 0x5ddd0626.
//
// Solidity: function REVEAL_SHARE_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) REVEALSHAREVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.REVEALSHAREVERIFIER(&_DKGManager.CallOpts)
}

// REVEALSUBMITVERIFIER is a free data retrieval call binding the contract method 0x070c7492.
//
// Solidity: function REVEAL_SUBMIT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCaller) REVEALSUBMITVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "REVEAL_SUBMIT_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REVEALSUBMITVERIFIER is a free data retrieval call binding the contract method 0x070c7492.
//
// Solidity: function REVEAL_SUBMIT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerSession) REVEALSUBMITVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.REVEALSUBMITVERIFIER(&_DKGManager.CallOpts)
}

// REVEALSUBMITVERIFIER is a free data retrieval call binding the contract method 0x070c7492.
//
// Solidity: function REVEAL_SUBMIT_VERIFIER() view returns(address)
func (_DKGManager *DKGManagerCallerSession) REVEALSUBMITVERIFIER() (common.Address, error) {
	return _DKGManager.Contract.REVEALSUBMITVERIFIER(&_DKGManager.CallOpts)
}

// ROUNDPREFIX is a free data retrieval call binding the contract method 0x56664d15.
//
// Solidity: function ROUND_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerCaller) ROUNDPREFIX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "ROUND_PREFIX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ROUNDPREFIX is a free data retrieval call binding the contract method 0x56664d15.
//
// Solidity: function ROUND_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerSession) ROUNDPREFIX() (uint32, error) {
	return _DKGManager.Contract.ROUNDPREFIX(&_DKGManager.CallOpts)
}

// ROUNDPREFIX is a free data retrieval call binding the contract method 0x56664d15.
//
// Solidity: function ROUND_PREFIX() view returns(uint32)
func (_DKGManager *DKGManagerCallerSession) ROUNDPREFIX() (uint32, error) {
	return _DKGManager.Contract.ROUNDPREFIX(&_DKGManager.CallOpts)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 roundId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetCiphertextHash(opts *bind.CallOpts, roundId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCiphertextHash", roundId, ciphertextIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 roundId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetCiphertextHash(roundId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetCiphertextHash is a free data retrieval call binding the contract method 0x373877a6.
//
// Solidity: function getCiphertextHash(bytes12 roundId, uint16 ciphertextIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetCiphertextHash(roundId [12]byte, ciphertextIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetCiphertextHash(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 roundId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerCaller) GetCollectivePublicKey(opts *bind.CallOpts, roundId [12]byte) (DKGTypesPoint, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCollectivePublicKey", roundId)

	if err != nil {
		return *new(DKGTypesPoint), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesPoint)).(*DKGTypesPoint)

	return out0, err

}

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 roundId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerSession) GetCollectivePublicKey(roundId [12]byte) (DKGTypesPoint, error) {
	return _DKGManager.Contract.GetCollectivePublicKey(&_DKGManager.CallOpts, roundId)
}

// GetCollectivePublicKey is a free data retrieval call binding the contract method 0x3353ec6e.
//
// Solidity: function getCollectivePublicKey(bytes12 roundId) view returns((uint256,uint256))
func (_DKGManager *DKGManagerCallerSession) GetCollectivePublicKey(roundId [12]byte) (DKGTypesPoint, error) {
	return _DKGManager.Contract.GetCollectivePublicKey(&_DKGManager.CallOpts, roundId)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCaller) GetCombinedDecryption(opts *bind.CallOpts, roundId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getCombinedDecryption", roundId, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesCombinedDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesCombinedDecryptionRecord)).(*DKGTypesCombinedDecryptionRecord)

	return out0, err

}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerSession) GetCombinedDecryption(roundId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bool,uint256))
func (_DKGManager *DKGManagerCallerSession) GetCombinedDecryption(roundId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 roundId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerCaller) GetContribution(opts *bind.CallOpts, roundId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getContribution", roundId, contributor)

	if err != nil {
		return *new(DKGTypesContributionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesContributionRecord)).(*DKGTypesContributionRecord)

	return out0, err

}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 roundId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetContribution(roundId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	return _DKGManager.Contract.GetContribution(&_DKGManager.CallOpts, roundId, contributor)
}

// GetContribution is a free data retrieval call binding the contract method 0xd3720aac.
//
// Solidity: function getContribution(bytes12 roundId, address contributor) view returns((address,uint16,bytes32,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerCallerSession) GetContribution(roundId [12]byte, contributor common.Address) (DKGTypesContributionRecord, error) {
	return _DKGManager.Contract.GetContribution(&_DKGManager.CallOpts, roundId, contributor)
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
// Solidity: function getDecryptionPolicy(bytes12 roundId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerCaller) GetDecryptionPolicy(opts *bind.CallOpts, roundId [12]byte) (DKGTypesDecryptionPolicy, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getDecryptionPolicy", roundId)

	if err != nil {
		return *new(DKGTypesDecryptionPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesDecryptionPolicy)).(*DKGTypesDecryptionPolicy)

	return out0, err

}

// GetDecryptionPolicy is a free data retrieval call binding the contract method 0x4554c0be.
//
// Solidity: function getDecryptionPolicy(bytes12 roundId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerSession) GetDecryptionPolicy(roundId [12]byte) (DKGTypesDecryptionPolicy, error) {
	return _DKGManager.Contract.GetDecryptionPolicy(&_DKGManager.CallOpts, roundId)
}

// GetDecryptionPolicy is a free data retrieval call binding the contract method 0x4554c0be.
//
// Solidity: function getDecryptionPolicy(bytes12 roundId) view returns((bool,uint16,uint64,uint64,uint64,uint64))
func (_DKGManager *DKGManagerCallerSession) GetDecryptionPolicy(roundId [12]byte) (DKGTypesDecryptionPolicy, error) {
	return _DKGManager.Contract.GetDecryptionPolicy(&_DKGManager.CallOpts, roundId)
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
// Solidity: function getPartialDecryption(bytes12 roundId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerCaller) GetPartialDecryption(opts *bind.CallOpts, roundId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPartialDecryption", roundId, participant, ciphertextIndex)

	if err != nil {
		return *new(DKGTypesPartialDecryptionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesPartialDecryptionRecord)).(*DKGTypesPartialDecryptionRecord)

	return out0, err

}

// GetPartialDecryption is a free data retrieval call binding the contract method 0x70f2469b.
//
// Solidity: function getPartialDecryption(bytes12 roundId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerSession) GetPartialDecryption(roundId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, roundId, participant, ciphertextIndex)
}

// GetPartialDecryption is a free data retrieval call binding the contract method 0x70f2469b.
//
// Solidity: function getPartialDecryption(bytes12 roundId, address participant, uint16 ciphertextIndex) view returns((address,uint16,uint16,bytes32,(uint256,uint256),bool))
func (_DKGManager *DKGManagerCallerSession) GetPartialDecryption(roundId [12]byte, participant common.Address, ciphertextIndex uint16) (DKGTypesPartialDecryptionRecord, error) {
	return _DKGManager.Contract.GetPartialDecryption(&_DKGManager.CallOpts, roundId, participant, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 roundId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCaller) GetPlaintext(opts *bind.CallOpts, roundId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getPlaintext", roundId, ciphertextIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 roundId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerSession) GetPlaintext(roundId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetPlaintext is a free data retrieval call binding the contract method 0x6759e0e1.
//
// Solidity: function getPlaintext(bytes12 roundId, uint16 ciphertextIndex) view returns(uint256)
func (_DKGManager *DKGManagerCallerSession) GetPlaintext(roundId [12]byte, ciphertextIndex uint16) (*big.Int, error) {
	return _DKGManager.Contract.GetPlaintext(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetRevealShareVerifierVKeyHash is a free data retrieval call binding the contract method 0xc2440e16.
//
// Solidity: function getRevealShareVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetRevealShareVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getRevealShareVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRevealShareVerifierVKeyHash is a free data retrieval call binding the contract method 0xc2440e16.
//
// Solidity: function getRevealShareVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetRevealShareVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetRevealShareVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetRevealShareVerifierVKeyHash is a free data retrieval call binding the contract method 0xc2440e16.
//
// Solidity: function getRevealShareVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetRevealShareVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetRevealShareVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetRevealSubmitVerifierVKeyHash is a free data retrieval call binding the contract method 0xb18730c2.
//
// Solidity: function getRevealSubmitVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetRevealSubmitVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getRevealSubmitVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRevealSubmitVerifierVKeyHash is a free data retrieval call binding the contract method 0xb18730c2.
//
// Solidity: function getRevealSubmitVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetRevealSubmitVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetRevealSubmitVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetRevealSubmitVerifierVKeyHash is a free data retrieval call binding the contract method 0xb18730c2.
//
// Solidity: function getRevealSubmitVerifierVKeyHash() view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetRevealSubmitVerifierVKeyHash() ([32]byte, error) {
	return _DKGManager.Contract.GetRevealSubmitVerifierVKeyHash(&_DKGManager.CallOpts)
}

// GetRevealedShare is a free data retrieval call binding the contract method 0x53d72184.
//
// Solidity: function getRevealedShare(bytes12 roundId, address participant) view returns((address,uint16,uint256,bytes32,bool))
func (_DKGManager *DKGManagerCaller) GetRevealedShare(opts *bind.CallOpts, roundId [12]byte, participant common.Address) (DKGTypesRevealedShareRecord, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getRevealedShare", roundId, participant)

	if err != nil {
		return *new(DKGTypesRevealedShareRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DKGTypesRevealedShareRecord)).(*DKGTypesRevealedShareRecord)

	return out0, err

}

// GetRevealedShare is a free data retrieval call binding the contract method 0x53d72184.
//
// Solidity: function getRevealedShare(bytes12 roundId, address participant) view returns((address,uint16,uint256,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetRevealedShare(roundId [12]byte, participant common.Address) (DKGTypesRevealedShareRecord, error) {
	return _DKGManager.Contract.GetRevealedShare(&_DKGManager.CallOpts, roundId, participant)
}

// GetRevealedShare is a free data retrieval call binding the contract method 0x53d72184.
//
// Solidity: function getRevealedShare(bytes12 roundId, address participant) view returns((address,uint16,uint256,bytes32,bool))
func (_DKGManager *DKGManagerCallerSession) GetRevealedShare(roundId [12]byte, participant common.Address) (DKGTypesRevealedShareRecord, error) {
	return _DKGManager.Contract.GetRevealedShare(&_DKGManager.CallOpts, roundId, participant)
}

// GetRound is a free data retrieval call binding the contract method 0xf4e34945.
//
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerCaller) GetRound(opts *bind.CallOpts, roundId [12]byte) (IDKGManagerRound, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getRound", roundId)

	if err != nil {
		return *new(IDKGManagerRound), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGManagerRound)).(*IDKGManagerRound)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0xf4e34945.
//
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetRound(roundId [12]byte) (IDKGManagerRound, error) {
	return _DKGManager.Contract.GetRound(&_DKGManager.CallOpts, roundId)
}

// GetRound is a free data retrieval call binding the contract method 0xf4e34945.
//
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerCallerSession) GetRound(roundId [12]byte) (IDKGManagerRound, error) {
	return _DKGManager.Contract.GetRound(&_DKGManager.CallOpts, roundId)
}

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 roundId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCaller) GetShareCommitmentHash(opts *bind.CallOpts, roundId [12]byte, participantIndex uint16) ([32]byte, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "getShareCommitmentHash", roundId, participantIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 roundId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerSession) GetShareCommitmentHash(roundId [12]byte, participantIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetShareCommitmentHash(&_DKGManager.CallOpts, roundId, participantIndex)
}

// GetShareCommitmentHash is a free data retrieval call binding the contract method 0x510ba2df.
//
// Solidity: function getShareCommitmentHash(bytes12 roundId, uint16 participantIndex) view returns(bytes32)
func (_DKGManager *DKGManagerCallerSession) GetShareCommitmentHash(roundId [12]byte, participantIndex uint16) ([32]byte, error) {
	return _DKGManager.Contract.GetShareCommitmentHash(&_DKGManager.CallOpts, roundId, participantIndex)
}

// RoundNonce is a free data retrieval call binding the contract method 0x415a1b86.
//
// Solidity: function roundNonce() view returns(uint64)
func (_DKGManager *DKGManagerCaller) RoundNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "roundNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RoundNonce is a free data retrieval call binding the contract method 0x415a1b86.
//
// Solidity: function roundNonce() view returns(uint64)
func (_DKGManager *DKGManagerSession) RoundNonce() (uint64, error) {
	return _DKGManager.Contract.RoundNonce(&_DKGManager.CallOpts)
}

// RoundNonce is a free data retrieval call binding the contract method 0x415a1b86.
//
// Solidity: function roundNonce() view returns(uint64)
func (_DKGManager *DKGManagerCallerSession) RoundNonce() (uint64, error) {
	return _DKGManager.Contract.RoundNonce(&_DKGManager.CallOpts)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 roundId) view returns(address[])
func (_DKGManager *DKGManagerCaller) SelectedParticipants(opts *bind.CallOpts, roundId [12]byte) ([]common.Address, error) {
	var out []interface{}
	err := _DKGManager.contract.Call(opts, &out, "selectedParticipants", roundId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 roundId) view returns(address[])
func (_DKGManager *DKGManagerSession) SelectedParticipants(roundId [12]byte) ([]common.Address, error) {
	return _DKGManager.Contract.SelectedParticipants(&_DKGManager.CallOpts, roundId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0xca3c0458.
//
// Solidity: function selectedParticipants(bytes12 roundId) view returns(address[])
func (_DKGManager *DKGManagerCallerSession) SelectedParticipants(roundId [12]byte) ([]common.Address, error) {
	return _DKGManager.Contract.SelectedParticipants(&_DKGManager.CallOpts, roundId)
}

// AbortRound is a paid mutator transaction binding the contract method 0x349181a2.
//
// Solidity: function abortRound(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactor) AbortRound(opts *bind.TransactOpts, roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "abortRound", roundId)
}

// AbortRound is a paid mutator transaction binding the contract method 0x349181a2.
//
// Solidity: function abortRound(bytes12 roundId) returns()
func (_DKGManager *DKGManagerSession) AbortRound(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.AbortRound(&_DKGManager.TransactOpts, roundId)
}

// AbortRound is a paid mutator transaction binding the contract method 0x349181a2.
//
// Solidity: function abortRound(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactorSession) AbortRound(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.AbortRound(&_DKGManager.TransactOpts, roundId)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactor) ClaimSlot(opts *bind.TransactOpts, roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "claimSlot", roundId)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 roundId) returns()
func (_DKGManager *DKGManagerSession) ClaimSlot(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimSlot(&_DKGManager.TransactOpts, roundId)
}

// ClaimSlot is a paid mutator transaction binding the contract method 0xd9933767.
//
// Solidity: function claimSlot(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactorSession) ClaimSlot(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ClaimSlot(&_DKGManager.TransactOpts, roundId)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0xb58aab90.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) CombineDecryption(opts *bind.TransactOpts, roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "combineDecryption", roundId, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0xb58aab90.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) CombineDecryption(roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, roundId, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0xb58aab90.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, uint256 plaintext, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) CombineDecryption(roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintext *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, roundId, ciphertextIndex, combineHash, plaintext, transcript, proof, input)
}

// CreateRound is a paid mutator transaction binding the contract method 0x3caf4487.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateRound(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createRound", threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, disclosureAllowed, decryptionPolicy)
}

// CreateRound is a paid mutator transaction binding the contract method 0x3caf4487.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, disclosureAllowed, decryptionPolicy)
}

// CreateRound is a paid mutator transaction binding the contract method 0x3caf4487.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, uint64 finalizeNotBeforeBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, finalizeNotBeforeBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock, disclosureAllowed, decryptionPolicy)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactor) ExtendRegistration(opts *bind.TransactOpts, roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "extendRegistration", roundId)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 roundId) returns()
func (_DKGManager *DKGManagerSession) ExtendRegistration(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ExtendRegistration(&_DKGManager.TransactOpts, roundId)
}

// ExtendRegistration is a paid mutator transaction binding the contract method 0x0b1451f0.
//
// Solidity: function extendRegistration(bytes12 roundId) returns()
func (_DKGManager *DKGManagerTransactorSession) ExtendRegistration(roundId [12]byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ExtendRegistration(&_DKGManager.TransactOpts, roundId)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0x058994a1.
//
// Solidity: function finalizeRound(bytes12 roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) FinalizeRound(opts *bind.TransactOpts, roundId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "finalizeRound", roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0x058994a1.
//
// Solidity: function finalizeRound(bytes12 roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) FinalizeRound(roundId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeRound(&_DKGManager.TransactOpts, roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0x058994a1.
//
// Solidity: function finalizeRound(bytes12 roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) FinalizeRound(roundId [12]byte, aggregateCommitmentsHash [32]byte, collectivePublicKeyHash [32]byte, shareCommitmentHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.FinalizeRound(&_DKGManager.TransactOpts, roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)
}

// ReconstructSecret is a paid mutator transaction binding the contract method 0x0e2c53f7.
//
// Solidity: function reconstructSecret(bytes12 roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) ReconstructSecret(opts *bind.TransactOpts, roundId [12]byte, disclosureHash [32]byte, reconstructedSecretHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "reconstructSecret", roundId, disclosureHash, reconstructedSecretHash, transcript, proof, input)
}

// ReconstructSecret is a paid mutator transaction binding the contract method 0x0e2c53f7.
//
// Solidity: function reconstructSecret(bytes12 roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) ReconstructSecret(roundId [12]byte, disclosureHash [32]byte, reconstructedSecretHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ReconstructSecret(&_DKGManager.TransactOpts, roundId, disclosureHash, reconstructedSecretHash, transcript, proof, input)
}

// ReconstructSecret is a paid mutator transaction binding the contract method 0x0e2c53f7.
//
// Solidity: function reconstructSecret(bytes12 roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) ReconstructSecret(roundId [12]byte, disclosureHash [32]byte, reconstructedSecretHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.ReconstructSecret(&_DKGManager.TransactOpts, roundId, disclosureHash, reconstructedSecretHash, transcript, proof, input)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 roundId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactor) SubmitCiphertext(opts *bind.TransactOpts, roundId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitCiphertext", roundId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 roundId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerSession) SubmitCiphertext(roundId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, roundId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitCiphertext is a paid mutator transaction binding the contract method 0xa9c4b25f.
//
// Solidity: function submitCiphertext(bytes12 roundId, uint16 ciphertextIndex, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitCiphertext(roundId [12]byte, ciphertextIndex uint16, c1x *big.Int, c1y *big.Int, c2x *big.Int, c2y *big.Int) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitCiphertext(&_DKGManager.TransactOpts, roundId, ciphertextIndex, c1x, c1y, c2x, c2y)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitContribution(opts *bind.TransactOpts, roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitContribution", roundId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitContribution(roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, roundId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xd6c29c9e.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, uint256 commitment0X, uint256 commitment0Y, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitContribution(roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, commitment0X *big.Int, commitment0Y *big.Int, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, roundId, contributorIndex, commitmentsHash, encryptedSharesHash, commitment0X, commitment0Y, transcript, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x802ae231.
//
// Solidity: function submitPartialDecryption(bytes12 roundId, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitPartialDecryption(opts *bind.TransactOpts, roundId [12]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitPartialDecryption", roundId, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x802ae231.
//
// Solidity: function submitPartialDecryption(bytes12 roundId, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitPartialDecryption(roundId [12]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, roundId, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// SubmitPartialDecryption is a paid mutator transaction binding the contract method 0x802ae231.
//
// Solidity: function submitPartialDecryption(bytes12 roundId, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitPartialDecryption(roundId [12]byte, participantIndex uint16, ciphertextIndex uint16, deltaHash [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitPartialDecryption(&_DKGManager.TransactOpts, roundId, participantIndex, ciphertextIndex, deltaHash, proof, input)
}

// SubmitRevealedShare is a paid mutator transaction binding the contract method 0xc9396bf0.
//
// Solidity: function submitRevealedShare(bytes12 roundId, uint16 participantIndex, uint256 shareValue, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitRevealedShare(opts *bind.TransactOpts, roundId [12]byte, participantIndex uint16, shareValue *big.Int, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitRevealedShare", roundId, participantIndex, shareValue, proof, input)
}

// SubmitRevealedShare is a paid mutator transaction binding the contract method 0xc9396bf0.
//
// Solidity: function submitRevealedShare(bytes12 roundId, uint16 participantIndex, uint256 shareValue, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitRevealedShare(roundId [12]byte, participantIndex uint16, shareValue *big.Int, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitRevealedShare(&_DKGManager.TransactOpts, roundId, participantIndex, shareValue, proof, input)
}

// SubmitRevealedShare is a paid mutator transaction binding the contract method 0xc9396bf0.
//
// Solidity: function submitRevealedShare(bytes12 roundId, uint16 participantIndex, uint256 shareValue, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitRevealedShare(roundId [12]byte, participantIndex uint16, shareValue *big.Int, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitRevealedShare(&_DKGManager.TransactOpts, roundId, participantIndex, shareValue, proof, input)
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
	RoundId         [12]byte
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
// Solidity: event CiphertextSubmitted(bytes12 indexed roundId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) FilterCiphertextSubmitted(opts *bind.FilterOpts, roundId [][12]byte, ciphertextIndex []uint16, submitter []common.Address) (*DKGManagerCiphertextSubmittedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "CiphertextSubmitted", roundIdRule, ciphertextIndexRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerCiphertextSubmittedIterator{contract: _DKGManager.contract, event: "CiphertextSubmitted", logs: logs, sub: sub}, nil
}

// WatchCiphertextSubmitted is a free log subscription operation binding the contract event 0xa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf.
//
// Solidity: event CiphertextSubmitted(bytes12 indexed roundId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
func (_DKGManager *DKGManagerFilterer) WatchCiphertextSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerCiphertextSubmitted, roundId [][12]byte, ciphertextIndex []uint16, submitter []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "CiphertextSubmitted", roundIdRule, ciphertextIndexRule, submitterRule)
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
// Solidity: event CiphertextSubmitted(bytes12 indexed roundId, uint16 indexed ciphertextIndex, address indexed submitter, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
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
	RoundId             [12]byte
	Contributor         common.Address
	ContributorIndex    uint16
	CommitmentsHash     [32]byte
	EncryptedSharesHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterContributionSubmitted is a free log retrieval operation binding the contract event 0x8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5.
//
// Solidity: event ContributionSubmitted(bytes12 indexed roundId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
func (_DKGManager *DKGManagerFilterer) FilterContributionSubmitted(opts *bind.FilterOpts, roundId [][12]byte, contributor []common.Address) (*DKGManagerContributionSubmittedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "ContributionSubmitted", roundIdRule, contributorRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerContributionSubmittedIterator{contract: _DKGManager.contract, event: "ContributionSubmitted", logs: logs, sub: sub}, nil
}

// WatchContributionSubmitted is a free log subscription operation binding the contract event 0x8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5.
//
// Solidity: event ContributionSubmitted(bytes12 indexed roundId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
func (_DKGManager *DKGManagerFilterer) WatchContributionSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerContributionSubmitted, roundId [][12]byte, contributor []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "ContributionSubmitted", roundIdRule, contributorRule)
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
// Solidity: event ContributionSubmitted(bytes12 indexed roundId, address indexed contributor, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash)
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
	RoundId         [12]byte
	CiphertextIndex uint16
	CombineHash     [32]byte
	Plaintext       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecryptionCombined is a free log retrieval operation binding the contract event 0xf00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336.
//
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) FilterDecryptionCombined(opts *bind.FilterOpts, roundId [][12]byte, ciphertextIndex []uint16) (*DKGManagerDecryptionCombinedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "DecryptionCombined", roundIdRule, ciphertextIndexRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerDecryptionCombinedIterator{contract: _DKGManager.contract, event: "DecryptionCombined", logs: logs, sub: sub}, nil
}

// WatchDecryptionCombined is a free log subscription operation binding the contract event 0xf00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336.
//
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) WatchDecryptionCombined(opts *bind.WatchOpts, sink chan<- *DKGManagerDecryptionCombined, roundId [][12]byte, ciphertextIndex []uint16) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var ciphertextIndexRule []interface{}
	for _, ciphertextIndexItem := range ciphertextIndex {
		ciphertextIndexRule = append(ciphertextIndexRule, ciphertextIndexItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "DecryptionCombined", roundIdRule, ciphertextIndexRule)
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
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext)
func (_DKGManager *DKGManagerFilterer) ParseDecryptionCombined(log types.Log) (*DKGManagerDecryptionCombined, error) {
	event := new(DKGManagerDecryptionCombined)
	if err := _DKGManager.contract.UnpackLog(event, "DecryptionCombined", log); err != nil {
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
	RoundId          [12]byte
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	DeltaHash        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPartialDecryptionSubmitted is a free log retrieval operation binding the contract event 0x39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
func (_DKGManager *DKGManagerFilterer) FilterPartialDecryptionSubmitted(opts *bind.FilterOpts, roundId [][12]byte, participant []common.Address) (*DKGManagerPartialDecryptionSubmittedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "PartialDecryptionSubmitted", roundIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerPartialDecryptionSubmittedIterator{contract: _DKGManager.contract, event: "PartialDecryptionSubmitted", logs: logs, sub: sub}, nil
}

// WatchPartialDecryptionSubmitted is a free log subscription operation binding the contract event 0x39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22.
//
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
func (_DKGManager *DKGManagerFilterer) WatchPartialDecryptionSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerPartialDecryptionSubmitted, roundId [][12]byte, participant []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "PartialDecryptionSubmitted", roundIdRule, participantRule)
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
// Solidity: event PartialDecryptionSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, uint16 ciphertextIndex, bytes32 deltaHash)
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
	RoundId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistrationClosed is a free log retrieval operation binding the contract event 0xca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f1358.
//
// Solidity: event RegistrationClosed(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) FilterRegistrationClosed(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerRegistrationClosedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RegistrationClosed", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRegistrationClosedIterator{contract: _DKGManager.contract, event: "RegistrationClosed", logs: logs, sub: sub}, nil
}

// WatchRegistrationClosed is a free log subscription operation binding the contract event 0xca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f1358.
//
// Solidity: event RegistrationClosed(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) WatchRegistrationClosed(opts *bind.WatchOpts, sink chan<- *DKGManagerRegistrationClosed, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RegistrationClosed", roundIdRule)
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
// Solidity: event RegistrationClosed(bytes12 indexed roundId)
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
	RoundId                 [12]byte
	NewSeedBlock            uint64
	NewRegistrationDeadline uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRegistrationExtended is a free log retrieval operation binding the contract event 0x9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c.
//
// Solidity: event RegistrationExtended(bytes12 indexed roundId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) FilterRegistrationExtended(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerRegistrationExtendedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RegistrationExtended", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRegistrationExtendedIterator{contract: _DKGManager.contract, event: "RegistrationExtended", logs: logs, sub: sub}, nil
}

// WatchRegistrationExtended is a free log subscription operation binding the contract event 0x9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c.
//
// Solidity: event RegistrationExtended(bytes12 indexed roundId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) WatchRegistrationExtended(opts *bind.WatchOpts, sink chan<- *DKGManagerRegistrationExtended, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RegistrationExtended", roundIdRule)
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
// Solidity: event RegistrationExtended(bytes12 indexed roundId, uint64 newSeedBlock, uint64 newRegistrationDeadline)
func (_DKGManager *DKGManagerFilterer) ParseRegistrationExtended(log types.Log) (*DKGManagerRegistrationExtended, error) {
	event := new(DKGManagerRegistrationExtended)
	if err := _DKGManager.contract.UnpackLog(event, "RegistrationExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerRevealedShareSubmittedIterator is returned from FilterRevealedShareSubmitted and is used to iterate over the raw logs and unpacked data for RevealedShareSubmitted events raised by the DKGManager contract.
type DKGManagerRevealedShareSubmittedIterator struct {
	Event *DKGManagerRevealedShareSubmitted // Event containing the contract specifics and raw log

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
func (it *DKGManagerRevealedShareSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRevealedShareSubmitted)
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
		it.Event = new(DKGManagerRevealedShareSubmitted)
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
func (it *DKGManagerRevealedShareSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRevealedShareSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRevealedShareSubmitted represents a RevealedShareSubmitted event raised by the DKGManager contract.
type DKGManagerRevealedShareSubmitted struct {
	RoundId          [12]byte
	Participant      common.Address
	ParticipantIndex uint16
	ShareHash        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRevealedShareSubmitted is a free log retrieval operation binding the contract event 0x5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac.
//
// Solidity: event RevealedShareSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, bytes32 shareHash)
func (_DKGManager *DKGManagerFilterer) FilterRevealedShareSubmitted(opts *bind.FilterOpts, roundId [][12]byte, participant []common.Address) (*DKGManagerRevealedShareSubmittedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RevealedShareSubmitted", roundIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRevealedShareSubmittedIterator{contract: _DKGManager.contract, event: "RevealedShareSubmitted", logs: logs, sub: sub}, nil
}

// WatchRevealedShareSubmitted is a free log subscription operation binding the contract event 0x5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac.
//
// Solidity: event RevealedShareSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, bytes32 shareHash)
func (_DKGManager *DKGManagerFilterer) WatchRevealedShareSubmitted(opts *bind.WatchOpts, sink chan<- *DKGManagerRevealedShareSubmitted, roundId [][12]byte, participant []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RevealedShareSubmitted", roundIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRevealedShareSubmitted)
				if err := _DKGManager.contract.UnpackLog(event, "RevealedShareSubmitted", log); err != nil {
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

// ParseRevealedShareSubmitted is a log parse operation binding the contract event 0x5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac.
//
// Solidity: event RevealedShareSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, bytes32 shareHash)
func (_DKGManager *DKGManagerFilterer) ParseRevealedShareSubmitted(log types.Log) (*DKGManagerRevealedShareSubmitted, error) {
	event := new(DKGManagerRevealedShareSubmitted)
	if err := _DKGManager.contract.UnpackLog(event, "RevealedShareSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerRoundAbortedIterator is returned from FilterRoundAborted and is used to iterate over the raw logs and unpacked data for RoundAborted events raised by the DKGManager contract.
type DKGManagerRoundAbortedIterator struct {
	Event *DKGManagerRoundAborted // Event containing the contract specifics and raw log

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
func (it *DKGManagerRoundAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRoundAborted)
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
		it.Event = new(DKGManagerRoundAborted)
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
func (it *DKGManagerRoundAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRoundAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRoundAborted represents a RoundAborted event raised by the DKGManager contract.
type DKGManagerRoundAborted struct {
	RoundId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoundAborted is a free log retrieval operation binding the contract event 0x97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e410.
//
// Solidity: event RoundAborted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) FilterRoundAborted(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerRoundAbortedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RoundAborted", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRoundAbortedIterator{contract: _DKGManager.contract, event: "RoundAborted", logs: logs, sub: sub}, nil
}

// WatchRoundAborted is a free log subscription operation binding the contract event 0x97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e410.
//
// Solidity: event RoundAborted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) WatchRoundAborted(opts *bind.WatchOpts, sink chan<- *DKGManagerRoundAborted, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RoundAborted", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRoundAborted)
				if err := _DKGManager.contract.UnpackLog(event, "RoundAborted", log); err != nil {
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

// ParseRoundAborted is a log parse operation binding the contract event 0x97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e410.
//
// Solidity: event RoundAborted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) ParseRoundAborted(log types.Log) (*DKGManagerRoundAborted, error) {
	event := new(DKGManagerRoundAborted)
	if err := _DKGManager.contract.UnpackLog(event, "RoundAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerRoundCreatedIterator is returned from FilterRoundCreated and is used to iterate over the raw logs and unpacked data for RoundCreated events raised by the DKGManager contract.
type DKGManagerRoundCreatedIterator struct {
	Event *DKGManagerRoundCreated // Event containing the contract specifics and raw log

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
func (it *DKGManagerRoundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRoundCreated)
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
		it.Event = new(DKGManagerRoundCreated)
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
func (it *DKGManagerRoundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRoundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRoundCreated represents a RoundCreated event raised by the DKGManager contract.
type DKGManagerRoundCreated struct {
	RoundId          [12]byte
	Organizer        common.Address
	SeedBlock        uint64
	LotteryThreshold *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoundCreated is a free log retrieval operation binding the contract event 0xcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d7.
//
// Solidity: event RoundCreated(bytes12 indexed roundId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) FilterRoundCreated(opts *bind.FilterOpts, roundId [][12]byte, organizer []common.Address) (*DKGManagerRoundCreatedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var organizerRule []interface{}
	for _, organizerItem := range organizer {
		organizerRule = append(organizerRule, organizerItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RoundCreated", roundIdRule, organizerRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRoundCreatedIterator{contract: _DKGManager.contract, event: "RoundCreated", logs: logs, sub: sub}, nil
}

// WatchRoundCreated is a free log subscription operation binding the contract event 0xcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d7.
//
// Solidity: event RoundCreated(bytes12 indexed roundId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) WatchRoundCreated(opts *bind.WatchOpts, sink chan<- *DKGManagerRoundCreated, roundId [][12]byte, organizer []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var organizerRule []interface{}
	for _, organizerItem := range organizer {
		organizerRule = append(organizerRule, organizerItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RoundCreated", roundIdRule, organizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRoundCreated)
				if err := _DKGManager.contract.UnpackLog(event, "RoundCreated", log); err != nil {
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

// ParseRoundCreated is a log parse operation binding the contract event 0xcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d7.
//
// Solidity: event RoundCreated(bytes12 indexed roundId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold)
func (_DKGManager *DKGManagerFilterer) ParseRoundCreated(log types.Log) (*DKGManagerRoundCreated, error) {
	event := new(DKGManagerRoundCreated)
	if err := _DKGManager.contract.UnpackLog(event, "RoundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerRoundEvictedIterator is returned from FilterRoundEvicted and is used to iterate over the raw logs and unpacked data for RoundEvicted events raised by the DKGManager contract.
type DKGManagerRoundEvictedIterator struct {
	Event *DKGManagerRoundEvicted // Event containing the contract specifics and raw log

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
func (it *DKGManagerRoundEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRoundEvicted)
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
		it.Event = new(DKGManagerRoundEvicted)
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
func (it *DKGManagerRoundEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRoundEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRoundEvicted represents a RoundEvicted event raised by the DKGManager contract.
type DKGManagerRoundEvicted struct {
	RoundId [12]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoundEvicted is a free log retrieval operation binding the contract event 0x98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef26.
//
// Solidity: event RoundEvicted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) FilterRoundEvicted(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerRoundEvictedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RoundEvicted", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRoundEvictedIterator{contract: _DKGManager.contract, event: "RoundEvicted", logs: logs, sub: sub}, nil
}

// WatchRoundEvicted is a free log subscription operation binding the contract event 0x98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef26.
//
// Solidity: event RoundEvicted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) WatchRoundEvicted(opts *bind.WatchOpts, sink chan<- *DKGManagerRoundEvicted, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RoundEvicted", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRoundEvicted)
				if err := _DKGManager.contract.UnpackLog(event, "RoundEvicted", log); err != nil {
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

// ParseRoundEvicted is a log parse operation binding the contract event 0x98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef26.
//
// Solidity: event RoundEvicted(bytes12 indexed roundId)
func (_DKGManager *DKGManagerFilterer) ParseRoundEvicted(log types.Log) (*DKGManagerRoundEvicted, error) {
	event := new(DKGManagerRoundEvicted)
	if err := _DKGManager.contract.UnpackLog(event, "RoundEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerRoundFinalizedIterator is returned from FilterRoundFinalized and is used to iterate over the raw logs and unpacked data for RoundFinalized events raised by the DKGManager contract.
type DKGManagerRoundFinalizedIterator struct {
	Event *DKGManagerRoundFinalized // Event containing the contract specifics and raw log

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
func (it *DKGManagerRoundFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerRoundFinalized)
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
		it.Event = new(DKGManagerRoundFinalized)
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
func (it *DKGManagerRoundFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerRoundFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerRoundFinalized represents a RoundFinalized event raised by the DKGManager contract.
type DKGManagerRoundFinalized struct {
	RoundId                  [12]byte
	AggregateCommitmentsHash [32]byte
	CollectivePublicKeyHash  [32]byte
	ShareCommitmentHash      [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRoundFinalized is a free log retrieval operation binding the contract event 0x5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a72455.
//
// Solidity: event RoundFinalized(bytes12 indexed roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) FilterRoundFinalized(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerRoundFinalizedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "RoundFinalized", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerRoundFinalizedIterator{contract: _DKGManager.contract, event: "RoundFinalized", logs: logs, sub: sub}, nil
}

// WatchRoundFinalized is a free log subscription operation binding the contract event 0x5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a72455.
//
// Solidity: event RoundFinalized(bytes12 indexed roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) WatchRoundFinalized(opts *bind.WatchOpts, sink chan<- *DKGManagerRoundFinalized, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "RoundFinalized", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerRoundFinalized)
				if err := _DKGManager.contract.UnpackLog(event, "RoundFinalized", log); err != nil {
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

// ParseRoundFinalized is a log parse operation binding the contract event 0x5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a72455.
//
// Solidity: event RoundFinalized(bytes12 indexed roundId, bytes32 aggregateCommitmentsHash, bytes32 collectivePublicKeyHash, bytes32 shareCommitmentHash)
func (_DKGManager *DKGManagerFilterer) ParseRoundFinalized(log types.Log) (*DKGManagerRoundFinalized, error) {
	event := new(DKGManagerRoundFinalized)
	if err := _DKGManager.contract.UnpackLog(event, "RoundFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGManagerSecretReconstructedIterator is returned from FilterSecretReconstructed and is used to iterate over the raw logs and unpacked data for SecretReconstructed events raised by the DKGManager contract.
type DKGManagerSecretReconstructedIterator struct {
	Event *DKGManagerSecretReconstructed // Event containing the contract specifics and raw log

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
func (it *DKGManagerSecretReconstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGManagerSecretReconstructed)
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
		it.Event = new(DKGManagerSecretReconstructed)
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
func (it *DKGManagerSecretReconstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGManagerSecretReconstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGManagerSecretReconstructed represents a SecretReconstructed event raised by the DKGManager contract.
type DKGManagerSecretReconstructed struct {
	RoundId                 [12]byte
	DisclosureHash          [32]byte
	ReconstructedSecretHash [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSecretReconstructed is a free log retrieval operation binding the contract event 0xbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b482.
//
// Solidity: event SecretReconstructed(bytes12 indexed roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash)
func (_DKGManager *DKGManagerFilterer) FilterSecretReconstructed(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerSecretReconstructedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "SecretReconstructed", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerSecretReconstructedIterator{contract: _DKGManager.contract, event: "SecretReconstructed", logs: logs, sub: sub}, nil
}

// WatchSecretReconstructed is a free log subscription operation binding the contract event 0xbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b482.
//
// Solidity: event SecretReconstructed(bytes12 indexed roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash)
func (_DKGManager *DKGManagerFilterer) WatchSecretReconstructed(opts *bind.WatchOpts, sink chan<- *DKGManagerSecretReconstructed, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "SecretReconstructed", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGManagerSecretReconstructed)
				if err := _DKGManager.contract.UnpackLog(event, "SecretReconstructed", log); err != nil {
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

// ParseSecretReconstructed is a log parse operation binding the contract event 0xbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b482.
//
// Solidity: event SecretReconstructed(bytes12 indexed roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash)
func (_DKGManager *DKGManagerFilterer) ParseSecretReconstructed(log types.Log) (*DKGManagerSecretReconstructed, error) {
	event := new(DKGManagerSecretReconstructed)
	if err := _DKGManager.contract.UnpackLog(event, "SecretReconstructed", log); err != nil {
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
	RoundId [12]byte
	Seed    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSeedResolved is a free log retrieval operation binding the contract event 0xc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652.
//
// Solidity: event SeedResolved(bytes12 indexed roundId, bytes32 seed)
func (_DKGManager *DKGManagerFilterer) FilterSeedResolved(opts *bind.FilterOpts, roundId [][12]byte) (*DKGManagerSeedResolvedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "SeedResolved", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerSeedResolvedIterator{contract: _DKGManager.contract, event: "SeedResolved", logs: logs, sub: sub}, nil
}

// WatchSeedResolved is a free log subscription operation binding the contract event 0xc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652.
//
// Solidity: event SeedResolved(bytes12 indexed roundId, bytes32 seed)
func (_DKGManager *DKGManagerFilterer) WatchSeedResolved(opts *bind.WatchOpts, sink chan<- *DKGManagerSeedResolved, roundId [][12]byte) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "SeedResolved", roundIdRule)
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
// Solidity: event SeedResolved(bytes12 indexed roundId, bytes32 seed)
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
	RoundId [12]byte
	Claimer common.Address
	Slot    uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlotClaimed is a free log retrieval operation binding the contract event 0x80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b.
//
// Solidity: event SlotClaimed(bytes12 indexed roundId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) FilterSlotClaimed(opts *bind.FilterOpts, roundId [][12]byte, claimer []common.Address) (*DKGManagerSlotClaimedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DKGManager.contract.FilterLogs(opts, "SlotClaimed", roundIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &DKGManagerSlotClaimedIterator{contract: _DKGManager.contract, event: "SlotClaimed", logs: logs, sub: sub}, nil
}

// WatchSlotClaimed is a free log subscription operation binding the contract event 0x80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b.
//
// Solidity: event SlotClaimed(bytes12 indexed roundId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) WatchSlotClaimed(opts *bind.WatchOpts, sink chan<- *DKGManagerSlotClaimed, roundId [][12]byte, claimer []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DKGManager.contract.WatchLogs(opts, "SlotClaimed", roundIdRule, claimerRule)
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
// Solidity: event SlotClaimed(bytes12 indexed roundId, address indexed claimer, uint16 slot)
func (_DKGManager *DKGManagerFilterer) ParseSlotClaimed(log types.Log) (*DKGManagerSlotClaimed, error) {
	event := new(DKGManagerSlotClaimed)
	if err := _DKGManager.contract.UnpackLog(event, "SlotClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
