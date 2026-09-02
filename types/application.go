package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AppPolicy gates submitCiphertext for one application. Mirrors
// `DKGTypes.AppPolicy`. A zero AuthorizedSubmitter is resolved to the
// registering address on chain, so the stored value is never zero.
type AppPolicy struct {
	AuthorizedSubmitter common.Address
	MaxCiphertexts      uint16 // 0 = unlimited
	NotBeforeBlock      uint64
	NotAfterBlock       uint64
}

// Application is the cached on-chain record. Stored locally so workers can
// route ciphertext events to the right per-application worker without an
// extra getApplication round-trip per ciphertext.
//
// Every application carries an organizer key: the application encryption key
// is PK_aid = PK_ep + PK_org, so decryption needs both the committee's
// threshold and the organizer's Δ = sk_org·C1.
type Application struct {
	EpochID        string
	AID            [32]byte
	Creator        common.Address
	OrganizerPK    CurvePoint
	Policy         AppPolicy
	CreatedAtBlock uint64
	Exists         bool
}

// Validate checks that the application record is internally coherent.
func (a Application) Validate() error {
	if a.EpochID == "" {
		return fmt.Errorf("epoch id is required")
	}
	if a.AID == ([32]byte{}) {
		return fmt.Errorf("application id is required")
	}
	if err := a.OrganizerPK.Validate(); err != nil {
		return fmt.Errorf("organizer public key: %w", err)
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

// OrganizerShare is the organizer's Δ = sk_org·C1 for one ciphertext, with
// the Chaum-Pedersen DLEQ (A1, A2, z) that binds it to the registered
// PK_org. The contract stores only the keccak of these words; the combine
// SNARK verifies the DLEQ.
type OrganizerShare struct {
	EpochID         string
	AID             [32]byte
	CiphertextIndex uint16
	DeltaOrg        CurvePoint
	A1              CurvePoint
	A2              CurvePoint
	Response        *big.Int
}
