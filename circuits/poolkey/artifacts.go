package poolkey

import (
	"encoding/hex"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/config"
)

// Artifacts contains the pool-key circuit artifact configuration.
var Artifacts = circuits.NewCircuitArtifacts(
	"poolkey",
	ecc.BN254,
	[]backend.ProverOption{solidity.WithProverTargetSolidityVerifier(backend.GROTH16)},
	[]backend.VerifierOption{solidity.WithVerifierTargetSolidityVerifier(backend.GROTH16)},
	&circuits.Artifact{RemoteURL: config.PoolKeyCircuitURL, Hash: mustArtifactHash(config.PoolKeyCircuitHash)},
	&circuits.Artifact{RemoteURL: config.PoolKeyProvingKeyURL, Hash: mustArtifactHash(config.PoolKeyProvingKeyHash)},
	&circuits.Artifact{
		RemoteURL: config.PoolKeyVerificationKeyURL,
		Hash:      mustArtifactHash(config.PoolKeyVerificationKeyHash),
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
