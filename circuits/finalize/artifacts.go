package finalize

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/config"
)

// Artifacts contains the finalization circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"finalize",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{RemoteURLs: config.FinalizeCircuitURLs, Hash: mustArtifactHash(config.FinalizeCircuitHash)},
	&circuits.Artifact{RemoteURLs: config.FinalizeProvingKeyURLs, Hash: mustArtifactHash(config.FinalizeProvingKeyHash)},
	&circuits.Artifact{
		RemoteURLs: config.FinalizeVerificationKeyURLs,
		Hash:       mustArtifactHash(config.FinalizeVerificationKeyHash),
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
