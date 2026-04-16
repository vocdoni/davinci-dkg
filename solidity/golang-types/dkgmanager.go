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
	CombineHash     [32]byte
	PlaintextHash   [32]byte
	Completed       bool
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
	Status                 uint8
	Nonce                  uint64
	SeedBlock              uint64
	Seed                   [32]byte
	LotteryThreshold       *big.Int
	ClaimedCount           uint16
	ContributionCount      uint16
	PartialDecryptionCount uint16
	RevealedShareCount     uint16
}

// DKGManagerMetaData contains all meta data concerning the DKGManager contract.
var DKGManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealSubmitVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealShareVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SHARE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SUBMIT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROUND_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRound\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCollectivePublicKey\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delta\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealShareVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealSubmitVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RevealedShareRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Round\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RoundPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.RoundStatus\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"revealedShareCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reconstructSecret\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roundNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitment0X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment0Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationExtended\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"newSeedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newRegistrationDeadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevealedShareSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundAborted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundCreated\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundEvicted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundFinalized\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecretReconstructed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisclosureDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRevealedShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReconstruction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRevealedShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6101a0806040523461029657610100816151cf8038038091610021828561029a565b8339810103126102965780519063ffffffff82169182810361029657610049602083016102d1565b610055604084016102d1565b610061606085016102d1565b9061006e608086016102d1565b9261007b60a087016102d1565b9461009460e061008d60c08a016102d1565b98016102d1565b9763ffffffff461603610287576001600160a01b03821615610278576001600160a01b038316158015610267575b8015610256575b8015610245575b8015610234575b8015610223575b6102145763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b60248201526018815261011a60388261029a565b5190201660c05260e0526101005261012052610140526101605261018052604051614ee990816102e682396080518161134d015260a05181818161033e01528181611d23015281816130c901528181613ce50152614c39015260c051818181610a47015261317f015260e0518181816103ca01528181610b6e0152613a9e015261010051818181610fd201528181611396015261153a015261012051818181610bb7015281816113f80152612652015261014051818181610dae015281816122310152613682015261016051818181610381015281816114d80152611747015261018051818181610a8a015281816115830152612a930152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038816156100de565b506001600160a01b038716156100d7565b506001600160a01b038616156100d0565b506001600160a01b038516156100c9565b506001600160a01b038416156100c2565b63e6c4247b60e01b5f5260045ffd5b633d23e4d160e11b5f5260045ffd5b5f80fd5b601f909101601f19168101906001600160401b038211908210176102bd57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102965756fe60806040526004361015610011575f80fd5b5f3560e01c8063058994a11461023457806306433b1b1461022f578063070c74921461022a578063074a75e1146102255780630b1451f0146102205780630e2c53f71461021b5780633353ec6e14610216578063349181a214610211578063415a1b861461020c578063510ba2df1461020757806353d721841461020257806356664d15146101fd5780635ddd0626146101f857806362c41927146101f357806363f314cd146101ee578063669a76a9146101e957806370f2469b146101e457806372517b4b146101df578063802ae231146101da57806385e1f4d0146101d55780638dc1f53a146101d057806393c3d3a8146101cb5780639f431549146101c6578063b18730c2146101c1578063bf192209146101bc578063c2440e16146101b7578063c9396bf0146101b2578063ca3c0458146101ad578063d3720aac146101a8578063d6c29c9e146101a3578063d99337671461019e578063f4e3494514610199578063fe1604b5146101945763fe2348971461018f575f80fd5b612255565b612212565b612142565b611c1f565b611b82565b611a56565b6119ce565b6115c6565b61155e565b61151b565b6114b3565b61141c565b6113d9565b611371565b611331565b610df1565b610d89565b610c5e565b610b92565b610b4f565b610abf565b610a6b565b610a2b565b610917565b6108b1565b610848565b61074b565b610710565b610683565b610442565b6103a5565b610362565b61031f565b610281565b600435906001600160a01b03198216820361025057565b5f80fd5b9181601f84011215610250578235916001600160401b038311610250576020838186019501011161025057565b346102505760e03660031901126102505761029a610239565b602435604435916064356084356001600160401b038111610250576102c3903690600401610254565b60a4929192356001600160401b038111610250576102e5903690600401610254565b93909260c435976001600160401b0389116102505761030b610313993690600401610254565b989097612583565b005b5f91031261025057565b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e575b50604051908152602090f35b610430915060203d602011610436575b61042881836123ee565b81019061293a565b5f610402565b503d61041e565b612459565b346102505760203660031901126102505761045b610239565b610477816001600160a01b0319165f52602260205260405f2090565b6001600160a01b0361049082546001600160a01b031690565b161561067457600281019081549260016104aa8560ff1690565b6104b381612065565b0361066557600582015461ffff16600183018054909161ffff6104de601084901c82165b61ffff1690565b911614610665576001600160401b03605082901c1695438710156106655761056861055861053e61051c610574946001600160401b039060481c1690565b99610538610531604088901c61ffff166104d7565b809c612949565b90612949565b6105526001600160401b0343169a8b612969565b99612969565b9260901c6001600160401b031690565b6001600160401b031690565b6001600160401b03821610156106565761062b81610651936105f5897f9f2b9abf7edf3bc2ca127de52d5e6f818ee43f02fa41ffd5ef9d24e45130cb9c995f60036001600160a01b03199b01559070ffffffffffffffff00000000000000000082549160481b169070ffffffffffffffff0000000000000000001916179055565b9071ffffffffffffffff0000000000000000000082549160501b169071ffffffffffffffff000000000000000000001916179055565b6040519384931695839092916001600160401b0360209181604085019616845216910152565b0390a2005b63d06b96b160e01b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b6328ad4a9560e21b5f5260045ffd5b346102505760c03660031901126102505761069c610239565b602435604435916064356001600160401b038111610250576106c2903690600401610254565b906084356001600160401b038111610250576106e2903690600401610254565b92909160a435966001600160401b03881161025057610708610313983690600401610254565b9790966129da565b3461025057602036600319011261025057604061073361072e610239565b612cac565b6107498251809260208091805184520151910152565bf35b3461025057602036600319011261025057610764610239565b610780816001600160a01b0319165f52602260205260405f2090565b80546001600160a01b03168015610674576001600160a01b0316330361083a57600201906107af825460ff1690565b6107b881612065565b60038114908115610825575b8115610811575b50610665576107ea6001600160a01b031992600460ff19825416179055565b167f97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e4105f80a2005b6004915061081e81612065565b145f6107cb565b905061083081612065565b60058114906107c4565b6282b42960e81b5f5260045ffd5b34610250575f3660031901126102505760206001600160401b035f5416604051908152f35b6024359061ffff8216820361025057565b6044359061ffff8216820361025057565b6064359061ffff8216820361025057565b6084359061ffff8216820361025057565b346102505760403660031901126102505760206108fd6108cf610239565b6001600160a01b03196108e061086d565b91165f52602a835260405f209061ffff165f5260205260405f2090565b54604051908152f35b6001600160a01b0381160361025057565b3461025057604036600319011261025057610a27610994610936610239565b6001600160a01b03196024359161094c83610906565b5f608060405161095b81612345565b8281528260208201528260408201528260608201520152165f52602960205260405f20906001600160a01b03165f5260205260405f2090565b60ff6003604051926109a584612345565b61ffff81546001600160a01b038116865260a01c1660208501526001810154604085015260028101546060850152015416151560808201526040519182918291909160808060a08301946001600160a01b03815116845261ffff6020820151166020850152604081015160408501526060810151606085015201511515910152565b0390f35b34610250575f36600319011261025057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b6001600160401b0381160361025057565b34610250576101003660031901126102505760043561ffff8116810361025057610ae761086d565b90610af061087e565b610af861088f565b610b006108a0565b60a43590610b0d82610aae565b60c43592610b1a84610aae565b60e43594851515860361025057610a2797610b3497613060565b6040516001600160a01b031990911681529081906020820190565b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e5750604051908152602090f35b91909160c060a060e08301946001600160a01b03815116845261ffff602082015116602085015261ffff604082015116604085015260608101516060850152610c556080820151608086019060208091805184520151910152565b01511515910152565b3461025057606036600319011261025057610a27610d05610c7d610239565b610cef60243591610c8d83610906565b6001600160a01b0319610c9e61087e565b915f60a0604051610cae81612365565b828152826020820152826040820152826060820152610ccb612c76565b60808201520152165f52602660205260405f209061ffff165f5260205260405f2090565b906001600160a01b03165f5260205260405f2090565b610d7d610d74600460405193610d1a85612365565b610d51610d4682546001600160a01b038116885261ffff8160a01c16602089015261ffff9060b01c1690565b61ffff166040870152565b60018101546060860152610d6760028201612c8e565b6080860152015460ff1690565b151560a0830152565b60405191829182610bfa565b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e5750604051908152602090f35b346102505760c036600319011261025057610e0a610239565b610e1261086d565b610e1a61087e565b6064356084356001600160401b03811161025057610e3c903690600401610254565b60a4939193356001600160401b03811161025057610e5e903690600401610254565b90610e7b886001600160a01b0319165f52602260205260405f2090565b956001600160a01b03610e9588546001600160a01b031690565b1615610674576003610eab600289015460ff1690565b610eb481612065565b0361066557610eff610efb610ef4610ede8c6001600160a01b0319165f52602360205260405f2090565b336001600160a01b03165f5260205260405f2090565b5460ff1690565b1590565b6113225761ffff88169384158015611305575b80156112f9575b80156112ea575b80156112e2575b6112d457610f7b6001600160a01b03610f64610f558d6001600160a01b0319165f52602460205260405f2090565b610f5e8d6134f2565b90613506565b9290923393546001600160a01b039160031b1c1690565b160361124a57610fc46004610fbc8c610cef8a610fab33936001600160a01b0319165f52602660205260405f2090565b9061ffff165f5260205260405f2090565b015460ff1690565b6112c5576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f92859261102060405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d5761103d936112ab575b5081019061351b565b9061105e86610fab896001600160a01b0319165f52602a60205260405f2090565b54825161107f61106e8a60a01c90565b6bffffffffffffffffffffffff1690565b149182159261129c575b508115611293575b8115611259575b5061124a5760c081018051909160e001906110cd6110db83516040519283916020830195869091604092825260208201520190565b03601f1981018352826123ee565b519020840361124a577f39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f22946111d89260036005936111348b610cef89610fab33936001600160a01b0319165f52602660205260405f2090565b805460a08c901b61ffff60a01b167fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff9091161760b089901b77ffff000000000000000000000000000000000000000000001617815560048101805460ff1916600117905592516002840155519101550180546111bb9060201c61ffff1661356d565b61356d565b65ffff0000000082549160201b169065ffff000000001916179055565b61121a6111fb82610fab886001600160a01b0319165f52602760205260405f2090565b61120a6111b6825461ffff1690565b61ffff1661ffff19825416179055565b6040805161ffff958616815294909116602085015283015233926001600160a01b0319169180606081015b0390a3005b63d1fed5fd60e01b5f5260045ffd5b905060808201516110cd61128860a085015b516040805160208101958652908101919091529182906060820190565b51902014155f611098565b80159150611091565b6020840151141591505f611089565b806112b95f6112bf936123ee565b80610315565b5f611034565b633466526160e01b5f5260045ffd5b62d949df60e51b5f5260045ffd5b508615610f27565b5061010061ffff871611610f20565b5061ffff861615610f19565b50600188015461131b9060101c61ffff166104d7565b8511610f12565b63965c290d60e01b5f5260045ffd5b34610250575f36600319011261025057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e5750604051908152602090f35b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102505760e036600319011261025057611435610239565b61143d61086d565b604435916064356084356001600160401b03811161025057611463903690600401610254565b60a4929192356001600160401b03811161025057611485903690600401610254565b93909260c435976001600160401b038911610250576114ab610313993690600401610254565b989097613580565b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e5750604051908152602090f35b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610250575f3660031901126102505760405163233ace1160e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d575f9161040e5750604051908152602090f35b346102505760a0366003190112610250576115df610239565b6115e761086d565b6044356064356001600160401b03811161025057611609903690600401610254565b6084356001600160401b03811161025057611628903690600401610254565b90611645876001600160a01b0319165f52602260205260405f2090565b936001600160a01b0361165f86546001600160a01b031690565b16156106745760018501549360ff60d086901c161561197d576003611688600288015460ff1690565b61169181612065565b03610665576116bb610efb610ef4610ede8c6001600160a01b0319165f52602360205260405f2090565b6113225761ffff8816948515908115611964575b50801561195c575b61194d5761170e6001600160a01b03610f646117058c6001600160a01b0319165f52602460205260405f2090565b610f5e8c6134f2565b160361124a576117396003610fbc33610cef8d6001600160a01b0319165f52602960205260405f2090565b61193e576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f92859261179560405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d576117b29361192a575b50810190613867565b6117d285610fab886001600160a01b0319165f52602a60205260405f2090565b549080516117e361106e8960a01c90565b149283159361191b575b50821561190c575b8215611903575b82156118df575b505061124a57816118b56005611245937f5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac95600161185733610cef8c6001600160a01b0319165f52602960205260405f2090565b805461ffff60a01b191660a08b901b61ffff60a01b1617815560038101805460ff1916600117905501550180546118949060301c61ffff1661356d565b67ffff00000000000082549160301b169067ffff0000000000001916179055565b6040519182916001600160a01b031933971695836020909392919361ffff60408201951681520152565b9091506110cd6118f761126b60608401519360800190565b51902014155f80611803565b811592506117fc565b604081015185141592506117f5565b6020820151141592505f6117ed565b806112b95f611938936123ee565b5f6117a9565b63a89ac15160e01b5f5260045ffd5b639eae062d60e01b5f5260045ffd5b5086156116d7565b611975915060101c61ffff166104d7565b85115f6116cf565b630ba0cb2f60e21b5f5260045ffd5b60206040818301928281528451809452019201905f5b8181106119af5750505090565b82516001600160a01b03168452602093840193909201916001016119a2565b34610250576020366003190112610250576001600160a01b03196119f0610239565b165f52602460205260405f206040519081602082549182815201915f5260205f20905f5b818110611a3757610a2785611a2b818703826123ee565b6040519182918261198c565b82546001600160a01b0316845260209093019260019283019201611a14565b3461025057604036600319011261025057610a27611ad9611a75610239565b6001600160a01b031960243591611a8b83610906565b5f60a0604051611a9a81612365565b8281528260208201528260408201528260608201528260808201520152165f52602560205260405f20906001600160a01b03165f5260205260405f2090565b611b2e610d74600460405193611aee85612365565b80546001600160a01b038116865260a01c61ffff166020860152600181015460408601526002810154606086015260038101546080860152015460ff1690565b6040519182918291909160a08060c08301946001600160a01b03815116845261ffff602082015116602085015260408101516040850152606081015160608501526080810151608085015201511515910152565b346102505761012036600319011261025057611b9c610239565b611ba461086d565b6044359160843560643560a43560c4356001600160401b03811161025057611bd0903690600401610254565b9160e4356001600160401b03811161025057611bf0903690600401610254565b95909461010435996001600160401b038b1161025057611c176103139b3690600401610254565b9a9099613975565b3461025057602036600319011261025057611c38610239565b611c54816001600160a01b0319165f52602260205260405f2090565b906001600160a01b03611c6e83546001600160a01b031690565b16156106745760028201805460018401805494909391611ca1610efb6001600160401b03605089901c1660ff8416614b52565b61066557600582019561ffff611cca6104d7611cbf8a5461ffff1690565b9360101c61ffff1690565b91161015611fc257611cf4610ef4610ede856001600160a01b0319165f52602360205260405f2090565b611fb357600382018054918215611f2b575b50506040516313a4120960e31b815233600482015260a0816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561043d576001916060915f91611efc575b500151611d6d81613ec4565b611d7681613ec4565b03611eed5760408051602081019283523360601b6bffffffffffffffffffffffff19169181019190915260049190611db181605481016110cd565b5190209101541115611ede57611e2b611dcc855461ffff1690565b94611df233611ded856001600160a01b0319165f52602460205260405f2090565b613ece565b611e22611e1533610cef866001600160a01b0319165f52602360205260405f2090565b805460ff19166001179055565b61120a86613f0e565b611e8e611e826001600160a01b0319831695604051877f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b339280611e7a868291909161ffff6020820193169052565b0390a3613f0e565b935460101c61ffff1690565b9261ffff808516911614611e9e57005b611eb892611eab91614c07565b805460ff19166002179055565b7fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13585f80a2005b637c75aa6f60e11b5f5260045ffd5b63aba4733960e01b5f5260045ffd5b611f1e915060a03d60a011611f24575b611f1681836123ee565b810190613e61565b5f611d61565b503d611f0c565b909150611f439060481c6001600160401b0316610568565b80431115611fa45740908115611f95578190556040518181526001600160a01b03198416907fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b5013965290602090a25f80611d06565b6302504bb360e61b5f5260045ffd5b63172181cb60e21b5f5260045ffd5b630c8d9eab60e31b5f5260045ffd5b63848084dd60e01b5f5260045ffd5b9060e08061204f9361ffff815116845261ffff602082015116602085015261ffff60408201511660408501526120126060820151606086019061ffff169052565b60808181015161ffff169085015260a0818101516001600160401b03169085015260c0818101516001600160401b03169085015201511515910152565b565b634e487b7160e01b5f52602160045260245ffd5b6006111561206f57565b612051565b90600682101561206f5752565b61204f909291926102206101406102408301956120a78482516001600160a01b03169052565b6120b960208201516020860190611fd1565b6120cc6040820151610120860190612074565b60608101516001600160401b03168483015260808101516001600160401b031661016085015260a081015161018085015260c08101516101a085015260e081015161ffff166101c085015261010081015161ffff166101e085015261012081015161ffff16610200850152015161ffff16910152565b3461025057602036600319011261025057610a27612206612201612164610239565b5f61014060405161217481612380565b8281526040516121838161239c565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e082015260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015201526001600160a01b0319165f52602260205260405f2090565b613fb9565b60405191829182612081565b34610250575f3660031901126102505760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461025057604036600319011261025057610a276122c3612274610239565b6001600160a01b031961228561086d565b915f6060604051612295816123b8565b8281528260208201528260408201520152165f52602860205260405f209061ffff165f5260205260405f2090565b60ff6003604051926122d4846123b8565b61ffff815416845260018101546020850152600281015460408501520154161515606082015260405191829182919091606080608083019461ffff8151168452602081015160208501526040810151604085015201511515910152565b634e487b7160e01b5f52604160045260245ffd5b60a081019081106001600160401b0382111761236057604052565b612331565b60c081019081106001600160401b0382111761236057604052565b61016081019081106001600160401b0382111761236057604052565b61010081019081106001600160401b0382111761236057604052565b608081019081106001600160401b0382111761236057604052565b604081019081106001600160401b0382111761236057604052565b90601f801991011681019081106001600160401b0382111761236057604052565b908060209392818452848401375f828201840152601f01601f1916010190565b929061244890612456959360408652604086019161240f565b92602081850391015261240f565b90565b6040513d5f823e3d90fd5b6040519061204f610100836123ee565b6040519061204f610160836123ee565b90610120828203126102505780601f8301121561025057610120604051926124ac82856123ee565b8391810192831161025057905b8282106124c65750505090565b81358152602091820191016124b9565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b908160051b918083046020149015171561251457565b6124ea565b908160011b918083046002149015171561251457565b8181029291811591840414171561251457565b906001820180921161251457565b906080820180921161251457565b9190820180921161251457565b90600681101561206f5760ff80198354169116179055565b97919590989692969493946125aa896001600160a01b0319165f52602260205260405f2090565b936001600160a01b036125c486546001600160a01b031690565b16156106745760028501916125da835460ff1690565b6125e381612065565b6003811461292b57806125f7600292612065565b03610665576005860194612611865461ffff9060101c1690565b9460018801549361262a6104d78661ffff9060201c1690565b61ffff88161061291c578e158015612914575b801561290c575b6128fd576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f9285926126a060405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d576126bd936128e9575b50810190612484565b9283516126cd61106e8d60a01c90565b14918215926128cd575b82156128aa575b508115612897575b508015612889575b801561287b575b801561286d575b61124a576127348a60405161272b816110cd8d8d60208401968791606093918352602083015260408201520190565b5190208a6140a4565b8060e08401510361124a5761274a6108a06124fe565b860361124a5761278b95878b612773936104d79861276c612780986101000190565b5192614223565b805460ff19166003179055565b5460101c61ffff1690565b906127a061279a6108606124fe565b8261255e565b5f5b8381106127f757505060408051968752602087019390935250508301526001600160a01b031916907f5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a724559080606081015b0390a2565b8060019160061b8301604051612829816110cd6020820194602081013590358660209093929193604081019481520152565b51902061286661284b8a6001600160a01b0319165f52602a60205260405f2090565b61ffff8460051b8801351661ffff165f5260205260405f2090565b55016127a2565b508760c083015114156126fc565b508660a083015114156126f5565b5089608083015114156126ee565b6060840151915061ffff1614155f6126e6565b9091506128c46104d7604086015b519260101c61ffff1690565b1415905f6126de565b915060208401516128e161ffff84166104d7565b1415916126d7565b806112b95f6128f7936123ee565b5f6126b4565b63c5f680ed60e01b5f5260045ffd5b508c15612644565b508b1561263d565b63368f2d7d60e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b90816020910312610250575190565b906001600160401b03809116911603906001600160401b03821161251457565b906001600160401b03809116911601906001600160401b03821161251457565b9060e0828203126102505780601f8301121561025057604051916129ae60e0846123ee565b829060e0810192831161025057905b8282106129ca5750505090565b81358152602091820191016129bd565b949195969290976129fd866001600160a01b0319165f52602260205260405f2090565b936001600160a01b03612a1786546001600160a01b031690565b16156106745760018501549360ff60d086901c161561197d5760028601996003612a428c5460ff1690565b612a4b81612065565b03610665578b158015612c6e575b612c5f57612a73600561ffff98015461ffff9060301c1690565b612a8061ffff88166104d7565b9788911610612c50576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f928592612ae160405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d57612afe93612c3c575b50810190612989565b938451612b0e61106e8860a01c90565b14801590612c2e575b8015612c20575b8015612c12575b61124a57604085019384511061124a5760408051602081018b8152918101899052612b619190612b5881606081016110cd565b51902087614109565b908160a08701510361124a57612b7760406124fe565b0361124a57612b9e82612ba395612b9560409661ffff9060101c1690565b90519089614396565b6144af565b60c08201510361124a578290608001510361124a576127f26001600160a01b031992612bf77fbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b48295600560ff19825416179055565b60405193849316958360209093929193604081019481520152565b508660808601511415612b25565b508860608601511415612b1e565b508360208601511415612b17565b806112b95f612c4a936123ee565b5f612af5565b63957674fd60e01b5f5260045ffd5b6314141ce560e21b5f5260045ffd5b508915612a59565b60405190612c83826123d3565b5f6020838281520152565b90604051612c9b816123d3565b602060018294805484520154910152565b6001600160a01b031990612cbe612c76565b50165f52602c60205260405f20600181015415612cde5761245690612c8e565b50604051612ceb816123d3565b5f81526001602082015290565b90816020910312610250575161245681610aae565b8115612d17570490565b634e487b7160e01b5f52601260045260245ffd5b6001600160401b03166001600160401b0381146125145760010190565b906040821015612d6157600c600183811c810193160290565b6124d6565b600682101561206f5752565b612e9360e061204f93612d9661ffff825116859061ffff1661ffff19825416179055565b60208181015185546040808501516060860151608087015160a0880151931b69ffff00000000000000001660309190911b67ffff000000000000169190951b65ffff000000001660109490941b63ffff00001671ffffffffffffffffffffffffffffffff0000199093169290921792909217179190911760509190911b71ffffffffffffffff0000000000000000000016178455612e8c612e4160c08301516001600160401b031690565b85547fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff1660909190911b79ffffffffffffffff00000000000000000000000000000000000016178555565b0151151590565b81547fffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560d01b7aff000000000000000000000000000000000000000000000000000016179055565b90611894610140600561204f94612f1c612f0086516001600160a01b031690565b82906001600160a01b03166001600160a01b0319825416179055565b612f2d602086015160018301612d72565b612fbf60028201612f4b6040880151612f4581612065565b8261256b565b612f86612f6260608901516001600160401b031690565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b6080870151815470ffffffffffffffff000000000000000000191660489190911b70ffffffffffffffff00000000000000000016179055565b60a0850151600382015560c085015160048201550192612ff9612fe760e083015161ffff1690565b855461ffff191661ffff909116178555565b61302661300c61010083015161ffff1690565b855463ffff0000191660109190911b63ffff000016178555565b61305761303961012083015161ffff1690565b855465ffff00000000191660209190911b65ffff0000000016178555565b015161ffff1690565b939490969295919561ffff851680159081156134e5575b81156134d7575b5080156134cb575b80156134ba575b80156134ab575b801561349f575b8015613490575b8015613464575b8015613448575b61065657604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561043d576001600160401b03915f91613419575b501680156106565761312361ffff8a1661ffff841661252f565b9061271082106133e25750505f19965b61316961314f61314a5f546001600160401b031690565b612d2b565b6001600160401b03166001600160401b03195f5416175f55565b6131cb61317d5f546001600160401b031690565b7f00000000000000000000000000000000000000000000000000000000000000006bffffffff00000000000000006001600160401b036001600160a01b031993169160401b161760a01b1690565b98896131fd6131e26021546001600160401b031690565b60406001600160401b03603f8316921610156133a457612d48565b61322492906bffffffffffffffffffffffff83549160031b9260a01c831b921b1916179055565b602180546001600160401b038082166001011667ffffffffffffffff1990911617905561324f612464565b61ffff909716875261ffff16602087015261ffff16604086015261ffff16606085015261ffff851660808501526001600160401b031660a08401526001600160401b031660c0830152151560e08201525f546001600160401b031690436001600160401b03169261ffff16916132c58385612969565b906132ce612474565b338152926020840152600160408401526001600160401b031660608301526001600160401b0316608082015260a081015f90528360c082015260e081015f905261010081015f905261012081015f905261014081015f9052613342856001600160a01b0319165f52602260205260405f2090565b9061334c91612edf565b61335591612969565b604080516001600160401b03929092168252602082019290925233916001600160a01b03198416917fcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d79190a390565b6133bd6133b082612d48565b90549060031b1c60a01b90565b6001600160a01b031981166133d3575b50612d48565b6133dc90614533565b5f6133cd565b61340e613413927e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa4361252f565b612d0d565b96613133565b61343b915060203d602011613441575b61343381836123ee565b810190612cf8565b5f613109565b503d613429565b506001600160401b0382166001600160401b03841611156130b0565b5061347f61056861ffff88166001600160401b034316612969565b6001600160401b03831611156130a9565b5061010061ffff8716116130a2565b5061ffff86161561309b565b5061271061ffff821610613094565b5061ffff881661ffff88161161308d565b5061ffff871615613086565b905061ffff8916105f61307e565b61ffff8a16159150613077565b61ffff5f199116019061ffff821161251457565b8054821015612d61575f5260205f2001905f90565b906101a0828203126102505780601f83011215610250576101a06040519261354382856123ee565b8391810192831161025057905b82821061355d5750505090565b8135815260209182019101613550565b61ffff1661ffff81146125145760010190565b949295989791979690966135a6866001600160a01b0319165f52602260205260405f2090565b926001600160a01b036135c085546001600160a01b031690565b16156106745760036135d6600286015460ff1690565b6135df81612065565b036106655761ffff89169a8b15801561385c575b8015613854575b801561384c575b61383d5761ffff9361363461362c8c610fab8c6001600160a01b0319165f52602760205260405f2090565b5461ffff1690565b6136466104d7600189015461ffff1690565b958691161061382e576136746003610fbc8d610fab8d6001600160a01b0319165f52602860205260405f2090565b61381f576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f9285926136d060405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d576136ec93612c3c5750810190612989565b9384516136fc61106e8860a01c90565b14801590613811575b8015613803575b80156137f5575b61124a57604085019182511061124a5760408051602081018b815291810189905261374f919061374681606081016110cd565b51902087614167565b938460a08701510361124a5761376560646124fe565b0361124a5761378993612b9e826137829560649551908c8b614888565b9160c00190565b510361124a576137f06001600160a01b031992612bf760036137e27f451276810ef520579055672046d83aad5adae5e72513ec6b904ac15cd449611597610fab876001600160a01b0319165f52602860205260405f2090565b01805460ff19166001179055565b0390a3565b508660808601511415613713565b50886060860151141561370c565b508160208601511415613705565b63955c0c4960e01b5f5260045ffd5b63032cddf960e11b5f5260045ffd5b636d28699160e01b5f5260045ffd5b508815613601565b508a156135fa565b506101008c116135f3565b9060a0828203126102505780601f83011215610250576040519161388c60a0846123ee565b829060a0810192831161025057905b8282106138a85750505090565b813581526020918201910161389b565b90610140828203126102505780601f8301121561025057610140604051926138e082856123ee565b8391810192831161025057905b8282106138fa5750505090565b81358152602091820191016138ed565b909291928311610250579190565b90939293848311610250578411610250578101920390565b9291926001600160401b0382116123605760405191613959601f8201601f1916602001846123ee565b829481845281830111610250578281602093845f960137010152565b9a999496939591989297909961399d8c6001600160a01b0319165f52602260205260405f2090565b956001600160a01b036139b788546001600160a01b031690565b161561067457600287015460ff16926139ec610efb60018a0154956139e6876001600160401b039060901c1690565b90614a01565b61066557613a16610efb8f610ede610ef4916001600160a01b0319165f52602360205260405f2090565b6113225761ffff8d169586158015613e48575b613e3957613a638f8f610f6490610f5e613a5d6001600160a01b03946001600160a01b0319165f52602460205260405f2090565b916134f2565b160361124a57613a908f610fbc600491610cef33916001600160a01b0319165f52602560205260405f2090565b613e2a576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102505784925f928592613aec60405198899586948594635c73957b60e11b86526004860161242f565b03915afa92831561043d57613b0993613e16575b508101906138b8565b928b613b1a61106e86519260a01c90565b1491821592613dfa575b8215613ddf575b508115613dd0575b508015613dc2575b8015613db4575b61124a5760408051602081018a8152918101899052613b729190613b6981606081016110cd565b5190208b6141c5565b8060c08401510361124a578561010084015114801590613da5575b61124a57610100613b9d816124fe565b850361124a57613bd4613bbb61080096613bc2613bbb89838961390a565b3691613930565b60208151910120976114009187613918565b60208151910120613bf78d6001600160a01b0319165f52602b60205260405f2090565b540361124a57613c1192613c0a926144af565b9160e00190565b510361124a57613c9d91613c716004600593613c448c610cef33916001600160a01b0319165f52602560205260405f2090565b805461ffff60a01b191660a08d901b61ffff60a01b1617815590600382015501805460ff19166001179055565b018054613c849060101c61ffff1661356d565b63ffff000082549160101b169063ffff00001916179055565b613cd8613cbc876001600160a01b0319165f52602c60205260405f2090565b928354926001850193845480155f14613d9f5750600190614a28565b9255556001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561025057604051633c1bcdef60e21b8152336004820152925f908490602490829084905af192831561043d577f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb593613d8b575b506040805161ffff9095168552602085019190915283015233926001600160a01b0319169180606081016137f0565b806112b95f613d99936123ee565b5f613d5c565b90614a28565b50866101208401511415613b8d565b508660a08301511415613b42565b508760808301511415613b3b565b9050606083015114155f613b33565b909150613df16104d7604086016128b8565b1415905f613b2b565b91506020840151613e0e61ffff84166104d7565b141591613b24565b806112b95f613e24936123ee565b5f613b00565b6305d252c360e01b5f5260045ffd5b63652122d960e01b5f5260045ffd5b50613e5a601086901c61ffff166104d7565b8711613a29565b908160a09103126102505760405190613e7982612345565b8051613e8481610906565b8252602081015160208301526040810151604083015260608101519060038210156102505760809160608401520151613ebc81610aae565b608082015290565b6003111561206f57565b80546801000000000000000081101561236057613ef091600182018155613506565b6001600160a01b0380839493549260031b9316831b921b1916179055565b61ffff60019116019061ffff821161251457565b9061204f604051613f328161239c565b60e0613fb282955461ffff8116845261ffff8160101c16602085015261ffff808260201c16166040850152613f7561ffff8260301c16606086019061ffff169052565b61ffff604082901c1660808501526001600160401b03605082901c1660a08501526001600160401b03609082901c1660c085015260d01c60ff1690565b1515910152565b9061204f6140986005613fca612474565b94613fec613fdf82546001600160a01b031690565b6001600160a01b03168752565b613ff860018201613f22565b6020870152614050614040600283015461401e6140158260ff1690565b60408b01612d66565b6001600160401b03600882901c1660608a015260481c6001600160401b031690565b6001600160401b03166080880152565b600381015460a0870152600481015460c0870152015461ffff811660e086015261ffff601082901c1661010086015261ffff602082901c1661012086015260301c61ffff1690565b61ffff16610140840152565b5f516020614e945f395f51905f5291604051906001600160a01b031960208301931683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c8152614102606c826123ee565b5190200690565b5f516020614e945f395f51905f5291604051906001600160a01b031960208301931683527fc5cb4182e179e0279f50e2d772929d40dc9d4db3b30ec2ebbefbe6b9bb543075602c830152604c820152604c8152614102606c826123ee565b5f516020614e945f395f51905f5291604051906001600160a01b031960208301931683527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c8152614102606c826123ee565b5f516020614e945f395f51905f5291604051906001600160a01b031960208301931683527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c8152614102606c826123ee565b90919392946142439061040061423c620100008261255e565b9186613918565b9161426f6104d760016142626104d7600589015461ffff9060101c1690565b96015460101c61ffff1690565b92610800915f5b868110614299575050505050505090614292916108a0916144af565b0361124a57565b61ffff8160051b890135168015801561438d575b61124a5761430e6142f26142df6142d6866001600160a01b0319165f52602460205260405f2090565b610f5e856134f2565b90546001600160a01b039160031b1c1690565b610cef856001600160a01b0319165f52602560205260405f2090565b90614320610efb600484015460ff1690565b908115614370575b5061124a576003614359613bbb61433f888661252f565b6143518961434c88612542565b61252f565b90888b613918565b602081519101209101540361124a57600101614276565b90506143856104d7835461ffff9060a01c1690565b14155f614328565b508681116142ad565b9190916143a56104008561255e565b925f5b8381106143ec575050505b602081106143c057505050565b8060051b8084013515908115916143df575b5061124a576001016143b3565b905082013515155f6143d2565b8060051b61ffff818801351690811580156144a2575b61124a576144506144346142df61442b886001600160a01b0319165f52602460205260405f2090565b610f5e866134f2565b610cef876001600160a01b0319165f52602960205260405f2090565b91614462610efb600385015460ff1690565b908115614485575b5061124a576001908701359101540361124a576001016143a8565b905061449a6104d7845461ffff9060a01c1690565b14155f61446a565b5061ffff84168211614402565b9291905f516020614e945f395f51905f525f940691829060051b8201915b8281106144da5750505050565b909192945f516020614e945f395f51905f5283816020938186358b0990089709939291016144cd565b8054905f815581614512575050565b5f5260205f20908101905b818110614528575050565b5f815560010161451d565b6001600160a01b0361456761455a836001600160a01b0319165f52602260205260405f2090565b546001600160a01b031690565b161561488557614589816001600160a01b0319165f52602460205260405f2090565b8054905f5b82811061472e5750505060015b61010061ffff8216111561465f5750806145d86145d36001600160a01b0319936001600160a01b0319165f52602460205260405f2090565b614503565b5f6145f5826001600160a01b0319165f52602b60205260405f2090565b55614638614615826001600160a01b0319165f52602260205260405f2090565b60055f918281558260018201558260028201558260038201558260048201550155565b167f98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef265f80a2565b806146896104d761362c6146bc94610fab876001600160a01b0319165f52602760205260405f2090565b6146fb575b6146b36003610fbc83610fab876001600160a01b0319165f52602860205260405f2090565b6146c15761356d565b61459b565b6111b66146e482610fab866001600160a01b0319165f52602860205260405f2090565b60035f918281558260018201558260028201550155565b61472961471e82610fab866001600160a01b0319165f52602760205260405f2090565b805461ffff19169055565b61468e565b61473b6142df8284613506565b61476861475e82610cef886001600160a01b0319165f52602360205260405f2090565b805460ff19169055565b5f6147a4614788876001600160a01b0319165f52602a60205260405f2090565b6147946104d786612542565b61ffff165f5260205260405f2090565b556147e56147c882610cef886001600160a01b0319165f52602560205260405f2090565b60045f918281558260018201558260028201558260038201550155565b6148086146e482610cef886001600160a01b0319165f52602960205260405f2090565b60015b61010061ffff8216111561482357505060010161458e565b806148506004610fbc85610cef61485996610fab8d6001600160a01b0319165f52602660205260405f2090565b61485e5761356d565b61480b565b6111b66147c884610cef84610fab8c6001600160a01b0319165f52602660205260405f2090565b50565b90939291936148b86104d760016148ab6148a188612550565b976104809061255e565b97015460101c61ffff1690565b905f5b84811061490e57505050505b602081106148d457505050565b8060061b83018160051b83013561124a57803515908115916148fe575b5061124a576001016148c7565b600191506020013514155f6148f1565b8060061b870161ffff8260051b8801351690811580156149f8575b61124a576149726149526142df61442b896001600160a01b0319165f52602460205260405f2090565b610cef86610fab8a6001600160a01b0319165f52602660205260405f2090565b91614984610efb600485015460ff1690565b9081156149db575b5061124a57815460b01c61ffff1661ffff80861691160361124a576002820154813514918215926149c6575b505061124a576001016148bb565b60209192506003015491013514155f806149b8565b90506149f06104d7845461ffff9060a01c1690565b14155f61498c565b50848211614929565b600681101561206f576002149081614a17575090565b6001600160401b0391501643111590565b9392919091841580614b48575b614b4057811580614b36575b614b31575f516020614e945f395f51905f52828609945f516020614e945f395f51905f528285095f516020614e945f395f51905f528188095f516020614e945f395f51905f5290620292f809965f516020614e945f395f51905f5290620292fc09614aab91614de6565b935f516020614e945f395f51905f5287600108614ac790614e2d565b935f516020614e945f395f51905f529109915f516020614e945f395f51905f529109905f516020614e945f395f51905f529108905f516020614e945f395f51905f52910992614b1590614d9e565b614b1e90614e2d565b5f516020614e945f395f51905f52910990565b505090565b5060018114614a41565b935090509190565b5060018314614a35565b600681101561206f576001149081614a17575090565b60405190610400614b7981846123ee565b368337565b60405190610800614b7981846123ee565b906020811015612d615760051b0190565b906040811015612d615760051b0190565b91905f835b60208210614bf15750505061040082015f905b60408210614bdb57505050610c000190565b6020806001928551815201930191019091614bc9565b6020806001928551815201930191019091614bb6565b919091614c12614b68565b614c1a614b7e565b93614c37836001600160a01b0319165f52602460205260405f2090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165f5b61ffff84168110614ce05750505061ffff165b60208110614cbb57506110cd614c9c614cb8939495604051928391602083019586614bb1565b519020916001600160a01b0319165f52602b60205260405f2090565b55565b806001614cd9614cd3614cce8395612519565b612542565b88614ba0565b5201614c76565b80614ced614d3092612542565b614cf78288614b8f565b5260a0614d076142df8387613506565b6040516313a4120960e31b81526001600160a01b03909116600482015292839081906024820190565b0381865afa91821561043d576001926040915f91614d80575b506020810151614d61614d5b85612519565b8d614ba0565b520151614d79614d73614cce84612519565b8b614ba0565b5201614c63565b614d98915060a03d8111611f2457611f1681836123ee565b5f614d49565b5f516020614e945f395f51905f5290065f516020614e945f395f51905f52035f516020614e945f395f51905f528111612514575f516020614e945f395f51905f529060010890565b905f516020614e945f395f51905f5290065f516020614e945f395f51905f52035f516020614e945f395f51905f528111612514575f516020614e945f395f51905f52910890565b60405190602082526020808301526020604083015260608201527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff60808201525f516020614e945f395f51905f5260a082015260208160c08160055afa1561025057519056fe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220129c2f76af753ff037ef7a271312e4da4c151c0215991f808754d181afcad25b64736f6c634300081c0033",
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
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bytes32,bytes32,bool))
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
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bytes32,bytes32,bool))
func (_DKGManager *DKGManagerSession) GetCombinedDecryption(roundId [12]byte, ciphertextIndex uint16) (DKGTypesCombinedDecryptionRecord, error) {
	return _DKGManager.Contract.GetCombinedDecryption(&_DKGManager.CallOpts, roundId, ciphertextIndex)
}

