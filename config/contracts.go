package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/types"
)

func ParseContractAddressesEnv(data []byte) (types.ContractAddresses, error) {
	values := map[string]string{}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return types.ContractAddresses{}, fmt.Errorf("invalid env line %q", line)
		}
		values[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		return types.ContractAddresses{}, err
	}

	// REGISTRY is optional — web3.New() derives it from the manager when absent.
	var registry common.Address
	if r := values["REGISTRY"]; r != "" {
		registry = common.HexToAddress(r)
	}
	// APP_MANAGER is optional for backward compatibility — when omitted,
	// callers can derive it on demand via manager.appManager() at runtime.
	var appManager common.Address
	if a := values["APP_MANAGER"]; a != "" {
		appManager = common.HexToAddress(a)
	}
	addresses := types.ContractAddresses{
		Registry:               registry,
		Manager:                common.HexToAddress(values["MANAGER"]),
		AppManager:             appManager,
		ContributionVerifier:   common.HexToAddress(values["CONTRIBUTION_VERIFIER"]),
		PoolKeyVerifier:        common.HexToAddress(values["POOL_KEY_VERIFIER"]),
		PartialDecryptVerifier: common.HexToAddress(values["PARTIAL_DECRYPT_VERIFIER"]),
		DecryptCombineVerifier: common.HexToAddress(values["DECRYPT_COMBINE_VERIFIER"]),
	}
	if err := addresses.Validate(); err != nil {
		return types.ContractAddresses{}, err
	}
	return addresses, nil
}

func LoadContractAddressesFile(path string) (types.ContractAddresses, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return types.ContractAddresses{}, err
	}
	return ParseContractAddressesEnv(data)
}
