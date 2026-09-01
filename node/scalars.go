package node

import (
	"crypto/rand"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// randomScalars draws k uniform non-zero elements of the BabyJubJub scalar
// field from crypto/rand. Used for polynomial coefficients, share-encryption
// nonces and DLEQ witnesses; none of them may ever be predictable.
func randomScalars(k int) ([]*big.Int, error) {
	order := group.ScalarField()
	out := make([]*big.Int, k)
	for i := range out {
		for {
			s, err := rand.Int(rand.Reader, order)
			if err != nil {
				return nil, err
			}
			if s.Sign() != 0 {
				out[i] = s
				break
			}
		}
	}
	return out, nil
}
