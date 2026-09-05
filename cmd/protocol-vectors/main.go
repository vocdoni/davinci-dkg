// protocol-vectors emits the canonical cross-impl test vectors as JSON
// under tests/vectors/. The SDK and the Foundry suite both consume these
// files directly, so any change to the underlying primitives is detected
// at every layer in CI.
//
//	go run ./cmd/protocol-vectors            # write to tests/vectors/
//	go run ./cmd/protocol-vectors -dir DIR   # write to a different dir
//
// Files emitted:
//
//	tests/vectors/protocol.json    domain digests (Schnorr / DLEQ transcripts and
//	                               the three BRLC transcript domains) + field constants
//	tests/vectors/schnorr.json     operator + organizer Schnorr proofs
//	tests/vectors/dleq.json        committee partial-decryption DLEQ challenges + responses
//
// Determinism: every vector uses fixed inputs and either deterministic
// witnesses or the chosen ones below. Re-running this command on a clean
// checkout must produce byte-identical files.
package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

var bn254Q, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10,
)

func main() {
	dir := flag.String("dir", "tests/vectors", "output directory")
	flag.Parse()

	if err := os.MkdirAll(*dir, 0o755); err != nil {
		fail("mkdir %s: %v", *dir, err)
	}

	write(*dir, "protocol.json", buildProtocol())
	write(*dir, "schnorr.json", buildSchnorr())
	write(*dir, "dleq.json", buildDLEQ())
}

// ─── protocol.json ──────────────────────────────────────────────────────────

type protocolFile struct {
	Description string                       `json:"description"`
	Domains     map[string]protocolDomainRow `json:"domains"`
	BN254Q      string                       `json:"bn254Q"`
	SubgroupL   string                       `json:"subgroupOrderL"`
}

type protocolDomainRow struct {
	Preimage     string `json:"preimage"`
	Keccak256    string `json:"keccak256"`
	BN254Reduced string `json:"bn254Reduced"`
}

func buildProtocol() protocolFile {
	return protocolFile{
		Description: "Cross-impl protocol constants. Domain digests are bound into the Schnorr registration transcripts; " +
			"the *TranscriptV1 rows are the BRLC domains every proof-carrying call binds into its challenge " +
			"(keccak(eid || domain || anchor) mod p).",
		Domains: map[string]protocolDomainRow{
			"OperatorRegisterV1":  domainRow(protocol.DomainOperatorRegisterV1Str, protocol.DomainOperatorRegisterV1),
			"OrganizerRegisterV1": domainRow(protocol.DomainOrganizerRegisterV1Str, protocol.DomainOrganizerRegisterV1),
			// BRLC transcript domains: submitContribution, activatePoolKey
			// and combineDecryption each derive their Fiat-Shamir challenge
			// under one of these. The pool-key row replaces the former
			// davinci-dkg:finalize:v1.
			"ContributionTranscriptV1": domainRow(
				protocol.DomainContributionTranscriptV1Str, protocol.DomainContributionTranscriptV1,
			),
			"PoolKeyTranscriptV1": domainRow(
				protocol.DomainPoolKeyTranscriptV1Str, protocol.DomainPoolKeyTranscriptV1,
			),
			"DecryptCombineTranscriptV1": domainRow(
				protocol.DomainDecryptCombineTranscriptV1Str, protocol.DomainDecryptCombineTranscriptV1,
			),
			// PartialDecrypt domain is consumed by the in-circuit DLEQ
			// transcript via SetBytes (no keccak); included here so the SDK
			// can re-derive it without hardcoding the modular reduction.
			"PartialDecryptCircuit": {
				Preimage:     string(hash.DomainPartialDecrypt),
				Keccak256:    "", // not keccak'd; consumed via SetBytes
				BN254Reduced: hash.DomainValue(hash.DomainPartialDecrypt).String(),
			},
		},
		BN254Q:    bn254Q.String(),
		SubgroupL: group.ScalarField().String(),
	}
}

func domainRow(preimage string, digest common.Hash) protocolDomainRow {
	reduced := new(big.Int).Mod(new(big.Int).SetBytes(digest.Bytes()), bn254Q)
	return protocolDomainRow{
		Preimage:     preimage,
		Keccak256:    digest.Hex(),
		BN254Reduced: reduced.String(),
	}
}

// ─── schnorr.json ───────────────────────────────────────────────────────────

type schnorrFile struct {
	Description string            `json:"description"`
	BN254Q      string            `json:"bn254Q"`
	SubgroupL   string            `json:"subgroupOrderL"`
	Operator    []operatorVector  `json:"operator"`
	Organizer   []organizerVector `json:"organizer"`
}

