package helpers

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/web3"
)

func WaitCombinedDecryption(
	ctx context.Context,
	services *TestServices,
	roundID [12]byte,
	ciphertextIndex uint16,
) (web3.CombinedDecryptionView, error) {
	var record web3.CombinedDecryptionView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		record, fetchErr = services.Contracts.GetCombinedDecryption(ctx, roundID, ciphertextIndex)
		return fetchErr == nil && record.Completed
	})
	if err != nil {
		return web3.CombinedDecryptionView{}, err
	}
	return record, nil
}

func WaitRevealedShare(
	ctx context.Context,
	services *TestServices,
	roundID [12]byte,
	participant common.Address,
) (web3.RevealedShareView, error) {
	var record web3.RevealedShareView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		record, fetchErr = services.Contracts.GetRevealedShare(ctx, roundID, participant)
		return fetchErr == nil && record.Accepted
	})
	if err != nil {
		return web3.RevealedShareView{}, err
	}
	return record, nil
}
