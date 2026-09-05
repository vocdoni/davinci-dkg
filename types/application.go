package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AppMode mirrors `DKGTypes.AppMode`: who is needed to decrypt.
type AppMode uint8

const (
	// AppModeOrganizerLocked: the organizer keeps sk_org until it calls
	// revealOrganizerSecret; the committee alone cannot decrypt before that.
	AppModeOrganizerLocked AppMode = iota
	// AppModeAutomatic: no organizer key at all (PK_org is the identity),
	// so the committee decrypts on its own as soon as it holds t partials.
	AppModeAutomatic
)

// AppPolicy is the per-application policy fixed at registration. Mirrors
// `DKGTypes.AppPolicy`. Submission is open to anyone (OpenSubmission), to
// the Submitters allow-list, or — when both are empty — to the registrant.
type AppPolicy struct {
	Mode             AppMode
	OpenSubmission   bool
	Submitters       []common.Address // at most 32
	MaxCiphertexts   uint16           // 0 = unlimited
	NotBeforeBlock   uint64
	NotAfterBlock    uint64
	DecryptNotBefore uint64 // unix seconds; partials and combines revert before it (0 = none)
	DecryptNotAfter  uint64 // unix seconds; partials and combines revert after it (0 = never)
}

// Application is the cached on-chain record. Stored locally so workers can
// route ciphertext events to the right per-application worker without an
// extra getApplication round-trip per ciphertext.
//
// Every application claims one of the epoch's MaxK pool keys at registration
// (PoolIndex): the encryption key is PK_aid = P_j for an Automatic
// application and P_j + PK_org for an organizer-locked one, so a ciphertext
// copied into another application decrypts under a different key. The
// organizer of a locked application reveals sk_org once, after which the
// committee combines on its own.
type Application struct {
	EpochID         string
	AID             [32]byte
	Creator         common.Address
	OrganizerPK     CurvePoint
	OrganizerSecret *big.Int // sk_org once revealed; nil or zero before that
	PoolIndex       uint8
	Policy          AppPolicy
	CreatedAtBlock  uint64
	Exists          bool
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
	// An organizer-locked application must carry a real key; the identity
	// (0, 1) is what an automatic registration stores.
	if a.Policy.Mode == AppModeOrganizerLocked &&
		a.OrganizerPK.X.Sign() == 0 && a.OrganizerPK.Y.Cmp(big.NewInt(1)) == 0 {
		return fmt.Errorf("organizer-locked application without an organizer key")
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