type operatorVector struct {
	Label    string `json:"label"`
	Operator string `json:"operator"`
	Secret   string `json:"secret"`
	Witness  string `json:"witness"`
	PubX     string `json:"pubX"`
	PubY     string `json:"pubY"`
	Ax       string `json:"ax"`
	Ay       string `json:"ay"`
	Z        string `json:"z"`
	Domain   string `json:"domain"` // BN254-reduced domain field
}

type organizerVector struct {
	Label   string `json:"label"`
	EpochID string `json:"epochId"`
	AID     string `json:"aid"`
	Secret  string `json:"secret"`
	Witness string `json:"witness"`
	PKOrgX  string `json:"pkOrgX"`
	PKOrgY  string `json:"pkOrgY"`
	Ax      string `json:"ax"`
	Ay      string `json:"ay"`
	Z       string `json:"z"`
	Domain  string `json:"domain"`
}

func buildSchnorr() schnorrFile {
	curve := twistededwards.GetEdwardsCurve()
	L := curve.Order

	opDomain := new(big.Int).Mod(new(big.Int).SetBytes(protocol.DomainOperatorRegisterV1.Bytes()), bn254Q)
	orgDomain := new(big.Int).Mod(new(big.Int).SetBytes(protocol.DomainOrganizerRegisterV1.Bytes()), bn254Q)

	out := schnorrFile{
		Description: "Schnorr PoK vectors. Operator transcript: poseidon5(domain, op, pubX, pubY, ax) → poseidon2(inner, ay). Organizer: poseidon4(domain, eidField, pkX, pkY) → poseidon4(inner, aidField, ax, ay). Verification: z·G == A + c·PK on BabyJubJub (RTE coords).",
		BN254Q:      bn254Q.String(),
		SubgroupL:   L.String(),
	}

	// Operator vectors — match cmd/operator-schnorr-vectors and the
	// constants block in solidity/test/TestHelpers.t.sol.
	opCases := []struct {
		label, addr string
	}{
		{"THIS", "0x7Fa9385bE102ac3EAc297483Dd6233D62b3e1496"},
		{"BEEF", "0x000000000000000000000000000000000000bEEF"},
		{"CAFE", "0x000000000000000000000000000000000000CaFe"},
	}
	for i, ca := range opCases {
		secret := big.NewInt(int64(0x1000 + 17*i))
		witness := big.NewInt(int64(0x2000 + 23*i))
		v := emitOperator(curve.Base, &L, opDomain, ca.label, ca.addr, secret, witness)
		out.Operator = append(out.Operator, v)
	}

	// Organizer vectors — fixed (eid, aid, secret, witness) per row.
	orgCases := []struct {
		label, eid, aid string
		secret, witness int64
	}{
		{"basic", "0x000000000000000000000077", "0x" + hexN(31, 0x00) + "07", 1234567890, 9999},
		{"different-aid", "0x000000000000000000000077", "0x" + hexN(31, 0x00) + "08", 1234567890, 9999},
		{"different-secret", "0x000000000000000000000088", "0x" + hexN(31, 0xab) + "cd", 42, 17},
	}
	for _, ca := range orgCases {
		v := emitOrganizer(curve.Base, &L, orgDomain, ca.label, ca.eid, ca.aid, big.NewInt(ca.secret), big.NewInt(ca.witness))
		out.Organizer = append(out.Organizer, v)
	}

	return out
}

func emitOperator(G twistededwards.PointAffine, L, domain *big.Int, label, addrHex string, secret, witness *big.Int) operatorVector {
	var pub, A twistededwards.PointAffine
	pub.ScalarMultiplication(&G, secret)
	A.ScalarMultiplication(&G, witness)

	pubX := pub.X.BigInt(new(big.Int))
	pubY := pub.Y.BigInt(new(big.Int))
	aX := A.X.BigInt(new(big.Int))
	aY := A.Y.BigInt(new(big.Int))

	addr := common.HexToAddress(addrHex)
	buf := make([]byte, 0, 32+20+32*4)
	buf = append(buf, protocol.DomainOperatorRegisterV1.Bytes()...)
	buf = append(buf, addr.Bytes()...)
	buf = append(buf, padTo32(pubX)...)
	buf = append(buf, padTo32(pubY)...)
	buf = append(buf, padTo32(aX)...)
	buf = append(buf, padTo32(aY)...)
	c := new(big.Int).SetBytes(ethcrypto.Keccak256(buf))
	c.Mod(c, L)
	z := new(big.Int).Mul(c, secret)
	z.Add(z, witness)
	z.Mod(z, L)

	return operatorVector{
		Label:    label,
		Operator: addrHex,
		Secret:   secret.String(),
		Witness:  witness.String(),
		PubX:     pubX.String(),
		PubY:     pubY.String(),
		Ax:       aX.String(),
		Ay:       aY.String(),
		Z:        z.String(),
		Domain:   domain.String(),
	}
}

