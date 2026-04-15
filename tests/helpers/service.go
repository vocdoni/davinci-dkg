package helpers

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/vocdoni/davinci-dkg/config"
	"github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

type HarnessConfig struct {
	RPCURL            string
	Addresses         types.ContractAddresses
	PrivateKey        string
	BootstrapNodeKeys bool
}

type TestServices struct {
	RPCURL    string
	Addresses types.ContractAddresses
	Contracts *web3.Contracts
	Registry  *golangtypes.DKGRegistry
	Manager   *golangtypes.DKGManager
	TxManager *txmanager.Manager
}

func (s *TestServices) CallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func LoadHarnessConfigFromEnv() (*HarnessConfig, error) {
	rpcURL := os.Getenv(TestRPCURLEnvVar)
	if rpcURL == "" {
		return nil, fmt.Errorf("rpc url is required")
	}

	addressesPath := os.Getenv(TestAddressesEnvVar)
	if addressesPath == "" {
		return nil, fmt.Errorf("addresses file is required")
	}

	addresses, err := config.LoadContractAddressesFile(addressesPath)
	if err != nil {
		return nil, err
	}

	return &HarnessConfig{
		RPCURL:            rpcURL,
		Addresses:         addresses,
		PrivateKey:        privateKeyFromEnv(),
		BootstrapNodeKeys: false,
	}, nil
}

func ParseAddressesEnvPayload(data []byte) (types.ContractAddresses, error) {
	return config.ParseContractAddressesEnv(data)
}

func NewTestServices(ctx context.Context) (*TestServices, func(), error) {
	cfg, cleanup, err := setupHarnessConfig(ctx)
	if err != nil {
		return nil, nil, err
	}

	contracts, err := web3.New([]string{cfg.RPCURL}, cfg.Addresses)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	txm, err := txmanager.New(contracts.Pool().Current, contracts.ChainID, cfg.PrivateKey)
	if err != nil {
		_ = contracts.Close()
		cleanup()
		return nil, nil, err
	}

	registry, err := golangtypes.NewDKGRegistry(cfg.Addresses.Registry, contracts.Client())
	if err != nil {
		_ = contracts.Close()
		cleanup()
		return nil, nil, err
	}

	manager, err := golangtypes.NewDKGManager(cfg.Addresses.Manager, contracts.Client())
	if err != nil {
		_ = contracts.Close()
		cleanup()
		return nil, nil, err
	}

	services := &TestServices{
		RPCURL:    cfg.RPCURL,
		Addresses: cfg.Addresses,
		Contracts: contracts,
		Registry:  registry,
		Manager:   manager,
		TxManager: txm,
	}

	if cfg.BootstrapNodeKeys {
		if err := bootstrapLocalNodeKeys(ctx, services); err != nil {
			_ = services.Contracts.Close()
			cleanup()
			return nil, nil, err
		}
	}

	finalCleanup := func() {
		_ = services.Contracts.Close()
		cleanup()
	}

	select {
	case <-ctx.Done():
		finalCleanup()
		return nil, nil, ctx.Err()
	default:
	}

	return services, finalCleanup, nil
}

func setupHarnessConfig(ctx context.Context) (*HarnessConfig, func(), error) {
	cfg, err := loadHarnessConfigIfPresent()
	if err != nil {
		return nil, nil, err
	}
	if cfg != nil {
		return cfg, func() {}, nil
	}
	return setupLocalHarness(ctx)
}

func loadHarnessConfigIfPresent() (*HarnessConfig, error) {
	rpcURL := os.Getenv(TestRPCURLEnvVar)
	addressesPath := os.Getenv(TestAddressesEnvVar)
	if rpcURL == "" && addressesPath == "" {
		return nil, nil
	}
	return LoadHarnessConfigFromEnv()
}

func setupLocalHarness(ctx context.Context) (*HarnessConfig, func(), error) {
	anvilPort, err := reserveTCPPort()
	if err != nil {
		return nil, nil, err
	}
	deployerPort, err := reserveTCPPort()
	if err != nil {
		return nil, nil, err
	}

	composeEnv := map[string]string{
		AnvilPortEnvVarName:          fmt.Sprintf("%d", anvilPort),
		DeployerServerPortEnvVarName: fmt.Sprintf("%d", deployerPort),
		"CONTRIBUTION_VERIFIER":      common.HexToAddress("0x3000000000000000000000000000000000000003").Hex(),
		"PARTIAL_DECRYPT_VERIFIER":   common.HexToAddress("0x4000000000000000000000000000000000000004").Hex(),
	}

	composePath, err := repoPath("tests", "docker", "docker-compose.yml")
	if err != nil {
		return nil, nil, err
	}

	compose, err := tc.NewDockerCompose(composePath)
	if err != nil {
		return nil, nil, fmt.Errorf("create docker compose: %w", err)
	}
	if err := compose.WithEnv(composeEnv).Up(ctx, tc.Wait(true), tc.RemoveOrphans(true)); err != nil {
		return nil, nil, fmt.Errorf("start docker compose: %w", err)
	}

	cleanup := func() {
		downCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		_ = compose.Down(downCtx, tc.RemoveOrphans(true), tc.RemoveVolumes(true))
	}

	rpcURL := fmt.Sprintf("http://127.0.0.1:%d", anvilPort)
	if err := waitRPCReady(ctx, rpcURL); err != nil {
		cleanup()
		return nil, nil, err
	}

	deployerURL := fmt.Sprintf("http://127.0.0.1:%d", deployerPort)
	addresses, err := waitAddresses(ctx, deployerURL)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	return &HarnessConfig{
		RPCURL:            rpcURL,
		Addresses:         addresses,
		PrivateKey:        LocalAccountPrivKey,
		BootstrapNodeKeys: true,
	}, cleanup, nil
}

func privateKeyFromEnv() string {
	if privateKey := os.Getenv(TestPrivateKeyEnvVar); privateKey != "" {
		return privateKey
	}
	return LocalAccountPrivKey
}

func reserveTCPPort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, fmt.Errorf("reserve tcp port: %w", err)
	}
	defer func() { _ = listener.Close() }()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func waitRPCReady(ctx context.Context, rpcURL string) error {
	client := &http.Client{Timeout: 2 * time.Second}
	body := `{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}`
	return WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, rpcURL, strings.NewReader(body))
		if err != nil {
			return false
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			return false
		}
		defer func() { _ = resp.Body.Close() }()
		return resp.StatusCode == http.StatusOK
	})
}

func waitAddresses(ctx context.Context, deployerURL string) (types.ContractAddresses, error) {
	var addresses types.ContractAddresses
	err := WaitUntilCondition(ctx, DefaultWaitInterval, func() bool {
		endpoint := fmt.Sprintf("%s/addresses.env", deployerURL)
		req, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if reqErr != nil {
			return false
		}
		resp, getErr := http.DefaultClient.Do(req)
		if getErr != nil {
			return false
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusOK {
			return false
		}
		payload, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return false
		}
		addresses, readErr = ParseAddressesEnvPayload(payload)
		return readErr == nil
	})
	if err != nil {
		return types.ContractAddresses{}, fmt.Errorf("wait addresses from deployer: %w", err)
	}
	return addresses, nil
}

func repoPath(parts ...string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	for {
		candidate := filepath.Join(dir, "go.mod")
		if _, statErr := os.Stat(candidate); statErr == nil {
			allParts := append([]string{dir}, parts...)
			return filepath.Join(allParts...), nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("repository root not found")
		}
		dir = parent
	}
}
