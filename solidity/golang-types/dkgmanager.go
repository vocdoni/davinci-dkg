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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contributionVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_partialDecryptVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalizeVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decryptCombineVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealSubmitVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_revealShareVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRIBUTION_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECRYPT_COMBINE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PARTIAL_DECRYPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SHARE_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVEAL_SUBMIT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROUND_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"abortRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimSlot\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"combineDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRound\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendRegistration\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCombinedDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.CombinedDecryptionRecord\",\"components\":[{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"completed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.ContributionRecord\",\"components\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"commitmentVectorDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContributionVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecryptCombineVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizeVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryptVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.PartialDecryptionRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delta\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.Point\",\"components\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealShareVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealSubmitVerifierVKeyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RevealedShareRecord\",\"components\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGManager.Round\",\"components\":[{\"name\":\"organizer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"policy\",\"type\":\"tuple\",\"internalType\":\"structDKGTypes.RoundPolicy\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"committeeSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minValidContributions\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lotteryAlphaBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seedDelay\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"registrationDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"contributionDeadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"disclosureAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDKGTypes.RoundStatus\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contributionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"partialDecryptionCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"revealedShareCount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getShareCommitmentHash\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reconstructSecret\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roundNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectedParticipants\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitContribution\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transcript\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitPartialDecryption\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRevealedShare\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"internalType\":\"bytes12\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"shareValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContributionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contributorIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"commitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"encryptedSharesHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecryptionCombined\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"combineHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"plaintextHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialDecryptionSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"ciphertextIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"deltaHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistrationClosed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevealedShareSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participantIndex\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"shareHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundAborted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundCreated\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"organizer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"seedBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"lotteryThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundEvicted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundFinalized\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"aggregateCommitmentsHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"collectivePublicKeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"shareCommitmentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecretReconstructed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"disclosureHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"reconstructedSecretHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SeedResolved\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"seed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotClaimed\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"bytes12\",\"indexed\":true,\"internalType\":\"bytes12\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyCombined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyContributed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPartiallyDecrypted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisclosureDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientContributions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPartialDecryptions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRevealedShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCombinedDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitteeSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinalization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPartialDecryption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPolicy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReconstruction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRevealedShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEligible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelectedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SeedNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlotsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x6101a08060405234610253576101008161453f80380380916100218285610257565b8339810103126102535780519063ffffffff82168203610253576100476020820161028e565b6100536040830161028e565b61005f6060840161028e565b61006b6080850161028e565b9161007860a0860161028e565b9361009160e061008a60c0890161028e565b970161028e565b966001600160a01b038316158015610242575b8015610231575b8015610220575b801561020f575b80156101fe575b6101ef5763ffffffff918160805260a05260405160208101918360e01b9060e01b1682523060601b6024820152601881526100fc603882610257565b5190201660c05260e052610100526101205261014052610160526101805260405161429c90816102a3823960805181611dc3015260a05181818161066801528181611619015281816125ed0152613649015260c05181818161269e0152612cc7015260e0518181816113a8015281816124e401526135a401526101005181818161127701528181611d650152611f56015261012051818181611d1a01528181612483015261377601526101405181818161026c0152818161192701526122bf015261016051818181610ead0152818161181b01526136050152610180518181816111d201528181612c8601526130400152f35b63baa3de5f60e01b5f5260045ffd5b506001600160a01b038816156100c0565b506001600160a01b038716156100b9565b506001600160a01b038616156100b2565b506001600160a01b038516156100ab565b506001600160a01b038416156100a4565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761027a57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102535756fe6080806040526004361015610012575f80fd5b5f905f3560e01c908163058994a11461366d5750806306433b1b14613629578063070c7492146135e5578063074a75e11461357d5780630b1451f0146133925780630e2c53f714612f33578063349181a214612e4d578063415a1b8614612e27578063510ba2df14612dd657806353d7218414612ceb57806356664d1514612caa5780635ddd062614612c6657806362c419271461250857806363f314cd146124c4578063669a76a91461245c57806370f2469b1461230057806372517b4b14612298578063802ae23114611de757806385e1f4d014611da65780638dc1f53a14611d3e57806393c3d3a814611cfa5780639f4315491461185c578063b18730c2146117f4578063b7bca6151461129b578063bf19220914611257578063c2440e16146111ab578063c9396bf014610d42578063ca3c045814610c7a578063d3720aac14610b76578063d993376714610599578063f4e3494514610290578063fe1604b51461024c5763fe23489714610189575f80fd5b346102495760403660031901126102495761ffff60406101a7613c72565b926001600160a01b03196101b9613cb6565b9482606085516101c881613e15565b8281528260208201528287820152015216815260286020522091165f52602052608060405f206040516101fa81613e15565b61ffff82541691828252600181015460208301908152606060ff600360028501549460408701958652015416930192151583526040519384525160208401525160408301525115156060820152f35b80fd5b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610249576020366003190112610249576001600160a01b03196102b3613c72565b826101406040516102c381613dc3565b8281526040516102d281613da7565b8381528360208201528360408201528360608201528360808201528360a08201528360c08201528360e082015260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152826101008201528261012082015201521681526022602052604081206040519061035182613dc3565b6001600160a01b0381541682526040519061036b82613da7565b60ff600182015461ffff8116845261ffff8160101c16602085015261ffff8160201c16604085015261ffff8160301c16606085015261ffff8160401c1660808501526001600160401b038160501c1660a08501526001600160401b038160901c1660c085015260d01c16151560e08301526020830191825260028101549360ff8516916040850192600681101561058557835260608501956001600160401b038160081c1687526001600160401b03608087019160481c16815260038201549060a08701918252600560048401549360c0890194855201549760e088019461ffff8a16865261010089019661ffff8b60101c16885260e06101208b019961ffff8d60201c168b5261ffff6101408d019d60301c168d526001600160a01b036040519c51168c525161ffff81511660208d015261ffff60208201511660408d015261ffff60408201511660608d015261ffff60608201511660808d015261ffff60808201511660a08d01526001600160401b0360a08201511660c08d01526001600160401b0360c082015116828d0152015115156101008b0152519060068210156105715750610120890152516001600160401b0390811661014089015290511661016087015251610180860152516101a08501525161ffff9081166101c0850152905181166101e084015290518116610200830152915190911661022082015261024090f35b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b83526021600452602483fd5b5034610249576020366003190112610249576001600160a01b03196105bc613c72565b168082526022602052604082206001600160a01b0381541615610b6757600281019081549060ff821660018201918254916006811015610b535760011480610b3d575b15610b2e57600581019061ffff808354169360101c16831015610b1f578688526023602052604088206001600160a01b0333165f5260205260ff60405f205416610b1057600381018054958615610a9d575b50506040516313a4120960e31b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316959060a0816024818a5afa8015610a05576060918b91610a7e575b5001516003811015610a6a575f1901610a5b57600490604051602081019182523360601b6040820152603481526106e3605482613e30565b5190209101541115610a4c5785875260246020526040872080549068010000000000000000821015610a38579261ffff9594939261072b83889560016107b396018155613f94565b81549060031b906001600160a01b0333831b921b1916179055888a52602360205260408a206001600160a01b0333165f5260205260405f20600160ff198254161790558361077883614080565b168419825416179055604051818152887f80d59d7599daf0493f96a2d1016163c29d85f5e4a8b59f3001f6e9a115a6c96b60203393a3614080565b915460101c16921682146107c5578480f35b604051916102006107d68185613e30565b368437604051916104006107ea8185613e30565b368437858752602460205260408720875b838110610904575050505b601081106108be57506040519060208201928387905b601082106108a857505050610220820186905b6020821061089257505050610600815261084b61062082613e30565b519020828452602b6020526040842055805460ff191660021790557fca89d7e15807c1ba6a0622215afe84b083f061c44c2e78e6e226709a8f5f13588280a25f8080808480f35b602080600192855181520193019101909161082f565b602080600192855181520193019101909161081c565b8060011b818104600214821517156108f057600181018091116108f0579060016108e9819385614255565b5201610806565b634e487b7160e01b87526011600452602487fd5b60018101808211610a24576010821015610a10578160051b8701526001600160a01b036109318284613f94565b90549060031b1c16604051906313a4120960e31b8252600482015260a081602481875afa908115610a05578a916109d7575b506020810151908260011b91838304600214841517156109c3579060409161098b848a614255565b52015190600181018091116109af57906109a86001939288614255565b52016107fb565b634e487b7160e01b8b52601160045260248bfd5b634e487b7160e01b8c52601160045260248cfd5b6109f8915060a03d81116109fe575b6109f08183613e30565b810190614019565b5f610963565b503d6109e6565b6040513d8c823e3d90fd5b634e487b7160e01b8a52603260045260248afd5b634e487b7160e01b8a52601160045260248afd5b634e487b7160e01b89526041600452602489fd5b637c75aa6f60e11b8752600487fd5b63aba4733960e01b8952600489fd5b634e487b7160e01b8a52602160045260248afd5b610a97915060a03d60a0116109fe576109f08183613e30565b5f6106ab565b6001600160401b0391965060481c1680431115610b015740948515610af257859055867fc16e97da5706abead845583dfc2e6126862a0c07801be8ac6027010b50139652602087604051908152a25f80610651565b6302504bb360e61b8952600489fd5b63172181cb60e21b8952600489fd5b630c8d9eab60e31b8852600488fd5b63848084dd60e01b8852600488fd5b63268dbf6760e21b8752600487fd5b506001600160401b038260501c164311156105ff565b634e487b7160e01b88526021600452602488fd5b6328ad4a9560e21b8352600483fd5b5034610249576040366003190112610249576001600160a01b036040610b9a613c72565b926001600160a01b0319610bac613cd8565b948260a08551610bbb81613ddf565b8281528260208201528287820152826060820152826080820152015216815260256020522091165f5260205260c060405f20604051610bf981613ddf565b8154916001600160a01b0383169283835261ffff602084019160a01c16815260018201546040840190815261ffff6002840154926060860193845260a060ff600460038801549760808a01988952015416960195151586526040519687525116602086015251604085015251606084015251608083015251151560a0820152f35b5034610249576020366003190112610249576001600160a01b0319610c9d613c72565b168152602460205260408120604051908160208254918281520190819285526020852090855b818110610d235750505082610cd9910383613e30565b604051928392602084019060208552518091526040840192915b818110610d01575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610cf3565b82546001600160a01b0316845260209093019260019283019201610cc3565b50346102495760a036600319011261024957610d5c613c72565b610d64613cb6565b906044356064356001600160401b0381116111a757610d87903690600401613c89565b936084356001600160401b0381116111a357610da7903690600401613c89565b946001600160a01b0319811695868952602260205260408920946001600160a01b03865416156111945760018601549860ff8a60d01c16156111855760ff60028801541660068110156111715760030361116257888b52602360205260408b206001600160a01b0333165f5260205260ff60405f205416156111535761ffff8616998a15908115611142575b50801561113a575b61112b57888b5260246020526001600160a01b03610e6560408d20610e5f89613f80565b90613f94565b9190913392549060031b1c160361111c57888b52602960205260408b206001600160a01b0333165f5260205260ff600360405f2001541661110d57908a916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561110957849284928892610efb60405196879586948594635c73957b60e11b865260048601613e71565b03915afa80156110fe576110e9575b5050820160a0838203126110e55780601f840112156110e55760405192610f3260a085613e30565b839060a081019283116110cd57905b8282106110bd57505050858852602a60205260408820875f5260205260405f20549082519060a01c14918215926110ae575b821561109f575b8215611096575b821561105e575b505061104f5761101b91836001600593878a52602960205260408a206001600160a01b0333165f52602052610fd860405f2091829061ffff60a01b1961ffff60a01b83549260a01b169116179055565b600381018260ff19825416179055015501610ffa61ffff825460301c16613fa9565b67ffff00000000000082549160301b169067ffff0000000000001916179055565b60405192835260208301527f5f16f25c2c3e0004ddf0924462825a82df1cd393546962d8a9757d1e71dc1dac60403393a380f35b63d1fed5fd60e01b8652600486fd5b90915060806060820151910151604051906020820192835260408201526040815261108a606082613e30565b51902014155f80610f88565b81159250610f81565b60408101518614159250610f7a565b60208101518814159250610f73565b8135815260209182019101610f41565b8a80fd5b634e487b7160e01b5f52604160045260245ffd5b8880fd5b816110f391613e30565b6110e557885f610f0a565b6040513d84823e3d90fd5b8380fd5b63a89ac15160e01b8b5260048bfd5b63d1fed5fd60e01b8b5260048bfd5b639eae062d60e01b8b5260048bfd5b508715610e3b565b61ffff915060101c168a115f610e33565b63965c290d60e01b8b5260048bfd5b63268dbf6760e21b8b5260048bfd5b634e487b7160e01b8c52602160045260248cfd5b630ba0cb2f60e21b8b5260048bfd5b6328ad4a9560e21b8a5260048afd5b8680fd5b8480fd5b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b5790611214575b602090604051908152f35b506020813d602011611243575b8161122e60209383613e30565b8101031261123f5760209051611209565b5f80fd5b3d9150611221565b604051903d90823e3d90fd5b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610249576112aa36613cee565b949091989793926001600160a01b0319899897981698898c52602260205260408c20966001600160a01b03885416156117e55760ff6002890154169260018901549360068110156117d157600214806117bb575b156117ac5760408e8d81526023602052206001600160a01b0333165f5260205260ff60405f2054161561179d5761ffff88169c8d15801561178d575b61177e576001600160a01b038f60408f91611360928152602460205220610e5f8c613f80565b9190913392549060031b1c160361176f5760408f8e81526025602052206001600160a01b0333165f5260205260ff600460405f2001541661176057908e916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561110957849284928a926113f660405196879586948594635c73957b60e11b865260048601613e71565b03915afa80156110fe5761174b575b50508401610100858203126117475780601f86011215611747576040519461142f61010087613e30565b8590610100810192831161174357905b8282106117335750505083519060a01c1490811591611720575b8115611709575b5080156116fb575b80156116ed575b80156116df575b6116d0577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016040516020810190888252896040820152604081526114bb606082613e30565b51902060405160208101918b83527f29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72602c830152604c820152604c8152611503606c82613e30565b519020068060c08401510361111c57611000840361111c576104008481116116cc57611530368285613fd4565b6020815191012094610a009081831161123f57811161123f578161155992850191033691613fd4565b60208151910120898c52602b60205260408c20540361111c57608060e09261158092614094565b910151036116bd5791600460059261160e94888b52602560205260408b206001600160a01b0333165f526020526115d260405f2092839061ffff60a01b1961ffff60a01b83549260a01b169116179055565b600382015501600160ff19825416179055016115f561ffff825460101c16613fa9565b63ffff000082549160101b169063ffff00001916179055565b846001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156116b957818091602460405180948193633c1bcdef60e21b83523360048401525af180156110fe576116a4575b5050604051938452602084015260408301527f8f25a636f27af2671bfd0f5c59da52b0495e5415d2e605b2d0994830aba13fb560603393a380f35b816116ae91613e30565b6111a757845f611669565b5080fd5b63d1fed5fd60e01b8852600488fd5b8b80fd5b63d1fed5fd60e01b8a5260048afd5b508660a08301511415611476565b50856080830151141561146f565b508860608301511415611468565b905061ffff60408401519160101c1614155f611460565b602084015161ffff821614159150611459565b813581526020918201910161143f565b8e80fd5b8c80fd5b8161175591613e30565b611747578c5f611405565b6305d252c360e01b8f5260048ffd5b63d1fed5fd60e01b8f5260048ffd5b63652122d960e01b8f5260048ffd5b508d61ffff8660101c161061133a565b63965c290d60e01b8e5260048efd5b63268dbf6760e21b8e5260048efd5b506001600160401b038460901c164311156112fe565b634e487b7160e01b8f52602160045260248ffd5b6328ad4a9560e21b8d5260048dfd5b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b579061121457602090604051908152f35b50346102495761186b36613cee565b6001600160a01b03198a9695949a9992991698898c52602260205260408c20926001600160a01b03845416156117e55760ff6002850154166006811015611ce657600303611cd75761ffff169a8b158015611ccf575b8015611cc7575b611cb8578a8d52602760205260408d208c5f52602052600161ffff60405f2054169401549461ffff8616809510611ca95760408e8d81526028602052208d5f5260205260ff600360405f20015416611c9a57908d916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156111095785928492869261197560405196879586948594635c73957b60e11b865260048601613e71565b03915afa80156110fe57611c85575b505061199291810190613eee565b9485519060a01c14801590611c77575b8015611c69575b8015611c5b575b6116d05760408501908151106116d0577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604051602081019089825288604082015260408152611a01606082613e30565b51902060405160208101918b83527fb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb602c830152604c820152604c8152611a49606c82613e30565b51902006928360a08701510361111c5761067f19016116d0575190608084018411610a245761028084018411610a245760101c61ffff16895b828110611b5f5750505b60108110611b0f5750603460c092611aa392614094565b91015103611b00577f451276810ef520579055672046d83aad5adae5e72513ec6b904ac15cd4496115916040918487526028602052828720865f526020526003835f2001600160ff1982541617905582519182526020820152a380f35b63d1fed5fd60e01b8552600485fd5b6102808160061b84010160808260051b850101356116d05780351590811591611b4f575b50611b4057600101611a8c565b63d1fed5fd60e01b8952600489fd5b600191506020013514155f611b33565b6102808160061b86010161ffff60808360051b88010135169081158015611c52575b611c30578a8d5260246020528c60406001600160a01b03611ba7828420610e5f87613f80565b90549060031b1c16918d81526026602052208d5f526020526001600160a01b0360405f2091165f5260205260405f209160ff60048401541615908115611c3f575b50611c30578b61ffff835460b01c1603611c3057600282015481351491821592611c1b575b505061111c57600101611a82565b60209192506003015491013514155f80611c0d565b63d1fed5fd60e01b8d5260048dfd5b905061ffff835460a01c1614155f611be8565b50838211611b81565b5085608086015114156119b0565b5086606086015114156119a9565b5080602086015114156119a2565b81611c8f91613e30565b6116cc578b5f611984565b63955c0c4960e01b8e5260048efd5b63032cddf960e11b8e5260048efd5b636d28699160e01b8d5260048dfd5b5088156118c8565b5089156118c1565b63268dbf6760e21b8d5260048dfd5b634e487b7160e01b8e52602160045260248efd5b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b579061121457602090604051908152f35b5034610249578060031936011261024957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102495760c036600319011261024957611e01613c72565b611e09613cb6565b90611e12613cc7565b606435906084356001600160401b03811161229457611e35903690600401613c89565b9460a4356001600160401b03811161229057611e55903690600401613c89565b956001600160a01b0319811696878a52602260205260408a20946001600160a01b03865416156122815760ff60028701541660068110156111715760030361116257888b52602360205260408b206001600160a01b0333165f5260205260ff60405f205416156111535761ffff8516998a15801561226d575b8015612261575b8015612259575b61224b576001600160a01b03611f0060408e8d8152602460205220610e5f89613f80565b9190913392549060031b1c160361223c57898c52602660205260408c2061ffff89165f5260205260405f206001600160a01b0333165f5260205260ff600460405f2001541661222d57908b916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561110957849284928892611fa460405196879586948594635c73957b60e11b865260048601613e71565b03915afa80156110fe57612218575b505082016101a0838203126122145780601f840112156122145760405192611fdd6101a085613e30565b83906101a081019283116116cc57905b82821061220457505050868952602a60205260408920885f5260205260405f20549082519060a01c14908115916121f5575b81156121ec575b81156121b6575b506116bd5760c0810160e081519201918251604051906020820192835260408201526040815261205e606082613e30565b5190208603611b40579261214e92600360059361ffff97968a8d52602660205260408d208989165f5260205260405f206001600160a01b0333165f526020526120c260405f2094859061ffff60a01b1961ffff60a01b83549260a01b169116179055565b83547fffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffff77ffff000000000000000000000000000000000000000000008a60b01b16911617845560048401600160ff19825416179055516002840155519101550161213184825460201c16613fa9565b65ffff0000000082549160201b169065ffff000000001916179055565b8386526027602052604086208282165f5260205260405f208261217381835416613fa9565b16831982541617905560405194855216602084015260408301527f39e01752de5471ef06952341613214369ee48b9bf21f57f7d6fcf9239f397f2260603393a380f35b9050608082015160a083015160405190602082019283526040820152604081526121e1606082613e30565b51902014155f61202d565b80159150612026565b6020830151891415915061201f565b8135815260209182019101611fed565b8980fd5b8161222291613e30565b61221457895f611fb3565b633466526160e01b8c5260048cfd5b63d1fed5fd60e01b8c5260048cfd5b62d949df60e51b8c5260048cfd5b508815611edc565b5061ffff881615611ed5565b5061ffff600188015460101c168b11611ece565b6328ad4a9560e21b8b5260048bfd5b8780fd5b8580fd5b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b579061121457602090604051908152f35b50346102495760603660031901126102495761231a613c72565b61ffff6040612327613cd8565b936001600160a01b0319612339613cc7565b948260a0855161234881613ddf565b8281528260208201528287820152826060820152865161236781613dfa565b8381528360208201526080820152015216815260266020522091165f526020526001600160a01b0360405f2091165f5260205260e060405f20604051906123ad82613ddf565b60208154916001600160a01b03831684528184019261ffff8160a01c16845261ffff604086019160b01c16815261ffff600183015491606087019283528160ff6004604051966123fc88613dfa565b6002810154885260038101548989015260808b019788520154169660a0890197151588526001600160a01b036040519951168952511685880152511660408601525160608501525180516080850152015160a083015251151560c0820152f35b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b579061121457602090604051908152f35b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461024957610100366003190112610249576004359061ffff821680920361024957612533613cb6565b9061253c613cc7565b906064359361ffff85168095036116b9576084359461ffff86168603612c625760a4356001600160401b0381168091036111095760c435906001600160401b0382168092036111a75760e435928315158094036122945784158015612c56575b8015612c49575b8015612c3d575b8015612c2c575b8015612c21575b8015612c15575b8015612c06575b8015612be4575b8015612bda575b612bcb57604051634331ed1f60e01b81526020816004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa8015612bc0578790612b80575b6001600160401b03915016978815612b715761ffff16976126458983613e9b565b6127108110612b0c5750505f19965b86546001600160401b0381166001600160401b038114612af857906001600160401b03600161ffff9493011680916001600160401b0319161789556bffffffff00000000000000007f000000000000000000000000000000000000000000000000000000000000000060401b1617998a6126e560215460406001600160401b03603f831692161015612ac057613f53565b6bffffffffffffffffffffffff829392549160031b92831b921b19161790556021546001600160401b0360018183160116906001600160401b031916176021556040519761273289613da7565b88526020880152166040860152606085015261ffff8816608085015260a084015260c083015260e08201526001600160401b038254169161278061ffff87166001600160401b034316613ece565b926040519361278e85613dc3565b33855260208501938452604085019060018252606086019283526001600160401b03608087019116815260a086019284845260c08701948886528060e08901528061010089015280610120890152806101408901526001600160a01b03198a60a01b168152602260205260408120966001600160a01b038951166001600160a01b031989541617885560018801905161ffff808251161661ffff1983541617825561285661ffff602083015116839063ffff000082549160101b169063ffff00001916179055565b6040810151825465ffff00000000191660209190911b65ffff00000000161782556060810151825467ffff000000000000191660309190911b67ffff000000000000161782556080810151825460a083015171ffffffffffffffffffff00000000000000001990911660409290921b69ffff0000000000000000169190911760509190911b71ffffffffffffffff000000000000000000001617825560c0810151907fffffffffff000000000000000000ffffffffffffffffffffffffffffffffffff79ffffffffffffffff0000000000000000000000000000000000007aff000000000000000000000000000000000000000000000000000060e08654940151151560d01b169360901b16911617179055600287019351906006821015610571575083549151925170ffffffffffffffffffffffffffffffffff1990921660ff919091161760089290921b68ffffffffffffffff00169190911760489190911b70ffffffffffffffff0000000000000000001617905551600383015551600482015560e0820151600590910180546101008401516101208501516101409095015167ffff00000000000060309190911b1667ffffffffffffffff1990921661ffff9485161763ffff000060109290921b919091161765ffff00000000602095861b16171790559093612a5391166001600160401b034316613ece565b7fcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d760405180612aa833956001600160a01b03198860a01b169583602090939291936001600160401b0360408201951681520152565b0390a36001600160a01b03196040519160a01b168152f35b612ac981613f53565b90549060031b1c60a01b6001600160a01b03198116612ae9575b50613f53565b612af29061410e565b5f612ae3565b634e487b7160e01b89526011600452602489fd5b807e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa4302907e068db8bac710cb295e9e1b089a027525460aa64c2f837b4a2339c0ebedfa43820403612b5d570496612654565b634e487b7160e01b88526011600452602488fd5b63d06b96b160e01b8752600487fd5b506020813d602011612bb8575b81612b9a60209383613e30565b810103126111a357612bb36001600160401b0391613f3f565b612624565b3d9150612b8d565b6040513d89823e3d90fd5b63d06b96b160e01b8652600486fd5b50818311156125d4565b506001600160401b03612bfd61ffff8b16824316613ece565b168211156125cd565b5061010061ffff8a16116125c6565b5061ffff8916156125bf565b5061271081106125b8565b5061ffff881661ffff8816116125b1565b5061ffff8716156125aa565b5061ffff881685116125a3565b5061ffff88161561259c565b8280fd5b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610249578060031936011261024957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610249576040366003190112610249576001600160a01b036040612d0f613c72565b926001600160a01b0319612d21613cd8565b948260808551612d3081613d8c565b8281528260208201528287820152826060820152015216815260296020522091165f5260205260a060405f20604051612d6881613d8c565b8154916001600160a01b0383169283835261ffff6020840191861c16815261ffff60018301549160408501928352608060ff60036002870154966060890197885201541695019415158552604051958652511660208501525160408401525160608301525115156080820152f35b50346102495760403660031901126102495761ffff6040612df5613c72565b926001600160a01b0319612e07613cb6565b94168152602a6020522091165f52602052602060405f2054604051908152f35b50346102495780600319360112610249576001600160401b036020915416604051908152f35b5034610249576020366003190112610249576001600160a01b0319612e70613c72565b168082526022602052604082206001600160a01b038154168015612f24573303612f165760020160ff8154166006811015612f025760058114908115612ef7575b50612ee857805460ff191660041790557f97d5ddda8e4d1dcdb9643b144637aeef99ca0f2efe328a2b8e2620776cf1e4108280a280f35b63268dbf6760e21b8352600483fd5b60049150145f612eb1565b634e487b7160e01b84526021600452602484fd5b6282b42960e81b8352600483fd5b6328ad4a9560e21b8452600484fd5b50346102495760c036600319011261024957612f4d613c72565b6024356044356064356001600160401b0381116111a757612f72903690600401613c89565b90936084356001600160401b0381116111a357612f93903690600401613c89565b91909560a4356001600160401b0381116110e557612fb5903690600401613c89565b9790936001600160a01b0319841698898b52602260205260408b20956001600160a01b03875416156133835760018701549360ff8560d01c161561337457600288019760ff89541660068110156117d1576003036117ac578b15801561336c575b61335d57600561ffff91015460301c169361ffff861680951061334e57908d916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156111095785928492869261308e60405196879586948594635c73957b60e11b865260048601613e71565b03915afa80156110fe57613339575b50506130ab91810190613eee565b9384519060a01c1480159061332b575b801561331d575b801561330f575b6116d05760408401908151106116d0577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160405160208101908a82528960408201526040815261311a606082613e30565b51902060405160208101918c83527fc5cb4182e179e0279f50e2d772929d40dc9d4db3b30ec2ebbefbe6b9bb543075602c830152604c820152604c8152613162606c82613e30565b51902006918260a08601510361111c576104006020970361111c57905161020084019160101c61ffff168285116132fb5790879392918c5b82811061324e5750505b6010811061321a575050916131b892614094565b60c08201510361320b576080849101510361104f57917fbc874c1da78f7646af98f19f16267e99af67534204f56955055f5a5a2d73b4829391604093600560ff198254161790558351928352820152a280f35b63d1fed5fd60e01b8752600487fd5b809192935060051b808501351590811591613241575b5061111c57600101908692916131a4565b905082013515155f613230565b8091929394955060051b61ffff818801351690811580156132f2575b61176f578d8f60408160298f856001600160a01b0395526024815284613295858520610e5f8b613f80565b90549060031b1c16958352522091165f528b5260405f209160ff600384015416159081156132df575b5061176f5760019086013591015403611c305760010190889493929161319a565b905061ffff835460a01c1614155f6132be565b5083821161326a565b634e487b7160e01b5f52601160045260245ffd5b5086608085015114156130c9565b5087606085015114156130c2565b5080602085015114156130bb565b8161334391613e30565b6116cc578b5f61309d565b63957674fd60e01b8e5260048efd5b6314141ce560e21b8e5260048efd5b508a15613016565b630ba0cb2f60e21b8d5260048dfd5b6328ad4a9560e21b8c5260048cfd5b5034610249576020366003190112610249576001600160a01b03196133b5613c72565b16808252602260205260408220906001600160a01b0382541615610b67576002820191825460ff811660068110156135695760010361355a5761ffff600583015416600183019182549161ffff8360101c1614610b2e576001600160401b038260501c16918243111561354b576001600160401b036134756135019461346f7fcba424d4ca0c24cfd724662848b8cf062529c48daf9734f804ebcfa51f5ea8d798979561ffff6134cc9660401c1694859160481c16613eae565b90613eae565b908960038701556134c76134936001600160401b0343169283613ece565b8a5470ffffffffffffffff000000000000000000191660489190911b70ffffffffffffffff00000000000000000016178a55565b613ece565b71ffffffffffffffff0000000000000000000082549160501b169071ffffffffffffffff000000000000000000001916179055565b60046001600160401b036001600160a01b03835416955460481c169101549061354560405192839283602090939291936001600160401b0360408201951681520152565b0390a380f35b63268dbf6760e21b8852600488fd5b63268dbf6760e21b8552600485fd5b634e487b7160e01b86526021600452602486fd5b503461024957806003193601126102495760405163233ace1160e01b8152906020826004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561124b579061121457602090604051908152f35b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461024957806003193601126102495760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b90503461123f5760e036600319011261123f57613688613c72565b906084356001600160401b03811161123f576136a8903690600401613c89565b60a4929192356001600160401b03811161123f576136ca903690600401613c89565b9160c4356001600160401b03811161123f576136ea903690600401613c89565b6001600160a01b031988949294165f52602260205260405f20956001600160a01b0387541615613c635760ff6002880154166006811015613c4f5760038114613c405760011901613c315761ffff600588015460101c169260018801549661ffff8860201c168510613c2257602435158015613c18575b8015613c0e575b613bff576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561123f5783806137bb878b83965f98635c73957b60e11b865260048601613e71565b03915afa8015613bf457613bdf575b508301610120848203126110e55780601f850112156110e557604051936137f361012086613e30565b849061012081019283116110cd57905b828210613bcf5750505082518760a01c14801590613bbd575b8015613ba8575b8015613b9a575b8015613b8a575b8015613b7a575b8015613b6a575b6116bd577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001604051602081019060243582526044356040820152606435606082015260608152613890608082613e30565b51902060405160208101916001600160a01b03198b1683527f7c20af5072936dabc40921b055b4668149175807f325ff0242bb400c2c186a39602c830152604c820152604c81526138e2606c82613e30565b51902006918260e085015103611b405761025093614a0082036116d0576101000151946142009161020090831161123f57908189949695939285019203908b5b878110613a2a57505050506139379350614094565b03613a1b5760028101805460ff191660031790556005015460101c61ffff169061460081018082116132fb57845b8381106139b95785857f5f329a28ec91a18b4e7904324a3c08646c7c3d433fe5e4a36c788af074a7245560606001600160a01b0319604051936024358552604435602086015260643560408601521692a280f35b8060019160061b83016040516020808201928035845201356040820152604081526139e5606082613e30565b5190206001600160a01b031987168852602a6020526040882061ffff808460051b88013516165f5260205260405f205501613965565b63d1fed5fd60e01b8452600484fd5b909192939596945061ffff8160051b8b013516158015613b50575b61223c578a6001600160a01b0360408e6001600160a01b0319838f613a849083881685526024602052610e5f61ffff878720928b60051b013516613f80565b90549060031b1c169416815260256020522091165f5260205260405f2060ff600482015416158015613b34575b611c3057610400820282810461040014831517156132fb5760018301808411613b205761040081029080820461040014901517156132fb57613afa613b0191600393888a613fbc565b3691613fd4565b602081519101209101540361223c576001019089949695939291613922565b634e487b7160e01b8f52601160045260248ffd5b5061ffff8260051b8c01351661ffff825460a01c161415613ab1565b5061ffff8260101c1661ffff8260051b8c01351611613a45565b5060643560c0840151141561383f565b5060443560a08401511415613838565b5060243560808401511415613831565b50806060840151141561382a565b50604083015161ffff8560101c161415613823565b50602083015161ffff8516141561381c565b8135815260209182019101613803565b613bec9199505f90613e30565b5f975f6137ca565b6040513d5f823e3d90fd5b63c5f680ed60e01b5f5260045ffd5b5060643515613768565b5060443515613761565b63368f2d7d60e21b5f5260045ffd5b63268dbf6760e21b5f5260045ffd5b63475a253560e01b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b6328ad4a9560e21b5f5260045ffd5b600435906001600160a01b03198216820361123f57565b9181601f8401121561123f578235916001600160401b03831161123f576020838186019501011161123f57565b6024359061ffff8216820361123f57565b6044359061ffff8216820361123f57565b602435906001600160a01b038216820361123f57565b60e060031982011261123f576004356001600160a01b03198116810361123f579160243561ffff8116810361123f579160443591606435916084356001600160401b03811161123f5781613d4491600401613c89565b9290929160a4356001600160401b03811161123f5781613d6691600401613c89565b9290929160c435906001600160401b03821161123f57613d8891600401613c89565b9091565b60a081019081106001600160401b038211176110d157604052565b61010081019081106001600160401b038211176110d157604052565b61016081019081106001600160401b038211176110d157604052565b60c081019081106001600160401b038211176110d157604052565b604081019081106001600160401b038211176110d157604052565b608081019081106001600160401b038211176110d157604052565b90601f801991011681019081106001600160401b038211176110d157604052565b908060209392818452848401375f828201840152601f01601f1916010190565b9290613e8a90613e989593604086526040860191613e51565b926020818503910152613e51565b90565b818102929181159184041417156132fb57565b906001600160401b03809116911603906001600160401b0382116132fb57565b906001600160401b03809116911601906001600160401b0382116132fb57565b9060e08282031261123f5780601f8301121561123f5760405191613f1360e084613e30565b829060e0810192831161123f57905b828210613f2f5750505090565b8135815260209182019101613f22565b51906001600160401b038216820361123f57565b906040821015613f6c57600c600183811c810193160290565b634e487b7160e01b5f52603260045260245ffd5b61ffff5f199116019061ffff82116132fb57565b8054821015613f6c575f5260205f2001905f90565b61ffff1661ffff81146132fb5760010190565b9093929384831161123f57841161123f578101920390565b9291926001600160401b0382116110d15760405191613ffd601f8201601f191660200184613e30565b82948184528183011161123f578281602093845f960137010152565b908160a091031261123f576040519061403182613d8c565b80516001600160a01b038116810361123f57825260208101516020830152604081015160408301526060810151600381101561123f57606083015261407890608001613f3f565b608082015290565b61ffff60019116019061ffff82116132fb57565b9291907f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f940691829060051b8201915b8281106140d25750505050565b909192947f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000183816020938186358b0990089709939291016140c5565b6001600160a01b031916805f5260226020526001600160a01b0360405f2054161561425257805f52602460205260405f208054905f5b8281106141e657505050805f52602460205260405f208054905f8155816141c8575b5050805f52602b6020525f6040812055805f5260226020525f60056040822082815582600182015582600282015582600382015582600482015501557f98a9ec8a25ae28f42f24e68ce0e89786ac50d95191ef5bbd9a4aef2a7eeaef265f80a2565b5f5260205f20908101905b81811015614166575f81556001016141d3565b835f52602360205260405f206001600160a01b03806142058486613f94565b90549060031b1c16165f5260205260405f2060ff198154169055835f52602a60205260405f209060018101918282116132fb5761ffff8060019416165f526020525f604081205501614144565b50565b906020811015613f6c5760051b019056fea2646970667358221220831ce8e936864126785b316590742b3a725d7de5f9e3f0d94cc0145ebd49b83c64736f6c634300081c0033",
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

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactor) SubmitContribution(opts *bind.TransactOpts, roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.contract.Transact(opts, "submitContribution", roundId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerSession) SubmitContribution(roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, roundId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
}

// SubmitContribution is a paid mutator transaction binding the contract method 0xb7bca615.
//
// Solidity: function submitContribution(bytes12 roundId, uint16 contributorIndex, bytes32 commitmentsHash, bytes32 encryptedSharesHash, bytes transcript, bytes proof, bytes input) returns()
func (_DKGManager *DKGManagerTransactorSession) SubmitContribution(roundId [12]byte, contributorIndex uint16, commitmentsHash [32]byte, encryptedSharesHash [32]byte, transcript []byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _DKGManager.Contract.SubmitContribution(&_DKGManager.TransactOpts, roundId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)
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
