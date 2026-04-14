package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

type PendingContribution struct {
	RoundID          string
	Operator         common.Address
	ContributorIndex uint16
}

type Contributor struct {
	operator common.Address
	storage  *storage.Storage
}

func NewContributor(operator common.Address, st *storage.Storage) *Contributor {
	return &Contributor{
		operator: operator,
		storage:  serviceStorage(st),
	}
}

func (c *Contributor) PendingContribution(roundID string) (*PendingContribution, error) {
	round, err := c.storage.Round(roundID)
	if err != nil {
		return nil, err
	}
	if round.Phase != types.RoundPhaseContribution {
		return nil, nil
	}

	index := participantIndex(round.SelectedParticipants, c.operator)
	if index == 0 {
		return nil, nil
	}
	if hasContribution(c.storage, roundID, c.operator) {
		return nil, nil
	}

	return &PendingContribution{
		RoundID:          roundID,
		Operator:         c.operator,
		ContributorIndex: index,
	}, nil
}