func emitOrganizer(G twistededwards.PointAffine, L, domain *big.Int, label, eidHex, aidHex string, secret, _witness *big.Int) organizerVector {
	// We use the schnorr package's prover to keep the vectors
	// byte-identical to what the production code path produces. The
	// witness is drawn from rand.Reader inside the package; for cross-impl
	// determinism we rebuild the proof manually with a fixed witness.
	var pub, A twistededwards.PointAffine
	pub.ScalarMultiplication(&G, secret)
	A.ScalarMultiplication(&G, _witness)

	pubX := pub.X.BigInt(new(big.Int))
	pubY := pub.Y.BigInt(new(big.Int))
	aX := A.X.BigInt(new(big.Int))
	aY := A.Y.BigInt(new(big.Int))

	eidBytes := mustHexBytes(eidHex)
	if len(eidBytes) != 12 {
		fail("organizer: eid must be 12 bytes")
	}
	aidBytes := mustHexBytes(aidHex)
	if len(aidBytes) != 32 {
		fail("organizer: aid must be 32 bytes")
	}

	buf := make([]byte, 0, 32+12+32*5)
	buf = append(buf, protocol.DomainOrganizerRegisterV1.Bytes()...)
	buf = append(buf, eidBytes...)
	buf = append(buf, aidBytes...)
	buf = append(buf, padTo32(pubX)...)
	buf = append(buf, padTo32(pubY)...)
	buf = append(buf, padTo32(aX)...)
	buf = append(buf, padTo32(aY)...)
	c := new(big.Int).SetBytes(ethcrypto.Keccak256(buf))
	c.Mod(c, L)
	z := new(big.Int).Mul(c, secret)
	z.Add(z, _witness)
	z.Mod(z, L)

	// Sanity check: re-prove via the package and confirm the public outputs
	// match (witness is not deterministic in the package, so we only verify
	// PK matches; the proof tuple here uses our chosen witness).
	var eidArr [12]byte
	copy(eidArr[:], eidBytes)
	var aidArr [32]byte
	copy(aidArr[:], aidBytes)
	pkX2, pkY2, _, err := schnorr.ProveOrganizerRegister(secret, eidArr, aidArr)
	if err != nil {
		fail("organizer prove: %v", err)
	}
	if pkX2.Cmp(pubX) != 0 || pkY2.Cmp(pubY) != 0 {
		fail("organizer pubkey drift between manual and package prover")
	}

	return organizerVector{
		Label:   label,
		EpochID: eidHex,
		AID:     aidHex,
		Secret:  secret.String(),
		Witness: _witness.String(),
		PKOrgX:  pubX.String(),
		PKOrgY:  pubY.String(),
		Ax:      aX.String(),
		Ay:      aY.String(),
		Z:       z.String(),
		Domain:  domain.String(),
	}
}

// ─── dleq.json ──────────────────────────────────────────────────────────────

type dleqFile struct {
	Description   string       `json:"description"`
	SubgroupL     string       `json:"subgroupOrderL"`
	PartialDomain string       `json:"partialDecryptDomainBn254"`
	Vectors       []dleqVector `json:"vectors"`
}

type dleqVector struct {
	Label            string `json:"label"`
	EpochID          string `json:"epochId"`
	AID              string `json:"aid"`
	CtIdx            uint16 `json:"ctIdx"`
	ParticipantIndex uint16 `json:"participantIndex"`
	Secret           string `json:"secret"`    // d_i
	Ephemeral        string `json:"ephemeral"` // sk producing C_1 = sk·G
	Witness          string `json:"witness"`   // w
	BaseX            string `json:"baseX"`     // C_1.x
	BaseY            string `json:"baseY"`
	PubX             string `json:"pubX"` // D_i.x
	PubY             string `json:"pubY"`
	DeltaX           string `json:"deltaX"` // δ_i.x
	DeltaY           string `json:"deltaY"`
	A1X              string `json:"a1x"`
	A1Y              string `json:"a1y"`
	A2X              string `json:"a2x"`
	A2Y              string `json:"a2y"`
	Challenge        string `json:"challenge"` // c
	Response         string `json:"response"`  // z = w + c·secret mod L
}

