package config

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// DefaultNetwork is the preset a node runs on when neither --network nor
// --manager is given.
const DefaultNetwork = "optimism"

// NetworkDeployment is the canonical on-chain deployment for a well-known network.
// Only the DKGManager address is stored here; all other contract addresses (registry,
// verifiers) are derived at startup by querying the manager's public immutable fields
// via the RPC endpoint.
type NetworkDeployment struct {
	// ChainID is the EIP-155 chain identifier. The node refuses RPC endpoints
	// that serve another chain.
	ChainID uint64
	// Manager is the deployed DKGManager contract address.
	Manager common.Address
	// StartBlock is the block at which the DKGManager was deployed. Used by
	// the UI and SDK to bound event log queries so they don't scan from
	// genesis (most free-tier RPC providers cap getLogs ranges at 10 000
	// blocks).
	StartBlock uint64
	// RPCs are public JSON-RPC endpoints, used when the operator sets none.
	RPCs []string
	// DecryptLookbackBlocks is how far behind the head a starting node scans
	// for ciphertexts awaiting decryption; 0 means the node's default
	// (50,400). On OP it matches the registry's INACTIVITY_WINDOW.
	DecryptLookbackBlocks uint64
}

// Deployed reports whether the preset has a DKGManager yet.
func (d NetworkDeployment) Deployed() bool {
	return d.Manager != (common.Address{})
}

// KnownNetworks maps canonical lowercase network names to their deployments.
// Add a new entry here after each production deployment.
var KnownNetworks = map[string]NetworkDeployment{
	"optimism": {
		ChainID:    10,
		Manager:    common.HexToAddress("0x401381a9711B473baf8757404B5514537A5216d7"),
		StartBlock: 158_592_323, // DKGManager deployment block on OP Mainnet (2026-10-23)
		// Public failover endpoints; operators put a private RPC first via
		// --web3.rpc. Both serve old log ranges and receipts, which share
		// rebuilds and transaction waits need (publicnode refuses both).
		RPCs: []string{
			"https://mainnet.optimism.io",
			"https://optimism.gateway.tenderly.co",
		},
		DecryptLookbackBlocks: 126_000,
	},
	"sepolia": {
		ChainID:    11155111,
		Manager:    common.HexToAddress("0xc73b7a868eca6ac7e3e647e2665aa16a793cf551"),
		StartBlock: 11_668_198, // DKGManager deployment block on Sepolia (v6 public testnet, 2026-09-09)
	},
}

// networkAliases maps short or alternative spellings to a canonical network name.
var networkAliases = map[string]string{
	"sep": "sepolia",
}

// NetworkByName resolves a network name (canonical or alias) to its deployment.
// The lookup is case-insensitive. Returns an error if the name is not recognised.
func NetworkByName(name string) (NetworkDeployment, error) {
	_, dep, err := ResolveNetwork(name)
	return dep, err
}

// ResolveNetwork returns the canonical network name and its deployment for the
// given name or alias. Returns an error if the name is not recognised.
func ResolveNetwork(name string) (string, NetworkDeployment, error) {
	lower := strings.ToLower(strings.TrimSpace(name))
	canonical := lower
	if alias, ok := networkAliases[lower]; ok {
		canonical = alias
	}
	dep, ok := KnownNetworks[canonical]
	if !ok {
		return "", NetworkDeployment{}, fmt.Errorf("unknown network %q — supported: %s", name, knownNetworkList())
	}
	return canonical, dep, nil
}

// knownNetworkList returns a human-readable comma-separated list of known
// canonical names and their aliases, e.g. "sepolia (sep)".
func knownNetworkList() string {
	// Build a reverse alias map: canonical → []alias
	reverseAliases := make(map[string][]string)
	for alias, canon := range networkAliases {
		reverseAliases[canon] = append(reverseAliases[canon], alias)
	}
	parts := make([]string, 0, len(KnownNetworks))
	for name := range KnownNetworks {
		entry := name
		if aliases := reverseAliases[name]; len(aliases) > 0 {
			entry += " (" + strings.Join(aliases, ", ") + ")"
		}
		parts = append(parts, entry)
	}
	return strings.Join(parts, ", ")
}
