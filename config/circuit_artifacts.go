package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "5cb3951874d4fa5d4f41dfc6f2a7cfe83cd7e762e284541ff61c028aff99e434"
	ContributionProvingKeyHash      = "406e7e943fbf4713e45e71d369bc810d4f3bfcc2e8b84ed97997cc1c4ac8d455"
	ContributionVerificationKeyHash = "d20d19d3d292708a9c77eb58772b918eb918a77fa73452d8c200ed4a869c297f"

	FinalizeCircuitHash         = "1b1c150d21c8e02790cce49956c4d26d5382f21033e211904c9392f90bae7257"
	FinalizeProvingKeyHash      = "252e7ff136ac4d02e9caac4aff0d8cc88d9b0d51aa53ee80cf7d857f52dfa8da"
	FinalizeVerificationKeyHash = "9a818d8bb2d01e484ebde3ca847ca6f69852a3e7b6d9fe129aec724db7439e17"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "1cf4d0748da1d61ec4c19bd4b10529b19aa385f1510deb53aa3b5a6d1041f7ad"
	PartialDecryptProvingKeyHash      = "8d5cb7d575354215d07b8b2721e1a5465424a85f78ca2d0beb56041efd2b0820"
	PartialDecryptVerificationKeyHash = "32bf56fd36212c9e52b0662c803e65a84236483b8497fad41eb6e1b60bc3ddf3"

	DecryptCombineCircuitHash         = "1849fae426b116e9fe9686b61c45f9294da7b667f4e9f091f4706d878a350e23"
	DecryptCombineProvingKeyHash      = "45db382e20ec94816f3fd8bb0fd1c54a5a550eb6a842c26790fac5d5583a442f"
	DecryptCombineVerificationKeyHash = "542df738d839dd71cb36c00389b8bf8db2916912f830a48615fa38b62221f686"

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
