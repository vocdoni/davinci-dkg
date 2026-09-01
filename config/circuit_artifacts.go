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
	ContributionProvingKeyHash      = "3b7f6cd34cff4e0a52cb97665f76a4c735cae0719e10bf2e5795a563ed47235e"
	ContributionVerificationKeyHash = "7330e4214e95ce91fcd9c7f11af1f34027b36705102c2e9c9fd85522f0410c50"

	FinalizeCircuitHash         = "b0a1953ede1f51ed6b6388f534035cbeab4ae1203835de9d06ce52bf46ad2b18"
	FinalizeProvingKeyHash      = "1c6f6eefac1e284eefbf48a54a6836ac287f4bce7b50bb165b49491b39965916"
	FinalizeVerificationKeyHash = "a96bac2cdf79cd32c521b6e6209f7fe0596b0be6175b5b58497315eb59ba9c7a"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "65ff22ca57e9288903060eacf961aeec186d097ec16d5f76f12afcece2323309"
	PartialDecryptProvingKeyHash      = "132b2e301c3d3b4c8914546c41dd20329d6fe6bab58a96faeffa5d6e48fec5d0"
	PartialDecryptVerificationKeyHash = "a966fa12ae6dab3b141520731dc91e13c5bfd5653e4b097f4e278548314db0e3"

	DecryptCombineCircuitHash         = "288f5004a4b74c80771f0658485c2f96f5c447395b08772df138dc9bde4bf30a"
	DecryptCombineProvingKeyHash      = "192977b7b1dd4f4acfe86cb9c81cefa38856e9500b2f2168742e609ab3f9bff7"
	DecryptCombineVerificationKeyHash = "9479a671b28382a97d95e04f739e9bd078eba8cb65fb091322b6ebe61ca6ac0f"

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
