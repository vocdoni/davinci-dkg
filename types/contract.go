package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ContractAddresses is the local address book for the DKG contract stack.
type ContractAddresses struct {
	Registry               common.Address
	Manager                common.Address
	ContributionVerifier   common.Address
	FinalizeVerifier       common.Address
	PartialDecryptVerifier common.Address
	DecryptCombineVerifier common.Address
	RevealSubmitVerifier   common.Address
	RevealShareVerifier    common.Address
}

// Validate checks that the address book contains the mandatory deployments.
// Registry is not required here because web3.New() derives it from the
// manager contract when it is not supplied.
func (c ContractAddresses) Validate() error {
	if c.Manager == (common.Address{}) {
		return fmt.Errorf("manager address is required")
	}
	if c.ContributionVerifier == (common.Address{}) {
		return fmt.Errorf("contribution verifier address is required")
	}
	if c.FinalizeVerifier == (common.Address{}) {
		return fmt.Errorf("finalize verifier address is required")
	}
	if c.PartialDecryptVerifier == (common.Address{}) {
		return fmt.Errorf("partial decrypt verifier address is required")
	}
	if c.DecryptCombineVerifier == (common.Address{}) {
		return fmt.Errorf("decrypt combine verifier address is required")
	}
	if c.RevealSubmitVerifier == (common.Address{}) {
		return fmt.Errorf("reveal submit verifier address is required")
	}
	if c.RevealShareVerifier == (common.Address{}) {
		return fmt.Errorf("reveal share verifier address is required")
	}
	return nil
}
