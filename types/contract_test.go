package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestContractAddressesValidate(t *testing.T) {
	c := qt.New(t)

	addresses := ContractAddresses{
		Registry:               common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Manager:                common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributionVerifier:   common.HexToAddress("0x3000000000000000000000000000000000000003"),
		FinalizeVerifier:       common.HexToAddress("0x4000000000000000000000000000000000000004"),
		PartialDecryptVerifier: common.HexToAddress("0x5000000000000000000000000000000000000005"),
		DecryptCombineVerifier: common.HexToAddress("0x6000000000000000000000000000000000000006"),
		RevealSubmitVerifier:   common.HexToAddress("0x6500000000000000000000000000000000000006"),
		RevealShareVerifier:    common.HexToAddress("0x7000000000000000000000000000000000000007"),
	}

	c.Assert(addresses.Validate(), qt.IsNil)
}

func TestContractAddressesValidateRejectsMissingVerifier(t *testing.T) {
	c := qt.New(t)

	addresses := ContractAddresses{
		Registry: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Manager:  common.HexToAddress("0x2000000000000000000000000000000000000002"),
	}

	err := addresses.Validate()
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "contribution verifier")
}

func TestContractAddressesValidateRejectsMissingPartialDecryptVerifier(t *testing.T) {
	c := qt.New(t)

	addresses := ContractAddresses{
		Registry:               common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Manager:                common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributionVerifier:   common.HexToAddress("0x3000000000000000000000000000000000000003"),
		FinalizeVerifier:       common.HexToAddress("0x4000000000000000000000000000000000000004"),
		DecryptCombineVerifier: common.HexToAddress("0x6000000000000000000000000000000000000006"),
		RevealSubmitVerifier:   common.HexToAddress("0x6500000000000000000000000000000000000006"),
		RevealShareVerifier:    common.HexToAddress("0x7000000000000000000000000000000000000007"),
	}

	err := addresses.Validate()
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "partial decrypt verifier")
}
