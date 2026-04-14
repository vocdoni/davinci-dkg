package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// NodeKey is a node operator identity plus its long-term share-encryption key.
type NodeKey struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
}

// Validate checks that the node key contains the minimal required fields.
func (k NodeKey) Validate() error {
	if k.Operator == (common.Address{}) {
		return fmt.Errorf("operator address is required")
	}
	if k.PubX == nil || k.PubY == nil {
		return fmt.Errorf("public key coordinates are required")
	}
	return nil
}