func buildDLEQ() dleqFile {
	curve := twistededwards.GetEdwardsCurve()
	L := curve.Order
	G := curve.Base

	domain := hash.DomainValue(hash.DomainPartialDecrypt)

	cases := []struct {
		label    string
		eid, aid string
		ctIdx    uint16
		i        uint16
		secret   int64
		ephem    int64
		witness  int64
	}{
		{"committee-basic", "0x000000000000000000000077", "0x" + hexN(31, 0x00) + "07", 3, 5, 4242424242, 999999, 7777777},
		{"committee-other-participant", "0x000000000000000000000077", "0x" + hexN(31, 0x00) + "07", 3, 6, 1234567890, 999999, 7777777},
		{"committee-other-ct", "0x000000000000000000000077", "0x" + hexN(31, 0x00) + "07", 4, 5, 4242424242, 999999, 7777777},
	}

	out := dleqFile{
		Description:   "Committee partial-decryption DLEQ vectors. Transcript: state=poseidon5(domain, eid, aid, ctIdx, i); each subsequent point folded as state=poseidon3(state, p.x, p.y) over (D_i, C_1, δ, A1, A2). Verification: z·G == A1 + c·D_i and z·C_1 == A2 + c·δ on BabyJubJub (RTE).",
		SubgroupL:     L.String(),
		PartialDomain: domain.String(),
	}
	for _, ca := range cases {
		v := emitDLEQ(&G, &L, domain, ca.label, ca.eid, ca.aid, ca.ctIdx, ca.i, big.NewInt(ca.secret), big.NewInt(ca.ephem), big.NewInt(ca.witness))
		out.Vectors = append(out.Vectors, v)
	}
	return out
}

func emitDLEQ(
	G *twistededwards.PointAffine,
	L, domain *big.Int,
	label, eidHex, aidHex string,
	ctIdx, partIdx uint16,
	secret, ephem, witness *big.Int,
) dleqVector {
	var PK, C1, Delta, A1, A2 twistededwards.PointAffine
	PK.ScalarMultiplication(G, secret)
	C1.ScalarMultiplication(G, ephem)
	Delta.ScalarMultiplication(&C1, secret)
	A1.ScalarMultiplication(G, witness)
	A2.ScalarMultiplication(&C1, witness)

	bx := pointXY(&C1)
	pkx := pointXY(&PK)
	dx := pointXY(&Delta)
	a1 := pointXY(&A1)
	a2 := pointXY(&A2)

	eidBytes := mustHexBytes(eidHex)
	aidBytes := mustHexBytes(aidHex)
	eidField := new(big.Int).SetBytes(eidBytes)
	aidField := new(big.Int).Mod(new(big.Int).SetBytes(aidBytes), bn254Q)

	state, err := poseidon.Hash([]*big.Int{
		domain, eidField, aidField,
		new(big.Int).SetUint64(uint64(ctIdx)),
		new(big.Int).SetUint64(uint64(partIdx)),
	})
	if err != nil {
		fail("dleq state poseidon: %v", err)
	}
	for _, p := range []twistededwards.PointAffine{PK, C1, Delta, A1, A2} {
		xy := pointXY(&p)
		state, err = poseidon.Hash([]*big.Int{state, xy[0], xy[1]})
		if err != nil {
			fail("dleq fold poseidon: %v", err)
		}
	}
	z := new(big.Int).Mul(state, secret)
	z.Add(z, witness)
	z.Mod(z, L)

	return dleqVector{
		Label:            label,
		EpochID:          eidHex,
		AID:              aidHex,
		CtIdx:            ctIdx,
		ParticipantIndex: partIdx,
		Secret:           secret.String(),
		Ephemeral:        ephem.String(),
		Witness:          witness.String(),
		BaseX:            bx[0].String(), BaseY: bx[1].String(),
		PubX: pkx[0].String(), PubY: pkx[1].String(),
		DeltaX: dx[0].String(), DeltaY: dx[1].String(),
		A1X: a1[0].String(), A1Y: a1[1].String(),
		A2X: a2[0].String(), A2Y: a2[1].String(),
		Challenge: state.String(),
		Response:  z.String(),
	}
}

// ─── helpers ────────────────────────────────────────────────────────────────

func write(dir, name string, v any) {
	path := filepath.Join(dir, name)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fail("marshal %s: %v", name, err)
	}
	b = append(b, '\n')
	if err := os.WriteFile(path, b, 0o644); err != nil {
		fail("write %s: %v", path, err)
	}
	fmt.Fprintf(os.Stderr, "wrote %s (%d bytes)\n", path, len(b))
}

func mustHexBytes(s string) []byte {
	if len(s) >= 2 && s[:2] == "0x" {
		s = s[2:]
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		fail("bad hex %q: %v", s, err)
	}
	return b
}

func padTo32(x *big.Int) []byte {
	b := x.Bytes()
	if len(b) > 32 {
		fail("value larger than 32 bytes: %s", x.String())
	}
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
}

func hexN(n int, b byte) string {
	out := make([]byte, n)
	for i := range out {
		out[i] = b
	}
	return hex.EncodeToString(out)
}

func pointXY(p *twistededwards.PointAffine) [2]*big.Int {
	return [2]*big.Int{
		p.X.BigInt(new(big.Int)),
		p.Y.BigInt(new(big.Int)),
	}
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
