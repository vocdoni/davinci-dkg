package contribution

import (
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

func hexKeccak(preimage string) string {
	return hex.EncodeToString(ethcrypto.Keccak256([]byte(preimage)))
}

// layoutAssignment builds a deterministic dealing for a (t, n) committee.
func layoutAssignment(threshold, committee, contributor int) Assignment {
	keys := make([]types.NodeKey, committee)
	indexes := make([]uint16, committee)
	nonces := make([]*big.Int, committee)
	for i := range committee {
		point := group.NewPoint()
		point.ScalarBaseMult(big.NewInt(int64(i*100 + 13)))
		encoded := group.Encode(point)
		keys[i] = types.NodeKey{PubX: encoded.X, PubY: encoded.Y}
		indexes[i] = uint16(i + 1)
		nonces[i] = big.NewInt(int64(1000 + i))
	}
	base := make([]*big.Int, threshold)
	for m := range threshold {
		base[m] = big.NewInt(int64(contributor*10 + m + 1))
	}
	return Assignment{
		RoundHash:        big.NewInt(0x77),
		Threshold:        uint16(threshold),
		CommitteeSize:    uint16(committee),
		ContributorIndex: uint16(contributor),
		Coefficients:     poolCoefficients(base),
		RecipientIndexes: indexes,
		RecipientKeys:    keys,
		EncryptionNonces: nonces,
	}
}

func assertWordIs(c *qt.C, words []*big.Int, offset int, want *big.Int) {
	c.Helper()
	c.Assert(words[offset].Cmp(want), qt.Equals, 0, qt.Commentf("word %d", offset))
}

// The compact transcript is exactly L_C = K·(2t+n) + 5n words in the region
// order of docs/pool-keys-v4.md §3, every word canonical, and the offsets of
// §5 land on the values a recipient or finalizer reads back.
func TestLayoutOffsetsAndLength(t *testing.T) {
	c := qt.New(t)
	modulus := ecc.BN254.ScalarField()

	for _, tc := range []struct{ t, n, contributor int }{
		{1, 1, 1}, {2, 3, 2}, {3, 4, 4}, {5, 7, 6},
	} {
		assignment := layoutAssignment(tc.t, tc.n, tc.contributor)
		_, pi, err := BuildWitness(assignment)
		c.Assert(err, qt.IsNil)

		layout, err := pi.Layout()
		c.Assert(err, qt.IsNil)
		c.Assert(layout, qt.Equals, Layout{Threshold: tc.t, CommitteeSize: tc.n})
		wantWords := MaxKeys*(2*tc.t+tc.n) + 5*tc.n
		c.Assert(layout.Words(), qt.Equals, wantWords)
		c.Assert(layout.Words(), qt.Equals, ccommon.CompactContributionWords(tc.t, tc.n))
		c.Assert(layout.Bytes(), qt.Equals, 32*wantWords)

		words, err := pi.TranscriptScalars()
		c.Assert(err, qt.IsNil)
		c.Assert(words, qt.HasLen, wantWords)
		for q, word := range words {
			c.Assert(word.Sign() >= 0 && word.Cmp(modulus) < 0, qt.IsTrue, qt.Commentf("word %d canonical", q))
		}

		// Region boundaries.
		c.Assert(layout.RecipientIndexesStart(), qt.Equals, 2*MaxKeys*tc.t)
		c.Assert(layout.RecipientKeysStart(), qt.Equals, 2*MaxKeys*tc.t+tc.n)
		c.Assert(layout.EphemeralsStart(), qt.Equals, 2*MaxKeys*tc.t+3*tc.n)
		c.Assert(layout.MaskedSharesStart(), qt.Equals, 2*MaxKeys*tc.t+5*tc.n)
		start, end := layout.CommitteeRegion()
		c.Assert(start, qt.Equals, 2*MaxKeys*tc.t)
		c.Assert(end, qt.Equals, 2*MaxKeys*tc.t+3*tc.n)
		c.Assert(layout.MaskedShareOffset(MaxKeys-1, tc.n-1)+1, qt.Equals, wantWords)

		// Finalizers read A[j][m] at 2(jt+m).
		for j := range MaxKeys {
			for m := range tc.t {
				c.Assert(layout.CommitmentOffset(j, m), qt.Equals, 2*(j*tc.t+m))
				assertWordIs(c, words, layout.CommitmentOffset(j, m), pi.Commitments[j][m].X)
				assertWordIs(c, words, layout.CommitmentOffset(j, m)+1, pi.Commitments[j][m].Y)
			}
		}
		// Recipients read their slot's index, key, ephemeral at 2Kt+3n+2i and
		// masked shares at 2Kt+5n+jn+i.
		for i := range tc.n {
			assertWordIs(c, words, layout.RecipientIndexOffset(i), big.NewInt(int64(i+1)))
			assertWordIs(c, words, layout.RecipientKeyOffset(i), pi.RecipientKeys[i].X)
			assertWordIs(c, words, layout.RecipientKeyOffset(i)+1, pi.RecipientKeys[i].Y)
			c.Assert(layout.EphemeralOffset(i), qt.Equals, 2*MaxKeys*tc.t+3*tc.n+2*i)
			assertWordIs(c, words, layout.EphemeralOffset(i), pi.EncryptedShares[0][i].Ephemeral.X)
			assertWordIs(c, words, layout.EphemeralOffset(i)+1, pi.EncryptedShares[0][i].Ephemeral.Y)
			for j := range MaxKeys {
				c.Assert(layout.MaskedShareOffset(j, i), qt.Equals, 2*MaxKeys*tc.t+5*tc.n+j*tc.n+i)
				assertWordIs(c, words, layout.MaskedShareOffset(j, i), pi.EncryptedShares[j][i].Ciphertext)
			}
		}

		// Decode round-trips, from words and from calldata bytes.
		decoded, err := layout.Decode(words)
		c.Assert(err, qt.IsNil)
		reencoded, err := layout.Encode(*decoded)
		c.Assert(err, qt.IsNil)
		c.Assert(reencoded, qt.HasLen, wantWords)
		for q := range words {
			c.Assert(reencoded[q].Cmp(words[q]), qt.Equals, 0)
		}
		data := make([]byte, 0, layout.Bytes())
		for _, word := range words {
			data = append(data, word.FillBytes(make([]byte, 32))...)
		}
		fromBytes, err := layout.DecodeBytes(data)
		c.Assert(err, qt.IsNil)
		c.Assert(fromBytes.MaskedShares[MaxKeys-1][tc.n-1].Cmp(pi.EncryptedShares[MaxKeys-1][tc.n-1].Ciphertext), qt.Equals, 0)
	}
}

// Decoders reject rather than reduce: a non-canonical word, a wrong length
// and a recipient slot out of committee order are all errors.
func TestLayoutDecodeRejectsMalformedTranscripts(t *testing.T) {
	c := qt.New(t)
	_, pi, err := BuildWitness(layoutAssignment(2, 3, 1))
	c.Assert(err, qt.IsNil)
	layout, err := pi.Layout()
	c.Assert(err, qt.IsNil)
	words, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)

	modulus := ecc.BN254.ScalarField()
	for _, bad := range []*big.Int{
		new(big.Int).Set(modulus),
		new(big.Int).Add(modulus, big.NewInt(1)),
		new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)),
	} {
		altered := append([]*big.Int{}, words...)
		altered[layout.MaskedShareOffset(3, 1)] = bad
		_, err := layout.Decode(altered)
		c.Assert(err, qt.ErrorMatches, ".*not a canonical field element.*")
	}
	_, err = layout.Decode(words[:len(words)-1])
	c.Assert(err, qt.ErrorMatches, ".*got .* words, expected .*")
	_, err = layout.DecodeBytes(make([]byte, layout.Bytes()+32))
	c.Assert(err, qt.ErrorMatches, ".*got .* bytes, expected .*")

	swapped := append([]*big.Int{}, words...)
	swapped[layout.RecipientIndexOffset(0)], swapped[layout.RecipientIndexOffset(1)] = swapped[layout.RecipientIndexOffset(1)], swapped[layout.RecipientIndexOffset(0)]
	_, err = layout.Decode(swapped)
	c.Assert(err, qt.ErrorMatches, ".*recipient slot 0 carries index 2.*")

	// Bounds.
	_, err = NewLayout(0, 3)
	c.Assert(err, qt.Not(qt.IsNil))
	_, err = NewLayout(4, 3)
	c.Assert(err, qt.Not(qt.IsNil))
	_, err = NewLayout(1, MaxRecipients+1)
	c.Assert(err, qt.Not(qt.IsNil))
	_, err = NewLayout(MaxRecipients, MaxRecipients)
	c.Assert(err, qt.IsNil)
}

