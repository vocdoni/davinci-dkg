package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "4dde6c9596cc0f48d88a049e06890d5a4fc5e5478515913932d69fbf562eba4a"
	ContributionProvingKeyHash      = "db44ff21a27995df43d8bf5123474b199a2c3851b64a2f2b039b828b6a5b9f47"
	ContributionVerificationKeyHash = "2373ff8dd55a040b0a4b3a90f79257816a4f786713aed69779e47e40d709bb90"

	FinalizeCircuitHash         = "b0a1953ede1f51ed6b6388f534035cbeab4ae1203835de9d06ce52bf46ad2b18"
	FinalizeProvingKeyHash      = "0cf61bd16ef3d481b31b40d655f1e901429f88911b1abdc675cf23dadcf86a09"
	FinalizeVerificationKeyHash = "af23767346af55a9d2968586484208f57d861e928cbc530d691f4e939def9023"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "65ff22ca57e9288903060eacf961aeec186d097ec16d5f76f12afcece2323309"
	PartialDecryptProvingKeyHash      = "fabc38153fad944fbc64c51867e8df0f0c3e8a73721f8ce449a0c99134129d64"
	PartialDecryptVerificationKeyHash = "326ff6651a27c9e3d101765f376df08933f6515eb0cfbd22c2155ff7caa11fea"

	DecryptCombineCircuitHash         = "ebd6543048720dbaad1d8a5e8734c37bbfd4492bcdb91ec25479fc9a1893de32"
	DecryptCombineProvingKeyHash      = "e16f45b822e85899f3d9f1972077f428531cb04e885c653d0716463d15433cf2"
	DecryptCombineVerificationKeyHash = "0182bb965daec8b180a310cd7f6115a60d44083fad1832c1f45f61985a7da68a"

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
