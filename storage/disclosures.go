package storage

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SaveRevealedShare stores one accepted emergency reveal-share payload.
func (s *Storage) SaveRevealedShare(share types.RevealedShare) error {
	if err := share.Validate(); err != nil {
		return err
	}
	if s.db != nil {
		if _, err := s.Round(share.RoundID); err != nil {
			return err
		}
		key := revealedShareKey(share.RoundID, share.Participant)
		if _, err := s.db.Get(key); err == nil {
			return fmt.Errorf("revealed share already exists")
		} else if err != db.ErrKeyNotFound {
			return err
		}
		payload, err := json.Marshal(share)
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
	if _, err := s.Round(share.RoundID); err != nil {
		return err
	}
	if _, ok := s.disclosures[share.RoundID][share.Participant]; ok {
		return fmt.Errorf("revealed share already exists")
	}
	s.disclosures[share.RoundID][share.Participant] = share
	return nil
}

// RevealedShares returns all stored revealed shares for the round.
func (s *Storage) RevealedShares(id string) []types.RevealedShare {
	if s.db != nil {
		result := []types.RevealedShare{}
		_ = s.db.Iterate(revealedSharePrefix(id), func(_, value []byte) bool {
			var share types.RevealedShare
			if err := json.Unmarshal(value, &share); err == nil {
				result = append(result, share)
			}
			return true
		})
		return result
	}
	result := make([]types.RevealedShare, 0, len(s.disclosures[id]))
	for _, share := range s.disclosures[id] {
		result = append(result, share)
	}
	return result
}

func revealedSharePrefix(id string) []byte {
	return []byte("revealed-share/" + id + "/")
}

func revealedShareKey(id string, participant common.Address) []byte {
	return append(revealedSharePrefix(id), participant.Bytes()...)
}
