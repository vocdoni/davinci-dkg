package types

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AppIDSaltBits is the width of the salt in an application id. An id is
// `salt << 160 | registrant`: `DKGAppManager.registerApplication` only accepts
// an id whose low 160 bits are the caller, so nobody can register an id in
// another account's namespace. A salt below 2^92 keeps every id below 2^252,
// inside the BN254 scalar field the decryption proofs bind it in.
const AppIDSaltBits = 92

// ErrAppIDSalt reports a salt outside [0, 2^AppIDSaltBits).
var ErrAppIDSalt = errors.New("application id salt must be in [0, 2^92)")

var appIDSaltBound = new(big.Int).Lsh(big.NewInt(1), AppIDSaltBits)

// ApplicationID returns the application id `salt << 160 | registrant`, the
// only shape the contract accepts from `registrant`.
func ApplicationID(registrant common.Address, salt *big.Int) ([32]byte, error) {
	var aid [32]byte
	if salt == nil || salt.Sign() < 0 || salt.Cmp(appIDSaltBound) >= 0 {
		return aid, ErrAppIDSalt
	}
	salt.FillBytes(aid[:common.HashLength-common.AddressLength])
	copy(aid[common.HashLength-common.AddressLength:], registrant.Bytes())
	return aid, nil
}

// RandomApplicationID returns a fresh application id in registrant's
// namespace, with a uniform 92-bit salt.
func RandomApplicationID(registrant common.Address) ([32]byte, error) {
	salt, err := rand.Int(rand.Reader, appIDSaltBound)
	if err != nil {
		return [32]byte{}, err
	}
	return ApplicationID(registrant, salt)
}

// ApplicationRegistrant returns the only account that may register aid: its
// low 160 bits.
func ApplicationRegistrant(aid [32]byte) common.Address {
	return common.BytesToAddress(aid[common.HashLength-common.AddressLength:])
}
