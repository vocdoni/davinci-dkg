package storage

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SaveEpoch stores a new epoch.
func (s *Storage) SaveEpoch(epoch types.Epoch) error {
	if epoch.ID == "" {
		return fmt.Errorf("epoch id is required")
	}
	if s.db != nil {
		if _, err := s.Epoch(epoch.ID); err == nil {
			return fmt.Errorf("epoch already exists")
		}
		payload, err := json.Marshal(epoch)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(roundKey(epoch.ID), payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, exists := s.epochs[epoch.ID]; exists {
		return fmt.Errorf("epoch already exists")
	}
	s.epochs[epoch.ID] = epoch
	s.ready[epoch.ID] = make(map[common.Address]struct{})
	s.contributions[epoch.ID] = make(map[common.Address]types.Contribution)
	s.decryptions[epoch.ID] = make(map[partialMemKey]types.PartialDecryption)
	return nil
}

// UpsertEpoch stores or replaces a epoch snapshot.
func (s *Storage) UpsertEpoch(epoch types.Epoch) error {
	if epoch.ID == "" {
		return fmt.Errorf("epoch id is required")
	}
	if s.db != nil {
		payload, err := json.Marshal(epoch)
		if err != nil {
			return err
		}
		tx := s.db.WriteTx()
		defer tx.Discard()
		if err := tx.Set(roundKey(epoch.ID), payload); err != nil {
			return err
		}
		return tx.Commit()
	}
	if _, exists := s.epochs[epoch.ID]; !exists {
		s.ready[epoch.ID] = make(map[common.Address]struct{})
		s.contributions[epoch.ID] = make(map[common.Address]types.Contribution)
		s.decryptions[epoch.ID] = make(map[partialMemKey]types.PartialDecryption)
	}
	s.epochs[epoch.ID] = epoch
	return nil
}

// Epoch returns the stored epoch.
func (s *Storage) Epoch(id string) (types.Epoch, error) {
	if s.db != nil {
		payload, err := s.db.Get(roundKey(id))
		if err != nil {
			if err == db.ErrKeyNotFound {
				return types.Epoch{}, fmt.Errorf("epoch not found")
			}
			return types.Epoch{}, err
		}
		var epoch types.Epoch
		if err := json.Unmarshal(payload, &epoch); err != nil {
			return types.Epoch{}, err
		}
		return epoch, nil
	}
	epoch, ok := s.epochs[id]
	if !ok {
		return types.Epoch{}, fmt.Errorf("epoch not found")
	}
	return epoch, nil
}

// MarkReady marks a participant as ready for the epoch.
func (s *Storage) MarkReady(id string, operator common.Address) error {
	if s.db != nil {
		if _, err := s.Epoch(id); err != nil {
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
	if _, ok := s.epochs[id]; !ok {
		return fmt.Errorf("epoch not found")
	}
	if _, ok := s.ready[id][operator]; ok {
		return fmt.Errorf("operator already marked ready")
	}
	s.ready[id][operator] = struct{}{}
	return nil
}

// ReadyCount returns the number of ready participants for the epoch.
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

// SetSelectedParticipants stores the ordered committee and advances the epoch phase.
func (s *Storage) SetSelectedParticipants(id string, participants []common.Address) error {
	if s.db != nil {
		stored, err := s.Epoch(id)
		if err != nil {
			return err
		}
		stored.SelectedParticipants = append([]common.Address(nil), participants...)
		stored.Phase = types.EpochPhaseContribution
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
	epoch, err := s.Epoch(id)
	if err != nil {
		return err
	}
	epoch.SelectedParticipants = append([]common.Address(nil), participants...)
	epoch.Phase = types.EpochPhaseContribution
	s.epochs[id] = epoch
	return nil
}

func roundKey(id string) []byte {
	return []byte("epoch/" + id)
}

func readyPrefix(id string) []byte {
	return []byte("ready/" + id + "/")
}

func readyKey(id string, operator common.Address) []byte {
	return append(readyPrefix(id), operator.Bytes()...)
}
