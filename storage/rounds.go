package storage

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SaveRound stores a new round.
func (s *Storage) SaveRound(round types.Round) error {
	if round.ID == "" {
		return fmt.Errorf("round id is required")
	}
	if s.db != nil {
		if _, err := s.Round(round.ID); err == nil {
			return fmt.Errorf("round already exists")
		}
		payload, err := json.Marshal(round)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(roundKey(round.ID), payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, exists := s.rounds[round.ID]; exists {
		return fmt.Errorf("round already exists")
	}
	s.rounds[round.ID] = round
	s.ready[round.ID] = make(map[common.Address]struct{})
	s.contributions[round.ID] = make(map[common.Address]types.Contribution)
	s.decryptions[round.ID] = make(map[common.Address]map[uint16]types.PartialDecryption)
	s.disclosures[round.ID] = make(map[common.Address]types.RevealedShare)
	return nil
}

// UpsertRound stores or replaces a round snapshot.
func (s *Storage) UpsertRound(round types.Round) error {
	if round.ID == "" {
		return fmt.Errorf("round id is required")
	}
	if s.db != nil {
		payload, err := json.Marshal(round)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(roundKey(round.ID), payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, exists := s.rounds[round.ID]; !exists {
		s.ready[round.ID] = make(map[common.Address]struct{})
		s.contributions[round.ID] = make(map[common.Address]types.Contribution)
		s.decryptions[round.ID] = make(map[common.Address]map[uint16]types.PartialDecryption)
		s.disclosures[round.ID] = make(map[common.Address]types.RevealedShare)
	}
	s.rounds[round.ID] = round
	return nil
}

// Round returns the stored round.
func (s *Storage) Round(id string) (types.Round, error) {
	if s.db != nil {
		payload, err := s.db.Get(roundKey(id))
		if err != nil {
			if err == db.ErrKeyNotFound {
				return types.Round{}, fmt.Errorf("round not found")
			}
			return types.Round{}, err
		}
		var round types.Round
		if err := json.Unmarshal(payload, &round); err != nil {
			return types.Round{}, err
		}
		return round, nil
	}
	round, ok := s.rounds[id]
	if !ok {
		return types.Round{}, fmt.Errorf("round not found")
	}
	return round, nil
}

// MarkReady marks a participant as ready for the round.
func (s *Storage) MarkReady(id string, operator common.Address) error {
	if s.db != nil {
		if _, err := s.Round(id); err != nil {
			return err
		}
		key := readyKey(id, operator)
		if _, err := s.db.Get(key); err == nil {
			return fmt.Errorf("operator already marked ready")
		} else if err != db.ErrKeyNotFound {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(key, []byte{1}); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, ok := s.rounds[id]; !ok {
		return fmt.Errorf("round not found")
	}
	if _, ok := s.ready[id][operator]; ok {
		return fmt.Errorf("operator already marked ready")
	}
	s.ready[id][operator] = struct{}{}
	return nil
}

// ReadyCount returns the number of ready participants for the round.
func (s *Storage) ReadyCount(id string) int {
	if s.db != nil {
		count := 0
		_ = s.db.Iterate(readyPrefix(id), func(_, _ []byte) bool {
			count++
			return true
		})
		return count
	}
	return len(s.ready[id])
}

// SetSelectedParticipants stores the ordered committee and advances the round phase.
func (s *Storage) SetSelectedParticipants(id string, participants []common.Address) error {
	if s.db != nil {
		stored, err := s.Round(id)
		if err != nil {
			return err
		}
		stored.SelectedParticipants = append([]common.Address(nil), participants...)
		stored.Phase = types.RoundPhaseContribution
		payload, err := json.Marshal(stored)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(roundKey(id), payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	round, err := s.Round(id)
	if err != nil {
		return err
	}
	round.SelectedParticipants = append([]common.Address(nil), participants...)
	round.Phase = types.RoundPhaseContribution
	s.rounds[id] = round
	return nil
}

func roundKey(id string) []byte {
	return []byte("round/" + id)
}

func readyPrefix(id string) []byte {
	return []byte("ready/" + id + "/")
}

func readyKey(id string, operator common.Address) []byte {
	return append(readyPrefix(id), operator.Bytes()...)
}
