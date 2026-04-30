package types

import "github.com/ethereum/go-ethereum/common"

// Epoch is the local typed representation of one DKG epoch.
type Epoch struct {
	ID                   string
	Organizer            common.Address
	Policy               EpochPolicy
	Phase                EpochPhase
	SelectedParticipants []common.Address
}
