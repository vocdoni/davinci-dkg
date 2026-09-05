package main

import (
	"encoding/hex"
	"math/big"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

// Vectors for the two v4 transcripts (docs/pool-keys-v4.md §3–§7): the compact
// contribution layout and the batched finalization transcript, with every
// intermediate value a Solidity or TypeScript implementation has to
// reproduce — words, keccaks, Poseidon digests, anchors, challenges, BRLC
// commitments and Merkle roots. Inputs are small fixed integers so the files
// stay readable and byte-identical across runs.

type pointJSON struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func pointOf(p types.CurvePoint) pointJSON { return pointJSON{X: p.X.String(), Y: p.Y.String()} }

func decStrings(values []*big.Int) []string {
	out := make([]string, len(values))
	for i, value := range values {
		out[i] = value.String()
	}
	return out
}

func keccakWordsHex(words []*big.Int) string {
	packed := make([]byte, 0, 32*len(words))
	for _, word := range words {
		packed = append(packed, word.FillBytes(make([]byte, 32))...)
	}
	return "0x" + hex.EncodeToString(ethcrypto.Keccak256(packed))
}

func bytes32Hex(value *big.Int) string {
	return "0x" + hex.EncodeToString(value.FillBytes(make([]byte, 32)))
}

// testRecipientKey derives committee member i's BabyJubJub key from a small
// fixed secret so the SDK can rebuild the keys (and decrypt the shares).
func testRecipientKey(i int) (*big.Int, types.NodeKey) {
	secret := big.NewInt(int64(i*100 + 13))
	point := group.NewPoint()
	point.ScalarBaseMult(secret)
	encoded := group.Encode(point)
	return secret, types.NodeKey{PubX: encoded.X, PubY: encoded.Y}
}

// testPolynomials builds one dealer's MaxK polynomials: coefficient m of key j
// is dealer·1000 + (j+1)·10 + m + 1, distinct per dealer and per key.
func testPolynomials(dealer, threshold int) [][]*big.Int {
	sets := make([][]*big.Int, ccommon.MaxK)
	for j := range ccommon.MaxK {
		sets[j] = make([]*big.Int, threshold)
		for m := range threshold {
			sets[j][m] = big.NewInt(int64(dealer*1000 + (j+1)*10 + m + 1))
		}
	}
	return sets
}

func testCommitments(coefficients [][]*big.Int) [][]types.CurvePoint {
	commitments := make([][]types.CurvePoint, len(coefficients))
	for j, set := range coefficients {
		commitments[j] = make([]types.CurvePoint, len(set))
		for m, coefficient := range set {
			point := group.NewPoint()
			point.ScalarBaseMult(coefficient)
			commitments[j][m] = group.Encode(point)
		}
	}
	return commitments
}

func domainJSON(preimage string, digest [32]byte) protocolDomainRow {
	return domainRow(preimage, digest)
}

// ─── contribution_compact.json ──────────────────────────────────────────────

type contributionCompactFile struct {
	Description string                      `json:"description"`
	Domain      protocolDomainRow           `json:"domain"`
	MaxN        int                         `json:"maxN"`
	MaxK        int                         `json:"maxK"`
	Layout      string                      `json:"layout"`
	Anchor      string                      `json:"anchor"`
	Vectors     []contributionCompactVector `json:"vectors"`
}

type contributionOffsets struct {
	Commitments      int `json:"commitments"`
	RecipientIndexes int `json:"recipientIndexes"`
	RecipientKeys    int `json:"recipientKeys"`
	Ephemerals       int `json:"ephemerals"`
	MaskedShares     int `json:"maskedShares"`
	Words            int `json:"words"`
}

type contributionCompactVector struct {
	Label            string   `json:"label"`
	EpochID          string   `json:"epochId"`
	Threshold        int      `json:"threshold"`
	CommitteeSize    int      `json:"committeeSize"`
	ContributorIndex int      `json:"contributorIndex"`
	RecipientSecrets []string `json:"recipientSecrets"`
	EncryptionNonces []string `json:"encryptionNonces"`
	// Coefficients[j][m]: key j, coefficient m (t per key).
	Coefficients [][]string `json:"coefficients"`
	// Shares[j][i]: key j's plaintext share for committee slot i.
	Shares                 [][]string          `json:"shares"`
	Offsets                contributionOffsets `json:"offsets"`
	Transcript             []string            `json:"transcript"`
	CommitteeSnapshotHash  string              `json:"committeeSnapshotKeccak"`
	TranscriptHash         string              `json:"transcriptKeccak"`
	CommitmentsHash        string              `json:"commitmentsHash"`
	EncryptedSharesHash    string              `json:"encryptedSharesHash"`
	ChallengeAnchor        string              `json:"anchor"`
	Challenge              string              `json:"challenge"`
	TranscriptCommitment   string              `json:"transcriptCommitment"`
	PublicInputs           []string            `json:"publicInputs"`
	PublicInputsDescriptor string              `json:"publicInputsOrder"`
}

func buildContributionCompact() contributionCompactFile {
	out := contributionCompactFile{
		Description: "Compact contribution transcript vectors (docs/pool-keys-v4.md §3–§5). transcript is the exact " +
			"L_C = maxK·(2t+n)+5n words submitContribution streams (decimal field elements, no padding); " +
			"offsets are word offsets of each region; committeeSnapshotKeccak is keccak256 over words " +
			"[2Kt, 2Kt+3n) (indexes then public keys), what DKGManager._snapshotCommittee must equal; " +
			"challenge = keccak(eid12 ‖ domain ‖ anchor) mod p; transcriptCommitment = Σ ρ^(q+1)·w[q].",
		Domain: domainJSON(protocol.DomainContributionTranscriptV2Str, protocol.DomainContributionTranscriptV2),
		MaxN:   ccommon.MaxN,
		MaxK:   ccommon.MaxK,
		Layout: "[0,2Kt) A[j][m].x,y key-major; [2Kt,2Kt+n) recipient indexes (slot i = i+1); " +
			"[2Kt+n,2Kt+3n) recipient keys x,y; [2Kt+3n,2Kt+5n) ephemerals x,y; [2Kt+5n,L_C) masked shares ms[j][i] key-major",
		Anchor: "keccak256(commitmentsHash ‖ encryptedSharesHash ‖ keccak256(transcript))",
	}
	cases := []struct {
		label       string
		eid         string
		t, n, index int
	}{
		{"t1-n1-contributor1", "0x000000000000000000000077", 1, 1, 1},
		{"t2-n3-contributor2", "0x000000000000000000000077", 2, 3, 2},
		{"t3-n4-contributor4", "0x0000000000000000000000a1", 3, 4, 4},
		{"t3-n5-contributor5", "0xdeadbeefcafebabe00000003", 3, 5, 5},
	}
	for _, ca := range cases {
		out.Vectors = append(out.Vectors, emitContributionCompact(ca.label, ca.eid, ca.t, ca.n, ca.index))
	}
	return out
}

func emitContributionCompact(label, eidHex string, t, n, index int) contributionCompactVector {
	eid := new(big.Int).SetBytes(mustHexBytes(eidHex))
	secrets := make([]*big.Int, n)
	keys := make([]types.NodeKey, n)
	indexes := make([]uint16, n)
	nonces := make([]*big.Int, n)
	for i := range n {
		secrets[i], keys[i] = testRecipientKey(i)
		indexes[i] = uint16(i + 1)
		nonces[i] = big.NewInt(int64(1000 + i))
	}
	coefficients := testPolynomials(index, t)
	_, pi, err := contribution.BuildWitness(contribution.Assignment{
		RoundHash:        eid,
		Threshold:        uint16(t),
		CommitteeSize:    uint16(n),
		ContributorIndex: uint16(index),
		Coefficients:     coefficients,
		RecipientIndexes: indexes,
		RecipientKeys:    keys,
		EncryptionNonces: nonces,
	})
	if err != nil {
		fail("contribution %s: %v", label, err)
	}
	layout, err := pi.Layout()
	if err != nil {
		fail("contribution %s layout: %v", label, err)
	}
	words, err := pi.TranscriptScalars()
	if err != nil {
		fail("contribution %s transcript: %v", label, err)
	}
	anchor, err := ccommon.ChallengeAnchor(words, pi.CommitmentHash, pi.ShareHash)
	if err != nil {
		fail("contribution %s anchor: %v", label, err)
	}
	start, end := layout.CommitteeRegion()
	coefficientStrings := make([][]string, len(coefficients))
	for j := range coefficients {
		coefficientStrings[j] = decStrings(coefficients[j])
	}
	shares := make([][]string, len(pi.Shares))
	for j := range pi.Shares {
		shares[j] = decStrings(pi.Shares[j])
	}
	return contributionCompactVector{
		Label:            label,
		EpochID:          eidHex,
		Threshold:        t,
		CommitteeSize:    n,
		ContributorIndex: index,
		RecipientSecrets: decStrings(secrets),
		EncryptionNonces: decStrings(nonces),
		Coefficients:     coefficientStrings,
		Shares:           shares,
		Offsets: contributionOffsets{
			Commitments:      0,
			RecipientIndexes: layout.RecipientIndexesStart(),
			RecipientKeys:    layout.RecipientKeysStart(),
			Ephemerals:       layout.EphemeralsStart(),
			MaskedShares:     layout.MaskedSharesStart(),
			Words:            layout.Words(),
		},
		Transcript:             decStrings(words),
		CommitteeSnapshotHash:  keccakWordsHex(words[start:end]),
		TranscriptHash:         keccakWordsHex(words),
		CommitmentsHash:        pi.CommitmentHash.String(),
		EncryptedSharesHash:    pi.ShareHash.String(),
		ChallengeAnchor:        bytes32Hex(anchor),
		Challenge:              pi.Challenge.String(),
		TranscriptCommitment:   pi.TranscriptCommitment.String(),
		PublicInputs:           decStrings(pi.Scalars()),
		PublicInputsDescriptor: "eid, threshold, committeeSize, contributorIndex, commitmentsHash, encryptedSharesHash, challenge, transcriptCommitment",
	}
}

// ─── finalize_transcript.json ───────────────────────────────────────────────

type finalizeFile struct {
	Description     string            `json:"description"`
	Domain          protocolDomainRow `json:"domain"`
	MaxN            int               `json:"maxN"`
	MaxK            int               `json:"maxK"`
	TranscriptWords int               `json:"transcriptWords"`
	KeyWords        int               `json:"keyWords"`
	Offsets         finalizeOffsets   `json:"offsets"`
	Layout          string            `json:"layout"`
	Digest          string            `json:"digest"`
	Anchor          string            `json:"anchor"`
	EmptyLeaf       string            `json:"merkleEmptyLeaf"`
	Vectors         []finalizeVector  `json:"vectors"`
}

type finalizeOffsets struct {
	Indexes int `json:"participantIndexes"`
	Hashes  int `json:"contributionHashes"`
	Keys    int `json:"keys"`
}

type finalizeDealer struct {
	Index           int        `json:"index"`
	Coefficients    [][]string `json:"coefficients"`
	CommitmentsHash string     `json:"commitmentsHash"`
}

type finalizeVector struct {
	Label                  string           `json:"label"`
	EpochID                string           `json:"epochId"`
	Threshold              int              `json:"threshold"`
	CommitteeSize          int              `json:"committeeSize"`
	AcceptedCount          int              `json:"acceptedCount"`
	Dealers                []finalizeDealer `json:"dealers"`
	PoolKeys               []pointJSON      `json:"poolKeys"`
	ShareCommitments       [][]pointJSON    `json:"shareCommitments"`
	Transcript             []string         `json:"transcript"`
	RowsDigest             string           `json:"rowsDigest"`
	KeyDigests             []string         `json:"keyDigests"`
	TranscriptDigest       string           `json:"transcriptDigest"`
	TranscriptHash         string           `json:"transcriptKeccak"`
	ChallengeAnchor        string           `json:"anchor"`
	Challenge              string           `json:"challenge"`
	TranscriptCommitment   string           `json:"transcriptCommitment"`
	PublicInputs           []string         `json:"publicInputs"`
	PublicInputsDescriptor string           `json:"publicInputsOrder"`
	ShareRoots             []string         `json:"shareRoots"`
}

func buildFinalizeTranscript() finalizeFile {
	out := finalizeFile{
		Description: "Batched finalization transcript vectors (docs/pool-keys-v4.md §6–§9). transcript is the fixed " +
			"L_F-word calldata finalizeEpoch streams (decimal field elements); dealers carry the accepted " +
			"contributions' coefficients (key-major, t per key) and the commitmentsHash the contract stores; " +
			"shareCommitments[j][i] is D_j,i for committee member i+1 (n per key here, identity (0,1) beyond in " +
			"the transcript); shareRoots[j] is the keccak Merkle root of key j's leaves; challenge = keccak(eid12 ‖ " +
			"domain ‖ anchor) mod p; transcriptCommitment = Σ ρ^(q+1)·w[q].",
		Domain:          domainJSON(protocol.DomainFinalizeTranscriptV2Str, protocol.DomainFinalizeTranscriptV2),
		MaxN:            ccommon.MaxN,
		MaxK:            ccommon.MaxK,
		TranscriptWords: finalize.TranscriptWords,
		KeyWords:        finalize.KeyWords,
		Offsets: finalizeOffsets{
			Indexes: finalize.IndexesStart,
			Hashes:  finalize.HashesStart,
			Keys:    finalize.KeysStart,
		},
		Layout: "[0,N) participant indexes (0 for rows ≥ a); [N,2N) contribution hashes (0 for rows ≥ a); " +
			"key j at 2N + j·(2+2N): P[j].x, P[j].y, D[j][0].x, D[j][0].y, …, D[j][N−1].x, D[j][N−1].y ((0,1) for i ≥ n)",
		Digest: "R = H(0, I[0..N), h[0..N)); B_j = H(1, j, P[j].x, P[j].y, D[j][0].x, …); " +
			"T = H(2, eid, t, n, a, K, L_F, R, B_0, …, B_(K−1)); H = Poseidon MultiHash (16-input chunks)",
		Anchor:    "keccak256(transcriptDigest ‖ keccak256(transcript))",
		EmptyLeaf: "0x" + hex.EncodeToString(ccommon.EmptyLeaf[:]),
	}
	cases := []struct {
		label    string
		eid      string
		t, n     int
		accepted []uint16
	}{
		{"t2-n3-all", "0x000000000000000000000077", 2, 3, []uint16{1, 2, 3}},
		{"t2-n4-member3-silent", "0x000000000000000000000077", 2, 4, []uint16{1, 2, 4}},
		{"t2-n4-descending-a2", "0x0000000000000000000000a1", 2, 4, []uint16{4, 2}},
		{"t3-n5-all", "0xdeadbeefcafebabe00000003", 3, 5, []uint16{1, 2, 3, 4, 5}},
	}
	for _, ca := range cases {
		out.Vectors = append(out.Vectors, emitFinalize(ca.label, ca.eid, ca.t, ca.n, ca.accepted))
	}
	return out
}

func emitFinalize(label, eidHex string, t, n int, accepted []uint16) finalizeVector {
	eid := new(big.Int).SetBytes(mustHexBytes(eidHex))
	dealers := make([]finalizeDealer, len(accepted))
	commitments := make([][][]types.CurvePoint, len(accepted))
	for d, index := range accepted {
		coefficients := testPolynomials(int(index), t)
		commitments[d] = testCommitments(coefficients)
		dealers[d] = finalizeDealer{Index: int(index)}
		dealers[d].Coefficients = make([][]string, len(coefficients))
		for j := range coefficients {
			dealers[d].Coefficients[j] = decStrings(coefficients[j])
		}
	}
	_, pi, err := finalize.BuildWitness(finalize.Assignment{
		RoundHash:          eid,
		Threshold:          uint16(t),
		CommitteeSize:      uint16(n),
		ParticipantIndexes: accepted,
		Commitments:        commitments,
	})
	if err != nil {
		fail("finalize %s: %v", label, err)
	}
	for d := range accepted {
		dealers[d].CommitmentsHash = pi.ContributionHashes[d].String()
	}
	words, err := pi.TranscriptScalars()
	if err != nil {
		fail("finalize %s transcript: %v", label, err)
	}
	parts, err := finalize.TranscriptDigestParts(pi.RoundHash, pi.Threshold, pi.CommitteeSize, pi.AcceptedCount, words)
	if err != nil {
		fail("finalize %s digest: %v", label, err)
	}
	anchor, err := ccommon.ChallengeAnchor(words, pi.TranscriptDigest)
	if err != nil {
		fail("finalize %s anchor: %v", label, err)
	}
	roots, err := pi.ShareRoots()
	if err != nil {
		fail("finalize %s roots: %v", label, err)
	}
	poolKeys := make([]pointJSON, ccommon.MaxK)
	shares := make([][]pointJSON, ccommon.MaxK)
	rootStrings := make([]string, ccommon.MaxK)
	for j := range ccommon.MaxK {
		poolKeys[j] = pointOf(pi.PoolKeys[j])
		shares[j] = make([]pointJSON, n)
		for i := range n {
			shares[j][i] = pointOf(pi.ShareCommitments[j][i])
		}
		rootStrings[j] = "0x" + hex.EncodeToString(roots[j][:])
	}
	return finalizeVector{
		Label:                  label,
		EpochID:                eidHex,
		Threshold:              t,
		CommitteeSize:          n,
		AcceptedCount:          len(accepted),
		Dealers:                dealers,
		PoolKeys:               poolKeys,
		ShareCommitments:       shares,
		Transcript:             decStrings(words),
		RowsDigest:             parts.Rows.String(),
		KeyDigests:             decStrings(parts.Keys),
		TranscriptDigest:       parts.Digest.String(),
		TranscriptHash:         keccakWordsHex(words),
		ChallengeAnchor:        bytes32Hex(anchor),
		Challenge:              pi.Challenge.String(),
		TranscriptCommitment:   pi.TranscriptCommitment.String(),
		PublicInputs:           decStrings(pi.Scalars()),
		PublicInputsDescriptor: "eid, threshold, committeeSize, acceptedCount, transcriptDigest, challenge, transcriptCommitment",
		ShareRoots:             rootStrings,
	}
}
