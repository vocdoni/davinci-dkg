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

	addresses := types.ContractAddresses{
		Registry:               common.HexToAddress(values["REGISTRY"]),
		Manager:                common.HexToAddress(values["MANAGER"]),
		ContributionVerifier:   common.HexToAddress(values["CONTRIBUTION_VERIFIER"]),
		FinalizeVerifier:       common.HexToAddress(values["FINALIZE_VERIFIER"]),
		PartialDecryptVerifier: common.HexToAddress(values["PARTIAL_DECRYPT_VERIFIER"]),
		DecryptCombineVerifier: common.HexToAddress(values["DECRYPT_COMBINE_VERIFIER"]),
		RevealSubmitVerifier:   common.HexToAddress(values["REVEAL_SUBMIT_VERIFIER"]),
		RevealShareVerifier:    common.HexToAddress(values["REVEAL_SHARE_VERIFIER"]),
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
