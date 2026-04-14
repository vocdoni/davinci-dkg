package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
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