// GetCombinedDecryption is a free data retrieval call binding the contract method 0xfe234897.
//
// Solidity: function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex) view returns((uint16,bytes32,bytes32,bool))
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
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
func (_DKGManager *DKGManagerSession) GetRound(roundId [12]byte) (IDKGManagerRound, error) {
	return _DKGManager.Contract.GetRound(&_DKGManager.CallOpts, roundId)
}

// GetRound is a free data retrieval call binding the contract method 0xf4e34945.
//
// Solidity: function getRound(bytes12 roundId) view returns((address,(uint16,uint16,uint16,uint16,uint16,uint64,uint64,bool),uint8,uint64,uint64,bytes32,uint256,uint16,uint16,uint16,uint16))
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

// CombineDecryption is a paid mutator transaction binding the contract method 0x9f431549.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) CombineDecryption(opts *bind.TransactOpts, roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintextHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "combineDecryption", roundId, ciphertextIndex, combineHash, plaintextHash, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0x9f431549.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) CombineDecryption(roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintextHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, roundId, ciphertextIndex, combineHash, plaintextHash, transcript, proof, input)
}

// CombineDecryption is a paid mutator transaction binding the contract method 0x9f431549.
//
// Solidity: function combineDecryption(bytes12 roundId, uint16 ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) CombineDecryption(roundId [12]byte, ciphertextIndex uint16, combineHash [32]byte, plaintextHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.CombineDecryption(&_DKGManager.TransactOpts, roundId, ciphertextIndex, combineHash, plaintextHash, transcript, proof, input)
}

