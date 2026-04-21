package web3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// GetCombinedDecryption returns the combined decryption record for one ciphertext.
func (c *Contracts) GetCombinedDecryption(
	ctx context.Context,
	roundID [12]byte,
	ciphertextIndex uint16,
) (CombinedDecryptionView, error) {
	input, err := c.managerABI.Pack("getCombinedDecryption", roundID, ciphertextIndex)
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

// GetRevealedShare returns the revealed-share record for one participant.
func (c *Contracts) GetRevealedShare(
	ctx context.Context,
	roundID [12]byte,
	participant common.Address,
) (RevealedShareView, error) {
	input, err := c.managerABI.Pack("getRevealedShare", roundID, participant)
	if err != nil {
		return RevealedShareView{}, fmt.Errorf("pack getRevealedShare: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Manager,
		Data: input,
	}, nil)
	if err != nil {
		return RevealedShareView{}, fmt.Errorf("call getRevealedShare: %w", err)
	}
	values, err := c.managerABI.Unpack("getRevealedShare", output)
	if err != nil {
		return RevealedShareView{}, fmt.Errorf("unpack getRevealedShare: %w", err)
	}
	if len(values) != 5 {
		return RevealedShareView{}, fmt.Errorf("unexpected output count for getRevealedShare")
	}
	shareHash := values[3].([32]byte)
	return RevealedShareView{
		Participant:      values[0].(common.Address),
		ParticipantIndex: values[1].(uint16),
		ShareValue:       values[2].(*big.Int),
		ShareHash:        common.BytesToHash(shareHash[:]),
		Accepted:         values[4].(bool),
	}, nil
}
