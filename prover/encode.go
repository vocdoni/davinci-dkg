package prover

import (
	"fmt"
	"math/big"
	"reflect"

	gnec "github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common"
)

// MarshalSolidityProof serialises a BN254 Groth16 proof in the byte layout the
// generated Solidity verifiers expect.
func MarshalSolidityProof(proof groth16.Proof) ([]byte, error) {
	p, ok := proof.(*groth16bn254.Proof)
	if !ok {
		return nil, fmt.Errorf("unexpected proof type %T", proof)
	}
	return p.MarshalSolidity(), nil
}

// EncodePublicWitness ABI-encodes the public inputs of a circuit assignment
// as consecutive 32-byte words, in declaration order.
func EncodePublicWitness(pub frontend.Circuit) ([]byte, error) {
	w, err := frontend.NewWitness(pub, gnec.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		return nil, fmt.Errorf("build public witness: %w", err)
	}
	rv := reflect.ValueOf(w.Vector())
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unexpected witness vector type %T", w.Vector())
	}
	vals := make([]*big.Int, rv.Len())
	for i := range vals {
		m := rv.Index(i).Addr().MethodByName("BigInt")
		if !m.IsValid() {
			return nil, fmt.Errorf("element %d missing BigInt", i)
		}
		v, ok := m.Call([]reflect.Value{reflect.ValueOf(new(big.Int))})[0].Interface().(*big.Int)
		if !ok {
			return nil, fmt.Errorf("element %d BigInt bad type", i)
		}
		vals[i] = new(big.Int).Set(v)
	}
	return EncodeWords(vals...)
}

// EncodeWords left-pads every value to 32 bytes and concatenates them.
func EncodeWords(values ...*big.Int) ([]byte, error) {
	out := make([]byte, 0, len(values)*32)
	for i, v := range values {
		if v == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}
		out = append(out, common.LeftPadBytes(v.Bytes(), 32)...)
	}
	return out, nil
}
