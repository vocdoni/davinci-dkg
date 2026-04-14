package storage

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SaveContribution stores one accepted contribution keyed by round and contributor.
func (s *Storage) SaveContribution(contribution types.Contribution) error {
	if err := contribution.Validate(); err != nil {
		return err
	}
	if s.db != nil {
		if _, err := s.Round(contribution.RoundID); err != nil {
			return err
		}
		key := contributionKey(contribution.RoundID, contribution.Contributor)
		if _, err := s.db.Get(key); err == nil {
			return fmt.Errorf("contribution already exists")
		} else if err != db.ErrKeyNotFound {
			return err
		}
		payload, err := json.Marshal(contribution)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(key, payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, err := s.Round(contribution.RoundID); err != nil {
		return err
	}
	if _, ok := s.contributions[contribution.RoundID][contribution.Contributor]; ok {
		return fmt.Errorf("contribution already exists")
	}
	s.contributions[contribution.RoundID][contribution.Contributor] = contribution
	return nil
}

// Contribution returns one stored contribution.
func (s *Storage) Contribution(id string, contributor common.Address) (types.Contribution, error) {
	if s.db != nil {
		payload, err := s.db.Get(contributionKey(id, contributor))
		if err != nil {
			if err == db.ErrKeyNotFound {
				return types.Contribution{}, fmt.Errorf("contribution not found")
			}
			return types.Contribution{}, err
		}
		var contribution types.Contribution
		if err := json.Unmarshal(payload, &contribution); err != nil {
			return types.Contribution{}, err
		}
		return contribution, nil
	}
	contribution, ok := s.contributions[id][contributor]
	if !ok {
		return types.Contribution{}, fmt.Errorf("contribution not found")
	}
	return contribution, nil
}

// Contributions returns all stored contributions for the round.
func (s *Storage) Contributions(id string) []types.Contribution {
	if s.db != nil {
		result := []types.Contribution{}
		_ = s.db.Iterate(contributionPrefix(id), func(_, value []byte) bool {
			var contribution types.Contribution
			if err := json.Unmarshal(value, &contribution); err == nil {
				result = append(result, contribution)
			}
			return true
		})
		return result
	}
	result := make([]types.Contribution, 0, len(s.contributions[id]))
	for _, contribution := range s.contributions[id] {
		result = append(result, contribution)
	}
	return result
}

func contributionPrefix(id string) []byte {
	return []byte("contribution/" + id + "/")
}

func contributionKey(id string, contributor common.Address) []byte {
	return append(contributionPrefix(id), contributor.Bytes()...)
}
