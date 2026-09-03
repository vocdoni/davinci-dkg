package node

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
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

	// AutoCreateEpochs makes this node race other nodes to fire `createEpoch`
	// once `nextEpochStartBlock()` is reached. Each candidate sleeps a random
	// jitter (0..AutoCreateJitter) before firing, so the population spreads
	// out and most calls succeed cheaply with one revert per loser. Default
	// true; disable for nodes that should only participate, not propose.
	AutoCreateEpochs bool          `mapstructure:"auto-create-epochs"`
	AutoCreateJitter time.Duration `mapstructure:"auto-create-jitter"`

	// DecryptLookbackBlocks is how far behind the chain head a freshly
	// started node scans for CiphertextSubmitted events it may still have
	// to serve. Ciphertexts older than this are ignored on restart.
	DecryptLookbackBlocks uint64 `mapstructure:"decrypt-lookback-blocks"`

	// EpochPolicy is the per-epoch policy this node proposes when it wins
	// the auto-create race. All fields are optional: missing fields fall
	// back to safe defaults (committee of 4, threshold 3, α=1.5, no
	// decryption-policy gating). Only consulted when AutoCreateEpochs is
	// true.
	EpochPolicy EpochPolicyConfig `mapstructure:"epoch-policy"`
}

type EpochPolicyConfig struct {
	Threshold             uint16 `mapstructure:"threshold"`
	CommitteeSize         uint16 `mapstructure:"committee-size"`
	MinValidContributions uint16 `mapstructure:"min-valid-contributions"`
	LotteryAlphaBps       uint16 `mapstructure:"lottery-alpha-bps"`
}

// validate mirrors the contract's InvalidPolicy checks and the circuit's
// committee cap so a bad policy fails at startup, not at createEpoch.
// Adaptive reports whether the node sizes the committee from the registry at
// creation time instead of using fixed numbers.
func (p EpochPolicyConfig) Adaptive() bool { return p.CommitteeSize == 0 }

func (p EpochPolicyConfig) validate() error {
	if p.Adaptive() {
		if p.Threshold != 0 || p.MinValidContributions != 0 {
			return fmt.Errorf("threshold and min valid contributions need an explicit committee size")
		}
		if p.LotteryAlphaBps < 10_000 {
			return fmt.Errorf("lottery alpha %d bps must be at least 10000 (1.0)", p.LotteryAlphaBps)
		}
		return nil
	}
	if p.Threshold < 1 {
		return fmt.Errorf("threshold must be at least 1")
	}
	if p.CommitteeSize < p.Threshold {
		return fmt.Errorf("committee size %d must be at least the threshold %d", p.CommitteeSize, p.Threshold)
	}
	if p.CommitteeSize > ccommon.MaxN {
		return fmt.Errorf("committee size %d exceeds the circuit cap MaxN=%d", p.CommitteeSize, ccommon.MaxN)
	}
	if p.MinValidContributions < p.Threshold {
		return fmt.Errorf("min valid contributions %d must be at least the threshold %d", p.MinValidContributions, p.Threshold)
	}
	if p.MinValidContributions > p.CommitteeSize {
		return fmt.Errorf("min valid contributions %d cannot exceed the committee size %d", p.MinValidContributions, p.CommitteeSize)
	}
	if p.LotteryAlphaBps < 10_000 {
		return fmt.Errorf("lottery alpha %d bps must be at least 10000 (1.0)", p.LotteryAlphaBps)
	}
	return nil
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
		Datadir:          filepath.Join(home, ".davinci-dkg"),
		PollInterval:     5 * time.Second,
		AutoCreateEpochs: true,
		AutoCreateJitter: 12 * time.Second,
		// ~7 days at 12 s blocks; matches the registry's default INACTIVITY_WINDOW.
		DecryptLookbackBlocks: 50_400,
		// Committee size 0 means "derive from the registry" (see adaptivePolicy).
		EpochPolicy: EpochPolicyConfig{
			Threshold:             0,
			CommitteeSize:         0,
			MinValidContributions: 0,
			LotteryAlphaBps:       15000,
		},
	}
}

// LoadConfig parses flags and DAVINCI_DKG_* environment variables.
func LoadConfig() (*Config, error) {
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
	fs.Bool("auto-create-epochs", cfg.AutoCreateEpochs, "race other nodes to fire createEpoch once nextEpochStartBlock() is reached (default true; disable to participate only)")
	fs.Duration("auto-create-jitter", cfg.AutoCreateJitter, "max random delay before firing the auto-create transaction (spreads contention)")
	fs.Uint64("decrypt-lookback-blocks", cfg.DecryptLookbackBlocks, "on startup, scan this many blocks behind head for ciphertexts still awaiting decryption")
	fs.Uint16("epoch-policy.threshold", cfg.EpochPolicy.Threshold, "Shamir threshold t when this node proposes an epoch")
	fs.Uint16("epoch-policy.committee-size", cfg.EpochPolicy.CommitteeSize, "committee size n when this node proposes an epoch (0 = derive from the registry)")
	fs.Uint16("epoch-policy.min-valid-contributions", cfg.EpochPolicy.MinValidContributions, "minValidContributions when this node proposes an epoch")
	fs.Uint16("epoch-policy.lottery-alpha-bps", cfg.EpochPolicy.LotteryAlphaBps, "lottery oversubscription α in basis points (10000 = 1.0)")
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
	if cfg.PollInterval <= 0 {
		return fmt.Errorf("poll interval must be greater than 0, got %s", cfg.PollInterval)
	}
	if err := cfg.EpochPolicy.validate(); err != nil {
		return fmt.Errorf("epoch policy: %w", err)
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
// the chain and participate in DKG epochs. A private key is always required; the
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

// ResolvedNetworkName returns the canonical network name for display/logging.
func (c *Config) ResolvedNetworkName() string {
	if c.Network != "" {
		canonical, _, err := config.ResolveNetwork(c.Network)
		if err == nil {
			return canonical
		}
	}
	return c.Web3.Network
}
