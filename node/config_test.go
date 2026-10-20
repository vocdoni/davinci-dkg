package node

import (
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/config"
)

// deployDefaultPreset installs a fake manager on the default preset (the real
// one is a placeholder until the deployment) and restores it afterwards.
func deployDefaultPreset(t *testing.T) config.NetworkDeployment {
	t.Helper()
	orig := config.KnownNetworks[config.DefaultNetwork]
	dep := orig
	dep.Manager = common.HexToAddress("0x00000000000000000000000000000000000000d1")
	config.KnownNetworks[config.DefaultNetwork] = dep
	t.Cleanup(func() { config.KnownNetworks[config.DefaultNetwork] = orig })
	return dep
}

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

	c.Run("rejects an undeployed preset unless --manager is given", func(c *qt.C) {
		orig := config.KnownNetworks[config.DefaultNetwork]
		config.KnownNetworks[config.DefaultNetwork] = config.NetworkDeployment{ChainID: orig.ChainID, RPCs: orig.RPCs}
		c.Cleanup(func() { config.KnownNetworks[config.DefaultNetwork] = orig })

		c.Assert(validateConfig(defaultConfig()), qt.ErrorMatches,
			"network "+config.DefaultNetwork+" has no DKGManager yet: pass --manager")

		cfg := defaultConfig()
		cfg.Network = config.DefaultNetwork
		c.Assert(validateConfig(cfg), qt.ErrorMatches, ".*no DKGManager yet.*")

		cfg.ManagerAddr = "0x00000000000000000000000000000000000000c0"
		c.Assert(validateConfig(cfg), qt.IsNil)
	})

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
	def := deployDefaultPreset(t)
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
		{"default", nil, config.DefaultNetwork, def.Manager.Hex(), def.RPCs, def.ChainID},
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
			[]string{"--network=" + config.DefaultNetwork, "--manager=" + custom},
			config.DefaultNetwork, custom, def.RPCs, 0,
		},
		{
			"custom rpc on the default preset",
			[]string{"--web3.rpc=https://mine"},
			config.DefaultNetwork, def.Manager.Hex(),
			[]string{"https://mine"},
			def.ChainID,
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
	deployDefaultPreset(t)
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

func TestDecryptLookbackFollowsThePreset(t *testing.T) {
	c := qt.New(t)
	for _, env := range []string{"DAVINCI_DKG_NETWORK", "DAVINCI_DKG_MANAGER", "DAVINCI_DKG_DECRYPT_LOOKBACK_BLOCKS"} {
		t.Setenv(env, "")
	}
	def := deployDefaultPreset(t)
	c.Assert(def.DecryptLookbackBlocks, qt.Equals, uint64(126_000))
	const custom = "0x00000000000000000000000000000000000000c0"

	lookback := func(c *qt.C, args ...string) uint64 {
		cfg, err := loadConfigFromArgs(args)
		c.Assert(err, qt.IsNil)
		return cfg.DecryptLookbackBlocks
	}
	c.Run("neither flag nor env uses the preset's value", func(c *qt.C) {
		c.Assert(lookback(c), qt.Equals, def.DecryptLookbackBlocks)
	})
	c.Run("a preset without a value keeps the default", func(c *qt.C) {
		c.Assert(lookback(c, "--network=sepolia", "--web3.rpc=https://a"), qt.Equals, uint64(50_400))
	})
	c.Run("the flag wins", func(c *qt.C) {
		c.Assert(lookback(c, "--decrypt-lookback-blocks=1000"), qt.Equals, uint64(1000))
	})
	c.Run("the env wins", func(c *qt.C) {
		c.Setenv("DAVINCI_DKG_DECRYPT_LOOKBACK_BLOCKS", "2000")
		c.Assert(lookback(c), qt.Equals, uint64(2000))
	})
	c.Run("a custom deployment keeps the default", func(c *qt.C) {
		c.Assert(lookback(c, "--manager="+custom), qt.Equals, uint64(50_400))
	})
}

// Misconfiguration must fail at startup with a clear message rather than
// as a revert (InvalidPolicy) on the first auto-created epoch or a panic in
// time.NewTicker.
func TestValidateConfigRejectsBadPollIntervalAndEpochPolicy(t *testing.T) {
	c := qt.New(t)
	deployDefaultPreset(t)
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
	deployDefaultPreset(t)
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

// A --web3.rpc value that does not parse (a stray quote) must not be echoed
// back: main prints the error, and the endpoints carry API keys.
func TestMalformedRPCFlagIsNotEchoed(t *testing.T) {
	c := qt.New(t)
	value := `https://rpcuser:RPCPASSWORD@rpc.example/v2/RPCPATHKEY?apikey=RPCQUERYKEY,"https://other.example`
	_, err := loadConfigFromArgs([]string{"--web3.rpc=" + value})
	c.Assert(err, qt.ErrorMatches, `.*--web3.rpc.*`)
	for _, secret := range []string{"rpcuser", "RPCPASSWORD", "RPCPATHKEY", "RPCQUERYKEY"} {
		c.Assert(strings.Contains(err.Error(), secret), qt.IsFalse, qt.Commentf("%s leaked: %v", secret, err))
	}
}
