package partialdecrypt

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	dleq "github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// PublicInputs is the native representation of the partial decrypt public inputs.
type PublicInputs struct {
	RoundHash        *big.Int
	ParticipantIndex *big.Int
	Base             types.CurvePoint
	PublicKey        types.CurvePoint
	Delta            types.CurvePoint
	A1               types.CurvePoint
	A2               types.CurvePoint
	Response         *big.Int
}

// BuildWitness materializes the partial decrypt native assignment.
func BuildWitness(a Assignment) (*PartialDecryptCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	participantIndex := big.NewInt(int64(a.ParticipantIndex))
	order := group.ScalarField()
	secret := new(big.Int).Mod(new(big.Int).Set(a.Secret), order)
	nonce := new(big.Int).Mod(new(big.Int).Set(a.Nonce), order)

	basePoint, err := group.Decode(a.Base)
	if err != nil {
		return nil, nil, fmt.Errorf("decode base point: %w", err)
	}
	publicKeyPoint := group.NewPoint()
	publicKeyPoint.ScalarBaseMult(secret)
	deltaPoint := group.NewPoint()
	deltaPoint.ScalarMult(basePoint, secret)
	a1Point := group.NewPoint()
	a1Point.ScalarBaseMult(nonce)
	a2Point := group.NewPoint()
	a2Point.ScalarMult(basePoint, nonce)

	// Build the proof skeleton from the caller-provided nonce-derived
	// commitments. The circuit binds (RoundHash, ParticipantIndex) into the
	// Fiat-Shamir transcript so the native challenge derivation must match.
	proof := &dleq.Proof{
		A1: group.Encode(a1Point),
		A2: group.Encode(a2Point),
	}
	challengeState, err := ccommon.HashFieldElementsNative(
		ccommon.PartialDecryptDomain(),
		a.RoundHash,
		participantIndex,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("hash challenge prefix: %w", err)
	}
	challenge, err := ccommon.HashPointTupleNative(
		challengeState,
		group.Encode(publicKeyPoint),
		group.Encode(basePoint),
		group.Encode(deltaPoint),
		proof.A1,
		proof.A2,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("hash challenge: %w", err)
	}
	response := new(big.Int).Mul(challenge, secret)
	response.Add(response, nonce)
	response.Mod(response, order)

	witness := &PartialDecryptCircuit{
		RoundHash:        new(big.Int).Set(a.RoundHash),
		ParticipantIndex: participantIndex,
		Base:             ccommon.CircuitPoint(a.Base),
		PublicKey:        ccommon.CircuitPoint(group.Encode(publicKeyPoint)),
		Delta:            ccommon.CircuitPoint(group.Encode(deltaPoint)),
		A1:               ccommon.CircuitPoint(proof.A1),
		A2:               ccommon.CircuitPoint(proof.A2),
		Response:         response,
		Secret:           secret,
		Nonce:            nonce,
	}
	publicInputs := &PublicInputs{
		RoundHash:        new(big.Int).Set(a.RoundHash),
		ParticipantIndex: new(big.Int).Set(participantIndex),
		Base:             a.Base,
		PublicKey:        group.Encode(publicKeyPoint),
		Delta:            group.Encode(deltaPoint),
		A1:               proof.A1,
		A2:               proof.A2,
		Response:         new(big.Int).Set(response),
	}
	return witness, publicInputs, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *PartialDecryptCircuit {
	return &PartialDecryptCircuit{
		RoundHash:        p.RoundHash,
		ParticipantIndex: p.ParticipantIndex,
		Base:             ccommon.CircuitPoint(p.Base),
		PublicKey:        ccommon.CircuitPoint(p.PublicKey),
		Delta:            ccommon.CircuitPoint(p.Delta),
		A1:               ccommon.CircuitPoint(p.A1),
		A2:               ccommon.CircuitPoint(p.A2),
		Response:         p.Response,
	}
}
