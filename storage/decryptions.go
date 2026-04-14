package storage

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// SavePartialDecryption stores one accepted partial decryption keyed by round, participant, and ciphertext.
func (s *Storage) SavePartialDecryption(decryption types.PartialDecryption) error {
	if err := decryption.Validate(); err != nil {
		return err
	}
	if s.db != nil {
		if _, err := s.Round(decryption.RoundID); err != nil {
			return err
		}
		key := partialDecryptionKey(decryption.RoundID, decryption.Participant, decryption.CiphertextIndex)
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
	if _, err := s.Round(decryption.RoundID); err != nil {
		return err
	}
	if _, ok := s.decryptions[decryption.RoundID][decryption.Participant]; !ok {
		s.decryptions[decryption.RoundID][decryption.Participant] = make(map[uint16]types.PartialDecryption)
	}
	if _, ok := s.decryptions[decryption.RoundID][decryption.Participant][decryption.CiphertextIndex]; ok {
		return fmt.Errorf("partial decryption already exists")
	}
	s.decryptions[decryption.RoundID][decryption.Participant][decryption.CiphertextIndex] = decryption
	return nil
}

// PartialDecryptions returns all stored partial decryptions for the round.
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
	for _, byCiphertext := range s.decryptions[id] {
		for _, decryption := range byCiphertext {
			result = append(result, decryption)
		}
	}
	return result
}

func partialDecryptionPrefix(id string) []byte {
	return []byte("partial-decryption/" + id + "/")
}

func partialDecryptionKey(id string, participant common.Address, ciphertextIndex uint16) []byte {
	key := append(partialDecryptionPrefix(id), participant.Bytes()...)
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, ciphertextIndex)
	return append(key, buf...)
}
