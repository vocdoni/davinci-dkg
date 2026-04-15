package web3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// GetNode returns the registered key material for one operator.
func (c *Contracts) GetNode(ctx context.Context, operator common.Address) (RegistryNode, error) {
	input, err := c.registryABI.Pack("getNode", operator)
	if err != nil {
		return RegistryNode{}, fmt.Errorf("pack getNode: %w", err)
	}
	output, err := c.pool.Current().CallContract(ctx, ethereum.CallMsg{
		To:   &c.Addresses.Registry,
		Data: input,
	}, nil)
	if err != nil {
		return RegistryNode{}, fmt.Errorf("call getNode: %w", err)
	}
	values, err := c.registryABI.Unpack("getNode", output)
	if err != nil {
		return RegistryNode{}, fmt.Errorf("unpack getNode: %w", err)
	}
	if len(values) != 4 {
		return RegistryNode{}, fmt.Errorf("unexpected output count for getNode")
	}
	return RegistryNode{
		Operator: values[0].(common.Address),
		PubX:     values[1].(*big.Int),
		PubY:     values[2].(*big.Int),
		Status:   values[3].(uint8),
	}, nil
}
