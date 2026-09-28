package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestDefaultNetworkIsGnosis(t *testing.T) {
	c := qt.New(t)

	name, dep, err := ResolveNetwork(DefaultNetwork)
	c.Assert(err, qt.IsNil)
	c.Assert(name, qt.Equals, "gnosis")
	c.Assert(dep.ChainID, qt.Equals, uint64(100))
	c.Assert(dep.Manager, qt.Equals, common.HexToAddress("0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B"))
	c.Assert(dep.StartBlock, qt.Equals, uint64(48_483_860))
	c.Assert(len(dep.RPCs) > 1, qt.IsTrue)

	sepolia, err := NetworkByName("SEP")
	c.Assert(err, qt.IsNil)
	c.Assert(sepolia.ChainID, qt.Equals, uint64(11155111))
	c.Assert(sepolia.RPCs, qt.HasLen, 0)
}

// sdk/src/networks.ts mirrors KnownNetworks: every preset's manager, start
// block and endpoints appear there.
func TestKnownNetworksMatchTheSDK(t *testing.T) {
	c := qt.New(t)

	raw, err := os.ReadFile("../sdk/src/networks.ts")
	if errors.Is(err, os.ErrNotExist) {
		t.Skip("sdk sources not present")
	}
	c.Assert(err, qt.IsNil)
	src := strings.ToLower(strings.ReplaceAll(string(raw), "_", ""))
	for name, dep := range KnownNetworks {
		c.Assert(src, qt.Contains, name+": {")
		c.Assert(src, qt.Contains, fmt.Sprintf("chainid: %d,", dep.ChainID))
		c.Assert(src, qt.Contains, "'"+strings.ToLower(dep.Manager.Hex())+"'")
		c.Assert(src, qt.Contains, fmt.Sprintf("startblock: %dn,", dep.StartBlock))
		for _, rpc := range dep.RPCs {
			c.Assert(src, qt.Contains, "'"+strings.ToLower(rpc)+"'")
		}
	}
}
