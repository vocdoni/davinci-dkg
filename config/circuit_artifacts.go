package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is where nodes fetch the pinned circuit artifacts:
	// the assets of a GitHub release of this repository, named `<sha256>.<ccs|pk|vk>`
	// (see `make circuits-release`). DefaultArtifactsRelease is that release's tag.
	DefaultArtifactsBaseURL = "https://github.com/vocdoni/davinci-dkg/releases/download"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "circuits-v5"
)

var (
	ContributionCircuitHash         = "67fc3d6c3e9abac0d019c2d0a27b15b9b70993ef699aae7935d7388eeab9b002"
	ContributionProvingKeyHash      = "c7fa389cb59d476c6d89df8f1b45fb4094afefed299f0e486d0f5dbf93288345"
	ContributionVerificationKeyHash = "c770faf0c5fe0700fc28601bd510394d3264aead45ceac44210049f90ba10277"

	FinalizeCircuitHash         = "30368e00297a19e07141802c50124527243340622ce11b529808d4154c5fff02"
	FinalizeProvingKeyHash      = "cb4132e18b0a76184a7904a926546cf2fd02fa27fac833383b28796db10429c6"
	FinalizeVerificationKeyHash = "65391b29e5f52b78931351bec8ad96442d5c915dd7808ede844ab68de5073c8a"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "503eca4896a825db777f9bb4e1beae8d27ede5cb923aaa4702ef451ec591ce8a"
	PartialDecryptProvingKeyHash      = "0e827380dafca1282092a4afd188d67cbdaced28146a590e77a992415b05e94f"
	PartialDecryptVerificationKeyHash = "ebfd45b83aec1c1e827db2daf81fc0aff902d309923b0bd8ff41a97c2dfac73f"

	DecryptCombineCircuitHash         = "c3d375b477c9ac80b27370ba6db64a00841372d50177d004d6fb504562db9243"
	DecryptCombineProvingKeyHash      = "660a642c9798c925b8d3c838995e217089c48ef7bb4e26fa090c7b4bd805199a"
	DecryptCombineVerificationKeyHash = "d2c499fc551c58d806fb3dc6fbf7dfb726e056037a992c6311f673092cf89a9b"

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
