package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
)

type PendingPartialDecryption struct {
	EpochID          string
	Operator         common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
}

type Decryptor struct {
	operator common.Address
	storage  *storage.Storage
}

func NewDecryptor(operator common.Address, st *storage.Storage) *Decryptor {
	return &Decryptor{
		operator: operator,
		storage:  serviceStorage(st),
	}
}

func (d *Decryptor) PendingPartialDecryption(epochID string, ciphertextIndex uint16) (*PendingPartialDecryption, error) {
	if ciphertextIndex == 0 {
		return nil, fmt.Errorf("ciphertext index is required")
	}

	epoch, err := d.storage.Epoch(epochID)
	if err != nil {
		return nil, err
	}
	if !allowsDecryption(epoch.Phase) {
		return nil, nil
	}

	index := participantIndex(epoch.SelectedParticipants, d.operator)
	if index == 0 {
		return nil, nil
	}
	if hasPartialDecryption(d.storage, epochID, d.operator, ciphertextIndex) {
		return nil, nil
	}

	return &PendingPartialDecryption{
		EpochID:          epochID,
		Operator:         d.operator,
		ParticipantIndex: index,
		CiphertextIndex:  ciphertextIndex,
	}, nil
}
