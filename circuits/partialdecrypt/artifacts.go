package partialdecrypt

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/config"
)

// Artifacts contains the partial decrypt circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"partialdecrypt",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{
		RemoteURL: config.PartialDecryptCircuitURL,
		Hash:      mustArtifactHash(config.PartialDecryptCircuitHash),
	},
	&circuits.Artifact{
		RemoteURL: config.PartialDecryptProvingKeyURL,
		Hash:      mustArtifactHash(config.PartialDecryptProvingKeyHash),
	},
	&circuits.Artifact{
		RemoteURL: config.PartialDecryptVerificationKeyURL,
		Hash:      mustArtifactHash(config.PartialDecryptVerificationKeyHash),
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
