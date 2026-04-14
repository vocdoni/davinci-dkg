package service

import (
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

type PendingFinalize struct {
	RoundID              string
	ContributionCount    int
	RequiredContributors uint16
}

type Finalizer struct {
	storage *storage.Storage
}

func NewFinalizer(st *storage.Storage) *Finalizer {
	return &Finalizer{storage: serviceStorage(st)}
}

func (f *Finalizer) PendingFinalize(roundID string) (*PendingFinalize, error) {
	round, err := f.storage.Round(roundID)
	if err != nil {
		return nil, err
	}
	if round.Phase != types.RoundPhaseContribution {
		return nil, nil
	}

	contributions := f.storage.Contributions(roundID)
	if len(contributions) < int(round.Policy.MinValidContributions) {
		return nil, nil
	}

	return &PendingFinalize{
		RoundID:              roundID,
		ContributionCount:    len(contributions),
		RequiredContributors: round.Policy.MinValidContributions,
	}, nil
}
