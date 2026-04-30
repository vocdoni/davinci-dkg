package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AppMode mirrors `solidity/src/libraries/DKGTypes.sol::AppMode` and
// `internal/protocol/protocol.go::AppMode`. See paper §4.3 / §6.
type AppMode uint8

const (
	AppModePublicDerivation AppMode = 0
	AppModeOrganizerCoDec   AppMode = 1
)

// AppPolicy gates submitCiphertext for one application. Mirrors
// `DKGTypes.AppPolicy`.
type AppPolicy struct {
	AuthorizedSubmitter common.Address // 0 = open
	MaxCiphertexts      uint16         // 0 = unlimited
	NotBeforeBlock      uint64
	NotAfterBlock       uint64
}

// Application is the cached on-chain record. Stored locally so workers can
// route ciphertext events to the right per-application worker without an
// extra getApplication round-trip per ciphertext.
type Application struct {
	EpochID        string
	AID            [32]byte
	Creator        common.Address
	Mode           AppMode
	DerivationS    *big.Int   // mode 0 only; nil/zero in mode 1
	OrganizerPK    CurvePoint // mode 1 only; identity (0,1) in mode 0
	Policy         AppPolicy
	CreatedAtBlock uint64
}

// Validate checks that the application record is internally coherent.
func (a Application) Validate() error {
	if a.EpochID == "" {
		return fmt.Errorf("epoch id is required")
	}
	if a.AID == ([32]byte{}) {
		return fmt.Errorf("application id is required")
	}
	switch a.Mode {
	case AppModePublicDerivation:
		if a.DerivationS == nil || a.DerivationS.Sign() == 0 {
			return fmt.Errorf("public-derivation mode requires non-zero S")
		}
	case AppModeOrganizerCoDec:
		if err := a.OrganizerPK.Validate(); err != nil {
			return fmt.Errorf("organizer public key: %w", err)
		}
	default:
		return fmt.Errorf("unknown app mode %d", a.Mode)
	}
	return nil
}

// CiphertextRecord is one (epoch, app, ctIdx) ciphertext.
type CiphertextRecord struct {
	EpochID         string
	AID             [32]byte
	CiphertextIndex uint16
	C1              CurvePoint
	C2              CurvePoint
	Submitter       common.Address
	SubmittedAt     uint64
}

// OrganizerShare is the organizer's Δ_org submission for a mode-1 ciphertext.
type OrganizerShare struct {
	EpochID         string
	AID             [32]byte
	CiphertextIndex uint16
	DeltaOrg        CurvePoint
	DLEQProof       []byte
}
