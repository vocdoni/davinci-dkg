package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// CurvePoint is an explicit affine point representation used across native and contract payloads.
type CurvePoint struct {
	X *big.Int
	Y *big.Int
}

// Validate checks that both coordinates are present.
func (p CurvePoint) Validate() error {
	if p.X == nil || p.Y == nil {
		return fmt.Errorf("point coordinates are required")
	}
	return nil
}

// EncryptedShare is one per-recipient ciphertext emitted by a contributor.
type EncryptedShare struct {
	Recipient      common.Address
	RecipientIndex uint16
	Ephemeral      CurvePoint
	Ciphertext     *big.Int
}

// Validate checks that the encrypted share is minimally well-formed.
func (s EncryptedShare) Validate() error {
	if s.Recipient == (common.Address{}) {
		return fmt.Errorf("recipient is required")
	}
	if s.RecipientIndex == 0 {
		return fmt.Errorf("recipient index is required")
	}
	if err := s.Ephemeral.Validate(); err != nil {
		return fmt.Errorf("ephemeral point: %w", err)
	}
	if s.Ciphertext == nil {
		return fmt.Errorf("ciphertext is required")
	}
	return nil
}

// Contribution is the off-chain typed representation of one accepted DKG contribution.
type Contribution struct {
	RoundID          string
	Contributor      common.Address
	ContributorIndex uint16
	Commitments      []CurvePoint
	EncryptedShares  []EncryptedShare
	Proof            []byte
	PublicInputHash  common.Hash
}

// Validate checks that the contribution contains the minimum data needed by storage and services.
func (c Contribution) Validate() error {
	if c.RoundID == "" {
		return fmt.Errorf("round id is required")
	}
	if c.Contributor == (common.Address{}) {
		return fmt.Errorf("contributor is required")
	}
	if c.ContributorIndex == 0 {
		return fmt.Errorf("contributor index is required")
	}
	if len(c.Commitments) == 0 {
		return fmt.Errorf("at least one commitment is required")
	}
	for i, point := range c.Commitments {
		if err := point.Validate(); err != nil {
			return fmt.Errorf("commitment %d: %w", i, err)
		}
	}
	if len(c.EncryptedShares) == 0 {
		return fmt.Errorf("at least one encrypted share is required")
	}
	for i, share := range c.EncryptedShares {
		if err := share.Validate(); err != nil {
			return fmt.Errorf("encrypted share %d: %w", i, err)
		}
	}
	if len(c.Proof) == 0 {
		return fmt.Errorf("proof is required")
	}
	return nil
}

// FinalizedOutput is the typed result of the contribution aggregation/finalization phase.
type FinalizedOutput struct {
	RoundID               string
	CollectivePublicKey   CurvePoint
	AggregateCommitments  []CurvePoint
	SelectedParticipantIX []uint16
}

// Validate checks that the finalized output is minimally coherent.
func (o FinalizedOutput) Validate() error {
	if o.RoundID == "" {
		return fmt.Errorf("round id is required")
	}
	if err := o.CollectivePublicKey.Validate(); err != nil {
		return fmt.Errorf("collective public key: %w", err)
	}
	if len(o.AggregateCommitments) == 0 {
		return fmt.Errorf("aggregate commitments are required")
	}
	for i, point := range o.AggregateCommitments {
		if err := point.Validate(); err != nil {
			return fmt.Errorf("aggregate commitment %d: %w", i, err)
		}
	}
	if len(o.SelectedParticipantIX) == 0 {
		return fmt.Errorf("selected participant indices are required")
	}
	return nil
}
