// operator-schnorr-vectors generates Schnorr-PoK proof vectors for
// operator registration tests. Output is a Solidity constants block to
// be pasted into solidity/test/TestHelpers.t.sol.
//
// Invocation:
//
//	go run ./cmd/operator-schnorr-vectors
//
// For each well-known test address we emit:
//
//	(privKey, pubX, pubY, A_x, A_y, z)
//
// where:
//
//	pubKey = privKey · G              (BabyJubJub generator)
//	A      = w · G                    (w is a deterministic test nonce)
//	c      = Poseidon( domain_field
//	                 , uint(address)
//	                 , pubX, pubY, A_x )    [T6, 5 inputs]
//	         then    Poseidon(c_inner, A_y) [T3, 2 inputs]
//	z      = (w + c · privKey) mod L
//
// `domain_field` is `keccak256("davinci-dkg:operator-register:v1") % Q`
// (BN254 scalar field prime), matching `_operatorSchnorrChallenge` in
// `solidity/src/DKGRegistry.sol`.
package main

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

func main() {
	curve := twistededwards.GetEdwardsCurve()
	G := curve.Base
	L := curve.Order

	// Test triples: (label, address-hex, secret-int, witness-int).
	// The address values come from the existing solidity tests: address(this)
	// (defaults to 0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84 in foundry tests),
	// 0xBEEF, 0xCAFE.  We emit one row per address per round below; the test
	// can map operator → vector by address lookup.
	addresses := []struct {
		label string
		addr  string
	}{
		{"THIS", "0x7Fa9385bE102ac3EAc297483Dd6233D62b3e1496"},
		{"BEEF", "0x000000000000000000000000000000000000bEEF"},
		{"CAFE", "0x000000000000000000000000000000000000CaFe"},
		{"ALICE", "0x00000000000000000000000000000000000A11ce"},
		{"BOB", "0x0000000000000000000000000000000000000B0b"},
		{"DEAD", "0x000000000000000000000000000000000000dEaD"},
	}

	for i, a := range addresses {
		secret := big.NewInt(int64(0x1000 + 17*i))
		witness := big.NewInt(int64(0x2000 + 23*i))
		vec := emit(curve, &G, L, a.addr, secret, witness)
		fmt.Printf(
			"// %s = %s, secret=%s, witness=%s\n"+
				"OperatorVector internal constant SCHNORR_%s = OperatorVector({\n"+
				"    pubX:  %s,\n"+
				"    pubY:  %s,\n"+
				"    aX:    %s,\n"+
				"    aY:    %s,\n"+
				"    z:     %s\n"+
				"});\n\n",
			a.label, a.addr, secret, witness, a.label,
			vec.pubX, vec.pubY, vec.aX, vec.aY, vec.z,
		)
	}
}

type vector struct {
	pubX, pubY, aX, aY, z string
}

func emit(curve twistededwards.CurveParams, G *twistededwards.PointAffine, L big.Int,
	addrHex string, secret, witness *big.Int) vector {
	_ = curve
	// pubKey = secret · G
	var pub twistededwards.PointAffine
	pub.ScalarMultiplication(G, secret)
	// A = witness · G
	var A twistededwards.PointAffine
	A.ScalarMultiplication(G, witness)

	pubX := pub.X.BigInt(new(big.Int))
	pubY := pub.Y.BigInt(new(big.Int))
	aX := A.X.BigInt(new(big.Int))
	aY := A.Y.BigInt(new(big.Int))

	addr := common.HexToAddress(addrHex)

	// c = keccak256(domain || op || pubX || pubY || aX || aY) mod L
	buf := make([]byte, 0, 32+20+32*4)
	buf = append(buf, protocol.DomainOperatorRegisterV1.Bytes()...)
	buf = append(buf, addr.Bytes()...)
	buf = append(buf, padTo32(pubX)...)
	buf = append(buf, padTo32(pubY)...)
	buf = append(buf, padTo32(aX)...)
	buf = append(buf, padTo32(aY)...)
	c := new(big.Int).SetBytes(ethcrypto.Keccak256(buf))
	c.Mod(c, &L)

	// z = witness + c · secret (mod L)
	cd := new(big.Int).Mul(c, secret)
	z := new(big.Int).Add(witness, cd)
	z.Mod(z, &L)

	return vector{
		pubX: pubX.String(),
		pubY: pubY.String(),
		aX:   aX.String(),
		aY:   aY.String(),
		z:    z.String(),
	}
}

func padTo32(v *big.Int) []byte {
	b := v.Bytes()
	if len(b) > 32 {
		return b[len(b)-32:]
	}
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
}
