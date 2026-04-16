package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "a83d50dc880614ce34b1b874eea8961bbbea6dbddb4b1634ca2adc568eaa12cf"
	ContributionProvingKeyHash      = "0dc0a8dafd381ea3c8130cd45e5d709f4451dc2c0ecb8d3a8409c149f24830df"
	ContributionVerificationKeyHash = "eef1e0a1f5f98374a2495faad60052ff1d3cfc5b0b68ecbe972f86aac17fa484"

	FinalizeCircuitHash         = "856e1e6d4d52362e78c8978afe92183443224198b9710eed9a36ec03d1b71035"
	FinalizeProvingKeyHash      = "59658e24dbee2cf345ab0a9b261f94530692858063b8d41f8409f30f35c1c6a2"
	FinalizeVerificationKeyHash = "2f0897f87d62ea93bd01010d995860bf4b83bdd5323daff6baf90b517be4203b"

	PartialDecryptCircuitHash         = "1613282f4e41c31ac10a830294bc866a25a70b132bb2ca6944717e25c08fe321"
	PartialDecryptProvingKeyHash      = "b7abf78675ebd66835387f8bdff10afbb2c5a489fa25b0f9ffc58e88b25e4938"
	PartialDecryptVerificationKeyHash = "5a0fd207c45eaf90e70e0215ba598041ad210876a9582212c933b47ce550bfec"

	DecryptCombineCircuitHash         = "556e91b4798a777d6b8cdb83947cf605d25e88657634c6a75dc9917e1baa5117"
	DecryptCombineProvingKeyHash      = "f2ec57eb51ebff3dea99e574ca09f572cfeefb0695be372e8e1213124c4f1cd9"
	DecryptCombineVerificationKeyHash = "b4fd6ae9a1b017179e66994bc326d530a2d0bea3737e94b2dba2ba81471b99ec"

	RevealShareCircuitHash         = "b942794d9d90da50c24b57eb6c2590f4e6ba92ebb359b9e5482d1daa817d8849"
	RevealShareProvingKeyHash      = "b17cae649172429ab7bd11b08d78a8db7f8141ebacd90a4f90106c1e4a03ee3c"
	RevealShareVerificationKeyHash = "478cf42495c1b6de90efcbb2b2fd17a9a6ee60c86ac2d7a8d4eae3bb0aadf90f"

	RevealSubmitCircuitHash         = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash      = "4d5514e3272098a990ed7edf8494fb96ab783d5c88e5b63755e72bdeeadbbfce"
	RevealSubmitVerificationKeyHash = "b73bd21a31ea80d1e3ce3a0d97a130b63dc44363291de245dbb8f6f41e4a5575"

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
