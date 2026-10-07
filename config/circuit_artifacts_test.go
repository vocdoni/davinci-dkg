package config

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

// TestArtifactURLsCDNFirst pins the download order: the CDN, then the GitHub
// release, both with the release's `<sha256>.<ext>` asset names.
func TestArtifactURLsCDNFirst(t *testing.T) {
	c := qt.New(t)
	c.Assert(FinalizeProvingKeyURLs, qt.DeepEquals, []string{
		"https://davinci-assets.fra1.cdn.digitaloceanspaces.com/dkg/" + DefaultArtifactsRelease + "/" +
			FinalizeProvingKeyHash + ".pk",
		"https://github.com/vocdoni/davinci-dkg/releases/download/" + DefaultArtifactsRelease + "/" +
			FinalizeProvingKeyHash + ".pk",
	})
	for _, urls := range [][]string{
		ContributionCircuitURLs, ContributionProvingKeyURLs, ContributionVerificationKeyURLs,
		FinalizeCircuitURLs, FinalizeProvingKeyURLs, FinalizeVerificationKeyURLs,
		PartialDecryptCircuitURLs, PartialDecryptProvingKeyURLs, PartialDecryptVerificationKeyURLs,
		DecryptCombineCircuitURLs, DecryptCombineProvingKeyURLs, DecryptCombineVerificationKeyURLs,
	} {
		c.Assert(urls, qt.HasLen, len(DefaultArtifactsMirrors))
	}
}
