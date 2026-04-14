package hash

var (
	// DomainShareEncryption separates share-encryption transcript hashing.
	DomainShareEncryption = []byte("davinci-dkg/share-encryption/v1")
	// DomainPartialDecrypt separates partial-decryption transcript hashing.
	DomainPartialDecrypt = []byte("davinci-dkg/partial-decrypt/v1")
	// DomainRoundSelection separates committee selection hashing.
	DomainRoundSelection = []byte("davinci-dkg/round-selection/v1")
)
