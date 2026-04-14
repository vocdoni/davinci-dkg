package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash         = "ce7ffd51cb46f13609b3cba3e1a765727000e48f2c3f4d51d771a2fb28ad6267"
	ContributionProvingKeyHash      = "6765b46d97f506138009f8757637b4254edce95c952701f6e228f81def0aa653"
	ContributionVerificationKeyHash = "aaef6669999b1337553ca062017118652b32dce90ffd930940a430f84166b9e8"

	FinalizeCircuitHash         = "689ca94187a711094ab5617d12d85f117caedc28064906338d1bc9eec4e50872"
	FinalizeProvingKeyHash      = "5b3413e711d99042e31a41ccf73064ac22d7252e7bd1525a1104532c400674a8"
	FinalizeVerificationKeyHash = "1728eeec379a573512321259121881c721dbcea12f3bc2d80549ff1628126644"

	PartialDecryptCircuitHash         = "1613282f4e41c31ac10a830294bc866a25a70b132bb2ca6944717e25c08fe321"
	PartialDecryptProvingKeyHash      = "bd5f3f36b675fb648f29c1ae0712fa84269bcae9511c4e9922965cc47b50d6aa"
	PartialDecryptVerificationKeyHash = "527c0932488c04fac96cae578e9b6dbaf1df50a95be9d8a7ebaee04463eb5543"

	DecryptCombineCircuitHash         = "bd141ea9ca9dd953346860d061ec7df13a48f2725687ab17942a08b8fd7260d8"
	DecryptCombineProvingKeyHash      = "79cf9295e21c191eb0aed912ebee2cbec5a0eedc398ecaceba22321432730b68"
	DecryptCombineVerificationKeyHash = "aeda683cf6cd2b223bee887efab475cfdd6046b048712a470669c76d9634214d"

	RevealShareCircuitHash         = "0efade6606114b84ff9a4ab6760bc7d2eb5859a9e06dd13d4f9bac1d89edab62"
	RevealShareProvingKeyHash      = "f8cc7c169bce70ebfd98c9cb12f9b1a722f88ea083f7a94f7c1b15a8e2a16f18"
	RevealShareVerificationKeyHash = "dbeded347a4689669b88bb1f5f6ef1f27943126daa9a78569ed877a3eb0909f7"

	RevealSubmitCircuitHash         = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash      = "9b8a9f5bbc20e7747bd70532f58f95e6048b53bdbc32aa1050c27555d433e0e0"
	RevealSubmitVerificationKeyHash = "1f7e19d2d7d3fbca713d90bc134b2515a29ba69adc43b2e5c707b5facab9ec46"

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
