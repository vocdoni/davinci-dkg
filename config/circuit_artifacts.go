package config

import "fmt"

const (
	// DefaultArtifactsRelease is the circuit artifact release the pinned
	// hashes below belong to: a GitHub release of this repository whose
	// assets are named `<sha256>.<ccs|pk|vk>` (see `make circuits-release`),
	// mirrored file for file on the CDN.
	DefaultArtifactsRelease = "circuits-v6"
	// ArtifactsCDNBaseURL serves `<release>/<sha256>.<ext>` from the DAVINCI
	// assets CDN. It is the first source nodes try.
	ArtifactsCDNBaseURL = "https://davinci-assets.fra1.cdn.digitaloceanspaces.com/dkg"
	// ArtifactsGitHubBaseURL serves the same files as GitHub release assets,
	// the fallback when the CDN fails.
	ArtifactsGitHubBaseURL = "https://github.com/vocdoni/davinci-dkg/releases/download"
)

// DefaultArtifactsMirrors are the base URLs a missing artifact is downloaded
// from, in order. Every copy is stream-verified against its pinned SHA-256
// before it is cached, so a mirror can make a download fail but never feed a
// node a different file.
var DefaultArtifactsMirrors = []string{ArtifactsCDNBaseURL, ArtifactsGitHubBaseURL}

var (
	ContributionCircuitHash         = "aec533457d3d4ca89b05c7a53791e66b310a6431382e170887f040ad49011059"
	ContributionProvingKeyHash      = "c5970bb317278755591e5ef36ffe9f104ba6fb53d6531465584378c4528d33ba"
	ContributionVerificationKeyHash = "70d3cfebceb198a2176dad7e6de786348af3aa6110fa86c6eb6ca4401f5b55ba"

	FinalizeCircuitHash         = "95916b110aa591f8cd13fc3c96ea6a22be6e4158919a8710f4066864d1d7af59"
	FinalizeProvingKeyHash      = "b359f97f8fef92093004dd4bf279d126e86b7c8d54b4b34fb6c714357abce759"
	FinalizeVerificationKeyHash = "8002a4eea9c83ef6e93bb3d7e0a0b93a8ae95681eea1f9a39efeabe617b20054"

	PartialDecryptCircuitHash         = "532b6c2746e51a7b334bcdbc066a83e54c9ea3b81b46b64a4582540c1263c586"
	PartialDecryptProvingKeyHash      = "d80bfa3d4d43e86204180d8884a3b1bc5c60b5f5832974f3867d11eafb22f865"
	PartialDecryptVerificationKeyHash = "fffa38ca38523a5165d94d57cd9189f4701da852e74a9476b51eab20a395c649"

	DecryptCombineCircuitHash         = "c58a9ff79c8d84ac06f9d96764dc90c241ca892b4d5376812c87321eaf3373e7"
	DecryptCombineProvingKeyHash      = "0a82f57b9a605ea4fb7f72d305887e7b441e405ef42893eb16c0dfa462d93785"
	DecryptCombineVerificationKeyHash = "befaa56b10e1fe2d076cf948ecb3fdc8f1d0fa0dece4498a4af9b281b2bc5c68"

	ContributionCircuitURLs         = artifactURLs(ContributionCircuitHash, "ccs")
	ContributionProvingKeyURLs      = artifactURLs(ContributionProvingKeyHash, "pk")
	ContributionVerificationKeyURLs = artifactURLs(ContributionVerificationKeyHash, "vk")

	FinalizeCircuitURLs         = artifactURLs(FinalizeCircuitHash, "ccs")
	FinalizeProvingKeyURLs      = artifactURLs(FinalizeProvingKeyHash, "pk")
	FinalizeVerificationKeyURLs = artifactURLs(FinalizeVerificationKeyHash, "vk")

	PartialDecryptCircuitURLs         = artifactURLs(PartialDecryptCircuitHash, "ccs")
	PartialDecryptProvingKeyURLs      = artifactURLs(PartialDecryptProvingKeyHash, "pk")
	PartialDecryptVerificationKeyURLs = artifactURLs(PartialDecryptVerificationKeyHash, "vk")

	DecryptCombineCircuitURLs         = artifactURLs(DecryptCombineCircuitHash, "ccs")
	DecryptCombineProvingKeyURLs      = artifactURLs(DecryptCombineProvingKeyHash, "pk")
	DecryptCombineVerificationKeyURLs = artifactURLs(DecryptCombineVerificationKeyHash, "vk")
)

// artifactURLs lists where the artifact with the given hash and extension is
// published, one URL per mirror in DefaultArtifactsMirrors order.
func artifactURLs(hash, ext string) []string {
	if hash == "" {
		return nil
	}
	urls := make([]string, 0, len(DefaultArtifactsMirrors))
	for _, base := range DefaultArtifactsMirrors {
		urls = append(urls, fmt.Sprintf("%s/%s/%s.%s", base, DefaultArtifactsRelease, hash, ext))
	}
	return urls
}