// The circuit's gated fold over its fixed-size arrays equals the plain BRLC
// over the compact transcript: the witness solves with the compact
// commitment and challenge, and would not with the padded v3.1 commitment
// (same words, exponent advanced through the inactive slots).
func TestFoldMatchesCompactTranscript(t *testing.T) {
	c := qt.New(t)
	field := ecc.BN254.ScalarField()

	witness, pi, err := BuildWitness(layoutAssignment(2, 3, 2))
	c.Assert(err, qt.IsNil)

	words, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	compact, err := ccommon.BRLCNative(pi.Challenge, words...)
	c.Assert(err, qt.IsNil)
	c.Assert(compact.Cmp(pi.TranscriptCommitment), qt.Equals, 0)

	// Padded layout: every fixed slot, inactive ones zero / identity.
	padded := make([]*big.Int, 0, 3*MaxKeys*MaxRecipients+5*MaxRecipients)
	appendPoint := func(x, y frontend.Variable) {
		padded = append(padded, x.(*big.Int), y.(*big.Int))
	}
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			appendPoint(witness.Commitments[j][m].X, witness.Commitments[j][m].Y)
		}
	}
	for i := range MaxRecipients {
		padded = append(padded, witness.RecipientIndexes[i].(*big.Int))
	}
	for i := range MaxRecipients {
		appendPoint(witness.RecipientPubKeys[i].X, witness.RecipientPubKeys[i].Y)
	}
	for i := range MaxRecipients {
		appendPoint(witness.Ephemerals[i].X, witness.Ephemerals[i].Y)
	}
	for j := range MaxKeys {
		for i := range MaxRecipients {
			padded = append(padded, witness.MaskedShares[j][i].(*big.Int))
		}
	}
	paddedCommitment, err := ccommon.BRLCNative(pi.Challenge, padded...)
	c.Assert(err, qt.IsNil)
	c.Assert(paddedCommitment.Cmp(compact), qt.Not(qt.Equals), 0)

	started := time.Now()
	c.Assert(test.IsSolved(&ContributionCircuit{}, witness, field), qt.IsNil)
	t.Logf("test-engine solve of the contribution circuit: %s", time.Since(started))

	wrong := *witness
	wrong.TranscriptCommitment = paddedCommitment
	c.Assert(test.IsSolved(&ContributionCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// Claiming a different committee size changes L_C and every gate: the
	// word count no longer matches.
	wrong = *witness
	wrong.CommitteeSize = big.NewInt(4)
	c.Assert(test.IsSolved(&ContributionCircuit{}, &wrong, field), qt.Not(qt.IsNil))

	// The v4 bounds: t = 0 and contributorIndex = 0 are rejected in-circuit.
	wrong = *witness
	wrong.Threshold = big.NewInt(0)
	c.Assert(test.IsSolved(&ContributionCircuit{}, &wrong, field), qt.Not(qt.IsNil))
	wrong = *witness
	wrong.ContributorIndex = big.NewInt(0)
	c.Assert(test.IsSolved(&ContributionCircuit{}, &wrong, field), qt.Not(qt.IsNil))
}

// The anchor and challenge are derived over the compact words under the v2
// domain; the digests still absorb the padded vectors.
func TestTranscriptAnchorAndChallenge(t *testing.T) {
	c := qt.New(t)
	assignment := layoutAssignment(3, 4, 3)
	_, pi, err := BuildWitness(assignment)
	c.Assert(err, qt.IsNil)

	words, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	anchor, err := ccommon.ChallengeAnchor(words, pi.CommitmentHash, pi.ShareHash)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(assignment.RoundHash, TranscriptDomain, anchor)
	c.Assert(err, qt.IsNil)
	c.Assert(challenge.Cmp(pi.Challenge), qt.Equals, 0)
	c.Assert(TranscriptDomain.Hex(), qt.Equals, "0x"+hexKeccak("davinci-dkg:contribution:v2"))

	scalars := pi.Scalars()
	c.Assert(scalars, qt.HasLen, 8)
	c.Assert(scalars[1].Int64(), qt.Equals, int64(3))
	c.Assert(scalars[2].Int64(), qt.Equals, int64(4))
	c.Assert(scalars[3].Int64(), qt.Equals, int64(3))
	c.Assert(scalars[4].Cmp(pi.CommitmentHash), qt.Equals, 0)
	c.Assert(scalars[5].Cmp(pi.ShareHash), qt.Equals, 0)
	c.Assert(scalars[6].Cmp(pi.Challenge), qt.Equals, 0)
	c.Assert(scalars[7].Cmp(pi.TranscriptCommitment), qt.Equals, 0)
}

// Compile-only: logs the R1CS size of the v4 contribution circuit. Skipped
// under -short; compiling the full-size circuit takes minutes.
func TestCompileContributionConstraintCount(t *testing.T) {
	if testing.Short() {
		t.Skip("compile-only constraint count skipped under -short")
	}
	started := time.Now()
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &ContributionCircuit{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("contribution circuit (MaxN=%d, MaxK=%d): %d constraints, %d public inputs, compiled in %s",
		MaxRecipients, MaxKeys, ccs.GetNbConstraints(), ccs.GetNbPublicVariables()-1, time.Since(started))
}
