package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/davinci-dkg/config"
)

type Config struct {
	Web3         Web3Config
	Log          LogConfig
	Datadir      string        `mapstructure:"datadir"`
	PrivKey      string        `mapstructure:"privkey"`
	Network      string        `mapstructure:"network"`
	ManagerAddr  string        `mapstructure:"manager"`
	PollInterval time.Duration `mapstructure:"poll-interval"`
}

type Web3Config struct {
	Network       string   `mapstructure:"network"`
	RPC           []string `mapstructure:"rpc"`
	GasMultiplier float64  `mapstructure:"gasMultiplier"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Output string `mapstructure:"output"`
}

func defaultConfig() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return &Config{
		Web3: Web3Config{
			Network:       "localhost",
			RPC:           []string{"http://127.0.0.1:8545"},
			GasMultiplier: 1.2,
		},
		Log: LogConfig{
			Level:  "info",
			Output: "stdout",
		},
		Datadir:      filepath.Join(home, ".davinci-dkg"),
		PollInterval: 5 * time.Second,
	}
}

func loadConfig() (*Config, error) {
	return loadConfigFromArgs(os.Args[1:])
}

func loadConfigFromArgs(args []string) (*Config, error) {
	cfg := defaultConfig()

	fs := flag.NewFlagSet("davinci-dkg-node", flag.ContinueOnError)
	fs.String("network", cfg.Network, "well-known network preset (e.g. sepolia, sep); sets the DKGManager address automatically")
	fs.String("web3.network", cfg.Web3.Network, "network display name (overridden by --network when a preset is matched)")
	fs.StringSlice("web3.rpc", cfg.Web3.RPC, "web3 rpc endpoints")
	fs.Float64("web3.gasMultiplier", cfg.Web3.GasMultiplier, "gas multiplier")
	fs.String("log.level", cfg.Log.Level, "log level")
	fs.String("log.output", cfg.Log.Output, "log output")
	fs.String("datadir", cfg.Datadir, "data directory")
	fs.String("privkey", cfg.PrivKey, "hex private key for signing transactions")
	fs.String("manager", cfg.ManagerAddr, "DKGManager contract address (optional when --network is set)")
	fs.Duration("poll-interval", cfg.PollInterval, "chain polling interval")
	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("parse flags: %w", err)
	}

	v := viper.New()
	v.SetEnvPrefix("DAVINCI_DKG")
	// Flag names use dots for nesting (e.g. web3.rpc) and dashes for
	// compounds (e.g. poll-interval). Both become underscores in env vars
	// so that DAVINCI_DKG_WEB3_RPC and DAVINCI_DKG_POLL_INTERVAL work.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.BindPFlags(fs); err != nil {
		return nil, fmt.Errorf("bind flags: %w", err)
	}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, validateConfig(cfg)
}

func validateConfig(cfg *Config) error {
	if cfg.Web3.GasMultiplier <= 0 {
		return fmt.Errorf("gas multiplier must be greater than 0")
	}
	if len(cfg.Web3.RPC) == 0 {
		return fmt.Errorf("at least one web3 rpc endpoint is required")
	}
	// Validate the network name early so the user gets a clear error message
	// rather than a confusing failure later during chain connection.
	if cfg.Network != "" {
		if _, err := config.NetworkByName(cfg.Network); err != nil {
			return err
		}
	}
	return nil
}

// HasChainConfig reports whether enough configuration is present to connect to
// the chain and participate in DKG rounds. A private key is always required; the
// DKGManager address may come from --manager or from a --network preset.
func (c *Config) HasChainConfig() bool {
	if c.PrivKey == "" {
		return false
	}
	if c.ManagerAddr != "" {
		return true
	}
	if c.Network != "" {
		_, err := config.NetworkByName(c.Network)
		return err == nil
	}
	return false
}

// resolvedManagerAddr returns the effective DKGManager address: the explicit
// --manager flag takes precedence; when absent the network preset is used.
func (c *Config) resolvedManagerAddr() string {
	if c.ManagerAddr != "" {
		return c.ManagerAddr
	}
	if c.Network != "" {
		dep, err := config.NetworkByName(c.Network)
		if err == nil {
			return dep.Manager.Hex()
		}
	}
	return ""
}

// resolvedNetworkName returns the canonical network name for display/logging.
func (c *Config) resolvedNetworkName() string {
	if c.Network != "" {
		canonical, _, err := config.ResolveNetwork(c.Network)
		if err == nil {
			return canonical
		}
	}
	return c.Web3.Network
}
