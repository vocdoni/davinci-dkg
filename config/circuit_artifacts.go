package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "161da14b379de8f7f5450b8f56e418052c1a0b0a96d042a70656eb2cf6b48d63"
	ContributionProvingKeyHash      = "29317a0df682000be9e499b1e1b6b00047fd97a1120c951d189afd3d7bbda95f"
	ContributionVerificationKeyHash = "2a0599b0b58a7a2dcb275301942383a55b4235adbfeddc3dc0b4c38d01d14900"

	FinalizeCircuitHash         = "1b1c150d21c8e02790cce49956c4d26d5382f21033e211904c9392f90bae7257"
	FinalizeProvingKeyHash      = "6dc5a8de96bdc68c7100bbaf7941e66b2c607fc21738d908204dbc9ddbfada75"
	FinalizeVerificationKeyHash = "12b6b1565d3c4977a602424832924d780e66f230d1a558c106eee75b591c0879"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "1cf4d0748da1d61ec4c19bd4b10529b19aa385f1510deb53aa3b5a6d1041f7ad"
	PartialDecryptProvingKeyHash      = "74e759dfc2477a128e4d830e7161f32d02f0cc8582b825d29a115c58201b7c0e"
	PartialDecryptVerificationKeyHash = "d4807144305d06d433a1362d75ec7e9dd3acd61fd0aeb891a50e9b24b23a6334"

	DecryptCombineCircuitHash         = "1849fae426b116e9fe9686b61c45f9294da7b667f4e9f091f4706d878a350e23"
	DecryptCombineProvingKeyHash      = "d8248e4c88328c26fd2135bae946e618485207eb738a69f8d2c56b07bd631a21"
	DecryptCombineVerificationKeyHash = "d54254c027394d13c673a2be14c97a9c6bf88bd32b16ab5b2a9c55d7e0efa6c1"

	ContributionCircuitURL         = artifactURL(ContributionCircuitHash, "ccs")
	ContributionProvingKeyURL      = artifactURL(ContributionProvingKeyHash, "pk")
	ContributionVerificationKeyURL = artifactURL(ContributionVerificationKeyHash, "vk")

	FinalizeCircuitURL         = artifactURL(FinalizeCircuitHash, "ccs")
	FinalizeProvingKeyURL      = artifactURL(FinalizeProvingKeyHash, "pk")
	FinalizeVerificationKeyURL = artifactURL(FinalizeVerificationKeyHash, "vk")

	PartialDecryptCircuitURL         = artifactURL(PartialDecryptCircuitHash, "ccs")
	PartialDecryptProvingKeyURL      = artifactURL(PartialDecryptProvingKeyHash, "pk")
	PartialDecryptVerificationKeyURL = artifactURL(PartialDecryptVerificationKeyHash, "vk")

	DecryptCombineCircuitURL         = artifactURL(DecryptCombineCircuitHash, "ccs")
	DecryptCombineProvingKeyURL      = artifactURL(DecryptCombineProvingKeyHash, "pk")
	DecryptCombineVerificationKeyURL = artifactURL(DecryptCombineVerificationKeyHash, "vk")
)

func artifactURL(hash, ext string) string {
	if hash == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s.%s", DefaultArtifactsBaseURL, DefaultArtifactsRelease, hash, ext)
}
