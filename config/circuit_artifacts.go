package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "11e527d3b9f4028bfeb6390ff2657cc11f9df9f759c7d88326509bbf5cdbdec4"
	ContributionProvingKeyHash      = "7876d261fb70a87bf4380dd6d9c8ce38c07478f90adedd7c7aa2cec961e708c1"
	ContributionVerificationKeyHash = "3f0c884b6100bdf3a1f57abe02c4a0c40948c97c7144f1888d299bf2357af13f"

	FinalizeCircuitHash         = "856e1e6d4d52362e78c8978afe92183443224198b9710eed9a36ec03d1b71035"
	FinalizeProvingKeyHash      = "6b559b363dd8b38ff2b84d409ce6b6dc1354714df0d64f7ae54c3af2d9f4ec67"
	FinalizeVerificationKeyHash = "3c25332a2cb9d4235e78be3c7a6b0b9aebee21609f08d2bf2bb0516b8eac9653"

	PartialDecryptCircuitHash         = "863bc4cf95ea534c14fa074c463aa55afd8a5dfad156db5cbec2faec9a3452ce"
	PartialDecryptProvingKeyHash      = "32fd77c360964dbc3723d9ac8c06fbd483e21f5db51eda83b57f5a3d08f6245c"
	PartialDecryptVerificationKeyHash = "b92904744850def22e126d049e36af56d977f0dc554b6d861b3d9da892b7810f"

	DecryptCombineCircuitHash         = "556e91b4798a777d6b8cdb83947cf605d25e88657634c6a75dc9917e1baa5117"
	DecryptCombineProvingKeyHash      = "67f87142268b13196ad33ce0e18a2d496ccf0abb8b2fa3f8773bea8d68a5f032"
	DecryptCombineVerificationKeyHash = "8cb823cc7d6e6d668e748e6e70fa238a8e34354428170e7566873e8a3a0752a0"

	RevealShareCircuitHash         = "b942794d9d90da50c24b57eb6c2590f4e6ba92ebb359b9e5482d1daa817d8849"
	RevealShareProvingKeyHash      = "dba7d6b1df657673d3458de8c6844ff71d0464032a8cb3326533139f6006f581"
	RevealShareVerificationKeyHash = "06110e5ebc8c4b21083d9319efa52385778aaeea449440c7f34cb4c59412ae23"

	RevealSubmitCircuitHash         = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash      = "9cb5cf0083b6d8ddc40cb3fa07af2bf9bf07f4dd6d7e654c7688be8a2780e615"
	RevealSubmitVerificationKeyHash = "063ca790909b2d7f27573b3a63476ea3b9ba8dcd54bba20df98c0c84e4adcbe2"

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

	RevealShareCircuitURL         = artifactURL(RevealShareCircuitHash, "ccs")
	RevealShareProvingKeyURL      = artifactURL(RevealShareProvingKeyHash, "pk")
	RevealShareVerificationKeyURL = artifactURL(RevealShareVerificationKeyHash, "vk")

	RevealSubmitCircuitURL         = artifactURL(RevealSubmitCircuitHash, "ccs")
	RevealSubmitProvingKeyURL      = artifactURL(RevealSubmitProvingKeyHash, "pk")
	RevealSubmitVerificationKeyURL = artifactURL(RevealSubmitVerificationKeyHash, "vk")
)

func artifactURL(hash, ext string) string {
	if hash == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s.%s", DefaultArtifactsBaseURL, DefaultArtifactsRelease, hash, ext)
}
