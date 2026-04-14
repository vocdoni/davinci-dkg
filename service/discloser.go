package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
)

type PendingRevealedShare struct {
	RoundID          string
	Operator         common.Address
	ParticipantIndex uint16
}

type Discloser struct {
	operator common.Address
	storage  *storage.Storage
}

func NewDiscloser(operator common.Address, st *storage.Storage) *Discloser {
	return &Discloser{
		operator: operator,
		storage:  serviceStorage(st),
	}
}

func (d *Discloser) PendingRevealedShare(roundID string) (*PendingRevealedShare, error) {
	round, err := d.storage.Round(roundID)
	if err != nil {
		return nil, err
	}
	if !round.Policy.DisclosureAllowed || !allowsDisclosure(round.Phase) {
		return nil, nil
	}

	index := participantIndex(round.SelectedParticipants, d.operator)
	if index == 0 {
		return nil, nil
	}
	if hasRevealedShare(d.storage, roundID, d.operator) {
		return nil, nil
	}

	return &PendingRevealedShare{
		RoundID:          roundID,
		Operator:         d.operator,
		ParticipantIndex: index,
	}, nil
}
