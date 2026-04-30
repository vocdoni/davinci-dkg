package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash = "11e527d3b9f4028bfeb6390ff2657cc11f9df9f759c7d88326509bbf5cdbdec4"
	ContributionProvingKeyHash = "5c77f8f911a3b7933667cc03856a40146de22b1461006e333f6e623b350bb48c"
	ContributionVerificationKeyHash = "8f858f4226de14efc4a938cc66e4c113782afdfe12bbe47849cde6b3e5be12bf"

	FinalizeCircuitHash = "856e1e6d4d52362e78c8978afe92183443224198b9710eed9a36ec03d1b71035"
	FinalizeProvingKeyHash = "b7d50091bbf5d4629605846b345d0113903db148242238c55d039eb0c88e51d9"
	FinalizeVerificationKeyHash = "e0bb00a05d2a34725cfa9864d625a72bf794485bab4a13faa8dbd73a8604f22f"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash = "ea643967ac06560492d2e2a52d433288d13c59aafe1fd2ccd7f756fdf66c4268"
	PartialDecryptProvingKeyHash = "e36767c950be07ed840c6f59c9a3c3cd54f08594a256d18eceb07997a9ff5245"
	PartialDecryptVerificationKeyHash = "dddca136f8b31f2508d4e5ce237f06f8738224a0d6ceb0afd48325e00eb07388"

	DecryptCombineCircuitHash = "f7e203480ffab16f1e55f39f61cdd05782be2d7ea0033b284293b49bee5d133b"
	DecryptCombineProvingKeyHash = "d1cd294544102bb194baa6a3255de8e23899df4c1f489c4d45c23b84bb44af43"
	DecryptCombineVerificationKeyHash = "a3af90bb13983d5dee88a0dc9fd133af13fefaab4af39ac8af87581a64d1c7ec"

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
