package decryptcombine

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a decrypt-combine
// witness.
//
// The organizer fields are the share the application's organizer posted on
// chain (`OrganizerShareSubmitted`) plus the PK_org from the application
// record. The builder recomputes the DLEQ challenge `e` itself with
// dleq.OrganizerShareChallenge, so callers never pass it: a caller-supplied
// `e` could silently disagree with what the contract recomputes.
type Assignment struct {
	RoundHash          *big.Int // semantically: eid, ≤ 12 bytes
	Aid                *big.Int // application identifier, < BN254 scalar field
	CtIdx              *big.Int // per-application ciphertext index, uint16
	DeltaOrg           types.CurvePoint
	OrganizerPK        types.CurvePoint
	OrganizerProof     dleq.Proof // A1, A2 and z of the organizer's DLEQ
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
	if len(a.ParticipantIndexes) != int(a.Threshold) {
		return fmt.Errorf("share count %d does not match threshold %d", len(a.ParticipantIndexes), a.Threshold)
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

// validateOrganizer mirrors the in-circuit organizer checks: the four points
// must be well-formed and z canonical. The DLEQ itself is verified by the
// circuit (and, before the node ever builds a witness, off chain by
// dleq.VerifyOrganizerShare).
func (a Assignment) validateOrganizer() error {
	for _, point := range []struct {
		kind  string
		value types.CurvePoint
	}{
		{"organizer share", a.DeltaOrg},
		{"organizer public key", a.OrganizerPK},
		{"organizer A1", a.OrganizerProof.A1},
		{"organizer A2", a.OrganizerProof.A2},
	} {
		if err := point.value.Validate(); err != nil {
			return fmt.Errorf("%s: %w", point.kind, err)
		}
	}
	z := a.OrganizerProof.Response
	if z == nil {
		return fmt.Errorf("organizer response is required")
	}
	if z.Sign() < 0 || z.Cmp(ccommon.SubgroupOrderMinusOne()) > 0 {
		return fmt.Errorf("organizer response is not canonical: must be in [0, r_bjj-1]")
	}
	return nil
}
