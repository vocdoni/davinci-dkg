package contribution

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/config"
)

// Artifacts contains the contribution circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"contribution",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{
		RemoteURLs: config.ContributionCircuitURLs,
		Hash:       mustArtifactHash(config.ContributionCircuitHash),
	},
	&circuits.Artifact{
		RemoteURLs: config.ContributionProvingKeyURLs,
		Hash:       mustArtifactHash(config.ContributionProvingKeyHash),
	},
	&circuits.Artifact{
		RemoteURLs: config.ContributionVerificationKeyURLs,
		Hash:       mustArtifactHash(config.ContributionVerificationKeyHash),
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
