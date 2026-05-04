package helpers

import (
	"context"

	"github.com/vocdoni/davinci-dkg/web3"
)

func WaitCombinedDecryption(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	ciphertextIndex uint16,
) (web3.CombinedDecryptionView, error) {
	var record web3.CombinedDecryptionView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		record, fetchErr = services.Contracts.GetCombinedDecryption(ctx, epochID, [32]byte{}, ciphertextIndex)
		return fetchErr == nil && record.Completed
	})
	if err != nil {
		return web3.CombinedDecryptionView{}, err
	}
	return record, nil
}
