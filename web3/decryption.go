package web3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
)

// GetCombinedDecryption returns the combined decryption record for one
// ciphertext under (epochID, aid). For the legacy per-epoch path pass
// `[32]byte{}` as `aid` (CIRCUITS_AUDIT2 #2).
func (c *Contracts) GetCombinedDecryption(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
) (CombinedDecryptionView, error) {
	input, err := c.managerABI.Pack("getCombinedDecryption", epochID, aid, ciphertextIndex)
	if err != nil {
		return CombinedDecryptionView{}, fmt.Errorf("pack getCombinedDecryption: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Manager,
		Data: input,
	}, nil)
	if err != nil {
		return CombinedDecryptionView{}, fmt.Errorf("call getCombinedDecryption: %w", err)
	}
	values, err := c.managerABI.Unpack("getCombinedDecryption", output)
	if err != nil {
		return CombinedDecryptionView{}, fmt.Errorf("unpack getCombinedDecryption: %w", err)
	}
	if len(values) != 3 {
		return CombinedDecryptionView{}, fmt.Errorf("unexpected output count for getCombinedDecryption: got %d", len(values))
	}
	return CombinedDecryptionView{
		CiphertextIndex: values[0].(uint16),
		Completed:       values[1].(bool),
		Plaintext:       new(big.Int).Set(values[2].(*big.Int)),
	}, nil
}
