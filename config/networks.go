package config

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// NetworkDeployment is the canonical on-chain deployment for a well-known network.
// Only the DKGManager address is stored here; all other contract addresses (registry,
// verifiers) are derived at startup by querying the manager's public immutable fields
// via the RPC endpoint.
type NetworkDeployment struct {
	// ChainID is the EIP-155 chain identifier, stored for display and validation.
	// The authoritative chain ID is always fetched from the RPC endpoint at runtime.
	ChainID uint64
	// Manager is the deployed DKGManager contract address.
	Manager common.Address
	// StartBlock is the block at which the DKGManager was deployed. Used by the
	// webapp and SDK to bound event log queries so they don't scan from genesis
	// (most free-tier RPC providers cap getLogs ranges at 10 000 blocks).
	StartBlock uint64
}

// KnownNetworks maps canonical lowercase network names to their deployments.
// Add a new entry here after each production deployment.
var KnownNetworks = map[string]NetworkDeployment{
	"sepolia": {
		ChainID:    11155111,
		Manager:    common.HexToAddress("0xb68be96a967672004370798459ab4e7a28541be4"),
		StartBlock: 10_702_755, // approximate DKGManager deployment block on Sepolia
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
