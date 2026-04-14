package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is the remote artifact bucket base.
	DefaultArtifactsBaseURL = "https://circuits.ams3.cdn.digitaloceanspaces.com"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "dev"
)

var (
	ContributionCircuitHash = "ce7ffd51cb46f13609b3cba3e1a765727000e48f2c3f4d51d771a2fb28ad6267"
	ContributionProvingKeyHash = "bded8ac56464b5739a2911cec81f0a64676e0fe113b7166f02db1bb7329a72f2"
	ContributionVerificationKeyHash = "e5c7885398480d1ffa4acb9e153cac604e5f952691c4a301598425a4d03c62ed"

	FinalizeCircuitHash = "689ca94187a711094ab5617d12d85f117caedc28064906338d1bc9eec4e50872"
	FinalizeProvingKeyHash = "86311dd591007244f53d866bae7b56e77d4ffcdf6a6dc510aa8350a367dd4e5b"
	FinalizeVerificationKeyHash = "cfb859fd24f2972fdfb20d40d8d1d677360070a91d33b52a6b1172c5c1c252df"

	PartialDecryptCircuitHash = "1613282f4e41c31ac10a830294bc866a25a70b132bb2ca6944717e25c08fe321"
	PartialDecryptProvingKeyHash = "7f580c4a71f93e660876921561cd109f28fe07006ac2ecb868ab41ce6a7eed8b"
	PartialDecryptVerificationKeyHash = "b68e47d0c45693a7294f0b6f4c0a57694ade6b19d92fc6a5a7118f5da9c7bca8"

	DecryptCombineCircuitHash = "bd141ea9ca9dd953346860d061ec7df13a48f2725687ab17942a08b8fd7260d8"
	DecryptCombineProvingKeyHash = "2f36d3c3fa4e82cc48a1994ba8fdb330a1a683a3718e6775a88c255276eb916f"
	DecryptCombineVerificationKeyHash = "da563a4ad86f4039ef74b76584944427e6dfb8c3feb483d465221e4d16d1947c"

	RevealShareCircuitHash = "0efade6606114b84ff9a4ab6760bc7d2eb5859a9e06dd13d4f9bac1d89edab62"
	RevealShareProvingKeyHash = "d4c2ebd15d855dc6f542106c2ee48f2dca599856f9e16f997bde18482ecf9abd"
	RevealShareVerificationKeyHash = "541c6029c7633dc0206b7bdc44f062ee62e59979ec308cdcf3a0c05d9942e3ff"

	RevealSubmitCircuitHash = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash = "a844d36ae9a2c979c21dc7e81cb842fccee3272f67affec5f381eef608501e76"
	RevealSubmitVerificationKeyHash = "38583cf450809831a71cc8cff4598c833b818a5ca613250c7111e88d651b2f92"

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
