package decryptcombine

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a decrypt-combine
// witness.
//
// OrganizerPK is the key from the on-chain application record and
// OrganizerSecret its discrete log: the revealed `sk_org` of an
// organizer-locked application, or 0 paired with the identity key (0, 1)
// for an automatic one. The circuit — not this struct — enforces
// PK_org = sk_org·G, so a mismatch fails at proving time.
type Assignment struct {
	RoundHash          *big.Int // semantically: eid, ≤ 12 bytes
	Aid                *big.Int // application identifier, < BN254 scalar field
	CtIdx              *big.Int // per-application ciphertext index, uint16
	OrganizerPK        types.CurvePoint
	OrganizerSecret    *big.Int
	Threshold          uint16
	CiphertextC1       types.CurvePoint
	CiphertextC2       types.CurvePoint
	ParticipantIndexes []uint16
	PartialDecryptions []types.CurvePoint
	Plaintext          *big.Int
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("epoch hash is required")
	}
	if a.RoundHash.Sign() < 0 || a.RoundHash.BitLen() > 96 {
		return fmt.Errorf("epoch hash must fit in 12 bytes")
	}
	if a.Aid != nil && (a.Aid.Sign() < 0 || a.Aid.BitLen() > 256) {
		return fmt.Errorf("application id must fit in 32 bytes")
	}
	if a.CtIdx != nil && (a.CtIdx.Sign() < 0 || a.CtIdx.BitLen() > 16) {
		return fmt.Errorf("ciphertext index must fit in a uint16")
	}
	if a.Threshold == 0 {
		return fmt.Errorf("threshold is required")
	}
	if err := a.CiphertextC1.Validate(); err != nil {
		return fmt.Errorf("ciphertext C1: %w", err)
	}
	if err := a.CiphertextC2.Validate(); err != nil {
		return fmt.Errorf("ciphertext C2: %w", err)
	}
	if err := a.validateOrganizer(); err != nil {
		return err
	}
	if len(a.ParticipantIndexes) == 0 || len(a.ParticipantIndexes) != len(a.PartialDecryptions) {
		return fmt.Errorf("participant indexes and partial decryptions must have the same non-zero length")
	}
	if len(a.ParticipantIndexes) < int(a.Threshold) {
		return fmt.Errorf("share count %d is below the threshold %d", len(a.ParticipantIndexes), a.Threshold)
	}
	if len(a.ParticipantIndexes) > MaxShares {
		return fmt.Errorf("share count %d exceeds max %d", len(a.ParticipantIndexes), MaxShares)
	}
	for i, index := range a.ParticipantIndexes {
		if index == 0 {
			return fmt.Errorf("participant index %d is zero", i)
		}
	}
	for i, value := range a.PartialDecryptions {
		if err := value.Validate(); err != nil {
			return fmt.Errorf("partial decryption %d: %w", i, err)
		}
	}
	if a.Plaintext == nil {
		return fmt.Errorf("plaintext is required")
	}
	// Mirror the in-circuit canonical-range bound so the
	// prover fails fast at witness build instead of inside the SNARK.
	if a.Plaintext.Sign() < 0 || a.Plaintext.Cmp(ccommon.SubgroupOrderMinusOne()) > 0 {
		return fmt.Errorf("plaintext is not canonical: must be in [0, r_bjj-1]")
	}
	return nil
}

// validateOrganizer mirrors the structural half of the in-circuit organizer
// checks: PK_org must be well-formed and sk_org canonical. The algebraic
// relation PK_org = sk_org·G is left to the circuit, which is the only
// place it has to hold.
func (a Assignment) validateOrganizer() error {
	if err := a.OrganizerPK.Validate(); err != nil {
		return fmt.Errorf("organizer public key: %w", err)
	}
	if a.OrganizerSecret == nil {
		return fmt.Errorf("organizer secret is required")
	}
	if a.OrganizerSecret.Sign() < 0 || a.OrganizerSecret.Cmp(ccommon.SubgroupOrderMinusOne()) > 0 {
		return fmt.Errorf("organizer secret is not canonical: must be in [0, r_bjj-1]")
	}
	return nil
}
