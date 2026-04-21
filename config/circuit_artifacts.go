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
	ContributionProvingKeyHash      = "1a4df626104a92627e8919a3f3a8e1777837aa64bd6896836c7839cd03315302"
	ContributionVerificationKeyHash = "2e7c197568d6227d4f5bdf41c050b89a55e99611b56f8dfb5498a29ce4edd20b"

	FinalizeCircuitHash         = "856e1e6d4d52362e78c8978afe92183443224198b9710eed9a36ec03d1b71035"
	FinalizeProvingKeyHash      = "1b8d1a37c4213616630f88f306700964aea396548ca06864c2968d9664319384"
	FinalizeVerificationKeyHash = "b1aac5c4fc181cd427cf781fde357406e2d897e71f88ddc0dcef7432c7185131"

	PartialDecryptCircuitHash         = "863bc4cf95ea534c14fa074c463aa55afd8a5dfad156db5cbec2faec9a3452ce"
	PartialDecryptProvingKeyHash      = "dcf6281e9f1e2e7f989ce76bc8c495b7538b2079e9dc3cfa333078f1eb0e59b1"
	PartialDecryptVerificationKeyHash = "c66684bb10be535c0ddc5b358824ebacc4220820bf9060dc25471ad3edcff2fd"

	DecryptCombineCircuitHash         = "556e91b4798a777d6b8cdb83947cf605d25e88657634c6a75dc9917e1baa5117"
	DecryptCombineProvingKeyHash      = "372e21cbb0782715e3f90e1d7f4b719549afe862079d87fc8004a43fdbadf2b2"
	DecryptCombineVerificationKeyHash = "1e9773d92a10e13d064bd44930920923cecb58a2e1a4cfccf735cfd726a83b6c"

	RevealShareCircuitHash         = "b942794d9d90da50c24b57eb6c2590f4e6ba92ebb359b9e5482d1daa817d8849"
	RevealShareProvingKeyHash      = "0b4aeed5e5188ebf298b2cb018fd3df2213f92fddda334557ac36957590f0755"
	RevealShareVerificationKeyHash = "fadabadeba9ca793dfb4d51adf6a0b9ccecaf604c2c4897856436b3b59ad817f"

	RevealSubmitCircuitHash         = "e8a000a5e8d40715d2c9364ab8708cb81c83171e95612a1c9e9e7dce8f0b657f"
	RevealSubmitProvingKeyHash      = "56288ec5b25ac49e83b527dded1a7835b5d485f5859197b31321da08011908e7"
	RevealSubmitVerificationKeyHash = "c0ab2b039010e9ddac118c3e28452e1bb70782565d9d390f8dcc9363a28405e9"

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
