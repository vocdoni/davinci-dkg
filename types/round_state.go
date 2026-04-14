package types

import "github.com/ethereum/go-ethereum/common"

// Round is the local typed representation of one DKG round.
type Round struct {
	ID                   string
	Organizer            common.Address
	Policy               RoundPolicy
	Phase                RoundPhase
	SelectedParticipants []common.Address
}
