package config

import "github.com/vocdoni/davinci-dkg/crypto/hash"

// TranscriptDomains collects the protocol-wide domain separators.
type TranscriptDomains struct {
	ShareEncryption []byte
	PartialDecrypt  []byte
	RoundSelection  []byte
}

// DefaultTranscriptDomains returns the active domain separators used by native
// crypto helpers and circuit witness builders.
func DefaultTranscriptDomains() TranscriptDomains {
	return TranscriptDomains{
		ShareEncryption: append([]byte(nil), hash.DomainShareEncryption...),
		PartialDecrypt:  append([]byte(nil), hash.DomainPartialDecrypt...),
		RoundSelection:  append([]byte(nil), hash.DomainRoundSelection...),
	}
}
