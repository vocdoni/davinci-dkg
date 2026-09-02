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
	ContributionProvingKeyHash      = "d747a935a6680c8b6446389ab994a7d939aa12b0def66c921f85fc984b1e69d9"
	ContributionVerificationKeyHash = "2d6a9bc67b0c7ec7c574d3f6d9154b8c5bee5251c5269ca09b92be14ac18ab75"

	FinalizeCircuitHash         = "1b1c150d21c8e02790cce49956c4d26d5382f21033e211904c9392f90bae7257"
	FinalizeProvingKeyHash      = "d1bffaabda7c70f45f45812593b9bd4ba3a599d1e5631d1659d5dafc3d7aa142"
	FinalizeVerificationKeyHash = "353ba8af766f1183d80a1b288d91f0cde3dbef5dd7d3f047f5f9ede1a855d804"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "bd17eca641d936ae3ef955873a3e25f9a6b6f206191ff93137de3ab81d11604d"
	PartialDecryptProvingKeyHash      = "9c536a045045acc6d2bc25066b8af160f7b2cc9d6affb3f1a5fe6a8b4299b1a1"
	PartialDecryptVerificationKeyHash = "9c63d60bdfd93d52d8a931187e8b6b8e29e1870560f78de3ed0671dbe059bbdf"

	DecryptCombineCircuitHash         = "21145a18917376d585ad9ab23e7917537221d8970179718ed7d166b96a333b07"
	DecryptCombineProvingKeyHash      = "23b50690255b6580e58c4f76addf9359f378856e884fec3bd3cc1e2c2960ecb3"
	DecryptCombineVerificationKeyHash = "06d704430edde466bc21c22916a0c45a31c1ed3da7841e7882e5678612a7676d"

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
