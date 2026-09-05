package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is where nodes fetch the pinned circuit artifacts:
	// the assets of a GitHub release of this repository, named `<sha256>.<ccs|pk|vk>`
	// (see `make circuits-release`). DefaultArtifactsRelease is that release's tag.
	DefaultArtifactsBaseURL = "https://github.com/vocdoni/davinci-dkg/releases/download"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "circuits-v4"
)

var (
	ContributionCircuitHash         = "ddc7238804648a9dcc58ad23357c2fbe7e3a677f5f1ae7f95d63be8ae7d26575"
	ContributionProvingKeyHash      = "bff0af27903c29e2e759e5b9653430e743cd8d56d24a2a4556fad311ec650a6a"
	ContributionVerificationKeyHash = "4433bd34f22742f6b76840a59c47816b9180fbfd3a040967cdca6a075bb887b9"

	FinalizeCircuitHash         = "70da88c6564e25f618c023f02562e42de2198c196b4c977bc4b0254f1cd4cb23"
	FinalizeProvingKeyHash      = "786fddb4f3190c1807ceb80865f0575c7f2f5ecf4ac1db52c890a37d280cea96"
	FinalizeVerificationKeyHash = "bbeb7769feab6345c236dedf9ac3ba2a39a4b803456797543925bbf7d658862b"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "7d7e54403dbde05297c09c40f1def8febaa38412d85248358aa4ab8b3d6dab00"
	PartialDecryptProvingKeyHash      = "db96fe2f3ef7ecd40c137f819812bd892d6a29f4046dd7337524571dec8c1d9d"
	PartialDecryptVerificationKeyHash = "6ed63ff4fa68e444811c14f0394b9b4c0ac3b9f9d2bfd08165fc61e79f743c7a"

	DecryptCombineCircuitHash         = "b1c67a2a1dbfab2ee84d451188e4faf26a9d02684ff4ca5c39557b3ff86db636"
	DecryptCombineProvingKeyHash      = "580a117c037fe7ca3465d9e87421e09eec03f7f9dc5d4965369bedd4ec6fad10"
	DecryptCombineVerificationKeyHash = "c493cfa7b8020299317c28f102d87969adc7cc6d2a698fbd72338cc3913c1aba"

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
