package helpers

import (
	"context"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// WaitCombinedDecryption polls until the combined decryption record for
// (epochID, aid, ciphertextIndex) is on chain.
func WaitCombinedDecryption(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
) (web3.CombinedDecryptionView, error) {
	var record web3.CombinedDecryptionView
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		var fetchErr error
		record, fetchErr = services.Contracts.GetCombinedDecryption(ctx, epochID, aid, ciphertextIndex)
		return fetchErr == nil && record.Completed
	})
	if err != nil {
		return web3.CombinedDecryptionView{}, err
	}
	return record, nil
}

// ScalarBasePoint returns s·G. Used by tests to derive PK_org = sk_org·G.
func ScalarBasePoint(s *big.Int) types.CurvePoint {
	p := group.NewPoint()
	p.ScalarBaseMult(s)
	return group.Encode(p)
}
