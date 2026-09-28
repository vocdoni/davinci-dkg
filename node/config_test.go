package node

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/config"
)

func TestDefaultConfig(t *testing.T) {
	c := qt.New(t)

	cfg := defaultConfig()

	c.Assert(cfg.Log.Level, qt.Equals, "info")
	c.Assert(cfg.Log.Output, qt.Equals, "stdout")
	c.Assert(cfg.Web3.GasMultiplier, qt.Equals, 1.2)
	c.Assert(cfg.Web3.Network, qt.Equals, "localhost")
}

func TestValidateConfig(t *testing.T) {
	c := qt.New(t)

	c.Run("rejects non positive gas multiplier", func(c *qt.C) {
		cfg := defaultConfig()
		cfg.Web3.GasMultiplier = 0

		err := validateConfig(cfg)

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "gas multiplier")
	})

	c.Run("rejects a preset without rpc endpoints", func(c *qt.C) {
		cfg := defaultConfig()
		cfg.Network = "sepolia"

		err := validateConfig(cfg)

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "web3 rpc")
	})

	c.Run("rejects an unknown network", func(c *qt.C) {
		cfg := defaultConfig()
		cfg.Network = "mainnet"
		c.Assert(validateConfig(cfg), qt.ErrorMatches, `unknown network "mainnet".*`)
	})
}

func TestNetworkResolution(t *testing.T) {
	c := qt.New(t)
	for _, env := range []string{"DAVINCI_DKG_NETWORK", "DAVINCI_DKG_MANAGER", "DAVINCI_DKG_WEB3_RPC"} {
		t.Setenv(env, "")
	}
	gnosis := config.KnownNetworks[config.DefaultNetwork]
	sepolia := config.KnownNetworks["sepolia"]
	const custom = "0x00000000000000000000000000000000000000c0"
	cases := []struct {
		name       string
		args       []string
		network    string
		manager    string
		rpc        []string
		chainCheck uint64
	}{
		{"default", nil, "gnosis", gnosis.Manager.Hex(), gnosis.RPCs, 100},
		{
			"explicit preset",
			[]string{"--network=sep", "--web3.rpc=https://a,https://b"},
			"sepolia", sepolia.Manager.Hex(),
			[]string{"https://a", "https://b"},
			sepolia.ChainID,
		},
		{"custom manager", []string{"--manager=" + custom}, "localhost", custom, []string{localRPC}, 0},
		{
			"custom manager on a preset chain",
			[]string{"--network=gnosis", "--manager=" + custom},
			"gnosis", custom, gnosis.RPCs, 0,
		},
		{
			"custom rpc on the default preset",
			[]string{"--web3.rpc=https://mine"},
			"gnosis", gnosis.Manager.Hex(),
			[]string{"https://mine"},
			100,
		},
	}
	for _, tc := range cases {
		c.Run(tc.name, func(c *qt.C) {
			cfg, err := loadConfigFromArgs(tc.args)
			c.Assert(err, qt.IsNil)
			c.Assert(cfg.ResolvedNetworkName(), qt.Equals, tc.network)
			c.Assert(cfg.resolvedManagerAddr(), qt.Equals, tc.manager)
			c.Assert(cfg.Web3.RPC, qt.DeepEquals, tc.rpc)
			c.Assert(cfg.requiredChainID(), qt.Equals, tc.chainCheck)
		})
	}

	// The preset's endpoint list is copied, never aliased.
	cfg, err := loadConfigFromArgs(nil)
	c.Assert(err, qt.IsNil)
	cfg.Web3.RPC[0] = "https://changed"
	c.Assert(config.KnownNetworks[config.DefaultNetwork].RPCs[0], qt.Not(qt.Equals), "https://changed")
}

func TestLoadConfigNetworkFromEnv(t *testing.T) {
	c := qt.New(t)
	t.Setenv("DAVINCI_DKG_NETWORK", "")
	t.Setenv("DAVINCI_DKG_MANAGER", "")
	cfg, err := loadConfigFromArgs(nil)
	c.Assert(err, qt.IsNil)
	c.Assert(cfg.ResolvedNetworkName(), qt.Equals, config.DefaultNetwork)

	t.Setenv("DAVINCI_DKG_MANAGER", "0x00000000000000000000000000000000000000c0")
	cfg, err = loadConfigFromArgs(nil)
	c.Assert(err, qt.IsNil)
	c.Assert(cfg.requiredChainID(), qt.Equals, uint64(0))
	c.Assert(cfg.Web3.RPC, qt.DeepEquals, []string{localRPC})
}

// Misconfiguration must fail at startup with a clear message rather than
// as a revert (InvalidPolicy) on the first auto-created epoch or a panic in
// time.NewTicker.
func TestValidateConfigRejectsBadPollIntervalAndEpochPolicy(t *testing.T) {
	c := qt.New(t)
	cases := []struct {
		name  string
		mut   func(*Config)
		match string
	}{
		{"zero poll interval", func(c *Config) { c.PollInterval = 0 }, ".*poll interval.*"},
		{"negative poll interval", func(c *Config) { c.PollInterval = -time.Second }, ".*poll interval.*"},
		{"zero threshold", func(c *Config) { fixedPolicy(c); c.EpochPolicy.Threshold = 0 }, ".*threshold.*"},
		{"committee below threshold", func(c *Config) { fixedPolicy(c); c.EpochPolicy.CommitteeSize = 2 }, ".*committee size.*"},
		{"committee above MaxN", func(c *Config) {
			fixedPolicy(c)
			c.EpochPolicy.CommitteeSize = ccommon.MaxN + 1
			c.EpochPolicy.MinValidContributions = ccommon.MaxN + 1
		}, ".*committee size.*"},
		{"min valid below threshold", func(c *Config) { fixedPolicy(c); c.EpochPolicy.MinValidContributions = 2 }, ".*min valid contributions.*"},
		{"min valid above committee", func(c *Config) { fixedPolicy(c); c.EpochPolicy.MinValidContributions = 5 }, ".*min valid contributions.*"},
		{"alpha below 1.0", func(c *Config) { c.EpochPolicy.LotteryAlphaBps = 9_999 }, ".*alpha.*"},
		{"adaptive with a threshold", func(c *Config) { c.EpochPolicy.Threshold = 2 }, ".*explicit committee size.*"},
	}
	for _, tc := range cases {
		c.Run(tc.name, func(c *qt.C) {
			cfg := defaultConfig()
			tc.mut(cfg)
			c.Assert(validateConfig(cfg), qt.ErrorMatches, tc.match)
		})
	}
	c.Assert(validateConfig(defaultConfig()), qt.IsNil)
}

func TestLoadConfigReportsInvalidFlags(t *testing.T) {
	c := qt.New(t)
	_, err := loadConfigFromArgs([]string{"--poll-interval=0s"})
	c.Assert(err, qt.ErrorMatches, ".*poll interval.*")
	_, err = loadConfigFromArgs([]string{"--epoch-policy.lottery-alpha-bps=100"})
	c.Assert(err, qt.ErrorMatches, ".*alpha.*")
}

// fixedPolicy replaces the adaptive default with explicit numbers so the
// per-field checks can be exercised.
func fixedPolicy(c *Config) {
	c.EpochPolicy = EpochPolicyConfig{Threshold: 3, CommitteeSize: 4, MinValidContributions: 3, LotteryAlphaBps: 15_000}
}
