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
	RoundHash        *big.Int // semantically: eid
	Aid              *big.Int
	CtIdx            *big.Int
	Role             *big.Int
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
	// Default the per-app transcript binding fields to zero / committee
	// when callers don't supply them. The circuit binds these into the
	// Fiat-Shamir state so the values must match between prover and verifier.
	aid := new(big.Int)
	if a.Aid != nil {
		aid.Set(a.Aid)
	}
	ctIdx := new(big.Int)
	if a.CtIdx != nil {
		ctIdx.Set(a.CtIdx)
	}
	role := big.NewInt(1) // ROLE_COMMITTEE
	if a.Role != nil {
		role.Set(a.Role)
	}
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
		aid,
		ctIdx,
		role,
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
		Aid:              new(big.Int).Set(aid),
		CtIdx:            new(big.Int).Set(ctIdx),
		Role:             new(big.Int).Set(role),
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
		Aid:              new(big.Int).Set(aid),
		CtIdx:            new(big.Int).Set(ctIdx),
		Role:             new(big.Int).Set(role),
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
		Aid:              p.Aid,
		CtIdx:            p.CtIdx,
		Role:             p.Role,
		ParticipantIndex: p.ParticipantIndex,
		Base:             ccommon.CircuitPoint(p.Base),
		PublicKey:        ccommon.CircuitPoint(p.PublicKey),
		Delta:            ccommon.CircuitPoint(p.Delta),
		A1:               ccommon.CircuitPoint(p.A1),
		A2:               ccommon.CircuitPoint(p.A2),
		Response:         p.Response,
	}
}
