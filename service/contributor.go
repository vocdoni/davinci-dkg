package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

type PendingContribution struct {
	EpochID          string
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

func (c *Contributor) PendingContribution(epochID string) (*PendingContribution, error) {
	epoch, err := c.storage.Epoch(epochID)
	if err != nil {
		return nil, err
	}
	if epoch.Phase != types.EpochPhaseContribution {
		return nil, nil
	}

	index := participantIndex(epoch.SelectedParticipants, c.operator)
	if index == 0 {
		return nil, nil
	}
	if hasContribution(c.storage, epochID, c.operator) {
		return nil, nil
	}

	return &PendingContribution{
		EpochID:          epochID,
		Operator:         c.operator,
		ContributorIndex: index,
	}, nil
}
