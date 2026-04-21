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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealSubmitVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealShareVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SHARE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SUBMIT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROUND_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRound\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCiphertextHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptionPolicy\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delta\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlaintext\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealShareVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealSubmitVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RevealedShareRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Round\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RoundPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"decryptionPolicy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.DecryptionPolicy\",\"components\":[{\"name\":\"ownerOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxDecryptions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"notBeforeBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notBeforeTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"notAfterTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.RoundStatus\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"revealedShareCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reconstructSecret\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roundNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitCiphertext\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitment0X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment0Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CiphertextSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"c1x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c1y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2x\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"c2y\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintext\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevealedShareSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundAborted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundCreated\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundEvicted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundFinalized\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecretReconstructed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CiphertextNotSubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DecryptionNotYetAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisclosureDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRevealedShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCiphertext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecryptionPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReconstruction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRevealedShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShareCount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LagrangeMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6101a0806040523461029657610100816161f08038038091610021828561029a565b8339810103126102965780519063ffffffff82169182810361029657610049602083016102d1565b610055604084016102d1565b610061606085016102d1565b9061006e608086016102d1565b9261007b60a087016102d1565b9461009460e061008d60c08a016102d1565b98016102d1565b9763ffffffff461603610287576001600160a01b03821615610278576001600160a01b038316158015610267575b8015610256575b8015610245575b8015610234575b8015610223575b6102145763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261011a60388261029a565b5190201660c05260e0526101005261012052610140526101605261018052604051615f0a90816102e6823960805181611417015260a05181818161037e0152818161224c01528181613e0c015281816147220152615b7f015260c051818181610b4a0152613ec2015260e05181818161040a01528181610bd001526144e201526101005181818161109c0152818161146001526119ae015261012051818181610c19015281816114c20152612be3015261014051818181610e70015281816127da01526135900152610160518181816103c1015281816118b10152611c93015261018051818181610b8d01528181611acb01526130240152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038816156100de565b506001600160a01b038716156100d7565b506001600160a01b038616156100d0565b506001600160a01b038516156100c9565b506001600160a01b038416156100c2565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b038211908210176102bd57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102965756fe60806040526004361015610011575f80fd5b5f3560e01c8063058994a11461027457806306433b1b1461026f578063070c74921461026a578063074a75e1146102655780630b1451f0146102605780630e2c53f71461025b5780633353ec6e14610256578063349181a214610251578063373877a61461024c578063415a1b86146102475780634554c0be14610242578063510ba2df1461023d57806353d721841461023857806356664d15146102335780635ddd06261461022e57806363f314cd14610229578063669a76a9146102245780636759e0e11461021f57806370f2469b1461021a57806372517b4b14610215578063802ae2311461021057806385e1f4d01461020b5780638dc1f53a1461020657806393c3d3a814610201578063a9c4b25f146101fc578063b18730c2146101f7578063b58aab90146101f2578063bf192209146101ed578063bfa78991146101e8578063c2440e16146101e3578063c9396bf0146101de578063ca3c0458146101d9578063d3720aac146101d4578063d6c29c9e146101cf578063d9933767146101ca578063f4e34945146101c5578063fe1604b5146101c05763fe234897146101bb575f80fd5b6127fe565b6127bb565b6126d6565b612148565b6120ab565b611fa2565b611f1a565b611b0e565b611aa6565b6119ff565b61198f565b6118f4565b61188c565b6114e6565b6114a3565b61143b565b6113fb565b610eb3565b610e4b565b610d1c565b610c5c565b610bf4565b610bb1565b610b6e565b610b2e565b610a1a565b6109b9565b610921565b6108fc565b6108a3565b61078b565b610750565b6106c3565b610482565b6103e5565b6103a2565b61035f565b6102c1565b600435906001600160a01b03198216820361029057565b5f80fd5b9181601f84011215610290578235916001600160401b038311610290576020838186019501011161029057565b346102905760e0366003190112610290576102da610279565b602435604435916064356084356001600160401b03811161029057610303903690600401610294565b60a4929192356001600160401b03811161029057610325903690600401610294565b93909260c435976001600160401b0389116102905761034b610353993690600401610294565b989097612b14565b005b5f91031261029057565b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e575b50604051908152602090f35b610470915060203d602011610476575b610468818361297f565b810190612ecb565b5f610442565b503d61045e565b6129ea565b346102905760203660031901126102905761049b610279565b6104b7816001600160a01b0319165f52602260205260405f2090565b6001600160a01b036104d082546001600160a01b031690565b16156106b457600481019081549260016104ea8560ff1690565b6104f38161258c565b036106a557600782015461ffff16600183018054909161ffff61051e601084901c82165b61ffff1690565b9116146106a5576001600160401b03605082901c1695438710156106a5576105a861059861057e61055c6105b4946001600160401b039060481c1690565b99610578610571604088901c61ffff16610517565b809c612eda565b90612eda565b6105926001600160401b0343169a8b612efa565b99612efa565b9260901c6001600160401b031690565b6001600160401b031690565b6001600160401b03821610156106965761066b8161069193610635897f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c995f60056001600160a01b03199b01559070ffffffffffffffff00000000000000000082549160481b169070ffffffffffffffff0000000000000000001916179055565b9071ffffffffffffffff0000000000000000000082549160501b169071ffffffffffffffff000000000000000000001916179055565b6040519384931695839092916001600160401b0360209181604085019616845216910152565b0390a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b6328ad4a9560e21b5f5260045ffd5b346102905760c0366003190112610290576106dc610279565b602435604435916064356001600160401b03811161029057610702903690600401610294565b906084356001600160401b03811161029057610722903690600401610294565b92909160a435966001600160401b03881161029057610748610353983690600401610294565b979096612f6b565b3461029057602036600319011261029057604061077361076e610279565b613244565b6107898251809260208091805184520151910152565bf35b34610290576020366003190112610290576107a4610279565b6107c0816001600160a01b0319165f52602260205260405f2090565b80546001600160a01b031680156106b4576001600160a01b0316330361087a57600401906107ef825460ff1690565b6107f88161258c565b60038114908115610865575b8115610851575b506106a55761082a6001600160a01b031992600460ff19825416179055565b167f97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e4105f80a2005b6004915061085e8161258c565b145f61080b565b90506108708161258c565b6005811490610804565b6282b42960e81b5f5260045ffd5b61ffff81160361029057565b602435906108a182610888565b565b346102905760403660031901126102905760206108f36108c1610279565b6001600160a01b0319602435916108d783610888565b165f52602d835260405f209061ffff165f5260205260405f2090565b54604051908152f35b34610290575f3660031901126102905760206001600160401b035f5416604051908152f35b34610290576020366003190112610290576001600160a01b0319610943610279565b61094b613290565b50165f52602260205260c0610965600260405f20016132c0565b61078960405180926001600160401b0360a0809280511515855261ffff6020820151166020860152826040820151166040860152826060820151166060860152826080820151166080860152015116910152565b346102905760403660031901126102905760206108f36109d7610279565b6001600160a01b0319602435916109ed83610888565b165f52602a835260405f209061ffff165f5260205260405f2090565b6001600160a01b0381160361029057565b3461029057604036600319011261029057610b2a610a97610a39610279565b6001600160a01b031960243591610a4f83610a09565b5f6080604051610a5e816128d6565b8281528260208201528260408201528260608201520152165f52602960205260405f20906001600160a01b03165f5260205260405f2090565b60ff600360405192610aa8846128d6565b61ffff81546001600160a01b038116865260a01c1660208501526001810154604085015260028101546060850152015416151560808201526040519182918291909160808060a08301946001600160a01b03815116845261ffff6020820151166020850152604081015160408501526060810151606085015201511515910152565b0390f35b34610290575f36600319011261029057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e5750604051908152602090f35b346102905760403660031901126102905760206001610cae610c7c610279565b6001600160a01b031960243591610c9283610888565b165f526028845260405f209061ffff165f5260205260405f2090565b0154604051908152f35b91909160c060a060e08301946001600160a01b03815116845261ffff602082015116602085015261ffff604082015116604085015260608101516060850152610d136080820151608086019060208091805184520151910152565b01511515910152565b3461029057606036600319011261029057610b2a610dc7610d3b610279565b610db160243591610d4b83610a09565b6001600160a01b031960443591610d6183610888565b5f60a0604051610d70816128f6565b828152826020820152826040820152826060820152610d8d61320e565b60808201520152165f52602660205260405f209061ffff165f5260205260405f2090565b906001600160a01b03165f5260205260405f2090565b610e3f610e36600460405193610ddc856128f6565b610e13610e0882546001600160a01b038116885261ffff8160a01c16602089015261ffff9060b01c1690565b61ffff166040870152565b60018101546060860152610e2960028201613226565b6080860152015460ff1690565b151560a0830152565b60405191829182610cb8565b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e5750604051908152602090f35b346102905760c036600319011261029057610ecc610279565b602435610ed881610888565b604435610ee481610888565b6064356084356001600160401b03811161029057610f06903690600401610294565b60a4939193356001600160401b03811161029057610f28903690600401610294565b90610f45886001600160a01b0319165f52602260205260405f2090565b956001600160a01b03610f5f88546001600160a01b031690565b16156106b4576003610f75600489015460ff1690565b610f7e8161258c565b036106a557610fc9610fc5610fbe610fa88c6001600160a01b0319165f52602360205260405f2090565b336001600160a01b03165f5260205260405f2090565b5460ff1690565b1590565b6113ec5761ffff881693841580156113cf575b80156113c3575b80156113b4575b80156113ac575b61139e576110456001600160a01b0361102e61101f8d6001600160a01b0319165f52602460205260405f2090565b6110288d61335a565b9061336e565b9290923393546001600160a01b039160031b1c1690565b16036113145761108e60046110868c610db18a61107533936001600160a01b0319165f52602660205260405f2090565b9061ffff165f5260205260405f2090565b015460ff1690565b61138f576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f9285926110ea60405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d5761110793611375575b50810190613388565b9061112886611075896001600160a01b0319165f52602a60205260405f2090565b5482516111496111388a60a01c90565b6bffffffffffffffffffffffff1690565b1491821592611366575b50811561135d575b8115611323575b506113145760c081018051909160e001906111976111a583516040519283916020830195869091604092825260208201520190565b03601f19810183528261297f565b5190208403611314577f39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22946112a29260036007936111fe8b610db18961107533936001600160a01b0319165f52602660205260405f2090565b805460a08c901b61ffff60a01b167fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff9091161760b089901b77ffff000000000000000000000000000000000000000000001617815560048101805460ff1916600117905592516002840155519101550180546112859060201c61ffff166133da565b6133da565b65ffff0000000082549160201b169065ffff000000001916179055565b6112e46112c582611075886001600160a01b0319165f52602760205260405f2090565b6112d4611280825461ffff1690565b61ffff1661ffff19825416179055565b6040805161ffff958616815294909116602085015283015233926001600160a01b0319169180606081015b0390a3005b63d1fed5fd60e01b5f5260045ffd5b9050608082015161119761135260a085015b516040805160208101958652908101919091529182906060820190565b51902014155f611162565b8015915061115b565b6020840151141591505f611153565b806113835f6113899361297f565b80610355565b5f6110fe565b633466526160e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b508615610ff1565b5061010061ffff871611610fea565b5061ffff861615610fe3565b5060018801546113e59060101c61ffff16610517565b8511610fdc565b63965c290d60e01b5f5260045ffd5b34610290575f36600319011261029057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e5750604051908152602090f35b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102905760c0366003190112610290576114ff610279565b60243561150b81610888565b6044356064359160843560a43593611535866001600160a01b0319165f52602260205260405f2090565b9261154784546001600160a01b031690565b956001600160a01b038716156106b4576003611567600487015460ff1690565b6115708161258c565b036106a55761ffff82169687158015611881575b61185f57611595610fc585896152f9565b801561186e575b61185f576115ac600287016132c0565b906115b78251151590565b908161184b575b5061183c5760408101516001600160401b03168015159081611820575b506117fe576115f76105a860608301516001600160401b031690565b801515908161180d575b506117fe5761161d6105a860808301516001600160401b031690565b80151590816117eb575b506117c9576116436105a860a08301516001600160401b031690565b80151590816117d8575b506117c957602001516116639061ffff16610517565b80151590816117a8575b5061179957611692826110758a6001600160a01b0319165f52602d60205260405f2090565b5461178a5761175560077fa5a7194c3409f675784ea2429410513d4d52c73d5fd751d44ed21da06cc643cf968a611717611785966110758c8b6116fb8c6111978c6040519485936020850197889094939260609260808301968352602083015260408201520152565b519020936001600160a01b0319165f52602d60205260405f2090565b550180546117309060401c61ffff1660010161ffff1690565b69ffff000000000000000082549160401b169069ffff00000000000000001916179055565b6040519384936001600160a01b0319339a1697859094939260609260808301968352602083015260408201520152565b0390a4005b6316feb18560e11b5f5260045ffd5b63464e67af60e01b5f5260045ffd5b905061ffff6117c0600788015461ffff9060401c1690565b1610155f61166d565b630410ff2960e31b5f5260045ffd5b90506001600160401b034216115f61164d565b90506001600160401b034316115f611627565b633deac39560e01b5f5260045ffd5b90506001600160401b034216105f611601565b6001600160401b031690506001600160401b034316105f6115db565b6330cd747160e01b5f5260045ffd5b6001600160a01b031690503314155f6115be565b634c4d29cd60e11b5f5260045ffd5b5061187c610fc583876152f9565b61159c565b506101008811611584565b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e5750604051908152602090f35b346102905760e03660031901126102905761190d610279565b60243561191981610888565b604435916064356084356001600160401b0381116102905761193f903690600401610294565b60a4929192356001600160401b03811161029057611961903690600401610294565b93909260c435976001600160401b03891161029057611987610353993690600401610294565b989097613466565b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b6001600160401b0381160361029057565b60e43590811515820361029057565b3590811515820361029057565b34610290576101c036600319011261029057600435611a1d81610888565b60243590611a2a82610888565b604435611a3681610888565b606435611a4281610888565b608435611a4e81610888565b60a43590611a5b826119d2565b60c43592611a68846119d2565b611a706119e3565b9460c03661010319011261029057610b2a97611a8b97613d6f565b6040516001600160a01b031990911681529081906020820190565b34610290575f3660031901126102905760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d575f9161044e5750604051908152602090f35b346102905760a036600319011261029057611b27610279565b602435611b3381610888565b6044356064356001600160401b03811161029057611b55903690600401610294565b6084356001600160401b03811161029057611b74903690600401610294565b90611b91876001600160a01b0319165f52602260205260405f2090565b936001600160a01b03611bab86546001600160a01b031690565b16156106b45760018501549360ff60d086901c1615611ec9576003611bd4600488015460ff1690565b611bdd8161258c565b036106a557611c07610fc5610fbe610fa88c6001600160a01b0319165f52602360205260405f2090565b6113ec5761ffff8816948515908115611eb0575b508015611ea8575b611e9957611c5a6001600160a01b0361102e611c518c6001600160a01b0319165f52602460205260405f2090565b6110288c61335a565b160361131457611c85600361108633610db18d6001600160a01b0319165f52602960205260405f2090565b611e8a576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f928592611ce160405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d57611cfe93611e76575b50810190614316565b611d1e85611075886001600160a01b0319165f52602a60205260405f2090565b54908051611d2f6111388960a01c90565b1492831593611e67575b508215611e58575b8215611e4f575b8215611e2b575b50506113145781611e01600761130f937f5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac956001611da333610db18c6001600160a01b0319165f52602960205260405f2090565b805461ffff60a01b191660a08b901b61ffff60a01b1617815560038101805460ff191660011790550155018054611de09060301c61ffff166133da565b67ffff00000000000082549160301b169067ffff0000000000001916179055565b6040519182916001600160a01b031933971695836020909392919361ffff60408201951681520152565b909150611197611e4361133560608401519360800190565b51902014155f80611d4f565b81159250611d48565b60408101518514159250611d41565b6020820151141592505f611d39565b806113835f611e849361297f565b5f611cf5565b63a89ac15160e01b5f5260045ffd5b639eae062d60e01b5f5260045ffd5b508615611c23565b611ec1915060101c61ffff16610517565b85115f611c1b565b630ba0cb2f60e21b5f5260045ffd5b60206040818301928281528451809452019201905f5b818110611efb5750505090565b82516001600160a01b0316845260209384019390920191600101611eee565b34610290576020366003190112610290576001600160a01b0319611f3c610279565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b818110611f8357610b2a85611f778187038261297f565b60405191829182611ed8565b82546001600160a01b0316845260209093019260019283019201611f60565b3461029057604036600319011261029057610b2a612002611fc1610279565b6001600160a01b031960243591611fd783610a09565b611fdf613290565b50165f52602560205260405f20906001600160a01b03165f5260205260405f2090565b612057610e36600460405193612017856128f6565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b6040519182918291909160a08060c08301946001600160a01b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b3461029057610120366003190112610290576120c5610279565b6120cd610894565b6044359160843560643560a43560c4356001600160401b038111610290576120f9903690600401610294565b9160e4356001600160401b03811161029057612119903690600401610294565b95909461010435996001600160401b038b11610290576121406103539b3690600401610294565b9a90996143b9565b3461029057602036600319011261029057612161610279565b61217d816001600160a01b0319165f52602260205260405f2090565b906001600160a01b0361219783546001600160a01b031690565b16156106b457600482018054600184018054949093916121ca610fc56001600160401b03605089901c1660ff8416615a98565b6106a557600782019561ffff6121f36105176121e88a5461ffff1690565b9360101c61ffff1690565b911610156124eb5761221d610fbe610fa8856001600160a01b0319165f52602360205260405f2090565b6124dc57600582018054918215612454575b50506040516313a4120960e31b815233600482015260a0816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561047d576001916060915f91612425575b50015161229681614901565b61229f81614901565b036124165760408051602081019283523360601b6bffffffffffffffffffffffff191691810191909152600691906122da8160548101611197565b5190209101541115612407576123546122f5855461ffff1690565b9461231b33612316856001600160a01b0319165f52602460205260405f2090565b61490b565b61234b61233e33610db1866001600160a01b0319165f52602360205260405f2090565b805460ff19166001179055565b6112d48661494b565b6123b76123ab6001600160a01b0319831695604051877f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b3392806123a3868291909161ffff6020820193169052565b0390a361494b565b935460101c61ffff1690565b9261ffff8085169116146123c757005b6123e1926123d491615b4d565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b612447915060a03d60a01161244d575b61243f818361297f565b81019061489e565b5f61228a565b503d612435565b90915061246c9060481c6001600160401b03166105a8565b804311156124cd57409081156124be578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f8061222f565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b9060e0806108a19361ffff815116845261ffff602082015116602085015261ffff604082015116604085015261253b6060820151606086019061ffff169052565b60808181015161ffff169085015260a0818101516001600160401b03169085015260c0818101516001600160401b03169085015201511515910152565b634e487b7160e01b5f52602160045260245ffd5b6006111561259657565b612578565b9060068210156125965752565b6108a1909291926103006101806103208301956125ce8482516001600160a01b03169052565b6125e0602082015160208601906124fa565b61263a60408201516101208601906001600160401b0360a0809280511515855261ffff6020820151166020860152826040820151166040860152826060820151166060860152826080820151166080860152015116910152565b61264d60608201516101e086019061259b565b60808101516001600160401b031661020085015260a08101516001600160401b031661022085015260c081015161024085015260e081015161026085015261010081015161ffff1661028085015261012081015161ffff166102a085015261014081015161ffff166102c085015261016081015161ffff166102e0850152015161ffff16910152565b3461029057602036600319011261029057610b2a6127af6127aa6126f8610279565b5f61018060405161270881612911565b8281526040516127178161292d565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e08201526020820152612751613290565b60408201528260608201528260808201528260a08201528260c08201528260e08201528261010082015282610120820152826101408201528261016082015201526001600160a01b0319165f52602260205260405f2090565b6149f6565b604051918291826125a8565b34610290575f3660031901126102905760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461029057604036600319011261029057610b2a61286961281d610279565b6001600160a01b03196024359161283383610888565b5f6040805161284181612949565b8281528260208201520152165f52602860205260405f209061ffff165f5260205260405f2090565b60016040519161287883612949565b60ff815461ffff8116855260101c16151560208401520154604082015260405191829182919091604080606083019461ffff81511684526020810151151560208501520151910152565b634e487b7160e01b5f52604160045260245ffd5b60a081019081106001600160401b038211176128f157604052565b6128c2565b60c081019081106001600160401b038211176128f157604052565b6101a081019081106001600160401b038211176128f157604052565b61010081019081106001600160401b038211176128f157604052565b606081019081106001600160401b038211176128f157604052565b604081019081106001600160401b038211176128f157604052565b90601f801991011681019081106001600160401b038211176128f157604052565b908060209392818452848401375f828201840152601f01601f1916010190565b92906129d9906129e795936040865260408601916129a0565b9260208185039101526129a0565b90565b6040513d5f823e3d90fd5b604051906108a16101008361297f565b604051906108a16101a08361297f565b90610120828203126102905780601f830112156102905761012060405192612a3d828561297f565b8391810192831161029057905b828210612a575750505090565b8135815260209182019101612a4a565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b908160051b9180830460201490151715612aa557565b612a7b565b908160011b9180830460021490151715612aa557565b81810292918115918404141715612aa557565b9060018201809211612aa557565b9060808201809211612aa557565b91908201809211612aa557565b9060068110156125965760ff80198354169116179055565b9791959098969296949394612b3b896001600160a01b0319165f52602260205260405f2090565b936001600160a01b03612b5586546001600160a01b031690565b16156106b4576004850191612b6b835460ff1690565b612b748161258c565b60038114612ebc5780612b8860029261258c565b036106a5576007860194612ba2865461ffff9060101c1690565b94600188015493612bbb6105178661ffff9060201c1690565b61ffff881610612ead578e158015612ea5575b8015612e9d575b612e8e576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f928592612c3160405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d57612c4e93612e7a575b50810190612a15565b928351612c5e6111388d60a01c90565b1491821592612e5e575b8215612e3b575b508115612e28575b508015612e1a575b8015612e0c575b8015612dfe575b61131457612cc58a604051612cbc816111978d8d60208401968791606093918352602083015260408201520190565b5190208a614b02565b8060e08401510361131457612cdb6108a0612a8f565b860361131457612d1c95878b612d049361051798612cfd612d11986101000190565b5192614c81565b805460ff19166003179055565b5460101c61ffff1690565b90612d31612d2b610860612a8f565b82612aef565b5f5b838110612d8857505060408051968752602087019390935250508301526001600160a01b031916907f5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a724559080606081015b0390a2565b8060019160061b8301604051612dba816111976020820194602081013590358660209093929193604081019481520152565b519020612df7612ddc8a6001600160a01b0319165f52602a60205260405f2090565b61ffff8460051b8801351661ffff165f5260205260405f2090565b5501612d33565b508760c08301511415612c8d565b508660a08301511415612c86565b508960808301511415612c7f565b6060840151915061ffff1614155f612c77565b909150612e55610517604086015b519260101c61ffff1690565b1415905f612c6f565b91506020840151612e7261ffff8416610517565b141591612c68565b806113835f612e889361297f565b5f612c45565b63c5f680ed60e01b5f5260045ffd5b508c15612bd5565b508b15612bce565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b90816020910312610290575190565b906001600160401b03809116911603906001600160401b038211612aa557565b906001600160401b03809116911601906001600160401b038211612aa557565b9060e0828203126102905780601f830112156102905760405191612f3f60e08461297f565b829060e0810192831161029057905b828210612f5b5750505090565b8135815260209182019101612f4e565b94919296939097612f8e866001600160a01b0319165f52602260205260405f2090565b946001600160a01b03612fa887546001600160a01b031690565b16156106b45760018601549360ff60d086901c1615611ec95760048701986003612fd38b5460ff1690565b612fdc8161258c565b036106a5578b158015613206575b6131f757613004600761ffff99015461ffff9060301c1690565b61301161ffff8816610517565b98899116106131e8576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f92859261307260405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d5761308f936131d4575b50810190612f1a565b91825161309f6111388860a01c90565b148015906131c6575b80156131b8575b80156131aa575b6113145760408301948551106113145760408051602081018b81529181018690526130f291906130e98160608101611197565b51902087614b67565b918260a085015103611314576131086040612a8f565b036113145760408861313f936131338261312a6131389661ffff9060101c1690565b8a51908c614df4565b614f0d565b9160c00190565b51036113145761318f6001600160a01b0319946131827fbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b4829784612d839651614fcb565b805460ff19166005179055565b60405193849316958360209093929193604081019481520152565b5083608084015114156130b6565b5088606084015114156130af565b5084602084015114156130a8565b806113835f6131e29361297f565b5f613086565b63957674fd60e01b5f5260045ffd5b6314141ce560e21b5f5260045ffd5b508615612fea565b6040519061321b82612964565b5f6020838281520152565b9060405161323381612964565b602060018294805484520154910152565b6001600160a01b03199061325661320e565b50165f52602c60205260405f20600181015415613276576129e790613226565b5060405161328381612964565b5f81526001602082015290565b6040519061329d826128f6565b5f60a0838281528260208201528260408201528260608201528260808201520152565b906108a16040516132d0816128f6565b60a061334c6001839661333e61332e825460ff81161515885261ffff8160081c1660208901526001600160401b03808260181c161660408901526001600160401b03808260581c161660608901526001600160401b039060981c1690565b6001600160401b03166080870152565b01546001600160401b031690565b6001600160401b0316910152565b61ffff5f199116019061ffff8211612aa557565b8054821015613383575f5260205f2001905f90565b612a67565b906101a0828203126102905780601f83011215610290576101a0604051926133b0828561297f565b8391810192831161029057905b8282106133ca5750505090565b81358152602091820191016133bd565b61ffff1661ffff8114612aa55760010190565b906080116102905790608090565b909291928311610290579190565b90939293848311610290578411610290578101920390565b9291926001600160401b0382116128f1576040519161344a601f8201601f19166020018461297f565b829481845281830111610290578281602093845f960137010152565b949295919793613488866001600160a01b0319165f52602260205260405f2090565b936001600160a01b036134a286546001600160a01b031690565b16156106b45760036134b8600487015460ff1690565b6134c18161258c565b036106a55761ffff83169a8b15801561378d575b8015613785575b61377657613500846110758a6001600160a01b0319165f52602d60205260405f2090565b549687156137675761ffff9561353761352f876110758d6001600160a01b0319165f52602760205260405f2090565b5461ffff1690565b61354961051760018b015461ffff1690565b978891161061375857613572866110758c6001600160a01b0319165f52602860205260405f2090565b9b6135828d5460ff9060101c1690565b613749576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f9285926135de60405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d576135fa936131d45750810190612f1a565b95865161360a6111388a60a01c90565b1480159061373b575b801561372d575b801561371f575b6113145760408701948551106113145760408051602081018d81529181018b905261365d91906136548160608101611197565b51902089614bc5565b918260a089015103611314576136736064612a8f565b81036113145761368661368d91856133ed565b3691613421565b602081519101200361131457613133826064946136b0976131389751918b61539c565b51036113145761371a826001600160a01b0319936001866136fd7ff00fbf9d648ee3274fc53f9f2eb67f1f6218a6bbc046de320813cdd0244b7336986201000062ff000019825416179055565b015560405193849316958360209093929193604081019481520152565b0390a3565b508860808801511415613621565b508a6060880151141561361a565b508460208801511415613613565b63955c0c4960e01b5f5260045ffd5b63032cddf960e11b5f5260045ffd5b6346f551f560e01b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508a156134dc565b506101008c116134d5565b356129e7816119d2565b356129e781610888565b9081602091031261029057516129e7816119d2565b634e487b7160e01b5f52601260045260245ffd5b81156137df570490565b6137c1565b6001600160401b03166001600160401b038114612aa55760010190565b90604082101561338357600c600183811c810193160290565b91908260c091031261029057604051613832816128f6565b60a0808294613840816119f2565b8452602081013561385081610888565b60208501526040810135613863816119d2565b60408501526060810135613876816119d2565b60608501526080810135613889816119d2565b608085015201359161389a836119d2565b0152565b60068210156125965752565b6139fc60e06108a1936138ce61ffff825116859061ffff1661ffff19825416179055565b6020810151845463ffff0000191660109190911b63ffff0000161784556040810151845465ffff00000000191660209190911b65ffff00000000161784556060810151845467ffff000000000000191660309190911b67ffff000000000000161784556080810151845469ffff0000000000000000191660409190911b69ffff00000000000000001617845560a0810151845471ffffffffffffffff00000000000000000000191660509190911b71ffffffffffffffff00000000000000000000161784556139f56139aa60c08301516001600160401b031690565b85547fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff1660909190911b79ffffffffffffffff00000000000000000000000000000000000016178555565b0151151590565b81547fffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560d01b7aff000000000000000000000000000000000000000000000000000016179055565b6001613b8a60a06108a1948051151560ff801987541691151516178555602081015162ffff0086549160081b169062ffff001916178555613abe613a9660408301516001600160401b031690565b86546affffffffffffffff000000191660189190911b6affffffffffffffff00000016178655565b613b19613ad560608301516001600160401b031690565b86547fffffffffffffffffffffffffff0000000000000000ffffffffffffffffffffff1660589190911b72ffffffffffffffff000000000000000000000016178655565b613b7c613b3060808301516001600160401b031690565b86547fffffffffff0000000000000000ffffffffffffffffffffffffffffffffffffff1660989190911b7affffffffffffffff0000000000000000000000000000000000000016178655565b01516001600160401b031690565b9101906001600160401b03166001600160401b0319825416179055565b9061173061018060076108a194613be4613bc886516001600160a01b031690565b82906001600160a01b03166001600160a01b0319825416179055565b613bf56020860151600183016138aa565b613c06604086015160028301613a48565b613c9860048201613c246060880151613c1e8161258c565b82612afc565b613c5f613c3b60808901516001600160401b031690565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b60a0870151815470ffffffffffffffff000000000000000000191660489190911b70ffffffffffffffff00000000000000000016179055565b60c0850151600582015560e085015160068201550192613cd3613cc161010083015161ffff1690565b855461ffff191661ffff909116178555565b613d00613ce661012083015161ffff1690565b855463ffff0000191660109190911b63ffff000016178555565b613d31613d1361014083015161ffff1690565b855465ffff00000000191660209190911b65ffff0000000016178555565b613d66613d4461016083015161ffff1690565b855467ffff000000000000191660309190911b67ffff00000000000016178555565b015161ffff1690565b939490969295919561ffff85168015908115614309575b81156142fb575b5080156142ef575b80156142de575b80156142cf575b80156142c3575b80156142b4575b8015614288575b801561426c575b610696576001600160401b03613dd6610144613798565b16151580614256575b8061422b575b80156141cd575b80156141b4575b6141a557604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561047d576001600160401b03915f91614176575b5016801561069657613e6661ffff8a1661ffff8416612ac0565b90612710821061413f5750505f19965b613eac613e92613e8d5f546001600160401b031690565b6137e4565b6001600160401b03166001600160401b03195f5416175f55565b613f0e613ec05f546001600160401b031690565b7f00000000000000000000000000000000000000000000000000000000000000006bffffffff00000000000000006001600160401b036001600160a01b031993169160401b161760a01b1690565b9889613f40613f256021546001600160401b031690565b60406001600160401b03603f83169216101561410157613801565b613f6792906bffffffffffffffffffffffff83549160031b9260a01c831b921b1916179055565b602180546001600160401b038082166001011667ffffffffffffffff19909116179055613f926129f5565b61ffff909716875261ffff16602087015261ffff16604086015261ffff16606085015261ffff851660808501526001600160401b031660a08401526001600160401b031660c0830152151560e08201525f546001600160401b031690436001600160401b03169261ffff16916140088385612efa565b90614011612a05565b3381529260208401526140263661010461381a565b6040840152600160608401526001600160401b031660808301526001600160401b031660a082015260c081015f90528360e082015261010081015f905261012081015f905261014081015f905261016081015f905261018081015f905261409f856001600160a01b0319165f52602260205260405f2090565b906140a991613ba7565b6140b291612efa565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917fcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d79190a390565b61411a61410d82613801565b90549060031b1c60a01b90565b6001600160a01b03198116614130575b50613801565b61413990615545565b5f61412a565b61416b614170927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43612ac0565b6137d5565b96613e76565b614198915060203d60201161419e575b614190818361297f565b8101906137ac565b5f613e4c565b503d614186565b63148b7e9360e31b5f5260045ffd5b5061010061ffff6141c66101246137a2565b1611613df3565b506141dc6105a8610164613798565b151580614215575b8015613dec57506141f66101a4613798565b6001600160401b0361420c6105a8610164613798565b91161115613dec565b506142246105a86101a4613798565b15156141e4565b50614237610184613798565b6001600160401b0361424d6105a8610144613798565b91161115613de5565b506142656105a8610184613798565b1515613ddf565b506001600160401b0382166001600160401b0384161115613dbf565b506142a36105a861ffff88166001600160401b034316612efa565b6001600160401b0383161115613db8565b5061010061ffff871611613db1565b5061ffff861615613daa565b5061271061ffff821610613da3565b5061ffff881661ffff881611613d9c565b5061ffff871615613d95565b905061ffff8916105f613d8d565b61ffff8a16159150613d86565b9060a0828203126102905780601f83011215610290576040519161433b60a08461297f565b829060a0810192831161029057905b8282106143575750505090565b813581526020918201910161434a565b90610140828203126102905780601f83011215610290576101406040519261438f828561297f565b8391810192831161029057905b8282106143a95750505090565b813581526020918201910161439c565b9a99949693959198929790996143e18c6001600160a01b0319165f52602260205260405f2090565b956001600160a01b036143fb88546001600160a01b031690565b16156106b457600487015460ff1692614430610fc560018a01549561442a876001600160401b039060901c1690565b90615947565b6106a55761445a610fc58f610fa8610fbe916001600160a01b0319165f52602360205260405f2090565b6113ec5761ffff8d169586158015614885575b614876576144a78f8f61102e906110286144a16001600160a01b03946001600160a01b0319165f52602460205260405f2090565b9161335a565b1603611314576144d48f611086600491610db133916001600160a01b0319165f52602560205260405f2090565b614867576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102905784925f92859261453060405198899586948594635c73957b60e11b8652600486016129c0565b03915afa92831561047d5761454d93614853575b50810190614367565b928b61455e61113886519260a01c90565b1491821592614837575b821561481c575b50811561480d575b5080156147ff575b80156147f1575b6113145760408051602081018a81529181018990526145b691906145ad8160608101611197565b5190208b614c23565b8060c0840151036113145785610100840151148015906147e2575b611314576101006145e181612a8f565b850361131457614611613686610800966145ff6136868983896133fb565b60208151910120976114009187613409565b602081519101206146348d6001600160a01b0319165f52602b60205260405f2090565b54036113145761464e9261464792614f0d565b9160e00190565b5103611314576146da916146ae60046007936146818c610db133916001600160a01b0319165f52602560205260405f2090565b805461ffff60a01b191660a08d901b61ffff60a01b1617815590600382015501805460ff19166001179055565b0180546146c19060101c61ffff166133da565b63ffff000082549160101b169063ffff00001916179055565b6147156146f9876001600160a01b0319165f52602c60205260405f2090565b928354926001850193845480155f146147dc575060019061596e565b9255556001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561029057604051633c1bcdef60e21b8152336004820152925f908490602490829084905af192831561047d577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb5936147c8575b506040805161ffff9095168552602085019190915283015233926001600160a01b03191691806060810161371a565b806113835f6147d69361297f565b5f614799565b9061596e565b508661012084015114156145d1565b508660a08301511415614586565b50876080830151141561457f565b9050606083015114155f614577565b90915061482e61051760408601612e49565b1415905f61456f565b9150602084015161484b61ffff8416610517565b141591614568565b806113835f6148619361297f565b5f614544565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b50614897601086901c61ffff16610517565b871161446d565b908160a091031261029057604051906148b6826128d6565b80516148c181610a09565b82526020810151602083015260408101516040830152606081015190600382101561029057608091606084015201516148f9816119d2565b608082015290565b6003111561259657565b8054680100000000000000008110156128f15761492d9160018201815561336e565b6001600160a01b0380839493549260031b9316831b921b1916179055565b61ffff60019116019061ffff8211612aa557565b906108a160405161496f8161292d565b60e06149ef82955461ffff8116845261ffff8160101c16602085015261ffff808260201c161660408501526149b261ffff8260301c16606086019061ffff169052565b61ffff604082901c1660808501526001600160401b03605082901c1660a08501526001600160401b03609082901c1660c085015260d01c60ff1690565b1515910152565b906108a1614af66007614a07612a05565b94614a29614a1c82546001600160a01b031690565b6001600160a01b03168752565b614a356001820161495f565b6020870152614a46600282016132c0565b6040870152614a9e614a8e6004830154614a6c614a638260ff1690565b60608b0161389e565b6001600160401b03600882901c1660808a015260481c6001600160401b031690565b6001600160401b031660a0880152565b600581015460c0870152600681015460e0870152015461ffff811661010086015261ffff601082901c1661012086015261ffff602082901c1661014086015261ffff603082901c1661016086015260401c61ffff1690565b61ffff16610180840152565b5f516020615e955f395f51905f5291604051906001600160a01b031960208301931683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c8152614b60606c8261297f565b5190200690565b5f516020615e955f395f51905f5291604051906001600160a01b031960208301931683527fc5cb4182e179e0279f50e2d772929d40dc9d4db3b30ec2ebbefbe6b9bb543075602c830152604c820152604c8152614b60606c8261297f565b5f516020615e955f395f51905f5291604051906001600160a01b031960208301931683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c8152614b60606c8261297f565b5f516020615e955f395f51905f5291604051906001600160a01b031960208301931683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c8152614b60606c8261297f565b9091939294614ca190610400614c9a6201000082612aef565b9186613409565b91614ccd6105176001614cc0610517600789015461ffff9060101c1690565b96015460101c61ffff1690565b92610800915f5b868110614cf7575050505050505090614cf0916108a091614f0d565b0361131457565b61ffff8160051b8901351680158015614deb575b61131457614d6c614d50614d3d614d34866001600160a01b0319165f52602460205260405f2090565b6110288561335a565b90546001600160a01b039160031b1c1690565b610db1856001600160a01b0319165f52602560205260405f2090565b90614d7e610fc5600484015460ff1690565b908115614dce575b50611314576003614db7613686614d9d8886612ac0565b614daf89614daa88612ad3565b612ac0565b90888b613409565b602081519101209101540361131457600101614cd4565b9050614de3610517835461ffff9060a01c1690565b14155f614d86565b50868111614d0b565b919091614e0361040085612aef565b925f5b838110614e4a575050505b60208110614e1e57505050565b8060051b808401351590811591614e3d575b5061131457600101614e11565b905082013515155f614e30565b8060051b61ffff81880135169081158015614f00575b61131457614eae614e92614d3d614e89886001600160a01b0319165f52602460205260405f2090565b6110288661335a565b610db1876001600160a01b0319165f52602960205260405f2090565b91614ec0610fc5600385015460ff1690565b908115614ee3575b50611314576001908701359101540361131457600101614e06565b9050614ef8610517845461ffff9060a01c1690565b14155f614ec8565b5061ffff84168211614e60565b9291905f516020615e955f395f51905f525f940691829060051b8201915b828110614f385750505050565b909192945f516020615e955f395f51905f5283816020938186358b099008970993929101614f2b565b6001600160401b0381116128f15760051b60200190565b90614f8282614f61565b614f8f604051918261297f565b8281528092614fa0601f1991614f61565b0190602036910137565b8051156133835760200190565b80518210156133835760209160051b010190565b91614fd583614f78565b92614fdf81614f78565b91610400810190818111612aa5575f5b8381106152cf57505050801580156152c5575b80156152bb575b6152ac5761501981949294614f78565b9061502381614f78565b945f93845b8381106151f957505061503a82614f78565b61504384614faa565b5161504d82614faa565b5260015b8381106151a5575061507461506e61506885615d0a565b83614fb7565b51615e19565b9061507e84614f78565b9484915b6001831161513d5750505061509684614faa565b525f955f945b8386106150d6575050505050506150c0905f516020615eb55f395f51905f52900690565b036150c757565b6373bdb71560e11b5f5260045ffd5b6150e68683999495969799614fb7565b516150f18988614fb7565b51916137df575f516020615eb55f395f51905f5280918160019461512c6151188e8b614fb7565b515f516020615eb55f395f51905f52900690565b9209095f940897019493929161509c565b9091929394959661514d84615d0a565b61515f61515982615d0a565b84614fb7565b51916137df5761518f5f516020615eb55f395f51905f529182886151999509615188828d614fb7565b5285614fb7565b515f960993615d18565b91909695949396615082565b94856151be6151b8839596979498615d0a565b88614fb7565b516151c98285614fb7565b51926137df576001925f516020615eb55f395f51905f5291096151ec8289614fb7565b5201949093929194615051565b61520b61511882849895969798614fb7565b600180915f905b88821061524157505090829161522a6001948c614fb7565b526152358289614fb7565b52019493929194615028565b909295918385146152a1576152596151188588614fb7565b92838314615292576137df575f516020615eb55f395f51905f52808085600194099461528485615ce4565b90085f9809935b0190615212565b63027639eb60e31b5f5260045ffd5b91959260019061528b565b630a4960f960e31b5f5260045ffd5b5081518111615009565b5083518111615002565b8060019160051b808401356152e4838b614fb7565b528401356152f28288614fb7565b5201614fef565b5f516020615e955f395f51905f528110801590615385575b61537f575f516020615e955f395f51905f528082819309928009818080808487097f1aee90f15f2189693df072d799fd11fc039b2959ebb7c867d075ca8cf4d7eb8e09600108937f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000009081490565b50505f90565b505f516020615e955f395f51905f52821015615311565b90939291936153cc61051760016153bf6153b588612ae1565b9761048090612aef565b97015460101c61ffff1690565b905f5b84811061542257505050505b602081106153e857505050565b8060061b83018160051b8301356113145780351590811591615412575b50611314576001016153db565b600191506020013514155f615405565b8060061b870161ffff8260051b88013516908115801561550c575b61131457615486615466614d3d614e89896001600160a01b0319165f52602460205260405f2090565b610db1866110758a6001600160a01b0319165f52602660205260405f2090565b91615498610fc5600485015460ff1690565b9081156154ef575b5061131457815460b01c61ffff1661ffff808616911603611314576002820154813514918215926154da575b5050611314576001016153cf565b60209192506003015491013514155f806154cc565b9050615504610517845461ffff9060a01c1690565b14155f6154a0565b5084821161543d565b8054905f815581615524575050565b5f5260205f20908101905b81811061553a575050565b5f815560010161552f565b6001600160a01b0361557961556c836001600160a01b0319165f52602260205260405f2090565b546001600160a01b031690565b16156159445761559b816001600160a01b0319165f52602460205260405f2090565b8054905f5b8281106157b95750505060015b61010061ffff821611156156a75750806155ea6155e56001600160a01b0319936001600160a01b0319165f52602460205260405f2090565b615515565b5f615607826001600160a01b0319165f52602b60205260405f2090565b55615632615627826001600160a01b0319165f52602c60205260405f2090565b60015f918281550155565b615680615651826001600160a01b0319165f52602260205260405f2090565b60075f918281558260018201558260028201558260038201558260048201558260058201558260068201550155565b167f98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef265f80a2565b806156d161051761352f61573294611075876001600160a01b0319165f52602760205260405f2090565b615786575b6157036156f982611075866001600160a01b0319165f52602860205260405f2090565b5460101c60ff1690565b61575e575b61572881611075856001600160a01b0319165f52602d60205260405f2090565b54615737576133da565b6155ad565b5f61575882611075866001600160a01b0319165f52602d60205260405f2090565b556133da565b61578161562782611075866001600160a01b0319165f52602860205260405f2090565b615708565b6157b46157a982611075866001600160a01b0319165f52602760205260405f2090565b805461ffff19169055565b6156d6565b6157c6614d3d828461336e565b6157f36157e982610db1886001600160a01b0319165f52602360205260405f2090565b805460ff19169055565b5f61582f615813876001600160a01b0319165f52602a60205260405f2090565b61581f61051786612ad3565b61ffff165f5260205260405f2090565b5561587061585382610db1886001600160a01b0319165f52602560205260405f2090565b60045f918281558260018201558260028201558260038201550155565b6158aa61589382610db1886001600160a01b0319165f52602960205260405f2090565b60035f918281558260018201558260028201550155565b60015b61010061ffff821611156158c55750506001016155a0565b806158f2600461108685610db16158fb966110758d6001600160a01b0319165f52602660205260405f2090565b615900576133da565b6158ad565b61128061592784610db1846110758c6001600160a01b0319165f52602660205260405f2090565b60045f918281558260018201558260016002830182815501550155565b50565b600681101561259657600214908161595d575090565b6001600160401b0391501643111590565b9392919091841580615a8e575b615a8657811580615a7c575b615a77575f516020615e955f395f51905f52828609945f516020615e955f395f51905f528285095f516020615e955f395f51905f528188095f516020615e955f395f51905f5290620292f809965f516020615e955f395f51905f5290620292fc096159f191615d6c565b935f516020615e955f395f51905f5287600108615a0d90615db3565b935f516020615e955f395f51905f529109915f516020615e955f395f51905f529109905f516020615e955f395f51905f529108905f516020615e955f395f51905f52910992615a5b90615d24565b615a6490615db3565b5f516020615e955f395f51905f52910990565b505090565b5060018114615987565b935090509190565b506001831461597b565b600681101561259657600114908161595d575090565b60405190610400615abf818461297f565b368337565b60405190610800615abf818461297f565b9060208110156133835760051b0190565b9060408110156133835760051b0190565b91905f835b60208210615b375750505061040082015f905b60408210615b2157505050610c000190565b6020806001928551815201930191019091615b0f565b6020806001928551815201930191019091615afc565b919091615b58615aae565b615b60615ac4565b93615b7d836001600160a01b0319165f52602460205260405f2090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110615c265750505061ffff165b60208110615c015750611197615be2615bfe939495604051928391602083019586615af7565b519020916001600160a01b0319165f52602b60205260405f2090565b55565b806001615c1f615c19615c148395612aaa565b612ad3565b88615ae6565b5201615bbc565b80615c33615c7692612ad3565b615c3d8288615ad5565b5260a0615c4d614d3d838761336e565b6040516313a4120960e31b81526001600160a01b03909116600482015292839081906024820190565b0381865afa91821561047d576001926040915f91615cc6575b506020810151615ca7615ca185612aaa565b8d615ae6565b520151615cbf615cb9615c1484612aaa565b8b615ae6565b5201615ba9565b615cde915060a03d811161244d5761243f818361297f565b5f615c8f565b5f516020615eb55f395f51905f5203905f516020615eb55f395f51905f528211612aa557565b5f19810191908211612aa557565b8015612aa5575f190190565b5f516020615e955f395f51905f5290065f516020615e955f395f51905f52035f516020615e955f395f51905f528111612aa5575f516020615e955f395f51905f529060010890565b905f516020615e955f395f51905f5290065f516020615e955f395f51905f52035f516020615e955f395f51905f528111612aa5575f516020615e955f395f51905f52910890565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f516020615e955f395f51905f5260a082015260208160c08160055afa15610290575190565b9060405191602083526020808401526020604084015260608301527f060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126ef60808301525f516020615eb55f395f51905f5260a083015260208260c08160055afa15615e875760c082519201604052565b639e44e6e05f526004601cfdfe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001060c89ce5c263405370a08b6d0302b0bab3eedb83920ee0a677297dc392126f1a26469706673582212205eaf9509ad751bbb5fd68ca073b8ad1bc7fc02083c97bada78fbc2753c73192764736f6c634300081c0033",
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
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
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
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetRound(roundId [12]byte) (IDKGManagerRound, error) {
	return _DKGManager.Contract.GetRound(&_DKGManager.CallOpts, roundId)
}

// GetRound is a free data retrieval call binding the contract method 0xf4e34945.
//
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),(bool,uint16,uint64,uint64,uint64,uint64),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16,uint16))
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

// CreateRound is a paid mutator transaction binding the contract method 0xbfa78991.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateRound(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createRound", threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed, decryptionPolicy)
}

// CreateRound is a paid mutator transaction binding the contract method 0xbfa78991.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed, decryptionPolicy)
}

// CreateRound is a paid mutator transaction binding the contract method 0xbfa78991.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed, (bool,uint16,uint64,uint64,uint64,uint64) decryptionPolicy) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool, decryptionPolicy DKGTypesDecryptionPolicy) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed, decryptionPolicy)
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
