package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestDefaultNetworkIsOptimism(t *testing.T) {
	c := qt.New(t)

	name, dep, err := ResolveNetwork(DefaultNetwork)
	c.Assert(err, qt.IsNil)
	c.Assert(name, qt.Equals, "optimism")
	c.Assert(dep.ChainID, qt.Equals, uint64(10))
	c.Assert(len(dep.RPCs) > 1, qt.IsTrue)
	c.Assert(dep.DecryptLookbackBlocks, qt.Equals, uint64(126_000))

	_, err = NetworkByName("gnosis")
	c.Assert(err, qt.ErrorMatches, `unknown network "gnosis".*`)

	sepolia, err := NetworkByName("SEP")
	c.Assert(err, qt.IsNil)
	c.Assert(sepolia.ChainID, qt.Equals, uint64(11155111))
	c.Assert(sepolia.RPCs, qt.HasLen, 0)
	c.Assert(sepolia.Deployed(), qt.IsTrue)
	c.Assert(NetworkDeployment{}.Deployed(), qt.IsFalse)
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

// Every endpoint of the OP preset serves old log ranges and receipts: a
// member rebuilding a key share reads ContributionSubmitted logs from its
// epoch's seed block, often a day back, and transaction waits poll
// receipts. publicnode refuses both with HTTP 403 ("Archive requests require
// a personal token", measured 2026-10-19), which stalled DKG decryptions on
// OP, so it stays out of the preset.
func TestOptimismPresetServesOldLogsAndReceipts(t *testing.T) {
	c := qt.New(t)
	refusesHistory := []string{"publicnode.com"}
	dep := KnownNetworks["optimism"]
	c.Assert(len(dep.RPCs) >= 2, qt.IsTrue, qt.Commentf("the preset needs a failover endpoint"))
	for _, rpc := range dep.RPCs {
		for _, host := range refusesHistory {
			c.Assert(strings.Contains(rpc, host), qt.IsFalse, qt.Commentf("%s refuses old logs and receipts", rpc))
		}
	}
}
