package config

import "fmt"

const (
	// DefaultArtifactsBaseURL is where nodes fetch the pinned circuit artifacts:
	// the assets of a GitHub release of this repository, named `<sha256>.<ccs|pk|vk>`
	// (see `make circuits-release`). DefaultArtifactsRelease is that release's tag.
	DefaultArtifactsBaseURL = "https://github.com/vocdoni/davinci-dkg/releases/download"
	// DefaultArtifactsRelease is the default remote artifact release channel.
	DefaultArtifactsRelease = "circuits-v3"
)

var (
	ContributionCircuitHash         = "e204d9fe5059d10eb68577bb62f38b2497790d9f8196e03858bddbd01c6f70ad"
	ContributionProvingKeyHash      = "3ff7007dbb761e27713637a497ef763e3a159564aeac58baaccbc3610ea0a6cf"
	ContributionVerificationKeyHash = "92cbf074d24cbac27852f3447e7fde341e5af8a9c5aae2cbbeb5fc75c1fd4c5b"

	PoolKeyCircuitHash         = "d47c0a8c4cb4e773ecb461d9c48621f5a3265a9ccd7f66c67ea3be2c4d3a1698"
	PoolKeyProvingKeyHash      = "190f9007ad878e9b8d8ef68770c984d3ada5d6ad32e4945393993317cb20b8d5"
	PoolKeyVerificationKeyHash = "6fac00af882c60445ca261220ed4104ce8fe593b20372fab0582974065ad64e6"

	// Hashes regenerated in P6 after the P5 circuit changes (added Aid,
	// CtIdx, Role public inputs to partialdecrypt; added Aid, CtIdx, Mode,
	// S, DeltaOrg + mode-aware T branch to decryptcombine). These pk/vk
	// hashes correspond to a DEV trusted setup; the production ceremony
	// in S2 will regenerate fresh keys and bump these again.
	PartialDecryptCircuitHash         = "7d7e54403dbde05297c09c40f1def8febaa38412d85248358aa4ab8b3d6dab00"
	PartialDecryptProvingKeyHash      = "80205078ebcf34a4bb2e68928450e1afd388a0dfd8a0810e6fedc1ae5f855b57"
	PartialDecryptVerificationKeyHash = "250dfb18431b5afdcf2b400bda8dd71b2c57aef9782c50975b4b8a7e1d5c5d1d"

	DecryptCombineCircuitHash         = "b1c67a2a1dbfab2ee84d451188e4faf26a9d02684ff4ca5c39557b3ff86db636"
	DecryptCombineProvingKeyHash      = "d70de162a56ac4801077f857f9491015916bb238db7fb9b0cfe5061eceae5305"
	DecryptCombineVerificationKeyHash = "daf3a4c487d7e743ee1217a6213b38e64db3cbffce94342eb835dd2f4145b619"

	ContributionCircuitURL         = artifactURL(ContributionCircuitHash, "ccs")
	ContributionProvingKeyURL      = artifactURL(ContributionProvingKeyHash, "pk")
	ContributionVerificationKeyURL = artifactURL(ContributionVerificationKeyHash, "vk")

	PoolKeyCircuitURL         = artifactURL(PoolKeyCircuitHash, "ccs")
	PoolKeyProvingKeyURL      = artifactURL(PoolKeyProvingKeyHash, "pk")
	PoolKeyVerificationKeyURL = artifactURL(PoolKeyVerificationKeyHash, "vk")

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
