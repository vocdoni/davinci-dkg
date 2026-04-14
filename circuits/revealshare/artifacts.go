package revealshare

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/config"
)

// MaxShares is an alias of the single shared constant `circuits/common.MaxN`.
const MaxShares = ccommon.MaxN

// Artifacts contains the reveal share circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"revealshare",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{RemoteURL: config.RevealShareCircuitURL, Hash: mustArtifactHash(config.RevealShareCircuitHash)},
	&circuits.Artifact{
		RemoteURL: config.RevealShareProvingKeyURL,
		Hash:      mustArtifactHash(config.RevealShareProvingKeyHash),
	},
	&circuits.Artifact{
		RemoteURL: config.RevealShareVerificationKeyURL,
		Hash:      mustArtifactHash(config.RevealShareVerificationKeyHash),
	},
)

func mustArtifactHash(value string) []byte {
	if value == "" {
		return nil
	}
	raw, err := hex.DecodeString(value)
	if err != nil {
		panic(err)
	}
	return raw
}
