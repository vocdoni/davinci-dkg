package config

import (
	"math/big"
	"time"
)

// Chain is what the node tunes per chain. The zero value means "unknown
// chain": the generic defaults apply.
type Chain struct {
	// BlockTime is the chain's block interval; 0 when unknown.
	BlockTime time.Duration
	// FallbackTipCap is the priority fee used when eth_maxPriorityFeePerGas
	// fails; nil means 1 gwei.
	FallbackTipCap *big.Int
	// MaxFeePerGas caps maxFeePerGas and the tip, fee bumps included; nil
	// means uncapped.
	MaxFeePerGas *big.Int
}

// ChainByID returns the profile of a chain, keyed by chain id so a custom
// --manager deployment on a known chain gets it too. The big.Ints are
// copies; callers may keep and mutate them.
func ChainByID(chainID uint64) Chain {
	switch chainID {
	case 10: // OP Mainnet: tip is ~0.001 gwei and the base fee a few thousand wei,
		// so the generic 1 gwei fallback would be ~1000x too high.
		return Chain{
			BlockTime:      2 * time.Second,
			FallbackTipCap: big.NewInt(1_000_000), // 0.001 gwei
			// 0.02 gwei: 5x the worst OP base fee in 4,145 samples (0.0039 gwei,
			// median 5.5e-7), and what an operator must hold per in-flight tx.
			MaxFeePerGas: big.NewInt(20_000_000),
		}
	case 100: // Gnosis Chain
		return Chain{BlockTime: 5 * time.Second}
	case 1, 11155111: // Ethereum mainnet, Sepolia
		return Chain{BlockTime: 12 * time.Second}
	}
	return Chain{}
}
