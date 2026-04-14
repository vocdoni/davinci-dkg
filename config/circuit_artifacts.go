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
	ContributionProvingKeyHash      = "e68c233277593a4e7020afa82a1b30043488c885cc1c22d31538fffe679ce014"
	ContributionVerificationKeyHash = "c710d032c62ec08f74defb83c8a70571aa1188018388322b8c9ef407916ba73d"

	FinalizeCircuitHash         = "689ca94187a711094ab5617d12d85f117caedc28064906338d1bc9eec4e50872"
	FinalizeProvingKeyHash      = "16a385e173d52e6e036e0579642a27824f990485de9bc938131372d3e5953ae5"
	FinalizeVerificationKeyHash = "c38a420d1319be6e6896634e168e390991649b85700d667c5b1d302d53a29c11"

	PartialDecryptCircuitHash         = "1613282f4e41c31ac10a830294bc866a25a70b132bb2ca6944717e25c08fe321"
	PartialDecryptProvingKeyHash      = "691402123aac55fbe94d85e9ace5acf3cc31b8b9663958c18aa47da12735f164"
	PartialDecryptVerificationKeyHash = "a4ee030ade29d96ba3669231102ee57ecdf681caf1ddfda36359e499676c9d95"

	DecryptCombineCircuitHash         = "bd141ea9ca9dd953346860d061ec7df13a48f2725687ab17942a08b8fd7260d8"
	DecryptCombineProvingKeyHash      = "07013ff530ec030738bb24816cd47a12e737ba0e4080987f99381f0a2f4a8c06"
	DecryptCombineVerificationKeyHash = "e3ba94b42cdb624b289ad23d01f8afc010edabd4571196b8c5cb430b651971dd"

	RevealShareCircuitHash         = "0efade6606114b84ff9a4ab6760bc7d2eb5859a9e06dd13d4f9bac1d89edab62"
	RevealShareProvingKeyHash      = "bb6c531412f7acc294c3a34a8c15cfc3c5037de12960cf333e1934c6096b7d9c"
	RevealShareVerificationKeyHash = "7c1ee9b9fa95034f0063ba49ca0540f93ddefe1da467e602975f05fab08719aa"

	RevealSubmitCircuitHash         = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash      = "035420a4f103d93c1d6f8405f9c9dcffb0dece80bac4699c670a7745f0ccde9d"
	RevealSubmitVerificationKeyHash = "7288d1e8aceb22b302290f672efa7792791bce0d5bd67ad307e31edd7faf27ba"

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