// CreateRound is a paid mutator transaction binding the contract method 0x62c41927.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed) returns(bytes12)
func (_DKGManager *DKGManagerTransactor) CreateRound(opts *bind.TransactOpts, threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "createRound", threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed)
}

// CreateRound is a paid mutator transaction binding the contract method 0x62c41927.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed) returns(bytes12)
func (_DKGManager *DKGManagerSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed)
}

// CreateRound is a paid mutator transaction binding the contract method 0x62c41927.
//
// Solidity: function createRound(uint16 threshold, uint16 committeeSize, uint16 minValidContributions, uint16 lotteryAlphaBps, uint16 seedDelay, uint64 registrationDeadlineBlock, uint64 contributionDeadlineBlock, bool disclosureAllowed) returns(bytes12)
func (_DKGManager *DKGManagerTransactorSession) CreateRound(threshold uint16, committeeSize uint16, minValidContributions uint16, lotteryAlphaBps uint16, seedDelay uint16, registrationDeadlineBlock uint64, contributionDeadlineBlock uint64, disclosureAllowed bool) (*types.Transaction, error) {
	return _DKGManager.Contract.CreateRound(&_DKGManager.TransactOpts, threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed)
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
	PlaintextHash   [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecryptionCombined is a free log retrieval operation binding the contract event 0x451276810ef520579055672046d83aad5adae5e72513ec6b904ac15cd4496115.
//
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash)
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

// WatchDecryptionCombined is a free log subscription operation binding the contract event 0x451276810ef520579055672046d83aad5adae5e72513ec6b904ac15cd4496115.
//
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash)
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

// ParseDecryptionCombined is a log parse operation binding the contract event 0x451276810ef520579055672046d83aad5adae5e72513ec6b904ac15cd4496115.
//
// Solidity: event DecryptionCombined(bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash)
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
