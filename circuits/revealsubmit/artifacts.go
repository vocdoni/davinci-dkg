package revealsubmit

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/config"
)

// Artifacts contains the reveal-submit circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"revealsubmit",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{RemoteURL: config.RevealSubmitCircuitURL, Hash: mustArtifactHash(config.RevealSubmitCircuitHash)},
	&circuits.Artifact{
		RemoteURL: config.RevealSubmitProvingKeyURL,
		Hash:      mustArtifactHash(config.RevealSubmitProvingKeyHash),
	},
	&circuits.Artifact{
		RemoteURL: config.RevealSubmitVerificationKeyURL,
		Hash:      mustArtifactHash(config.RevealSubmitVerificationKeyHash),
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
