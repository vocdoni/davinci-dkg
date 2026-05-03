package storage

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SavePartialDecryption stores one accepted partial decryption keyed by
// (epoch, AID, role, participant, ciphertextIndex). AID and Role are part
// of the key so that a participant may submit distinct shares for different
// applications and producer roles (committee vs. organizer) on the same
// (epoch, ciphertextIndex) without colliding.
func (s *Storage) SavePartialDecryption(decryption types.PartialDecryption) error {
	if err := decryption.Validate(); err != nil {
		return err
	}
	if s.db != nil {
		if _, err := s.Epoch(decryption.EpochID); err != nil {
			return err
		}
		key := partialDecryptionKey(
			decryption.EpochID,
			decryption.AID,
			decryption.Role,
			decryption.Participant,
			decryption.CiphertextIndex,
		)
		if _, err := s.db.Get(key); err == nil {
			return fmt.Errorf("partial decryption already exists")
		} else if err != db.ErrKeyNotFound {
			return err
		}
		payload, err := json.Marshal(decryption)
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
	if _, err := s.Epoch(decryption.EpochID); err != nil {
		return err
	}
	pk := partialMemKey{
		AID:             decryption.AID,
		Role:            decryption.Role,
		Participant:     decryption.Participant,
		CiphertextIndex: decryption.CiphertextIndex,
	}
	if _, ok := s.decryptions[decryption.EpochID]; !ok {
		s.decryptions[decryption.EpochID] = make(map[partialMemKey]types.PartialDecryption)
	}
	if _, ok := s.decryptions[decryption.EpochID][pk]; ok {
		return fmt.Errorf("partial decryption already exists")
	}
	s.decryptions[decryption.EpochID][pk] = decryption
	return nil
}

// PartialDecryptions returns all stored partial decryptions for the epoch.
func (s *Storage) PartialDecryptions(id string) []types.PartialDecryption {
	if s.db != nil {
		result := []types.PartialDecryption{}
		_ = s.db.Iterate(partialDecryptionPrefix(id), func(_, value []byte) bool {
			var decryption types.PartialDecryption
			if err := json.Unmarshal(value, &decryption); err == nil {
				result = append(result, decryption)
			}
			return true
		})
		return result
	}
	result := []types.PartialDecryption{}
	for _, decryption := range s.decryptions[id] {
		result = append(result, decryption)
	}
	return result
}

func partialDecryptionPrefix(id string) []byte {
	return []byte("partial-decryption/" + id + "/")
}

func partialDecryptionKey(
	id string,
	aid [32]byte,
	role types.Role,
	participant common.Address,
	ciphertextIndex uint16,
) []byte {
	key := append(partialDecryptionPrefix(id), aid[:]...)
	key = append(key, byte(role))
	key = append(key, participant.Bytes()...)
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, ciphertextIndex)
	return append(key, buf...)
}
