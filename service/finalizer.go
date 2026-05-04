package service

import (
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

type PendingFinalize struct {
	EpochID              string
	ContributionCount    int
	RequiredContributors uint16
}

type Finalizer struct {
	storage *storage.Storage
}

func NewFinalizer(st *storage.Storage) *Finalizer {
	return &Finalizer{storage: serviceStorage(st)}
}

func (f *Finalizer) PendingFinalize(epochID string) (*PendingFinalize, error) {
	epoch, err := f.storage.Epoch(epochID)
	if err != nil {
		return nil, err
	}
	if epoch.Phase != types.EpochPhaseKeyAssembly {
		return nil, nil
	}

	contributions := f.storage.Contributions(epochID)
	if len(contributions) < int(epoch.Policy.MinValidContributions) {
		return nil, nil
	}

	return &PendingFinalize{
		EpochID:              epochID,
		ContributionCount:    len(contributions),
		RequiredContributors: epoch.Policy.MinValidContributions,
	}, nil
}
