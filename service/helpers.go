package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/storage"
	"github.com/vocdoni/davinci-dkg/types"
)

func serviceStorage(st *storage.Storage) *storage.Storage {
	if st == nil {
		return storage.New()
	}
	return st
}

func participantIndex(participants []common.Address, operator common.Address) uint16 {
	for i, participant := range participants {
		if participant == operator {
			return uint16(i + 1)
		}
	}
	return 0
}

func hasContribution(st *storage.Storage, epochID string, operator common.Address) bool {
	for _, contribution := range st.Contributions(epochID) {
		if contribution.Contributor == operator {
			return true
		}
	}
	return false
}

func hasPartialDecryption(st *storage.Storage, epochID string, operator common.Address, ciphertextIndex uint16) bool {
	for _, decryption := range st.PartialDecryptions(epochID) {
		if decryption.Participant == operator && decryption.CiphertextIndex == ciphertextIndex {
			return true
		}
	}
	return false
}

func allowsDecryption(phase types.EpochPhase) bool {
	return phase == types.EpochPhaseFinalized || phase == types.EpochPhaseDecryption
}
