package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is where nodes fetch the pinned circuit artifacts:
	// the assets of a GitHub release of this repository, named `<sha256>.<ccs|pk|vk>`
	// (see `make circuits-release`). DefaultArtifactsRelease is that release's tag.
	DefaultArtifactsBaseURL = "https://github.com/vocdoni/davinci-dkg/releases/download"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "circuits-v6"
)

var (
	ContributionCircuitHash         = "aec533457d3d4ca89b05c7a53791e66b310a6431382e170887f040ad49011059"
	ContributionProvingKeyHash      = "c5970bb317278755591e5ef36ffe9f104ba6fb53d6531465584378c4528d33ba"
	ContributionVerificationKeyHash = "70d3cfebceb198a2176dad7e6de786348af3aa6110fa86c6eb6ca4401f5b55ba"

	FinalizeCircuitHash         = "95916b110aa591f8cd13fc3c96ea6a22be6e4158919a8710f4066864d1d7af59"
	FinalizeProvingKeyHash      = "b359f97f8fef92093004dd4bf279d126e86b7c8d54b4b34fb6c714357abce759"
	FinalizeVerificationKeyHash = "8002a4eea9c83ef6e93bb3d7e0a0b93a8ae95681eea1f9a39efeabe617b20054"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "532b6c2746e51a7b334bcdbc066a83e54c9ea3b81b46b64a4582540c1263c586"
	PartialDecryptProvingKeyHash      = "d80bfa3d4d43e86204180d8884a3b1bc5c60b5f5832974f3867d11eafb22f865"
	PartialDecryptVerificationKeyHash = "fffa38ca38523a5165d94d57cd9189f4701da852e74a9476b51eab20a395c649"

	DecryptCombineCircuitHash         = "c58a9ff79c8d84ac06f9d96764dc90c241ca892b4d5376812c87321eaf3373e7"
	DecryptCombineProvingKeyHash      = "0a82f57b9a605ea4fb7f72d305887e7b441e405ef42893eb16c0dfa462d93785"
	DecryptCombineVerificationKeyHash = "befaa56b10e1fe2d076cf948ecb3fdc8f1d0fa0dece4498a4af9b281b2bc5c68"

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
