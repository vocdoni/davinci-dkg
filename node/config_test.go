package node

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
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

	c.Run("rejects missing rpc endpoints", func(c *qt.C) {
		cfg := defaultConfig()
		cfg.Web3.RPC = nil

		err := validateConfig(cfg)

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "web3 rpc")
	})
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
		{"zero threshold", func(c *Config) { c.EpochPolicy.Threshold = 0 }, ".*threshold.*"},
		{"committee below threshold", func(c *Config) { c.EpochPolicy.CommitteeSize = 2 }, ".*committee size.*"},
		{"committee above MaxN", func(c *Config) {
			c.EpochPolicy.CommitteeSize = ccommon.MaxN + 1
			c.EpochPolicy.MinValidContributions = ccommon.MaxN + 1
		}, ".*committee size.*"},
		{"min valid below threshold", func(c *Config) { c.EpochPolicy.MinValidContributions = 2 }, ".*min valid contributions.*"},
		{"min valid above committee", func(c *Config) { c.EpochPolicy.MinValidContributions = 5 }, ".*min valid contributions.*"},
		{"alpha below 1.0", func(c *Config) { c.EpochPolicy.LotteryAlphaBps = 9_999 }, ".*alpha.*"},
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
