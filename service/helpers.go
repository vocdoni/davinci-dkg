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

func hasContribution(st *storage.Storage, roundID string, operator common.Address) bool {
	for _, contribution := range st.Contributions(roundID) {
		if contribution.Contributor == operator {
			return true
		}
	}
	return false
}

func hasPartialDecryption(st *storage.Storage, roundID string, operator common.Address, ciphertextIndex uint16) bool {
	for _, decryption := range st.PartialDecryptions(roundID) {
		if decryption.Participant == operator && decryption.CiphertextIndex == ciphertextIndex {
			return true
		}
	}
	return false
}

func hasRevealedShare(st *storage.Storage, roundID string, operator common.Address) bool {
	for _, share := range st.RevealedShares(roundID) {
		if share.Participant == operator {
			return true
		}
	}
	return false
}

func allowsDecryption(phase types.RoundPhase) bool {
	return phase == types.RoundPhaseFinalized || phase == types.RoundPhaseDecryption
}

func allowsDisclosure(phase types.RoundPhase) bool {
	// RoundPhaseCompleted is the terminal state after reconstructSecret;
	// disclosure is already done at that point.
	return phase == types.RoundPhaseFinalized || phase == types.RoundPhaseDisclosure
}
