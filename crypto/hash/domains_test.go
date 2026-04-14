package hash

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTranscriptDomainsAreStable(t *testing.T) {
	c := qt.New(t)

	c.Assert(DomainShareEncryption, qt.Not(qt.DeepEquals), []byte(nil))
	c.Assert(DomainShareEncryption, qt.Not(qt.DeepEquals), DomainPartialDecrypt)
	c.Assert(DomainPartialDecrypt, qt.Not(qt.DeepEquals), DomainRoundSelection)
}
